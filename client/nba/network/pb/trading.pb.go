// Code generated by protoc-gen-go.
// source: trading.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CompoundContractC2S_MsgID int32

const (
	CompoundContractC2S_eMsgID CompoundContractC2S_MsgID = 15105
)

var CompoundContractC2S_MsgID_name = map[int32]string{
	15105: "eMsgID",
}
var CompoundContractC2S_MsgID_value = map[string]int32{
	"eMsgID": 15105,
}

func (x CompoundContractC2S_MsgID) Enum() *CompoundContractC2S_MsgID {
	p := new(CompoundContractC2S_MsgID)
	*p = x
	return p
}
func (x CompoundContractC2S_MsgID) String() string {
	return proto.EnumName(CompoundContractC2S_MsgID_name, int32(x))
}
func (x *CompoundContractC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CompoundContractC2S_MsgID_value, data, "CompoundContractC2S_MsgID")
	if err != nil {
		return err
	}
	*x = CompoundContractC2S_MsgID(value)
	return nil
}
func (CompoundContractC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor80, []int{3, 0}
}

type CompoundContractS2C_MsgID int32

const (
	CompoundContractS2C_eMsgID CompoundContractS2C_MsgID = 15105
)

var CompoundContractS2C_MsgID_name = map[int32]string{
	15105: "eMsgID",
}
var CompoundContractS2C_MsgID_value = map[string]int32{
	"eMsgID": 15105,
}

func (x CompoundContractS2C_MsgID) Enum() *CompoundContractS2C_MsgID {
	p := new(CompoundContractS2C_MsgID)
	*p = x
	return p
}
func (x CompoundContractS2C_MsgID) String() string {
	return proto.EnumName(CompoundContractS2C_MsgID_name, int32(x))
}
func (x *CompoundContractS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CompoundContractS2C_MsgID_value, data, "CompoundContractS2C_MsgID")
	if err != nil {
		return err
	}
	*x = CompoundContractS2C_MsgID(value)
	return nil
}
func (CompoundContractS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor80, []int{4, 0}
}

type SellContractC2S_MsgID int32

const (
	SellContractC2S_eMsgID SellContractC2S_MsgID = 15107
)

var SellContractC2S_MsgID_name = map[int32]string{
	15107: "eMsgID",
}
var SellContractC2S_MsgID_value = map[string]int32{
	"eMsgID": 15107,
}

func (x SellContractC2S_MsgID) Enum() *SellContractC2S_MsgID {
	p := new(SellContractC2S_MsgID)
	*p = x
	return p
}
func (x SellContractC2S_MsgID) String() string {
	return proto.EnumName(SellContractC2S_MsgID_name, int32(x))
}
func (x *SellContractC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SellContractC2S_MsgID_value, data, "SellContractC2S_MsgID")
	if err != nil {
		return err
	}
	*x = SellContractC2S_MsgID(value)
	return nil
}
func (SellContractC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{5, 0} }

type SellContractS2C_MsgID int32

const (
	SellContractS2C_eMsgID SellContractS2C_MsgID = 15107
)

var SellContractS2C_MsgID_name = map[int32]string{
	15107: "eMsgID",
}
var SellContractS2C_MsgID_value = map[string]int32{
	"eMsgID": 15107,
}

func (x SellContractS2C_MsgID) Enum() *SellContractS2C_MsgID {
	p := new(SellContractS2C_MsgID)
	*p = x
	return p
}
func (x SellContractS2C_MsgID) String() string {
	return proto.EnumName(SellContractS2C_MsgID_name, int32(x))
}
func (x *SellContractS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SellContractS2C_MsgID_value, data, "SellContractS2C_MsgID")
	if err != nil {
		return err
	}
	*x = SellContractS2C_MsgID(value)
	return nil
}
func (SellContractS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{6, 0} }

type CancelContractC2S_MsgID int32

const (
	CancelContractC2S_eMsgID CancelContractC2S_MsgID = 15108
)

var CancelContractC2S_MsgID_name = map[int32]string{
	15108: "eMsgID",
}
var CancelContractC2S_MsgID_value = map[string]int32{
	"eMsgID": 15108,
}

func (x CancelContractC2S_MsgID) Enum() *CancelContractC2S_MsgID {
	p := new(CancelContractC2S_MsgID)
	*p = x
	return p
}
func (x CancelContractC2S_MsgID) String() string {
	return proto.EnumName(CancelContractC2S_MsgID_name, int32(x))
}
func (x *CancelContractC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CancelContractC2S_MsgID_value, data, "CancelContractC2S_MsgID")
	if err != nil {
		return err
	}
	*x = CancelContractC2S_MsgID(value)
	return nil
}
func (CancelContractC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{7, 0} }

type CancelContractS2C_MsgID int32

