package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case ch <- rand.Intn(5): // Create and send random number into channel
			case <-done: // If receive signal on done channel - Return
				return
			default:
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Rand Number = ", <-ch) // Print number received on standard output
		}
		done <- struct{}{} // Send Terminate Signal and return
		return
	}()
	<-done // Exit Main when Terminate Signal received
}
