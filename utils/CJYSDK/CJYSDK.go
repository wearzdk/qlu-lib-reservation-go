package CJYSDK

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"encoding/base64"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

// 超级鹰SDK

var cjyConf *config.CJYOption
var client *resty.Client

type CJYResp struct {
	ErrNo  int    `json:"err_no"`
	ErrStr string `json:"err_str"`
	PicId  string `json:"pic_id"`
	PicStr string `json:"pic_str"`
	Md5    string `json:"md5"`
}

func init() {
	cjyConf = &config.Config.ThirdOcr.CJY
	client = resty.New()
}

func GetPicVal(data []byte) *CJYResp {
	url := "https://upload.chaojiying.net/Upload/Processing.php"
	var dataBase64 string
	dataBase64 = base64.StdEncoding.EncodeToString(data)
	resp, err := client.R().
		SetFormData(map[string]string{
			"user":        cjyConf.Username,
			"pass2":       cjyConf.Password,
			"softid":      cjyConf.SoftId,
			"codetype":    "4004",
			"file_base64": dataBase64,
		}).
		Post(url)
	if err != nil {
		log.Error(err, "超级鹰请求失败")
		return nil
	}
	var cjyResp CJYResp
	err = json.Unmarshal(resp.Body(), &cjyResp)
	if err != nil {
		log.Error(err, "超级鹰响应解析失败")
		return nil
	}
	if cjyResp.ErrNo != 0 {
		log.Error(err, "超级鹰响应错误 "+cjyResp.ErrStr)
		return nil
	}
	return &cjyResp
}

func ReportError(picId string) {
	url := "https://upload.chaojiying.net/Upload/ReportError.php"
	_, err := client.R().
		SetFormData(map[string]string{
			"user":   cjyConf.Username,
			"pass2":  cjyConf.Password,
			"softid": cjyConf.SoftId,
			"id":     picId,
		}).
		Post(url)
	if err != nil {
		log.Error(err, "超级鹰报错失败")
	}
}
