package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-fileserver/api"
	"log"
	"time"
)

type Stresser struct {
	client pb.FileStoreClient
}

func (s *Stresser) DownloadAndUpload() {

}

func (s *Stresser) List() {
	for i := 0; i < 1000; i++ {
		go func() {
			list, err := s.client.ListAll(context.Background(), &pb.ListRequest{})
			if err != nil {
				log.Fatal(err)
			}
			log.Println(list.String())
		}()
	}
	time.Sleep(time.Minute)
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewFileStoreClient(conn)
	stresser := Stresser{client: client}

	stresser.DownloadAndUpload()
	stresser.List()
}
