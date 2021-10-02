package main

import (
	"fmt"
	"time"
)

type PS5 struct {
	Version string
	CPU     []string
	GPU     map[string]string
}

func (_ PS5) Modeling() {
	time.Sleep(time.Second) // 模擬耗時
}

func (_ PS5) LoadTest() {
	time.Sleep(time.Second) // 模擬耗時
}

func (_ PS5) Analysis() {
	time.Sleep(time.Second) // 模擬耗時
}

func (p *PS5) Clone() *PS5 {
	ps5 := PS5{}
	ps5 = *p
	cpu := make([]string, len(p.CPU))
	copy(cpu, p.CPU)
	ps5.CPU = cpu
	gpu := make(map[string]string)
	for k, v := range p.GPU {
		gpu[k] = v
	}
	ps5.GPU = gpu
	return &ps5
}

func CreatePrototypePS5() *PS5 {
	prototypePS5 := PS5{
		Version: "Prototype",
		CPU:     []string{"原型CPU"},
		GPU:     make(map[string]string),
	}
	prototypePS5.GPU["GPU"] = "原型GPU"
	prototypePS5.Modeling()
	prototypePS5.LoadTest()
	prototypePS5.Analysis()
	return &prototypePS5
}

func main() {
	prototypePS5 := CreatePrototypePS5()
	ps5 := prototypePS5.Clone()
	ps5.Version = "Version-1"
	ps5.CPU[0] = "量產CPU"
	ps5.GPU["GPU"] = "量產GPU2"
	fmt.Println(prototypePS5)
	fmt.Println(ps5)
}
