// Code generated by protoc-gen-go.
// source: byzantine.proto
// DO NOT EDIT!

/*
Package byzantine is a generated protocol buffer package.

It is generated from these files:
	byzantine.proto

It has these top-level messages:
	Publication
	Publisher
	Subscriber
	PubResponse
	SubRequest
	ChainMAC
	ChainResponse
	EchoResponse
	ReadyResponse
*/
package byzantine

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / Publication is the message sent out to the brokers, who pass it along to the subscribers.
type Publication struct {
	PubType       uint32      `protobuf:"varint,1,opt,name=PubType" json:"PubType,omitempty"`
	PublisherID   uint64      `protobuf:"varint,2,opt,name=PublisherID" json:"PublisherID,omitempty"`
	PublicationID int64       `protobuf:"zigzag64,3,opt,name=PublicationID" json:"PublicationID,omitempty"`
	TopicID       uint64      `protobuf:"varint,4,opt,name=TopicID" json:"TopicID,omitempty"`
	BrokerID      uint64      `protobuf:"varint,5,opt,name=BrokerID" json:"BrokerID,omitempty"`
	Contents      [][]byte    `protobuf:"bytes,6,rep,name=Contents,proto3" json:"Contents,omitempty"`
	MAC           []byte      `protobuf:"bytes,7,opt,name=MAC,proto3" json:"MAC,omitempty"`
	ChainMACs     []*ChainMAC `protobuf:"bytes,8,rep,name=ChainMACs" json:"ChainMACs,omitempty"`
}

func (m *Publication) Reset()                    { *m = Publication{} }
func (m *Publication) String() string            { return proto.CompactTextString(m) }
func (*Publication) ProtoMessage()               {}
func (*Publication) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Publication) GetPubType() uint32 {
	if m != nil {
		return m.PubType
	}
	return 0
}

func (m *Publication) GetPublisherID() uint64 {
	if m != nil {
		return m.PublisherID
	}
	return 0
}

func (m *Publication) GetPublicationID() int64 {
	if m != nil {
		return m.PublicationID
	}
	return 0
}

func (m *Publication) GetTopicID() uint64 {
	if m != nil {
		return m.TopicID
	}
	return 0
}

func (m *Publication) GetBrokerID() uint64 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

func (m *Publication) GetContents() [][]byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Publication) GetMAC() []byte {
	if m != nil {
		return m.MAC
	}
	return nil
}

func (m *Publication) GetChainMACs() []*ChainMAC {
	if m != nil {
		return m.ChainMACs
	}
	return nil
}

// / Publisher defines a publisher within a quorum.
type Publisher struct {
	Address     string `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	PublisherID uint64 `protobuf:"varint,2,opt,name=PublisherID" json:"PublisherID,omitempty"`
}

func (m *Publisher) Reset()                    { *m = Publisher{} }
func (m *Publisher) String() string            { return proto.CompactTextString(m) }
func (*Publisher) ProtoMessage()               {}
func (*Publisher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Publisher) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Publisher) GetPublisherID() uint64 {
	if m != nil {
		return m.PublisherID
	}
	return 0
}

// / Subscriber defines a subscriber within a subscriber pool.
type Subscriber struct {
	Address                    string `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	PoolID                     uint64 `protobuf:"varint,2,opt,name=PoolID" json:"PoolID,omitempty"`
	DestinationDistinguishment uint64 `protobuf:"varint,3,opt,name=DestinationDistinguishment" json:"DestinationDistinguishment,omitempty"`
}

func (m *Subscriber) Reset()                    { *m = Subscriber{} }
func (m *Subscriber) String() string            { return proto.CompactTextString(m) }
func (*Subscriber) ProtoMessage()               {}
func (*Subscriber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Subscriber) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Subscriber) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *Subscriber) GetDestinationDistinguishment() uint64 {
	if m != nil {
		return m.DestinationDistinguishment
	}
	return 0
}

type PubResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *PubResponse) Reset()                    { *m = PubResponse{} }
func (m *PubResponse) String() string            { return proto.CompactTextString(m) }
func (*PubResponse) ProtoMessage()               {}
func (*PubResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PubResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SubRequest struct {
	PublisherID uint64 `protobuf:"varint,1,opt,name=PublisherID" json:"PublisherID,omitempty"`
	BrokerID    uint64 `protobuf:"varint,2,opt,name=BrokerID" json:"BrokerID,omitempty"`
}

func (m *SubRequest) Reset()                    { *m = SubRequest{} }
func (m *SubRequest) String() string            { return proto.CompactTextString(m) }
func (*SubRequest) ProtoMessage()               {}
func (*SubRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SubRequest) GetPublisherID() uint64 {
	if m != nil {
		return m.PublisherID
	}
	return 0
}

func (m *SubRequest) GetBrokerID() uint64 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

type ChainMAC struct {
	From string `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=To" json:"To,omitempty"`
	MAC  []byte `protobuf:"bytes,3,opt,name=MAC,proto3" json:"MAC,omitempty"`
}

func (m *ChainMAC) Reset()                    { *m = ChainMAC{} }
func (m *ChainMAC) String() string            { return proto.CompactTextString(m) }
func (*ChainMAC) ProtoMessage()               {}
func (*ChainMAC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ChainMAC) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ChainMAC) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ChainMAC) GetMAC() []byte {
	if m != nil {
		return m.MAC
	}
	return nil
}

type ChainResponse struct {
	Valid bool `protobuf:"varint,1,opt,name=Valid" json:"Valid,omitempty"`
}

