// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: post.proto

package blogpost

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	Create(ctx context.Context, in *CreatePost, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *ReadPost, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdatePost, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *ReadPost, opts ...grpc.CallOption) (*Message, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) Create(ctx context.Context, in *CreatePost, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/BlogService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Read(ctx context.Context, in *ReadPost, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/BlogService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Update(ctx context.Context, in *UpdatePost, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/BlogService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Delete(ctx context.Context, in *ReadPost, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/BlogService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	Create(context.Context, *CreatePost) (*Response, error)
	Read(context.Context, *ReadPost) (*Response, error)
	Update(context.Context, *UpdatePost) (*Response, error)
	Delete(context.Context, *ReadPost) (*Message, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) Create(context.Context, *CreatePost) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBlogServiceServer) Read(context.Context, *ReadPost) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedBlogServiceServer) Update(context.Context, *UpdatePost) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBlogServiceServer) Delete(context.Context, *ReadPost) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Create(ctx, req.(*CreatePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Read(ctx, req.(*ReadPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Update(ctx, req.(*UpdatePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Delete(ctx, req.(*ReadPost))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BlogService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _BlogService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BlogService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BlogService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}
