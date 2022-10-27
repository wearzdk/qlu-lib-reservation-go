package app

import (
	"QluTakeLesson/utils/log"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ping/ping"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"time"
)

var client = resty.New()

type AreasStateResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		List struct {
			ChildArea []struct {
				Id               int    `json:"id"`
				Name             string `json:"name"`
				Enname           string `json:"enname"`
				ParentId         int    `json:"parentId"`
				Levels           int    `json:"levels"`
				IsValid          int    `json:"isValid"`
				Comment          string `json:"comment"`
				Sort             int    `json:"sort"`
				Type             int    `json:"type"`
				Color            string `json:"color"`
				TotalCount       int    `json:"TotalCount"`
				UnavailableSpace int    `json:"UnavailableSpace"`
				HeatOpen         int    `json:"heat_open"`
				Webhidden        int    `json:"webhidden"`
			} `json:"childArea"`
		} `json:"list"`
	} `json:"data"`
}
type AreaInfoResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		List []struct {
			Id         int         `json:"id"`
			No         string      `json:"no"`
			Name       string      `json:"name"`
			Area       int         `json:"area"`
			Category   int         `json:"category"`
			PointX     float64     `json:"point_x"`
			PointX2    interface{} `json:"point_x2"`
			PointX3    interface{} `json:"point_x3"`
			PointX4    interface{} `json:"point_x4"`
			PointY     float64     `json:"point_y"`
			PointY2    interface{} `json:"point_y2"`
			PointY3    interface{} `json:"point_y3"`
			PointY4    interface{} `json:"point_y4"`
			Width      float64     `json:"width"`
			Height     float64     `json:"height"`
			Status     int         `json:"status"`
			StatusName string      `json:"status_name"`
			AreaName   string      `json:"area_name"`
			AreaLevels int         `json:"area_levels"`
			AreaType   int         `json:"area_type"`
			AreaColor  interface{} `json:"area_color"`
		} `json:"list"`
	} `json:"data"`
}
type ReserveResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		List struct {
			Id            int    `json:"id"`
			No            string `json:"no"`
			Booker        string `json:"booker"`
			SpaceCategory int    `json:"spaceCategory"`
			Space         string `json:"space"`
			IsSingle      int    `json:"isSingle"`
			MemberCount   int    `json:"memberCount"`
			BeginTime     struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"beginTime"`
			EndTime struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"endTime"`
			Title       string `json:"title"`
			Application string `json:"application"`
			IsPublic    int    `json:"isPublic"`
			UpdateTime  struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"updateTime"`
			ExamTime struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"examTime"`
			Examinant  interface{} `json:"examinant"`
			ExamResult string      `json:"examResult"`
			SignIn     int         `json:"signIn"`
			SignOut    int         `json:"signOut"`
			Status     int         `json:"status"`
			ROWNUMBER  string      `json:"ROW_NUMBER"`
			SpaceInfo  struct {
				Id        int    `json:"id"`
				No        string `json:"no"`
				Name      string `json:"name"`
				Area      int    `json:"area"`
				Category  int    `json:"category"`
				Status    int    `json:"status"`
				Detail    string `json:"detail"`
				ROWNUMBER string `json:"ROW_NUMBER"`
				AreaInfo  struct {
					Id             int         `json:"id"`
					Name           string      `json:"name"`
					ParentId       int         `json:"parentId"`
					Levels         int         `json:"levels"`
					IsValid        int         `json:"isValid"`
					Comment        string      `json:"comment"`
					Sort           int         `json:"sort"`
					Type           int         `json:"type"`
					Color          interface{} `json:"color"`
					Enname         string      `json:"enname"`
					NameMerge      string      `json:"nameMerge"`
					EnnameMerge    string      `json:"ennameMerge"`
					ROWNUMBER      string      `json:"ROW_NUMBER"`
					ParentAreaInfo []struct {
						Id          int         `json:"id"`
						Name        string      `json:"name"`
						ParentId    int         `json:"parentId"`
						Levels      int         `json:"levels"`
						IsValid     int         `json:"isValid"`
						Comment     string      `json:"comment"`
						Sort        int         `json:"sort"`
						Type        int         `json:"type"`
						Color       interface{} `json:"color"`
						Enname      string      `json:"enname"`
						NameMerge   string      `json:"nameMerge"`
						EnnameMerge string      `json:"ennameMerge"`
						ROWNUMBER   string      `json:"ROW_NUMBER"`
					} `json:"parentAreaInfo"`
				} `json:"areaInfo"`
				CategoryInfo struct {
					Id           int    `json:"id"`
					Name         string `json:"name"`
					Type         int    `json:"type"`
					SpaceCount   int    `json:"spaceCount"`
					MinPerson    int    `json:"minPerson"`
					MaxPerson    int    `json:"maxPerson"`
					BookRule     int    `json:"bookRule"`
					RenegeRule   int    `json:"renegeRule"`
					Comment      string `json:"comment"`
					IsValid      int    `json:"isValid"`
					ROWNUMBER    string `json:"ROW_NUMBER"`
					BookRuleInfo struct {
						Id              int    `json:"id"`
						Name            string `json:"name"`
						NeedExam        int    `json:"needExam"`
						ReserveTime     int    `json:"reserveTime"`
						AttentionTime   int    `json:"attentionTime"`
						AlarmTime       int    `json:"alarmTime"`
						OpenTime        int    `json:"openTime"`
						CloseTime       int    `json:"closeTime"`
						MinTime         int    `json:"minTime"`
						MaxTime         int    `json:"maxTime"`
						ContinueTime    int    `json:"continueTime"`
						CancelTime      int    `json:"cancelTime"`
						LeaveTime       int    `json:"leaveTime"`
						BookDay         int    `json:"bookDay"`
						UpdateTime      int    `json:"updateTime"`
						SignIn          int    `json:"signIn"`
						SignInPerson    int    `json:"signInPerson"`
						SuperSignIn     int    `json:"superSignIn"`
						SignOut         int    `json:"signOut"`
						SignOutDelay    int    `json:"signOutDelay"`
						Light           int    `json:"light"`
						Power           int    `json:"power"`
						AutoSignOutTime int    `json:"autoSignOutTime"`
						SpaceOpenTime   struct {
							Date         string `json:"date"`
							TimezoneType int    `json:"timezone_type"`
							Timezone     string `json:"timezone"`
						} `json:"spaceOpenTime"`
						SpaceCloseTime struct {
							Date         string `json:"date"`
							TimezoneType int    `json:"timezone_type"`
							Timezone     string `json:"timezone"`
						} `json:"spaceCloseTime"`
						UpdateCount int    `json:"updateCount"`
						ROWNUMBER   string `json:"ROW_NUMBER"`
					} `json:"bookRuleInfo"`
					RenegeRuleInfo struct {
						Id                int    `json:"id"`
						Name              string `json:"name"`
						SignInCount       int    `json:"signInCount"`
						SignOutCount      int    `json:"signOutCount"`
						LeaveNoBackCount  int    `json:"leaveNoBackCount"`
						LeaveNoCheckCount int    `json:"leaveNoCheckCount"`
						LateCount         int    `json:"lateCount"`
						TotalCount        int    `json:"totalCount"`
						SignInHour        int    `json:"signInHour"`
						SignOutHour       int    `json:"signOutHour"`
						LeaveNoBackHour   int    `json:"leaveNoBackHour"`
						LeaveNoCheckHour  int    `json:"leaveNoCheckHour"`
						LateHour          int    `json:"lateHour"`
						TotalHour         int    `json:"totalHour"`
						IsValid           int    `json:"isValid"`
						ROWNUMBER         string `json:"ROW_NUMBER"`
					} `json:"renegeRuleInfo"`
				} `json:"categoryInfo"`
			} `json:"spaceInfo"`
			StatusName string `json:"statusName"`
			Renegeinfo struct {
				Renege int         `json:"renege"`
				Count  interface{} `json:"count"`
			} `json:"renegeinfo"`
			Starttime  string `json:"starttime"`
			Endingtime string `json:"endingtime"`
			Segment    string `json:"segment"`
			Linkurl    string `json:"linkurl"`
			PData      string `json:"p_data"`
		} `json:"list"`
		Hash struct {
			Userid string `json:"userid"`
			Hash   string `json:"hash"`
			Expire string `json:"expire"`
		} `json:"_hash_"`
	} `json:"data"`
	Checkinfo interface{} `json:"checkinfo"`
}
type AreaSegmentResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		List []struct {
			SpaceId    int    `json:"spaceId"`
			SpaceName  string `json:"spaceName"`
			Area       int    `json:"area"`
			BookTimeId int    `json:"bookTimeId"`
			BeginTime  struct {
				Date         string `json:"date"`
				TimezoneType int    `json:"timezone_type"`
				Timezone     string `json:"timezone"`
			} `json:"beginTime"`
			EndTime   string `json:"endTime"`
			Status    int    `json:"status"`
			Day       string `json:"day"`
			StartTime string `json:"startTime"`
			Id        int    `json:"id"`
		} `json:"list"`
	} `json:"data"`
}