const (
	CancelContractS2C_eMsgID CancelContractS2C_MsgID = 15108
)

var CancelContractS2C_MsgID_name = map[int32]string{
	15108: "eMsgID",
}
var CancelContractS2C_MsgID_value = map[string]int32{
	"eMsgID": 15108,
}

func (x CancelContractS2C_MsgID) Enum() *CancelContractS2C_MsgID {
	p := new(CancelContractS2C_MsgID)
	*p = x
	return p
}
func (x CancelContractS2C_MsgID) String() string {
	return proto.EnumName(CancelContractS2C_MsgID_name, int32(x))
}
func (x *CancelContractS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CancelContractS2C_MsgID_value, data, "CancelContractS2C_MsgID")
	if err != nil {
		return err
	}
	*x = CancelContractS2C_MsgID(value)
	return nil
}
func (CancelContractS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{8, 0} }

type BuyContractC2S_MsgID int32

const (
	BuyContractC2S_eMsgID BuyContractC2S_MsgID = 15109
)

var BuyContractC2S_MsgID_name = map[int32]string{
	15109: "eMsgID",
}
var BuyContractC2S_MsgID_value = map[string]int32{
	"eMsgID": 15109,
}

func (x BuyContractC2S_MsgID) Enum() *BuyContractC2S_MsgID {
	p := new(BuyContractC2S_MsgID)
	*p = x
	return p
}
func (x BuyContractC2S_MsgID) String() string {
	return proto.EnumName(BuyContractC2S_MsgID_name, int32(x))
}
func (x *BuyContractC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BuyContractC2S_MsgID_value, data, "BuyContractC2S_MsgID")
	if err != nil {
		return err
	}
	*x = BuyContractC2S_MsgID(value)
	return nil
}
func (BuyContractC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{9, 0} }

type BuyContractS2C_MsgID int32

const (
	BuyContractS2C_eMsgID BuyContractS2C_MsgID = 15109
)

var BuyContractS2C_MsgID_name = map[int32]string{
	15109: "eMsgID",
}
var BuyContractS2C_MsgID_value = map[string]int32{
	"eMsgID": 15109,
}

func (x BuyContractS2C_MsgID) Enum() *BuyContractS2C_MsgID {
	p := new(BuyContractS2C_MsgID)
	*p = x
	return p
}
func (x BuyContractS2C_MsgID) String() string {
	return proto.EnumName(BuyContractS2C_MsgID_name, int32(x))
}
func (x *BuyContractS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BuyContractS2C_MsgID_value, data, "BuyContractS2C_MsgID")
	if err != nil {
		return err
	}
	*x = BuyContractS2C_MsgID(value)
	return nil
}
func (BuyContractS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{10, 0} }

type FindContractC2S_MsgID int32

const (
	FindContractC2S_eMsgID FindContractC2S_MsgID = 15110
)

var FindContractC2S_MsgID_name = map[int32]string{
	15110: "eMsgID",
}
var FindContractC2S_MsgID_value = map[string]int32{
	"eMsgID": 15110,
}

func (x FindContractC2S_MsgID) Enum() *FindContractC2S_MsgID {
	p := new(FindContractC2S_MsgID)
	*p = x
	return p
}
func (x FindContractC2S_MsgID) String() string {
	return proto.EnumName(FindContractC2S_MsgID_name, int32(x))
}
func (x *FindContractC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FindContractC2S_MsgID_value, data, "FindContractC2S_MsgID")
	if err != nil {
		return err
	}
	*x = FindContractC2S_MsgID(value)
	return nil
}
func (FindContractC2S_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{11, 0} }

type FindContractS2C_MsgID int32

const (
	FindContractS2C_eMsgID FindContractS2C_MsgID = 15110
)

var FindContractS2C_MsgID_name = map[int32]string{
	15110: "eMsgID",
}
var FindContractS2C_MsgID_value = map[string]int32{
	"eMsgID": 15110,
}

func (x FindContractS2C_MsgID) Enum() *FindContractS2C_MsgID {
	p := new(FindContractS2C_MsgID)
	*p = x
	return p
}
func (x FindContractS2C_MsgID) String() string {
	return proto.EnumName(FindContractS2C_MsgID_name, int32(x))
}
func (x *FindContractS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FindContractS2C_MsgID_value, data, "FindContractS2C_MsgID")
	if err != nil {
		return err
	}
	*x = FindContractS2C_MsgID(value)
	return nil
}
func (FindContractS2C_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor80, []int{12, 0} }

type GetGamerTradeInfoC2S_MsgID int32

const (
	GetGamerTradeInfoC2S_eMsgID GetGamerTradeInfoC2S_MsgID = 15111
)

var GetGamerTradeInfoC2S_MsgID_name = map[int32]string{
	15111: "eMsgID",
}
var GetGamerTradeInfoC2S_MsgID_value = map[string]int32{
	"eMsgID": 15111,
}

