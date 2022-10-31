package app

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 预约列表管理

type Reservation struct {
	AreaId   int    `json:"area_id"`
	SeatId   int    `json:"seat_id"`
	AreaName string `json:"area_name"`
	SeatName string `json:"seat_name"`
}

// 计划预约列表
var reservationList []Reservation

// ReserveTime 预约时间  全局变量
var ReserveTime time.Time

func init() {
	// 初始化
	reservationList = []Reservation{}
	file := config.FileReading("reservation.json")
	if file == nil {
		return
	}
	err := json.Unmarshal(file, &reservationList)
	if err != nil {
		log.Warning("预约列表初始化失败")
		log.Error(err)
		return
	}
	// 初始化预约时间 为明天
	ReserveTime = time.Now().AddDate(0, 0, 1)
}

// AddReservation 添加预约
func AddReservation(areaId int, seatId int) {
	areaName := GetAreaName(areaId)
	seatName := GetSeatName(areaId, seatId)
	// 添加预约
	reservationList = append(reservationList, Reservation{
		AreaId:   areaId,
		SeatId:   seatId,
		AreaName: areaName,
		SeatName: seatName,
	})
	SaveReservationList()
}

// AddReservationBySeatIds 批量添加预约
func AddReservationBySeatIds(areaId int, seatIds string) {
	// 解析座位ID
	seatIdList := ParseSeatIds(seatIds)
	for _, seatId := range seatIdList {
		AddReservation(areaId, seatId)
	}
}

func ParseSeatIds(ids string) []int {
	var seatIdList []int
	// 解析座位ID
	// 首先逗号分割
	seatIdStr := strings.Split(ids, ",")
	// 然后解析
	for _, seatIdItem := range seatIdStr {
		// 是否有横杠
		if strings.Contains(seatIdItem, "-") {
			// 有横杠
			// 横杠分割
			seatIdRange := strings.Split(seatIdItem, "-")
			// 解析
			start, _ := strconv.Atoi(seatIdRange[0])
			end, _ := strconv.Atoi(seatIdRange[1])
			for j := start; j <= end; j++ {
				seatIdList = append(seatIdList, j)
			}
		} else {
			// 没有横杠
			seatId, _ := strconv.Atoi(seatIdItem)
			seatIdList = append(seatIdList, seatId)
		}
	}
	return seatIdList
}

// SaveReservationList 保存预约列表
func SaveReservationList() {
	// 保存预约列表
	file, err := json.Marshal(reservationList)
	if err != nil {
		log.Warning("预约列表保存失败")
		log.Error(err)
		return
	}
	config.FileSaving("reservation.json", file)
}

// GetReservationList 获取预约列表
func GetReservationList() []Reservation {
	return reservationList
}

func (r *Reservation) Print() {
	fmt.Printf("区域：%s 座位：%s\n", r.AreaName, r.SeatName)
}

// WaitForReservation 等待预约
func WaitForReservation() {
	for {
		WaitForReservationTime()
		// 遍历预约列表
		reservationResultChan := make(chan bool)
		for {
			isSuccess := false
			// 监听预约结果
			go func(isSuccess *bool) {
				for {
					*isSuccess = <-reservationResultChan
				}
			}(&isSuccess)
			// 遍历预约列表进行预约
			for _, reservation := range reservationList {
				// 执行预约
				seatId := reservation.SeatId
				log.Info("正在预约：" + reservation.AreaName + " " + reservation.SeatName)
				if reservation.SeatId == -1 {
					// 随机座位
					seatId = GetRandomSeat(reservation.AreaId)
					if seatId == -1 {
						log.Info("获取随机座位失败")
						continue
					}
				}
				go func(areaId int, seatId int) {
					reservationResultChan <- Reserve(areaId, seatId)
				}(reservation.AreaId, seatId)
				// 随机等待1-2秒
				time.Sleep(time.Duration(rand.Intn(1)+1) * time.Second)
				if isSuccess {
					break
				}
			}
			if isSuccess {
				break
			}
			// 如果已经预约30分钟，则停止预约
			if time.Now().Sub(ReserveTime).Minutes() > 30 {
				log.Info("已经预约30分钟，停止预约")
				break
			}
		}
		log.Info("今日预约已完成")
	}

}

