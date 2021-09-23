package main

import (
	"fmt"
	"time"
)

type UserInfo struct {
	ID   uint32
	Name string
}

var userInfos = []UserInfo{
	{
		1,
		"糖糖",
	},
	{
		2,
		"鹽鹽",
	},
	{
		3,
		"乖乖",
	},
}

func UberProducer(job chan<- UserInfo, i int) {
	go UberConsumer(userInfos[i], i)
}

func UberConsumer(userInfo UserInfo, id int) {
	fmt.Printf("uber consumer %d get %s user\n", id, userInfo.Name)
}

func main() {
	job := make(chan UserInfo)
	UberProducerCount := len(userInfos)

	for i := 0; i < UberProducerCount; i++ {
		go UberProducer(job, i)
	}

	time.Sleep(10 * time.Second) //等待goroutine執行完畢
}
