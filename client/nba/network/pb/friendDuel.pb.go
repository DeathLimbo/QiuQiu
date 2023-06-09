// Code generated by protoc-gen-go.
// source: friendDuel.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DuelInviteSendC2S_MsgID int32

const (
	DuelInviteSendC2S_eMsgID DuelInviteSendC2S_MsgID = 1559
)

var DuelInviteSendC2S_MsgID_name = map[int32]string{
	1559: "eMsgID",
}
var DuelInviteSendC2S_MsgID_value = map[string]int32{
	"eMsgID": 1559,
}

func (x DuelInviteSendC2S_MsgID) Enum() *DuelInviteSendC2S_MsgID {
	p := new(DuelInviteSendC2S_MsgID)
	*p = x
	return p
}
func (x DuelInviteSendC2S_MsgID) String() string {
	return proto.EnumName(DuelInviteSendC2S_MsgID_name, int32(x))
}
func (x *DuelInviteSendC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelInviteSendC2S_MsgID_value, data, "DuelInviteSendC2S_MsgID")
	if err != nil {
		return err
	}
	*x = DuelInviteSendC2S_MsgID(value)
	return nil
}
func (DuelInviteSendC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{0, 0} }

type DuelInviteSendS2C_MsgID int32

const (
	DuelInviteSendS2C_eMsgID DuelInviteSendS2C_MsgID = 1559
)

var DuelInviteSendS2C_MsgID_name = map[int32]string{
	1559: "eMsgID",
}
var DuelInviteSendS2C_MsgID_value = map[string]int32{
	"eMsgID": 1559,
}

func (x DuelInviteSendS2C_MsgID) Enum() *DuelInviteSendS2C_MsgID {
	p := new(DuelInviteSendS2C_MsgID)
	*p = x
	return p
}
func (x DuelInviteSendS2C_MsgID) String() string {
	return proto.EnumName(DuelInviteSendS2C_MsgID_name, int32(x))
}
func (x *DuelInviteSendS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelInviteSendS2C_MsgID_value, data, "DuelInviteSendS2C_MsgID")
	if err != nil {
		return err
	}
	*x = DuelInviteSendS2C_MsgID(value)
	return nil
}
func (DuelInviteSendS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{1, 0} }

type DuelInviteRecvS2C_MsgID int32

const (
	DuelInviteRecvS2C_eMsgID DuelInviteRecvS2C_MsgID = 1560
)

var DuelInviteRecvS2C_MsgID_name = map[int32]string{
	1560: "eMsgID",
}
var DuelInviteRecvS2C_MsgID_value = map[string]int32{
	"eMsgID": 1560,
}

func (x DuelInviteRecvS2C_MsgID) Enum() *DuelInviteRecvS2C_MsgID {
	p := new(DuelInviteRecvS2C_MsgID)
	*p = x
	return p
}
func (x DuelInviteRecvS2C_MsgID) String() string {
	return proto.EnumName(DuelInviteRecvS2C_MsgID_name, int32(x))
}
func (x *DuelInviteRecvS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelInviteRecvS2C_MsgID_value, data, "DuelInviteRecvS2C_MsgID")
	if err != nil {
		return err
	}
	*x = DuelInviteRecvS2C_MsgID(value)
	return nil
}
func (DuelInviteRecvS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{2, 0} }

type DuelAgreeC2S_MsgID int32

const (
	DuelAgreeC2S_eMsgID DuelAgreeC2S_MsgID = 1561
)

var DuelAgreeC2S_MsgID_name = map[int32]string{
	1561: "eMsgID",
}
var DuelAgreeC2S_MsgID_value = map[string]int32{
	"eMsgID": 1561,
}

func (x DuelAgreeC2S_MsgID) Enum() *DuelAgreeC2S_MsgID {
	p := new(DuelAgreeC2S_MsgID)
	*p = x
	return p
}
func (x DuelAgreeC2S_MsgID) String() string {
	return proto.EnumName(DuelAgreeC2S_MsgID_name, int32(x))
}
func (x *DuelAgreeC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelAgreeC2S_MsgID_value, data, "DuelAgreeC2S_MsgID")
	if err != nil {
		return err
	}
	*x = DuelAgreeC2S_MsgID(value)
	return nil
}
func (DuelAgreeC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{3, 0} }

