package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("bye bye")
				os.Exit(1)
			}
		}
	}()

	// do something

	select {}
}
