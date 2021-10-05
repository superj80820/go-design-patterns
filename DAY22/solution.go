package main

import "fmt"

type PS5 interface {
	StartGPUEngine()
}

type PS5Machine struct{}

func (p PS5Machine) StartGPUEngine() {
	fmt.Println("start normal gpu engine")
}

type PS5MachinePlus struct {
	ps5Machine PS5Machine
}

func (p PS5MachinePlus) StartGPUEngine() {
	p.ps5Machine.StartGPUEngine()
	fmt.Println("start plus gpu engine")
}

func main() {
	ps5MachinePlus := PS5MachinePlus{
		ps5Machine: PS5Machine{},
	}
	ps5MachinePlus.StartGPUEngine()
}