type DuelAgreeS2C_MsgID int32

const (
	DuelAgreeS2C_eMsgID DuelAgreeS2C_MsgID = 1561
)

var DuelAgreeS2C_MsgID_name = map[int32]string{
	1561: "eMsgID",
}
var DuelAgreeS2C_MsgID_value = map[string]int32{
	"eMsgID": 1561,
}

func (x DuelAgreeS2C_MsgID) Enum() *DuelAgreeS2C_MsgID {
	p := new(DuelAgreeS2C_MsgID)
	*p = x
	return p
}
func (x DuelAgreeS2C_MsgID) String() string {
	return proto.EnumName(DuelAgreeS2C_MsgID_name, int32(x))
}
func (x *DuelAgreeS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelAgreeS2C_MsgID_value, data, "DuelAgreeS2C_MsgID")
	if err != nil {
		return err
	}
	*x = DuelAgreeS2C_MsgID(value)
	return nil
}
func (DuelAgreeS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{4, 0} }

type DuelInviteRespS2C_MsgID int32

const (
	DuelInviteRespS2C_eMsgID DuelInviteRespS2C_MsgID = 1562
)

var DuelInviteRespS2C_MsgID_name = map[int32]string{
	1562: "eMsgID",
}
var DuelInviteRespS2C_MsgID_value = map[string]int32{
	"eMsgID": 1562,
}

func (x DuelInviteRespS2C_MsgID) Enum() *DuelInviteRespS2C_MsgID {
	p := new(DuelInviteRespS2C_MsgID)
	*p = x
	return p
}
func (x DuelInviteRespS2C_MsgID) String() string {
	return proto.EnumName(DuelInviteRespS2C_MsgID_name, int32(x))
}
func (x *DuelInviteRespS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelInviteRespS2C_MsgID_value, data, "DuelInviteRespS2C_MsgID")
	if err != nil {
		return err
	}
	*x = DuelInviteRespS2C_MsgID(value)
	return nil
}
func (DuelInviteRespS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{5, 0} }

type DuelAgreeRespS2C_MsgID int32

const (
	DuelAgreeRespS2C_eMsgID DuelAgreeRespS2C_MsgID = 1563
)

var DuelAgreeRespS2C_MsgID_name = map[int32]string{
	1563: "eMsgID",
}
var DuelAgreeRespS2C_MsgID_value = map[string]int32{
	"eMsgID": 1563,
}

func (x DuelAgreeRespS2C_MsgID) Enum() *DuelAgreeRespS2C_MsgID {
	p := new(DuelAgreeRespS2C_MsgID)
	*p = x
	return p
}
func (x DuelAgreeRespS2C_MsgID) String() string {
	return proto.EnumName(DuelAgreeRespS2C_MsgID_name, int32(x))
}
func (x *DuelAgreeRespS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelAgreeRespS2C_MsgID_value, data, "DuelAgreeRespS2C_MsgID")
	if err != nil {
		return err
	}
	*x = DuelAgreeRespS2C_MsgID(value)
	return nil
}
func (DuelAgreeRespS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor28, []int{6, 0} }

type DuelInviteCancelC2S_MsgID int32

const (
	DuelInviteCancelC2S_eMsgID DuelInviteCancelC2S_MsgID = 1564
)

var DuelInviteCancelC2S_MsgID_name = map[int32]string{
	1564: "eMsgID",
}
var DuelInviteCancelC2S_MsgID_value = map[string]int32{
	"eMsgID": 1564,
}

func (x DuelInviteCancelC2S_MsgID) Enum() *DuelInviteCancelC2S_MsgID {
	p := new(DuelInviteCancelC2S_MsgID)
	*p = x
	return p
}
func (x DuelInviteCancelC2S_MsgID) String() string {
	return proto.EnumName(DuelInviteCancelC2S_MsgID_name, int32(x))
}
func (x *DuelInviteCancelC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelInviteCancelC2S_MsgID_value, data, "DuelInviteCancelC2S_MsgID")
	if err != nil {
		return err
	}
	*x = DuelInviteCancelC2S_MsgID(value)
	return nil
}
func (DuelInviteCancelC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor28, []int{7, 0}
}

