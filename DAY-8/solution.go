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
	job <- userInfos[i]
}

func UberConsumer(job <-chan UserInfo, id int) {
	for userInfo := range job {
		fmt.Printf("uber consumer %d get %s user\n", id, userInfo.Name)
	}
}

func main() {
	job := make(chan UserInfo)
	UberProducerCount := len(userInfos)
	UberConsumerCount := 5

	for i := 0; i < UberProducerCount; i++ {
		go UberProducer(job, i)
	}

	for i := 0; i < UberConsumerCount; i++ {
		go UberConsumer(job, i)
	}

	time.Sleep(10 * time.Second)
}