func (x GetGamerTradeInfoC2S_MsgID) Enum() *GetGamerTradeInfoC2S_MsgID {
	p := new(GetGamerTradeInfoC2S_MsgID)
	*p = x
	return p
}
func (x GetGamerTradeInfoC2S_MsgID) String() string {
	return proto.EnumName(GetGamerTradeInfoC2S_MsgID_name, int32(x))
}
func (x *GetGamerTradeInfoC2S_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetGamerTradeInfoC2S_MsgID_value, data, "GetGamerTradeInfoC2S_MsgID")
	if err != nil {
		return err
	}
	*x = GetGamerTradeInfoC2S_MsgID(value)
	return nil
}
func (GetGamerTradeInfoC2S_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor80, []int{13, 0}
}

type GetGamerTradeInfoS2C_MsgID int32

const (
	GetGamerTradeInfoS2C_eMsgID GetGamerTradeInfoS2C_MsgID = 15111
)

var GetGamerTradeInfoS2C_MsgID_name = map[int32]string{
	15111: "eMsgID",
}
var GetGamerTradeInfoS2C_MsgID_value = map[string]int32{
	"eMsgID": 15111,
}

func (x GetGamerTradeInfoS2C_MsgID) Enum() *GetGamerTradeInfoS2C_MsgID {
	p := new(GetGamerTradeInfoS2C_MsgID)
	*p = x
	return p
}
func (x GetGamerTradeInfoS2C_MsgID) String() string {
	return proto.EnumName(GetGamerTradeInfoS2C_MsgID_name, int32(x))
}
func (x *GetGamerTradeInfoS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetGamerTradeInfoS2C_MsgID_value, data, "GetGamerTradeInfoS2C_MsgID")
	if err != nil {
		return err
	}
	*x = GetGamerTradeInfoS2C_MsgID(value)
	return nil
}
func (GetGamerTradeInfoS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor80, []int{14, 0}
}

type RemoveTradeInfoS2C_MsgID int32

const (
	RemoveTradeInfoS2C_eMsgID RemoveTradeInfoS2C_MsgID = 15112
)

var RemoveTradeInfoS2C_MsgID_name = map[int32]string{
	15112: "eMsgID",
}
var RemoveTradeInfoS2C_MsgID_value = map[string]int32{
	"eMsgID": 15112,
}

func (x RemoveTradeInfoS2C_MsgID) Enum() *RemoveTradeInfoS2C_MsgID {
	p := new(RemoveTradeInfoS2C_MsgID)
	*p = x
	return p
}
func (x RemoveTradeInfoS2C_MsgID) String() string {
	return proto.EnumName(RemoveTradeInfoS2C_MsgID_name, int32(x))
}
func (x *RemoveTradeInfoS2C_MsgID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RemoveTradeInfoS2C_MsgID_value, data, "RemoveTradeInfoS2C_MsgID")
	if err != nil {
		return err
	}
	*x = RemoveTradeInfoS2C_MsgID(value)
	return nil
}
func (RemoveTradeInfoS2C_MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor80, []int{15, 0}
}

// 交易基本信息entry
type TradeEntry struct {
	// id
	TradeId *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	// 卖家gid
	Seller *int32 `protobuf:"varint,2,opt,name=seller" json:"seller,omitempty"`
	// 下架时间戳
	ExpireTime *int64 `protobuf:"varint,3,opt,name=expireTime" json:"expireTime,omitempty"`
	// 交易类型
	SellType *int32 `protobuf:"varint,4,opt,name=sellType" json:"sellType,omitempty"`
	// 交易商品
	Goods *ItemVaryConfig `protobuf:"bytes,5,opt,name=goods" json:"goods,omitempty"`
	// 交易价格
	Price *ItemVaryConfig `protobuf:"bytes,6,opt,name=price" json:"price,omitempty"`
	// 球员id
	PlayerId         *int32 `protobuf:"varint,7,opt,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TradeEntry) Reset()                    { *m = TradeEntry{} }
func (m *TradeEntry) String() string            { return proto.CompactTextString(m) }
func (*TradeEntry) ProtoMessage()               {}
func (*TradeEntry) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{0} }

func (m *TradeEntry) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

func (m *TradeEntry) GetSeller() int32 {
	if m != nil && m.Seller != nil {
		return *m.Seller
	}
	return 0
}

func (m *TradeEntry) GetExpireTime() int64 {
	if m != nil && m.ExpireTime != nil {
		return *m.ExpireTime
	}
	return 0
}

func (m *TradeEntry) GetSellType() int32 {
	if m != nil && m.SellType != nil {
		return *m.SellType
	}
	return 0
}

func (m *TradeEntry) GetGoods() *ItemVaryConfig {
	if m != nil {
		return m.Goods
	}
	return nil
}

func (m *TradeEntry) GetPrice() *ItemVaryConfig {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *TradeEntry) GetPlayerId() int32 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

// 玩家交易相关次数限制
type GamerTradeNumData struct {
	// 下次刷新次数的时间戳
	NextRefTime *int64 `protobuf:"varint,1,opt,name=nextRefTime" json:"nextRefTime,omitempty"`
	// 每周已卖的次数
	WeekSellNum *int32 `protobuf:"varint,2,opt,name=weekSellNum" json:"weekSellNum,omitempty"`
	// 每周已买的次数
	WeekBuyNum *int32 `protobuf:"varint,3,opt,name=weekBuyNum" json:"weekBuyNum,omitempty"`
	// 每周已交换的次数
	WeekSwapNum      *int32 `protobuf:"varint,4,opt,name=weekSwapNum" json:"weekSwapNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GamerTradeNumData) Reset()                    { *m = GamerTradeNumData{} }
func (m *GamerTradeNumData) String() string            { return proto.CompactTextString(m) }
func (*GamerTradeNumData) ProtoMessage()               {}
func (*GamerTradeNumData) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{1} }

