// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.14.0
// source: people.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PeopleServiceClient is the client API for PeopleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeopleServiceClient interface {
	CreatePeople(ctx context.Context, in *CreatePeopleReq, opts ...grpc.CallOption) (*CreatePeopleRsp, error)
}

type peopleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeopleServiceClient(cc grpc.ClientConnInterface) PeopleServiceClient {
	return &peopleServiceClient{cc}
}

func (c *peopleServiceClient) CreatePeople(ctx context.Context, in *CreatePeopleReq, opts ...grpc.CallOption) (*CreatePeopleRsp, error) {
	out := new(CreatePeopleRsp)
	err := c.cc.Invoke(ctx, "/PeopleService/CreatePeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeopleServiceServer is the server API for PeopleService service.
// All implementations must embed UnimplementedPeopleServiceServer
// for forward compatibility
type PeopleServiceServer interface {
	CreatePeople(context.Context, *CreatePeopleReq) (*CreatePeopleRsp, error)
	mustEmbedUnimplementedPeopleServiceServer()
}

// UnimplementedPeopleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeopleServiceServer struct {
}

func (UnimplementedPeopleServiceServer) CreatePeople(context.Context, *CreatePeopleReq) (*CreatePeopleRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeople not implemented")
}
func (UnimplementedPeopleServiceServer) mustEmbedUnimplementedPeopleServiceServer() {}

// UnsafePeopleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeopleServiceServer will
// result in compilation errors.
type UnsafePeopleServiceServer interface {
	mustEmbedUnimplementedPeopleServiceServer()
}

func RegisterPeopleServiceServer(s grpc.ServiceRegistrar, srv PeopleServiceServer) {
	s.RegisterService(&PeopleService_ServiceDesc, srv)
}

func _PeopleService_CreatePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePeopleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeopleServiceServer).CreatePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeopleService/CreatePeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeopleServiceServer).CreatePeople(ctx, req.(*CreatePeopleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PeopleService_ServiceDesc is the grpc.ServiceDesc for PeopleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeopleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PeopleService",
	HandlerType: (*PeopleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePeople",
			Handler:    _PeopleService_CreatePeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "people.proto",
}
