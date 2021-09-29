package factory

import "fmt"

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

func main() {
	ps5 := &PS5WithCD{}
	ps5.AddCDMachine()
	ps5.AddCPU()
	ps5.AddGPU()

	// or
	// ps5 := &PS5WithDigital{}
	// ps5.AddCPU()
	// ps5.AddGPU()

	ps5.PlayGame()
}
