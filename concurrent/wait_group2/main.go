package main

import (
	"fmt"
	"sync"
)

func add(a, b int, done func()) {
	defer done()
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	fmt.Println(nil == nil)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add(i, 1, wg.Done)
	}
	wg.Wait()
}
