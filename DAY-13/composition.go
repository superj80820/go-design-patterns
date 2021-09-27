package main

import "fmt"

type PS5 struct {
	CPU
	Controller
	GPU
}

type CPU struct{}
type Controller struct{}
type GPU struct{}

func main() {
	ps5 := PS5{
		CPU{},
		Controller{},
		GPU{},
	}
	fmt.Println(ps5)
}
