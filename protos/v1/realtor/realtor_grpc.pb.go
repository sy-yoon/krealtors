// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/v1/realtor/realtor.proto

package realtor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RealtorServiceClient is the client API for RealtorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealtorServiceClient interface {
	GetReItem(ctx context.Context, in *GetReItemRequest, opts ...grpc.CallOption) (*ReItem, error)
	ListReItemHeader(ctx context.Context, in *ListReItemHeaderRequest, opts ...grpc.CallOption) (*ListReItemHeaderResponse, error)
	CreateReItem(ctx context.Context, in *ReItem, opts ...grpc.CallOption) (*ReItem, error)
	UpdateReItem(ctx context.Context, in *ReItem, opts ...grpc.CallOption) (*ReItem, error)
	DeleteReItem(ctx context.Context, in *DeleteReItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type realtorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRealtorServiceClient(cc grpc.ClientConnInterface) RealtorServiceClient {
	return &realtorServiceClient{cc}
}

func (c *realtorServiceClient) GetReItem(ctx context.Context, in *GetReItemRequest, opts ...grpc.CallOption) (*ReItem, error) {
	out := new(ReItem)
	err := c.cc.Invoke(ctx, "/v1.realtor.RealtorService/GetReItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realtorServiceClient) ListReItemHeader(ctx context.Context, in *ListReItemHeaderRequest, opts ...grpc.CallOption) (*ListReItemHeaderResponse, error) {
	out := new(ListReItemHeaderResponse)
	err := c.cc.Invoke(ctx, "/v1.realtor.RealtorService/ListReItemHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realtorServiceClient) CreateReItem(ctx context.Context, in *ReItem, opts ...grpc.CallOption) (*ReItem, error) {
	out := new(ReItem)
	err := c.cc.Invoke(ctx, "/v1.realtor.RealtorService/CreateReItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realtorServiceClient) UpdateReItem(ctx context.Context, in *ReItem, opts ...grpc.CallOption) (*ReItem, error) {
	out := new(ReItem)
	err := c.cc.Invoke(ctx, "/v1.realtor.RealtorService/UpdateReItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realtorServiceClient) DeleteReItem(ctx context.Context, in *DeleteReItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/v1.realtor.RealtorService/DeleteReItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealtorServiceServer is the server API for RealtorService service.
// All implementations must embed UnimplementedRealtorServiceServer
// for forward compatibility
type RealtorServiceServer interface {
	GetReItem(context.Context, *GetReItemRequest) (*ReItem, error)
	ListReItemHeader(context.Context, *ListReItemHeaderRequest) (*ListReItemHeaderResponse, error)
	CreateReItem(context.Context, *ReItem) (*ReItem, error)
	UpdateReItem(context.Context, *ReItem) (*ReItem, error)
	DeleteReItem(context.Context, *DeleteReItemRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRealtorServiceServer()
}

// UnimplementedRealtorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRealtorServiceServer struct {
}

func (UnimplementedRealtorServiceServer) GetReItem(context.Context, *GetReItemRequest) (*ReItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReItem not implemented")
}
func (UnimplementedRealtorServiceServer) ListReItemHeader(context.Context, *ListReItemHeaderRequest) (*ListReItemHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReItemHeader not implemented")
}
func (UnimplementedRealtorServiceServer) CreateReItem(context.Context, *ReItem) (*ReItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReItem not implemented")
}
func (UnimplementedRealtorServiceServer) UpdateReItem(context.Context, *ReItem) (*ReItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReItem not implemented")
}
func (UnimplementedRealtorServiceServer) DeleteReItem(context.Context, *DeleteReItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReItem not implemented")
}
func (UnimplementedRealtorServiceServer) mustEmbedUnimplementedRealtorServiceServer() {}

// UnsafeRealtorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealtorServiceServer will
// result in compilation errors.
type UnsafeRealtorServiceServer interface {
	mustEmbedUnimplementedRealtorServiceServer()
}

func RegisterRealtorServiceServer(s grpc.ServiceRegistrar, srv RealtorServiceServer) {
	s.RegisterService(&RealtorService_ServiceDesc, srv)
}

func _RealtorService_GetReItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtorServiceServer).GetReItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.realtor.RealtorService/GetReItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtorServiceServer).GetReItem(ctx, req.(*GetReItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealtorService_ListReItemHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReItemHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtorServiceServer).ListReItemHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.realtor.RealtorService/ListReItemHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtorServiceServer).ListReItemHeader(ctx, req.(*ListReItemHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealtorService_CreateReItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtorServiceServer).CreateReItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.realtor.RealtorService/CreateReItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtorServiceServer).CreateReItem(ctx, req.(*ReItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealtorService_UpdateReItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtorServiceServer).UpdateReItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.realtor.RealtorService/UpdateReItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtorServiceServer).UpdateReItem(ctx, req.(*ReItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealtorService_DeleteReItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtorServiceServer).DeleteReItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.realtor.RealtorService/DeleteReItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtorServiceServer).DeleteReItem(ctx, req.(*DeleteReItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealtorService_ServiceDesc is the grpc.ServiceDesc for RealtorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealtorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.realtor.RealtorService",
	HandlerType: (*RealtorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReItem",
			Handler:    _RealtorService_GetReItem_Handler,
		},
		{
			MethodName: "ListReItemHeader",
			Handler:    _RealtorService_ListReItemHeader_Handler,
		},
		{
			MethodName: "CreateReItem",
			Handler:    _RealtorService_CreateReItem_Handler,
		},
		{
			MethodName: "UpdateReItem",
			Handler:    _RealtorService_UpdateReItem_Handler,
		},
		{
			MethodName: "DeleteReItem",
			Handler:    _RealtorService_DeleteReItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/realtor/realtor.proto",
}
