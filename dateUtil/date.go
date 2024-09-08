package dateUtil

import (
	"time"
)

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func StartTimeOfMonth(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func EndTimeOfMonth(d time.Time) time.Time {
	return StartTimeOfMonth(d).AddDate(0, 1, 0).Add(-time.Second)
}

func StartTimeOfDate(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func EndTimeOfDate(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

func StartTimeOfWeek(t time.Time) time.Time {
	d := time.Time(t)
	offset := int(time.Monday - d.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location()).AddDate(0, 0, offset)
}

func EndTimeOfWeek(t time.Time) time.Time {
	d := time.Time(t)
	offset := int(7 - d.Weekday())
	if offset == 7 {
		offset = 0
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()).AddDate(0, 0, offset)
}
