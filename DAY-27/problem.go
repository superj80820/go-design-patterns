package main

import "fmt"

type Game struct {
	Name      string
	Type      string
	GraphType string
	AudioType string
}

type PS5 struct{}

func (PS5) PlayGame(game Game) {
	if game.Type == "3D遊戲" {
		fmt.Println("3D模式")
	}
	if game.GraphType == "高效能顯示" {
		fmt.Println("加強顯示晶片")
	}
	if game.AudioType == "環繞音效" {
		fmt.Println("環繞音效模式")
	}
	fmt.Printf("play %s", game.Name)
}

func main() {
	ps5 := PS5{}
	ps5.PlayGame(
		Game{
			Name:      "最終幻想",
			Type:      "3D遊戲",
			GraphType: "高效能顯示",
			AudioType: "環繞音效",
		},
	)
}