func (m *GamerTradeNumData) GetNextRefTime() int64 {
	if m != nil && m.NextRefTime != nil {
		return *m.NextRefTime
	}
	return 0
}

func (m *GamerTradeNumData) GetWeekSellNum() int32 {
	if m != nil && m.WeekSellNum != nil {
		return *m.WeekSellNum
	}
	return 0
}

func (m *GamerTradeNumData) GetWeekBuyNum() int32 {
	if m != nil && m.WeekBuyNum != nil {
		return *m.WeekBuyNum
	}
	return 0
}

func (m *GamerTradeNumData) GetWeekSwapNum() int32 {
	if m != nil && m.WeekSwapNum != nil {
		return *m.WeekSwapNum
	}
	return 0
}

// 玩家寄卖的合同id
type GamerTradeSellData struct {
	TradeDetails     []*GamerTradeSellData_TradeDetail `protobuf:"bytes,1,rep,name=tradeDetails" json:"tradeDetails,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *GamerTradeSellData) Reset()                    { *m = GamerTradeSellData{} }
func (m *GamerTradeSellData) String() string            { return proto.CompactTextString(m) }
func (*GamerTradeSellData) ProtoMessage()               {}
func (*GamerTradeSellData) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{2} }

func (m *GamerTradeSellData) GetTradeDetails() []*GamerTradeSellData_TradeDetail {
	if m != nil {
		return m.TradeDetails
	}
	return nil
}

type GamerTradeSellData_TradeDetail struct {
	// 交易id
	TradeId *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	// 交易类型
	TradeType        *int32 `protobuf:"varint,2,opt,name=tradeType" json:"tradeType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GamerTradeSellData_TradeDetail) Reset()         { *m = GamerTradeSellData_TradeDetail{} }
func (m *GamerTradeSellData_TradeDetail) String() string { return proto.CompactTextString(m) }
func (*GamerTradeSellData_TradeDetail) ProtoMessage()    {}
func (*GamerTradeSellData_TradeDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor80, []int{2, 0}
}

func (m *GamerTradeSellData_TradeDetail) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

func (m *GamerTradeSellData_TradeDetail) GetTradeType() int32 {
	if m != nil && m.TradeType != nil {
		return *m.TradeType
	}
	return 0
}

// 1.球员碎片合成球员合同c2s
type CompoundContractC2S struct {
	// 将合成的球员合同道具id
	ItemIds          []int32 `protobuf:"varint,1,rep,name=itemIds" json:"itemIds,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CompoundContractC2S) Reset()                    { *m = CompoundContractC2S{} }
func (m *CompoundContractC2S) String() string            { return proto.CompactTextString(m) }
func (*CompoundContractC2S) ProtoMessage()               {}
func (*CompoundContractC2S) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{3} }

func (m *CompoundContractC2S) GetItemIds() []int32 {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

// 球员碎片合成球员合同s2c
type CompoundContractS2C struct {
	// 合成的合同道具
	Contracts        []*ItemVaryConfig `protobuf:"bytes,1,rep,name=contracts" json:"contracts,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *CompoundContractS2C) Reset()                    { *m = CompoundContractS2C{} }
func (m *CompoundContractS2C) String() string            { return proto.CompactTextString(m) }
func (*CompoundContractS2C) ProtoMessage()               {}
func (*CompoundContractS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{4} }

func (m *CompoundContractS2C) GetContracts() []*ItemVaryConfig {
	if m != nil {
		return m.Contracts
	}
	return nil
}

// 3.寄买合同c2s
type SellContractC2S struct {
	// 寄卖的球员合同道具id
	ItemId *int32 `protobuf:"varint,1,opt,name=item_id" json:"item_id,omitempty"`
	// 类型(1:代币交易. 2:球员交换)
	Type *int32 `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	// 价格
	Price            *ItemVaryConfig `protobuf:"bytes,3,opt,name=price" json:"price,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SellContractC2S) Reset()                    { *m = SellContractC2S{} }
func (m *SellContractC2S) String() string            { return proto.CompactTextString(m) }
func (*SellContractC2S) ProtoMessage()               {}
func (*SellContractC2S) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{5} }

func (m *SellContractC2S) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *SellContractC2S) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *SellContractC2S) GetPrice() *ItemVaryConfig {
	if m != nil {
		return m.Price
	}
	return nil
}