type DuelInviteCancelS2C_MsgID int32

const (
	DuelInviteCancelS2C_eMsgID DuelInviteCancelS2C_MsgID = 1564
)

var DuelInviteCancelS2C_MsgID_name = map[int32]string{
	1564: "eMsgID",
}
var DuelInviteCancelS2C_MsgID_value = map[string]int32{
	"eMsgID": 1564,
}

func (x DuelInviteCancelS2C_MsgID) Enum() *DuelInviteCancelS2C_MsgID {
	p := new(DuelInviteCancelS2C_MsgID)
	*p = x
	return p
}
func (x DuelInviteCancelS2C_MsgID) String() string {
	return proto.EnumName(DuelInviteCancelS2C_MsgID_name, int32(x))
}
func (x *DuelInviteCancelS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DuelInviteCancelS2C_MsgID_value, data, "DuelInviteCancelS2C_MsgID")
	if err != nil {
		return err
	}
	*x = DuelInviteCancelS2C_MsgID(value)
	return nil
}
func (DuelInviteCancelS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor28, []int{8, 0}
}

// 玩家1发送邀请
type DuelInviteSendC2S struct {
	// 接收者gid
	Gid              *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInviteSendC2S) Reset()                    { *m = DuelInviteSendC2S{} }
func (m *DuelInviteSendC2S) String() string            { return proto.CompactTextString(m) }
func (*DuelInviteSendC2S) ProtoMessage()               {}
func (*DuelInviteSendC2S) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{0} }

func (m *DuelInviteSendC2S) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

// @Gamer
type DuelInviteSendS2C struct {
	// 接收者gid
	Gid *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	// 1:对方不在线，
	// 2：等待对方同意
	Result           *int32 `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInviteSendS2C) Reset()                    { *m = DuelInviteSendS2C{} }
func (m *DuelInviteSendS2C) String() string            { return proto.CompactTextString(m) }
func (*DuelInviteSendS2C) ProtoMessage()               {}
func (*DuelInviteSendS2C) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{1} }

func (m *DuelInviteSendS2C) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *DuelInviteSendS2C) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

// 玩家2收到邀请
// @Gamer
type DuelInviteRecvS2C struct {
	// 发起者gid
	Gid *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	// 发送时间
	SendTime         *int64 `protobuf:"varint,2,opt,name=sendTime" json:"sendTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInviteRecvS2C) Reset()                    { *m = DuelInviteRecvS2C{} }
func (m *DuelInviteRecvS2C) String() string            { return proto.CompactTextString(m) }
func (*DuelInviteRecvS2C) ProtoMessage()               {}
func (*DuelInviteRecvS2C) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{2} }

func (m *DuelInviteRecvS2C) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *DuelInviteRecvS2C) GetSendTime() int64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

// 玩家2同意切磋
type DuelAgreeC2S struct {
	// true：同意， false：拒绝
	Agree            *bool  `protobuf:"varint,2,opt,name=agree" json:"agree,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelAgreeC2S) Reset()                    { *m = DuelAgreeC2S{} }
func (m *DuelAgreeC2S) String() string            { return proto.CompactTextString(m) }
func (*DuelAgreeC2S) ProtoMessage()               {}
func (*DuelAgreeC2S) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{3} }

func (m *DuelAgreeC2S) GetAgree() bool {
	if m != nil && m.Agree != nil {
		return *m.Agree
	}
	return false
}

// 玩家2同意切磋
// @Gamer
type DuelAgreeS2C struct {
	// 1:对方不在线，
	// 2：等待结果
	Result           *int32 `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelAgreeS2C) Reset()                    { *m = DuelAgreeS2C{} }
func (m *DuelAgreeS2C) String() string            { return proto.CompactTextString(m) }
func (*DuelAgreeS2C) ProtoMessage()               {}
func (*DuelAgreeS2C) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{4} }

func (m *DuelAgreeS2C) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

