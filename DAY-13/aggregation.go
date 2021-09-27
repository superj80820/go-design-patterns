package main

import "fmt"

type NintendoSwitch struct {
	Controller string
}

func (n *NintendoSwitch) SetController(controller string) {
	n.Controller = controller
}

func (n NintendoSwitch) PlayGame() {
	if n.Controller != "" {
		fmt.Printf("use %s to play game", n.Controller)
	} else {
		fmt.Println("use default controller to play game")
	}
}

func main() {
	nintendoSwitch := NintendoSwitch{}
	nintendoSwitch.SetController("drum")
	nintendoSwitch.PlayGame()
}
