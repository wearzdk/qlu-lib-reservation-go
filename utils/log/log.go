package log

import (
	"fmt"
	"github.com/fatih/color"
)

// 日志

type log struct {
	Info string
	Type string
}

var logs []log

func init() {
	logs = []log{}
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
