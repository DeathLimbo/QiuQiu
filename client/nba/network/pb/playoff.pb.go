// Code generated by protoc-gen-go.
// source: playoff.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UpdatePlayoffMatchInfoS2C_MsgID int32

const (
	UpdatePlayoffMatchInfoS2C_eMsgID UpdatePlayoffMatchInfoS2C_MsgID = 12558
)

var UpdatePlayoffMatchInfoS2C_MsgID_name = map[int32]string{
	12558: "eMsgID",
}
var UpdatePlayoffMatchInfoS2C_MsgID_value = map[string]int32{
	"eMsgID": 12558,
}

func (x UpdatePlayoffMatchInfoS2C_MsgID) Enum() *UpdatePlayoffMatchInfoS2C_MsgID {
	p := new(UpdatePlayoffMatchInfoS2C_MsgID)
	*p = x
	return p
}
func (x UpdatePlayoffMatchInfoS2C_MsgID) String() string {
	return proto.EnumName(UpdatePlayoffMatchInfoS2C_MsgID_name, int32(x))
}
func (x *UpdatePlayoffMatchInfoS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UpdatePlayoffMatchInfoS2C_MsgID_value, data, "UpdatePlayoffMatchInfoS2C_MsgID")
	if err != nil {
		return err
	}
	*x = UpdatePlayoffMatchInfoS2C_MsgID(value)
	return nil
}
func (UpdatePlayoffMatchInfoS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor50, []int{3, 0}
}

type LoginSendPlayoffS2C_MsgID int32

const (
	LoginSendPlayoffS2C_eMsgID LoginSendPlayoffS2C_MsgID = 12559
)

var LoginSendPlayoffS2C_MsgID_name = map[int32]string{
	12559: "eMsgID",
}
var LoginSendPlayoffS2C_MsgID_value = map[string]int32{
	"eMsgID": 12559,
}

func (x LoginSendPlayoffS2C_MsgID) Enum() *LoginSendPlayoffS2C_MsgID {
	p := new(LoginSendPlayoffS2C_MsgID)
	*p = x
	return p
}
func (x LoginSendPlayoffS2C_MsgID) String() string {
	return proto.EnumName(LoginSendPlayoffS2C_MsgID_name, int32(x))
}
func (x *LoginSendPlayoffS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LoginSendPlayoffS2C_MsgID_value, data, "LoginSendPlayoffS2C_MsgID")
	if err != nil {
		return err
	}
	*x = LoginSendPlayoffS2C_MsgID(value)
	return nil
}
func (LoginSendPlayoffS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor50, []int{4, 0}
}

type PlayoffSweepC2S_MsgID int32

const (
	PlayoffSweepC2S_eMsgID PlayoffSweepC2S_MsgID = 12560
)

var PlayoffSweepC2S_MsgID_name = map[int32]string{
	12560: "eMsgID",
}
var PlayoffSweepC2S_MsgID_value = map[string]int32{
	"eMsgID": 12560,
}

func (x PlayoffSweepC2S_MsgID) Enum() *PlayoffSweepC2S_MsgID {
	p := new(PlayoffSweepC2S_MsgID)
	*p = x
	return p
}
func (x PlayoffSweepC2S_MsgID) String() string {
	return proto.EnumName(PlayoffSweepC2S_MsgID_name, int32(x))
}
func (x *PlayoffSweepC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayoffSweepC2S_MsgID_value, data, "PlayoffSweepC2S_MsgID")
	if err != nil {
		return err
	}
	*x = PlayoffSweepC2S_MsgID(value)
	return nil
}
func (PlayoffSweepC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor50, []int{5, 0} }

type PlayoffSweepS2C_MsgID int32

const (
	PlayoffSweepS2C_eMsgID PlayoffSweepS2C_MsgID = 12560
)

var PlayoffSweepS2C_MsgID_name = map[int32]string{
	12560: "eMsgID",
}
var PlayoffSweepS2C_MsgID_value = map[string]int32{
	"eMsgID": 12560,
}

func (x PlayoffSweepS2C_MsgID) Enum() *PlayoffSweepS2C_MsgID {
	p := new(PlayoffSweepS2C_MsgID)
	*p = x
	return p
}
func (x PlayoffSweepS2C_MsgID) String() string {
	return proto.EnumName(PlayoffSweepS2C_MsgID_name, int32(x))
}
func (x *PlayoffSweepS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PlayoffSweepS2C_MsgID_value, data, "PlayoffSweepS2C_MsgID")
	if err != nil {
		return err
	}
	*x = PlayoffSweepS2C_MsgID(value)
	return nil
}
func (PlayoffSweepS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor50, []int{6, 0} }

