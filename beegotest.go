package main

import (
	"github.com/beego/beego/v2/server/web"
	"gotest/myweb"
)

func main() {
	myweb.InitConfig()
	web.Run()
}
