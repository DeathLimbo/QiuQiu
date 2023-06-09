// Code generated by protoc-gen-go.
// source: mascot.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MascotGetDataC2S_MsgID int32

const (
	MascotGetDataC2S_eMsgID MascotGetDataC2S_MsgID = 14337
)

var MascotGetDataC2S_MsgID_name = map[int32]string{
	14337: "eMsgID",
}
var MascotGetDataC2S_MsgID_value = map[string]int32{
	"eMsgID": 14337,
}

func (x MascotGetDataC2S_MsgID) Enum() *MascotGetDataC2S_MsgID {
	p := new(MascotGetDataC2S_MsgID)
	*p = x
	return p
}
func (x MascotGetDataC2S_MsgID) String() string {
	return proto.EnumName(MascotGetDataC2S_MsgID_name, int32(x))
}
func (x *MascotGetDataC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotGetDataC2S_MsgID_value, data, "MascotGetDataC2S_MsgID")
	if err != nil {
		return err
	}
	*x = MascotGetDataC2S_MsgID(value)
	return nil
}
func (MascotGetDataC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{3, 0} }

type MascotGetDataS2C_MsgID int32

const (
	MascotGetDataS2C_eMsgID MascotGetDataS2C_MsgID = 14337
)

var MascotGetDataS2C_MsgID_name = map[int32]string{
	14337: "eMsgID",
}
var MascotGetDataS2C_MsgID_value = map[string]int32{
	"eMsgID": 14337,
}

func (x MascotGetDataS2C_MsgID) Enum() *MascotGetDataS2C_MsgID {
	p := new(MascotGetDataS2C_MsgID)
	*p = x
	return p
}
func (x MascotGetDataS2C_MsgID) String() string {
	return proto.EnumName(MascotGetDataS2C_MsgID_name, int32(x))
}
func (x *MascotGetDataS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotGetDataS2C_MsgID_value, data, "MascotGetDataS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotGetDataS2C_MsgID(value)
	return nil
}
func (MascotGetDataS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{4, 0} }

type MascotResetCntS2C_MsgID int32

const (
	MascotResetCntS2C_eMsgID MascotResetCntS2C_MsgID = 14338
)

var MascotResetCntS2C_MsgID_name = map[int32]string{
	14338: "eMsgID",
}
var MascotResetCntS2C_MsgID_value = map[string]int32{
	"eMsgID": 14338,
}

func (x MascotResetCntS2C_MsgID) Enum() *MascotResetCntS2C_MsgID {
	p := new(MascotResetCntS2C_MsgID)
	*p = x
	return p
}
func (x MascotResetCntS2C_MsgID) String() string {
	return proto.EnumName(MascotResetCntS2C_MsgID_name, int32(x))
}
func (x *MascotResetCntS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotResetCntS2C_MsgID_value, data, "MascotResetCntS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotResetCntS2C_MsgID(value)
	return nil
}
func (MascotResetCntS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{5, 0} }

type MascotSettleDataS2C_MsgID int32

const (
	MascotSettleDataS2C_eMsgID MascotSettleDataS2C_MsgID = 14339
)

var MascotSettleDataS2C_MsgID_name = map[int32]string{
	14339: "eMsgID",
}
var MascotSettleDataS2C_MsgID_value = map[string]int32{
	"eMsgID": 14339,
}

func (x MascotSettleDataS2C_MsgID) Enum() *MascotSettleDataS2C_MsgID {
	p := new(MascotSettleDataS2C_MsgID)
	*p = x
	return p
}
func (x MascotSettleDataS2C_MsgID) String() string {
	return proto.EnumName(MascotSettleDataS2C_MsgID_name, int32(x))
}
func (x *MascotSettleDataS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotSettleDataS2C_MsgID_value, data, "MascotSettleDataS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotSettleDataS2C_MsgID(value)
	return nil
}
func (MascotSettleDataS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor43, []int{6, 0}
}

type MascotGetSectionRewardC2S_MsgID int32

const (
	MascotGetSectionRewardC2S_eMsgID MascotGetSectionRewardC2S_MsgID = 14340
)

var MascotGetSectionRewardC2S_MsgID_name = map[int32]string{
	14340: "eMsgID",
}
var MascotGetSectionRewardC2S_MsgID_value = map[string]int32{
	"eMsgID": 14340,
}