type RefreshPlayoffOnlineS2C_MsgID int32

const (
	RefreshPlayoffOnlineS2C_eMsgID RefreshPlayoffOnlineS2C_MsgID = 12562
)

var RefreshPlayoffOnlineS2C_MsgID_name = map[int32]string{
	12562: "eMsgID",
}
var RefreshPlayoffOnlineS2C_MsgID_value = map[string]int32{
	"eMsgID": 12562,
}

func (x RefreshPlayoffOnlineS2C_MsgID) Enum() *RefreshPlayoffOnlineS2C_MsgID {
	p := new(RefreshPlayoffOnlineS2C_MsgID)
	*p = x
	return p
}
func (x RefreshPlayoffOnlineS2C_MsgID) String() string {
	return proto.EnumName(RefreshPlayoffOnlineS2C_MsgID_name, int32(x))
}
func (x *RefreshPlayoffOnlineS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RefreshPlayoffOnlineS2C_MsgID_value, data, "RefreshPlayoffOnlineS2C_MsgID")
	if err != nil {
		return err
	}
	*x = RefreshPlayoffOnlineS2C_MsgID(value)
	return nil
}
func (RefreshPlayoffOnlineS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor50, []int{7, 0}
}

type AddNewPlayoffS2C_MsgID int32

const (
	AddNewPlayoffS2C_eMsgID AddNewPlayoffS2C_MsgID = 12561
)

var AddNewPlayoffS2C_MsgID_name = map[int32]string{
	12561: "eMsgID",
}
var AddNewPlayoffS2C_MsgID_value = map[string]int32{
	"eMsgID": 12561,
}

func (x AddNewPlayoffS2C_MsgID) Enum() *AddNewPlayoffS2C_MsgID {
	p := new(AddNewPlayoffS2C_MsgID)
	*p = x
	return p
}
func (x AddNewPlayoffS2C_MsgID) String() string {
	return proto.EnumName(AddNewPlayoffS2C_MsgID_name, int32(x))
}
func (x *AddNewPlayoffS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AddNewPlayoffS2C_MsgID_value, data, "AddNewPlayoffS2C_MsgID")
	if err != nil {
		return err
	}
	*x = AddNewPlayoffS2C_MsgID(value)
	return nil
}
func (AddNewPlayoffS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor50, []int{8, 0} }

type GetPlayoffTaskRewardC2S_MsgID int32

const (
	GetPlayoffTaskRewardC2S_eMsgID GetPlayoffTaskRewardC2S_MsgID = 12563
)

var GetPlayoffTaskRewardC2S_MsgID_name = map[int32]string{
	12563: "eMsgID",
}
var GetPlayoffTaskRewardC2S_MsgID_value = map[string]int32{
	"eMsgID": 12563,
}

func (x GetPlayoffTaskRewardC2S_MsgID) Enum() *GetPlayoffTaskRewardC2S_MsgID {
	p := new(GetPlayoffTaskRewardC2S_MsgID)
	*p = x
	return p
}
func (x GetPlayoffTaskRewardC2S_MsgID) String() string {
	return proto.EnumName(GetPlayoffTaskRewardC2S_MsgID_name, int32(x))
}
func (x *GetPlayoffTaskRewardC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetPlayoffTaskRewardC2S_MsgID_value, data, "GetPlayoffTaskRewardC2S_MsgID")
	if err != nil {
		return err
	}
	*x = GetPlayoffTaskRewardC2S_MsgID(value)
	return nil
}
func (GetPlayoffTaskRewardC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor50, []int{9, 0}
}

type GetPlayoffTaskRewardS2C_MsgID int32

const (
	GetPlayoffTaskRewardS2C_eMsgID GetPlayoffTaskRewardS2C_MsgID = 12563
)

var GetPlayoffTaskRewardS2C_MsgID_name = map[int32]string{
	12563: "eMsgID",
}
var GetPlayoffTaskRewardS2C_MsgID_value = map[string]int32{
	"eMsgID": 12563,
}

