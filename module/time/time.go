package main

import (
	"fmt"
	"time"
)

func PrintNow() {
	p := fmt.Println

	// 得到当前时间。
	now := time.Now()
	p(now)
	return now

}
