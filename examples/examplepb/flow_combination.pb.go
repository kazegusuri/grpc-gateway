// Code generated by protoc-gen-go.
// source: examples/examplepb/flow_combination.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EmptyProto struct {
}

func (m *EmptyProto) Reset()                    { *m = EmptyProto{} }
func (m *EmptyProto) String() string            { return proto.CompactTextString(m) }
func (*EmptyProto) ProtoMessage()               {}
func (*EmptyProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type NonEmptyProto struct {
	A string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C string `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *NonEmptyProto) Reset()                    { *m = NonEmptyProto{} }
func (m *NonEmptyProto) String() string            { return proto.CompactTextString(m) }
func (*NonEmptyProto) ProtoMessage()               {}
func (*NonEmptyProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type UnaryProto struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
}

func (m *UnaryProto) Reset()                    { *m = UnaryProto{} }
func (m *UnaryProto) String() string            { return proto.CompactTextString(m) }
func (*UnaryProto) ProtoMessage()               {}
func (*UnaryProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type NestedProto struct {
	A *UnaryProto `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string      `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C string      `protobuf:"bytes,3,opt,name=c" json:"c,omitempty"`
}

func (m *NestedProto) Reset()                    { *m = NestedProto{} }
func (m *NestedProto) String() string            { return proto.CompactTextString(m) }
func (*NestedProto) ProtoMessage()               {}
func (*NestedProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *NestedProto) GetA() *UnaryProto {
	if m != nil {
		return m.A
	}
	return nil
}

type SingleNestedProto struct {
	A *UnaryProto `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
}

func (m *SingleNestedProto) Reset()                    { *m = SingleNestedProto{} }
func (m *SingleNestedProto) String() string            { return proto.CompactTextString(m) }
func (*SingleNestedProto) ProtoMessage()               {}
func (*SingleNestedProto) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *SingleNestedProto) GetA() *UnaryProto {
	if m != nil {
		return m.A
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyProto)(nil), "gengo.grpc.gateway.examples.examplepb.EmptyProto")
	proto.RegisterType((*NonEmptyProto)(nil), "gengo.grpc.gateway.examples.examplepb.NonEmptyProto")
	proto.RegisterType((*UnaryProto)(nil), "gengo.grpc.gateway.examples.examplepb.UnaryProto")
	proto.RegisterType((*NestedProto)(nil), "gengo.grpc.gateway.examples.examplepb.NestedProto")
	proto.RegisterType((*SingleNestedProto)(nil), "gengo.grpc.gateway.examples.examplepb.SingleNestedProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for FlowCombination service

type FlowCombinationClient interface {
	RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcEmptyStreamClient, error)
	StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyRpcClient, error)
	StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyStreamClient, error)
	RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error)
	RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcBodyStreamClient, error)
	RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathSingleNestedStreamClient, error)
	RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathNestedStreamClient, error)
}

type flowCombinationClient struct {
	cc *grpc.ClientConn
}

func NewFlowCombinationClient(cc *grpc.ClientConn) FlowCombinationClient {
	return &flowCombinationClient{cc}
}

