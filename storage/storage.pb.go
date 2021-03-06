// Code generated by protoc-gen-go.
// source: storage/storage.proto
// DO NOT EDIT!

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	storage/storage.proto

It has these top-level messages:
	NodeIDProto
*/
package storage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// NodeIDProto is the serialized form of NodeID. It's used only for persistence in storage.
// As this is long-term we prefer not to use a Go specific format.
type NodeIDProto struct {
	Path             []byte `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	PrefixLenBits    *int32 `protobuf:"varint,2,opt,name=prefix_len_bits" json:"prefix_len_bits,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NodeIDProto) Reset()                    { *m = NodeIDProto{} }
func (m *NodeIDProto) String() string            { return proto.CompactTextString(m) }
func (*NodeIDProto) ProtoMessage()               {}
func (*NodeIDProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NodeIDProto) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *NodeIDProto) GetPrefixLenBits() int32 {
	if m != nil && m.PrefixLenBits != nil {
		return *m.PrefixLenBits
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeIDProto)(nil), "storage.NodeIDProto")
}

var fileDescriptor0 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x09, 0x17, 0xb7, 0x5f, 0x7e, 0x4a, 0xaa, 0xa7, 0x4b, 0x00, 0x58, 0x9c, 0x87, 0x8b, 0xa5,
	0x20, 0xb1, 0x24, 0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0x47, 0x48, 0x9c, 0x8b, 0xbf, 0xa0, 0x28,
	0x35, 0x2d, 0xb3, 0x22, 0x3e, 0x27, 0x35, 0x2f, 0x3e, 0x29, 0xb3, 0xa4, 0x58, 0x82, 0x09, 0x28,
	0xc1, 0x0a, 0x08, 0x00, 0x00, 0xff, 0xff, 0x30, 0x3b, 0x1f, 0xdb, 0x56, 0x00, 0x00, 0x00,
}
