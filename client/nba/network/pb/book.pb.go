// Code generated by protoc-gen-go.
// source: book.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ActiveBookC2S_MsgID int32

const (
	ActiveBookC2S_eMsgID ActiveBookC2S_MsgID = 13825
)

var ActiveBookC2S_MsgID_name = map[int32]string{
	13825: "eMsgID",
}
var ActiveBookC2S_MsgID_value = map[string]int32{
	"eMsgID": 13825,
}

func (x ActiveBookC2S_MsgID) Enum() *ActiveBookC2S_MsgID {
	p := new(ActiveBookC2S_MsgID)
	*p = x
	return p
}
func (x ActiveBookC2S_MsgID) String() string {
	return proto.EnumName(ActiveBookC2S_MsgID_name, int32(x))
}
func (x *ActiveBookC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ActiveBookC2S_MsgID_value, data, "ActiveBookC2S_MsgID")
	if err != nil {
		return err
	}
	*x = ActiveBookC2S_MsgID(value)
	return nil
}
func (ActiveBookC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1, 0} }

type ActiveBookS2C_MsgID int32

const (
	ActiveBookS2C_eMsgID ActiveBookS2C_MsgID = 13825
)

var ActiveBookS2C_MsgID_name = map[int32]string{
	13825: "eMsgID",
}
var ActiveBookS2C_MsgID_value = map[string]int32{
	"eMsgID": 13825,
}

func (x ActiveBookS2C_MsgID) Enum() *ActiveBookS2C_MsgID {
	p := new(ActiveBookS2C_MsgID)
	*p = x
	return p
}
func (x ActiveBookS2C_MsgID) String() string {
	return proto.EnumName(ActiveBookS2C_MsgID_name, int32(x))
}
func (x *ActiveBookS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ActiveBookS2C_MsgID_value, data, "ActiveBookS2C_MsgID")
	if err != nil {
		return err
	}
	*x = ActiveBookS2C_MsgID(value)
	return nil
}
func (ActiveBookS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{2, 0} }

type BookListS2C_MsgID int32

const (
	BookListS2C_eMsgID BookListS2C_MsgID = 13826
)

var BookListS2C_MsgID_name = map[int32]string{
	13826: "eMsgID",
}
var BookListS2C_MsgID_value = map[string]int32{
	"eMsgID": 13826,
}

func (x BookListS2C_MsgID) Enum() *BookListS2C_MsgID {
	p := new(BookListS2C_MsgID)
	*p = x
	return p
}
func (x BookListS2C_MsgID) String() string {
	return proto.EnumName(BookListS2C_MsgID_name, int32(x))
}
func (x *BookListS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BookListS2C_MsgID_value, data, "BookListS2C_MsgID")
	if err != nil {
		return err
	}
	*x = BookListS2C_MsgID(value)
	return nil
}
func (BookListS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{3, 0} }

// 图鉴
type BookData struct {
	Coachs           []int32 `protobuf:"varint,1,rep,name=coachs" json:"coachs,omitempty"`
	Mascots          []int32 `protobuf:"varint,2,rep,name=mascots" json:"mascots,omitempty"`
	Titles           []int32 `protobuf:"varint,3,rep,name=titles" json:"titles,omitempty"`
	FashionDresses   []int32 `protobuf:"varint,4,rep,name=fashionDresses" json:"fashionDresses,omitempty"`
	Sneakers         []int32 `protobuf:"varint,5,rep,name=sneakers" json:"sneakers,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BookData) Reset()                    { *m = BookData{} }
func (m *BookData) String() string            { return proto.CompactTextString(m) }
func (*BookData) ProtoMessage()               {}
func (*BookData) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *BookData) GetCoachs() []int32 {
	if m != nil {
		return m.Coachs
	}
	return nil
}

func (m *BookData) GetMascots() []int32 {
	if m != nil {
		return m.Mascots
	}
	return nil
}

func (m *BookData) GetTitles() []int32 {
	if m != nil {
		return m.Titles
	}
	return nil
}

func (m *BookData) GetFashionDresses() []int32 {
	if m != nil {
		return m.FashionDresses
	}
	return nil
}

func (m *BookData) GetSneakers() []int32 {
	if m != nil {
		return m.Sneakers
	}
	return nil
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
// 激活教练图鉴
type ActiveBookC2S struct {
	Type             *int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Id               *int32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActiveBookC2S) Reset()                    { *m = ActiveBookC2S{} }
func (m *ActiveBookC2S) String() string            { return proto.CompactTextString(m) }
func (*ActiveBookC2S) ProtoMessage()               {}
func (*ActiveBookC2S) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *ActiveBookC2S) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ActiveBookC2S) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

// 返回 激活教练图鉴
type ActiveBookS2C struct {
	Type             *int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Id               *int32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActiveBookS2C) Reset()                    { *m = ActiveBookS2C{} }
func (m *ActiveBookS2C) String() string            { return proto.CompactTextString(m) }
func (*ActiveBookS2C) ProtoMessage()               {}
func (*ActiveBookS2C) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *ActiveBookS2C) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ActiveBookS2C) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

// 登录下发图鉴信息
type BookListS2C struct {
	Book             *BookData `protobuf:"bytes,1,opt,name=book" json:"book,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BookListS2C) Reset()                    { *m = BookListS2C{} }
func (m *BookListS2C) String() string            { return proto.CompactTextString(m) }
func (*BookListS2C) ProtoMessage()               {}
func (*BookListS2C) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *BookListS2C) GetBook() *BookData {
	if m != nil {
		return m.Book
	}
	return nil
}

