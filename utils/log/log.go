package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

// 日志

var logFile *os.File

func init() {
	// 打开文件
	var err error
	logFile, err = os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Info 日志

func Info(format string, a ...interface{}) {
	out := "INFO: " + fmt.Sprintf(format, a...)
	color.Green(out)
	WriteLog(out, "INFO: ")

}

// Error 日志
func Error(err error, info ...string) {
	out := "ERROR: " + fmt.Sprint(info) + " " + err.Error()
	color.Red(out)
	WriteLog(out, "ERROR: ")
}

// Warning 日志
func Warning(info string) {
	out := "WARNING: " + info
	color.Yellow(out)
	WriteLog(out, "WARNING: ")
}

func Debug(info ...interface{}) {
	out := "DEBUG: " + fmt.Sprint(info)
	color.Blue(out)
	WriteLog(out, "DEBUG: ")
}

// WriteLog 写入日志文件
func WriteLog(log, logType string) {
	// 写入文件
	_, err := logFile.WriteString(time.Now().Format("2006-01-02 15:04:05") + log + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}
