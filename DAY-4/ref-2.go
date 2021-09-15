package main

import (

"fmt"

"sync"

)

func produce(c chan int) {

for i := 0; i 

}

close(c)

}

func consume(id int, c chan int, wg *sync.WaitGroup) {

defer wg.Done()

for {

val, ok := 

if ok {

fmt.Printf("Consumer:%d, val:%d\n", id, val)

} else {

return

}

}

}

func main() {

var c = make(chan int, 10)

var wg sync.WaitGroup

go produce(c)

consumeCount := 2

for i:=1; i

wg.Add(1)

go consume(i, c, &wg)
    

wg.Wait()

}

