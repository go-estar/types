package dateUtil

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	date, _ := time.Parse("2006-01-02 15:04:05", "2020-12-27 00:00:00")
	//date= time.Now()
	t.Log(int(date.Weekday()))
	t.Log(date.Day())
}
