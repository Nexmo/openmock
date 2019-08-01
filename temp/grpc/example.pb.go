// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package openmock

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MyRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyRequest) Reset()         { *m = MyRequest{} }
func (m *MyRequest) String() string { return proto.CompactTextString(m) }
func (*MyRequest) ProtoMessage()    {}
func (*MyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *MyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyRequest.Unmarshal(m, b)
}
func (m *MyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyRequest.Marshal(b, m, deterministic)
}
func (m *MyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyRequest.Merge(m, src)
}
func (m *MyRequest) XXX_Size() int {
	return xxx_messageInfo_MyRequest.Size(m)
}
func (m *MyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MyRequest proto.InternalMessageInfo

func (m *MyRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type MyResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyResponse) Reset()         { *m = MyResponse{} }
func (m *MyResponse) String() string { return proto.CompactTextString(m) }
func (*MyResponse) ProtoMessage()    {}
func (*MyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *MyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyResponse.Unmarshal(m, b)
}
func (m *MyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyResponse.Marshal(b, m, deterministic)
}
func (m *MyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyResponse.Merge(m, src)
}
func (m *MyResponse) XXX_Size() int {
	return xxx_messageInfo_MyResponse.Size(m)
}
func (m *MyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MyResponse proto.InternalMessageInfo

func (m *MyResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MyRequest)(nil), "openmock.MyRequest")
	proto.RegisterType((*MyResponse)(nil), "openmock.MyResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x2f, 0x48, 0xcd, 0xcb,
	0xcd, 0x4f, 0xce, 0x56, 0x92, 0xe7, 0xe2, 0xf4, 0xad, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x95, 0xd4, 0xb8, 0xb8, 0x40, 0x0a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8,
	0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xa1, 0x8a, 0x60, 0x5c, 0x23, 0x17, 0x90, 0x41,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xe6, 0x5c, 0x1c, 0xbe, 0x95, 0xbe, 0xa9, 0x25,
	0x19, 0xf9, 0x29, 0x42, 0xc2, 0x7a, 0x30, 0xcb, 0xf4, 0xe0, 0x36, 0x49, 0x89, 0xa0, 0x0a, 0x42,
	0x4c, 0x57, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xef, 0x9d,
	0xaf, 0x57, 0xb0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyServiceClient interface {
	MyMethod(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
}

type myServiceClient struct {
	cc *grpc.ClientConn
}

func NewMyServiceClient(cc *grpc.ClientConn) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) MyMethod(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/openmock.MyService/MyMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
type MyServiceServer interface {
	MyMethod(context.Context, *MyRequest) (*MyResponse, error)
}

// UnimplementedMyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (*UnimplementedMyServiceServer) MyMethod(ctx context.Context, req *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyMethod not implemented")
}

func RegisterMyServiceServer(s *grpc.Server, srv MyServiceServer) {
	s.RegisterService(&_MyService_serviceDesc, srv)
}

func _MyService_MyMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).MyMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmock.MyService/MyMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).MyMethod(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmock.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MyMethod",
			Handler:    _MyService_MyMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}