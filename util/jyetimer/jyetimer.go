package jyetimer

import (
	"fmt"
	"time"
)

// YYYY-MM-DD -> unix
// YYYY-MM-DD hh:mm:ss -> unix
func TimeString2UnixInt64(t string) (data int64, err error) {
	if len(t) != 10 && len(t) != 19 {
		return 0, fmt.Errorf("can not parse %s", t)
	}

	timeTemplate := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, err := time.ParseInLocation(timeTemplate, t, loc)

	return timer.Unix(), err
}

// YYYY-MM-DD -> unix
// YYYY-MM-DD hh:mm:ss -> unix
func TimeString2UnixString(t string) (data string, err error) {
	if len(t) != 10 && len(t) != 19 {
		return "0", fmt.Errorf("can not parse %s", t)
	}
	timeTemplate := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, err := time.ParseInLocation(timeTemplate, t, loc)

	return fmt.Sprintf("%d", timer.Unix()), err
}
