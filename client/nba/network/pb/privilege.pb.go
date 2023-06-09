// Code generated by protoc-gen-go.
// source: privilege.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetPrivilegeAwardC2S_MsgID int32

const (
	GetPrivilegeAwardC2S_eMsgID GetPrivilegeAwardC2S_MsgID = 15617
)

var GetPrivilegeAwardC2S_MsgID_name = map[int32]string{
	15617: "eMsgID",
}
var GetPrivilegeAwardC2S_MsgID_value = map[string]int32{
	"eMsgID": 15617,
}

func (x GetPrivilegeAwardC2S_MsgID) Enum() *GetPrivilegeAwardC2S_MsgID {
	p := new(GetPrivilegeAwardC2S_MsgID)
	*p = x
	return p
}
func (x GetPrivilegeAwardC2S_MsgID) String() string {
	return proto.EnumName(GetPrivilegeAwardC2S_MsgID_name, int32(x))
}
func (x *GetPrivilegeAwardC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetPrivilegeAwardC2S_MsgID_value, data, "GetPrivilegeAwardC2S_MsgID")
	if err != nil {
		return err
	}
	*x = GetPrivilegeAwardC2S_MsgID(value)
	return nil
}
func (GetPrivilegeAwardC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor52, []int{1, 0}
}

type GetPrivilegeAwardS2C_MsgID int32

const (
	GetPrivilegeAwardS2C_eMsgID GetPrivilegeAwardS2C_MsgID = 15617
)

var GetPrivilegeAwardS2C_MsgID_name = map[int32]string{
	15617: "eMsgID",
}
var GetPrivilegeAwardS2C_MsgID_value = map[string]int32{
	"eMsgID": 15617,
}

func (x GetPrivilegeAwardS2C_MsgID) Enum() *GetPrivilegeAwardS2C_MsgID {
	p := new(GetPrivilegeAwardS2C_MsgID)
	*p = x
	return p
}
func (x GetPrivilegeAwardS2C_MsgID) String() string {
	return proto.EnumName(GetPrivilegeAwardS2C_MsgID_name, int32(x))
}
func (x *GetPrivilegeAwardS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetPrivilegeAwardS2C_MsgID_value, data, "GetPrivilegeAwardS2C_MsgID")
	if err != nil {
		return err
	}
	*x = GetPrivilegeAwardS2C_MsgID(value)
	return nil
}
func (GetPrivilegeAwardS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor52, []int{2, 0}
}

type PrivilegeDataS2C_MsgID int32

const (
	PrivilegeDataS2C_eMsgID PrivilegeDataS2C_MsgID = 15618
)

var PrivilegeDataS2C_MsgID_name = map[int32]string{
	15618: "eMsgID",
}
var PrivilegeDataS2C_MsgID_value = map[string]int32{
	"eMsgID": 15618,
}

func (x PrivilegeDataS2C_MsgID) Enum() *PrivilegeDataS2C_MsgID {
	p := new(PrivilegeDataS2C_MsgID)
	*p = x
	return p
}
func (x PrivilegeDataS2C_MsgID) String() string {
	return proto.EnumName(PrivilegeDataS2C_MsgID_name, int32(x))
}
func (x *PrivilegeDataS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PrivilegeDataS2C_MsgID_value, data, "PrivilegeDataS2C_MsgID")
	if err != nil {
		return err
	}
	*x = PrivilegeDataS2C_MsgID(value)
	return nil
}
func (PrivilegeDataS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor52, []int{3, 0} }

// 订阅特权
type Privilege struct {
	// 月卡到期时间
	EndTime *int64 `protobuf:"varint,1,opt,name=endTime" json:"endTime,omitempty"`
	// 下次刷新时间
	NextResetTime *int64 `protobuf:"varint,2,opt,name=nextResetTime" json:"nextResetTime,omitempty"`
	// 购买的特权类型id
	Id *int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	// 今日领奖标记 true已领取
	AwardFlag *bool `protobuf:"varint,4,opt,name=awardFlag" json:"awardFlag,omitempty"`
	// 已经发送续订提醒邮件天数
	BeforeSendRenewMailDays *int32 `protobuf:"varint,5,opt,name=beforeSendRenewMailDays" json:"beforeSendRenewMailDays,omitempty"`
	XXX_unrecognized        []byte `json:"-"`
}

func (m *Privilege) Reset()                    { *m = Privilege{} }
func (m *Privilege) String() string            { return proto.CompactTextString(m) }
func (*Privilege) ProtoMessage()               {}
func (*Privilege) Descriptor() ([]byte, []int) { return fileDescriptor52, []int{0} }

func (m *Privilege) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *Privilege) GetNextResetTime() int64 {
	if m != nil && m.NextResetTime != nil {
		return *m.NextResetTime
	}
	return 0
}

func (m *Privilege) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Privilege) GetAwardFlag() bool {
	if m != nil && m.AwardFlag != nil {
		return *m.AwardFlag
	}
	return false
}

func (m *Privilege) GetBeforeSendRenewMailDays() int32 {
	if m != nil && m.BeforeSendRenewMailDays != nil {
		return *m.BeforeSendRenewMailDays
	}
	return 0
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~订阅特权~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
// 订阅特权 领取每日奖励
type GetPrivilegeAwardC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetPrivilegeAwardC2S) Reset()                    { *m = GetPrivilegeAwardC2S{} }
func (m *GetPrivilegeAwardC2S) String() string            { return proto.CompactTextString(m) }
func (*GetPrivilegeAwardC2S) ProtoMessage()               {}
func (*GetPrivilegeAwardC2S) Descriptor() ([]byte, []int) { return fileDescriptor52, []int{1} }

