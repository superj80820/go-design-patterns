package main

import "fmt"

type Me struct {
	Mac
}

type Mac struct{}

func (m Mac) WriteArticle() {
	fmt.Println("write article")
}

func CreateMe(mac Mac) *Me {
	return &Me{mac}
}

func (m Me) WriteIronmanArticle() {
	m.Mac.WriteArticle()
}

func main() {
	me := CreateMe(Mac{})
	me.WriteIronmanArticle()
}
