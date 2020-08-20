package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-fileserver/api"
	"grpc-fileserver/internal/config"
	"grpc-fileserver/internal/file"
	"grpc-fileserver/internal/server"
	"log"
	"net"
)

func main() {
	cfg := config.FromFlags()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcPort))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	store := file.InMemoryStore{}
	pb.RegisterFileStoreServer(grpcServer, server.New(cfg, store))
	grpcServer.Serve(lis)
}
