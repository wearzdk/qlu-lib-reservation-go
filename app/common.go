package app

import (
	"QluTakeLesson/utils/log"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func CheckFileExist(s string) bool {
	_, err := os.Stat(s)
	return err == nil || os.IsExist(err)
}

// ShowAreaList 展示区域列表
func ShowAreaList() {
	// 获取可预约列表
	areasState, err := GetAreasState()
	if areasState == nil {
		log.Warning("获取可预约列表失败")
		log.Error(err)
		return
	}
	println("可预约列表：")
	for _, area := range areasState.Data.List.ChildArea {
		if area.IsValid == 1 && area.Type == 1 {
			color.Cyan("ID：%d\t区域：%s\t\t总计：%d\t可用：%d\n", area.Id, area.Name, area.TotalCount, area.TotalCount-area.UnavailableSpace)
		}
	}
}

// ShowSeatList
// 展示座位列表
// areaId 区域ID
// seatType 座位类型 0：全部 1：空闲 2：已占用
// /*
func ShowSeatList(areaId int, seatType int) {
	// 获取区域信息
	areaInfo, err := GetAreaInfo(areaId)
	if err != nil {
		log.Error(err)
		return
	}
	if areaInfo == nil {
		log.Warning("获取区域信息失败")
		return
	}
	for _, seat := range areaInfo.Data.List {
		if seatType == 1 && seat.Status == 1 {
			// 空闲
			color.Cyan("ID：%d\t座位：%s\t\t状态：%s\n", seat.Id, seat.Name, seat.StatusName)
		} else if seatType == 2 && seat.Status != 1 {
			// 已占用
			color.Cyan("ID：%d\t座位：%s\t\t状态：%s\n", seat.Id, seat.Name, seat.StatusName)
		} else if seatType == 0 {
			// 全部
			color.Cyan("ID：%d\t座位：%s\t\t状态：%s\n", seat.Id, seat.Name, seat.StatusName)
		}
	}
}

// ClearScreen 清空屏幕
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
