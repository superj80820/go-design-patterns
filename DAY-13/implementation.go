package main

import "fmt"

type Contestant interface {
	Sing()
}

type Mike struct{}

func (m Mike) Sing() { fmt.Println("mike singing") }

type Kevin struct{}

func (m Kevin) Sing() { fmt.Println("kevin singing") }

type York struct{}

func (m York) Sing() { fmt.Println("york singing") }

func Show(contestant Contestant) {
	contestant.Sing()
}

func main() {
	Show(Mike{})
	Show(Kevin{})
	Show(York{})
}