func (x MascotGetSectionRewardC2S_MsgID) Enum() *MascotGetSectionRewardC2S_MsgID {
	p := new(MascotGetSectionRewardC2S_MsgID)
	*p = x
	return p
}
func (x MascotGetSectionRewardC2S_MsgID) String() string {
	return proto.EnumName(MascotGetSectionRewardC2S_MsgID_name, int32(x))
}
func (x *MascotGetSectionRewardC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotGetSectionRewardC2S_MsgID_value, data, "MascotGetSectionRewardC2S_MsgID")
	if err != nil {
		return err
	}
	*x = MascotGetSectionRewardC2S_MsgID(value)
	return nil
}
func (MascotGetSectionRewardC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor43, []int{7, 0}
}

type MascotGetSectionRewardS2C_MsgID int32

const (
	MascotGetSectionRewardS2C_eMsgID MascotGetSectionRewardS2C_MsgID = 14340
)

var MascotGetSectionRewardS2C_MsgID_name = map[int32]string{
	14340: "eMsgID",
}
var MascotGetSectionRewardS2C_MsgID_value = map[string]int32{
	"eMsgID": 14340,
}

func (x MascotGetSectionRewardS2C_MsgID) Enum() *MascotGetSectionRewardS2C_MsgID {
	p := new(MascotGetSectionRewardS2C_MsgID)
	*p = x
	return p
}
func (x MascotGetSectionRewardS2C_MsgID) String() string {
	return proto.EnumName(MascotGetSectionRewardS2C_MsgID_name, int32(x))
}
func (x *MascotGetSectionRewardS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotGetSectionRewardS2C_MsgID_value, data, "MascotGetSectionRewardS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotGetSectionRewardS2C_MsgID(value)
	return nil
}
func (MascotGetSectionRewardS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor43, []int{8, 0}
}

type MascotGetAchieveRewardC2S_MsgID int32

const (
	MascotGetAchieveRewardC2S_eMsgID MascotGetAchieveRewardC2S_MsgID = 14341
)

var MascotGetAchieveRewardC2S_MsgID_name = map[int32]string{
	14341: "eMsgID",
}
var MascotGetAchieveRewardC2S_MsgID_value = map[string]int32{
	"eMsgID": 14341,
}

func (x MascotGetAchieveRewardC2S_MsgID) Enum() *MascotGetAchieveRewardC2S_MsgID {
	p := new(MascotGetAchieveRewardC2S_MsgID)
	*p = x
	return p
}
func (x MascotGetAchieveRewardC2S_MsgID) String() string {
	return proto.EnumName(MascotGetAchieveRewardC2S_MsgID_name, int32(x))
}
func (x *MascotGetAchieveRewardC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotGetAchieveRewardC2S_MsgID_value, data, "MascotGetAchieveRewardC2S_MsgID")
	if err != nil {
		return err
	}
	*x = MascotGetAchieveRewardC2S_MsgID(value)
	return nil
}
func (MascotGetAchieveRewardC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor43, []int{9, 0}
}

type MascotGetAchieveRewardS2C_MsgID int32

const (
	MascotGetAchieveRewardS2C_eMsgID MascotGetAchieveRewardS2C_MsgID = 14341
)

var MascotGetAchieveRewardS2C_MsgID_name = map[int32]string{
	14341: "eMsgID",
}
var MascotGetAchieveRewardS2C_MsgID_value = map[string]int32{
	"eMsgID": 14341,
}

func (x MascotGetAchieveRewardS2C_MsgID) Enum() *MascotGetAchieveRewardS2C_MsgID {
	p := new(MascotGetAchieveRewardS2C_MsgID)
	*p = x
	return p
}
func (x MascotGetAchieveRewardS2C_MsgID) String() string {
	return proto.EnumName(MascotGetAchieveRewardS2C_MsgID_name, int32(x))
}
func (x *MascotGetAchieveRewardS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotGetAchieveRewardS2C_MsgID_value, data, "MascotGetAchieveRewardS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotGetAchieveRewardS2C_MsgID(value)
	return nil
}
func (MascotGetAchieveRewardS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor43, []int{10, 0}
}

type MascotSweepC2S_MsgID int32

const (
	MascotSweepC2S_eMsgID MascotSweepC2S_MsgID = 14342
)

var MascotSweepC2S_MsgID_name = map[int32]string{
	14342: "eMsgID",
}
var MascotSweepC2S_MsgID_value = map[string]int32{
	"eMsgID": 14342,
}

