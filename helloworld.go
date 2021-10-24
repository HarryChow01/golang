
package main

import (
    "fmt"
    "reflect"
)

func fun1() {
    sum := 0
    for i := 0; i <= 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}

func fun2() {
    sum := 0
    for i := 0; i <= 10; i++ {
        sum += i
    }
    fmt.Println("sum:", sum)
}

func fun3() {
    arr := []int64{}
    // arr = append(arr, 5)
    fmt.Printf(reflect.TypeOf(arr))
    fmt.Printf("len(arr): %+v", len(arr))
}

func main() {
    fmt.Println("Hello, world")
    fun3()
}
