// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: grpc.proto

package gRPCEncryptedStorage

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

// PasswordsStorageClient is the client API for PasswordsStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PasswordsStorageClient interface {
	GetStorage(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Storage, error)
	AddUpdateStorageRecord(ctx context.Context, in *Storage, opts ...grpc.CallOption) (*AnswerResultCode, error)
	DeleteRecord(ctx context.Context, in *Storage, opts ...grpc.CallOption) (*AnswerResultCode, error)
}

type passwordsStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewPasswordsStorageClient(cc grpc.ClientConnInterface) PasswordsStorageClient {
	return &passwordsStorageClient{cc}
}

func (c *passwordsStorageClient) GetStorage(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Storage, error) {
	out := new(Storage)
	err := c.cc.Invoke(ctx, "/DevEncryptionRPC.PasswordsStorage/GetStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsStorageClient) AddUpdateStorageRecord(ctx context.Context, in *Storage, opts ...grpc.CallOption) (*AnswerResultCode, error) {
	out := new(AnswerResultCode)
	err := c.cc.Invoke(ctx, "/DevEncryptionRPC.PasswordsStorage/AddUpdateStorageRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passwordsStorageClient) DeleteRecord(ctx context.Context, in *Storage, opts ...grpc.CallOption) (*AnswerResultCode, error) {
	out := new(AnswerResultCode)
	err := c.cc.Invoke(ctx, "/DevEncryptionRPC.PasswordsStorage/DeleteRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasswordsStorageServer is the server API for PasswordsStorage service.
// All implementations must embed UnimplementedPasswordsStorageServer
// for forward compatibility
type PasswordsStorageServer interface {
	GetStorage(context.Context, *Token) (*Storage, error)
	AddUpdateStorageRecord(context.Context, *Storage) (*AnswerResultCode, error)
	DeleteRecord(context.Context, *Storage) (*AnswerResultCode, error)
	mustEmbedUnimplementedPasswordsStorageServer()
}

// UnimplementedPasswordsStorageServer must be embedded to have forward compatible implementations.
type UnimplementedPasswordsStorageServer struct {
}

func (UnimplementedPasswordsStorageServer) GetStorage(context.Context, *Token) (*Storage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (UnimplementedPasswordsStorageServer) AddUpdateStorageRecord(context.Context, *Storage) (*AnswerResultCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUpdateStorageRecord not implemented")
}
func (UnimplementedPasswordsStorageServer) DeleteRecord(context.Context, *Storage) (*AnswerResultCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedPasswordsStorageServer) mustEmbedUnimplementedPasswordsStorageServer() {}

// UnsafePasswordsStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PasswordsStorageServer will
// result in compilation errors.
type UnsafePasswordsStorageServer interface {
	mustEmbedUnimplementedPasswordsStorageServer()
}

func RegisterPasswordsStorageServer(s grpc.ServiceRegistrar, srv PasswordsStorageServer) {
	s.RegisterService(&PasswordsStorage_ServiceDesc, srv)
}

func _PasswordsStorage_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsStorageServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DevEncryptionRPC.PasswordsStorage/GetStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsStorageServer).GetStorage(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasswordsStorage_AddUpdateStorageRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Storage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsStorageServer).AddUpdateStorageRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DevEncryptionRPC.PasswordsStorage/AddUpdateStorageRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsStorageServer).AddUpdateStorageRecord(ctx, req.(*Storage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasswordsStorage_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Storage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasswordsStorageServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DevEncryptionRPC.PasswordsStorage/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasswordsStorageServer).DeleteRecord(ctx, req.(*Storage))
	}
	return interceptor(ctx, in, info, handler)
}

// PasswordsStorage_ServiceDesc is the grpc.ServiceDesc for PasswordsStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PasswordsStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DevEncryptionRPC.PasswordsStorage",
	HandlerType: (*PasswordsStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStorage",
			Handler:    _PasswordsStorage_GetStorage_Handler,
		},
		{
			MethodName: "AddUpdateStorageRecord",
			Handler:    _PasswordsStorage_AddUpdateStorageRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _PasswordsStorage_DeleteRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