func (x GetPlayoffTaskRewardS2C_MsgID) Enum() *GetPlayoffTaskRewardS2C_MsgID {
	p := new(GetPlayoffTaskRewardS2C_MsgID)
	*p = x
	return p
}
func (x GetPlayoffTaskRewardS2C_MsgID) String() string {
	return proto.EnumName(GetPlayoffTaskRewardS2C_MsgID_name, int32(x))
}
func (x *GetPlayoffTaskRewardS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetPlayoffTaskRewardS2C_MsgID_value, data, "GetPlayoffTaskRewardS2C_MsgID")
	if err != nil {
		return err
	}
	*x = GetPlayoffTaskRewardS2C_MsgID(value)
	return nil
}
func (GetPlayoffTaskRewardS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor50, []int{10, 0}
}

// 季后赛离散数据
type PlayoffDiscreteData struct {
	NextRefreshTime  *int64 `protobuf:"varint,1,opt,name=nextRefreshTime" json:"nextRefreshTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayoffDiscreteData) Reset()                    { *m = PlayoffDiscreteData{} }
func (m *PlayoffDiscreteData) String() string            { return proto.CompactTextString(m) }
func (*PlayoffDiscreteData) ProtoMessage()               {}
func (*PlayoffDiscreteData) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{0} }

func (m *PlayoffDiscreteData) GetNextRefreshTime() int64 {
	if m != nil && m.NextRefreshTime != nil {
		return *m.NextRefreshTime
	}
	return 0
}

// 单个季后赛赛季数据
type SinglePlayoffData struct {
	SeasonCfgId        *int32              `protobuf:"varint,1,opt,name=seasonCfgId" json:"seasonCfgId,omitempty"`
	MatchInfoList      []*PlayoffMatchInfo `protobuf:"bytes,2,rep,name=matchInfoList" json:"matchInfoList,omitempty"`
	TakeTheCrownIsShow *bool               `protobuf:"varint,3,opt,name=takeTheCrownIsShow" json:"takeTheCrownIsShow,omitempty"`
	RewardIds          []int32             `protobuf:"varint,4,rep,name=rewardIds" json:"rewardIds,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *SinglePlayoffData) Reset()                    { *m = SinglePlayoffData{} }
func (m *SinglePlayoffData) String() string            { return proto.CompactTextString(m) }
func (*SinglePlayoffData) ProtoMessage()               {}
func (*SinglePlayoffData) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{1} }

func (m *SinglePlayoffData) GetSeasonCfgId() int32 {
	if m != nil && m.SeasonCfgId != nil {
		return *m.SeasonCfgId
	}
	return 0
}

func (m *SinglePlayoffData) GetMatchInfoList() []*PlayoffMatchInfo {
	if m != nil {
		return m.MatchInfoList
	}
	return nil
}

func (m *SinglePlayoffData) GetTakeTheCrownIsShow() bool {
	if m != nil && m.TakeTheCrownIsShow != nil {
		return *m.TakeTheCrownIsShow
	}
	return false
}

func (m *SinglePlayoffData) GetRewardIds() []int32 {
	if m != nil {
		return m.RewardIds
	}
	return nil
}

// 季后赛比赛信息
type PlayoffMatchInfo struct {
	CfgId            *int32 `protobuf:"varint,1,opt,name=cfgId" json:"cfgId,omitempty"`
	MyScore          *int32 `protobuf:"varint,2,opt,name=myScore" json:"myScore,omitempty"`
	OppScore         *int32 `protobuf:"varint,3,opt,name=oppScore" json:"oppScore,omitempty"`
	FightCount       *int32 `protobuf:"varint,4,opt,name=fightCount" json:"fightCount,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayoffMatchInfo) Reset()                    { *m = PlayoffMatchInfo{} }
func (m *PlayoffMatchInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayoffMatchInfo) ProtoMessage()               {}
func (*PlayoffMatchInfo) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{2} }

func (m *PlayoffMatchInfo) GetCfgId() int32 {
	if m != nil && m.CfgId != nil {
		return *m.CfgId
	}
	return 0
}

func (m *PlayoffMatchInfo) GetMyScore() int32 {
	if m != nil && m.MyScore != nil {
		return *m.MyScore
	}
	return 0
}

func (m *PlayoffMatchInfo) GetOppScore() int32 {
	if m != nil && m.OppScore != nil {
		return *m.OppScore
	}
	return 0
}

func (m *PlayoffMatchInfo) GetFightCount() int32 {
	if m != nil && m.FightCount != nil {
		return *m.FightCount
	}
	return 0
}

// 更新季后赛比赛信息
type UpdatePlayoffMatchInfoS2C struct {
	MatchInfo        *PlayoffMatchInfo `protobuf:"bytes,1,opt,name=matchInfo" json:"matchInfo,omitempty"`
	SeasonId         *int32            `protobuf:"varint,2,opt,name=seasonId" json:"seasonId,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *UpdatePlayoffMatchInfoS2C) Reset()                    { *m = UpdatePlayoffMatchInfoS2C{} }
func (m *UpdatePlayoffMatchInfoS2C) String() string            { return proto.CompactTextString(m) }
func (*UpdatePlayoffMatchInfoS2C) ProtoMessage()               {}
func (*UpdatePlayoffMatchInfoS2C) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{3} }

