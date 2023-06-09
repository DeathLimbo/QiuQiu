// Code generated by protoc-gen-go.
// source: mail.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SendAllMailS2C_MsgID int32

const (
	SendAllMailS2C_eMsgID SendAllMailS2C_MsgID = 1281
)

var SendAllMailS2C_MsgID_name = map[int32]string{
	1281: "eMsgID",
}
var SendAllMailS2C_MsgID_value = map[string]int32{
	"eMsgID": 1281,
}

func (x SendAllMailS2C_MsgID) Enum() *SendAllMailS2C_MsgID {
	p := new(SendAllMailS2C_MsgID)
	*p = x
	return p
}
func (x SendAllMailS2C_MsgID) String() string {
	return proto.EnumName(SendAllMailS2C_MsgID_name, int32(x))
}
func (x *SendAllMailS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SendAllMailS2C_MsgID_value, data, "SendAllMailS2C_MsgID")
	if err != nil {
		return err
	}
	*x = SendAllMailS2C_MsgID(value)
	return nil
}
func (SendAllMailS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor42, []int{3, 0} }

type GamerMailDeleteMailC2S_MsgID int32

const (
	GamerMailDeleteMailC2S_eMsgID GamerMailDeleteMailC2S_MsgID = 1282
)

var GamerMailDeleteMailC2S_MsgID_name = map[int32]string{
	1282: "eMsgID",
}
var GamerMailDeleteMailC2S_MsgID_value = map[string]int32{
	"eMsgID": 1282,
}

func (x GamerMailDeleteMailC2S_MsgID) Enum() *GamerMailDeleteMailC2S_MsgID {
	p := new(GamerMailDeleteMailC2S_MsgID)
	*p = x
	return p
}
func (x GamerMailDeleteMailC2S_MsgID) String() string {
	return proto.EnumName(GamerMailDeleteMailC2S_MsgID_name, int32(x))
}
func (x *GamerMailDeleteMailC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GamerMailDeleteMailC2S_MsgID_value, data, "GamerMailDeleteMailC2S_MsgID")
	if err != nil {
		return err
	}
	*x = GamerMailDeleteMailC2S_MsgID(value)
	return nil
}
func (GamerMailDeleteMailC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor42, []int{4, 0}
}

type GamerMailDeleteMailS2C_MsgID int32

const (
	GamerMailDeleteMailS2C_eMsgID GamerMailDeleteMailS2C_MsgID = 1282
)

var GamerMailDeleteMailS2C_MsgID_name = map[int32]string{
	1282: "eMsgID",
}
var GamerMailDeleteMailS2C_MsgID_value = map[string]int32{
	"eMsgID": 1282,
}

func (x GamerMailDeleteMailS2C_MsgID) Enum() *GamerMailDeleteMailS2C_MsgID {
	p := new(GamerMailDeleteMailS2C_MsgID)
	*p = x
	return p
}
func (x GamerMailDeleteMailS2C_MsgID) String() string {
	return proto.EnumName(GamerMailDeleteMailS2C_MsgID_name, int32(x))
}
func (x *GamerMailDeleteMailS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GamerMailDeleteMailS2C_MsgID_value, data, "GamerMailDeleteMailS2C_MsgID")
	if err != nil {
		return err
	}
	*x = GamerMailDeleteMailS2C_MsgID(value)
	return nil
}
func (GamerMailDeleteMailS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor42, []int{5, 0}
}

type GamerMailDrawRewardC2S_MsgID int32

const (
	GamerMailDrawRewardC2S_eMsgID GamerMailDrawRewardC2S_MsgID = 1283
)

var GamerMailDrawRewardC2S_MsgID_name = map[int32]string{
	1283: "eMsgID",
}
var GamerMailDrawRewardC2S_MsgID_value = map[string]int32{
	"eMsgID": 1283,
}

