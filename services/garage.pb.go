// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/garage.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	math "math"
	models "models"
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

type GarageUserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GarageUserId) Reset()         { *m = GarageUserId{} }
func (m *GarageUserId) String() string { return proto.CompactTextString(m) }
func (*GarageUserId) ProtoMessage()    {}
func (*GarageUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8eddb3c5524f29, []int{0}
}

func (m *GarageUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageUserId.Unmarshal(m, b)
}
func (m *GarageUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageUserId.Marshal(b, m, deterministic)
}
func (m *GarageUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageUserId.Merge(m, src)
}
func (m *GarageUserId) XXX_Size() int {
	return xxx_messageInfo_GarageUserId.Size(m)
}
func (m *GarageUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageUserId.DiscardUnknown(m)
}

var xxx_messageInfo_GarageUserId proto.InternalMessageInfo

func (m *GarageUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GarageAndUserId struct {
	UserId               string         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Garage               *models.Garage `protobuf:"bytes,2,opt,name=garage,proto3" json:"garage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GarageAndUserId) Reset()         { *m = GarageAndUserId{} }
func (m *GarageAndUserId) String() string { return proto.CompactTextString(m) }
func (*GarageAndUserId) ProtoMessage()    {}
func (*GarageAndUserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8eddb3c5524f29, []int{1}
}

func (m *GarageAndUserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarageAndUserId.Unmarshal(m, b)
}
func (m *GarageAndUserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarageAndUserId.Marshal(b, m, deterministic)
}
func (m *GarageAndUserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarageAndUserId.Merge(m, src)
}
func (m *GarageAndUserId) XXX_Size() int {
	return xxx_messageInfo_GarageAndUserId.Size(m)
}
func (m *GarageAndUserId) XXX_DiscardUnknown() {
	xxx_messageInfo_GarageAndUserId.DiscardUnknown(m)
}

var xxx_messageInfo_GarageAndUserId proto.InternalMessageInfo

func (m *GarageAndUserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GarageAndUserId) GetGarage() *models.Garage {
	if m != nil {
		return m.Garage
	}
	return nil
}

func init() {
	proto.RegisterType((*GarageUserId)(nil), "services.GarageUserId")
	proto.RegisterType((*GarageAndUserId)(nil), "services.GarageAndUserId")
}

func init() { proto.RegisterFile("services/garage.proto", fileDescriptor_2a8eddb3c5524f29) }

var fileDescriptor_2a8eddb3c5524f29 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0x4f, 0x2c, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x80, 0x09, 0x4b, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0xc5,
	0x93, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0xca, 0xa4, 0x84, 0x73, 0xf3, 0x53,
	0x52, 0x73, 0x50, 0xf5, 0x2a, 0xa9, 0x73, 0xf1, 0xb8, 0x83, 0xf9, 0xa1, 0xc5, 0xa9, 0x45, 0x9e,
	0x29, 0x42, 0xe2, 0x5c, 0xec, 0xa5, 0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x6c, 0xa5, 0x60, 0x09, 0xa5, 0x20, 0x2e, 0x7e, 0x88, 0x42, 0xc7, 0xbc, 0x14,
	0x02, 0x6a, 0x85, 0xd4, 0xb8, 0xd8, 0x20, 0x96, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0xf1,
	0xe9, 0x41, 0xac, 0xd6, 0x83, 0x98, 0x10, 0x04, 0x95, 0x35, 0xaa, 0xe6, 0x62, 0x87, 0x88, 0x14,
	0x0b, 0x99, 0x70, 0xb1, 0xf8, 0x64, 0x16, 0x97, 0x08, 0x89, 0xe9, 0xc1, 0x3c, 0xa3, 0x87, 0xec,
	0x2e, 0x29, 0x21, 0x54, 0x23, 0x40, 0x6a, 0x95, 0x18, 0x84, 0xac, 0xb8, 0x98, 0x1d, 0x53, 0x52,
	0x84, 0x24, 0xd1, 0x35, 0xc1, 0xdd, 0x28, 0x25, 0xa6, 0x07, 0x09, 0x12, 0x3d, 0x58, 0x90, 0xe8,
	0xb9, 0x82, 0x82, 0x44, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x51, 0xbe, 0x7e, 0x55, 0x01, 0x00, 0x00,
}