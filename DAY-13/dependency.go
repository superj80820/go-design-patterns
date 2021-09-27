package main

import "fmt"

type Me struct{}

type Whiteboard struct{}

func (w Whiteboard) DrawOnWhiteboard() {
	fmt.Println("write")
}

func (m Me) Discuss(whiteboard Whiteboard) {
	whiteboard.DrawOnWhiteboard()
}

func main() {
	me := Me{}
	me.Discuss(Whiteboard{})
}