func (x GamerMailDrawRewardC2S_MsgID) Enum() *GamerMailDrawRewardC2S_MsgID {
	p := new(GamerMailDrawRewardC2S_MsgID)
	*p = x
	return p
}
func (x GamerMailDrawRewardC2S_MsgID) String() string {
	return proto.EnumName(GamerMailDrawRewardC2S_MsgID_name, int32(x))
}
func (x *GamerMailDrawRewardC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GamerMailDrawRewardC2S_MsgID_value, data, "GamerMailDrawRewardC2S_MsgID")
	if err != nil {
		return err
	}
	*x = GamerMailDrawRewardC2S_MsgID(value)
	return nil
}
func (GamerMailDrawRewardC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor42, []int{6, 0}
}

type GamerMailDrawRewardS2C_MsgID int32

const (
	GamerMailDrawRewardS2C_eMsgID GamerMailDrawRewardS2C_MsgID = 1283
)

var GamerMailDrawRewardS2C_MsgID_name = map[int32]string{
	1283: "eMsgID",
}
var GamerMailDrawRewardS2C_MsgID_value = map[string]int32{
	"eMsgID": 1283,
}

func (x GamerMailDrawRewardS2C_MsgID) Enum() *GamerMailDrawRewardS2C_MsgID {
	p := new(GamerMailDrawRewardS2C_MsgID)
	*p = x
	return p
}
func (x GamerMailDrawRewardS2C_MsgID) String() string {
	return proto.EnumName(GamerMailDrawRewardS2C_MsgID_name, int32(x))
}
func (x *GamerMailDrawRewardS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GamerMailDrawRewardS2C_MsgID_value, data, "GamerMailDrawRewardS2C_MsgID")
	if err != nil {
		return err
	}
	*x = GamerMailDrawRewardS2C_MsgID(value)
	return nil
}
func (GamerMailDrawRewardS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor42, []int{7, 0}
}

type GamerMailSetReadC2S_MsgID int32

const (
	GamerMailSetReadC2S_eMsgID GamerMailSetReadC2S_MsgID = 1284
)

var GamerMailSetReadC2S_MsgID_name = map[int32]string{
	1284: "eMsgID",
}
var GamerMailSetReadC2S_MsgID_value = map[string]int32{
	"eMsgID": 1284,
}

func (x GamerMailSetReadC2S_MsgID) Enum() *GamerMailSetReadC2S_MsgID {
	p := new(GamerMailSetReadC2S_MsgID)
	*p = x
	return p
}
func (x GamerMailSetReadC2S_MsgID) String() string {
	return proto.EnumName(GamerMailSetReadC2S_MsgID_name, int32(x))
}
func (x *GamerMailSetReadC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GamerMailSetReadC2S_MsgID_value, data, "GamerMailSetReadC2S_MsgID")
	if err != nil {
		return err
	}
	*x = GamerMailSetReadC2S_MsgID(value)
	return nil
}
func (GamerMailSetReadC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor42, []int{8, 0}
}

type GamerMailSetReadS2C_MsgID int32

const (
	GamerMailSetReadS2C_eMsgID GamerMailSetReadS2C_MsgID = 1284
)

var GamerMailSetReadS2C_MsgID_name = map[int32]string{
	1284: "eMsgID",
}
var GamerMailSetReadS2C_MsgID_value = map[string]int32{
	"eMsgID": 1284,
}

func (x GamerMailSetReadS2C_MsgID) Enum() *GamerMailSetReadS2C_MsgID {
	p := new(GamerMailSetReadS2C_MsgID)
	*p = x
	return p
}
func (x GamerMailSetReadS2C_MsgID) String() string {
	return proto.EnumName(GamerMailSetReadS2C_MsgID_name, int32(x))
}
func (x *GamerMailSetReadS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GamerMailSetReadS2C_MsgID_value, data, "GamerMailSetReadS2C_MsgID")
	if err != nil {
		return err
	}
	*x = GamerMailSetReadS2C_MsgID(value)
	return nil
}
func (GamerMailSetReadS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor42, []int{9, 0}
}

type UpdateMailS2C_MsgID int32

const (
	UpdateMailS2C_eMsgID UpdateMailS2C_MsgID = 1286
)

var UpdateMailS2C_MsgID_name = map[int32]string{
	1286: "eMsgID",
}
var UpdateMailS2C_MsgID_value = map[string]int32{
	"eMsgID": 1286,
}

