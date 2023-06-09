package network

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gobot/back"
	"gobot/pkg/logger"
	"nba/network/pb"
	"nba/types"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BaseAuther struct {
	*types.AuthData
}

type Authreq struct {
	SdkName       string `json:"sdkName,omitempty"`
	SdkUid        string `json:"sdkUid,omitempty"`
	ClientVersion string `json:"clientVersion,omitempty"`
	ValidateInfo  string `json:"validateInfo,omitempty"`
	Channel       string `json:"channel,omitempty"`
	Device        string `json:"device,omitempty"`
}

type Authrsp struct {
	Code          int32                 `json:"error"`
	Session       string                `json:"session"`
	UID           int64                 `json:"uid"`
	Addr          string                `json:"server"`
	Zones         []*pb.LogicServerInfo `json:"zones"`
	LastLoginZone int32                 `json:"lastzoneid"`
	Notice        string                `json:"notice"`
	SdkUID        string
}

func (req Authreq) Data() string {
	d, _ := json.Marshal(req)
	return base64.StdEncoding.EncodeToString(d)
}

func NewAuther(auth *types.AuthData) back.IAuth {
	base := &BaseAuther{
		AuthData: auth,
	}

	if types.SdkConf.Name == "lmd" {
		return &LmdAuther{BaseAuther: *base}
	}
	return base
}

// sdk 认证
func (auth *BaseAuther) SdkAuth(cfg interface{}) (rsp interface{}, err error) {
	rsp = &Authreq{
		SdkName:       "mine",
		Channel:       "mine",
		SdkUid:        fmt.Sprintf("%v", auth.SdkAccount),
		ClientVersion: auth.Conf.ClientVersion,
		Device:        fmt.Sprintf("%v", auth.Conf.Device),
	}
	return
}

func httpGet(url string) (rsp *http.Response, err error) {
	retry := 10
	for i := 0; i < retry; i++ {
		rsp, err = http.Get(url)
		if err == nil {
			return
		}
		time.Sleep(time.Millisecond * 50 * (time.Duration(i + 1)))
	}
	return
}

// 游戏服认证
func (auth *BaseAuther) GameAuth(reqinfo interface{}) (data interface{}, err error) {
	reqdata := reqinfo.(*Authreq)
	url := fmt.Sprintf("%v?data=%v", auth.Conf.URL, reqdata.Data())
	resp, err := httpGet(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var info Authrsp
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return
	}
	info.SdkUID = reqdata.SdkUid
	data = &info
	return
}

type LmdAuther struct {
	BaseAuther
}

// sdk 认证
func (auth *LmdAuther) SdkAuth(cfg interface{}) (rsp interface{}, err error) {
	info := types.SdkConf
	data := url.Values{}
	account := auth.SdkAccount
	data.Set("appID", info.AppID)
	data.Set("timestamp", strconv.Itoa(int(time.Now().Unix())))
	data.Set("accountID", account)
	data.Set("deviceID", account)
	data.Set("channelID", info.ChannelID)
	signStr := data.Encode() + "&secretKey=" + info.AppKey
	md5Str := strings.ToUpper(MD5Str(signStr))
	data.Set("sign", md5Str)
	data.Set("sign", md5Str)
	for _, k := range []string{
		"appID", "timestamp",
		"accountID", "deviceID",
		"channelID", "sign",
	} {
		val := data.Get(k)
		data.Set(k, decrypt(val, 2))
	}
	dataEncode := data.Encode()

	resp, err := http.Post(info.URL,
		"application/x-www-form-urlencoded",
		strings.NewReader(dataEncode))
	if err != nil {
		logger.Error("lmd Post", "err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	var bodyResult struct {
		Body string `json:"body"` // 加密字符串
	}
	err = json.NewDecoder(resp.Body).Decode(&bodyResult)
	if err != nil {
		logger.Error("lmdSdk", "Body", bodyResult.Body)
		return nil, err
	}
	logger.Debug("lmdSdk", "body", bodyResult.Body)

	var result struct {
		Code int32                  `json:"code"` //0(登录认证成功)；其他失败
		Data map[string]interface{} `json:"data"`
		Msg  string                 `json:"msg"` //描述信息
	}

	base64Body, err := base64.RawStdEncoding.DecodeString(bodyResult.Body)
	if err != nil {
		logger.Error("lmdSdk", "base64Body err", err.Error())
		return nil, err
	}
	logger.Debug("lmdSdk", "base64body", string(base64Body))
	for i := 0; i < len(base64Body); i++ {
		base64Body[i] ^= types.SdkConf.Secret
	}
	logger.Debug("lmdSdk", "decodedResult", string(base64Body))

	err = json.Unmarshal(base64Body, &result)
	if err != nil {
		logger.Error("lmdSdk", "err", err.Error())
		return nil, err
	}

	logger.Debug("lmdSdk", "res", result)

	if result.Code != 0 {
		return nil, fmt.Errorf("%v", result)
	}
	accountData := result.Data["accounts"].(map[string]interface{})
	infoData := result.Data["info"].(map[string]interface{})

	rsp = &Authreq{
		SdkName:       "lmd",
		Channel:       "lmd",
		SdkUid:        fmt.Sprintf("%v", accountData["uid"]),
		ValidateInfo:  fmt.Sprintf("%v", infoData["access_token"]),
		ClientVersion: auth.Conf.ClientVersion,
		Device:        account,
	}
	return
}

func MD5Str(s string) string {
	return MD5Bytes([]byte(s))
}

func MD5Bytes(s []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(s)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func decrypt(value string, secret byte) string {
	bytes := []byte(value)
	for i := range bytes {
		bytes[i] ^= secret
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
