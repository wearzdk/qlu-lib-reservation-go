package app

import (
	"QluTakeLesson/utils/config"
	"QluTakeLesson/utils/log"
	"encoding/json"
)

// 可预约时间段

type Segment struct {
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	AreaId     int    `json:"area_id"`
	BookTimeId int    `json:"book_time_id"`
	Day        string `json:"day"`
	Status     int    `json:"status"`
}

var segmentList []Segment

func InitSegment() {
	segmentList = []Segment{}
	// 在配置中读取
	file := config.FileReading("segment.json")
	if file != nil {
		err := json.Unmarshal(file, &segmentList)
		if err != nil {
			segmentList = []Segment{}
			log.Warning("读取可预约时间段列表失败")
			log.Error(err)
			UpdateSegmentList()
		}
	} else {
		UpdateSegmentList()
	}
}

// UpdateSegmentList 更新可预约时间段
func UpdateSegmentList() {
	log.Info("正在更新可预约时间段")
	// 获取所有区域
	areaList := GetAreaList()
	// 清空列表
	segmentList = []Segment{}
	// 遍历区域
	for _, area := range areaList {
		// 获取可预约时间段
		AreaSegmentRes, _ := GetAreaSegment(area.AreaId)
		// 遍历时间段
		for _, segment := range AreaSegmentRes.Data.List {
			// 添加到列表
			segmentList = append(segmentList, Segment{
				StartTime:  segment.StartTime,
				EndTime:    segment.EndTime,
				AreaId:     segment.Area,
				BookTimeId: segment.BookTimeId,
				Day:        segment.Day,
				Status:     segment.Status,
			})
		}
	}
	saveSegmentList()
}

// saveSegmentList 保存可预约时间段
func saveSegmentList() {
	// 序列化
	data, err := json.Marshal(segmentList)
	if err != nil {
		log.Error(err)
		return
	}
	// 保存到文件
	config.FileSaving("segment.json", data)
	log.Info("可预约时间段更新完成")
}

// GetSegmentList 获取可预约时间段列表
func GetSegmentList() []Segment {
	return segmentList
}

// GetSegmentListByAreaId 根据区域 ID 获取可预约时间段列表
func GetSegmentListByAreaId(areaId int) []Segment {
	var list []Segment
	for _, segment := range segmentList {
		if segment.AreaId == areaId {
			list = append(list, segment)
		}
	}
	return list
}

// GetSegmentListByAreaIdAndDay 根据区域 ID 和日期获取可预约时间段列表
func GetSegmentListByAreaIdAndDay(areaId int, day string) *Segment {
	for _, segment := range segmentList {
		//log.Debug("segment.Day:"+segment.Day, "day:"+day, "segment.AreaId:"+strconv.Itoa(segment.AreaId), "areaId:"+strconv.Itoa(areaId))
		if segment.AreaId == areaId && segment.Day == day {
			//log.Debug("根据区域 ID 和日期获取可预约时间段列表成功")
			return &segment
		}
	}
	return nil
}