func (x UpdateMailS2C_MsgID) Enum() *UpdateMailS2C_MsgID {
	p := new(UpdateMailS2C_MsgID)
	*p = x
	return p
}
func (x UpdateMailS2C_MsgID) String() string {
	return proto.EnumName(UpdateMailS2C_MsgID_name, int32(x))
}
func (x *UpdateMailS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UpdateMailS2C_MsgID_value, data, "UpdateMailS2C_MsgID")
	if err != nil {
		return err
	}
	*x = UpdateMailS2C_MsgID(value)
	return nil
}
func (UpdateMailS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor42, []int{10, 0} }

// 邮件基本信息
type MailItem struct {
	// 邮件唯一ID(uuid)
	MailID *int64 `protobuf:"varint,1,opt,name=mailID" json:"mailID,omitempty"`
	// 标题
	Title *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// 发件人
	SendPlayer *string `protobuf:"bytes,3,opt,name=sendPlayer" json:"sendPlayer,omitempty"`
	// 内容
	Content *string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	// 资源
	Reward []*ItemVaryConfig `protobuf:"bytes,5,rep,name=reward" json:"reward,omitempty"`
	// 结束时间
	EndTime *int64 `protobuf:"varint,6,opt,name=endTime" json:"endTime,omitempty"`
	// 发送时间
	SendTime *int64 `protobuf:"varint,7,opt,name=sendTime" json:"sendTime,omitempty"`
	// 是否已领取奖励
	IsDrawReward *bool `protobuf:"varint,8,opt,name=isDrawReward" json:"isDrawReward,omitempty"`
	// 是否已读
	IsRead *bool `protobuf:"varint,9,opt,name=isRead" json:"isRead,omitempty"`
	// 发资源原因（有资源才用）
	ReasonId *int32 `protobuf:"varint,10,opt,name=reasonId" json:"reasonId,omitempty"`
	// 导致发资源的id
	ReasonSubId      *int32 `protobuf:"varint,11,opt,name=reasonSubId" json:"reasonSubId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MailItem) Reset()                    { *m = MailItem{} }
func (m *MailItem) String() string            { return proto.CompactTextString(m) }
func (*MailItem) ProtoMessage()               {}
func (*MailItem) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{0} }

func (m *MailItem) GetMailID() int64 {
	if m != nil && m.MailID != nil {
		return *m.MailID
	}
	return 0
}

func (m *MailItem) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *MailItem) GetSendPlayer() string {
	if m != nil && m.SendPlayer != nil {
		return *m.SendPlayer
	}
	return ""
}

func (m *MailItem) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *MailItem) GetReward() []*ItemVaryConfig {
	if m != nil {
		return m.Reward
	}
	return nil
}

func (m *MailItem) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *MailItem) GetSendTime() int64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

func (m *MailItem) GetIsDrawReward() bool {
	if m != nil && m.IsDrawReward != nil {
		return *m.IsDrawReward
	}
	return false
}

func (m *MailItem) GetIsRead() bool {
	if m != nil && m.IsRead != nil {
		return *m.IsRead
	}
	return false
}

func (m *MailItem) GetReasonId() int32 {
	if m != nil && m.ReasonId != nil {
		return *m.ReasonId
	}
	return 0
}

func (m *MailItem) GetReasonSubId() int32 {
	if m != nil && m.ReasonSubId != nil {
		return *m.ReasonSubId
	}
	return 0
}

// 收到邮件信息
type ReceiveMailC2L struct {
	// 邮件信息
	MailId           *int64 `protobuf:"varint,1,opt,name=mailId" json:"mailId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReceiveMailC2L) Reset()                    { *m = ReceiveMailC2L{} }
func (m *ReceiveMailC2L) String() string            { return proto.CompactTextString(m) }
func (*ReceiveMailC2L) ProtoMessage()               {}
func (*ReceiveMailC2L) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{1} }

func (m *ReceiveMailC2L) GetMailId() int64 {
	if m != nil && m.MailId != nil {
		return *m.MailId
	}
	return 0
}