func (m *UpdatePlayoffMatchInfoS2C) GetMatchInfo() *PlayoffMatchInfo {
	if m != nil {
		return m.MatchInfo
	}
	return nil
}

func (m *UpdatePlayoffMatchInfoS2C) GetSeasonId() int32 {
	if m != nil && m.SeasonId != nil {
		return *m.SeasonId
	}
	return 0
}

// 登录推送
type LoginSendPlayoffS2C struct {
	PlayoffDataList  []*SinglePlayoffData `protobuf:"bytes,1,rep,name=playoffDataList" json:"playoffDataList,omitempty"`
	DiscreteData     *PlayoffDiscreteData `protobuf:"bytes,2,opt,name=discreteData" json:"discreteData,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *LoginSendPlayoffS2C) Reset()                    { *m = LoginSendPlayoffS2C{} }
func (m *LoginSendPlayoffS2C) String() string            { return proto.CompactTextString(m) }
func (*LoginSendPlayoffS2C) ProtoMessage()               {}
func (*LoginSendPlayoffS2C) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{4} }

func (m *LoginSendPlayoffS2C) GetPlayoffDataList() []*SinglePlayoffData {
	if m != nil {
		return m.PlayoffDataList
	}
	return nil
}

func (m *LoginSendPlayoffS2C) GetDiscreteData() *PlayoffDiscreteData {
	if m != nil {
		return m.DiscreteData
	}
	return nil
}

// 季后赛扫荡
type PlayoffSweepC2S struct {
	ScheduleId       *int32 `protobuf:"varint,1,opt,name=scheduleId" json:"scheduleId,omitempty"`
	SweepCount       *int32 `protobuf:"varint,2,opt,name=sweepCount" json:"sweepCount,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayoffSweepC2S) Reset()                    { *m = PlayoffSweepC2S{} }
func (m *PlayoffSweepC2S) String() string            { return proto.CompactTextString(m) }
func (*PlayoffSweepC2S) ProtoMessage()               {}
func (*PlayoffSweepC2S) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{5} }

func (m *PlayoffSweepC2S) GetScheduleId() int32 {
	if m != nil && m.ScheduleId != nil {
		return *m.ScheduleId
	}
	return 0
}

func (m *PlayoffSweepC2S) GetSweepCount() int32 {
	if m != nil && m.SweepCount != nil {
		return *m.SweepCount
	}
	return 0
}

