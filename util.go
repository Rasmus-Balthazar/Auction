package main

import (
	"net"

	"github.com/Rasmus-Balthazar/Auction/auctionService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Dial grpc server listening on the given port
func (server *Server) ConnectAutionClient(port string) auctionService.AuctionServiceClient {
	l.Printf("dialing %s", port)

	conn, err := grpc.Dial(net.JoinHostPort("localhost", port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		l.Fatalf("fail to dial: %v", err)
	}

	client := auctionService.NewAuctionServiceClient(conn)
	return client
}
