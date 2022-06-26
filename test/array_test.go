package test

import (
	"fmt"
	"testing"
)

func TestArray1(t *testing.T) {
	// 声明数组，不提供显式初始化
	var arr1 [5]int // 数组声明即定义，自动分配内存并初始化
	// fmt.Printf("arr1 == nil: %t\n", arr1 == nil), error: 不能和nil比较
	fmt.Printf("len(arr1): %d\n", len(arr1)) // 数组声明即定义，自动分配内存并初始化
	fmt.Println(arr1)

	// 左侧 [5]int 可以省略
	// var arr2 [5]int = [5]int{1, 2, 3, 4, 5} // ok
	var arr2 = [5]int{1, 2, 3, 4, 5}
	// arr3 := [5]int{1,2,3,4,5} // 同上
	fmt.Printf("len(arr2): %d\n", len(arr2)) // 数组声明即定义，自动分配内存并初始化
	fmt.Println(arr2)

	arr4 := [...]int{1, 2, 3, 4, 5}
	for index, value := range arr4 {
		fmt.Printf("index = %+v, value = %+v\n", index, value)
	}

	// 数组部分赋值
	arr5 := [4]int{10, 20, 30}
	for index, value := range arr5 {
		fmt.Printf("index = %+v, value = %+v\n", index, value)
	}
}

func TestArray2(t *testing.T) {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5}
	// var arr3 = [3]int{1, 2, 3}
	// fmt.Println("arr1 == arr3: ", arr1 == arr3)	// error，类型不同不能比较；

	fmt.Println("arr1 == arr2: ", arr1 == arr2)
	fmt.Println("arr1 != arr2: ", arr1 != arr2)

	// fmt.Println("arr1 < arr2: ", arr1 < arr2)  // error，不支持<和>比较；
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func TestArray3(t *testing.T) {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}

func testPointer1(intValue int32) {
	fmt.Printf("func intValue: %v\n", intValue)
	intValue = 15
	fmt.Printf("func intValue: %v\n", intValue)
}

func testPointer2(intValue *int32) {
	fmt.Printf("func intValue: %v\n", *intValue)
	*intValue = 15
	fmt.Printf("func intValue: %v\n", *intValue)
}

func Test222(t *testing.T) {
	// ArrayTest()
	var intValue int32 = 10
	fmt.Printf("main intValue: %v\n", intValue)
	testPointer2(&intValue)
	fmt.Printf("main intValue: %v\n", intValue)
}