// 返回订阅特权 领取每日奖励
type GetPrivilegeAwardS2C struct {
	AwardFlag        *bool  `protobuf:"varint,1,opt,name=awardFlag" json:"awardFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetPrivilegeAwardS2C) Reset()                    { *m = GetPrivilegeAwardS2C{} }
func (m *GetPrivilegeAwardS2C) String() string            { return proto.CompactTextString(m) }
func (*GetPrivilegeAwardS2C) ProtoMessage()               {}
func (*GetPrivilegeAwardS2C) Descriptor() ([]byte, []int) { return fileDescriptor52, []int{2} }

func (m *GetPrivilegeAwardS2C) GetAwardFlag() bool {
	if m != nil && m.AwardFlag != nil {
		return *m.AwardFlag
	}
	return false
}

// 登录下发订阅特权数据
type PrivilegeDataS2C struct {
	Data             *Privilege `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *PrivilegeDataS2C) Reset()                    { *m = PrivilegeDataS2C{} }
func (m *PrivilegeDataS2C) String() string            { return proto.CompactTextString(m) }
func (*PrivilegeDataS2C) ProtoMessage()               {}
func (*PrivilegeDataS2C) Descriptor() ([]byte, []int) { return fileDescriptor52, []int{3} }

func (m *PrivilegeDataS2C) GetData() *Privilege {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Privilege)(nil), "Privilege")
	proto.RegisterType((*GetPrivilegeAwardC2S)(nil), "GetPrivilegeAwardC2S")
	proto.RegisterType((*GetPrivilegeAwardS2C)(nil), "GetPrivilegeAwardS2C")
	proto.RegisterType((*PrivilegeDataS2C)(nil), "PrivilegeDataS2C")
	proto.RegisterEnum("GetPrivilegeAwardC2S_MsgID", GetPrivilegeAwardC2S_MsgID_name, GetPrivilegeAwardC2S_MsgID_value)
	proto.RegisterEnum("GetPrivilegeAwardS2C_MsgID", GetPrivilegeAwardS2C_MsgID_name, GetPrivilegeAwardS2C_MsgID_value)
	proto.RegisterEnum("PrivilegeDataS2C_MsgID", PrivilegeDataS2C_MsgID_name, PrivilegeDataS2C_MsgID_value)
}

func init() { proto.RegisterFile("privilege.proto", fileDescriptor52) }

var fileDescriptor52 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2c,
	0xcb, 0xcc, 0x49, 0x4d, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x49, 0xce, 0xcf,
	0xcd, 0xcd, 0xcf, 0x83, 0xf0, 0x94, 0x8a, 0xb9, 0x38, 0x03, 0x60, 0x0a, 0x84, 0xf8, 0xb9, 0xd8,
	0x53, 0xf3, 0x52, 0x42, 0x32, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x85, 0x44, 0xb9,
	0x78, 0xf3, 0x52, 0x2b, 0x4a, 0x82, 0x52, 0x8b, 0x53, 0x4b, 0xc0, 0xc2, 0x4c, 0x60, 0x61, 0x2e,
	0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x66, 0x20, 0x9b, 0x55, 0x48, 0x90, 0x8b, 0x33, 0xb1, 0x3c, 0xb1,
	0x28, 0xc5, 0x2d, 0x27, 0x31, 0x5d, 0x82, 0x05, 0x28, 0xc4, 0x21, 0x24, 0xcf, 0x25, 0x9e, 0x94,
	0x9a, 0x96, 0x5f, 0x94, 0x1a, 0x0c, 0x34, 0x2d, 0x28, 0x35, 0x2f, 0xb5, 0xdc, 0x37, 0x31, 0x33,
	0xc7, 0x25, 0xb1, 0xb2, 0x58, 0x82, 0x15, 0xa4, 0x47, 0x49, 0x87, 0x4b, 0xc4, 0x3d, 0xb5, 0x04,
	0x6e, 0xaf, 0x23, 0x48, 0xbf, 0xb3, 0x51, 0xb0, 0x92, 0x08, 0x17, 0xab, 0x6f, 0x71, 0xba, 0xa7,
	0x8b, 0x10, 0x37, 0x17, 0x5b, 0x2a, 0x98, 0x25, 0xd0, 0x58, 0xa5, 0x64, 0x8f, 0x45, 0x75, 0xb0,
	0x91, 0x33, 0xaa, 0xcd, 0x20, 0xf7, 0x72, 0xe0, 0x30, 0xc0, 0x89, 0x4b, 0x00, 0xae, 0xdb, 0x25,
	0xb1, 0x24, 0x11, 0xa4, 0x59, 0x82, 0x8b, 0x25, 0x05, 0xc8, 0x04, 0xeb, 0xe3, 0x36, 0xe2, 0xd2,
	0x83, 0x2b, 0xc0, 0x66, 0x46, 0x53, 0x15, 0x20, 0x00, 0x00, 0xff, 0xff, 0x28, 0x9f, 0x3b, 0xe4,
	0x47, 0x01, 0x00, 0x00,
}
