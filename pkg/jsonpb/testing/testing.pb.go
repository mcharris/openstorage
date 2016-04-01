// Code generated by protoc-gen-go.
// source: pkg/jsonpb/testing/testing.proto
// DO NOT EDIT!

package testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Status int32

const (
	Status_STATUS_NONE    Status = 0
	Status_STATUS_INIT    Status = 1
	Status_STATUS_OK      Status = 2
	Status_STATUS_OFFLINE Status = 3
	Status_STATUS_ERROR   Status = 4
)

var Status_name = map[int32]string{
	0: "STATUS_NONE",
	1: "STATUS_INIT",
	2: "STATUS_OK",
	3: "STATUS_OFFLINE",
	4: "STATUS_ERROR",
}
var Status_value = map[string]int32{
	"STATUS_NONE":    0,
	"STATUS_INIT":    1,
	"STATUS_OK":      2,
	"STATUS_OFFLINE": 3,
	"STATUS_ERROR":   4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Foo struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Status    Status                     `protobuf:"varint,2,opt,name=status,enum=testing.Status" json:"status,omitempty"`
}

func (m *Foo) Reset()                    { *m = Foo{} }
func (m *Foo) String() string            { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()               {}
func (*Foo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Foo) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Foo)(nil), "testing.Foo")
	proto.RegisterEnum("testing.Status", Status_name, Status_value)
}

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc8, 0x4e, 0xd7,
	0xcf, 0x2a, 0xce, 0xcf, 0x2b, 0x48, 0xd2, 0x2f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x87, 0xd1,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x94, 0x7c, 0x7a, 0x7e, 0x7e, 0x7a,
	0x4e, 0xaa, 0x3e, 0x58, 0x38, 0xa9, 0x34, 0x4d, 0xbf, 0x24, 0x33, 0x17, 0x28, 0x95, 0x98, 0x5b,
	0x00, 0x51, 0xa9, 0x94, 0xc1, 0xc5, 0xec, 0x96, 0x9f, 0x2f, 0x64, 0xc1, 0xc5, 0x09, 0x97, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0xe8, 0xd5, 0x83, 0xe9, 0xd5, 0x0b, 0x81,
	0xa9, 0x08, 0x42, 0x28, 0x16, 0x52, 0xe7, 0x62, 0x03, 0x32, 0x4a, 0x4a, 0x8b, 0x25, 0x98, 0x80,
	0xda, 0xf8, 0x8c, 0xf8, 0xf5, 0x60, 0x4e, 0x09, 0x06, 0x0b, 0x07, 0x41, 0xa5, 0xb5, 0xe2, 0xb9,
	0xd8, 0x20, 0x22, 0x42, 0xfc, 0x5c, 0xdc, 0xc1, 0x21, 0x8e, 0x21, 0xa1, 0xc1, 0xf1, 0x7e, 0xfe,
	0x7e, 0xae, 0x02, 0x0c, 0x48, 0x02, 0x9e, 0x7e, 0x9e, 0x21, 0x02, 0x8c, 0x42, 0xbc, 0x5c, 0x9c,
	0x50, 0x01, 0x7f, 0x6f, 0x01, 0x26, 0x21, 0x21, 0x2e, 0x3e, 0x18, 0xd7, 0xcd, 0xcd, 0xc7, 0x13,
	0xa8, 0x87, 0x59, 0x48, 0x80, 0x8b, 0x07, 0x2a, 0xe6, 0x1a, 0x14, 0xe4, 0x1f, 0x24, 0xc0, 0x92,
	0xc4, 0x06, 0x76, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x99, 0x1f, 0xa5, 0x1f, 0x01,
	0x00, 0x00,
}
