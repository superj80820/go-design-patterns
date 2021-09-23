package main

import (
	"fmt"
	"time"
)

type SendInfo struct {
	NewsName   string
	FinishTime time.Time
}

func Worker(id int, jobs <-chan string, results chan<- SendInfo, startTime time.Time) {
	time.Sleep(time.Duration(10 * time.Second)) //模擬與某service建立socket的時間
	for job := range jobs {
		time.Sleep(time.Duration(3 * time.Second)) //模擬推播運行的時間
		fmt.Printf("%s cost %s by worker %d\n", job, time.Since(startTime), id)
		results <- SendInfo{job, time.Now()}
	}
}

func main() {
	start := time.Now()
	allNews := []string{
		"中秋節來了",
		"記得",
		"不要戶外烤肉～",
		"也不要吃太撐",
	}
	jobs := make(chan string, len(allNews))
	results := make(chan SendInfo, len(allNews))

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results, start)
	}

	for _, news := range allNews {
		jobs <- news
	}

	// do something

	for r := 1; r <= len(allNews); r++ {
		result := <-results
		fmt.Printf("news %s is sent at %s\n", result.NewsName, result.FinishTime)
	}
}
