package config

import "flag"

type Config struct {
	GrpcPort    int
	MaxFileSize int
}

func FromFlags() *Config {
	port := flag.Int("port", 8080, "grpc port")
	maxFileSize := flag.Int("max_size", 12345, "maximum size of a file")
	flag.Parse()
	return &Config{GrpcPort: *port, MaxFileSize: *maxFileSize}
}
