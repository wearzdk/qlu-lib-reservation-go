package app

import (
	"QluTakeLesson/utils/log"
	"fmt"
	"strconv"
	"strings"
)

// 输入输出验证

func InputInt(name string) int {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Warning(name + "输入错误")
		return -1
	}
	return input
}
func InputStr(name string) string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Warning(name + "输入错误")
		return ""
	}
	return input
}
func InputDecimals(name string) float32 {
	var input float32
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(name + "输入错误")
		return 0
	}
	return input
}

// InputDayTime 输入时间
func InputDayTime(name string) (int, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Warning(name + "输入错误")
		return 0, err
	}
	// 将输入的时间转换为秒数 输入格式为 00:00:00
	var secondSum int
	// 以:分割字符串
	inputArr := strings.Split(input, ":")
	if len(inputArr) == 0 {
		inputArr = strings.Split(input, "：")
	}
	if len(inputArr) == 3 {
		// 将字符串转换为int
		hour, _ := strconv.Atoi(inputArr[0])
		minute, _ := strconv.Atoi(inputArr[1])
		second, _ := strconv.Atoi(inputArr[2])
		secondSum = hour*3600 + minute*60 + second
	} else if len(inputArr) == 2 {
		hour, _ := strconv.Atoi(inputArr[0])
		minute, _ := strconv.Atoi(inputArr[1])
		secondSum = hour*3600 + minute*60
	} else {
		fmt.Println("输入的时间格式有误")
		err = fmt.Errorf("输入错误")
		return 0, err
	}
	return secondSum, nil
}

// inputChoose 输入选择 Y/n
func inputBool(name string) bool {
	println("是否" + name + "？(Y/n)")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Warning("输入错误")
		return false
	}
	if input == "y" || input == "Y" {
		return true
	}
	return false
}
