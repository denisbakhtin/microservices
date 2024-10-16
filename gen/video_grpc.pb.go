// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: video.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MetadataService_GetMetadata_FullMethodName = "/MetadataService/GetMetadata"
	MetadataService_PutMetadata_FullMethodName = "/MetadataService/PutMetadata"
)

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataServiceClient interface {
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error)
	PutMetadata(ctx context.Context, in *PutMetadataRequest, opts ...grpc.CallOption) (*PutMetadataResponse, error)
}

type metadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataServiceClient(cc grpc.ClientConnInterface) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*GetMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetadataResponse)
	err := c.cc.Invoke(ctx, MetadataService_GetMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) PutMetadata(ctx context.Context, in *PutMetadataRequest, opts ...grpc.CallOption) (*PutMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutMetadataResponse)
	err := c.cc.Invoke(ctx, MetadataService_PutMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
// All implementations must embed UnimplementedMetadataServiceServer
// for forward compatibility.
type MetadataServiceServer interface {
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)
	PutMetadata(context.Context, *PutMetadataRequest) (*PutMetadataResponse, error)
	mustEmbedUnimplementedMetadataServiceServer()
}

// UnimplementedMetadataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetadataServiceServer struct{}

func (UnimplementedMetadataServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedMetadataServiceServer) PutMetadata(context.Context, *PutMetadataRequest) (*PutMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutMetadata not implemented")
}
func (UnimplementedMetadataServiceServer) mustEmbedUnimplementedMetadataServiceServer() {}
func (UnimplementedMetadataServiceServer) testEmbeddedByValue()                         {}

// UnsafeMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServiceServer will
// result in compilation errors.
type UnsafeMetadataServiceServer interface {
	mustEmbedUnimplementedMetadataServiceServer()
}

func RegisterMetadataServiceServer(s grpc.ServiceRegistrar, srv MetadataServiceServer) {
	// If the following call pancis, it indicates UnimplementedMetadataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MetadataService_ServiceDesc, srv)
}

func _MetadataService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetadataService_GetMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_PutMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).PutMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetadataService_PutMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).PutMetadata(ctx, req.(*PutMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetadataService_ServiceDesc is the grpc.ServiceDesc for MetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetadata",
			Handler:    _MetadataService_GetMetadata_Handler,
		},
		{
			MethodName: "PutMetadata",
			Handler:    _MetadataService_PutMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}

const (
	RatingService_GetAggregatedRating_FullMethodName = "/RatingService/GetAggregatedRating"
	RatingService_PutRating_FullMethodName           = "/RatingService/PutRating"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetAggregatedRating(ctx context.Context, in *GetAggregatedRatingRequest, opts ...grpc.CallOption) (*GetAggregatedRatingResponse, error)
	PutRating(ctx context.Context, in *PutRatingRequest, opts ...grpc.CallOption) (*PutRatingResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetAggregatedRating(ctx context.Context, in *GetAggregatedRatingRequest, opts ...grpc.CallOption) (*GetAggregatedRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAggregatedRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_GetAggregatedRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) PutRating(ctx context.Context, in *PutRatingRequest, opts ...grpc.CallOption) (*PutRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_PutRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility.
type RatingServiceServer interface {
	GetAggregatedRating(context.Context, *GetAggregatedRatingRequest) (*GetAggregatedRatingResponse, error)
	PutRating(context.Context, *PutRatingRequest) (*PutRatingResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRatingServiceServer struct{}

func (UnimplementedRatingServiceServer) GetAggregatedRating(context.Context, *GetAggregatedRatingRequest) (*GetAggregatedRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAggregatedRating not implemented")
}
func (UnimplementedRatingServiceServer) PutRating(context.Context, *PutRatingRequest) (*PutRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRating not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}
func (UnimplementedRatingServiceServer) testEmbeddedByValue()                       {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	// If the following call pancis, it indicates UnimplementedRatingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_GetAggregatedRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAggregatedRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAggregatedRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetAggregatedRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAggregatedRating(ctx, req.(*GetAggregatedRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_PutRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).PutRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_PutRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).PutRating(ctx, req.(*PutRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAggregatedRating",
			Handler:    _RatingService_GetAggregatedRating_Handler,
		},
		{
			MethodName: "PutRating",
			Handler:    _RatingService_PutRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}

const (
	VideoService_GetVideoDetails_FullMethodName = "/VideoService/GetVideoDetails"
)

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	GetVideoDetails(ctx context.Context, in *GetVideoDetailsRequest, opts ...grpc.CallOption) (*GetVideoDetailsResponse, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideoDetails(ctx context.Context, in *GetVideoDetailsRequest, opts ...grpc.CallOption) (*GetVideoDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVideoDetailsResponse)
	err := c.cc.Invoke(ctx, VideoService_GetVideoDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility.
type VideoServiceServer interface {
	GetVideoDetails(context.Context, *GetVideoDetailsRequest) (*GetVideoDetailsResponse, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVideoServiceServer struct{}

func (UnimplementedVideoServiceServer) GetVideoDetails(context.Context, *GetVideoDetailsRequest) (*GetVideoDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoDetails not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}
func (UnimplementedVideoServiceServer) testEmbeddedByValue()                      {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	// If the following call pancis, it indicates UnimplementedVideoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_GetVideoDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideoDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_GetVideoDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideoDetails(ctx, req.(*GetVideoDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideoDetails",
			Handler:    _VideoService_GetVideoDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
