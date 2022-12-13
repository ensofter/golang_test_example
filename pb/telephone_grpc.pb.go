// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.11
// source: pb/telephone.proto

package pb

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

// TelephoneClient is the client API for Telephone service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelephoneClient interface {
	GetContact(ctx context.Context, in *GetContactRequest, opts ...grpc.CallOption) (*GetContactResponse, error)
	ListContacts(ctx context.Context, in *ListContactsRequest, opts ...grpc.CallOption) (Telephone_ListContactsClient, error)
	RecordCallHistrory(ctx context.Context, opts ...grpc.CallOption) (Telephone_RecordCallHistroryClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Telephone_SendMessageClient, error)
}

type telephoneClient struct {
	cc grpc.ClientConnInterface
}

func NewTelephoneClient(cc grpc.ClientConnInterface) TelephoneClient {
	return &telephoneClient{cc}
}

func (c *telephoneClient) GetContact(ctx context.Context, in *GetContactRequest, opts ...grpc.CallOption) (*GetContactResponse, error) {
	out := new(GetContactResponse)
	err := c.cc.Invoke(ctx, "/pb.Telephone/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telephoneClient) ListContacts(ctx context.Context, in *ListContactsRequest, opts ...grpc.CallOption) (Telephone_ListContactsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Telephone_ServiceDesc.Streams[0], "/pb.Telephone/ListContacts", opts...)
	if err != nil {
		return nil, err
	}
	x := &telephoneListContactsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Telephone_ListContactsClient interface {
	Recv() (*ListContactsResponse, error)
	grpc.ClientStream
}

type telephoneListContactsClient struct {
	grpc.ClientStream
}

func (x *telephoneListContactsClient) Recv() (*ListContactsResponse, error) {
	m := new(ListContactsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *telephoneClient) RecordCallHistrory(ctx context.Context, opts ...grpc.CallOption) (Telephone_RecordCallHistroryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Telephone_ServiceDesc.Streams[1], "/pb.Telephone/RecordCallHistrory", opts...)
	if err != nil {
		return nil, err
	}
	x := &telephoneRecordCallHistroryClient{stream}
	return x, nil
}

type Telephone_RecordCallHistroryClient interface {
	Send(*RecordCallHistoryRequest) error
	CloseAndRecv() (*RecordCallHistoryResponse, error)
	grpc.ClientStream
}

type telephoneRecordCallHistroryClient struct {
	grpc.ClientStream
}

func (x *telephoneRecordCallHistroryClient) Send(m *RecordCallHistoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *telephoneRecordCallHistroryClient) CloseAndRecv() (*RecordCallHistoryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RecordCallHistoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *telephoneClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Telephone_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Telephone_ServiceDesc.Streams[2], "/pb.Telephone/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &telephoneSendMessageClient{stream}
	return x, nil
}

type Telephone_SendMessageClient interface {
	Send(*SendMessageRequest) error
	Recv() (*SendMessageResponse, error)
	grpc.ClientStream
}

type telephoneSendMessageClient struct {
	grpc.ClientStream
}

func (x *telephoneSendMessageClient) Send(m *SendMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *telephoneSendMessageClient) Recv() (*SendMessageResponse, error) {
	m := new(SendMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TelephoneServer is the server API for Telephone service.
// All implementations must embed UnimplementedTelephoneServer
// for forward compatibility
type TelephoneServer interface {
	GetContact(context.Context, *GetContactRequest) (*GetContactResponse, error)
	ListContacts(*ListContactsRequest, Telephone_ListContactsServer) error
	RecordCallHistrory(Telephone_RecordCallHistroryServer) error
	SendMessage(Telephone_SendMessageServer) error
	mustEmbedUnimplementedTelephoneServer()
}

// UnimplementedTelephoneServer must be embedded to have forward compatible implementations.
type UnimplementedTelephoneServer struct {
}

func (UnimplementedTelephoneServer) GetContact(context.Context, *GetContactRequest) (*GetContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedTelephoneServer) ListContacts(*ListContactsRequest, Telephone_ListContactsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListContacts not implemented")
}
func (UnimplementedTelephoneServer) RecordCallHistrory(Telephone_RecordCallHistroryServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordCallHistrory not implemented")
}
func (UnimplementedTelephoneServer) SendMessage(Telephone_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedTelephoneServer) mustEmbedUnimplementedTelephoneServer() {}

// UnsafeTelephoneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelephoneServer will
// result in compilation errors.
type UnsafeTelephoneServer interface {
	mustEmbedUnimplementedTelephoneServer()
}

func RegisterTelephoneServer(s grpc.ServiceRegistrar, srv TelephoneServer) {
	s.RegisterService(&Telephone_ServiceDesc, srv)
}

func _Telephone_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelephoneServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Telephone/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelephoneServer).GetContact(ctx, req.(*GetContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Telephone_ListContacts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListContactsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TelephoneServer).ListContacts(m, &telephoneListContactsServer{stream})
}

type Telephone_ListContactsServer interface {
	Send(*ListContactsResponse) error
	grpc.ServerStream
}

type telephoneListContactsServer struct {
	grpc.ServerStream
}

func (x *telephoneListContactsServer) Send(m *ListContactsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Telephone_RecordCallHistrory_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TelephoneServer).RecordCallHistrory(&telephoneRecordCallHistroryServer{stream})
}

type Telephone_RecordCallHistroryServer interface {
	SendAndClose(*RecordCallHistoryResponse) error
	Recv() (*RecordCallHistoryRequest, error)
	grpc.ServerStream
}

type telephoneRecordCallHistroryServer struct {
	grpc.ServerStream
}

func (x *telephoneRecordCallHistroryServer) SendAndClose(m *RecordCallHistoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *telephoneRecordCallHistroryServer) Recv() (*RecordCallHistoryRequest, error) {
	m := new(RecordCallHistoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Telephone_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TelephoneServer).SendMessage(&telephoneSendMessageServer{stream})
}

type Telephone_SendMessageServer interface {
	Send(*SendMessageResponse) error
	Recv() (*SendMessageRequest, error)
	grpc.ServerStream
}

type telephoneSendMessageServer struct {
	grpc.ServerStream
}

func (x *telephoneSendMessageServer) Send(m *SendMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *telephoneSendMessageServer) Recv() (*SendMessageRequest, error) {
	m := new(SendMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Telephone_ServiceDesc is the grpc.ServiceDesc for Telephone service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Telephone_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Telephone",
	HandlerType: (*TelephoneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContact",
			Handler:    _Telephone_GetContact_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListContacts",
			Handler:       _Telephone_ListContacts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordCallHistrory",
			Handler:       _Telephone_RecordCallHistrory_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _Telephone_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/telephone.proto",
}
