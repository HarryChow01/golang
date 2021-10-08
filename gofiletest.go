package main

import (
	"embed"
	"fmt"
)

//go:embed web/config.json
var s string

//go:embed web/config.json
var b []byte

//go:embed web/config.json
var fs embed.FS
func main() {
	fmt.Println(s)
	fmt.Println(string(b))
	data, err := fs.ReadFile("config.json")
	fmt.Println(err,string(data))
}