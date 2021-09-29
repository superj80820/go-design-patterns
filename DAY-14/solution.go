package factory

import "fmt"

type PS5 interface {
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

func CreatePS5(style string) PS5 {
	switch style {
	case "PS5WithCD":
		ps5 := &PS5WithCD{}
		ps5.AddCDMachine()
		ps5.AddCPU()
		ps5.AddGPU()
		return ps5
	case "PS5WithDigital":
		ps5 := &PS5WithDigital{}
		ps5.AddCPU()
		ps5.AddGPU()
		return &PS5WithDigital{}
	}
	return nil
}

func main() {
	ps5 := CreatePS5("PS5WithCD")
	ps5.PlayGame()
}
