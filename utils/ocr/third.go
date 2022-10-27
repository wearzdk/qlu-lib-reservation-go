package ocr

import (
	"QluTakeLesson/utils/CJYSDK"
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"errors"
	"fmt"
	"os"
	"time"
)

// 使用第三方ocr平台

// 已经请求的验证码列表
type codeLog struct {
	code string
	time int64
	id   string
}

var codeLogs map[string]codeLog
var thirdOcr *config.ThirdOcrOption

func init() {
	codeLogs = make(map[string]codeLog)
	thirdOcr = &config.Config.ThirdOcr
}

// GetPicCode 获取验证码结果
func GetPicCode(file []byte) string {
	if !thirdOcr.Enable {
		// 不使用第三方ocr
		log.Debug("使用本地ocr")
		// 保存到本地
		err := os.WriteFile("code.png", file, 0666)
		if err != nil {
			log.Error(err, "保存验证码失败")
		}
		// 检测是否存在本地ocr
		if _, err = os.Stat("ocr/tesseract.exe"); os.IsNotExist(err) {
			log.Error(errors.New("未找到本地ocr"), "请打开程序目录下的code.png手动输入验证码")
			var code string
			_, err = fmt.Scanln(&code)
			if err != nil {
				log.Error(err, "输入验证码失败")
				return ""
			}
			return code
		}
		ocr, err := OCR("code.png")
		if err != nil {
			log.Error(err, "识别验证码失败")
			return ""
		}
		return ocr
	}
	if thirdOcr.Enable && thirdOcr.Name == "cjy" {
		// 使用超级鹰
		cjyResp := CJYSDK.GetPicVal(file)
		if cjyResp != nil {
			// 请求成功
			codeLogs[cjyResp.PicStr] = codeLog{
				code: cjyResp.PicStr,
				time: time.Now().Unix(),
				id:   cjyResp.PicId,
			}
			return cjyResp.PicStr
		}
	}
	return ""
}

// ReportResult 报告结果
func ReportResult(code string, isSuccess bool) {
	if !thirdOcr.Enable {
		return
	}
	if thirdOcr.Enable && thirdOcr.Name == "cjy" {
		// 使用超级鹰
		if codeLogs[code].code == code {
			// 有记录
			if !isSuccess {
				// 失败
				CJYSDK.ReportError(codeLogs[code].id)
			}
			delete(codeLogs, code)
		}
	}

}