func (x MascotSweepC2S_MsgID) Enum() *MascotSweepC2S_MsgID {
	p := new(MascotSweepC2S_MsgID)
	*p = x
	return p
}
func (x MascotSweepC2S_MsgID) String() string {
	return proto.EnumName(MascotSweepC2S_MsgID_name, int32(x))
}
func (x *MascotSweepC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotSweepC2S_MsgID_value, data, "MascotSweepC2S_MsgID")
	if err != nil {
		return err
	}
	*x = MascotSweepC2S_MsgID(value)
	return nil
}
func (MascotSweepC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{11, 0} }

type MascotSweepS2C_MsgID int32

const (
	MascotSweepS2C_eMsgID MascotSweepS2C_MsgID = 14342
)

var MascotSweepS2C_MsgID_name = map[int32]string{
	14342: "eMsgID",
}
var MascotSweepS2C_MsgID_value = map[string]int32{
	"eMsgID": 14342,
}

func (x MascotSweepS2C_MsgID) Enum() *MascotSweepS2C_MsgID {
	p := new(MascotSweepS2C_MsgID)
	*p = x
	return p
}
func (x MascotSweepS2C_MsgID) String() string {
	return proto.EnumName(MascotSweepS2C_MsgID_name, int32(x))
}
func (x *MascotSweepS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotSweepS2C_MsgID_value, data, "MascotSweepS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotSweepS2C_MsgID(value)
	return nil
}
func (MascotSweepS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{12, 0} }

type MascotBuyCntC2S_MsgID int32

const (
	MascotBuyCntC2S_eMsgID MascotBuyCntC2S_MsgID = 14343
)

var MascotBuyCntC2S_MsgID_name = map[int32]string{
	14343: "eMsgID",
}
var MascotBuyCntC2S_MsgID_value = map[string]int32{
	"eMsgID": 14343,
}

func (x MascotBuyCntC2S_MsgID) Enum() *MascotBuyCntC2S_MsgID {
	p := new(MascotBuyCntC2S_MsgID)
	*p = x
	return p
}
func (x MascotBuyCntC2S_MsgID) String() string {
	return proto.EnumName(MascotBuyCntC2S_MsgID_name, int32(x))
}
func (x *MascotBuyCntC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotBuyCntC2S_MsgID_value, data, "MascotBuyCntC2S_MsgID")
	if err != nil {
		return err
	}
	*x = MascotBuyCntC2S_MsgID(value)
	return nil
}
func (MascotBuyCntC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{13, 0} }

type MascotBuyCntS2C_MsgID int32

const (
	MascotBuyCntS2C_eMsgID MascotBuyCntS2C_MsgID = 14343
)

var MascotBuyCntS2C_MsgID_name = map[int32]string{
	14343: "eMsgID",
}
var MascotBuyCntS2C_MsgID_value = map[string]int32{
	"eMsgID": 14343,
}

func (x MascotBuyCntS2C_MsgID) Enum() *MascotBuyCntS2C_MsgID {
	p := new(MascotBuyCntS2C_MsgID)
	*p = x
	return p
}
func (x MascotBuyCntS2C_MsgID) String() string {
	return proto.EnumName(MascotBuyCntS2C_MsgID_name, int32(x))
}
func (x *MascotBuyCntS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MascotBuyCntS2C_MsgID_value, data, "MascotBuyCntS2C_MsgID")
	if err != nil {
		return err
	}
	*x = MascotBuyCntS2C_MsgID(value)
	return nil
}
func (MascotBuyCntS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor43, []int{14, 0} }

