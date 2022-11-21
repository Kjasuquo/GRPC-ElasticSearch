// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: search.proto

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	// Menu Search
	SearchService(ctx context.Context, in *InSearchRequest, opts ...grpc.CallOption) (*InSearchResponse, error)
	SearchSuggestions(ctx context.Context, in *InSearchSuggestionsRequest, opts ...grpc.CallOption) (*InSearchSuggestionsResponse, error)
	SearchItemSuggestions(ctx context.Context, in *InSearchItemSuggestionsRequest, opts ...grpc.CallOption) (*InSearchItemSuggestionsResponse, error)
	SearchPackageSuggestions(ctx context.Context, in *InSearchPackageSuggestionsRequest, opts ...grpc.CallOption) (*InSearchPackageSuggestionsResponse, error)
	SearchItems(ctx context.Context, in *InSearchItemRequest, opts ...grpc.CallOption) (*InSearchItemResponse, error)
	SearchPackages(ctx context.Context, in *InSearchPackageRequest, opts ...grpc.CallOption) (*InSearchPackageResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchService(ctx context.Context, in *InSearchRequest, opts ...grpc.CallOption) (*InSearchResponse, error) {
	out := new(InSearchResponse)
	err := c.cc.Invoke(ctx, "/search_service.v1.searchService/SearchService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchSuggestions(ctx context.Context, in *InSearchSuggestionsRequest, opts ...grpc.CallOption) (*InSearchSuggestionsResponse, error) {
	out := new(InSearchSuggestionsResponse)
	err := c.cc.Invoke(ctx, "/search_service.v1.searchService/SearchSuggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchItemSuggestions(ctx context.Context, in *InSearchItemSuggestionsRequest, opts ...grpc.CallOption) (*InSearchItemSuggestionsResponse, error) {
	out := new(InSearchItemSuggestionsResponse)
	err := c.cc.Invoke(ctx, "/search_service.v1.searchService/SearchItemSuggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchPackageSuggestions(ctx context.Context, in *InSearchPackageSuggestionsRequest, opts ...grpc.CallOption) (*InSearchPackageSuggestionsResponse, error) {
	out := new(InSearchPackageSuggestionsResponse)
	err := c.cc.Invoke(ctx, "/search_service.v1.searchService/SearchPackageSuggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchItems(ctx context.Context, in *InSearchItemRequest, opts ...grpc.CallOption) (*InSearchItemResponse, error) {
	out := new(InSearchItemResponse)
	err := c.cc.Invoke(ctx, "/search_service.v1.searchService/SearchItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchPackages(ctx context.Context, in *InSearchPackageRequest, opts ...grpc.CallOption) (*InSearchPackageResponse, error) {
	out := new(InSearchPackageResponse)
	err := c.cc.Invoke(ctx, "/search_service.v1.searchService/SearchPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	// Menu Search
	SearchService(context.Context, *InSearchRequest) (*InSearchResponse, error)
	SearchSuggestions(context.Context, *InSearchSuggestionsRequest) (*InSearchSuggestionsResponse, error)
	SearchItemSuggestions(context.Context, *InSearchItemSuggestionsRequest) (*InSearchItemSuggestionsResponse, error)
	SearchPackageSuggestions(context.Context, *InSearchPackageSuggestionsRequest) (*InSearchPackageSuggestionsResponse, error)
	SearchItems(context.Context, *InSearchItemRequest) (*InSearchItemResponse, error)
	SearchPackages(context.Context, *InSearchPackageRequest) (*InSearchPackageResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SearchService(context.Context, *InSearchRequest) (*InSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchService not implemented")
}
func (UnimplementedSearchServiceServer) SearchSuggestions(context.Context, *InSearchSuggestionsRequest) (*InSearchSuggestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSuggestions not implemented")
}
func (UnimplementedSearchServiceServer) SearchItemSuggestions(context.Context, *InSearchItemSuggestionsRequest) (*InSearchItemSuggestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchItemSuggestions not implemented")
}
func (UnimplementedSearchServiceServer) SearchPackageSuggestions(context.Context, *InSearchPackageSuggestionsRequest) (*InSearchPackageSuggestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPackageSuggestions not implemented")
}
func (UnimplementedSearchServiceServer) SearchItems(context.Context, *InSearchItemRequest) (*InSearchItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchItems not implemented")
}
func (UnimplementedSearchServiceServer) SearchPackages(context.Context, *InSearchPackageRequest) (*InSearchPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPackages not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SearchService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_service.v1.searchService/SearchService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchService(ctx, req.(*InSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InSearchSuggestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_service.v1.searchService/SearchSuggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchSuggestions(ctx, req.(*InSearchSuggestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchItemSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InSearchItemSuggestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchItemSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_service.v1.searchService/SearchItemSuggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchItemSuggestions(ctx, req.(*InSearchItemSuggestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchPackageSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InSearchPackageSuggestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchPackageSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_service.v1.searchService/SearchPackageSuggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchPackageSuggestions(ctx, req.(*InSearchPackageSuggestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InSearchItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_service.v1.searchService/SearchItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchItems(ctx, req.(*InSearchItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InSearchPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search_service.v1.searchService/SearchPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchPackages(ctx, req.(*InSearchPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search_service.v1.searchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchService",
			Handler:    _SearchService_SearchService_Handler,
		},
		{
			MethodName: "SearchSuggestions",
			Handler:    _SearchService_SearchSuggestions_Handler,
		},
		{
			MethodName: "SearchItemSuggestions",
			Handler:    _SearchService_SearchItemSuggestions_Handler,
		},
		{
			MethodName: "SearchPackageSuggestions",
			Handler:    _SearchService_SearchPackageSuggestions_Handler,
		},
		{
			MethodName: "SearchItems",
			Handler:    _SearchService_SearchItems_Handler,
		},
		{
			MethodName: "SearchPackages",
			Handler:    _SearchService_SearchPackages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
