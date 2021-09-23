
package main

import "fmt"

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


func main() {
    fmt.Println("Hello, world")
    fun2()
}

