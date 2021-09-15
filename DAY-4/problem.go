package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MessageBoard struct {
	messages []string
}

func (m *MessageBoard) SendMessage(message string) {
	m.messages = append(m.messages, message)
}

func (m *MessageBoard) Show() {
	for {
		fmt.Println(m.messages[0])
		m.messages = m.messages[1:]
	}
}

func main() {
	messageBoard := new(MessageBoard)
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模擬各個client發送的隨機處理時間
			messageBoard.SendMessage("cool")
		}
	}()
	go func() {
		messageBoard.Show()
	}()
	time.Sleep(10 * time.Second) //等待goroutine執行完畢
}
