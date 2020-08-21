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
	streamLimiter := server.NewRateLimiter([]string{"DownloadFile", "UploadFile"}, cfg.StreamLimit)
	listLimiter := server.NewRateLimiter([]string{"ListAll"}, cfg.ListLimit)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(listLimiter.Unary()),
		grpc.StreamInterceptor(streamLimiter.Stream()),
	)
	store := file.NewSystemStore("/files")
	pb.RegisterFileStoreServer(grpcServer, server.New(cfg, store))

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
