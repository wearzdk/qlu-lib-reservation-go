package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

// 日志

type log struct {
	Info string
	Type string
}

var logs []log

func init() {
	logs = []log{}
	WriteLog()
	// 每隔一分钟写入一次日志
	go func() {
		for {
			WriteLog()
			time.Sleep(time.Minute)
		}
	}()

}

// Info 日志

func Info(format string, a ...interface{}) {
	out := "INFO: " + fmt.Sprintf(format, a...)
	logs = append(logs, log{Info: out, Type: "info"})
	color.Green(out)
}

// Error 日志
func Error(err error, info ...string) {
	out := "ERROR: " + fmt.Sprint(info) + " " + err.Error()
	logs = append(logs, log{Info: out, Type: "error"})
	color.Red(out)
}

// Warning 日志
func Warning(info string) {
	out := "WARNING: " + info
	logs = append(logs, log{Info: out, Type: "warning"})
	color.Yellow(out)
}

func Debug(info ...interface{}) {
	out := "DEBUG: " + fmt.Sprint(info)
	logs = append(logs, log{Info: out, Type: "debug"})
	color.Blue(out)
}

// WriteLog 定期写入日志文件
func WriteLog() {
	logfile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	for _, logItem := range logs {
		_, err = logfile.WriteString(logItem.Info + "\n")
		if err != nil {
			return
		}
	}
	err = logfile.Close()
	if err != nil {
		return
	}
	logs = []log{}
}
