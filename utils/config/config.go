package config

import (
	"QluTakeLesson/utils/log"
	"encoding/json"
	"fmt"
	"os"
)

// 配置文件

type RuiJieOption struct {
	Enable    bool   `json:"enable"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	UserIndex string `json:"userIndex,omitempty"`
}

type CJYOption struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	SoftId   string `json:"softId,omitempty"`
}

type ThirdOcrOption struct {
	Enable bool      `json:"enable"`
	Name   string    `json:"name,omitempty"`
	CJY    CJYOption `json:"cjy,omitempty"`
}

type Option struct {
	ReservationStartTime int    `json:"reservation_start_time,omitempty"`
	LastReservationTime  string `json:"last_reservation_time,omitempty"`
	LastReservation      struct {
		AreaName string `json:"area_name,omitempty"`
		Seat     string `json:"seat,omitempty"`
	} `json:"last_reservation,omitempty"`
	RuiJie   RuiJieOption   `json:"rui_jie"`
	ThirdOcr ThirdOcrOption `json:"third_ocr,omitempty"`
}

var Config Option

func init() {
	// 初始化
	Config = Option{}
	// 在配置中读取
	file := FileReading("config.json")
	if file != nil {
		err := json.Unmarshal(file, &Config)
		if err != nil {
			log.Error(err, "读取配置失败")
		}
	}
}

func FileSaving(name string, file []byte) {
	// 检查config目录是否存在
	if _, err := os.Stat("config"); os.IsNotExist(err) {
		//	不存在则创建
		err = os.Mkdir("config", 0755)
	}
	// 保存
	err := os.WriteFile("config/"+name, file, 0644)
	if err != nil {
		log.Warning("配置文件保存失败")
		log.Error(err)
	}
}

func FileReading(name string) []byte {
	file, err := os.ReadFile("config/" + name)
	if err != nil {
		return nil
	}
	return file
}

// SaveConfig 保存配置
func SaveConfig() {
	// 序列化
	data, err := json.Marshal(Config)
	if err != nil {
		log.Warning("配置文件保存失败")
		log.Error(err)
		return
	}
	// 保存
	FileSaving("config.json", data)
}

// SetReservationStartTime 设置预约开始时间
func SetReservationStartTime(time int) {
	Config.ReservationStartTime = time
	SaveConfig()
}

// GetReservationStartTime 获取预约开始时间
func GetReservationStartTime() int {
	return Config.ReservationStartTime
}

// GetReservationStartTimeFormat 获取格式化后预约开始时间
func GetReservationStartTimeFormat() string {
	hours := Config.ReservationStartTime / 3600
	minutes := Config.ReservationStartTime / 60 % 60
	seconds := Config.ReservationStartTime % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