func (c *flowCombinationClient) RpcEmptyRpc(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcEmptyStream(ctx context.Context, in *EmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcEmptyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[0], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcEmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcEmptyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcEmptyStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcEmptyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) StreamEmptyRpc(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyRpcClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[1], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/StreamEmptyRpc", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationStreamEmptyRpcClient{stream}
	return x, nil
}

type FlowCombination_StreamEmptyRpcClient interface {
	Send(*EmptyProto) error
	CloseAndRecv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationStreamEmptyRpcClient struct {
	grpc.ClientStream
}

func (x *flowCombinationStreamEmptyRpcClient) Send(m *EmptyProto) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyRpcClient) CloseAndRecv() (*EmptyProto, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) StreamEmptyStream(ctx context.Context, opts ...grpc.CallOption) (FlowCombination_StreamEmptyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[2], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/StreamEmptyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationStreamEmptyStreamClient{stream}
	return x, nil
}

type FlowCombination_StreamEmptyStreamClient interface {
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationStreamEmptyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationStreamEmptyStreamClient) Send(m *EmptyProto) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcBodyRpc(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedRpc(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcPathNestedRpc(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (*EmptyProto, error) {
	out := new(EmptyProto)
	err := grpc.Invoke(ctx, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedRpc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowCombinationClient) RpcBodyStream(ctx context.Context, in *NonEmptyProto, opts ...grpc.CallOption) (FlowCombination_RpcBodyStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[3], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcBodyStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcBodyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcBodyStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcBodyStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcBodyStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcPathSingleNestedStream(ctx context.Context, in *SingleNestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathSingleNestedStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[4], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathSingleNestedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcPathSingleNestedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcPathSingleNestedStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcPathSingleNestedStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcPathSingleNestedStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowCombinationClient) RpcPathNestedStream(ctx context.Context, in *NestedProto, opts ...grpc.CallOption) (FlowCombination_RpcPathNestedStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FlowCombination_serviceDesc.Streams[5], c.cc, "/gengo.grpc.gateway.examples.examplepb.FlowCombination/RpcPathNestedStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowCombinationRpcPathNestedStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlowCombination_RpcPathNestedStreamClient interface {
	Recv() (*EmptyProto, error)
	grpc.ClientStream
}

type flowCombinationRpcPathNestedStreamClient struct {
	grpc.ClientStream
}

func (x *flowCombinationRpcPathNestedStreamClient) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FlowCombination service

type FlowCombinationServer interface {
	RpcEmptyRpc(context.Context, *EmptyProto) (*EmptyProto, error)
	RpcEmptyStream(*EmptyProto, FlowCombination_RpcEmptyStreamServer) error
	StreamEmptyRpc(FlowCombination_StreamEmptyRpcServer) error
	StreamEmptyStream(FlowCombination_StreamEmptyStreamServer) error
	RpcBodyRpc(context.Context, *NonEmptyProto) (*EmptyProto, error)
	RpcPathSingleNestedRpc(context.Context, *SingleNestedProto) (*EmptyProto, error)
	RpcPathNestedRpc(context.Context, *NestedProto) (*EmptyProto, error)
	RpcBodyStream(*NonEmptyProto, FlowCombination_RpcBodyStreamServer) error
	RpcPathSingleNestedStream(*SingleNestedProto, FlowCombination_RpcPathSingleNestedStreamServer) error
	RpcPathNestedStream(*NestedProto, FlowCombination_RpcPathNestedStreamServer) error
}

func RegisterFlowCombinationServer(s *grpc.Server, srv FlowCombinationServer) {
	s.RegisterService(&_FlowCombination_serviceDesc, srv)
}

func _FlowCombination_RpcEmptyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EmptyProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcEmptyRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcEmptyStream(m, &flowCombinationRpcEmptyStreamServer{stream})
}

type FlowCombination_RpcEmptyStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcEmptyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_StreamEmptyRpc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyRpc(&flowCombinationStreamEmptyRpcServer{stream})
}

type FlowCombination_StreamEmptyRpcServer interface {
	SendAndClose(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ServerStream
}

type flowCombinationStreamEmptyRpcServer struct {
	grpc.ServerStream
}

func (x *flowCombinationStreamEmptyRpcServer) SendAndClose(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyRpcServer) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowCombination_StreamEmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowCombinationServer).StreamEmptyStream(&flowCombinationStreamEmptyStreamServer{stream})
}

type FlowCombination_StreamEmptyStreamServer interface {
	Send(*EmptyProto) error
	Recv() (*EmptyProto, error)
	grpc.ServerStream
}

type flowCombinationStreamEmptyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationStreamEmptyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowCombinationStreamEmptyStreamServer) Recv() (*EmptyProto, error) {
	m := new(EmptyProto)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowCombination_RpcBodyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NonEmptyProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcBodyRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcPathSingleNestedRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SingleNestedProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcPathSingleNestedRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcPathNestedRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NestedProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(FlowCombinationServer).RpcPathNestedRpc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _FlowCombination_RpcBodyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NonEmptyProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcBodyStream(m, &flowCombinationRpcBodyStreamServer{stream})
}

type FlowCombination_RpcBodyStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcBodyStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcBodyStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_RpcPathSingleNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SingleNestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathSingleNestedStream(m, &flowCombinationRpcPathSingleNestedStreamServer{stream})
}

type FlowCombination_RpcPathSingleNestedStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcPathSingleNestedStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcPathSingleNestedStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

func _FlowCombination_RpcPathNestedStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NestedProto)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlowCombinationServer).RpcPathNestedStream(m, &flowCombinationRpcPathNestedStreamServer{stream})
}

type FlowCombination_RpcPathNestedStreamServer interface {
	Send(*EmptyProto) error
	grpc.ServerStream
}

type flowCombinationRpcPathNestedStreamServer struct {
	grpc.ServerStream
}

func (x *flowCombinationRpcPathNestedStreamServer) Send(m *EmptyProto) error {
	return x.ServerStream.SendMsg(m)
}

