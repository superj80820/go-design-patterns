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

type Command interface {
	Execute()
}

type ACommand struct {
	cpu CPU
}

func (a ACommand) Execute() {
	a.cpu.ADoSomething()
	a.cpu.CDoSomething()
}

type BCommand struct {
	cpu CPU
}

func (b BCommand) Execute() {
	b.cpu.ADoSomething()
	b.cpu.BDoSomething()
}

type PS5 struct {
	commands map[string]Command
}

func (p *PS5) SetCommand(name string, command Command) {
	p.commands[name] = command
}

func (p *PS5) DoCommand(name string) {
	p.commands[name].Execute()
}

func main() {
	cpu := CPU{}
	aCommand := ACommand{cpu}
	bCommand := BCommand{cpu}
	ps5 := PS5{make(map[string]Command)}
	ps5.SetCommand("a", aCommand)
	ps5.SetCommand("b", bCommand)
	ps5.DoCommand("a")
	ps5.DoCommand("b")
}
