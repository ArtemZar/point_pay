// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: accounts.proto

package proto

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

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	CreateAccount(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Accounts_GetAccountsClient, error)
	GenerateAddress(ctx context.Context, in *NewWalletRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	Deposit(ctx context.Context, in *ChangeBalanceRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	Withdrawal(ctx context.Context, in *ChangeBalanceRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) CreateAccount(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/transport.Accounts/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetAccounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Accounts_GetAccountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Accounts_ServiceDesc.Streams[0], "/transport.Accounts/GetAccounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountsGetAccountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Accounts_GetAccountsClient interface {
	Recv() (*AccountResponse, error)
	grpc.ClientStream
}

type accountsGetAccountsClient struct {
	grpc.ClientStream
}

func (x *accountsGetAccountsClient) Recv() (*AccountResponse, error) {
	m := new(AccountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *accountsClient) GenerateAddress(ctx context.Context, in *NewWalletRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/transport.Accounts/GenerateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Deposit(ctx context.Context, in *ChangeBalanceRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/transport.Accounts/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Withdrawal(ctx context.Context, in *ChangeBalanceRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/transport.Accounts/Withdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	CreateAccount(context.Context, *NewUserRequest) (*AccountResponse, error)
	GetAccounts(*emptypb.Empty, Accounts_GetAccountsServer) error
	GenerateAddress(context.Context, *NewWalletRequest) (*AccountResponse, error)
	Deposit(context.Context, *ChangeBalanceRequest) (*AccountResponse, error)
	Withdrawal(context.Context, *ChangeBalanceRequest) (*AccountResponse, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) CreateAccount(context.Context, *NewUserRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountsServer) GetAccounts(*emptypb.Empty, Accounts_GetAccountsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedAccountsServer) GenerateAddress(context.Context, *NewWalletRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAddress not implemented")
}
func (UnimplementedAccountsServer) Deposit(context.Context, *ChangeBalanceRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedAccountsServer) Withdrawal(context.Context, *ChangeBalanceRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdrawal not implemented")
}
func (UnimplementedAccountsServer) mustEmbedUnimplementedAccountsServer() {}

// UnsafeAccountsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServer will
// result in compilation errors.
type UnsafeAccountsServer interface {
	mustEmbedUnimplementedAccountsServer()
}

func RegisterAccountsServer(s grpc.ServiceRegistrar, srv AccountsServer) {
	s.RegisterService(&Accounts_ServiceDesc, srv)
}

func _Accounts_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.Accounts/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).CreateAccount(ctx, req.(*NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetAccounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountsServer).GetAccounts(m, &accountsGetAccountsServer{stream})
}

type Accounts_GetAccountsServer interface {
	Send(*AccountResponse) error
	grpc.ServerStream
}

type accountsGetAccountsServer struct {
	grpc.ServerStream
}

func (x *accountsGetAccountsServer) Send(m *AccountResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Accounts_GenerateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GenerateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.Accounts/GenerateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GenerateAddress(ctx, req.(*NewWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.Accounts/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Deposit(ctx, req.(*ChangeBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Withdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Withdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transport.Accounts/Withdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Withdrawal(ctx, req.(*ChangeBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounts_ServiceDesc is the grpc.ServiceDesc for Accounts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transport.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Accounts_CreateAccount_Handler,
		},
		{
			MethodName: "GenerateAddress",
			Handler:    _Accounts_GenerateAddress_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Accounts_Deposit_Handler,
		},
		{
			MethodName: "Withdrawal",
			Handler:    _Accounts_Withdrawal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAccounts",
			Handler:       _Accounts_GetAccounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "accounts.proto",
}
