package main

import (
	"log"
	"time"
	"github.com/shun/go-linkname-example/timeutil"
	_ "unsafe"
)

//go:linkname now timeutil.TimeNow
func now() time.Time {
	return time.Date(2022, 12, 1, 1, 0, 0, 0, time.Local)
}

func main() {
	now := timeutil.TimeNow()
	log.Println(now)
}