// 寄卖合同s2c
type SellContractS2C struct {
	// 寄卖的合同交易
	Trade            *TradeEntry `protobuf:"bytes,1,opt,name=trade" json:"trade,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SellContractS2C) Reset()                    { *m = SellContractS2C{} }
func (m *SellContractS2C) String() string            { return proto.CompactTextString(m) }
func (*SellContractS2C) ProtoMessage()               {}
func (*SellContractS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{6} }

func (m *SellContractS2C) GetTrade() *TradeEntry {
	if m != nil {
		return m.Trade
	}
	return nil
}

// 4.取消寄卖c2s
type CancelContractC2S struct {
	// 交易id
	TradeId          *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CancelContractC2S) Reset()                    { *m = CancelContractC2S{} }
func (m *CancelContractC2S) String() string            { return proto.CompactTextString(m) }
func (*CancelContractC2S) ProtoMessage()               {}
func (*CancelContractC2S) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{7} }

func (m *CancelContractC2S) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

// 取消寄卖s2c
type CancelContractS2C struct {
	// 交易id
	TradeId          *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CancelContractS2C) Reset()                    { *m = CancelContractS2C{} }
func (m *CancelContractS2C) String() string            { return proto.CompactTextString(m) }
func (*CancelContractS2C) ProtoMessage()               {}
func (*CancelContractS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{8} }

func (m *CancelContractS2C) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

// 5.购买合同c2s
type BuyContractC2S struct {
	// 交易id
	TradeId          *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuyContractC2S) Reset()                    { *m = BuyContractC2S{} }
func (m *BuyContractC2S) String() string            { return proto.CompactTextString(m) }
func (*BuyContractC2S) ProtoMessage()               {}
func (*BuyContractC2S) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{9} }

func (m *BuyContractC2S) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

// 购买合同s2c
type BuyContractS2C struct {
	// 交易id
	TradeId          *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BuyContractS2C) Reset()                    { *m = BuyContractS2C{} }
func (m *BuyContractS2C) String() string            { return proto.CompactTextString(m) }
func (*BuyContractS2C) ProtoMessage()               {}
func (*BuyContractS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{10} }

func (m *BuyContractS2C) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

// 6.查找合同c2s
type FindContractC2S struct {
	// 交易类型（1：代币 2：球员）
	SellType *int32 `protobuf:"varint,1,opt,name=sellType" json:"sellType,omitempty"`
	// 按位置查找(0:不含位置 1-5位置)
	Place *int32 `protobuf:"varint,4,opt,name=place" json:"place,omitempty"`
	// 按球员名称查找
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// 按品质查找
	Quality          *int32 `protobuf:"varint,5,opt,name=quality" json:"quality,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FindContractC2S) Reset()                    { *m = FindContractC2S{} }
func (m *FindContractC2S) String() string            { return proto.CompactTextString(m) }
func (*FindContractC2S) ProtoMessage()               {}
func (*FindContractC2S) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{11} }

func (m *FindContractC2S) GetSellType() int32 {
	if m != nil && m.SellType != nil {
		return *m.SellType
	}
	return 0
}

func (m *FindContractC2S) GetPlace() int32 {
	if m != nil && m.Place != nil {
		return *m.Place
	}
	return 0
}

func (m *FindContractC2S) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FindContractC2S) GetQuality() int32 {
	if m != nil && m.Quality != nil {
		return *m.Quality
	}
	return 0
}

