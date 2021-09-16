package main

import (
	"fmt"
	"time"
)

func PushNews(news string, startTime time.Time) time.Time {
	time.Sleep(time.Duration(3 * time.Second)) //模擬推播運行的時間
	fmt.Printf("%s Cost %s\n", news, time.Since(startTime))
	return time.Now()
}

func main() {
	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
	}
	for _, news := range allNews {
		go PushNews(news, start)
	}
	time.Sleep(10 * time.Second) //等待goroutine執行完畢
}