var userInfo UserInfo

func InitRequest(info *UserInfo) {
	log.Info("初始化网络请求...")
	userInfo = *info
	client.SetCookie(&http.Cookie{
		Name:  "PHPSESSID",
		Value: info.PHPSESSID,
	})
	client.SetCookie(&http.Cookie{
		Name:  "access_token",
		Value: info.AccessToken,
	})
	client.SetCookie(&http.Cookie{
		Name:  "expire",
		Value: info.Expire,
	})
	client.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: info.UserId,
	})
	client.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	client.SetHeader("Accept", "application/json, text/javascript, */*; q=0.01")
	client.SetHeader("Accept-Language", "zh-CN,zh;q=0.9")
	client.SetHeader("Accept-Encoding", "gzip, deflate")
	client.SetHeader("Referer", "http://yuyue.lib.qlu.edu.cn/")
}

func Get(url string) (*resty.Response, error) {
	return client.R().Get(url)
}

// GetWithParams 带paras的Get请求
func GetWithParams(url string, params map[string]string) (*resty.Response, error) {
	return client.R().SetQueryParams(params).Get(url)
}

func Post(url string, data map[string]string) (*resty.Response, error) {
	return client.R().SetFormData(data).Post(url)
}

func GetAreasState() (*AreasStateResponse, error) {
	// 日期字符串
	date := time.Now().Format("2006-01-02")
	var areasStateResponse AreasStateResponse
	resp, err := Get("http://yuyue.lib.qlu.edu.cn/api.php/areas/0/date/" + date)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &areasStateResponse)
	if err != nil {
		return nil, err
	}
	return &areasStateResponse, nil
}

