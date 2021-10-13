package main

import "fmt"

type GPU interface {
	Draw()
}

type AGPU struct{}

func (a AGPU) Draw() {
	fmt.Println("draw!")
}

type BGPU struct{}

func (b BGPU) Draw() {
	fmt.Println("draw!")
}

type CGPU struct{}

func (b CGPU) Draw() {
	fmt.Println("draw!")
}

type PS5 struct {
	gpu GPU
}

func CreatePS5(gpu GPU) PS5 {
	ps5 := PS5{
		gpu: gpu,
	}
	return ps5
}

func (p PS5) PlayGame() {
	p.gpu.Draw()
	fmt.Println("play game!")
}

func main() {
	gpu := AGPU{} // BGPU{} or CGPU{}
	ps5 := CreatePS5(&gpu)
	ps5.PlayGame()
}
