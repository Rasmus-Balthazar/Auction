package main

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/Rasmus-Balthazar/Auction/auctionService"
	"google.golang.org/grpc"
)

type ClientConnection struct {
	stream auctionService.AuctionService_ConnectClient
	pid    uint32
	name   string
}

type Server struct {
	auctionService.UnimplementedAuctionServiceServer
	connections map[uint32]*ClientConnection
	chBids      chan *auctionService.BidMessage
	pid         uint32
	result      auctionService.Outcome
}

type Auction struct {
	highestBidder auctionService.Outcome
	//State         auctionService.AuctionState_GOING
	AuctionLocker sync.Mutex
}

var (
	auctionLocker sync.Mutex

	auction = auctionService.Outcome{
		State:    auctionService.AuctionState_GOING,
		Amount:   "",
		BidderId: 0,
	}
)

func server() {
	// We need a listener for grpc
	listener, err := net.Listen("tcp", net.JoinHostPort("localhost", *port))
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

func (a *Auction) endAuction() {
	a.AuctionLocker.Lock()

	defer a.AuctionLocker.Unlock()

	//a.status = AuctionService.AuctionStatus_AUCTION_OVER
	log.Printf("Auction ended - the winning bid was %v by %v\n!", a.Amount, highestBid.MadeBy)
}