// 吉祥物关卡
type MascotChapterData struct {
	// 关卡id
	Id *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// 是否通关
	IsPass           *bool  `protobuf:"varint,2,opt,name=is_pass" json:"is_pass,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotChapterData) Reset()                    { *m = MascotChapterData{} }
func (m *MascotChapterData) String() string            { return proto.CompactTextString(m) }
func (*MascotChapterData) ProtoMessage()               {}
func (*MascotChapterData) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{0} }

func (m *MascotChapterData) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *MascotChapterData) GetIsPass() bool {
	if m != nil && m.IsPass != nil {
		return *m.IsPass
	}
	return false
}

// 吉祥物基础数据
type MascotBaseData struct {
	// 挑战次数下次刷新时间
	NextRefreshTime *int64 `protobuf:"varint,1,opt,name=nextRefreshTime" json:"nextRefreshTime,omitempty"`
	// 当日已购买次数
	DailyBuyCnt *int32 `protobuf:"varint,2,opt,name=dailyBuyCnt" json:"dailyBuyCnt,omitempty"`
	// 当日已挑战次数（总次数为初始次数+购买次数）
	DailyFightCnt *int32 `protobuf:"varint,3,opt,name=dailyFightCnt" json:"dailyFightCnt,omitempty"`
	// 累计挑战次数
	TotalFightCnt *int32 `protobuf:"varint,4,opt,name=totalFightCnt" json:"totalFightCnt,omitempty"`
	// 已领取章节奖励的章节id列表
	SectionIds []int32 `protobuf:"varint,5,rep,name=sectionIds" json:"sectionIds,omitempty"`
	// 已领取成就奖励的成就id列表
	AchieveTasks     []*TaskItem `protobuf:"bytes,6,rep,name=achieveTasks" json:"achieveTasks,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MascotBaseData) Reset()                    { *m = MascotBaseData{} }
func (m *MascotBaseData) String() string            { return proto.CompactTextString(m) }
func (*MascotBaseData) ProtoMessage()               {}
func (*MascotBaseData) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{1} }

func (m *MascotBaseData) GetNextRefreshTime() int64 {
	if m != nil && m.NextRefreshTime != nil {
		return *m.NextRefreshTime
	}
	return 0
}

func (m *MascotBaseData) GetDailyBuyCnt() int32 {
	if m != nil && m.DailyBuyCnt != nil {
		return *m.DailyBuyCnt
	}
	return 0
}

func (m *MascotBaseData) GetDailyFightCnt() int32 {
	if m != nil && m.DailyFightCnt != nil {
		return *m.DailyFightCnt
	}
	return 0
}

func (m *MascotBaseData) GetTotalFightCnt() int32 {
	if m != nil && m.TotalFightCnt != nil {
		return *m.TotalFightCnt
	}
	return 0
}

func (m *MascotBaseData) GetSectionIds() []int32 {
	if m != nil {
		return m.SectionIds
	}
	return nil
}

func (m *MascotBaseData) GetAchieveTasks() []*TaskItem {
	if m != nil {
		return m.AchieveTasks
	}
	return nil
}

// 吉祥物数据（服务器）
type MascotData struct {
	// 吉祥物基础数据
	Bases *MascotBaseData `protobuf:"bytes,1,opt,name=bases" json:"bases,omitempty"`
	// 吉祥物关卡
	Chapters         []*MascotChapterData `protobuf:"bytes,2,rep,name=chapters" json:"chapters,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *MascotData) Reset()                    { *m = MascotData{} }
func (m *MascotData) String() string            { return proto.CompactTextString(m) }
func (*MascotData) ProtoMessage()               {}
func (*MascotData) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{2} }

func (m *MascotData) GetBases() *MascotBaseData {
	if m != nil {
		return m.Bases
	}
	return nil
}

func (m *MascotData) GetChapters() []*MascotChapterData {
	if m != nil {
		return m.Chapters
	}
	return nil
}

// 客户端无数据时，在功能开放的前提下，请求数据
type MascotGetDataC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotGetDataC2S) Reset()                    { *m = MascotGetDataC2S{} }
func (m *MascotGetDataC2S) String() string            { return proto.CompactTextString(m) }
func (*MascotGetDataC2S) ProtoMessage()               {}
func (*MascotGetDataC2S) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{3} }

// 上线时，通知吉祥物基础数据（客户端无数据时，请求应答）
type MascotGetDataS2C struct {
	// 吉祥物基础数据
	Bases *MascotBaseData `protobuf:"bytes,1,opt,name=bases" json:"bases,omitempty"`
	// 吉祥物关卡数据
	Chapters         []*MascotChapterData `protobuf:"bytes,2,rep,name=chapters" json:"chapters,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *MascotGetDataS2C) Reset()                    { *m = MascotGetDataS2C{} }
func (m *MascotGetDataS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotGetDataS2C) ProtoMessage()               {}
func (*MascotGetDataS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{4} }

func (m *MascotGetDataS2C) GetBases() *MascotBaseData {
	if m != nil {
		return m.Bases
	}
	return nil
}

func (m *MascotGetDataS2C) GetChapters() []*MascotChapterData {
	if m != nil {
		return m.Chapters
	}
	return nil
}

// 定时刷新（重置）吉祥物挑战次数
type MascotResetCntS2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotResetCntS2C) Reset()                    { *m = MascotResetCntS2C{} }
func (m *MascotResetCntS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotResetCntS2C) ProtoMessage()               {}
func (*MascotResetCntS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{5} }

