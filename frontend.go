package main

import (
	"context"
	"net"
	"os"

	"github.com/Rasmus-Balthazar/Auction/auctionService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
writes to servers, written to from clients.
*/
/*
type ReplicationManager struct {
	pid    uint32
	server auctionService.AuctionService_ConnectServer
} */

type FrontEnd struct {
	auctionService.UnimplementedAuctionServiceServer
	replicationManagers map[uint32]*Server
	clients             map[uint32]*Client
	pid                 uint32
	chBids              chan *auctionService.BidMessage
}

/*

 */

/* func NewReplicationManager(pid uint32, server auctionService.AuctionService_ConnectServer) *ReplicationManager {
return &ReplicationManager{
	pid:    pid,
	server: server,
} */

func frontend() {
	l.Printf("i was called")

	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", *port))
	if err != nil {
		l.Fatalf("fail to listen on port %s: %v", *port, err)
	}
	defer listener.Close()

	conn, err := grpc.Dial(net.JoinHostPort("localhost", "50050"), grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	frontend := &FrontEnd{
		replicationManagers: make(map[uint32]*Server),
		pid:                 uint32(os.Getpid()),
	}

	// The usual gRPC server setup
	go func() {
		grpcServer := grpc.NewServer()
		auctionService.RegisterAuctionServiceServer(grpcServer, frontend)
		l.Printf("server %s is running on port %s", *name, *port)
		if err := grpcServer.Serve(listener); err != nil {
			l.Fatalf("stopped serving: %v", err)
		}
	}()

	// Recv messages and send them to everyone
	for msg := range frontend.chBids {
		l.Printf("send '%s'", msg)

		// Adopt time from server in msg.
		// Uncomment to make sure all messages always appear as the newest message.
		// This might not be what we want though - if a message had latency, we might want to appear
		// in the order from its original context.
		//
		// msg.time = server.time

		for _, rm := range frontend.replicationManagers {
			// Check if this message was randomly "lost"
			l.Print(rm.pid)
			//client.stream.Send(msg)
		}
	}
}

func (frontend *FrontEnd) Bid(ctx context.Context, bid *auctionService.BidMessage) (*auctionService.Outcome, error) {
	l.Printf("something was bid %v", bid)
	highestBidder := &auctionService.Outcome{}
	for _, rm := range frontend.replicationManagers {
		rm.stream.Send(bid)
		highestBidder, _ = rm.Bid(ctx, bid)
	}

	for _, rm := range frontend.replicationManagers {
		rm.stream.Send(bid)
		highestBidder, _ = rm.Bid(ctx, bid)
	}
	return highestBidder, nil
}
