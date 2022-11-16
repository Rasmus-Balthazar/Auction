package main

import (
	"bufio"
	"context"
	"net"
	"os"
	"strings"

	"github.com/Rasmus-Balthazar/Auction/auctionService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	pid    uint32 // use pids because we dont care to generate uuids
	stream auctionService.AuctionService_ConnectClient
}

func NewClient(stream auctionService.AuctionService_ConnectClient) *Client {
	client := &Client{
		pid:    uint32(os.Getpid()),
		stream: stream,
	}

	// Recv msgs from server in background
	go func() {
		for {
			bid, err := stream.Recv()
			if err != nil {
				l.Println(bid.BidAmount)
				l.Println("lost connection to server")
				// Quit because there's no option to reestablish connection besides running the
				// process again.
				return
			}
		}
	}()

	return client
}

//see hw 3 and 4 for interface usage.

// Handle user commands typed in the text box.
// We only really need these because the requirements say clients should be able to quit.

// Send handler for messages. Makes sure we remember to increment time.
func (client *Client) Send(msg string) {
	// Check if this message was randomly "lost"
	if Lost() {
		return
	}

}

// Recv handler for messages. Makes sure we remember to increment time.
func (client *Client) Recv(msg *auctionService.Outcome) {

}

func client() {
	// Connect to server
	conn, err := grpc.Dial(net.JoinHostPort("localhost", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Fatalf("fail to dial: %v", err)
	}

	auctionClient := auctionService.NewAuctionServiceClient(conn)
	stream, err := auctionClient.Connect(context.Background())
	if err != nil {
		l.Fatalf("fail to connect: %v", err)
	}

	client := NewClient(stream)

	client.stream.Send(&auctionService.BidMessage{BidderId: 1, BidAmount: 1})
	// Tell server our name so that it can tell everyone else

	// Main loop. Handles user input and displaying new messages from the server.
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		switch input[0] {
		case "bid":
			//frondend.bid
			/* select {
			/* case e := <-client.events:
			l.Printf("client event '%s'", e)
			switch e.ID {
			case "msg":
				client.Recv(e.Message)
			case "quit":
				return
			}
			} */
		case "quit":
			return
		}
	}
}

// make bid multicast
func (client *Client) Bid(bid *auctionService.BidMessage) {
	client.stream.Send(bid)
}