// 查找合同s2c
// @Gamer
type FindContractS2C struct {
	// 满足条件的交易合同列表
	Entries []*TradeEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// 请求频率过高,返回冷却剩余时间（毫秒）
	Cd               *int64 `protobuf:"varint,2,opt,name=cd" json:"cd,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FindContractS2C) Reset()                    { *m = FindContractS2C{} }
func (m *FindContractS2C) String() string            { return proto.CompactTextString(m) }
func (*FindContractS2C) ProtoMessage()               {}
func (*FindContractS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{12} }

func (m *FindContractS2C) GetEntries() []*TradeEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *FindContractS2C) GetCd() int64 {
	if m != nil && m.Cd != nil {
		return *m.Cd
	}
	return 0
}

// 7.获取玩家交易相关数据c2s
type GetGamerTradeInfoC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetGamerTradeInfoC2S) Reset()                    { *m = GetGamerTradeInfoC2S{} }
func (m *GetGamerTradeInfoC2S) String() string            { return proto.CompactTextString(m) }
func (*GetGamerTradeInfoC2S) ProtoMessage()               {}
func (*GetGamerTradeInfoC2S) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{13} }

// 获取玩家交易相关数据s2c
type GetGamerTradeInfoS2C struct {
	// 次数限制
	TradeNum *GamerTradeNumData `protobuf:"bytes,1,opt,name=tradeNum" json:"tradeNum,omitempty"`
	// 自己寄卖的合同列表
	Entries          []*TradeEntry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GetGamerTradeInfoS2C) Reset()                    { *m = GetGamerTradeInfoS2C{} }
func (m *GetGamerTradeInfoS2C) String() string            { return proto.CompactTextString(m) }
func (*GetGamerTradeInfoS2C) ProtoMessage()               {}
func (*GetGamerTradeInfoS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{14} }

func (m *GetGamerTradeInfoS2C) GetTradeNum() *GamerTradeNumData {
	if m != nil {
		return m.TradeNum
	}
	return nil
}

func (m *GetGamerTradeInfoS2C) GetEntries() []*TradeEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// 通知玩家移除交易
type RemoveTradeInfoS2C struct {
	TradeId          *int64 `protobuf:"varint,1,opt,name=tradeId" json:"tradeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RemoveTradeInfoS2C) Reset()                    { *m = RemoveTradeInfoS2C{} }
func (m *RemoveTradeInfoS2C) String() string            { return proto.CompactTextString(m) }
func (*RemoveTradeInfoS2C) ProtoMessage()               {}
func (*RemoveTradeInfoS2C) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{15} }

func (m *RemoveTradeInfoS2C) GetTradeId() int64 {
	if m != nil && m.TradeId != nil {
		return *m.TradeId
	}
	return 0
}

// 交易结束、通知卖家
type TradingFinishData struct {
	// true:交易成功 false:交易超时退还
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// 交易
	Entry            *TradeEntry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TradingFinishData) Reset()                    { *m = TradingFinishData{} }
func (m *TradingFinishData) String() string            { return proto.CompactTextString(m) }
func (*TradingFinishData) ProtoMessage()               {}
func (*TradingFinishData) Descriptor() ([]byte, []int) { return fileDescriptor80, []int{16} }

func (m *TradingFinishData) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *TradingFinishData) GetEntry() *TradeEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*TradeEntry)(nil), "TradeEntry")
	proto.RegisterType((*GamerTradeNumData)(nil), "GamerTradeNumData")
	proto.RegisterType((*GamerTradeSellData)(nil), "GamerTradeSellData")
	proto.RegisterType((*GamerTradeSellData_TradeDetail)(nil), "GamerTradeSellData.TradeDetail")
	proto.RegisterType((*CompoundContractC2S)(nil), "CompoundContractC2S")
	proto.RegisterType((*CompoundContractS2C)(nil), "CompoundContractS2C")
	proto.RegisterType((*SellContractC2S)(nil), "SellContractC2S")
	proto.RegisterType((*SellContractS2C)(nil), "SellContractS2C")
	proto.RegisterType((*CancelContractC2S)(nil), "CancelContractC2S")
	proto.RegisterType((*CancelContractS2C)(nil), "CancelContractS2C")
	proto.RegisterType((*BuyContractC2S)(nil), "BuyContractC2S")
	proto.RegisterType((*BuyContractS2C)(nil), "BuyContractS2C")
	proto.RegisterType((*FindContractC2S)(nil), "FindContractC2S")
	proto.RegisterType((*FindContractS2C)(nil), "FindContractS2C")
	proto.RegisterType((*GetGamerTradeInfoC2S)(nil), "GetGamerTradeInfoC2S")
	proto.RegisterType((*GetGamerTradeInfoS2C)(nil), "GetGamerTradeInfoS2C")
	proto.RegisterType((*RemoveTradeInfoS2C)(nil), "RemoveTradeInfoS2C")
	proto.RegisterType((*TradingFinishData)(nil), "TradingFinishData")
	proto.RegisterEnum("CompoundContractC2S_MsgID", CompoundContractC2S_MsgID_name, CompoundContractC2S_MsgID_value)
	proto.RegisterEnum("CompoundContractS2C_MsgID", CompoundContractS2C_MsgID_name, CompoundContractS2C_MsgID_value)
	proto.RegisterEnum("SellContractC2S_MsgID", SellContractC2S_MsgID_name, SellContractC2S_MsgID_value)
	proto.RegisterEnum("SellContractS2C_MsgID", SellContractS2C_MsgID_name, SellContractS2C_MsgID_value)
	proto.RegisterEnum("CancelContractC2S_MsgID", CancelContractC2S_MsgID_name, CancelContractC2S_MsgID_value)
	proto.RegisterEnum("CancelContractS2C_MsgID", CancelContractS2C_MsgID_name, CancelContractS2C_MsgID_value)
	proto.RegisterEnum("BuyContractC2S_MsgID", BuyContractC2S_MsgID_name, BuyContractC2S_MsgID_value)
	proto.RegisterEnum("BuyContractS2C_MsgID", BuyContractS2C_MsgID_name, BuyContractS2C_MsgID_value)
	proto.RegisterEnum("FindContractC2S_MsgID", FindContractC2S_MsgID_name, FindContractC2S_MsgID_value)
	proto.RegisterEnum("FindContractS2C_MsgID", FindContractS2C_MsgID_name, FindContractS2C_MsgID_value)
	proto.RegisterEnum("GetGamerTradeInfoC2S_MsgID", GetGamerTradeInfoC2S_MsgID_name, GetGamerTradeInfoC2S_MsgID_value)
	proto.RegisterEnum("GetGamerTradeInfoS2C_MsgID", GetGamerTradeInfoS2C_MsgID_name, GetGamerTradeInfoS2C_MsgID_value)
	proto.RegisterEnum("RemoveTradeInfoS2C_MsgID", RemoveTradeInfoS2C_MsgID_name, RemoveTradeInfoS2C_MsgID_value)
}

