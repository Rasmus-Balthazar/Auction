package main

import (
	"log"
	"net"
	"os"

	"github.com/Rasmus-Balthazar/Auction/auctionService"
	"google.golang.org/grpc"
)

/*
writes to servers, written to from clients.
*/

type ReplicationManager struct {
	pid    uint32
	server auctionService.AuctionService_ConnectServer
}

type FrontEnd struct {
	replicationManagers map[uint32]*ReplicationManager
}

func NewReplicationManager(pid uint32, server auctionService.AuctionService_ConnectServer) *ReplicationManager {
	return &ReplicationManager{
		pid:    pid,
		server: server,
	}
}

func frontend() {

	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", "8080"))
	if err != nil {
		l.Fatalf("fail to listen on port %s: %v", *port, err)
	}
	defer listener.Close()

	server := &Server{
		connections: make(map[uint32]*ClientConnection),
		pid:         uint32(os.Getpid()),
	}

	// The usual gRPC server setup
	go func() {
		grpcServer := grpc.NewServer()
		auctionService.RegisterAuctionServiceServer(grpcServer, server)
		l.Printf("server %s is running on port %s", *name, *port)
		if err := grpcServer.Serve(listener); err != nil {
			l.Fatalf("stopped serving: %v", err)
		}
	}()

	// Recv messages and send them to everyone
	for msg := range server.chBids {
		l.Printf("send '%s'", msg)

		// Adopt time from server in msg.
		// Uncomment to make sure all messages always appear as the newest message.
		// This might not be what we want though - if a message had latency, we might want to appear
		// in the order from its original context.
		//
		// msg.time = server.time

		for _, client := range server.connections {
			// Check if this message was randomly "lost"
			log.Print(client.name)
			//client.stream.Send(msg)
		}
	}
}

/* func (server *Server) Multicast(req ra.Request) chan ra.Reply {
	l.Printf("multicasting %v", req)
	ra.Send(server)
	req.SetTime(server.GetTime())

	for _, p := range server.peers {
		p.client.SendMessage(context.Background(), ToMessage(req))
	}

	return server.replyQueue
} */