// 比赛结束
type MascotSettleDataS2C struct {
	// 关卡
	Chapter *MascotChapterData `protobuf:"bytes,1,opt,name=chapter" json:"chapter,omitempty"`
	// 首通奖励
	FirstPassRewards []*ItemVaryConfig `protobuf:"bytes,2,rep,name=firstPassRewards" json:"firstPassRewards,omitempty"`
	// 比赛结果
	IsWin            *bool  `protobuf:"varint,3,opt,name=isWin" json:"isWin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotSettleDataS2C) Reset()                    { *m = MascotSettleDataS2C{} }
func (m *MascotSettleDataS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotSettleDataS2C) ProtoMessage()               {}
func (*MascotSettleDataS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{6} }

func (m *MascotSettleDataS2C) GetChapter() *MascotChapterData {
	if m != nil {
		return m.Chapter
	}
	return nil
}

func (m *MascotSettleDataS2C) GetFirstPassRewards() []*ItemVaryConfig {
	if m != nil {
		return m.FirstPassRewards
	}
	return nil
}

func (m *MascotSettleDataS2C) GetIsWin() bool {
	if m != nil && m.IsWin != nil {
		return *m.IsWin
	}
	return false
}

// 领取章节奖励c2s
type MascotGetSectionRewardC2S struct {
	// 章节id
	SectionId        *int32 `protobuf:"varint,1,opt,name=sectionId" json:"sectionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotGetSectionRewardC2S) Reset()                    { *m = MascotGetSectionRewardC2S{} }
func (m *MascotGetSectionRewardC2S) String() string            { return proto.CompactTextString(m) }
func (*MascotGetSectionRewardC2S) ProtoMessage()               {}
func (*MascotGetSectionRewardC2S) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{7} }

func (m *MascotGetSectionRewardC2S) GetSectionId() int32 {
	if m != nil && m.SectionId != nil {
		return *m.SectionId
	}
	return 0
}

// 领取章节奖励s2c
type MascotGetSectionRewardS2C struct {
	// 章节id
	SectionId        *int32 `protobuf:"varint,1,opt,name=sectionId" json:"sectionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotGetSectionRewardS2C) Reset()                    { *m = MascotGetSectionRewardS2C{} }
func (m *MascotGetSectionRewardS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotGetSectionRewardS2C) ProtoMessage()               {}
func (*MascotGetSectionRewardS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{8} }

func (m *MascotGetSectionRewardS2C) GetSectionId() int32 {
	if m != nil && m.SectionId != nil {
		return *m.SectionId
	}
	return 0
}

// 领取成就奖励c2s
type MascotGetAchieveRewardC2S struct {
	// 成就id
	AchieveId        *int32 `protobuf:"varint,2,opt,name=achieveId" json:"achieveId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotGetAchieveRewardC2S) Reset()                    { *m = MascotGetAchieveRewardC2S{} }
func (m *MascotGetAchieveRewardC2S) String() string            { return proto.CompactTextString(m) }
func (*MascotGetAchieveRewardC2S) ProtoMessage()               {}
func (*MascotGetAchieveRewardC2S) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{9} }

func (m *MascotGetAchieveRewardC2S) GetAchieveId() int32 {
	if m != nil && m.AchieveId != nil {
		return *m.AchieveId
	}
	return 0
}

// 领取成就奖励s2c
type MascotGetAchieveRewardS2C struct {
	// 成就id
	AchieveId        *int32 `protobuf:"varint,2,opt,name=achieveId" json:"achieveId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotGetAchieveRewardS2C) Reset()                    { *m = MascotGetAchieveRewardS2C{} }
func (m *MascotGetAchieveRewardS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotGetAchieveRewardS2C) ProtoMessage()               {}
func (*MascotGetAchieveRewardS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{10} }

func (m *MascotGetAchieveRewardS2C) GetAchieveId() int32 {
	if m != nil && m.AchieveId != nil {
		return *m.AchieveId
	}
	return 0
}

// 吉祥物扫荡
type MascotSweepC2S struct {
	// 关卡id
	ChapterId *int32 `protobuf:"varint,1,opt,name=chapterId" json:"chapterId,omitempty"`
	// 扫荡次数
	SweepCnt         *int32 `protobuf:"varint,2,opt,name=sweepCnt" json:"sweepCnt,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotSweepC2S) Reset()                    { *m = MascotSweepC2S{} }
func (m *MascotSweepC2S) String() string            { return proto.CompactTextString(m) }
func (*MascotSweepC2S) ProtoMessage()               {}
func (*MascotSweepC2S) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{11} }

