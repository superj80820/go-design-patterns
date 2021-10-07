package main

import "fmt"

type CPUFacade interface {
	Work()
}

type StructA struct{}

func (s StructA) DoAction() {}

type StructB struct{}

func (s StructB) DoAction() {}

type CPU struct{}

func (c CPU) Work() {
	strcutA := StructA{}
	strcutB := StructB{}
	strcutA.DoAction()
	strcutB.DoAction()
}

type PS5 struct {
	cpu CPUFacade
}

func (p PS5) Start() {
	p.cpu.Work()
	fmt.Println("start ps5...done!")
}

func main() {
	ps5 := PS5{cpu: CPU{}}
	ps5.Start()
}
