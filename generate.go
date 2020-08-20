package grpc_fileserver

//go:generate protoc -I api/ api/filestorage.proto --go_out=plugins=grpc:api
