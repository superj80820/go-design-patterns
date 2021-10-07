package main

import "fmt"

type CPU interface {
	Run()
}

type SingleCPU struct{}

func (SingleCPU) Run() {
	fmt.Println("run cpu")
}

type MultiCPUs struct {
	subCPUs []CPU
}

func (d MultiCPUs) Run() {
	for _, cpu := range d.subCPUs {
		cpu.Run()
	}
}

func (m *MultiCPUs) AddSubCPU(cpu CPU) {
	m.subCPUs = append(m.subCPUs, cpu)
}

func PS5Start(cpu CPU) {
	cpu.Run()
}

func main() {
	singleCPU1 := SingleCPU{}
	PS5Start(singleCPU1)

	singleCPU2 := SingleCPU{}
	PS5Start(singleCPU2)

	multiCPUs := MultiCPUs{}
	multiCPUs.AddSubCPU(&singleCPU1)
	multiCPUs.AddSubCPU(&singleCPU2)
	PS5Start(multiCPUs)
}
