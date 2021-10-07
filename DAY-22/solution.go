package main

import "fmt"

type PS5 interface {
	StartGPUEngine()
}

type PS5WithCD struct{}

func (p PS5WithCD) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

type PS5WithDigital struct{}

func (p PS5WithDigital) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

type PS5MachinePlus struct {
	ps5Machine PS5
}

func (p *PS5MachinePlus) SetPS5Machine(ps5 PS5) {
	p.ps5Machine = ps5
}

func (p PS5MachinePlus) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("start plus plugin")
}

func main() {
	ps5MachinePlus := PS5MachinePlus{
		ps5Machine: PS5WithCD{}, // or PS5WithDigital{}
	}
	// ps5MachinePlus.SetPS5Machine(PS5WithDigital{}) // 可以在更換主機
	ps5MachinePlus.StartGPUEngine()
}