// 季后赛扫荡
type PlayoffSweepS2C struct {
	SeasonId             *int32                `protobuf:"varint,1,opt,name=seasonId" json:"seasonId,omitempty"`
	RegularScheduleCfgId *int32                `protobuf:"varint,2,opt,name=regularScheduleCfgId" json:"regularScheduleCfgId,omitempty"`
	ItemList             []*ItemVaryConfigList `protobuf:"bytes,3,rep,name=itemList" json:"itemList,omitempty"`
	Count                *int32                `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized     []byte                `json:"-"`
}

func (m *PlayoffSweepS2C) Reset()                    { *m = PlayoffSweepS2C{} }
func (m *PlayoffSweepS2C) String() string            { return proto.CompactTextString(m) }
func (*PlayoffSweepS2C) ProtoMessage()               {}
func (*PlayoffSweepS2C) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{6} }

func (m *PlayoffSweepS2C) GetSeasonId() int32 {
	if m != nil && m.SeasonId != nil {
		return *m.SeasonId
	}
	return 0
}

func (m *PlayoffSweepS2C) GetRegularScheduleCfgId() int32 {
	if m != nil && m.RegularScheduleCfgId != nil {
		return *m.RegularScheduleCfgId
	}
	return 0
}

func (m *PlayoffSweepS2C) GetItemList() []*ItemVaryConfigList {
	if m != nil {
		return m.ItemList
	}
	return nil
}

func (m *PlayoffSweepS2C) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

// 通知客户端刷新季后赛战斗次数
type RefreshPlayoffOnlineS2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshPlayoffOnlineS2C) Reset()                    { *m = RefreshPlayoffOnlineS2C{} }
func (m *RefreshPlayoffOnlineS2C) String() string            { return proto.CompactTextString(m) }
func (*RefreshPlayoffOnlineS2C) ProtoMessage()               {}
func (*RefreshPlayoffOnlineS2C) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{7} }

// 新增季后赛
type AddNewPlayoffS2C struct {
	Data             *SinglePlayoffData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *AddNewPlayoffS2C) Reset()                    { *m = AddNewPlayoffS2C{} }
func (m *AddNewPlayoffS2C) String() string            { return proto.CompactTextString(m) }
func (*AddNewPlayoffS2C) ProtoMessage()               {}
func (*AddNewPlayoffS2C) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{8} }

func (m *AddNewPlayoffS2C) GetData() *SinglePlayoffData {
	if m != nil {
		return m.Data
	}
	return nil
}

// 领季后赛任务奖励
type GetPlayoffTaskRewardC2S struct {
	// 赛季id
	SeasonId *int32 `protobuf:"varint,1,opt,name=seasonId" json:"seasonId,omitempty"`
	// 任务id
	TaskId           *int32 `protobuf:"varint,2,opt,name=taskId" json:"taskId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetPlayoffTaskRewardC2S) Reset()                    { *m = GetPlayoffTaskRewardC2S{} }
func (m *GetPlayoffTaskRewardC2S) String() string            { return proto.CompactTextString(m) }
func (*GetPlayoffTaskRewardC2S) ProtoMessage()               {}
func (*GetPlayoffTaskRewardC2S) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{9} }

func (m *GetPlayoffTaskRewardC2S) GetSeasonId() int32 {
	if m != nil && m.SeasonId != nil {
		return *m.SeasonId
	}
	return 0
}

func (m *GetPlayoffTaskRewardC2S) GetTaskId() int32 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