// 玩家1收到结果
// @Gamer
type DuelInviteRespS2C struct {
	// 对手id
	Gid *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	// 是否同意
	Agree *bool `protobuf:"varint,2,opt,name=agree" json:"agree,omitempty"`
	// 不同意的原因
	// 1:玩家不在线，无法切磋
	// 2:玩家拒绝了你的切磋请求
	// 3:该玩家正在比赛中，无法被邀请
	// 4:玩家正在考虑是否与其他玩家切磋
	Reason           *int32 `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInviteRespS2C) Reset()                    { *m = DuelInviteRespS2C{} }
func (m *DuelInviteRespS2C) String() string            { return proto.CompactTextString(m) }
func (*DuelInviteRespS2C) ProtoMessage()               {}
func (*DuelInviteRespS2C) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{5} }

func (m *DuelInviteRespS2C) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *DuelInviteRespS2C) GetAgree() bool {
	if m != nil && m.Agree != nil {
		return *m.Agree
	}
	return false
}

func (m *DuelInviteRespS2C) GetReason() int32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

// 玩家2收到结果
// @Gamer
type DuelAgreeRespS2C struct {
	Gid              *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	Result           *bool  `protobuf:"varint,2,opt,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelAgreeRespS2C) Reset()                    { *m = DuelAgreeRespS2C{} }
func (m *DuelAgreeRespS2C) String() string            { return proto.CompactTextString(m) }
func (*DuelAgreeRespS2C) ProtoMessage()               {}
func (*DuelAgreeRespS2C) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{6} }

func (m *DuelAgreeRespS2C) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *DuelAgreeRespS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

// 玩家1取消邀请
type DuelInviteCancelC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInviteCancelC2S) Reset()                    { *m = DuelInviteCancelC2S{} }
func (m *DuelInviteCancelC2S) String() string            { return proto.CompactTextString(m) }
func (*DuelInviteCancelC2S) ProtoMessage()               {}
func (*DuelInviteCancelC2S) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{7} }

// @Gamer
type DuelInviteCancelS2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInviteCancelS2C) Reset()                    { *m = DuelInviteCancelS2C{} }
func (m *DuelInviteCancelS2C) String() string            { return proto.CompactTextString(m) }
func (*DuelInviteCancelS2C) ProtoMessage()               {}
func (*DuelInviteCancelS2C) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{8} }

type DuelInfo struct {
	// 对手id
	Gid *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	// 状态(0:无 1：邀请中, 2,比赛中)
	Status *int32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	// 发起时间
	SendTime *int64 `protobuf:"varint,3,opt,name=sendTime" json:"sendTime,omitempty"`
	// 已处理标记
	Flag             *bool  `protobuf:"varint,4,opt,name=flag" json:"flag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelInfo) Reset()                    { *m = DuelInfo{} }
func (m *DuelInfo) String() string            { return proto.CompactTextString(m) }
func (*DuelInfo) ProtoMessage()               {}
func (*DuelInfo) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{9} }

func (m *DuelInfo) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *DuelInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *DuelInfo) GetSendTime() int64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

func (m *DuelInfo) GetFlag() bool {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return false
}

type CreateDuelRoomInfoM2F struct {
	Self             *DuelUser `protobuf:"bytes,1,opt,name=self" json:"self,omitempty"`
	Other            *DuelUser `protobuf:"bytes,2,opt,name=other" json:"other,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *CreateDuelRoomInfoM2F) Reset()                    { *m = CreateDuelRoomInfoM2F{} }
func (m *CreateDuelRoomInfoM2F) String() string            { return proto.CompactTextString(m) }
func (*CreateDuelRoomInfoM2F) ProtoMessage()               {}
func (*CreateDuelRoomInfoM2F) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{10} }

func (m *CreateDuelRoomInfoM2F) GetSelf() *DuelUser {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *CreateDuelRoomInfoM2F) GetOther() *DuelUser {
	if m != nil {
		return m.Other
	}
	return nil
}