var _FlowCombination_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.FlowCombination",
	HandlerType: (*FlowCombinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcEmptyRpc",
			Handler:    _FlowCombination_RpcEmptyRpc_Handler,
		},
		{
			MethodName: "RpcBodyRpc",
			Handler:    _FlowCombination_RpcBodyRpc_Handler,
		},
		{
			MethodName: "RpcPathSingleNestedRpc",
			Handler:    _FlowCombination_RpcPathSingleNestedRpc_Handler,
		},
		{
			MethodName: "RpcPathNestedRpc",
			Handler:    _FlowCombination_RpcPathNestedRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RpcEmptyStream",
			Handler:       _FlowCombination_RpcEmptyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamEmptyRpc",
			Handler:       _FlowCombination_StreamEmptyRpc_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamEmptyStream",
			Handler:       _FlowCombination_StreamEmptyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RpcBodyStream",
			Handler:       _FlowCombination_RpcBodyStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathSingleNestedStream",
			Handler:       _FlowCombination_RpcPathSingleNestedStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RpcPathNestedStream",
			Handler:       _FlowCombination_RpcPathNestedStream_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor2 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x96, 0x3f, 0x6f, 0xd3, 0x4e,
	0x18, 0xc7, 0x75, 0xa9, 0xf4, 0x93, 0xf2, 0xe4, 0x97, 0x92, 0x3a, 0x28, 0x24, 0x21, 0xd0, 0xf6,
	0x28, 0x22, 0x20, 0x61, 0x87, 0xc0, 0x80, 0x10, 0x53, 0x11, 0x8c, 0x55, 0xa1, 0x62, 0xf1, 0x82,
	0xce, 0xae, 0x71, 0x23, 0x25, 0x3e, 0xd7, 0x31, 0x0a, 0x55, 0xd5, 0x85, 0x95, 0xb1, 0x23, 0x2f,
	0x81, 0x91, 0x09, 0x36, 0x06, 0x66, 0x86, 0x4a, 0xbc, 0x02, 0x16, 0x5e, 0x01, 0x2b, 0xe7, 0xe7,
	0xce, 0xbe, 0x18, 0x12, 0x1a, 0x0a, 0xf2, 0x50, 0xf5, 0xec, 0xe7, 0x8f, 0x3f, 0xf7, 0x7d, 0xbe,
	0x77, 0x0a, 0x5c, 0xf7, 0x5e, 0xb2, 0x51, 0x38, 0xf4, 0xc6, 0x96, 0x5a, 0x84, 0x8e, 0xf5, 0x7c,
	0xc8, 0x27, 0xcf, 0x5c, 0x3e, 0x72, 0x06, 0x01, 0x8b, 0x07, 0x3c, 0x30, 0xc3, 0x88, 0xc7, 0xdc,
	0xb8, 0xea, 0x7b, 0x81, 0xcf, 0x4d, 0x3f, 0x0a, 0x5d, 0xd3, 0x67, 0xb1, 0x37, 0x61, 0x07, 0x66,
	0x5a, 0x6d, 0x66, 0xd5, 0xed, 0x8e, 0xcf, 0xb9, 0x3f, 0xf4, 0x2c, 0x16, 0x0e, 0x2c, 0x16, 0x04,
	0x3c, 0xc6, 0x1e, 0x63, 0xd9, 0x84, 0xfe, 0x0f, 0xf0, 0x70, 0x14, 0xc6, 0x07, 0xdb, 0xf8, 0xd4,
	0x83, 0xea, 0x16, 0x0f, 0xf4, 0x0b, 0xa3, 0x0c, 0x84, 0x35, 0xc9, 0x1a, 0xe9, 0x96, 0x93, 0xa5,
	0xd3, 0x2c, 0xa5, 0x4b, 0xb7, 0xb9, 0x94, 0x2c, 0x69, 0x0b, 0xe0, 0x69, 0xc0, 0x22, 0x95, 0x5e,
	0x81, 0xa5, 0x71, 0x1c, 0xc9, 0x02, 0xca, 0xa0, 0xb2, 0xe5, 0x8d, 0x63, 0x6f, 0x57, 0xc6, 0xee,
	0xa7, 0xad, 0x2a, 0xfd, 0x5b, 0xe6, 0x42, 0xe8, 0xe6, 0x54, 0xe7, 0xd9, 0x5f, 0x7f, 0x0c, 0x2b,
	0x3b, 0x83, 0x40, 0xec, 0xee, 0x9f, 0x7d, 0xa8, 0xff, 0xa9, 0x06, 0xe7, 0x1e, 0x09, 0xc1, 0x1f,
	0x68, 0xbd, 0x8d, 0xd7, 0x04, 0x2a, 0x4f, 0x42, 0x17, 0x75, 0x11, 0xff, 0x8d, 0x45, 0xdb, 0x6a,
	0x21, 0xdb, 0x7f, 0x5e, 0x42, 0x1b, 0xaf, 0x4e, 0xbe, 0x1e, 0x97, 0x6a, 0x74, 0xd9, 0x12, 0x35,
	0x96, 0x97, 0x04, 0x92, 0x95, 0x71, 0x4c, 0x60, 0x39, 0xa5, 0xd9, 0x89, 0x23, 0x8f, 0x8d, 0x0a,
	0x02, 0x6a, 0x21, 0x50, 0x9d, 0xae, 0x4c, 0x01, 0x8d, 0x11, 0xa0, 0x47, 0x90, 0x4a, 0xd2, 0x14,
	0x2c, 0x93, 0xa6, 0x92, 0x2c, 0x5a, 0xa9, 0x2e, 0x31, 0xde, 0x10, 0xe1, 0x10, 0x4d, 0x55, 0xa8,
	0x5c, 0x1d, 0x04, 0x6b, 0xd0, 0xf3, 0x79, 0x30, 0xf9, 0xd0, 0x25, 0x42, 0xb3, 0xf7, 0x25, 0x00,
	0x21, 0xd4, 0x26, 0xdf, 0x45, 0xbd, 0xee, 0x2c, 0xf8, 0x8d, 0xdc, 0x11, 0x3d, 0x0b, 0xd9, 0x47,
	0x82, 0x68, 0x1f, 0x08, 0xad, 0xe2, 0x28, 0x1d, 0x01, 0x90, 0x2c, 0xee, 0x91, 0x1b, 0xf6, 0x45,
	0xda, 0xc2, 0x77, 0x21, 0x8b, 0xf7, 0xac, 0x43, 0x76, 0x64, 0x1d, 0x3a, 0xe2, 0xcf, 0x3d, 0x4a,
	0x5e, 0xda, 0xa9, 0x19, 0xf7, 0x5f, 0x78, 0x11, 0x56, 0xd8, 0xab, 0xb4, 0xad, 0x5b, 0xe4, 0x6a,
	0xb0, 0x9f, 0x6b, 0x37, 0x69, 0x5d, 0x27, 0x64, 0x75, 0x49, 0x64, 0x9d, 0x76, 0x66, 0x94, 0xe6,
	0x52, 0x5a, 0xf4, 0x42, 0x1e, 0x26, 0x8b, 0x1a, 0x6f, 0x09, 0x34, 0x84, 0x68, 0xdb, 0x22, 0x32,
	0x7d, 0x05, 0x24, 0x3a, 0xde, 0x5d, 0x50, 0x91, 0x5f, 0xae, 0x8e, 0xb3, 0x68, 0xb9, 0x81, 0x52,
	0x5e, 0x56, 0x7b, 0x49, 0x40, 0x6f, 0x06, 0xd8, 0x51, 0xf0, 0x9a, 0x62, 0xd6, 0x28, 0x84, 0xf1,
	0x8d, 0x40, 0x4d, 0xd1, 0x6a, 0xce, 0xfe, 0xa2, 0xf3, 0xfe, 0x3b, 0xc2, 0x00, 0x09, 0xf7, 0xe8,
	0xda, 0x5c, 0xc2, 0xa9, 0x71, 0x9d, 0xb2, 0x91, 0x6c, 0x68, 0x73, 0xe2, 0xa2, 0x85, 0x71, 0x52,
	0x82, 0xaa, 0x72, 0xb5, 0x3a, 0x6f, 0x85, 0x19, 0xfb, 0x8b, 0x34, 0xf6, 0x67, 0x42, 0x6b, 0xda,
	0x5a, 0xf2, 0xc0, 0x25, 0xde, 0x9e, 0xde, 0x5c, 0xce, 0xdb, 0x32, 0xc5, 0x4e, 0xaf, 0x36, 0xe9,
	0x32, 0xf5, 0x92, 0xd2, 0x4b, 0x73, 0x1c, 0x9e, 0x36, 0x76, 0xc5, 0xa1, 0x69, 0xfc, 0x6c, 0x72,
	0x1d, 0xdc, 0xa0, 0xab, 0x73, 0x7d, 0xae, 0xb3, 0x3a, 0xea, 0x20, 0xcd, 0x4c, 0x10, 0x77, 0xc5,
	0x3b, 0x02, 0xad, 0x19, 0x7e, 0x57, 0x0a, 0x17, 0x6a, 0xf9, 0x6b, 0x28, 0xf2, 0xba, 0xda, 0xd6,
	0x2c, 0x27, 0x64, 0xd4, 0xdf, 0x09, 0xd4, 0x73, 0xbe, 0x57, 0xbc, 0x05, 0x59, 0x7f, 0x82, 0xa4,
	0xfb, 0xf4, 0xca, 0x6f, 0xad, 0xaf, 0x87, 0x70, 0xfa, 0x9e, 0xb2, 0x69, 0xce, 0x4f, 0x11, 0x8d,
	0x7a, 0x64, 0xb3, 0x62, 0x97, 0x33, 0x24, 0xe7, 0x3f, 0xfc, 0xb1, 0x75, 0xfb, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x94, 0xb8, 0xb0, 0x35, 0xde, 0x09, 0x00, 0x00,
}
