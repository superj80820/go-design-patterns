package main

import (
	"fmt"
	"time"
)

func PushNews(news string, startTime time.Time) {
	time.Sleep(time.Duration(3 * time.Second)) //模擬推播運行的時間
	fmt.Printf("%s cost %s\n", news, time.Since(startTime))
}

func main() {
	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
	}
	for _, news := range allNews {
		PushNews(news, start)
	}
	fmt.Printf("cost %s", time.Since(start))
}