// 邮件对象
type ReceiveMailObj struct {
	// 邮件信息
	Mail *MailItem `protobuf:"bytes,1,opt,name=mail" json:"mail,omitempty"`
	// 收件人
	GamerID *int32 `protobuf:"varint,2,opt,name=gamerID" json:"gamerID,omitempty"`
	// pvp参数
	PvpParam         *int32 `protobuf:"varint,3,opt,name=pvpParam" json:"pvpParam,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReceiveMailObj) Reset()                    { *m = ReceiveMailObj{} }
func (m *ReceiveMailObj) String() string            { return proto.CompactTextString(m) }
func (*ReceiveMailObj) ProtoMessage()               {}
func (*ReceiveMailObj) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{2} }

func (m *ReceiveMailObj) GetMail() *MailItem {
	if m != nil {
		return m.Mail
	}
	return nil
}

func (m *ReceiveMailObj) GetGamerID() int32 {
	if m != nil && m.GamerID != nil {
		return *m.GamerID
	}
	return 0
}

func (m *ReceiveMailObj) GetPvpParam() int32 {
	if m != nil && m.PvpParam != nil {
		return *m.PvpParam
	}
	return 0
}

// 上线推送所有邮件
type SendAllMailS2C struct {
	// 邮件信息
	Mail             []*MailItem `protobuf:"bytes,1,rep,name=mail" json:"mail,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SendAllMailS2C) Reset()                    { *m = SendAllMailS2C{} }
func (m *SendAllMailS2C) String() string            { return proto.CompactTextString(m) }
func (*SendAllMailS2C) ProtoMessage()               {}
func (*SendAllMailS2C) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{3} }

func (m *SendAllMailS2C) GetMail() []*MailItem {
	if m != nil {
		return m.Mail
	}
	return nil
}

// 删除邮件
type GamerMailDeleteMailC2S struct {
	// 邮件ID(一键删除id传0)
	MailID           *int64 `protobuf:"varint,1,opt,name=mailID" json:"mailID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GamerMailDeleteMailC2S) Reset()                    { *m = GamerMailDeleteMailC2S{} }
func (m *GamerMailDeleteMailC2S) String() string            { return proto.CompactTextString(m) }
func (*GamerMailDeleteMailC2S) ProtoMessage()               {}
func (*GamerMailDeleteMailC2S) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{4} }

func (m *GamerMailDeleteMailC2S) GetMailID() int64 {
	if m != nil && m.MailID != nil {
		return *m.MailID
	}
	return 0
}

// 删除邮件
type GamerMailDeleteMailS2C struct {
	// 被删除的邮件id列表
	DeletedMailIds   []int64 `protobuf:"varint,1,rep,name=deletedMailIds" json:"deletedMailIds,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GamerMailDeleteMailS2C) Reset()                    { *m = GamerMailDeleteMailS2C{} }
func (m *GamerMailDeleteMailS2C) String() string            { return proto.CompactTextString(m) }
func (*GamerMailDeleteMailS2C) ProtoMessage()               {}
func (*GamerMailDeleteMailS2C) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{5} }

func (m *GamerMailDeleteMailS2C) GetDeletedMailIds() []int64 {
	if m != nil {
		return m.DeletedMailIds
	}
	return nil
}

// 领取附件
type GamerMailDrawRewardC2S struct {
	// 邮件ID(一键领取id传0)
	MailID           *int64 `protobuf:"varint,1,opt,name=mailID" json:"mailID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GamerMailDrawRewardC2S) Reset()                    { *m = GamerMailDrawRewardC2S{} }
func (m *GamerMailDrawRewardC2S) String() string            { return proto.CompactTextString(m) }
func (*GamerMailDrawRewardC2S) ProtoMessage()               {}
func (*GamerMailDrawRewardC2S) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{6} }

func (m *GamerMailDrawRewardC2S) GetMailID() int64 {
	if m != nil && m.MailID != nil {
		return *m.MailID
	}
	return 0
}

// 领取附件
type GamerMailDrawRewardS2C struct {
	// 邮件信息
	MailIdList       []int64 `protobuf:"varint,1,rep,name=mailIdList" json:"mailIdList,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GamerMailDrawRewardS2C) Reset()                    { *m = GamerMailDrawRewardS2C{} }
func (m *GamerMailDrawRewardS2C) String() string            { return proto.CompactTextString(m) }
func (*GamerMailDrawRewardS2C) ProtoMessage()               {}
func (*GamerMailDrawRewardS2C) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{7} }

