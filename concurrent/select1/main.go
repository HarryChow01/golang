package main

import (
	"fmt"
)

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	select {
	case s1 := <-ch1:
		fmt.Println(s1)
	case s2 := <-ch2:
		fmt.Println(s2)
	}

	fmt.Println("first")
	var ch3 chan string
	// nil的channel发送和接收永远阻塞
	select {
	case ch3 <- "to ch3":
		fmt.Println("to ch3")
	}
	fmt.Println("second")
}
