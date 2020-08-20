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
	limit10 := server.NewRateLimiter([]string{"DownloadFile", "UploadFile"}, 10)
	limit100 := server.NewRateLimiter([]string{"ListAll"}, 100)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(limit100.Unary()),
		grpc.StreamInterceptor(limit10.Stream()),
	)
	store := file.InMemoryStore{}
	pb.RegisterFileStoreServer(grpcServer, server.New(cfg, store))

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
