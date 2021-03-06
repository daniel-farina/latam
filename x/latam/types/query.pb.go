// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: latam/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("latam/query.proto", fileDescriptor_a163cd3855719498) }

var fileDescriptor_a163cd3855719498 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0x31, 0xae, 0xc2, 0x30,
	0x0c, 0x40, 0xdb, 0xe1, 0x7f, 0xa4, 0x6e, 0x30, 0xa1, 0x0a, 0xe5, 0x00, 0x08, 0x6a, 0x15, 0x6e,
	0xc0, 0xc0, 0xce, 0xca, 0xe6, 0x94, 0x10, 0x22, 0xb5, 0x71, 0x68, 0x52, 0x44, 0x6f, 0xc1, 0xb1,
	0x18, 0x3b, 0x32, 0xa2, 0xf6, 0x22, 0x88, 0x9a, 0xc5, 0x1e, 0xfc, 0xfc, 0xf4, 0x92, 0x69, 0x89,
	0x01, 0x2b, 0xb8, 0x36, 0xaa, 0x6e, 0x33, 0x57, 0x53, 0xa0, 0xd9, 0xfc, 0x84, 0xd6, 0xa8, 0xf2,
	0x8c, 0xb5, 0xb1, 0x98, 0x8d, 0x77, 0x9e, 0xe9, 0x42, 0x13, 0xe9, 0x52, 0x01, 0x3a, 0x03, 0x68,
	0x2d, 0x05, 0x0c, 0x86, 0xac, 0xe7, 0xbf, 0x74, 0x59, 0x90, 0xaf, 0xc8, 0x83, 0x44, 0xaf, 0x58,
	0x08, 0xb7, 0x5c, 0xaa, 0x80, 0x39, 0x38, 0xd4, 0xc6, 0x8e, 0x30, 0xb3, 0x9b, 0x49, 0xf2, 0x77,
	0xf8, 0x12, 0xbb, 0xfd, 0xb3, 0x17, 0x71, 0xd7, 0x8b, 0xf8, 0xdd, 0x8b, 0xf8, 0x31, 0x88, 0xa8,
	0x1b, 0x44, 0xf4, 0x1a, 0x44, 0x74, 0x5c, 0x69, 0x13, 0x2e, 0x8d, 0xcc, 0x0a, 0xaa, 0x80, 0x8b,
	0xd6, 0x9c, 0x04, 0x9c, 0x7c, 0xff, 0xed, 0xd0, 0x3a, 0xe5, 0xe5, 0xff, 0xe8, 0xdd, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x01, 0x3d, 0x69, 0xd0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "danielfarina.latam.latam.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "latam/query.proto",
}