type DuelUser struct {
	// 玩家id
	Gid *int32 `protobuf:"varint,1,opt,name=gid" json:"gid,omitempty"`
	// psId
	PsId *int32 `protobuf:"varint,2,opt,name=psId" json:"psId,omitempty"`
	// 区服id
	LsId *int32 `protobuf:"varint,3,opt,name=lsId" json:"lsId,omitempty"`
	// 是否是机器人
	IsRobot          *bool  `protobuf:"varint,4,opt,name=isRobot" json:"isRobot,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DuelUser) Reset()                    { *m = DuelUser{} }
func (m *DuelUser) String() string            { return proto.CompactTextString(m) }
func (*DuelUser) ProtoMessage()               {}
func (*DuelUser) Descriptor() ([]byte, []int) { return fileDescriptor28, []int{11} }

func (m *DuelUser) GetGid() int32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *DuelUser) GetPsId() int32 {
	if m != nil && m.PsId != nil {
		return *m.PsId
	}
	return 0
}

func (m *DuelUser) GetLsId() int32 {
	if m != nil && m.LsId != nil {
		return *m.LsId
	}
	return 0
}

func (m *DuelUser) GetIsRobot() bool {
	if m != nil && m.IsRobot != nil {
		return *m.IsRobot
	}
	return false
}

func init() {
	proto.RegisterType((*DuelInviteSendC2S)(nil), "DuelInviteSendC2S")
	proto.RegisterType((*DuelInviteSendS2C)(nil), "DuelInviteSendS2C")
	proto.RegisterType((*DuelInviteRecvS2C)(nil), "DuelInviteRecvS2C")
	proto.RegisterType((*DuelAgreeC2S)(nil), "DuelAgreeC2S")
	proto.RegisterType((*DuelAgreeS2C)(nil), "DuelAgreeS2C")
	proto.RegisterType((*DuelInviteRespS2C)(nil), "DuelInviteRespS2C")
	proto.RegisterType((*DuelAgreeRespS2C)(nil), "DuelAgreeRespS2C")
	proto.RegisterType((*DuelInviteCancelC2S)(nil), "DuelInviteCancelC2S")
	proto.RegisterType((*DuelInviteCancelS2C)(nil), "DuelInviteCancelS2C")
	proto.RegisterType((*DuelInfo)(nil), "duelInfo")
	proto.RegisterType((*CreateDuelRoomInfoM2F)(nil), "CreateDuelRoomInfoM2F")
	proto.RegisterType((*DuelUser)(nil), "DuelUser")
	proto.RegisterEnum("DuelInviteSendC2S_MsgID", DuelInviteSendC2S_MsgID_name, DuelInviteSendC2S_MsgID_value)
	proto.RegisterEnum("DuelInviteSendS2C_MsgID", DuelInviteSendS2C_MsgID_name, DuelInviteSendS2C_MsgID_value)
	proto.RegisterEnum("DuelInviteRecvS2C_MsgID", DuelInviteRecvS2C_MsgID_name, DuelInviteRecvS2C_MsgID_value)
	proto.RegisterEnum("DuelAgreeC2S_MsgID", DuelAgreeC2S_MsgID_name, DuelAgreeC2S_MsgID_value)
	proto.RegisterEnum("DuelAgreeS2C_MsgID", DuelAgreeS2C_MsgID_name, DuelAgreeS2C_MsgID_value)
	proto.RegisterEnum("DuelInviteRespS2C_MsgID", DuelInviteRespS2C_MsgID_name, DuelInviteRespS2C_MsgID_value)
	proto.RegisterEnum("DuelAgreeRespS2C_MsgID", DuelAgreeRespS2C_MsgID_name, DuelAgreeRespS2C_MsgID_value)
	proto.RegisterEnum("DuelInviteCancelC2S_MsgID", DuelInviteCancelC2S_MsgID_name, DuelInviteCancelC2S_MsgID_value)
	proto.RegisterEnum("DuelInviteCancelS2C_MsgID", DuelInviteCancelS2C_MsgID_name, DuelInviteCancelS2C_MsgID_value)
}

func init() { proto.RegisterFile("friendDuel.proto", fileDescriptor28) }

var fileDescriptor28 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x5f, 0x4f, 0xc2, 0x30,
	0x14, 0xc5, 0x83, 0x63, 0x38, 0x2f, 0x55, 0xe7, 0xd4, 0xb8, 0xf8, 0x64, 0xfa, 0x64, 0x62, 0xc2,
	0xc3, 0x34, 0xbe, 0x9b, 0x21, 0x8a, 0x09, 0x2f, 0x20, 0x1f, 0x60, 0xb2, 0xbb, 0xb9, 0x64, 0xac,
	0xa4, 0x2d, 0x7c, 0x15, 0xff, 0x7e, 0x58, 0xdb, 0xce, 0x20, 0x93, 0x61, 0x7c, 0x5b, 0xd7, 0x73,
	0x7e, 0xe7, 0xdc, 0x5b, 0x70, 0x13, 0x9e, 0x61, 0x11, 0x77, 0xe7, 0x98, 0x77, 0x66, 0x9c, 0x49,
	0x76, 0xda, 0x9e, 0x46, 0x13, 0xce, 0xca, 0x03, 0xbd, 0x86, 0x03, 0x7d, 0xd5, 0x2f, 0x16, 0x99,
	0xc4, 0x91, 0x12, 0x86, 0xc1, 0xc8, 0x6b, 0x83, 0x95, 0x66, 0xb1, 0xdf, 0x38, 0x6b, 0x9c, 0xdb,
	0xf4, 0x08, 0xec, 0x81, 0x48, 0xfb, 0x5d, 0xf5, 0xb7, 0x85, 0xe6, 0xcb, 0x7d, 0x21, 0xb4, 0xf7,
	0xdb, 0x37, 0x0a, 0xc2, 0x8a, 0xcf, 0xdb, 0x83, 0x16, 0x47, 0x31, 0xcf, 0xa5, 0xbf, 0xf5, 0x07,
	0xe7, 0x7e, 0x95, 0x33, 0xc4, 0xc9, 0x62, 0x8d, 0xe3, 0x82, 0x23, 0x14, 0xff, 0x31, 0x9b, 0xa2,
	0x21, 0x59, 0x75, 0xa4, 0x57, 0x42, 0x2f, 0x81, 0x68, 0xd2, 0x4d, 0xca, 0x11, 0xf5, 0x10, 0xbb,
	0x60, 0x47, 0xfa, 0xdb, 0x98, 0x9c, 0x3a, 0xd3, 0x1b, 0xa1, 0x57, 0x2b, 0x26, 0x9d, 0xfc, 0x8f,
	0xd2, 0xca, 0x35, 0xae, 0x96, 0x16, 0xb3, 0xb5, 0xd2, 0xd5, 0xf0, 0x12, 0x1b, 0x09, 0x56, 0xf8,
	0xd6, 0x26, 0xec, 0x3b, 0xa1, 0xb7, 0xe0, 0x2e, 0xcb, 0xd4, 0x52, 0xab, 0xed, 0x6a, 0x67, 0xfa,
	0x20, 0xf4, 0x02, 0x0e, 0x7f, 0xda, 0x85, 0x51, 0x31, 0xc1, 0x5c, 0xed, 0xa3, 0x4e, 0xfc, 0x59,
	0x2b, 0x56, 0xb1, 0x1b, 0xc4, 0x77, 0xe0, 0xc4, 0x46, 0x9c, 0xb0, 0xb5, 0x62, 0x42, 0x46, 0x72,
	0x2e, 0xca, 0xb5, 0x55, 0xde, 0x4c, 0x4f, 0x6c, 0x79, 0x04, 0x9a, 0x49, 0x1e, 0xa5, 0x7e, 0xd3,
	0x14, 0x7f, 0x80, 0xe3, 0x50, 0x2d, 0x44, 0xa2, 0xce, 0x1e, 0x32, 0x36, 0xd5, 0xc8, 0x41, 0xd0,
	0xf3, 0x4e, 0xa0, 0x29, 0x30, 0x4f, 0x0c, 0xb6, 0x1d, 0xec, 0x74, 0xf4, 0xfd, 0x58, 0x20, 0xf7,
	0x7c, 0xb0, 0x99, 0x7c, 0x46, 0x6e, 0x02, 0x56, 0x6f, 0x68, 0x17, 0x9c, 0xa5, 0xaa, 0x52, 0x4a,
	0x45, 0xce, 0x44, 0x3f, 0xfe, 0xae, 0xa4, 0x4e, 0xb9, 0x3e, 0x99, 0x07, 0xf0, 0xf6, 0x61, 0x3b,
	0x13, 0x43, 0xf6, 0xc4, 0x64, 0xd9, 0xe8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x97, 0x04, 0x50,
	0x27, 0x03, 0x00, 0x00,
}
