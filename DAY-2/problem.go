package main

import (
	"fmt"
	"time"
)

type Like struct {
	Count uint16
}

func (l *Like) Add(writerID string) {
	l.Count++
}

func AddLikes(writerID string, like *Like) {
	for i := 0; i < 10000; i++ {
		like.Add(writerID)
	}
}

func main() {
	like := new(Like)
	go AddLikes("A", like)
	go AddLikes("B", like)
	go AddLikes("C", like)
	go AddLikes("D", like)
	time.Sleep(1 * time.Second) //等待goroutine執行完畢
	fmt.Println(like.Count)
}
