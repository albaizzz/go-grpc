// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/user.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserGender int32

const (
	UserGender_UNDEFINED UserGender = 0
	UserGender_MALE      UserGender = 1
	UserGender_FEMALE    UserGender = 2
)

var UserGender_name = map[int32]string{
	0: "UNDEFINED",
	1: "MALE",
	2: "FEMALE",
}
var UserGender_value = map[string]int32{
	"UNDEFINED": 0,
	"MALE":      1,
	"FEMALE":    2,
}

func (x UserGender) String() string {
	return proto.EnumName(UserGender_name, int32(x))
}
func (UserGender) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type User struct {
	Id       string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password string     `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Gender   UserGender `protobuf:"varint,4,opt,name=gender,enum=model.UserGender" json:"gender,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetGender() UserGender {
	if m != nil {
		return m.Gender
	}
	return UserGender_UNDEFINED
}

type UserList struct {
	List []*User `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *UserList) Reset()                    { *m = UserList{} }
func (m *UserList) String() string            { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()               {}
func (*UserList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*UserList)(nil), "model.UserList")
	proto.RegisterEnum("model.UserGender", UserGender_name, UserGender_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	List(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*UserList, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/model.Users/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) List(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := grpc.Invoke(ctx, "/model.Users/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	Register(context.Context, *User) (*google_protobuf.Empty, error)
	List(context.Context, *google_protobuf.Empty) (*UserList, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Users/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).List(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Users_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models/user.proto",
}

func init() { proto.RegisterFile("models/user.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x6b, 0xb3, 0x40,
	0x10, 0xc6, 0xd5, 0x6c, 0xc4, 0x4c, 0x78, 0xf3, 0x26, 0x73, 0x28, 0x62, 0x0f, 0x15, 0x4f, 0xb6,
	0x85, 0x15, 0xed, 0x27, 0x28, 0xc4, 0x94, 0x42, 0x9a, 0x83, 0x90, 0x0f, 0x90, 0xe0, 0x54, 0x04,
	0xcd, 0xda, 0x5d, 0x43, 0xe9, 0xb7, 0x2f, 0x4e, 0xfa, 0xc7, 0x4b, 0x6f, 0xe3, 0x6f, 0x7e, 0xf3,
	0xb8, 0x3c, 0xb0, 0x6a, 0x55, 0x49, 0x8d, 0x49, 0xce, 0x86, 0xb4, 0xec, 0xb4, 0xea, 0x15, 0x4e,
	0x19, 0x05, 0xd7, 0x95, 0x52, 0x55, 0x43, 0x09, 0xc3, 0xe3, 0xf9, 0x35, 0xa1, 0xb6, 0xeb, 0x3f,
	0x2e, 0x4e, 0xf4, 0x06, 0x62, 0x6f, 0x48, 0xe3, 0x02, 0x9c, 0xba, 0xf4, 0xed, 0xd0, 0x8e, 0x67,
	0x85, 0x53, 0x97, 0x88, 0x20, 0x4e, 0x87, 0x96, 0x7c, 0x87, 0x09, 0xcf, 0x18, 0x80, 0xd7, 0x1d,
	0x8c, 0x79, 0x57, 0xba, 0xf4, 0x27, 0xcc, 0x7f, 0xbe, 0xf1, 0x16, 0xdc, 0x8a, 0x4e, 0x25, 0x69,
	0x5f, 0x84, 0x76, 0xbc, 0xc8, 0x56, 0x92, 0x7f, 0x2e, 0x87, 0xf0, 0x27, 0x5e, 0x14, 0x5f, 0x42,
	0x74, 0x0f, 0xde, 0x40, 0xb7, 0xb5, 0xe9, 0xf1, 0x06, 0x44, 0x53, 0x9b, 0xde, 0xb7, 0xc3, 0x49,
	0x3c, 0xcf, 0xe6, 0xa3, 0xa3, 0x82, 0x17, 0x77, 0x29, 0xc0, 0x6f, 0x04, 0xfe, 0x83, 0xd9, 0x7e,
	0xb7, 0xce, 0x37, 0xcf, 0xbb, 0x7c, 0xbd, 0xb4, 0xd0, 0x03, 0xf1, 0xf2, 0xb8, 0xcd, 0x97, 0x36,
	0x02, 0xb8, 0x9b, 0x9c, 0x67, 0x27, 0x6b, 0x61, 0x3a, 0x9c, 0x18, 0x4c, 0xc1, 0x2b, 0xa8, 0xaa,
	0x4d, 0x4f, 0x1a, 0xc7, 0xd1, 0xc1, 0x95, 0xbc, 0x54, 0x22, 0xbf, 0x2b, 0x91, 0xf9, 0x50, 0x49,
	0x64, 0x61, 0x0a, 0x82, 0xdf, 0xf5, 0x87, 0x11, 0xfc, 0x1f, 0xc5, 0x0c, 0x62, 0x64, 0x1d, 0x5d,
	0x56, 0x1e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x16, 0x2f, 0x04, 0x81, 0x01, 0x00, 0x00,
}
