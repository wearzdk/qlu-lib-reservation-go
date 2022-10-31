package app

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"QluTakeLesson/utils/ocr"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"os"
	"time"
)

type LoginRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		List struct {
			Id                string      `json:"id"`
			Card              string      `json:"card"`
			Name              string      `json:"name"`
			IdCard            string      `json:"idCard"`
			Gender            int         `json:"gender"`
			Birthday          string      `json:"birthday"`
			JoinTime          string      `json:"joinTime"`
			Wallet            string      `json:"wallet"`
			Saving            string      `json:"saving"`
			FillScore         int         `json:"fillScore"`
			TotalFillScore    int         `json:"totalFillScore"`
			ConsumeScore      int         `json:"consumeScore"`
			TotalConsumeScore int         `json:"totalConsumeScore"`
			Role              interface{} `json:"role"`
			RoleName          string      `json:"roleName"`
			Dept              interface{} `json:"dept"`
			DeptName          string      `json:"deptName"`
			SubDept           interface{} `json:"subDept"`
			SubDeptName       interface{} `json:"subDeptName"`
			Tel               string      `json:"tel"`
			Mobile            interface{} `json:"mobile"`
			Email             string      `json:"email"`
			Qq                interface{} `json:"qq"`
			Status            int         `json:"status"`
			Weixin            interface{} `json:"weixin"`
			HwUpdateFlag      int         `json:"hw_update_flag"`
			SkedbUpdateFlag   int         `json:"skedb_update_flag"`
			ROWNUMBER         string      `json:"ROW_NUMBER"`
			Renegeinfo        interface{} `json:"renegeinfo"`
		} `json:"list"`
		Hash struct {
			Userid      string `json:"userid"`
			AccessToken string `json:"access_token"`
			Expire      string `json:"expire"`
		} `json:"_hash_"`
	} `json:"data"`
}
type UserInfo struct {
	UserId      string
	AccessToken string
	Expire      string
	PHPSESSID   string
	UserName    string
	PassWord    string
}

// LoadUserConfig 加载用户配置
func LoadUserConfig() *UserInfo {
	file := config.FileReading("userInfo.json")
	if file == nil {
		return nil
	}
	userInfo := &UserInfo{}
	err := json.Unmarshal(file, userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}

func ReloadUserConfig() {
	// 加载用户配置
	userInfo := LoadUserConfig()
	if userInfo == nil || userInfo.UserId == "" || userInfo.AccessToken == "" || userInfo.PHPSESSID == "" {
		log.Info("未找到用户配置，请先登录")
		Login()
		userInfo = LoadUserConfig()
	}
	// 初始化网络请求
	InitRequest(userInfo)
}

func SaveUserConfig(userInfo *UserInfo) {
	file, err := json.Marshal(userInfo)
	if err != nil {
		return
	}
	config.FileSaving("userInfo.json", file)
}

func getVerifyCode(client *resty.Client) []byte {
	codeReq, err := client.R().
		SetHeader("Referer", "http://yuyue.lib.qlu.edu.cn/").Get("http://yuyue.lib.qlu.edu.cn/api.php/check")
	if err != nil {
		log.Error(err, "获取验证码失败")
		return nil
	}
	//sessionId := codeReq.Cookies()[0].Value
	if len(codeReq.Cookies()) > 0 {
		err := os.Setenv("PHPSESSID", codeReq.Cookies()[0].Value)
		if err != nil {
			log.Error(err, "设置PHPSESSID失败")
			return nil
		}
	}
	// 保存验证码
	codeImg := codeReq.Body()
	if codeImg == nil {
		log.Error(errors.New("验证码为空"), "获取验证码失败")
		return nil
	}
	return codeImg
}

func Login() {
	// 加载用户配置
	userInfo := LoadUserConfig()
	var account string
	var password string
	if userInfo == nil {
		// 输入账号
		println("请输入账号：")
		account = InputStr("账号")
		// 输入密码
		println("请输入密码：")
		password = InputStr("密码")
	} else {
		account = userInfo.UserName
		password = userInfo.PassWord
		log.Info("已加载配置： 用户名：" + account + " 密码：" + password)
	}
	var sessionId string
	var loginRes LoginRes
	var code string
	var tryCount int
	var codeTryCount int
	client := resty.New()
	for {
		tryCount++
		// 获取验证码
		for {
			codeTryCount++
			if codeTryCount > 100 {
				log.Warning("OCR识别失败，请手动输入验证码")
				code = InputStr("验证码")
				break
			}
			file := getVerifyCode(client)
			if file == nil {
				continue
			}
			sessionId = os.Getenv("PHPSESSID")
			log.Info("PHPSESSID为：" + sessionId)
			// 执行OCR
			code = ocr.GetPicCode(file)
			log.Info("验证码为：" + code)
			if len(code) == 4 {
				break
			} else {
				ocr.ReportResult(code, false)
				log.Warning("识别失败，重新识别")
			}

		}
		//client.SetCookie(&http.Cookie{
		//	Name:  "PHPSESSID",
		//	Value: sessionId,
		//})
		// 登录
		log.Info("正在登录..." + code)
		loginReq, err := client.R().
			SetHeader("Referer", "http://yuyue.lib.qlu.edu.cn/").
			SetFormData(map[string]string{
				"username": account,
				"password": password,
				"verify":   code,
			}).
			Post("http://yuyue.lib.qlu.edu.cn/api.php/login")
		if err != nil {
			panic(err)
		}
		// 获取Json
		err = json.Unmarshal(loginReq.Body(), &loginRes)
		if err != nil {
			panic(err)
		}
		if loginRes.Status != 1 {
			log.Warning("登录失败")
			log.Warning("错误信息：" + loginRes.Msg)
			if loginRes.Msg == "验证码错误，请重新输入" {
				ocr.ReportResult(code, false)
			}
			log.Info("PHPSESSID为：" + sessionId)
		} else {
			log.Info("登录成功")
			ocr.ReportResult(code, true)
			break
		}
	}
	// 保存用户信息
	userId := loginRes.Data.List.Id
	accessToken := loginRes.Data.Hash.AccessToken
	expire := loginRes.Data.Hash.Expire
	// 保存用户信息
	userInfoConf := &UserInfo{
		UserId:      userId,
		AccessToken: accessToken,
		Expire:      expire,
		UserName:    account,
		PassWord:    password,
		PHPSESSID:   sessionId,
	}
	SaveUserConfig(userInfoConf)
	log.Info("用户信息已保存至userInfo.json")
	ReloadUserConfig()
}

// CheckLoginExpire 检查登录是否过期
func CheckLoginExpire() bool {
	userInfo := LoadUserConfig()
	// 获取当前时间戳
	now := time.Now().Add(time.Minute * 20).Unix()
	// 格式化时间
	parse, err := time.Parse("2006-01-02 15:04:05", userInfo.Expire)
	//log.Debug(userInfo.Expire, parse.Unix(), now)
	if err != nil {
		return true
	}
	// 获取过期时间戳
	expire := parse.Unix()
	if now > expire {
		return true
	}
	return false
}
