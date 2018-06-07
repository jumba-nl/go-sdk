// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package api_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type AuthGetOptions struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthGetOptions) Reset()         { *m = AuthGetOptions{} }
func (m *AuthGetOptions) String() string { return proto.CompactTextString(m) }
func (*AuthGetOptions) ProtoMessage()    {}
func (*AuthGetOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_e1c8515303b2b671, []int{0}
}
func (m *AuthGetOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthGetOptions.Unmarshal(m, b)
}
func (m *AuthGetOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthGetOptions.Marshal(b, m, deterministic)
}
func (dst *AuthGetOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthGetOptions.Merge(dst, src)
}
func (m *AuthGetOptions) XXX_Size() int {
	return xxx_messageInfo_AuthGetOptions.Size(m)
}
func (m *AuthGetOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthGetOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AuthGetOptions proto.InternalMessageInfo

func (m *AuthGetOptions) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Profile struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_e1c8515303b2b671, []int{1}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (dst *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(dst, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthGetOptions)(nil), "api_v1.AuthGetOptions")
	proto.RegisterType((*Profile)(nil), "api_v1.Profile")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Get(ctx context.Context, in *AuthGetOptions, opts ...grpc.CallOption) (*Profile, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Get(ctx context.Context, in *AuthGetOptions, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/api_v1.AuthService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Get(context.Context, *AuthGetOptions) (*Profile, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthGetOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_v1.AuthService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Get(ctx, req.(*AuthGetOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api_v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AuthService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_e1c8515303b2b671) }

var fileDescriptor_auth_e1c8515303b2b671 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2c, 0xc8, 0x8c, 0x2f, 0x33, 0x54, 0x52,
	0xe3, 0xe2, 0x73, 0x2c, 0x2d, 0xc9, 0x70, 0x4f, 0x2d, 0xf1, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b,
	0x16, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0xcf, 0x4e, 0xcd, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x82, 0x70, 0x94, 0x74, 0xb9, 0xd8, 0x03, 0x8a, 0xf2, 0xd3, 0x32, 0x73, 0x52, 0x85, 0xf8, 0xb8,
	0x98, 0x32, 0x53, 0xa0, 0xb2, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x4c, 0x60, 0x11, 0x30, 0xdb, 0xc8, 0x9e, 0x8b, 0x1b, 0x64, 0x6c, 0x70, 0x6a, 0x51, 0x59,
	0x66, 0x72, 0xaa, 0x90, 0x01, 0x17, 0xb3, 0x7b, 0x6a, 0x89, 0x90, 0x98, 0x1e, 0xc4, 0x56, 0x3d,
	0x54, 0x2b, 0xa5, 0xf8, 0x61, 0xe2, 0x50, 0x2b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xce, 0x34, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x58, 0xd3, 0x95, 0xd2, 0xb4, 0x00, 0x00, 0x00,
}
