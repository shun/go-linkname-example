package timeutil

import (
	"time"
	_ "unsafe"
)

var faketime time.Time

func SetFaketime(timestr string) {
	const FORMAT string = "2006-01-02 15:04:05"
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t1, _ := time.ParseInLocation(FORMAT, timestr, jst)

	faketime = t1
}
//go:linkname now time.Now
func now() time.Time {
	return faketime
}