// WaitForReservationTime 等待预约时间
func WaitForReservationTime() {
	startTime := config.GetReservationStartTime()
	if config.Config.LastReservationTime == time.Now().Add(+time.Hour*24).Format("2006-01-02") {
		// 今天已经预约过了
		log.Info("今天已经预约过了")
		time.Sleep(60 * time.Minute)
	}
	log.Info("等待预约中...")
	log.Info("预约开始时间：%s", config.GetReservationStartTimeFormat())
	for {
		// 获取当前时间 (当前距0点的秒数)
		currentTime := time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()
		// 是否到达预约时间
		if currentTime >= startTime && currentTime < startTime+1200 {
			// 到达预约时间
			break
		}
		if startTime-currentTime > 0 {
			log.Info("距离预约时间还有%d秒\n", startTime-currentTime)
		} else {
			log.Info("距离预约时间还有%d秒\n", startTime-currentTime+86400)
		}
		time.Sleep(10 * time.Second)
	}
}

// ShowReservationOptions 展示预约选项
func ShowReservationOptions() {
	reservationList := GetReservationList()
	if reservationList == nil {
		println("获取预约列表失败")
		return
	}
	println("预约列表：")
	for _, reservation := range reservationList {
		reservation.Print()
	}
}

// ShowReservationList 展示预约列表
func ShowReservationList() {
	ShowReservationOptions()
}

// ClearReservationList 清空预约列表
func ClearReservationList() {
	reservationList = nil
	SaveReservationList()
}

// Reserve 预约
func Reserve(areaId, seatId int) bool {
	// 获取可预约时间
	segmentInfo := GetSegmentListByAreaIdAndDay(areaId, ReserveTime.Format("2006-01-02"))
	for {
		if segmentInfo == nil {
			log.Warning("预约时间信息已过期,正在重新获取")
			UpdateSegmentList()
			ReserveTime = time.Now().Add(time.Hour * 24)
			segmentInfo = GetSegmentListByAreaIdAndDay(areaId, ReserveTime.Format("2006-01-02"))
		} else {
			break
		}
	}

	// 预约
	reserveResponse, err := PostReserve(seatId, segmentInfo.BookTimeId)
	if err != nil {
		log.Warning("预约失败,请求时发生错误")
		log.Error(err)
		return false
	}
	if reserveResponse == nil {
		log.Warning("预约失败,返回结果为空")
		return false
	}
	// 预约成功
	if reserveResponse.Status == 1 {
		reserveResponse.Print()
		// 保存预约信息
		config.Config.LastReservationTime = ReserveTime.Format("2006-01-02")
		config.Config.LastReservation.AreaName = reserveResponse.Data.List.SpaceInfo.AreaInfo.NameMerge
		config.Config.LastReservation.Seat = reserveResponse.Data.List.SpaceInfo.Name
		config.SaveConfig()
		return true
	} else {
		log.Warning(reserveResponse.Msg)
		if reserveResponse.Msg == "没有登录或登录已超时" {
			// 重新登录
			Login()
		}
		return false
	}
}

// Print 预约结果信息提示
func (r *ReserveResponse) Print() {
	color.Cyan("预约结果：%s\n", r.Msg)
	color.Cyan("预约开始时间：%s\t 结束时间：%s\t\n", r.Data.List.Starttime, r.Data.List.EndTime.Date)
	color.Cyan("区域信息：%s %s\n", r.Data.List.SpaceInfo.AreaInfo.NameMerge, r.Data.List.SpaceInfo.Name)
}
