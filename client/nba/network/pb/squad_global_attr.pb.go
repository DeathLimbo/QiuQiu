// Code generated by protoc-gen-go.
// source: squad_global_attr.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 阵容全局属性
type SquadGlobalAttr struct {
	AttrValues       []*PlayerAttrValue `protobuf:"bytes,1,rep,name=attrValues" json:"attrValues,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SquadGlobalAttr) Reset()                    { *m = SquadGlobalAttr{} }
func (m *SquadGlobalAttr) String() string            { return proto.CompactTextString(m) }
func (*SquadGlobalAttr) ProtoMessage()               {}
func (*SquadGlobalAttr) Descriptor() ([]byte, []int) { return fileDescriptor71, []int{0} }

func (m *SquadGlobalAttr) GetAttrValues() []*PlayerAttrValue {
	if m != nil {
		return m.AttrValues
	}
	return nil
}

func init() {
	proto.RegisterType((*SquadGlobalAttr)(nil), "SquadGlobalAttr")
}

func init() { proto.RegisterFile("squad_global_attr.proto", fileDescriptor71) }

var fileDescriptor71 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x2e, 0x2c, 0x4d,
	0x4c, 0x89, 0x4f, 0xcf, 0xc9, 0x4f, 0x4a, 0xcc, 0x89, 0x4f, 0x2c, 0x29, 0x29, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xce, 0x4d, 0x4c, 0x2e, 0xca, 0x87, 0x72, 0xb8, 0x92, 0xf3, 0xf3,
	0xd2, 0x20, 0x6c, 0x25, 0x73, 0x2e, 0xfe, 0x60, 0x90, 0x1e, 0x77, 0xb0, 0x16, 0x47, 0xa0, 0x0e,
	0x21, 0x15, 0x2e, 0x2e, 0x90, 0xce, 0xb0, 0xc4, 0x9c, 0xd2, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66,
	0x0d, 0x6e, 0x23, 0x01, 0xbd, 0x80, 0x9c, 0xc4, 0xca, 0xd4, 0x22, 0x47, 0x98, 0x04, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0xeb, 0xb2, 0x37, 0x6b, 0x00, 0x00, 0x00,
}
