package main

import (
	"fmt"
	"time"
)

type Like struct {
	count uint8
}

func (l *Like) Add(writerID string) {
	fmt.Printf("%s change count: %d\n", writerID, l.count+1)
	l.count++
}

func (l *Like) Show(readerID string) {
	fmt.Printf("%s read count: %d\n", readerID, l.count)
}

func AddLikes(writerID string, like *Like) {
	for i := 0; i < 100; i++ {
		like.Add(writerID)
	}
}

func ReadLikes(readerID string, like *Like) {
	for i := 0; i < 200; i++ {
		like.Show(readerID)
	}
}

func main() {
	like := new(Like)
	go AddLikes("A", like)
	go ReadLikes("B", like)
	go ReadLikes("C", like)
	go AddLikes("D", like)
	time.Sleep(10 * time.Second) //等待goroutine執行完畢
}
