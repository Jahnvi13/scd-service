package main

import (
	"log"
	"net"
	"scd-service/db"
	"scd-service/proto"
	"scd-service/server"

	"google.golang.org/grpc"
)

func main() {
	db.InitDB()

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSCDServiceServer(grpcServer, server.NewServer())

	log.Println("gRPC server listening on :8000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
