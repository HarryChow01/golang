package test

import (
    "fmt"
    "testing"
)

func TestArray1(t *testing.T) {
    var arr1 [4]int
    for index, value := range arr1 {
        fmt.Printf("index: %+v, value: %+v\n", index, value)
    }
    arr2 := [4]int{10, 20, 30}
    for index, value := range arr2 {
        fmt.Printf("index: %+v, value: %+v\n", index, value)
    }
    arr3 := [...]int{1, 2}
    for index, value := range arr3 {
    fmt.Printf("index: %+v, value: %+v\n", index, value)
    }
}

type Currency int

const (
    USD Currency = iota // 美元
    EUR                 // 欧元
    GBP                 // 英镑
    RMB                 // 人民币
)

func TestArray2(t *testing.T) {
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
