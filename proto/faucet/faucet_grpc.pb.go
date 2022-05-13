// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: faucet/faucet.proto

package faucet

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

// FaucetClient is the client API for Faucet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaucetClient interface {
	RequestFunds(ctx context.Context, in *FundingRequest, opts ...grpc.CallOption) (*FundingResponse, error)
}

type faucetClient struct {
	cc grpc.ClientConnInterface
}

func NewFaucetClient(cc grpc.ClientConnInterface) FaucetClient {
	return &faucetClient{cc}
}

func (c *faucetClient) RequestFunds(ctx context.Context, in *FundingRequest, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/faucet.Faucet/RequestFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaucetServer is the server API for Faucet service.
// All implementations must embed UnimplementedFaucetServer
// for forward compatibility
type FaucetServer interface {
	RequestFunds(context.Context, *FundingRequest) (*FundingResponse, error)
	mustEmbedUnimplementedFaucetServer()
}

// UnimplementedFaucetServer must be embedded to have forward compatible implementations.
type UnimplementedFaucetServer struct {
}

func (UnimplementedFaucetServer) RequestFunds(context.Context, *FundingRequest) (*FundingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFunds not implemented")
}
func (UnimplementedFaucetServer) mustEmbedUnimplementedFaucetServer() {}

// UnsafeFaucetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaucetServer will
// result in compilation errors.
type UnsafeFaucetServer interface {
	mustEmbedUnimplementedFaucetServer()
}

func RegisterFaucetServer(s grpc.ServiceRegistrar, srv FaucetServer) {
	s.RegisterService(&Faucet_ServiceDesc, srv)
}

func _Faucet_RequestFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServer).RequestFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.Faucet/RequestFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServer).RequestFunds(ctx, req.(*FundingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Faucet_ServiceDesc is the grpc.ServiceDesc for Faucet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Faucet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "faucet.Faucet",
	HandlerType: (*FaucetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestFunds",
			Handler:    _Faucet_RequestFunds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faucet/faucet.proto",
}
