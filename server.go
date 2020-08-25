package main

import (
	"log"
	"net"

	"github.com/gervasioamy/go-grpc-poc/proto"
	"github.com/gervasioamy/go-grpc-poc/server"

	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := server.NotificationsServer{}
	proto.RegisterNotificationServiceServer(grpcServer, &s)
	log.Printf("Starting server at port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
