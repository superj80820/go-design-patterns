package main

import "fmt"

type Game struct {
	Name      string
	Type      string
	GraphType string
	AudioType string
}

type PS5 struct {
	middlewares []func(game Game)
}

func (p *PS5) AddMiddleware(middleware func(game Game)) *PS5 {
	p.middlewares = append(p.middlewares, middleware)
	return p
}

func (p PS5) PlayGame(game Game) {
	for _, middleware := range p.middlewares {
		middleware(game)
	}
	fmt.Printf("play %s", game.Name)
}

func GameMiddleware(game Game) {
	if game.Type == "3D遊戲" {
		fmt.Println("3D模式")
	}
}

func GraphMiddleware(game Game) {
	if game.GraphType == "高效能顯示" {
		fmt.Println("加強顯示晶片")
	}
}

func AudioMiddleware(game Game) {
	if game.AudioType == "環繞音效" {
		fmt.Println("環繞音效模式")
	}
}

func main() {
	ps5 := PS5{}
	ps5.
		AddMiddleware(GameMiddleware).
		AddMiddleware(GraphMiddleware).
		AddMiddleware(AudioMiddleware)
	ps5.PlayGame(
		Game{
			Name:      "最終幻想",
			Type:      "3D遊戲",
			GraphType: "高效能顯示",
			AudioType: "環繞音效",
		},
	)
}
