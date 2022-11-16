package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Event struct {
	ID      string
	Message *chittychat.Message
}

type Client struct {
	time   uint64
	pid    uint32 // use pids because we dont care to generate uuids
	stream chittychat.Chat_ConnectClient
	events chan *Event
}

func NewClient(stream chittychat.Chat_ConnectClient) *Client {
	client := &Client{
		pid:      uint32(os.Getpid()),
		stream:   stream,
		messages: []*chittychat.Message{},
		events:   make(chan *Event, 1),
	}

	// Recv msgs from server in background
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				l.Println("lost connection to server")
				// Quit because there's no option to reestablish connection besides running the
				// process again.
				client.events <- &Event{"quit", nil}
				return
			}

			// Ignore messages from ourself
			if msg.Pid != client.pid {
				client.events <- &Event{"msg", msg}
			}
		}
	}()

	return client
}

// -- start Lamport interface --

func (client *Client) GetTime() uint64 {
	return client.time
}

func (client *Client) GetPid() uint32 {
	return client.pid
}

// -- end Lamport interface --

// Handle user commands typed in the text box.
// We only really need these because the requirements say clients should be able to quit.
func (client *Client) Handle(cmd string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "help":
		client.Log(`[Type messages and press enter to send.](fg:magenta)
Available commands:
/[quit](fg:green) - Gracefully exits the chatroom
/[loss](fg:green) [<percent>](fg:blue) - Set message loss percent
/[clear](fg:green) - Forgets all messages
/[help](fg:green) - Displays this message`)
	case "loss":
		if len(parts) < 2 {
			client.Log(fmt.Sprintf("Current loss is [%d%%](fg:green)", *loss))
			return
		}

		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			client.Log(fmt.Sprintf("['%s' is not a valid integer](fg:red)", parts[1]))
			return
		}

		client.Log(fmt.Sprintf("Changed loss from [%d%%](fg:red) to [%d%%](fg:green)", *loss, amount))
		*loss = amount
	case "clear":
		client.messages = []*chittychat.Message{}
		client.events <- &Event{"clear", nil}
	case "quit":
		err := client.stream.CloseSend()
		if err != nil {
			l.Fatalf("fail to close stream: %v", err)
		}
		client.events <- &Event{"quit", nil}
	default:
		client.Log(fmt.Sprintf("[Unknown command '%s'](fg:red)", cmd))
	}
}

// Send handler for messages. Makes sure we remember to increment time.
func (client *Client) Send(msg string) {

	client.Log(ownMsg)

	// Check if this message was randomly "lost"
	if Lost() {
		return
	}

	client.stream.Send(lamport.MakeMessage(client, sendMsg))
}

// Recv handler for messages. Makes sure we remember to increment time.
func (client *Client) Recv(msg *chittychat.Message) {
	time := lamport.LamportRecv(client, msg)
	l.Printf("ticking client time (%d -> %d)", client.time, time)
	client.time = time
	client.messages = append(client.messages, msg)
}

func client() {
	// Connect to server
	conn, err := grpc.Dial(net.JoinHostPort("localhost", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Fatalf("fail to dial: %v", err)
	}
	// Init client FIXME: Change names
	chatclient := chittychat.NewChatClient(conn)
	stream, err := chatclient.Connect(context.Background())
	if err != nil {
		l.Fatalf("fail to connect: %v", err)
	}

	client := NewClient(stream)

	// Tell server our name so that it can tell everyone else
	client.time = lamport.LamportSend(client)
	client.stream.Send(lamport.MakeMessage(client, *name))

	// Main loop. Handles user input and displaying new messages from the server.
	for {
		select {
		case e := <-client.events:
			l.Printf("client event '%s'", e)
			switch e.ID {
			case "msg":
				client.Recv(e.Message)
			case "quit":
				return
			}
		}
	}
}
