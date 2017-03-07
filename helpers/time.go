package helpers

import (
	"strings"
	"time"
)

func GetDateUnix(date ...string) (int64, int64) {
	var selectDate string
	if len(date) > 0 {
		selectDate = date[0]
	} else {
		selectDate = time.Now().Format("2006-01-02")
	}

	selectDate = strings.TrimSpace(selectDate)
	startTime := selectDate + " 00:00:00"
	endTime := selectDate + " 23:59:59"
	return StrToBeijingTime(startTime).Unix(), StrToBeijingTime(endTime).Unix()
}

func StrToBeijingTime(strDateTime string) time.Time {
	setLocation, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", strDateTime, setLocation)
	return t
}

func DateHourMap(date string) map[int][]int64 {
	startTime, _ := GetDateUnix(date)
	var hoursMap map[int][]int64

	for i := 0; i < 24; i++ {
		hourStartTime, hourEndTime := int64(i*60)+startTime, int64((i+1)*60-1)*startTime
		hoursMap[i] = []int64{hourStartTime, hourEndTime}
	}

	return hoursMap
}
