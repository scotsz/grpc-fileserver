package config

import "flag"

type Config struct {
	GrpcPort    int
	MaxFileSize int
	StreamLimit int64
	ListLimit   int64
	ChunkSize   int
}

func FromFlags() *Config {
	port := flag.Int("port", 8080, "grpc port")
	maxFileSize := flag.Int("max_size", 12345, "maximum size of a file")
	streams := flag.Int64("max_streams", 10, "maximum download/upload concurrent requests")
	lists := flag.Int64("max_list", 100, "maximum ListAll concurrent requests")
	chunkSize := flag.Int("chunk_size", 4096, "length of a single chunk")
	flag.Parse()
	return &Config{
		GrpcPort:    *port,
		MaxFileSize: *maxFileSize,
		StreamLimit: *streams,
		ListLimit:   *lists,
		ChunkSize:   *chunkSize,
	}
}