// 领季后赛任务奖励
type GetPlayoffTaskRewardS2C struct {
	// 赛季id
	SeasonId *int32 `protobuf:"varint,1,opt,name=seasonId" json:"seasonId,omitempty"`
	// 任务id
	TaskId           *int32 `protobuf:"varint,2,opt,name=taskId" json:"taskId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetPlayoffTaskRewardS2C) Reset()                    { *m = GetPlayoffTaskRewardS2C{} }
func (m *GetPlayoffTaskRewardS2C) String() string            { return proto.CompactTextString(m) }
func (*GetPlayoffTaskRewardS2C) ProtoMessage()               {}
func (*GetPlayoffTaskRewardS2C) Descriptor() ([]byte, []int) { return fileDescriptor50, []int{10} }

func (m *GetPlayoffTaskRewardS2C) GetSeasonId() int32 {
	if m != nil && m.SeasonId != nil {
		return *m.SeasonId
	}
	return 0
}

func (m *GetPlayoffTaskRewardS2C) GetTaskId() int32 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func init() {
	proto.RegisterType((*PlayoffDiscreteData)(nil), "PlayoffDiscreteData")
	proto.RegisterType((*SinglePlayoffData)(nil), "SinglePlayoffData")
	proto.RegisterType((*PlayoffMatchInfo)(nil), "PlayoffMatchInfo")
	proto.RegisterType((*UpdatePlayoffMatchInfoS2C)(nil), "UpdatePlayoffMatchInfoS2C")
	proto.RegisterType((*LoginSendPlayoffS2C)(nil), "LoginSendPlayoffS2C")
	proto.RegisterType((*PlayoffSweepC2S)(nil), "PlayoffSweepC2S")
	proto.RegisterType((*PlayoffSweepS2C)(nil), "PlayoffSweepS2C")
	proto.RegisterType((*RefreshPlayoffOnlineS2C)(nil), "RefreshPlayoffOnlineS2C")
	proto.RegisterType((*AddNewPlayoffS2C)(nil), "AddNewPlayoffS2C")
	proto.RegisterType((*GetPlayoffTaskRewardC2S)(nil), "GetPlayoffTaskRewardC2S")
	proto.RegisterType((*GetPlayoffTaskRewardS2C)(nil), "GetPlayoffTaskRewardS2C")
	proto.RegisterEnum("UpdatePlayoffMatchInfoS2C_MsgID", UpdatePlayoffMatchInfoS2C_MsgID_name, UpdatePlayoffMatchInfoS2C_MsgID_value)
	proto.RegisterEnum("LoginSendPlayoffS2C_MsgID", LoginSendPlayoffS2C_MsgID_name, LoginSendPlayoffS2C_MsgID_value)
	proto.RegisterEnum("PlayoffSweepC2S_MsgID", PlayoffSweepC2S_MsgID_name, PlayoffSweepC2S_MsgID_value)
	proto.RegisterEnum("PlayoffSweepS2C_MsgID", PlayoffSweepS2C_MsgID_name, PlayoffSweepS2C_MsgID_value)
	proto.RegisterEnum("RefreshPlayoffOnlineS2C_MsgID", RefreshPlayoffOnlineS2C_MsgID_name, RefreshPlayoffOnlineS2C_MsgID_value)
	proto.RegisterEnum("AddNewPlayoffS2C_MsgID", AddNewPlayoffS2C_MsgID_name, AddNewPlayoffS2C_MsgID_value)
	proto.RegisterEnum("GetPlayoffTaskRewardC2S_MsgID", GetPlayoffTaskRewardC2S_MsgID_name, GetPlayoffTaskRewardC2S_MsgID_value)
	proto.RegisterEnum("GetPlayoffTaskRewardS2C_MsgID", GetPlayoffTaskRewardS2C_MsgID_name, GetPlayoffTaskRewardS2C_MsgID_value)
}

func init() { proto.RegisterFile("playoff.proto", fileDescriptor50) }

var fileDescriptor50 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0x76, 0xbb, 0xee, 0xde, 0x5a, 0xdb, 0xa6, 0x85, 0xd6, 0xc5, 0x87, 0x12, 0x14, 0x8a,
	0x42, 0x84, 0xfe, 0x03, 0x49, 0x41, 0x22, 0xbb, 0xae, 0x6e, 0xaa, 0xf8, 0x3a, 0xcd, 0xdc, 0x24,
	0xc3, 0xb6, 0x33, 0x61, 0x66, 0x4a, 0xed, 0xa3, 0xe0, 0xab, 0xdf, 0x7f, 0xd8, 0xc9, 0x24, 0x5d,
	0xab, 0xcd, 0x8a, 0x6f, 0xc3, 0xbd, 0x77, 0xce, 0xb9, 0xe7, 0x9c, 0x0b, 0xed, 0x7c, 0x49, 0xb6,
	0x22, 0x49, 0xfc, 0x5c, 0x0a, 0x2d, 0xce, 0x20, 0x16, 0xbc, 0x7a, 0x7b, 0x3e, 0xf4, 0x5f, 0x97,
	0xcd, 0x19, 0x53, 0xb1, 0x44, 0x8d, 0x33, 0xa2, 0x89, 0x3b, 0x84, 0x0e, 0xc7, 0x0f, 0xfa, 0x0a,
	0x13, 0x89, 0x2a, 0x9b, 0xb3, 0x15, 0x8e, 0x9c, 0xb1, 0x33, 0x69, 0x78, 0x1f, 0x1d, 0xe8, 0x45,
	0x8c, 0xa7, 0x4b, 0xdc, 0x7d, 0x2b, 0xc6, 0xfb, 0xd0, 0x52, 0x48, 0x94, 0xe0, 0x41, 0x92, 0x86,
	0xd4, 0x8e, 0x36, 0xdd, 0x09, 0xb4, 0x57, 0x44, 0xc7, 0x59, 0xc8, 0x13, 0x71, 0xce, 0x94, 0x1e,
	0xdd, 0x19, 0x37, 0x26, 0xad, 0x69, 0xcf, 0xaf, 0x7e, 0x5e, 0xec, 0x9a, 0xee, 0x19, 0xb8, 0x9a,
	0x5c, 0xe3, 0x3c, 0xc3, 0x40, 0x8a, 0x0d, 0x0f, 0x55, 0x94, 0x89, 0xcd, 0xa8, 0x61, 0x50, 0x4e,
	0xdc, 0x1e, 0x9c, 0x4a, 0xdc, 0x10, 0x49, 0x43, 0xaa, 0x46, 0x47, 0x06, 0xa1, 0xe9, 0xbd, 0x87,
	0xee, 0x01, 0x44, 0x1b, 0x9a, 0xf1, 0x1e, 0x77, 0x07, 0xee, 0xae, 0xb6, 0x51, 0x2c, 0x24, 0x1a,
	0xd6, 0xa2, 0xd0, 0x85, 0x13, 0x91, 0xe7, 0x65, 0xa5, 0x61, 0x2b, 0x2e, 0x40, 0xc2, 0xd2, 0x4c,
	0x07, 0x62, 0xcd, 0xb5, 0x41, 0x36, 0x35, 0x8f, 0xc1, 0x83, 0xb7, 0x39, 0x25, 0x1a, 0xff, 0xc6,
	0x8f, 0xa6, 0x81, 0xfb, 0x08, 0x4e, 0x6f, 0xf4, 0x58, 0x9a, 0x5a, 0x2d, 0x86, 0xa8, 0xb4, 0xc2,
	0xec, 0x62, 0xa9, 0xbd, 0x01, 0x34, 0x2f, 0x54, 0x1a, 0xce, 0xdc, 0x16, 0x1c, 0xa3, 0x7d, 0x75,
	0xbf, 0x2c, 0xbc, 0x4f, 0x0e, 0xf4, 0xcf, 0x45, 0xca, 0x78, 0x84, 0x9c, 0x56, 0x28, 0x05, 0xcb,
	0x53, 0xe8, 0xe4, 0xbf, 0x9d, 0xb5, 0xbe, 0x39, 0xd6, 0x37, 0xd7, 0x3f, 0xf4, 0xfd, 0x09, 0xdc,
	0xa3, 0x7b, 0xb1, 0x59, 0xc2, 0xd6, 0x74, 0xe0, 0xd7, 0x44, 0x5a, 0xb7, 0xc6, 0xd7, 0x85, 0x77,
	0x09, 0x9d, 0x1d, 0xf9, 0x06, 0x31, 0x0f, 0xa6, 0x51, 0x61, 0x8c, 0x8a, 0x33, 0xa4, 0xeb, 0x25,
	0xde, 0xf8, 0x59, 0xd4, 0x6c, 0xdf, 0x9a, 0x75, 0xab, 0xae, 0x6f, 0x0b, 0xef, 0xb3, 0xf3, 0x27,
	0x62, 0xa1, 0x69, 0xdf, 0x93, 0x12, 0xef, 0x21, 0x0c, 0x24, 0xa6, 0xeb, 0x25, 0x91, 0x51, 0x45,
	0x55, 0x5e, 0x4e, 0x19, 0xd6, 0x63, 0x38, 0x61, 0x1a, 0x57, 0x56, 0x7c, 0xc3, 0x8a, 0xef, 0xfb,
	0xa1, 0x29, 0xbc, 0x23, 0x72, 0x1b, 0x98, 0xdb, 0x65, 0x69, 0xd1, 0xb2, 0x99, 0xef, 0x85, 0x57,
	0xbf, 0xcf, 0x33, 0x18, 0x56, 0x57, 0x5c, 0x6d, 0x75, 0xc9, 0x97, 0x8c, 0xa3, 0x59, 0xab, 0xee,
	0xc3, 0x8f, 0x85, 0xf7, 0x12, 0xba, 0xcf, 0x29, 0x7d, 0x85, 0x9b, 0xbd, 0x50, 0xc6, 0x70, 0x44,
	0x0b, 0x7f, 0xcb, 0xd4, 0x6b, 0x92, 0xa8, 0xc3, 0xfa, 0xbe, 0xf0, 0xde, 0xc0, 0xf0, 0x05, 0xea,
	0x6a, 0x6e, 0x4e, 0xd4, 0xf5, 0x95, 0x3d, 0xe5, 0xc2, 0xe5, 0x43, 0x4f, 0xee, 0xc3, 0xb1, 0x36,
	0x23, 0xff, 0xba, 0x9b, 0x9f, 0xb7, 0x42, 0xd6, 0xdb, 0xfc, 0x5f, 0x90, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x7f, 0xc8, 0x2f, 0x54, 0x1f, 0x04, 0x00, 0x00,
}