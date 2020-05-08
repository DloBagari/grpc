// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.5.1
// source: greet/greetpb/streamclient.proto

package greetpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GreetClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GreetClient) Reset() {
	*x = GreetClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_streamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetClient) ProtoMessage() {}

func (x *GreetClient) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_streamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetClient.ProtoReflect.Descriptor instead.
func (*GreetClient) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_streamclient_proto_rawDescGZIP(), []int{0}
}

func (x *GreetClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GreetClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *GreetClient `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetClientRequest) Reset() {
	*x = GreetClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_streamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetClientRequest) ProtoMessage() {}

func (x *GreetClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_streamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetClientRequest.ProtoReflect.Descriptor instead.
func (*GreetClientRequest) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_streamclient_proto_rawDescGZIP(), []int{1}
}

func (x *GreetClientRequest) GetGreeting() *GreetClient {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetClientResponse) Reset() {
	*x = GreetClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greetpb_streamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetClientResponse) ProtoMessage() {}

func (x *GreetClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greetpb_streamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetClientResponse.ProtoReflect.Descriptor instead.
func (*GreetClientResponse) Descriptor() ([]byte, []int) {
	return file_greet_greetpb_streamclient_proto_rawDescGZIP(), []int{2}
}

func (x *GreetClientResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_greet_greetpb_streamclient_proto protoreflect.FileDescriptor

var file_greet_greetpb_streamclient_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x21, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x12, 0x47, 0x72, 0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x2d, 0x0a, 0x13, 0x47, 0x72, 0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x65, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x50, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x20, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_greetpb_streamclient_proto_rawDescOnce sync.Once
	file_greet_greetpb_streamclient_proto_rawDescData = file_greet_greetpb_streamclient_proto_rawDesc
)

func file_greet_greetpb_streamclient_proto_rawDescGZIP() []byte {
	file_greet_greetpb_streamclient_proto_rawDescOnce.Do(func() {
		file_greet_greetpb_streamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_greetpb_streamclient_proto_rawDescData)
	})
	return file_greet_greetpb_streamclient_proto_rawDescData
}

var file_greet_greetpb_streamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_greet_greetpb_streamclient_proto_goTypes = []interface{}{
	(*GreetClient)(nil),         // 0: streamclient.GreetClient
	(*GreetClientRequest)(nil),  // 1: streamclient.GreetClientRequest
	(*GreetClientResponse)(nil), // 2: streamclient.GreetClientResponse
}
var file_greet_greetpb_streamclient_proto_depIdxs = []int32{
	0, // 0: streamclient.GreetClientRequest.greeting:type_name -> streamclient.GreetClient
	1, // 1: streamclient.GreetClientStream.Greet:input_type -> streamclient.GreetClientRequest
	2, // 2: streamclient.GreetClientStream.Greet:output_type -> streamclient.GreetClientResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_greet_greetpb_streamclient_proto_init() }
func file_greet_greetpb_streamclient_proto_init() {
	if File_greet_greetpb_streamclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_greetpb_streamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetClient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_greet_greetpb_streamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetClientRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_greet_greetpb_streamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetClientResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greet_greetpb_streamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_greetpb_streamclient_proto_goTypes,
		DependencyIndexes: file_greet_greetpb_streamclient_proto_depIdxs,
		MessageInfos:      file_greet_greetpb_streamclient_proto_msgTypes,
	}.Build()
	File_greet_greetpb_streamclient_proto = out.File
	file_greet_greetpb_streamclient_proto_rawDesc = nil
	file_greet_greetpb_streamclient_proto_goTypes = nil
	file_greet_greetpb_streamclient_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetClientStreamClient is the client API for GreetClientStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetClientStreamClient interface {
	Greet(ctx context.Context, opts ...grpc.CallOption) (GreetClientStream_GreetClient, error)
}

type greetClientStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClientStreamClient(cc grpc.ClientConnInterface) GreetClientStreamClient {
	return &greetClientStreamClient{cc}
}

func (c *greetClientStreamClient) Greet(ctx context.Context, opts ...grpc.CallOption) (GreetClientStream_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetClientStream_serviceDesc.Streams[0], "/streamclient.GreetClientStream/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetClientStreamGreetClient{stream}
	return x, nil
}

type GreetClientStream_GreetClient interface {
	Send(*GreetClientRequest) error
	CloseAndRecv() (*GreetClientResponse, error)
	grpc.ClientStream
}

type greetClientStreamGreetClient struct {
	grpc.ClientStream
}

func (x *greetClientStreamGreetClient) Send(m *GreetClientRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetClientStreamGreetClient) CloseAndRecv() (*GreetClientResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetClientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetClientStreamServer is the server API for GreetClientStream service.
type GreetClientStreamServer interface {
	Greet(GreetClientStream_GreetServer) error
}

// UnimplementedGreetClientStreamServer can be embedded to have forward compatible implementations.
type UnimplementedGreetClientStreamServer struct {
}

func (*UnimplementedGreetClientStreamServer) Greet(GreetClientStream_GreetServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetClientStreamServer(s *grpc.Server, srv GreetClientStreamServer) {
	s.RegisterService(&_GreetClientStream_serviceDesc, srv)
}

func _GreetClientStream_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetClientStreamServer).Greet(&greetClientStreamGreetServer{stream})
}

type GreetClientStream_GreetServer interface {
	SendAndClose(*GreetClientResponse) error
	Recv() (*GreetClientRequest, error)
	grpc.ServerStream
}

type greetClientStreamGreetServer struct {
	grpc.ServerStream
}

func (x *greetClientStreamGreetServer) SendAndClose(m *GreetClientResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetClientStreamGreetServer) Recv() (*GreetClientRequest, error) {
	m := new(GreetClientRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetClientStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "streamclient.GreetClientStream",
	HandlerType: (*GreetClientStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _GreetClientStream_Greet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/streamclient.proto",
}