func (m *MascotSweepC2S) GetChapterId() int32 {
	if m != nil && m.ChapterId != nil {
		return *m.ChapterId
	}
	return 0
}

func (m *MascotSweepC2S) GetSweepCnt() int32 {
	if m != nil && m.SweepCnt != nil {
		return *m.SweepCnt
	}
	return 0
}

// 吉祥物扫荡
type MascotSweepS2C struct {
	// 奖励列表（列表长度为扫荡次数）
	Rewards          []*ItemVaryConfigList `protobuf:"bytes,1,rep,name=rewards" json:"rewards,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *MascotSweepS2C) Reset()                    { *m = MascotSweepS2C{} }
func (m *MascotSweepS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotSweepS2C) ProtoMessage()               {}
func (*MascotSweepS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{12} }

func (m *MascotSweepS2C) GetRewards() []*ItemVaryConfigList {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// 吉祥物购买挑战次数
type MascotBuyCntC2S struct {
	// 购买次数
	BuyCnt           *int32 `protobuf:"varint,1,opt,name=buyCnt" json:"buyCnt,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotBuyCntC2S) Reset()                    { *m = MascotBuyCntC2S{} }
func (m *MascotBuyCntC2S) String() string            { return proto.CompactTextString(m) }
func (*MascotBuyCntC2S) ProtoMessage()               {}
func (*MascotBuyCntC2S) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{13} }

func (m *MascotBuyCntC2S) GetBuyCnt() int32 {
	if m != nil && m.BuyCnt != nil {
		return *m.BuyCnt
	}
	return 0
}

