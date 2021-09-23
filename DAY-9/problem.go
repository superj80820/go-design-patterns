package main

import (
	"fmt"
	"time"
)

func PushNews(news string, startTime time.Time) <-chan time.Time {
	newsCh := make(chan time.Time)
	go func() {
		time.Sleep(time.Duration(10 * time.Second)) //模擬與某service建立socket的時間
		time.Sleep(time.Duration(3 * time.Second))  //模擬推播運行的時間
		fmt.Printf("%s cost %s\n", news, time.Since(startTime))
		newsCh <- time.Now()
	}()
	return newsCh
}

func main() {
	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
		"也不要吃太撐",
	}
	newsChs := []<-chan time.Time{}
	for _, news := range allNews {
		newsChs = append(newsChs, PushNews(news, start))
	}

	// do something

	for index, newsCh := range newsChs {
		fmt.Printf("news %d is sent at %s\n", index, <-newsCh)
	}
}
