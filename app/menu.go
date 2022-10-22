package app

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"crypto/md5"
	"fmt"
	"github.com/fatih/color"
	"os"
)

// 菜单

// ShowMainMenu 输出主功能菜单
func ShowMainMenu() string {
	var menu = []string{
		"0. 退出",
		"1. 登录",
		"2. 区域列表(实时)",
		"3. 区域列表(离线)",
		"4. 查看预约选项",
		"5. 等待预约",
		"6. 更新区域信息",
		"7. 设置预约开始时间",
		"8. 配置锐捷认证",
		"9. 验证码识别配置",
	}
	for {
		for _, item := range menu {
			println(item)
		}
		println("请输入选项：")
		inputCode := InputInt("菜单")
		switch inputCode {
		case 1:
			Login()
		case 2:
			ShowAreaListMenu(true)
		case 3:
			ShowAreaListMenu(false)
		case 4:
			ShowReservationListMenu()
		case 5:
			return "serve"
		case 6:
			UpdateAreaList()
		case 7:
			println("请输入预约开始时间（格式：06:00:00 或 5:40）")
			time, err := InputDayTime("预约开始时间")
			if err == nil {
				config.SetReservationStartTime(time)
				println("设定开始时间为" + config.GetReservationStartTimeFormat())
			} else {
				println(err.Error())
			}
		case 8:
			SetRuiJieConfig()
		case 9:
			SetCaptchaConfigMenu()
		case 0:
			os.Exit(0)
		default:
			println("输入错误")
		}
	}
}

// ShowAreaListMenu 展示区域列表菜单
func ShowAreaListMenu(onLine bool) {
	if onLine {
		ShowAreaList()
	} else {
		ShowAreaListOffline()
	}
	menu := []string{
		"输入ID选择区域",
		"输入0返回上级菜单",
	}
	for _, item := range menu {
		print(item + "\t")
	}
	println("请输入选项：")
	inputCode := InputInt("选项")
	switch inputCode {
	case 0:
		return
	default:
		if onLine {
			ShowSeatListMenu(inputCode)
		} else {
			ShowSeatListOffline(inputCode)
		}
	}
}

func ShowSeatListOffline(areaId int) {
	seats := GetSeatList(areaId)
	for _, seat := range seats {
		color.Cyan("ID：%d\t座位：%s\t\n", seat.SeatId, seat.SeatName)
	}
	ShowSeatListMenuAfter(areaId)
}

func ShowAreaListOffline() {
	println("区域列表(离线)")
	for _, area := range GetAreaList() {
		color.Cyan("ID：%d\t区域：%s\t\t总计：%d\t\n", area.AreaId, area.AreaName, len(area.SeatList))
	}
}

// ShowSeatListMenu 展示座位列表菜单
func ShowSeatListMenu(areaId int) {
	menu := []string{
		"0. 返回上级菜单",
		"1. 查看全部座位列表",
		"2. 查看空闲座位列表",
		"3. 将此区域加入预约列表（随机选择座位）",
	}

	for {
		for _, item := range menu {
			print(item + "\t")
		}
		println("请输入选项：")
		inputCode := InputInt("选项")
		switch inputCode {
		case 0:
			return
		case 1:
			ShowSeatList(areaId, 0)
			ShowSeatListMenuAfter(areaId)
		case 2:
			ShowSeatList(areaId, 1)
			ShowSeatListMenuAfter(areaId)
		case 3:
			AddReservation(areaId, -1)
		default:
			println("输入错误")
		}
	}
}

// ShowSeatListMenuAfter 展示列表后可选择的操作
func ShowSeatListMenuAfter(areaId int) {
	menu := []string{
		"0. 返回上级菜单",
		"1. 将指定座位加入预约列表",
		"2. 批量指定座位加入预约列表（逗号分隔，支持区间）",
		"3. 将此区域加入预约列表（随机选择座位）",
	}
	for {
		for _, item := range menu {
			print(item + "\t")
		}
		println("请输入选项：")
		inputCode := InputInt("选项")
		switch inputCode {
		case 0:
			return
		case 1:
			println("请输入座位号：")
			seatId := InputInt("座位号")
			AddReservation(areaId, seatId)
			println("添加成功")
		case 2:
			println("请输入座位号（逗号分隔，支持区间）：")
			seatIds := InputStr("座位号")
			AddReservationBySeatIds(areaId, seatIds)
			println("添加成功")
		case 3:
			AddReservation(areaId, -1)
		}
	}
}

// ShowReservationListMenu 展示预约列表菜单
func ShowReservationListMenu() {
	for {
		menu := []string{
			"0. 返回上级菜单",
			"1. 查看预约列表",
			"2. 清空预约列表",
		}
		for _, item := range menu {
			print(item + "\t")
		}
		println("请输入选项：")
		inputCode := InputInt("选项")
		switch inputCode {
		case 0:
			return
		case 1:
			ShowReservationList()
		case 2:
			ClearReservationList()
			println("清空成功")
		default:
			println("输入错误")
		}
	}
}

// SetRuiJieConfig 设置锐捷配置
func SetRuiJieConfig() {
	isEnable := inputBool("启用校园网认证")
	if isEnable {
		println("输入校园网登录用户名")
		userName := InputStr("用户名")
		println("输入校园网登录密码 默认身份证后六位")
		password := InputStr("密码")
		config.Config.RuiJie.Enable = true
		config.Config.RuiJie.Username = userName
		config.Config.RuiJie.Password = password
		config.SaveConfig()
	} else {
		config.Config.RuiJie.Enable = false
		config.SaveConfig()
	}
}

// SetCaptchaConfigMenu 设置验证码配置
func SetCaptchaConfigMenu() {
	menu := []string{
		"0. 返回上级菜单",
		"1. 使用本地验证码识别",
		"2. 使用超级鹰",
	}
	for _, item := range menu {
		print(item + "\t")
	}
	println("请输入选项：")
	inputCode := InputInt("选项")
	switch inputCode {
	case 0:
		return
	case 1:
		config.Config.ThirdOcr.Enable = false
		config.SaveConfig()
		log.Info("使用本地验证码识别")
	case 2:
		SetCJYConfig()
	}
}

// SetCJYConfig 设置超级鹰配置
func SetCJYConfig() {
	isEnable := inputBool("启用超级鹰")
	isRewrite := true
	if config.Config.ThirdOcr.CJY.Username != "" {
		println("检测到已有超级鹰配置，是否覆盖？")
		isRewrite = inputBool("覆盖")
	}
	if isEnable {
		if isRewrite {
			println("输入超级鹰用户名")
			userName := InputStr("用户名")
			println("输入超级鹰密码")
			password := InputStr("密码")
			// 转换为32位小写md5
			password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
			println("输入超级鹰软件ID")
			appID := InputStr("软件ID")
			config.Config.ThirdOcr.Enable = true
			config.Config.ThirdOcr.Name = "cjy"
			config.Config.ThirdOcr.CJY.Username = userName
			config.Config.ThirdOcr.CJY.Password = password
			config.Config.ThirdOcr.CJY.SoftId = appID
			config.SaveConfig()
		}
		log.Info("超级鹰配置成功")
	} else {
		config.Config.ThirdOcr.Enable = false
		config.SaveConfig()
		log.Info("已关闭超级鹰")
	}
}
