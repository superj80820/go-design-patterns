package main

import "fmt"

// 定義抽象工廠
type GameRoomFactory interface {
	GameMachineFactory() GameMachine
	GameFactory() Game
}

// 定義抽象產品
type Game interface {
	Start()
}

type GameMachine interface {
	PlayGame()
}

// 實作產品
type PS5 struct{}

func (p PS5) PlayGame() {
	fmt.Println("loading cd...play!")
}
func (p PS5) addCDMachine() {
	fmt.Println("adding cd machine...done!")
}
func (p PS5) addCPU() {
	fmt.Println("adding cpu...done!")
}
func (p PS5) addGPU() {
	fmt.Println("adding gpu...done!")
}

type GameFinalFantasy struct{}

func (s *GameFinalFantasy) build() {
	fmt.Println("build game...done!")
}
func (s *GameFinalFantasy) Start() {
	fmt.Println("start game...done!")
}

type Switch struct{}

func (s Switch) PlayGame() {
	fmt.Println("loading cd...play!")
}
func (s Switch) addCDMachine() {
	fmt.Println("adding cd machine...done!")
}
func (s Switch) addCPU() {
	fmt.Println("adding cpu...done!")
}
func (s Switch) addGPU() {
	fmt.Println("adding gpu...done!")
}

type GameMario struct{}

func (s *GameMario) build() {
	fmt.Println("build game...done!")
}
func (s *GameMario) Start() {
	fmt.Println("start game...done!")
}

// 實作工廠
type SonyFactory struct{}

func (f *SonyFactory) GameMachineFactory() GameMachine {
	ps5 := &PS5{}
	ps5.addCDMachine()
	ps5.addCPU()
	ps5.addGPU()
	return &PS5{}
}

func (f *SonyFactory) GameFactory() Game {
	game := &GameFinalFantasy{}
	game.build()
	return game
}

type NintendoFactory struct{}

func (n *NintendoFactory) GameMachineFactory() GameMachine {
	switchMachine := &Switch{}
	switchMachine.addCDMachine()
	switchMachine.addCPU()
	switchMachine.addGPU()
	return &PS5{}
}

func (n *NintendoFactory) GameFactory() Game {
	game := &GameMario{}
	game.build()
	return game
}

func User(gameHomeFactory GameRoomFactory) {
	gameMachine := gameHomeFactory.GameMachineFactory()
	game := gameHomeFactory.GameFactory()
	game.Start()
	gameMachine.PlayGame()
}

func main() {
	User(&SonyFactory{})
}
