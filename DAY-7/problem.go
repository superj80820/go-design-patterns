package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetServerData(serverName string) string {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second) //模擬取得server data消耗的時間
	return fmt.Sprintf("%s server data", serverName)
}

func ShowNews(news ...interface{}) {
	fmt.Println(news...)
}

func main() {
	start := time.Now()
	responseByServerA := GetServerData("A")
	responseByServerB := GetServerData("B")
	responseByServerC := GetServerData("C")
	ShowNews(responseByServerA, responseByServerB, responseByServerC)
	fmt.Printf("cost %s", time.Since(start))
}
