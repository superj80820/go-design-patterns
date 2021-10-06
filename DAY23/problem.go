package main

import "fmt"

type StructA struct{}

func (s StructA) DoAction() {}

type StructB struct{}

func (s StructB) DoAction() {}

type PS5 struct {
}

func (p PS5) Start() {
	strcutA := StructA{}
	strcutB := StructB{}
	strcutA.DoAction()
	strcutB.DoAction()
	fmt.Println("start ps5...done!")
}

func main() {
	ps5 := PS5{}
	ps5.Start()
}
