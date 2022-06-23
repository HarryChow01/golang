package main

import "fmt"

var ch1 chan int
var ch2 chan int
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

// 1. select从上到下,从左到右，计算每一个case的表达式
// 2. nil的channel永远不会返回
// 3. 所有case不命中，则走default

func main() {
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)

	return chs[i]
}
