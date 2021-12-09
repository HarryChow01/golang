package test

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Test2(t *testing.T) {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}
