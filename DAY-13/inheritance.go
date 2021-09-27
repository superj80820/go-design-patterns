package main

import "fmt"

type PS4 struct {
	PS4Game string
}

func (p PS4) PlayPS4Game() {
	fmt.Printf("play %s", p.PS4Game)
}

type PS5 struct {
	PS4
	PS5Game string
}

func (p PS5) PlayPS5Game() {
	fmt.Printf("play %s", p.PS5Game)
}

func main() {
	ps5 := PS5{
		PS4: PS4{PS4Game: "KOF14"},
	}
	ps5.PlayPS4Game()
}
