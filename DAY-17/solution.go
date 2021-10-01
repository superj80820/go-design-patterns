package main

import (
	"fmt"
	"strings"
)

type PS5Builder interface {
	SetController(isBuy bool) PS5Builder
	SetBluetoothHeadphones(isBuy bool) PS5Builder
	Build() *PS5
}

type PS5 struct {
	cpu                 string
	gpu                 string
	controller          string
	bluetoothHeadphones string
}

func CreatePS5() *PS5 {
	return &PS5{
		cpu: "cpu",
		gpu: "gpu",
	}
}

func (p PS5) PlayGame() {
	var (
		accessories     []string
		withAccessories string
	)
	if p.controller != "" {
		accessories = append(accessories, p.controller)
	}
	if p.bluetoothHeadphones != "" {
		accessories = append(accessories, p.bluetoothHeadphones)
	}
	if len(accessories) != 0 {
		withAccessories = " with " + strings.Join(accessories, ", ")
	}
	fmt.Printf("loading...play%s!\n", withAccessories)
}

type PS5Director struct {
	Builder PS5Builder
}

func CreatePS5Director(concretePS5Builder *ConcretePS5Builder) *PS5Director {
	return &PS5Director{
		Builder: concretePS5Builder,
	}
}

func (p PS5Director) Construct() *PS5 {
	return p.Builder.
		SetController(true).
		Build()
}

type ConcretePS5Builder struct {
	ps5 *PS5
}

func CreateConcretePS5Builder() *ConcretePS5Builder {
	return &ConcretePS5Builder{
		ps5: CreatePS5(),
	}
}

func (p *ConcretePS5Builder) SetController(isBuy bool) PS5Builder {
	if isBuy {
		p.ps5.controller = "FightingStick"
	}
	return p
}

func (p *ConcretePS5Builder) SetBluetoothHeadphones(isBuy bool) PS5Builder {
	if isBuy {
		p.ps5.bluetoothHeadphones = "BluetoothHeadphones"
	}
	return p
}

func (p *ConcretePS5Builder) Build() *PS5 {
	return p.ps5
}

func main() {
	concretePS5Builder := CreateConcretePS5Builder()
	ps5Director := CreatePS5Director(concretePS5Builder)
	ps5 := ps5Director.Construct()

	ps5.PlayGame()
}
