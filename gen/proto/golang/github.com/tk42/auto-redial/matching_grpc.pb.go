// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/matching.proto

package apiv1

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

// MatchingStoreServiceClient is the client API for MatchingStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchingStoreServiceClient interface {
	GetMatching(ctx context.Context, in *GetMatchingRequest, opts ...grpc.CallOption) (*GetMatchingResponse, error)
	PutMatching(ctx context.Context, in *PutMatchingRequest, opts ...grpc.CallOption) (*PutMatchingResponse, error)
	DeleteMatching(ctx context.Context, in *DeleteMatchingRequest, opts ...grpc.CallOption) (*DeleteMatchingResponse, error)
}

type matchingStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchingStoreServiceClient(cc grpc.ClientConnInterface) MatchingStoreServiceClient {
	return &matchingStoreServiceClient{cc}
}

func (c *matchingStoreServiceClient) GetMatching(ctx context.Context, in *GetMatchingRequest, opts ...grpc.CallOption) (*GetMatchingResponse, error) {
	out := new(GetMatchingResponse)
	err := c.cc.Invoke(ctx, "/api.v1.MatchingStoreService/GetMatching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingStoreServiceClient) PutMatching(ctx context.Context, in *PutMatchingRequest, opts ...grpc.CallOption) (*PutMatchingResponse, error) {
	out := new(PutMatchingResponse)
	err := c.cc.Invoke(ctx, "/api.v1.MatchingStoreService/PutMatching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchingStoreServiceClient) DeleteMatching(ctx context.Context, in *DeleteMatchingRequest, opts ...grpc.CallOption) (*DeleteMatchingResponse, error) {
	out := new(DeleteMatchingResponse)
	err := c.cc.Invoke(ctx, "/api.v1.MatchingStoreService/DeleteMatching", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchingStoreServiceServer is the server API for MatchingStoreService service.
// All implementations should embed UnimplementedMatchingStoreServiceServer
// for forward compatibility
type MatchingStoreServiceServer interface {
	GetMatching(context.Context, *GetMatchingRequest) (*GetMatchingResponse, error)
	PutMatching(context.Context, *PutMatchingRequest) (*PutMatchingResponse, error)
	DeleteMatching(context.Context, *DeleteMatchingRequest) (*DeleteMatchingResponse, error)
}

// UnimplementedMatchingStoreServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMatchingStoreServiceServer struct {
}

func (UnimplementedMatchingStoreServiceServer) GetMatching(context.Context, *GetMatchingRequest) (*GetMatchingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatching not implemented")
}
func (UnimplementedMatchingStoreServiceServer) PutMatching(context.Context, *PutMatchingRequest) (*PutMatchingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutMatching not implemented")
}
func (UnimplementedMatchingStoreServiceServer) DeleteMatching(context.Context, *DeleteMatchingRequest) (*DeleteMatchingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMatching not implemented")
}

// UnsafeMatchingStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchingStoreServiceServer will
// result in compilation errors.
type UnsafeMatchingStoreServiceServer interface {
	mustEmbedUnimplementedMatchingStoreServiceServer()
}

func RegisterMatchingStoreServiceServer(s grpc.ServiceRegistrar, srv MatchingStoreServiceServer) {
	s.RegisterService(&MatchingStoreService_ServiceDesc, srv)
}

func _MatchingStoreService_GetMatching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingStoreServiceServer).GetMatching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.MatchingStoreService/GetMatching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingStoreServiceServer).GetMatching(ctx, req.(*GetMatchingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingStoreService_PutMatching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutMatchingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingStoreServiceServer).PutMatching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.MatchingStoreService/PutMatching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingStoreServiceServer).PutMatching(ctx, req.(*PutMatchingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchingStoreService_DeleteMatching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMatchingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchingStoreServiceServer).DeleteMatching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.MatchingStoreService/DeleteMatching",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchingStoreServiceServer).DeleteMatching(ctx, req.(*DeleteMatchingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchingStoreService_ServiceDesc is the grpc.ServiceDesc for MatchingStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchingStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.MatchingStoreService",
	HandlerType: (*MatchingStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMatching",
			Handler:    _MatchingStoreService_GetMatching_Handler,
		},
		{
			MethodName: "PutMatching",
			Handler:    _MatchingStoreService_PutMatching_Handler,
		},
		{
			MethodName: "DeleteMatching",
			Handler:    _MatchingStoreService_DeleteMatching_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/matching.proto",
}
