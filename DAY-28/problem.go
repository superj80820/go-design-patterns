package main

import "fmt"

type CPU struct{}

func (CPU) ADoSomething() {
	fmt.Println("a do something")
}
func (CPU) BDoSomething() {
	fmt.Println("b do something")
}
func (CPU) CDoSomething() {
	fmt.Println("c do something")
}

type PS5 struct {
	cpu CPU
}

func (p PS5) ACommand() {
	p.cpu.ADoSomething()
	p.cpu.CDoSomething()
}
func (p PS5) BCommand() {
	p.cpu.ADoSomething()
	p.cpu.BDoSomething()
}
func main() {
	cpu := CPU{}
	ps5 := PS5{cpu}
	ps5.ACommand()
	ps5.BCommand()
}