// GetAreaInfo 获取区域信息
func GetAreaInfo(areaId int) (*AreaInfoResponse, error) {
	var areaInfoResponse AreaInfoResponse
	resp, err := GetWithParams("http://yuyue.lib.qlu.edu.cn/api.php/spaces_old", map[string]string{
		"day":       time.Now().Format("2006-01-02"),
		"area":      strconv.Itoa(areaId),
		"segment":   "1551649",
		"startTime": time.Now().Format("15:04"),
		"endTime":   "22:00",
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &areaInfoResponse)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if areaInfoResponse.Status != 1 {
		return nil, errors.New(areaInfoResponse.Msg)
	}
	return &areaInfoResponse, nil
}

// PostReserve 提交预约请求
func PostReserve(spaceId, segment int) (*ReserveResponse, error) {
	log.Info("运行预约 : ", spaceId, "segment:", segment)
	var reserveResponse ReserveResponse
	resp, err := Post(fmt.Sprintf("http://yuyue.lib.qlu.edu.cn/api.php/spaces/%d/book", spaceId), map[string]string{
		"access_token": userInfo.AccessToken,
		"userid":       userInfo.UserId,
		"segment":      strconv.Itoa(segment),
		"type":         "1",
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &reserveResponse)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &reserveResponse, nil
}

// GetAreaSegment 获取区域可预约时间段
func GetAreaSegment(areaId int) (*AreaSegmentResponse, error) {
	var areaSegmentResponse AreaSegmentResponse
	resp, err := GetWithParams("http://yuyue.lib.qlu.edu.cn/api.php/space_time_buckets", map[string]string{
		"area": strconv.Itoa(areaId),
		"day":  ReserveTime.Format("2006-01-02"),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &areaSegmentResponse)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &areaSegmentResponse, nil
}

// CheckNetwork 检查网络连接
func CheckNetwork() bool {
	log.Info("检查网络连接...")
	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		log.Error(err, "网络连接失败")
		return false
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Timeout = 500 * time.Millisecond
	err = pinger.Run()
	if err != nil {
		log.Error(err, "网络连接失败")
		return false
	}
	if pinger.Statistics().PacketLoss > 0 {
		log.Error(errors.New("网络连接失败"))
		return false
	}
	log.Info("网络连接正常")
	return true
}
