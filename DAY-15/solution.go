package main

import "fmt"

type GameMachineFactory interface {
	Create() GameMachine
}

type PS5WithCDFactory struct{}

func (f *PS5WithCDFactory) Create() GameMachine {
	ps5 := &PS5WithCD{}
	ps5.AddCDMachine()
	ps5.AddCPU()
	ps5.AddGPU()
	return &PS5WithCD{}
}

type PS5WithDigitalFactory struct{}

func (f *PS5WithDigitalFactory) Create() GameMachine {
	ps5 := &PS5WithDigital{}
	ps5.AddCPU()
	ps5.AddGPU()
	return &PS5WithDigital{}
}

type GameMachine interface {
	PlayGame()
}

type PS5WithCD struct{}

func (p PS5WithCD) PlayGame() {
	fmt.Println("loading cd...play!")
}
func (p PS5WithCD) AddCDMachine() {
	fmt.Println("adding cd machine...done!")
}
func (p PS5WithCD) AddCPU() {
	fmt.Println("adding cpu...done!")
}
func (p PS5WithCD) AddGPU() {
	fmt.Println("adding gpu...done!")
}

type PS5WithDigital struct{}

func (p PS5WithDigital) PlayGame() {
	fmt.Println("loading digital file...play!")
}
func (p PS5WithDigital) AddCPU() {
	fmt.Println("adding cpu...done!")
}
func (p PS5WithDigital) AddGPU() {
	fmt.Println("adding gpu...done!")
}

func User(gameMachineFactory GameMachineFactory) {
	gameMachine := gameMachineFactory.Create()
	gameMachine.PlayGame()
}

func main() {
	User(&PS5WithCDFactory{})
}
