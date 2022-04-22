package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect1(t *testing.T) {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		t := <-tick
		fmt.Println(t)
	}
	t.Log("end")
}