func init() {
	proto.RegisterType((*BookData)(nil), "BookData")
	proto.RegisterType((*ActiveBookC2S)(nil), "ActiveBookC2S")
	proto.RegisterType((*ActiveBookS2C)(nil), "ActiveBookS2C")
	proto.RegisterType((*BookListS2C)(nil), "BookListS2C")
	proto.RegisterEnum("ActiveBookC2S_MsgID", ActiveBookC2S_MsgID_name, ActiveBookC2S_MsgID_value)
	proto.RegisterEnum("ActiveBookS2C_MsgID", ActiveBookS2C_MsgID_name, ActiveBookS2C_MsgID_value)
	proto.RegisterEnum("BookListS2C_MsgID", BookListS2C_MsgID_name, BookListS2C_MsgID_value)
}

func init() { proto.RegisterFile("book.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x90, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x86, 0xb9, 0xbb, 0xe4, 0x3c, 0x27, 0xe7, 0x79, 0x2c, 0xa2, 0x8b, 0x95, 0xa4, 0xb2, 0x4a,
	0x91, 0x5a, 0x10, 0x4d, 0x1a, 0x41, 0xab, 0xfc, 0x82, 0x75, 0x1d, 0xcd, 0x92, 0x8f, 0x09, 0x99,
	0x41, 0xb0, 0xd4, 0x5f, 0xee, 0xee, 0x8a, 0x85, 0x60, 0x73, 0xdd, 0xcc, 0xb3, 0xfb, 0xcc, 0x0b,
	0x2f, 0xc0, 0x33, 0x51, 0x57, 0x4c, 0x33, 0x09, 0x5d, 0x6e, 0x2d, 0x0d, 0x03, 0x8d, 0x3f, 0x5b,
	0x8e, 0xb0, 0xb9, 0xf7, 0x6f, 0xb5, 0x11, 0xa3, 0x76, 0xb0, 0xb6, 0x64, 0x6c, 0xcb, 0x7a, 0x71,
	0xb5, 0xba, 0x4e, 0xd5, 0x29, 0x1c, 0x0d, 0x86, 0x2d, 0x09, 0xeb, 0x65, 0x04, 0xfe, 0x83, 0x38,
	0xe9, 0x91, 0xf5, 0x2a, 0xee, 0xe7, 0xb0, 0x7b, 0x35, 0xdc, 0x3a, 0x1a, 0xeb, 0x19, 0x99, 0x3d,
	0x4f, 0x22, 0xdf, 0xc3, 0x86, 0x47, 0x34, 0x1d, 0xce, 0xac, 0xd3, 0x40, 0xf2, 0x5b, 0x38, 0xb9,
	0xb3, 0xe2, 0xde, 0x31, 0x84, 0x55, 0x65, 0xa3, 0xb6, 0x90, 0xc8, 0xc7, 0x84, 0x3e, 0x69, 0xe1,
	0x05, 0x80, 0xa5, 0x7b, 0xf1, 0x21, 0x7e, 0xce, 0xcf, 0x20, 0x7d, 0xe2, 0xb7, 0x87, 0x5a, 0x65,
	0xb0, 0xc6, 0x38, 0xed, 0x3f, 0xfb, 0xbf, 0x07, 0x9a, 0xb2, 0x3a, 0xf8, 0xc0, 0x0d, 0x64, 0x41,
	0x7d, 0x74, 0x2c, 0x41, 0xbf, 0x80, 0x24, 0x74, 0x12, 0xf5, 0xac, 0x3c, 0x2e, 0x7e, 0x4b, 0xf8,
	0xcf, 0xfe, 0xea, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0x20, 0xdb, 0x4c, 0x41, 0x01, 0x00,
	0x00,
}