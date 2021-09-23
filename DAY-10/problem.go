package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("bye bye")
				os.Exit(1)
			}
		}
	}()

	// do something

	select {}
}
