package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Factory struct {
	active   bool
	capacity int
	queue    []int
}

func (f *Factory) produce(c *sync.Cond) {
	for f.active {
		c.L.Lock()
		for len(f.queue) == f.capacity && f.active {
			c.Wait()
		}
		if f.active {
			n := rand.Intn(100)
			f.queue = append(f.queue, n)
			fmt.Printf("Generating %d\n", n)
			c.Broadcast()
		}
		c.L.Unlock()
	}

}

func (f *Factory) consume(id int, c *sync.Cond, wg *sync.WaitGroup) {
	for {
		c.L.Lock()
		for len(f.queue) == 0 && f.active {
			c.Wait()
		}
		if len(f.queue) > 0 {
			n := f.queue[0]
			f.queue = f.queue[1:]
			fmt.Printf("Consumer: %d, Received: %d\n", id, n)
			c.Broadcast()
		} else if !f.active {
			c.L.Unlock()
			wg.Done()
			break
		}
		c.L.Unlock()
	}
}

func main() {

    factory := Factory
    
    c := sync.NewCond(&sync.Mutex{})
    
    var consumerCount int = 2
    
    var wg sync.WaitGroup
    
    go factory.produce(c)
    
    for i:=1; i 
    
    wg.Add(1)
    
    go factory.consume(i, c, &wg)
    
    }
    
    time.Sleep(20*time.Millisecond)
    
    factory.active = false
    
    c.Broadcast()
    
    wg.Wait()
    
    }
