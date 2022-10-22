package ocr

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// OCR 使用tesseract识别图片中的数字
func OCR(path string) (string, error) {
	// 识别图片
	// 执行命令  图片中数字长度4位
	osCmd := exec.Command("ocr\\tesseract", path, "codeResult", "digits")
	// 执行命令
	_, err := osCmd.Output()
	// 读取结果
	result, err := os.ReadFile("codeResult.txt")
	if err != nil {
		return "", err
	}
	// 去除空格
	resultStr := strings.ReplaceAll(string(result), " ", "")
	// 去除换行符
	resultStr = strings.ReplaceAll(resultStr, "\r", "")
	resultStr = strings.ReplaceAll(resultStr, "\n", "")
	// 转为数字
	resultInt, _ := strconv.Atoi(resultStr)
	// 转为字符串
	resultStr = strconv.Itoa(resultInt)
	// 删除文件
	_ = os.Remove("codeResult.txt")
	return resultStr, nil
}
