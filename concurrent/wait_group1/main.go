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
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			c := x + 1
			fmt.Printf("xxx  %d + %d = %d\n", x, 1, c)
		}(i)
	}
	wg.Wait()
}
