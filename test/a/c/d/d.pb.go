// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/a/c/d/d.proto

package d

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

type BaseRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cd17698b66b054, []int{0}
}

func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (m *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(m, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

type BaseResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error"`
	Details              []string `protobuf:"bytes,4,rep,name=details,proto3" json:"details"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4cd17698b66b054, []int{1}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BaseResponse) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "d.BaseRequest")
	proto.RegisterType((*BaseResponse)(nil), "d.BaseResponse")
}

func init() { proto.RegisterFile("test/a/c/d/d.proto", fileDescriptor_d4cd17698b66b054) }

var fileDescriptor_d4cd17698b66b054 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0xb1, 0x8a, 0xc2, 0x40,
	0x10, 0xc6, 0x71, 0xf6, 0x92, 0xdc, 0x71, 0x73, 0xa7, 0xc2, 0x60, 0xb1, 0x58, 0x85, 0x54, 0x69,
	0xcc, 0x82, 0xe2, 0x0b, 0xd8, 0xdb, 0x04, 0x5f, 0x60, 0xcd, 0x0c, 0x2a, 0x44, 0x37, 0xee, 0x8c,
	0xef, 0x2f, 0x49, 0x54, 0xac, 0x76, 0x7f, 0xc3, 0x57, 0xfc, 0x01, 0x95, 0x45, 0x9d, 0x77, 0x8d,
	0x23, 0x47, 0x55, 0x17, 0x83, 0x06, 0x34, 0x54, 0x4c, 0xe0, 0x6f, 0xeb, 0x85, 0x6b, 0xbe, 0xdd,
	0x59, 0xb4, 0x68, 0xe1, 0x7f, 0xa4, 0x74, 0xe1, 0x2a, 0x8c, 0x08, 0x69, 0x13, 0x88, 0xad, 0xc9,
	0x4d, 0x99, 0xd5, 0xc3, 0x1f, 0x2d, 0xfc, 0x5c, 0x58, 0xc4, 0x1f, 0xd9, 0x7e, 0xe5, 0xa6, 0xfc,
	0xad, 0x5f, 0xc4, 0x39, 0x64, 0x1c, 0x63, 0x88, 0x36, 0x19, 0xee, 0x23, 0xfa, 0x3d, 0xb1, 0xfa,
	0x73, 0x2b, 0x36, 0xcd, 0x93, 0x7e, 0xff, 0xe4, 0x6a, 0x03, 0xe9, 0x9e, 0x45, 0x71, 0x09, 0xd0,
	0xbf, 0x3b, 0xd6, 0x53, 0x20, 0x9c, 0x56, 0x54, 0x7d, 0x34, 0x2d, 0x66, 0x6f, 0x8f, 0x51, 0x87,
	0xef, 0xa1, 0x7e, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xa0, 0x4c, 0xc4, 0xd3, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	TestMethod(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) TestMethod(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/d.Test/TestMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	TestMethod(context.Context, *BaseRequest) (*BaseResponse, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) TestMethod(ctx context.Context, req *BaseRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_TestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).TestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d.Test/TestMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).TestMethod(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "d.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMethod",
			Handler:    _Test_TestMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/a/c/d/d.proto",
}