func (m *GamerMailDrawRewardS2C) GetMailIdList() []int64 {
	if m != nil {
		return m.MailIdList
	}
	return nil
}

// 设置已读
type GamerMailSetReadC2S struct {
	// 邮件ID
	MailID           *int64 `protobuf:"varint,1,opt,name=mailID" json:"mailID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GamerMailSetReadC2S) Reset()                    { *m = GamerMailSetReadC2S{} }
func (m *GamerMailSetReadC2S) String() string            { return proto.CompactTextString(m) }
func (*GamerMailSetReadC2S) ProtoMessage()               {}
func (*GamerMailSetReadC2S) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{8} }

func (m *GamerMailSetReadC2S) GetMailID() int64 {
	if m != nil && m.MailID != nil {
		return *m.MailID
	}
	return 0
}

// 设置已读
type GamerMailSetReadS2C struct {
	// 邮件信息
	Mail             *MailItem `protobuf:"bytes,1,opt,name=mail" json:"mail,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *GamerMailSetReadS2C) Reset()                    { *m = GamerMailSetReadS2C{} }
func (m *GamerMailSetReadS2C) String() string            { return proto.CompactTextString(m) }
func (*GamerMailSetReadS2C) ProtoMessage()               {}
func (*GamerMailSetReadS2C) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{9} }

func (m *GamerMailSetReadS2C) GetMail() *MailItem {
	if m != nil {
		return m.Mail
	}
	return nil
}

// 更新邮件
type UpdateMailS2C struct {
	// 邮件信息
	Mail             []*MailItem `protobuf:"bytes,1,rep,name=mail" json:"mail,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *UpdateMailS2C) Reset()                    { *m = UpdateMailS2C{} }
func (m *UpdateMailS2C) String() string            { return proto.CompactTextString(m) }
func (*UpdateMailS2C) ProtoMessage()               {}
func (*UpdateMailS2C) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{10} }

func (m *UpdateMailS2C) GetMail() []*MailItem {
	if m != nil {
		return m.Mail
	}
	return nil
}

// 全服邮件结构
type GlobalMailItem struct {
	// 邮件唯一ID(uuid)
	MailUuid *int64 `protobuf:"varint,1,opt,name=mailUuid" json:"mailUuid,omitempty"`
	// 发件人
	SendPlayer *string `protobuf:"bytes,2,opt,name=sendPlayer" json:"sendPlayer,omitempty"`
	// 标题
	Title *string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	// 内容
	Content *string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	// 资源附件
	Rewards []*ItemVaryConfig `protobuf:"bytes,5,rep,name=rewards" json:"rewards,omitempty"`
	// 生成时间
	CreateTime *int64 `protobuf:"varint,6,opt,name=createTime" json:"createTime,omitempty"`
	// 生效时间(如果定时发送的邮件, 生效时间从定时发送时开始生效)
	StartTime *int64 `protobuf:"varint,7,opt,name=startTime" json:"startTime,omitempty"`
	// 结束时间（全局邮件有30天的存活期）
	EndTime *int64 `protobuf:"varint,8,opt,name=endTime" json:"endTime,omitempty"`
	// 定时发送时间(0:直接发送)
	TimingSendTime *int64 `protobuf:"varint,9,opt,name=timingSendTime" json:"timingSendTime,omitempty"`
	// 区服Id列表(全服邮件适用在那些区服)（0：全区服适用）
	ZoneIds []int32 `protobuf:"varint,10,rep,name=zoneIds" json:"zoneIds,omitempty"`
	// 渠道（暂时无用，预留 有渠道后使用）
	Channel *string `protobuf:"bytes,11,opt,name=channel" json:"channel,omitempty"`
	// 收件人条件类型（1：玩家等级 2：注册时间区间 3：最近上线时间区间）
	CondType *int32 `protobuf:"varint,12,opt,name=condType" json:"condType,omitempty"`
	// 收件人条件参数区间左值(开始时间戳、开始等级)
	CondValLeft *int64 `protobuf:"varint,13,opt,name=condValLeft" json:"condValLeft,omitempty"`
	// 收件人条件参数区间右值（结束时间戳，结束等级）
	CondValRight *int64 `protobuf:"varint,14,opt,name=condValRight" json:"condValRight,omitempty"`
	// 邮件描述
	Describe         *string `protobuf:"bytes,15,opt,name=describe" json:"describe,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GlobalMailItem) Reset()                    { *m = GlobalMailItem{} }
func (m *GlobalMailItem) String() string            { return proto.CompactTextString(m) }
func (*GlobalMailItem) ProtoMessage()               {}
func (*GlobalMailItem) Descriptor() ([]byte, []int) { return fileDescriptor42, []int{11} }

func (m *GlobalMailItem) GetMailUuid() int64 {
	if m != nil && m.MailUuid != nil {
		return *m.MailUuid
	}
	return 0
}

func (m *GlobalMailItem) GetSendPlayer() string {
	if m != nil && m.SendPlayer != nil {
		return *m.SendPlayer
	}
	return ""
}

func (m *GlobalMailItem) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *GlobalMailItem) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *GlobalMailItem) GetRewards() []*ItemVaryConfig {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *GlobalMailItem) GetCreateTime() int64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *GlobalMailItem) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *GlobalMailItem) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *GlobalMailItem) GetTimingSendTime() int64 {
	if m != nil && m.TimingSendTime != nil {
		return *m.TimingSendTime
	}
	return 0
}

