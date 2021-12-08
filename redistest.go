
package main

import (
    "fmt"
    "path"
    "runtime"
)

func getCurrentAbPathByCaller() string {
    var abPath string
    _, filename, _, ok := runtime.Caller(0)
    fmt.Printf("cur file: %s\n", filename)
    if ok {
        abPath = path.Dir(filename)
    }
    fmt.Printf("absolute path: %s\n", abPath)
    return abPath
}


func redisTest1() {

}


func main() {
    fmt.Println("start:")
    getCurrentAbPathByCaller()
}
