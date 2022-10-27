package app

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"encoding/json"
	"math/rand"
	"time"
)

// 区域数据

type Area struct {
	AreaId   int    `json:"area_id"`
	AreaName string `json:"area_name"`
	SeatList []Seat `json:"seat_list"`
}

type Seat struct {
	SeatId   int    `json:"seat_id"`
	SeatName string `json:"seat_name"`
}

var areas []Area

func InitArea() {
	log.Info("正在初始化区域列表信息....")
	// 初始化
	areas = []Area{}
	// 在配置中读取
	file := config.FileReading("area.json")
	if file == nil {
		FetchAreaList()
	} else {
		err := json.Unmarshal(file, &areas)
		if err != nil {
			log.Info("区域列表信息初始化失败，正在重新初始化")
			areas = []Area{}
			FetchAreaList()
		}
	}
}

// FetchAreaList 获取区域列表
func FetchAreaList() {
	log.Info("正在更新区域列表信息")
	// 先清空
	areas = []Area{}
	var areaList *AreasStateResponse
	var err error
	for {
		// 获取区域列表
		areaList, err = GetAreasState()
		if err != nil {
			log.Error(err, "获取区域列表失败 正在重试")
			time.Sleep(time.Second * 5)
			continue
		}
		if areaList == nil {
			log.Error(err, "获取区域列表失败 正在重试")
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	for _, l1Area := range areaList.Data.List.ChildArea {
		if l1Area.Type == 1 {
			// 获取座位列表
			seatListRes, err := GetAreaInfo(l1Area.Id)
			if err != nil {
				log.Error(err, "获取区域列表信息失败，正在重试")
				continue
			}
			if seatListRes == nil {
				log.Warning("获取区域列表信息失败，正在重试")
				continue
			}
			// 添加区域
			var seatList []Seat
			for _, seat := range seatListRes.Data.List {
				seatList = append(seatList, Seat{
					SeatId:   seat.Id,
					SeatName: seat.Name,
				})
			}
			areas = append(areas, Area{
				AreaId:   l1Area.Id,
				AreaName: l1Area.Name,
				SeatList: seatList,
			})
		}
	}
	SaveAreaList()
}

// UpdateAreaList 更新区域列表
func UpdateAreaList() {
	FetchAreaList()
}

// SaveAreaList 保存区域列表
func SaveAreaList() {
	// 保存到配置文件
	file, _ := json.MarshalIndent(areas, "", "  ")
	config.FileSaving("area.json", file)
	log.Info("区域列表信息更新完毕，文档已保存")
}

// GetAreaList 获取区域列表
func GetAreaList() []Area {
	// 检查区域列表是否为空
	if len(areas) == 0 {
		UpdateAreaList()
	}
	return areas
}

// GetSeatList 获取座位列表
func GetSeatList(areaId int) []Seat {
	for _, area := range areas {
		if area.AreaId == areaId {
			return area.SeatList
		}
	}
	return []Seat{}
}

// GetAreaName 获取区域名称
func GetAreaName(areaId int) string {
	for _, area := range areas {
		if area.AreaId == areaId {
			return area.AreaName
		}
	}
	return ""
}

// GetSeatName 获取座位名称
func GetSeatName(areaId int, seatId int) string {
	if seatId == -1 {
		return "随机"
	}
	for _, area := range areas {
		if area.AreaId == areaId {
			for _, seat := range area.SeatList {
				if seat.SeatId == seatId {
					return seat.SeatName
				}
			}
		}
	}
	return ""
}

// GetRandomSeat 获取随机座位
func GetRandomSeat(areaId int) int {
	// 取随机数
	randInt := int16(rand.Float64() * 1000)
	for _, area := range areas {
		if area.AreaId == areaId {
			return area.SeatList[randInt%int16(len(area.SeatList))].SeatId
		}
	}
	return -1
}
