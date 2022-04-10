package test

import (
	"fmt"
	"testing"
)

func TestChannel1(t *testing.T) {
	var ch1 chan int
	t.Log(ch1)

	ch2 := make(chan int)
	t.Log(ch2)

	ch3 := ch2
	t.Log(ch2 == ch3)

	ch4 := make(chan int, 3)
	t.Log(ch4)
}

func TestChannel2(t *testing.T) {
	// var ch1 chan int		// nil
	ch1 := make(chan int)
	var xx = 5
	var yy = 1

	// 启动一个协程
	go func() {
		ch1 <- xx // 发送操作推荐右分开写
	}()

	yy = <-ch1 // 接收操作，推荐右连接写
	// <-ch1      // 空接收语句，结果丢弃
	t.Log(yy)
	close(ch1)
}

func TestChannel3(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
