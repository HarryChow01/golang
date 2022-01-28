package test

import (
    "fmt"
    "testing"
)

func TestArray1(t *testing.T) {
    var arr1 [12]int
    for index, value := range arr1 {
        fmt.Printf("index: %+v, value: %+v\n", index, value)
    }
    arr2 := [12]int{10, 20, 30}
    for index, value := range arr2 {
        fmt.Printf("index: %+v, value: %+v\n", index, value)
    }
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