// 吉祥物购买挑战次数
type MascotBuyCntS2C struct {
	// 已购买挑战次数
	DailyBuyCnt      *int32 `protobuf:"varint,1,opt,name=dailyBuyCnt" json:"dailyBuyCnt,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MascotBuyCntS2C) Reset()                    { *m = MascotBuyCntS2C{} }
func (m *MascotBuyCntS2C) String() string            { return proto.CompactTextString(m) }
func (*MascotBuyCntS2C) ProtoMessage()               {}
func (*MascotBuyCntS2C) Descriptor() ([]byte, []int) { return fileDescriptor43, []int{14} }

func (m *MascotBuyCntS2C) GetDailyBuyCnt() int32 {
	if m != nil && m.DailyBuyCnt != nil {
		return *m.DailyBuyCnt
	}
	return 0
}

func init() {
	proto.RegisterType((*MascotChapterData)(nil), "mascotChapterData")
	proto.RegisterType((*MascotBaseData)(nil), "mascotBaseData")
	proto.RegisterType((*MascotData)(nil), "mascotData")
	proto.RegisterType((*MascotGetDataC2S)(nil), "mascotGetDataC2S")
	proto.RegisterType((*MascotGetDataS2C)(nil), "mascotGetDataS2C")
	proto.RegisterType((*MascotResetCntS2C)(nil), "mascotResetCntS2C")
	proto.RegisterType((*MascotSettleDataS2C)(nil), "mascotSettleDataS2C")
	proto.RegisterType((*MascotGetSectionRewardC2S)(nil), "mascotGetSectionRewardC2S")
	proto.RegisterType((*MascotGetSectionRewardS2C)(nil), "mascotGetSectionRewardS2C")
	proto.RegisterType((*MascotGetAchieveRewardC2S)(nil), "mascotGetAchieveRewardC2S")
	proto.RegisterType((*MascotGetAchieveRewardS2C)(nil), "mascotGetAchieveRewardS2C")
	proto.RegisterType((*MascotSweepC2S)(nil), "mascotSweepC2S")
	proto.RegisterType((*MascotSweepS2C)(nil), "mascotSweepS2C")
	proto.RegisterType((*MascotBuyCntC2S)(nil), "mascotBuyCntC2S")
	proto.RegisterType((*MascotBuyCntS2C)(nil), "mascotBuyCntS2C")
	proto.RegisterEnum("MascotGetDataC2S_MsgID", MascotGetDataC2S_MsgID_name, MascotGetDataC2S_MsgID_value)
	proto.RegisterEnum("MascotGetDataS2C_MsgID", MascotGetDataS2C_MsgID_name, MascotGetDataS2C_MsgID_value)
	proto.RegisterEnum("MascotResetCntS2C_MsgID", MascotResetCntS2C_MsgID_name, MascotResetCntS2C_MsgID_value)
	proto.RegisterEnum("MascotSettleDataS2C_MsgID", MascotSettleDataS2C_MsgID_name, MascotSettleDataS2C_MsgID_value)
	proto.RegisterEnum("MascotGetSectionRewardC2S_MsgID", MascotGetSectionRewardC2S_MsgID_name, MascotGetSectionRewardC2S_MsgID_value)
	proto.RegisterEnum("MascotGetSectionRewardS2C_MsgID", MascotGetSectionRewardS2C_MsgID_name, MascotGetSectionRewardS2C_MsgID_value)
	proto.RegisterEnum("MascotGetAchieveRewardC2S_MsgID", MascotGetAchieveRewardC2S_MsgID_name, MascotGetAchieveRewardC2S_MsgID_value)
	proto.RegisterEnum("MascotGetAchieveRewardS2C_MsgID", MascotGetAchieveRewardS2C_MsgID_name, MascotGetAchieveRewardS2C_MsgID_value)
	proto.RegisterEnum("MascotSweepC2S_MsgID", MascotSweepC2S_MsgID_name, MascotSweepC2S_MsgID_value)
	proto.RegisterEnum("MascotSweepS2C_MsgID", MascotSweepS2C_MsgID_name, MascotSweepS2C_MsgID_value)
	proto.RegisterEnum("MascotBuyCntC2S_MsgID", MascotBuyCntC2S_MsgID_name, MascotBuyCntC2S_MsgID_value)
	proto.RegisterEnum("MascotBuyCntS2C_MsgID", MascotBuyCntS2C_MsgID_name, MascotBuyCntS2C_MsgID_value)
}

func init() { proto.RegisterFile("mascot.proto", fileDescriptor43) }

var fileDescriptor43 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x4d, 0x6f, 0x12, 0x41,
	0x18, 0xc7, 0x03, 0xb8, 0x40, 0x9f, 0x6d, 0xbb, 0x74, 0xd0, 0x88, 0x1e, 0xb4, 0x19, 0x7b, 0x68,
	0x2f, 0x1b, 0xc3, 0xc5, 0x8b, 0x17, 0xbb, 0x44, 0x43, 0xac, 0x89, 0x61, 0x1b, 0x3d, 0x9a, 0xe9,
	0xf2, 0x00, 0x13, 0x61, 0x77, 0xb3, 0x33, 0x5a, 0x39, 0xfa, 0x9e, 0xf8, 0x2d, 0xfc, 0xa6, 0xce,
	0xdb, 0x92, 0x80, 0x0b, 0x31, 0xa6, 0xb7, 0xe1, 0xbf, 0xf3, 0xfc, 0xe7, 0xf7, 0xbc, 0x01, 0xfb,
	0x0b, 0x26, 0x92, 0x4c, 0x86, 0x79, 0x91, 0xc9, 0xec, 0x3e, 0x24, 0x59, 0x3a, 0x71, 0x67, 0x7f,
	0xc1, 0x92, 0x22, 0x2b, 0x3f, 0x48, 0x26, 0xde, 0xdb, 0x33, 0x7d, 0x0c, 0x47, 0x36, 0x28, 0x9a,
	0xb1, 0x5c, 0x62, 0x31, 0x60, 0x92, 0x11, 0x80, 0x3a, 0x1f, 0xf7, 0x6a, 0xc7, 0xb5, 0x53, 0x8f,
	0x04, 0xd0, 0xe2, 0xe2, 0x5d, 0xce, 0x84, 0xe8, 0xd5, 0x95, 0xd0, 0xa6, 0xbf, 0x6b, 0x70, 0x68,
	0x43, 0xce, 0x99, 0x40, 0x73, 0xff, 0x2e, 0x04, 0x29, 0x7e, 0x92, 0x23, 0x9c, 0x14, 0x28, 0x66,
	0x97, 0x7c, 0x81, 0x26, 0xb8, 0x41, 0xba, 0xe0, 0x8f, 0x19, 0x9f, 0x2f, 0xcf, 0x3f, 0x2c, 0xa3,
	0x54, 0x1a, 0x03, 0x8f, 0xdc, 0x81, 0x03, 0x23, 0x3e, 0xe7, 0xd3, 0x99, 0xd4, 0x72, 0xa3, 0x94,
	0x65, 0x26, 0xd9, 0x7c, 0x25, 0xdf, 0x32, 0xb2, 0x82, 0x11, 0x98, 0x48, 0x9e, 0xa5, 0xc3, 0xb1,
	0xe8, 0x79, 0xc7, 0x0d, 0xa5, 0x3d, 0x84, 0x7d, 0x96, 0xcc, 0x38, 0x7e, 0xc4, 0x4b, 0x95, 0x89,
	0xe8, 0x35, 0x95, 0xea, 0xf7, 0xf7, 0x42, 0xfd, 0x6b, 0x28, 0x71, 0x41, 0x47, 0x00, 0x16, 0xd1,
	0xe0, 0x3d, 0x00, 0xef, 0x4a, 0xa1, 0x0a, 0x03, 0xe5, 0xf7, 0x83, 0x70, 0x03, 0xff, 0x04, 0xda,
	0x89, 0xcd, 0x5e, 0xe7, 0xa8, 0xad, 0x48, 0xf8, 0x57, 0x51, 0xe8, 0x29, 0x74, 0xac, 0xf8, 0x02,
	0x8d, 0x6d, 0xd4, 0x8f, 0xe9, 0x6d, 0xf0, 0x5e, 0x89, 0xe9, 0x70, 0x40, 0x7c, 0x68, 0xa2, 0x39,
	0x75, 0x3e, 0xe7, 0x34, 0xdd, 0xb8, 0x19, 0xf7, 0xa3, 0x1b, 0x62, 0xa8, 0x7e, 0xef, 0xac, 0xec,
	0xe1, 0x08, 0x05, 0xea, 0xd2, 0xa9, 0x07, 0xab, 0xae, 0x7e, 0xc9, 0xe9, 0xaf, 0x1a, 0x74, 0xed,
	0xdd, 0x18, 0xa5, 0x9c, 0x63, 0x89, 0xf7, 0x08, 0x5a, 0xee, 0x79, 0x07, 0x58, 0xf1, 0x3a, 0x39,
	0x83, 0xce, 0x84, 0x17, 0x42, 0xbe, 0x56, 0xc3, 0x30, 0xc2, 0x6b, 0x56, 0x8c, 0x4b, 0xd6, 0x20,
	0xd4, 0x65, 0x7f, 0xc3, 0x8a, 0x65, 0xa4, 0x66, 0x8e, 0x4f, 0xc9, 0x01, 0x78, 0x5c, 0xbc, 0xe5,
	0xa9, 0xe9, 0x6d, 0xbb, 0x0a, 0xe6, 0x6b, 0x4e, 0x07, 0x70, 0x6f, 0x55, 0xa7, 0xd8, 0xf6, 0xd8,
	0xda, 0xaa, 0xd2, 0x92, 0x23, 0xd8, 0x5b, 0xf5, 0xdd, 0x8e, 0x62, 0x95, 0xcb, 0xb7, 0x1d, 0x2e,
	0x3a, 0xaf, 0xff, 0x72, 0x79, 0x66, 0x67, 0x6b, 0x8d, 0xc5, 0xcd, 0x9b, 0x72, 0xa9, 0x6f, 0x73,
	0xf9, 0xbe, 0xc3, 0xc5, 0xb1, 0xfc, 0x9b, 0xcb, 0xcb, 0x72, 0xc1, 0xe2, 0x6b, 0xc4, 0xdc, 0x01,
	0xb8, 0xf6, 0x94, 0x69, 0x90, 0x0e, 0xb4, 0x85, 0xf9, 0x5c, 0xee, 0x55, 0x95, 0xd9, 0x8f, 0x9c,
	0x5e, 0xac, 0x99, 0x69, 0x8e, 0x13, 0x68, 0x15, 0xae, 0x7b, 0x35, 0xd3, 0xbd, 0xee, 0x46, 0xf7,
	0x2e, 0xb8, 0x90, 0x5b, 0xdc, 0x9e, 0x40, 0xe0, 0x06, 0xd7, 0x6c, 0xb4, 0x66, 0x3b, 0x84, 0xe6,
	0x95, 0x5d, 0xef, 0xad, 0xf5, 0xfd, 0x99, 0xd3, 0xa7, 0xeb, 0x81, 0x9a, 0x63, 0xe3, 0xcf, 0x61,
	0x47, 0xf4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xdb, 0x10, 0xa7, 0xd9, 0x04, 0x00, 0x00,
}