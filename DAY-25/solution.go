package main

import "fmt"

var pockerCards = map[int]*Card{
	1: {
		Name:  "A",
		Color: "紅",
	},
	2: {
		Name:  "A",
		Color: "黑",
	},
	// 其他卡牌
}

type Card struct {
	Name  string
	Color string
}

type PockerGame struct {
	Cards map[int]*Card
}

func NewPockerGame() *PockerGame {
	board := &PockerGame{Cards: map[int]*Card{}}
	for id := range pockerCards {
		board.Cards[id] = pockerCards[id]
	}
	return board
}

func main() {
	game1 := NewPockerGame()
	game2 := NewPockerGame()
	fmt.Println(game1.Cards[1] == game2.Cards[1])
}
