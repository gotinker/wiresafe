// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: wiresafe.proto

package wiresafe

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

// WireSafeClient is the client API for WireSafe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WireSafeClient interface {
	InitKey(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*InitReply, error)
	Encrypt(ctx context.Context, in *EncryptReq, opts ...grpc.CallOption) (*EncryptResp, error)
	Decrypt(ctx context.Context, in *DecryptReq, opts ...grpc.CallOption) (*DecryptResp, error)
}

type wireSafeClient struct {
	cc grpc.ClientConnInterface
}

func NewWireSafeClient(cc grpc.ClientConnInterface) WireSafeClient {
	return &wireSafeClient{cc}
}

func (c *wireSafeClient) InitKey(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*InitReply, error) {
	out := new(InitReply)
	err := c.cc.Invoke(ctx, "/wiresafe.WireSafe/InitKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireSafeClient) Encrypt(ctx context.Context, in *EncryptReq, opts ...grpc.CallOption) (*EncryptResp, error) {
	out := new(EncryptResp)
	err := c.cc.Invoke(ctx, "/wiresafe.WireSafe/encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireSafeClient) Decrypt(ctx context.Context, in *DecryptReq, opts ...grpc.CallOption) (*DecryptResp, error) {
	out := new(DecryptResp)
	err := c.cc.Invoke(ctx, "/wiresafe.WireSafe/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WireSafeServer is the server API for WireSafe service.
// All implementations must embed UnimplementedWireSafeServer
// for forward compatibility
type WireSafeServer interface {
	InitKey(context.Context, *KeySpec) (*InitReply, error)
	Encrypt(context.Context, *EncryptReq) (*EncryptResp, error)
	Decrypt(context.Context, *DecryptReq) (*DecryptResp, error)
	mustEmbedUnimplementedWireSafeServer()
}

// UnimplementedWireSafeServer must be embedded to have forward compatible implementations.
type UnimplementedWireSafeServer struct {
}

func (UnimplementedWireSafeServer) InitKey(context.Context, *KeySpec) (*InitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitKey not implemented")
}
func (UnimplementedWireSafeServer) Encrypt(context.Context, *EncryptReq) (*EncryptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedWireSafeServer) Decrypt(context.Context, *DecryptReq) (*DecryptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedWireSafeServer) mustEmbedUnimplementedWireSafeServer() {}

// UnsafeWireSafeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WireSafeServer will
// result in compilation errors.
type UnsafeWireSafeServer interface {
	mustEmbedUnimplementedWireSafeServer()
}

func RegisterWireSafeServer(s grpc.ServiceRegistrar, srv WireSafeServer) {
	s.RegisterService(&WireSafe_ServiceDesc, srv)
}

func _WireSafe_InitKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireSafeServer).InitKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiresafe.WireSafe/InitKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireSafeServer).InitKey(ctx, req.(*KeySpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WireSafe_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireSafeServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiresafe.WireSafe/encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireSafeServer).Encrypt(ctx, req.(*EncryptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WireSafe_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireSafeServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wiresafe.WireSafe/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireSafeServer).Decrypt(ctx, req.(*DecryptReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WireSafe_ServiceDesc is the grpc.ServiceDesc for WireSafe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WireSafe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiresafe.WireSafe",
	HandlerType: (*WireSafeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitKey",
			Handler:    _WireSafe_InitKey_Handler,
		},
		{
			MethodName: "encrypt",
			Handler:    _WireSafe_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _WireSafe_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wiresafe.proto",
}