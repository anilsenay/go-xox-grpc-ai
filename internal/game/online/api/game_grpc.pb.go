// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: game.proto

package api

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

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	GetGameBoard(ctx context.Context, in *GameBoardRequest, opts ...grpc.CallOption) (*GameBoardResponse, error)
	ServerMove(ctx context.Context, in *ServerMoveRequest, opts ...grpc.CallOption) (*ServerMoveResponse, error)
	ClientMove(ctx context.Context, in *ClientMoveRequest, opts ...grpc.CallOption) (*ClientMoveResponse, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/online.GameService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) GetGameBoard(ctx context.Context, in *GameBoardRequest, opts ...grpc.CallOption) (*GameBoardResponse, error) {
	out := new(GameBoardResponse)
	err := c.cc.Invoke(ctx, "/online.GameService/GetGameBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ServerMove(ctx context.Context, in *ServerMoveRequest, opts ...grpc.CallOption) (*ServerMoveResponse, error) {
	out := new(ServerMoveResponse)
	err := c.cc.Invoke(ctx, "/online.GameService/ServerMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ClientMove(ctx context.Context, in *ClientMoveRequest, opts ...grpc.CallOption) (*ClientMoveResponse, error) {
	out := new(ClientMoveResponse)
	err := c.cc.Invoke(ctx, "/online.GameService/ClientMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	GetGameBoard(context.Context, *GameBoardRequest) (*GameBoardResponse, error)
	ServerMove(context.Context, *ServerMoveRequest) (*ServerMoveResponse, error)
	ClientMove(context.Context, *ClientMoveRequest) (*ClientMoveResponse, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedGameServiceServer) GetGameBoard(context.Context, *GameBoardRequest) (*GameBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameBoard not implemented")
}
func (UnimplementedGameServiceServer) ServerMove(context.Context, *ServerMoveRequest) (*ServerMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerMove not implemented")
}
func (UnimplementedGameServiceServer) ClientMove(context.Context, *ClientMoveRequest) (*ClientMoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientMove not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online.GameService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_GetGameBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameBoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetGameBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online.GameService/GetGameBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetGameBoard(ctx, req.(*GameBoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ServerMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ServerMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online.GameService/ServerMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ServerMove(ctx, req.(*ServerMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ClientMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ClientMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online.GameService/ClientMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ClientMove(ctx, req.(*ClientMoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "online.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _GameService_Join_Handler,
		},
		{
			MethodName: "GetGameBoard",
			Handler:    _GameService_GetGameBoard_Handler,
		},
		{
			MethodName: "ServerMove",
			Handler:    _GameService_ServerMove_Handler,
		},
		{
			MethodName: "ClientMove",
			Handler:    _GameService_ClientMove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}