func init() { proto.RegisterFile("trading.proto", fileDescriptor80) }

var fileDescriptor80 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xc7, 0xd5, 0xba, 0xce, 0x63, 0xdc, 0x36, 0x8d, 0xdb, 0x83, 0x15, 0x21, 0x40, 0x16, 0x07,
	0x0e, 0xc8, 0x12, 0x91, 0x10, 0x12, 0x70, 0x40, 0x38, 0x50, 0xe5, 0x00, 0x88, 0x26, 0xe2, 0x8a,
	0x56, 0xf6, 0x24, 0xac, 0xf0, 0x0b, 0x7b, 0xe3, 0xd6, 0x17, 0x24, 0xc4, 0xf3, 0x8b, 0xf0, 0x3d,
	0x99, 0x5d, 0x27, 0x38, 0x4e, 0x1c, 0x29, 0xdc, 0xbc, 0xb3, 0x33, 0xf3, 0xff, 0xed, 0x3c, 0x0c,
	0x27, 0x22, 0x65, 0x3e, 0x8f, 0xe6, 0x4e, 0x92, 0xc6, 0x22, 0x1e, 0x80, 0x17, 0x47, 0xb3, 0xf2,
	0xdb, 0xfe, 0x73, 0x00, 0x30, 0xa5, 0x5b, 0x7c, 0x19, 0x89, 0xb4, 0x30, 0x7b, 0xd0, 0x96, 0xbe,
	0x38, 0xf6, 0xad, 0x83, 0xbb, 0x07, 0xf7, 0x35, 0xf3, 0x14, 0x5a, 0x19, 0x06, 0x01, 0xa6, 0xd6,
	0x21, 0x9d, 0x75, 0xd3, 0x04, 0xc0, 0x9b, 0x84, 0xa7, 0x38, 0xe5, 0x21, 0x5a, 0x9a, 0xf2, 0x39,
	0x83, 0x8e, 0xf4, 0x99, 0x16, 0x09, 0x5a, 0x47, 0xca, 0xeb, 0x36, 0xe8, 0xf3, 0x38, 0xf6, 0x33,
	0x4b, 0xa7, 0xa3, 0x31, 0xec, 0x39, 0x63, 0x81, 0xe1, 0x7b, 0x96, 0x16, 0x2e, 0x29, 0xf3, 0xb9,
	0xbc, 0x4f, 0x52, 0xee, 0xa1, 0xd5, 0x6a, 0xbe, 0xa7, 0x8c, 0x49, 0xc0, 0x0a, 0x4c, 0x89, 0xa3,
	0x2d, 0x33, 0xda, 0x33, 0xe8, 0x5f, 0xb2, 0x10, 0x53, 0xc5, 0xfa, 0x66, 0x11, 0x8e, 0x98, 0x60,
	0xe6, 0x39, 0x18, 0x11, 0xde, 0x88, 0x2b, 0x9c, 0x29, 0x9a, 0x92, 0x98, 0x8c, 0xd7, 0x88, 0x9f,
	0x26, 0x44, 0x44, 0x7e, 0x15, 0xb6, 0x34, 0xbe, 0x58, 0x14, 0xd2, 0xa6, 0x29, 0xdb, 0xca, 0xf1,
	0x9a, 0x25, 0xd2, 0xa8, 0xc8, 0xed, 0x2f, 0x60, 0x56, 0x3a, 0x32, 0x87, 0x12, 0x7a, 0x04, 0xc7,
	0xaa, 0x2c, 0x23, 0x14, 0x8c, 0x07, 0x19, 0x29, 0x69, 0x84, 0x7d, 0xc7, 0xd9, 0x76, 0x75, 0xa6,
	0x95, 0xdf, 0xe0, 0x21, 0x18, 0x6b, 0xc7, 0xed, 0xe2, 0xf6, 0xa1, 0xab, 0x0c, 0xaa, 0x72, 0x0a,
	0xd4, 0x7e, 0x06, 0xe7, 0x6e, 0x1c, 0x26, 0xf1, 0x22, 0xf2, 0xa9, 0x16, 0x74, 0xeb, 0x09, 0x77,
	0x38, 0x91, 0xa1, 0x9c, 0x4a, 0x34, 0xf6, 0x4b, 0x6d, 0xdd, 0xbe, 0x00, 0xfd, 0x75, 0x36, 0x1f,
	0x8f, 0x4c, 0x03, 0x5a, 0xa8, 0xbe, 0xce, 0xbe, 0xe6, 0xf6, 0xdb, 0xed, 0xe8, 0xc9, 0xd0, 0x35,
	0x6d, 0xe8, 0x7a, 0xcb, 0xe3, 0x8a, 0x7d, 0xb3, 0xe4, 0x3b, 0x12, 0xce, 0xa0, 0x27, 0x5f, 0xd6,
	0x80, 0xf2, 0x81, 0x97, 0xaf, 0xd0, 0xcd, 0x63, 0x38, 0x12, 0xff, 0x1e, 0x50, 0xb5, 0x56, 0x6b,
	0x6c, 0x6d, 0x93, 0xce, 0xb7, 0xdc, 0x76, 0xeb, 0x3a, 0x12, 0x7a, 0x00, 0xba, 0x2a, 0x8e, 0x52,
	0x31, 0x86, 0x86, 0x53, 0x8d, 0xe9, 0x8e, 0x24, 0x4f, 0xa0, 0xef, 0xb2, 0xc8, 0xc3, 0x4d, 0xdc,
	0x5a, 0xd1, 0x9b, 0x62, 0xbf, 0x37, 0xc4, 0x4a, 0x84, 0x3d, 0x63, 0x1f, 0xc3, 0x29, 0x0d, 0xd6,
	0xff, 0x8a, 0xfe, 0xd8, 0x0c, 0xdc, 0x57, 0x91, 0x02, 0xe7, 0xd0, 0x7b, 0xc5, 0xeb, 0x13, 0xb2,
	0xbe, 0x84, 0x65, 0x5f, 0x4e, 0xa8, 0x13, 0x01, 0xf3, 0x56, 0x3b, 0x49, 0x6d, 0x8a, 0x68, 0x5c,
	0x55, 0x9b, 0xba, 0x52, 0xe8, 0xf3, 0x82, 0x05, 0x5c, 0x14, 0x6a, 0x47, 0x1b, 0x07, 0xea, 0x67,
	0x6e, 0xbf, 0xab, 0x0b, 0x49, 0xc4, 0x5b, 0xd0, 0x46, 0x3a, 0x71, 0x5c, 0x8d, 0xd2, 0x7a, 0x67,
	0x68, 0xcf, 0x0e, 0x3d, 0x5f, 0x69, 0x68, 0x3b, 0x52, 0x3e, 0x80, 0x8b, 0x4b, 0x14, 0xd5, 0xe6,
	0x8c, 0xa3, 0x59, 0x4c, 0x0f, 0x68, 0xf2, 0xfe, 0x95, 0xdb, 0x49, 0x83, 0xb7, 0xa4, 0xb8, 0x07,
	0x1d, 0xb1, 0xfc, 0x15, 0x2c, 0x07, 0xc4, 0x74, 0xb6, 0x7f, 0x10, 0x6b, 0xac, 0x87, 0x5b, 0xac,
	0x3b, 0x14, 0x9f, 0x82, 0x79, 0x85, 0x61, 0x9c, 0x63, 0x4d, 0x6f, 0x9f, 0xc6, 0xfc, 0xce, 0xed,
	0xe7, 0xd0, 0x9f, 0x96, 0xff, 0x5a, 0x2a, 0x1b, 0xcf, 0x3e, 0x2a, 0x0a, 0x8a, 0xcd, 0x16, 0x9e,
	0x87, 0x59, 0xa6, 0x62, 0x3b, 0x72, 0xb4, 0x25, 0x56, 0xa1, 0xea, 0x54, 0x87, 0xfa, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0x83, 0x8e, 0x5d, 0xac, 0x05, 0x00, 0x00,
}