func (m *GlobalMailItem) GetZoneIds() []int32 {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *GlobalMailItem) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *GlobalMailItem) GetCondType() int32 {
	if m != nil && m.CondType != nil {
		return *m.CondType
	}
	return 0
}

func (m *GlobalMailItem) GetCondValLeft() int64 {
	if m != nil && m.CondValLeft != nil {
		return *m.CondValLeft
	}
	return 0
}

func (m *GlobalMailItem) GetCondValRight() int64 {
	if m != nil && m.CondValRight != nil {
		return *m.CondValRight
	}
	return 0
}

func (m *GlobalMailItem) GetDescribe() string {
	if m != nil && m.Describe != nil {
		return *m.Describe
	}
	return ""
}

func init() {
	proto.RegisterType((*MailItem)(nil), "MailItem")
	proto.RegisterType((*ReceiveMailC2L)(nil), "ReceiveMailC2L")
	proto.RegisterType((*ReceiveMailObj)(nil), "ReceiveMailObj")
	proto.RegisterType((*SendAllMailS2C)(nil), "SendAllMailS2C")
	proto.RegisterType((*GamerMailDeleteMailC2S)(nil), "GamerMailDeleteMailC2S")
	proto.RegisterType((*GamerMailDeleteMailS2C)(nil), "GamerMailDeleteMailS2C")
	proto.RegisterType((*GamerMailDrawRewardC2S)(nil), "GamerMailDrawRewardC2S")
	proto.RegisterType((*GamerMailDrawRewardS2C)(nil), "GamerMailDrawRewardS2C")
	proto.RegisterType((*GamerMailSetReadC2S)(nil), "GamerMailSetReadC2S")
	proto.RegisterType((*GamerMailSetReadS2C)(nil), "GamerMailSetReadS2C")
	proto.RegisterType((*UpdateMailS2C)(nil), "UpdateMailS2C")
	proto.RegisterType((*GlobalMailItem)(nil), "GlobalMailItem")
	proto.RegisterEnum("SendAllMailS2C_MsgID", SendAllMailS2C_MsgID_name, SendAllMailS2C_MsgID_value)
	proto.RegisterEnum("GamerMailDeleteMailC2S_MsgID", GamerMailDeleteMailC2S_MsgID_name, GamerMailDeleteMailC2S_MsgID_value)
	proto.RegisterEnum("GamerMailDeleteMailS2C_MsgID", GamerMailDeleteMailS2C_MsgID_name, GamerMailDeleteMailS2C_MsgID_value)
	proto.RegisterEnum("GamerMailDrawRewardC2S_MsgID", GamerMailDrawRewardC2S_MsgID_name, GamerMailDrawRewardC2S_MsgID_value)
	proto.RegisterEnum("GamerMailDrawRewardS2C_MsgID", GamerMailDrawRewardS2C_MsgID_name, GamerMailDrawRewardS2C_MsgID_value)
	proto.RegisterEnum("GamerMailSetReadC2S_MsgID", GamerMailSetReadC2S_MsgID_name, GamerMailSetReadC2S_MsgID_value)
	proto.RegisterEnum("GamerMailSetReadS2C_MsgID", GamerMailSetReadS2C_MsgID_name, GamerMailSetReadS2C_MsgID_value)
	proto.RegisterEnum("UpdateMailS2C_MsgID", UpdateMailS2C_MsgID_name, UpdateMailS2C_MsgID_value)
}

