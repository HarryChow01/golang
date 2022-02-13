package test

import (
	"testing"
)

func TestChannel1(t *testing.T) {
	var ch1 chan int
	t.Log(ch1)
	ch2 := make(chan int)
	t.Log(ch2)
}




