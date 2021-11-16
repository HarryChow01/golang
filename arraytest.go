package main

import "fmt"

func ArrayTest() {
    arr := []int64{}
    // var arr []int64
    // arr = append(arr, 5)
    fmt.Printf("len(arr): %+v", len(arr))
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

func main() {
    // ArrayTest()
    var intValue int32 = 10
    fmt.Printf("main intValue: %v\n", intValue)
    testPointer2(&intValue)
    fmt.Printf("main intValue: %v\n", intValue)
}