func (m *ChainResponse) Reset()                    { *m = ChainResponse{} }
func (m *ChainResponse) String() string            { return proto.CompactTextString(m) }
func (*ChainResponse) ProtoMessage()               {}
func (*ChainResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ChainResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type EchoResponse struct {
	Hello bool `protobuf:"varint,1,opt,name=Hello" json:"Hello,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EchoResponse) GetHello() bool {
	if m != nil {
		return m.Hello
	}
	return false
}

type ReadyResponse struct {
	Ready bool `protobuf:"varint,1,opt,name=Ready" json:"Ready,omitempty"`
}

func (m *ReadyResponse) Reset()                    { *m = ReadyResponse{} }
func (m *ReadyResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadyResponse) ProtoMessage()               {}
func (*ReadyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadyResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func init() {
	proto.RegisterType((*Publication)(nil), "byzantine.Publication")
	proto.RegisterType((*Publisher)(nil), "byzantine.Publisher")
	proto.RegisterType((*Subscriber)(nil), "byzantine.Subscriber")
	proto.RegisterType((*PubResponse)(nil), "byzantine.PubResponse")
	proto.RegisterType((*SubRequest)(nil), "byzantine.SubRequest")
	proto.RegisterType((*ChainMAC)(nil), "byzantine.ChainMAC")
	proto.RegisterType((*ChainResponse)(nil), "byzantine.ChainResponse")
	proto.RegisterType((*EchoResponse)(nil), "byzantine.EchoResponse")
	proto.RegisterType((*ReadyResponse)(nil), "byzantine.ReadyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Subscribe service

type SubscribeClient interface {
	Subscribe(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*ReadyResponse, error)
}

type subscribeClient struct {
	cc *grpc.ClientConn
}

func NewSubscribeClient(cc *grpc.ClientConn) SubscribeClient {
	return &subscribeClient{cc}
}

func (c *subscribeClient) Subscribe(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*ReadyResponse, error) {
	out := new(ReadyResponse)
	err := grpc.Invoke(ctx, "/byzantine.Subscribe/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Subscribe service

type SubscribeServer interface {
	Subscribe(context.Context, *SubRequest) (*ReadyResponse, error)
}

func RegisterSubscribeServer(s *grpc.Server, srv SubscribeServer) {
	s.RegisterService(&_Subscribe_serviceDesc, srv)
}

func _Subscribe_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/byzantine.Subscribe/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Subscribe(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Subscribe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "byzantine.Subscribe",
	HandlerType: (*SubscribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Subscribe_Subscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "byzantine.proto",
}

// Client API for Broker service

type BrokerClient interface {
	Echo(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*EchoResponse, error)
	GetSubscribers(ctx context.Context, in *Subscriber, opts ...grpc.CallOption) (Broker_GetSubscribersClient, error)
	RegisterSubscriber(ctx context.Context, in *Subscriber, opts ...grpc.CallOption) (*ReadyResponse, error)
	Ready(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*ReadyResponse, error)
	Receive(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*PubResponse, error)
	Push(ctx context.Context, opts ...grpc.CallOption) (Broker_PushClient, error)
	Chain(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*ChainResponse, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Echo(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := grpc.Invoke(ctx, "/byzantine.Broker/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) GetSubscribers(ctx context.Context, in *Subscriber, opts ...grpc.CallOption) (Broker_GetSubscribersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Broker_serviceDesc.Streams[0], c.cc, "/byzantine.Broker/GetSubscribers", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerGetSubscribersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broker_GetSubscribersClient interface {
	Recv() (*Subscriber, error)
	grpc.ClientStream
}

type brokerGetSubscribersClient struct {
	grpc.ClientStream
}

func (x *brokerGetSubscribersClient) Recv() (*Subscriber, error) {
	m := new(Subscriber)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerClient) RegisterSubscriber(ctx context.Context, in *Subscriber, opts ...grpc.CallOption) (*ReadyResponse, error) {
	out := new(ReadyResponse)
	err := grpc.Invoke(ctx, "/byzantine.Broker/RegisterSubscriber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Ready(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*ReadyResponse, error) {
	out := new(ReadyResponse)
	err := grpc.Invoke(ctx, "/byzantine.Broker/Ready", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Receive(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*PubResponse, error) {
	out := new(PubResponse)
	err := grpc.Invoke(ctx, "/byzantine.Broker/Receive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Push(ctx context.Context, opts ...grpc.CallOption) (Broker_PushClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Broker_serviceDesc.Streams[1], c.cc, "/byzantine.Broker/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerPushClient{stream}
	return x, nil
}

type Broker_PushClient interface {
	Send(*SubRequest) error
	Recv() (*Publication, error)
	grpc.ClientStream
}

type brokerPushClient struct {
	grpc.ClientStream
}

func (x *brokerPushClient) Send(m *SubRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerPushClient) Recv() (*Publication, error) {
	m := new(Publication)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerClient) Chain(ctx context.Context, in *Publication, opts ...grpc.CallOption) (*ChainResponse, error) {
	out := new(ChainResponse)
	err := grpc.Invoke(ctx, "/byzantine.Broker/Chain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Broker service

type BrokerServer interface {
	Echo(context.Context, *Publication) (*EchoResponse, error)
	GetSubscribers(*Subscriber, Broker_GetSubscribersServer) error
	RegisterSubscriber(context.Context, *Subscriber) (*ReadyResponse, error)
	Ready(context.Context, *Publication) (*ReadyResponse, error)
	Receive(context.Context, *Publication) (*PubResponse, error)
	Push(Broker_PushServer) error
	Chain(context.Context, *Publication) (*ChainResponse, error)
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/byzantine.Broker/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Echo(ctx, req.(*Publication))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_GetSubscribers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscriber)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerServer).GetSubscribers(m, &brokerGetSubscribersServer{stream})
}

type Broker_GetSubscribersServer interface {
	Send(*Subscriber) error
	grpc.ServerStream
}

type brokerGetSubscribersServer struct {
	grpc.ServerStream
}

func (x *brokerGetSubscribersServer) Send(m *Subscriber) error {
	return x.ServerStream.SendMsg(m)
}

func _Broker_RegisterSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscriber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).RegisterSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/byzantine.Broker/RegisterSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).RegisterSubscriber(ctx, req.(*Subscriber))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/byzantine.Broker/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Ready(ctx, req.(*Publication))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/byzantine.Broker/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Receive(ctx, req.(*Publication))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerServer).Push(&brokerPushServer{stream})
}

type Broker_PushServer interface {
	Send(*Publication) error
	Recv() (*SubRequest, error)
	grpc.ServerStream
}

type brokerPushServer struct {
	grpc.ServerStream
}

func (x *brokerPushServer) Send(m *Publication) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerPushServer) Recv() (*SubRequest, error) {
	m := new(SubRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Broker_Chain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Chain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/byzantine.Broker/Chain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Chain(ctx, req.(*Publication))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "byzantine.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Broker_Echo_Handler,
		},
		{
			MethodName: "RegisterSubscriber",
			Handler:    _Broker_RegisterSubscriber_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _Broker_Ready_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _Broker_Receive_Handler,
		},
		{
			MethodName: "Chain",
			Handler:    _Broker_Chain_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSubscribers",
			Handler:       _Broker_GetSubscribers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Push",
			Handler:       _Broker_Push_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "byzantine.proto",
}

func init() { proto.RegisterFile("byzantine.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x9c, 0xb8, 0x69, 0x72, 0x9b, 0xf4, 0x83, 0xe9, 0x9f, 0x31, 0x08, 0xac, 0x51, 0x2b,
	0xac, 0x2e, 0xea, 0x52, 0x76, 0x45, 0x42, 0x94, 0x04, 0x4a, 0x91, 0x2a, 0xa2, 0x69, 0xc4, 0x0e,
	0x84, 0xed, 0x8c, 0x62, 0xab, 0xae, 0x27, 0x78, 0x6c, 0xa4, 0x50, 0xc1, 0x82, 0x05, 0x2f, 0xc0,
	0xa3, 0xf1, 0x0a, 0xec, 0x78, 0x09, 0xe4, 0x89, 0xc7, 0x9e, 0x14, 0x5c, 0xd8, 0xcd, 0x39, 0xbe,
	0x39, 0xe7, 0xde, 0x73, 0x67, 0x02, 0xff, 0x7b, 0xb3, 0x8f, 0x6e, 0x9c, 0x86, 0x31, 0xdd, 0x9b,
	0x26, 0x2c, 0x65, 0xa8, 0x53, 0x12, 0xe6, 0x9d, 0x09, 0x63, 0x93, 0x88, 0x3a, 0xee, 0x34, 0x74,
	0xdc, 0x38, 0x66, 0xa9, 0x9b, 0x86, 0x2c, 0xe6, 0xf3, 0x42, 0xfc, 0xb5, 0x01, 0x2b, 0xc3, 0xcc,
	0x8b, 0x42, 0x5f, 0xd0, 0xc8, 0x80, 0xe5, 0x61, 0xe6, 0x8d, 0x66, 0x53, 0x6a, 0x68, 0x96, 0x66,
	0xf7, 0x88, 0x84, 0xc8, 0x2a, 0x0a, 0x79, 0x40, 0x93, 0x93, 0x81, 0xd1, 0xb0, 0x34, 0x5b, 0x27,
	0x2a, 0x85, 0xb6, 0xa1, 0xa7, 0x48, 0x9d, 0x0c, 0x8c, 0xa6, 0xa5, 0xd9, 0x88, 0x2c, 0x92, 0xb9,
	0xc3, 0x88, 0x4d, 0x43, 0xff, 0x64, 0x60, 0xe8, 0x42, 0x43, 0x42, 0x64, 0x42, 0xfb, 0x69, 0xc2,
	0xce, 0x85, 0xfc, 0x92, 0xf8, 0x54, 0xe2, 0xfc, 0x5b, 0x9f, 0xc5, 0x29, 0x8d, 0x53, 0x6e, 0xb4,
	0xac, 0xa6, 0xdd, 0x25, 0x25, 0x46, 0x37, 0xa0, 0x79, 0x7a, 0xd4, 0x37, 0x96, 0x2d, 0xcd, 0xee,
	0x92, 0xfc, 0x88, 0x1e, 0x40, 0xa7, 0x1f, 0xb8, 0x61, 0x7c, 0x7a, 0xd4, 0xe7, 0x46, 0xdb, 0x6a,
	0xda, 0x2b, 0x07, 0x6b, 0x7b, 0x55, 0x46, 0xf2, 0x1b, 0xa9, 0xaa, 0xf0, 0x31, 0x74, 0xca, 0x59,
	0xf2, 0x1e, 0x8f, 0xc6, 0xe3, 0x84, 0x72, 0x2e, 0x52, 0xe8, 0x10, 0x09, 0xff, 0x9e, 0x02, 0xfe,
	0x0c, 0x70, 0x96, 0x79, 0xdc, 0x4f, 0x42, 0xef, 0x5a, 0xa5, 0x4d, 0x68, 0x0d, 0x19, 0x8b, 0x4a,
	0x91, 0x02, 0xa1, 0xc7, 0x60, 0x0e, 0x28, 0x4f, 0xc3, 0x58, 0x04, 0x36, 0x08, 0xf3, 0xe3, 0x24,
	0x0b, 0x79, 0x70, 0x41, 0xe3, 0x54, 0x44, 0xaa, 0x93, 0x6b, 0x2a, 0xf0, 0x7d, 0xd1, 0x21, 0xa1,
	0x7c, 0xca, 0x62, 0x4e, 0xf3, 0x06, 0xce, 0x32, 0xdf, 0x97, 0x0d, 0xb4, 0x89, 0x84, 0xf8, 0xa5,
	0x68, 0x94, 0xd0, 0xf7, 0x19, 0xe5, 0xe9, 0xd5, 0xc1, 0xb4, 0xdf, 0xd7, 0xab, 0xae, 0xa7, 0xb1,
	0xb8, 0x1e, 0xfc, 0x04, 0xda, 0x32, 0x4a, 0x84, 0x40, 0x7f, 0x9e, 0xb0, 0x8b, 0x62, 0x5e, 0x71,
	0x46, 0xab, 0xd0, 0x18, 0x31, 0xf1, 0xab, 0x0e, 0x69, 0x8c, 0x98, 0x5c, 0x59, 0xb3, 0x5c, 0x19,
	0xde, 0x81, 0x9e, 0x50, 0x28, 0x1b, 0x5f, 0x87, 0xa5, 0xd7, 0x6e, 0x14, 0x8e, 0x8b, 0xb6, 0xe7,
	0x00, 0x6f, 0x43, 0xf7, 0x99, 0x1f, 0x30, 0xb5, 0xea, 0x05, 0x8d, 0x22, 0x26, 0xab, 0x04, 0xc8,
	0xc5, 0x08, 0x75, 0xc7, 0x33, 0xb5, 0x4c, 0x10, 0xb2, 0x4c, 0x80, 0x83, 0x73, 0xe8, 0x94, 0xab,
	0x42, 0x6f, 0x55, 0xb0, 0xa1, 0xdc, 0x96, 0x2a, 0x24, 0xd3, 0x50, 0xe8, 0x05, 0x03, 0x6c, 0x7d,
	0xf9, 0xfe, 0xe3, 0x5b, 0xc3, 0xc4, 0x1b, 0x0e, 0x2f, 0x97, 0x5f, 0x1d, 0x0f, 0xb5, 0xdd, 0x83,
	0x9f, 0x3a, 0xb4, 0xe6, 0x79, 0xa1, 0x21, 0xe8, 0xf9, 0x10, 0x68, 0x53, 0x91, 0x53, 0x1e, 0x89,
	0xb9, 0xa5, 0xf0, 0xea, 0xb4, 0x78, 0x4b, 0xb8, 0xdc, 0xc4, 0x5d, 0xc7, 0x13, 0x4a, 0x0e, 0xf5,
	0x03, 0x76, 0xa8, 0xed, 0xa2, 0x37, 0xb0, 0x7a, 0x4c, 0xd3, 0xea, 0xde, 0xf1, 0xab, 0x13, 0x14,
	0xbc, 0xf9, 0x67, 0x1a, 0xdf, 0x16, 0xc2, 0x1b, 0x68, 0x4d, 0x0a, 0x57, 0x53, 0xf0, 0x7d, 0x0d,
	0x9d, 0x03, 0x22, 0x74, 0x12, 0xf2, 0x94, 0x26, 0xca, 0xdd, 0xae, 0xb1, 0xa8, 0x0f, 0x69, 0x5b,
	0xb8, 0xdc, 0xc5, 0xb7, 0xa4, 0x4b, 0x52, 0x88, 0x3a, 0x97, 0xf3, 0xcb, 0xff, 0x29, 0x9f, 0xe5,
	0xac, 0xd8, 0x55, 0x6d, 0x3c, 0xf5, 0x06, 0x86, 0x30, 0x40, 0xb8, 0x57, 0x19, 0xb8, 0xe3, 0x59,
	0x2e, 0xfa, 0x0e, 0x96, 0x09, 0xf5, 0x69, 0xf8, 0x81, 0xd6, 0xca, 0x5e, 0xe1, 0x4b, 0xd1, 0x1d,
	0x21, 0x7a, 0x0f, 0x9b, 0x52, 0x74, 0x9a, 0x45, 0x91, 0x73, 0xa9, 0x3c, 0x0d, 0xd1, 0xf6, 0x2b,
	0xd0, 0x87, 0x19, 0x0f, 0xea, 0xae, 0x4e, 0x8d, 0x2b, 0x5e, 0x17, 0xea, 0xab, 0xa8, 0x5b, 0xa9,
	0xf3, 0xc0, 0xd6, 0xf6, 0x35, 0xf4, 0x08, 0x96, 0xc4, 0x8b, 0xf8, 0xa7, 0x1c, 0x16, 0xde, 0x0e,
	0xfe, 0xcf, 0x6b, 0x89, 0xbf, 0xf7, 0x87, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe3, 0x53,
	0xea, 0x1a, 0x06, 0x00, 0x00,
}
