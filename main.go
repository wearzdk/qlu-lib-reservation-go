package main

import (
	"QluTakeLesson/app"
	"QluTakeLesson/task"
	"QluTakeLesson/utils/RuiJIeNet"
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"os"
)

func bootStrap() {
	log.Info("程序启动")
	if config.Config.RuiJie.Enable {
		RuiJIeNet.ExecuteLogin()
	}
	// 加载用户配置
	app.ReloadUserConfig()
	// 初始化区域信息
	app.InitArea()
	// 初始化预约时间信息
	app.InitSegment()
}

func serve() {

	// 启动定时任务
	go task.BootStrap()
	// 进入预约模式
	app.WaitForReservation()
}

func main() {
	bootStrap()
	// 获取运行参数
	args := os.Args
	if len(args) == 1 {
		// 进入主菜单
		action := app.ShowMainMenu()
		if action == "serve" {
			serve()
		}
	} else {
		if args[1] == "service" {
			// 服务模式
			serve()
		}
	}
}
