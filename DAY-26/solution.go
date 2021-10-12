package main

import (
	"fmt"
	"time"
)

type PS5 interface {
	PlayGame()
}

type PS5Machine struct{}

func NewPS5Machine() *PS5Machine {
	return &PS5Machine{}
}

func (u *PS5Machine) PlayGame() {
	fmt.Println("play game")
}

type PS5MachineProxy struct {
	ps5 *PS5Machine
}

func NewPS5MachineProxy(ps5 *PS5Machine) *PS5MachineProxy {
	return &PS5MachineProxy{
		ps5: ps5,
	}
}

func (p *PS5MachineProxy) PlayGame() {
	start := time.Now()
	p.ps5.PlayGame()
	fmt.Printf("play game cost time: %s", time.Since(start))
}

func main() {
	ps5 := NewPS5MachineProxy(NewPS5Machine())
	ps5.PlayGame()
}
