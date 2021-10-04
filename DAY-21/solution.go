package main

type Controller interface {
	MethodA()
	MethodB()
	MethodC()
	MethodD()
}

type ControllerA struct{}

func (c ControllerA) MethodA() {}
func (c ControllerA) MethodB() {}
func (c ControllerA) MethodC() {}
func (c ControllerA) MethodD() {}

type ControllerB struct{}

func (c ControllerB) MethodA() {}
func (c ControllerB) MethodB() {}
func (c ControllerB) MethodC() {}
func (c ControllerB) MethodD() {}

type PS5 interface {
	Start()
	SetController()
	Play()
}

type PS5Machine struct {
	ps5Controller Controller
}

func (p PS5Machine) Start() {
	p.ps5Controller.MethodA()
	p.ps5Controller.MethodB()
}

func (p *PS5Machine) SetController(controller Controller) {
	p.ps5Controller = controller
}

func (p PS5Machine) Play() {
	p.ps5Controller.MethodC()
	p.ps5Controller.MethodD()
}

func main() {
	ps5 := PS5Machine{}
	ps5.Start()
	ps5.SetController(ControllerA{})
	ps5.Play()
	ps5.SetController(ControllerB{})
	ps5.Play()
}
