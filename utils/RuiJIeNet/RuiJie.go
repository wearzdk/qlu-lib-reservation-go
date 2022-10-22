package RuiJIeNet

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"time"
)

// 锐捷认证

type RuiJieLoginResp struct {
	UserIndex         string      `json:"userIndex"`
	Result            string      `json:"result"`
	Message           string      `json:"message"`
	Forwordurl        interface{} `json:"forwordurl"`
	KeepaliveInterval int         `json:"keepaliveInterval"`
	ValidCodeUrl      string      `json:"validCodeUrl"`
}

type RuiJieOnlineInfoResp struct {
	UserIndex         string      `json:"userIndex"`
	Result            string      `json:"result"`
	Message           string      `json:"message"`
	KeepaliveInterval interface{} `json:"keepaliveInterval"`
	MaxLeavingTime    string      `json:"maxLeavingTime"`
	MaxFlow           interface{} `json:"maxFlow"`
	UserName          string      `json:"userName"`
	UserId            string      `json:"userId"`
	UserIp            string      `json:"userIp"`
	UserMac           interface{} `json:"userMac"`
	WebGatePort       interface{} `json:"webGatePort"`
	WebGateIp         interface{} `json:"webGateIp"`
	Service           string      `json:"service"`
	RealServiceName   string      `json:"realServiceName"`
	ApMac             interface{} `json:"apMac"`
	VlanId            interface{} `json:"vlanId"`
	Ssid              interface{} `json:"ssid"`
	AccountInfo       interface{} `json:"accountInfo"`
	LoginType         string      `json:"loginType"`
	UtrustUrl         string      `json:"utrustUrl"`
	UserUrl           string      `json:"userUrl"`
	PubMessage        interface{} `json:"pubMessage"`
	UserGroup         string      `json:"userGroup"`
	AccountFee        string      `json:"accountFee"`
	WlanAcName        interface{} `json:"wlanAcName"`
	HasMabInfo        bool        `json:"hasMabInfo"`
	IsAlowMab         bool        `json:"isAlowMab"`
	UserPackage       string      `json:"userPackage"`
	SelfUrl           interface{} `json:"selfUrl"`
	BallInfo          string      `json:"ballInfo"`
	BallIsDisplay     string      `json:"ballIsDisplay"`
	Notify            string      `json:"notify"`
	Resource          string      `json:"resource"`
	Offlineurl        string      `json:"offlineurl"`
	RedirectUrl       string      `json:"redirectUrl"`
	Announcement      string      `json:"announcement"`
	UserNotice        string      `json:"userNotice"`
	SystemNotice      string      `json:"systemNotice"`
	PortalIp          string      `json:"portalIp"`
	PortalUrl         string      `json:"portalUrl"`
	RghallUrl         string      `json:"rghallUrl"`
	WelcomeTip        string      `json:"welcomeTip"`
	ServiceList       string      `json:"serviceList"`
	IsSuccessService  string      `json:"isSuccessService"`
	IsCloseWinAllowed string      `json:"isCloseWinAllowed"`
	CheckUserLogout   string      `json:"checkUserLogout"`
	SamEdition        string      `json:"samEdition"`
	IsFaq             string      `json:"isFaq"`
	PcClient          string      `json:"pcClient"`
	PcClientUrl       string      `json:"pcClientUrl"`
	PhoneClient       string      `json:"phoneClient"`
	IsAutoLogin       string      `json:"isAutoLogin"`
	DomianName        string      `json:"domianName"`
	NetFlowKey        interface{} `json:"netFlowKey"`
	Errorflowurl      string      `json:"errorflowurl"`
	IsErrorMsg        string      `json:"isErrorMsg"`
	SuccessUrl        string      `json:"successUrl"`
	MabInfo           string      `json:"mabInfo"`
	MabInfoMaxCount   int         `json:"mabInfoMaxCount"`
}

var ruiJieConfig *config.RuiJieOption

// init 初始化
func init() {
	// 初始化
	ruiJieConfig = &config.Config.RuiJie
}

// SaveConfig 保存配置
func SaveConfig() {
	config.SaveConfig()
}

var client = resty.New()

// ExecuteLogin 执行认证
func ExecuteLogin() {
	log.Info("执行校园网认证")
	for {
		url := "http://172.20.255.1:9090/eportal/InterFace.do?method=login"
		resp, err := client.R().
			SetQueryParams(map[string]string{
				"userId":         ruiJieConfig.Username,
				"password":       ruiJieConfig.Password,
				"service":        "internet",
				"queryString":    "&t=wireless-v2-plain&nasip=10.191.0.2",
				"operatorPwd":    "",
				"operatorUserId": "",
				"validcode":      "",
			}).
			Post(url)
		if err != nil {
			log.Error(err)
			return
		}
		var loginResp RuiJieLoginResp
		err = json.Unmarshal(resp.Body(), &loginResp)
		if err != nil {
			log.Error(err, "认证失败")
			return
		}
		if loginResp.Result == "success" {
			log.Info("校园网登录成功!")
			ruiJieConfig.UserIndex = loginResp.UserIndex
			if loginResp.UserIndex == "" {
				time.Sleep(300 * time.Millisecond)
				continue
			}
			SaveConfig()
			break
		} else {
			log.Warning("校园网登录失败! 错误信息: " + loginResp.Message)
		}
	}

}

// QueryLoginResult 查询认证结果
func QueryLoginResult() bool {
	url := "http://172.20.255.1:9090/eportal/InterFace.do?method=getOnlineUserInfo"
	maxRetry := 15
	tryCount := 0
	for {
		tryCount++
		if tryCount >= maxRetry {
			break
		}
		//log.Debug(ruiJieConfig.UserIndex)
		resp, err := client.R().
			SetQueryParams(map[string]string{
				"userIndex": ruiJieConfig.UserIndex,
			}).
			Post(url)
		if err != nil {
			log.Error(err, "查询认证结果失败")
			return false
		}
		var loginResp RuiJieOnlineInfoResp
		err = json.Unmarshal(resp.Body(), &loginResp)
		if err != nil {
			log.Error(err, "查询认证结果失败")
			return false
		}
		if loginResp.Result == "fail" {
			log.Warning("校园网登录失败! 错误信息: " + loginResp.Message)
		} else if loginResp.UserName != "" {
			color.Cyan("%s,%s 已连接到%s", loginResp.UserName, loginResp.WelcomeTip, loginResp.Service)
			return true
		} else {
			log.Info("正在认证...")
			time.Sleep(time.Millisecond * 300)
		}
	}
	return true

}