func init() { proto.RegisterFile("mail.proto", fileDescriptor42) }

var fileDescriptor42 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x15, 0x21, 0x06, 0x3c, 0x04, 0x93, 0x3a, 0x51, 0xba, 0xea, 0xa5, 0xc8, 0xa7, 0x9c, 0x38,
	0x70, 0xad, 0xd4, 0xaa, 0x05, 0x35, 0x42, 0x22, 0x6a, 0x84, 0x93, 0xdc, 0x17, 0x76, 0x42, 0xb6,
	0x5a, 0x7f, 0x68, 0x77, 0x93, 0x88, 0xde, 0xda, 0x46, 0xfd, 0x4f, 0xfd, 0x35, 0xfd, 0x2b, 0xdd,
	0x5d, 0x63, 0x17, 0x52, 0x50, 0xd3, 0x9b, 0xfd, 0x3c, 0xf3, 0xe6, 0xcd, 0xc7, 0x33, 0x40, 0x42,
	0xb9, 0xe8, 0xe7, 0x32, 0xd3, 0xd9, 0x2b, 0x98, 0x67, 0xe9, 0x4d, 0xf1, 0x1c, 0xfd, 0xaa, 0x41,
	0xeb, 0xdc, 0x7c, 0x1a, 0x6b, 0x4c, 0xc2, 0x00, 0x1a, 0x36, 0x6c, 0x3c, 0x22, 0xb5, 0x5e, 0xed,
	0xb4, 0x1e, 0x76, 0xc0, 0xd3, 0x5c, 0x0b, 0x24, 0x7b, 0xe6, 0xd5, 0x0f, 0x43, 0x00, 0x85, 0x29,
	0xbb, 0x10, 0x74, 0x89, 0x92, 0xd4, 0x1d, 0xd6, 0x85, 0xa6, 0x61, 0xd3, 0x98, 0x6a, 0xb2, 0xef,
	0x80, 0xd7, 0xd0, 0x90, 0xf8, 0x40, 0x25, 0x23, 0x5e, 0xaf, 0x7e, 0xda, 0x1e, 0x74, 0xfb, 0x96,
	0xfa, 0x9a, 0xca, 0xe5, 0xd0, 0x54, 0xe5, 0x0b, 0x9b, 0x61, 0x48, 0x2e, 0x79, 0x82, 0xa4, 0xe1,
	0xaa, 0x1c, 0x42, 0x4b, 0x95, 0x48, 0xd3, 0x21, 0xc7, 0x70, 0xc0, 0xd5, 0x48, 0xd2, 0x87, 0x69,
	0xc1, 0xd4, 0x32, 0x68, 0xcb, 0xaa, 0xe3, 0x6a, 0x8a, 0x94, 0x11, 0xdf, 0xbd, 0x9b, 0x3c, 0x89,
	0x54, 0x65, 0xe9, 0x98, 0x11, 0x30, 0x88, 0x17, 0x1e, 0x41, 0xbb, 0x40, 0xe2, 0xbb, 0x99, 0x01,
	0xdb, 0x16, 0x8c, 0x7a, 0x10, 0x4c, 0x71, 0x8e, 0xfc, 0x1e, 0x6d, 0x9f, 0xc3, 0xc1, 0xa4, 0x6a,
	0x93, 0x15, 0x6d, 0x46, 0x93, 0x8d, 0x88, 0x4f, 0xb3, 0xcf, 0xe1, 0x4b, 0xd8, 0xb7, 0x11, 0xee,
	0x7b, 0x7b, 0xe0, 0xf7, 0xab, 0x09, 0x19, 0xf1, 0x0b, 0x9a, 0xa0, 0x34, 0x23, 0xda, 0x73, 0x25,
	0x8d, 0x88, 0xfc, 0x3e, 0xbf, 0xa0, 0x92, 0x26, 0x6e, 0x22, 0x5e, 0xf4, 0x0e, 0x82, 0xd8, 0xb4,
	0xf3, 0x5e, 0x08, 0x9b, 0x15, 0x0f, 0x86, 0x6b, 0x6c, 0xf5, 0x0d, 0xb6, 0xe8, 0x18, 0xbc, 0x73,
	0xb5, 0x18, 0x8f, 0xc2, 0x36, 0x34, 0xd0, 0x3d, 0x1d, 0x7e, 0x85, 0xe8, 0x2d, 0x9c, 0x9c, 0xd9,
	0x1a, 0x36, 0x6c, 0x84, 0x02, 0xf5, 0x4a, 0x78, 0xfc, 0x74, 0x3f, 0xdb, 0xf2, 0xbf, 0x41, 0xf4,
	0x71, 0x6b, 0xbe, 0x15, 0x72, 0x02, 0x01, 0x73, 0x00, 0x73, 0x12, 0x98, 0x72, 0x92, 0x76, 0xf1,
	0x6c, 0xe8, 0xa8, 0x96, 0xf1, 0x4c, 0x1d, 0xdf, 0x21, 0xfa, 0xb0, 0x35, 0xdf, 0xea, 0x08, 0x8b,
	0x73, 0x1c, 0xb3, 0x09, 0x57, 0x7a, 0xb7, 0x06, 0xc3, 0xf1, 0x06, 0x8e, 0x2a, 0x8e, 0x18, 0xb5,
	0xdd, 0xfe, 0x33, 0x05, 0x3c, 0x42, 0x34, 0xfa, 0x3b, 0x79, 0x73, 0x1d, 0xb5, 0x7f, 0xae, 0xe3,
	0xd1, 0x8e, 0xa1, 0x73, 0x95, 0x33, 0xfa, 0x67, 0x8a, 0xff, 0xb3, 0xce, 0x1f, 0x10, 0xfd, 0xdc,
	0x83, 0xe0, 0x4c, 0x64, 0x33, 0x2a, 0xaa, 0x2b, 0x32, 0x47, 0x63, 0x19, 0xae, 0xee, 0xf8, 0xea,
	0x04, 0x9f, 0x58, 0xab, 0xb0, 0x5b, 0xe5, 0xbe, 0x1d, 0x4e, 0xeb, 0x41, 0xb3, 0x70, 0x9a, 0xda,
	0x65, 0x35, 0xc3, 0x3a, 0x37, 0x86, 0xd0, 0xb8, 0xe6, 0xb6, 0x17, 0xe0, 0x2b, 0x4d, 0xa5, 0x5e,
	0xb3, 0xdb, 0x9a, 0x23, 0x5b, 0x0e, 0x30, 0x77, 0xa2, 0x79, 0xc2, 0xd3, 0x45, 0x5c, 0xe2, 0x7e,
	0x19, 0xf8, 0x25, 0x4b, 0xd1, 0x1e, 0x0e, 0x98, 0x8a, 0x9e, 0xd3, 0x74, 0x4b, 0xd3, 0x14, 0x85,
	0x33, 0x9b, 0x6f, 0x3b, 0x33, 0x22, 0xd9, 0xe5, 0x32, 0x47, 0x72, 0x50, 0x7a, 0xd2, 0x22, 0xd7,
	0x54, 0x4c, 0xf0, 0x46, 0x93, 0x4e, 0x69, 0xf0, 0x15, 0x38, 0xe5, 0x8b, 0x5b, 0x4d, 0x82, 0xf2,
	0x47, 0xc0, 0x50, 0xcd, 0x25, 0x9f, 0x21, 0xe9, 0x5a, 0xba, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0x6a, 0x40, 0x60, 0xb6, 0x04, 0x00, 0x00,
}