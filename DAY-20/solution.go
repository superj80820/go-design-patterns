package main

import "fmt"

type SignalHandler interface {
	ClickButton()
}

type PS5 struct{}

func (_ PS5) ClickPS5Button() {
	fmt.Println("click ps5 button")
}

type PS5Adapter struct {
	ps5Machine *PS5
}

func (p PS5Adapter) ClickButton() {
	p.ps5Machine.ClickPS5Button()
}

type Switch struct{}

func (_ Switch) ClickSwitchButton() {
	fmt.Println("click switch button")
}

type SwitchAdapter struct {
	switchMachine *Switch
}

func (p SwitchAdapter) ClickButton() {
	p.switchMachine.ClickSwitchButton()
}

func CreateSignalHandler(platform string) SignalHandler {
	var signalHandler SignalHandler
	switch platform {
	case "ps5":
		signalHandler = PS5Adapter{
			ps5Machine: &PS5{},
		}
	case "switch":
		signalHandler = SwitchAdapter{
			switchMachine: &Switch{},
		}
	}
	return signalHandler
}

func main() {
	signalHandler := CreateSignalHandler("ps5")
	signalHandler.ClickButton()
}
