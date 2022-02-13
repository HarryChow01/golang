package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine1(t *testing.T) {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	t.Logf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `@#$%` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
