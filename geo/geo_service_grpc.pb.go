// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: geo_service.proto

package geo

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

// GeoServiceClient is the client API for GeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoServiceClient interface {
	AddressSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GeoCode(ctx context.Context, in *GeoCodeRequest, opts ...grpc.CallOption) (*GeoCodeResponse, error)
}

type geoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoServiceClient(cc grpc.ClientConnInterface) GeoServiceClient {
	return &geoServiceClient{cc}
}

func (c *geoServiceClient) AddressSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/geo.GeoService/AddressSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) GeoCode(ctx context.Context, in *GeoCodeRequest, opts ...grpc.CallOption) (*GeoCodeResponse, error) {
	out := new(GeoCodeResponse)
	err := c.cc.Invoke(ctx, "/geo.GeoService/GeoCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoServiceServer is the server API for GeoService service.
// All implementations must embed UnimplementedGeoServiceServer
// for forward compatibility
type GeoServiceServer interface {
	AddressSearch(context.Context, *SearchRequest) (*SearchResponse, error)
	GeoCode(context.Context, *GeoCodeRequest) (*GeoCodeResponse, error)
	mustEmbedUnimplementedGeoServiceServer()
}

// UnimplementedGeoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeoServiceServer struct {
}

func (UnimplementedGeoServiceServer) AddressSearch(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressSearch not implemented")
}
func (UnimplementedGeoServiceServer) GeoCode(context.Context, *GeoCodeRequest) (*GeoCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCode not implemented")
}
func (UnimplementedGeoServiceServer) mustEmbedUnimplementedGeoServiceServer() {}

// UnsafeGeoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoServiceServer will
// result in compilation errors.
type UnsafeGeoServiceServer interface {
	mustEmbedUnimplementedGeoServiceServer()
}

func RegisterGeoServiceServer(s grpc.ServiceRegistrar, srv GeoServiceServer) {
	s.RegisterService(&GeoService_ServiceDesc, srv)
}

func _GeoService_AddressSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).AddressSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.GeoService/AddressSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).AddressSearch(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_GeoCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).GeoCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.GeoService/GeoCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).GeoCode(ctx, req.(*GeoCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoService_ServiceDesc is the grpc.ServiceDesc for GeoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geo.GeoService",
	HandlerType: (*GeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddressSearch",
			Handler:    _GeoService_AddressSearch_Handler,
		},
		{
			MethodName: "GeoCode",
			Handler:    _GeoService_GeoCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geo_service.proto",
}
