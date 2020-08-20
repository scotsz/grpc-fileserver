// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filestorage.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UploadRequest struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1da47343321951, []int{0}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type UploadResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1da47343321951, []int{1}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DownloadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1da47343321951, []int{2}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

type File struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1da47343321951, []int{3}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1da47343321951, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type FileList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileList) Reset()         { *m = FileList{} }
func (m *FileList) String() string { return proto.CompactTextString(m) }
func (*FileList) ProtoMessage()    {}
func (*FileList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f1da47343321951, []int{5}
}

func (m *FileList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileList.Unmarshal(m, b)
}
func (m *FileList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileList.Marshal(b, m, deterministic)
}
func (m *FileList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileList.Merge(m, src)
}
func (m *FileList) XXX_Size() int {
	return xxx_messageInfo_FileList.Size(m)
}
func (m *FileList) XXX_DiscardUnknown() {
	xxx_messageInfo_FileList.DiscardUnknown(m)
}

var xxx_messageInfo_FileList proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UploadRequest)(nil), "fileserver.v1.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "fileserver.v1.UploadResponse")
	proto.RegisterType((*DownloadRequest)(nil), "fileserver.v1.DownloadRequest")
	proto.RegisterType((*File)(nil), "fileserver.v1.File")
	proto.RegisterType((*ListRequest)(nil), "fileserver.v1.ListRequest")
	proto.RegisterType((*FileList)(nil), "fileserver.v1.FileList")
}

func init() {
	proto.RegisterFile("filestorage.proto", fileDescriptor_2f1da47343321951)
}

var fileDescriptor_2f1da47343321951 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x09, 0xb6, 0xd5, 0x8e, 0x5d, 0xa5, 0xa3, 0xa0, 0x2c, 0x2a, 0x12, 0x14, 0x7a, 0x0a,
	0x7e, 0xdc, 0x05, 0x45, 0xbd, 0xe8, 0x69, 0xc5, 0x8b, 0xb7, 0xa8, 0xa3, 0x06, 0x63, 0x12, 0x93,
	0xb4, 0x3e, 0xae, 0xaf, 0x22, 0xc9, 0xb2, 0xd0, 0x06, 0x7b, 0xcb, 0xfc, 0x3f, 0x26, 0xbf, 0x10,
	0x18, 0xbf, 0x29, 0x4d, 0x21, 0x5a, 0x2f, 0xdf, 0x49, 0x38, 0x6f, 0xa3, 0xc5, 0x2a, 0x4b, 0xe4,
	0x67, 0xe4, 0xc5, 0xec, 0x94, 0x1f, 0x43, 0xf5, 0xe8, 0xb4, 0x95, 0xaf, 0x0d, 0x7d, 0x4f, 0x29,
	0x44, 0xdc, 0x86, 0xfe, 0xcb, 0xc7, 0xd4, 0x7c, 0xee, 0xb2, 0x43, 0x36, 0x19, 0x35, 0xed, 0xc0,
	0x8f, 0x60, 0xa3, 0x8b, 0x05, 0x67, 0x4d, 0x20, 0x44, 0xe8, 0x19, 0xf9, 0x45, 0x39, 0x36, 0x6c,
	0xf2, 0x99, 0x8f, 0x61, 0xf3, 0xda, 0xfe, 0x98, 0xb9, 0x75, 0x7c, 0x00, 0xbd, 0x5b, 0xa5, 0x89,
	0x57, 0xb0, 0x7e, 0xaf, 0x42, 0xec, 0x64, 0x80, 0xb5, 0x24, 0x27, 0xe9, 0xec, 0x97, 0xc1, 0x30,
	0x0d, 0x0f, 0xd1, 0x7a, 0xc2, 0x3b, 0x80, 0xf6, 0xa6, 0x24, 0xe1, 0x9e, 0x58, 0xc0, 0x15, 0x0b,
	0xac, 0xf5, 0xfe, 0x12, 0xb7, 0x45, 0x9c, 0x30, 0xbc, 0x81, 0x51, 0x07, 0x94, 0xd7, 0x1d, 0x14,
	0x85, 0x82, 0xb6, 0xde, 0x2a, 0xfc, 0x54, 0x3a, 0x61, 0x78, 0x01, 0xab, 0x89, 0xf4, 0x52, 0x6b,
	0xac, 0x8b, 0xc4, 0xdc, 0xa3, 0xea, 0x9d, 0x7f, 0xda, 0xc9, 0xbf, 0xea, 0x3f, 0xad, 0x48, 0xa7,
	0x9e, 0x07, 0xf9, 0x07, 0xce, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x2f, 0x63, 0xb8, 0x96,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileStoreClient is the client API for FileStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileStoreClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileStore_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileStore_DownloadFileClient, error)
	ListAll(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*FileList, error)
}

type fileStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStoreClient(cc grpc.ClientConnInterface) FileStoreClient {
	return &fileStoreClient{cc}
}

func (c *fileStoreClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileStore_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileStore_serviceDesc.Streams[0], "/fileserver.v1.FileStore/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStoreUploadFileClient{stream}
	return x, nil
}

type FileStore_UploadFileClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type fileStoreUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileStoreUploadFileClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileStoreUploadFileClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileStoreClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileStore_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileStore_serviceDesc.Streams[1], "/fileserver.v1.FileStore/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStoreDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileStore_DownloadFileClient interface {
	Recv() (*File, error)
	grpc.ClientStream
}

type fileStoreDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileStoreDownloadFileClient) Recv() (*File, error) {
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileStoreClient) ListAll(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*FileList, error) {
	out := new(FileList)
	err := c.cc.Invoke(ctx, "/fileserver.v1.FileStore/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileStoreServer is the server API for FileStore service.
type FileStoreServer interface {
	UploadFile(FileStore_UploadFileServer) error
	DownloadFile(*DownloadRequest, FileStore_DownloadFileServer) error
	ListAll(context.Context, *ListRequest) (*FileList, error)
}

// UnimplementedFileStoreServer can be embedded to have forward compatible implementations.
type UnimplementedFileStoreServer struct {
}

func (*UnimplementedFileStoreServer) UploadFile(srv FileStore_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedFileStoreServer) DownloadFile(req *DownloadRequest, srv FileStore_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (*UnimplementedFileStoreServer) ListAll(ctx context.Context, req *ListRequest) (*FileList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}

func RegisterFileStoreServer(s *grpc.Server, srv FileStoreServer) {
	s.RegisterService(&_FileStore_serviceDesc, srv)
}

func _FileStore_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileStoreServer).UploadFile(&fileStoreUploadFileServer{stream})
}

type FileStore_UploadFileServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type fileStoreUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileStoreUploadFileServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileStoreUploadFileServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileStore_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileStoreServer).DownloadFile(m, &fileStoreDownloadFileServer{stream})
}

type FileStore_DownloadFileServer interface {
	Send(*File) error
	grpc.ServerStream
}

type fileStoreDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileStoreDownloadFileServer) Send(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func _FileStore_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStoreServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileserver.v1.FileStore/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStoreServer).ListAll(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fileserver.v1.FileStore",
	HandlerType: (*FileStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAll",
			Handler:    _FileStore_ListAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileStore_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileStore_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "filestorage.proto",
}
