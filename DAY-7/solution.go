package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Producer(serverNames ...string) <-chan string {
	producerCh := make(chan string, len(serverNames))
	go func() {
		defer close(producerCh)
		for _, serverName := range serverNames {
			producerCh <- serverName
		}
	}()
	return producerCh
}

func Task(producerCh <-chan string) <-chan string {
	taskCh := make(chan string)
	go func() {
		defer close(taskCh)
		for serverName := range producerCh {
			taskCh <- GetServerData(serverName)
		}
	}()
	return taskCh
}

func Consumer(taskChs ...<-chan string) <-chan string {
	consumerCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(taskChs))
	go func() {
		wg.Wait()
		close(consumerCh)
	}()

	for _, task := range taskChs {
		go func(task <-chan string) {
			defer wg.Done()
			for new := range task {
				consumerCh <- new
			}
		}(task)
	}

	return consumerCh
}

func GetServerData(serverName string) string {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second) //模擬取得server data消耗的時間
	return fmt.Sprintf("%s server data", serverName)
}

func ShowNews(news ...interface{}) {
	fmt.Println(news...)
}

func main() {
	start := time.Now()
	producerCh := Producer("A", "B", "C")

	task1 := Task(producerCh)
	task2 := Task(producerCh)
	task3 := Task(producerCh)

	consumerCh := Consumer(task1, task2, task3)

	for news := range consumerCh {
		ShowNews(news)
	}
	fmt.Printf("cost %s", time.Since(start))
}
