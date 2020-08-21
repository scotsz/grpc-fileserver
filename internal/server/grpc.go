package server

import (
	"bytes"
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-fileserver/api"
	"grpc-fileserver/internal/config"
	"grpc-fileserver/internal/file"
	"io"
	"log"
	"time"
)

type server struct {
	store   file.Store
	cfg     *config.Config
	nameGen NameGenerator
}

func (s server) UploadFile(fs pb.FileStore_UploadFileServer) error {
	fileSize := 0
	fileData := bytes.Buffer{}

	for {
		req, err := fs.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		chunk := req.GetChunk()
		fileSize += len(chunk)
		if fileSize > s.cfg.MaxFileSize {
			return status.Errorf(codes.InvalidArgument, "file should not be bigger than %d", s.cfg.MaxFileSize)
		}
		_, err = fileData.Write(chunk)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	fileName := s.nameGen.Generate()
	err := s.store.Save(file.New(fileName, time.Now(), fileData))
	if err != nil {
		return status.Errorf(codes.Internal, "couldn't save file: %v", err)
	}

	err = fs.SendAndClose(&pb.UploadResponse{Name: fileName})
	if err != nil {
		return status.Errorf(codes.Internal, "couldn't send: %v", err)
	}

	log.Printf("saved %s, %d", fileName, fileSize)
	return nil
}

func (s server) DownloadFile(req *pb.DownloadRequest, fs pb.FileStore_DownloadFileServer) error {
	reader, err := s.store.ReadFile(req.Name)

	if err != nil {
		return status.Errorf(codes.NotFound, "couldn't read file: %v", err)
	}
	buffer := make([]byte, s.cfg.ChunkSize)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "couldn't read chunk: %v", err)
		}
		err = fs.Send(&pb.FileData{Chunk: buffer[:n]})
		if err != nil {
			return status.Errorf(codes.Internal, "couldn't send chunk: %v", err)
		}
	}
	log.Println(reader.Close())
	return nil
}

func (s server) ListAll(ctx context.Context, req *pb.ListRequest) (*pb.FileList, error) {
	files := s.store.ListAll()
	list := &pb.FileList{}
	for _, f := range files {
		//list.WriteString(fmt.Sprintf("%s | %v | %v", f.Name, f.CreatedAt, f.UpdatedAt))
		created, err := ptypes.TimestampProto(f.CreatedAt)
		if err != nil {
			created = nil
		}
		updated, err := ptypes.TimestampProto(f.UpdatedAt)
		if err != nil {
			updated = nil
		}
		list.Files = append(list.Files, &pb.FileInfo{
			Name:      f.Name,
			CreatedAt: created,
			UpdatedAt: updated,
		})
	}
	return list, nil
}

func New(cfg *config.Config, store file.Store) pb.FileStoreServer {
	return &server{store, cfg, NewUuidGenerator()}
}
