package web

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

//go:embed config.json
var configFile []byte

type BeegoConfig struct {
	ListenConf          ListenConfig  `json:"listen_conf"`
}

type ListenConfig struct {
	HttpAddr    string `json:"http_addr"`
	HttpPort    int    `json:"http_port"`
	EnableAdmin bool   `json:"enable_admin"`
	AdminAddr   string `json:"admin_addr"`
	AdminPort   int    `json:"admin_port"`
}

func initConfig() {
	beegoConfig := new(BeegoConfig)
	fmt.Println("listenConfig", beegoConfig)
	json.Unmarshal(configFile, beegoConfig)
	fmt.Println("listenConfig", beegoConfig)

	web.BConfig.CopyRequestBody = true
	web.BConfig.ServerName = "beego_test"
	web.BConfig.AppName = "beego_test"

	web.BConfig.Listen.HTTPAddr = beegoConfig.ListenConf.HttpAddr
	web.BConfig.Listen.HTTPPort = beegoConfig.ListenConf.HttpPort
	web.BConfig.Listen.EnableAdmin = beegoConfig.ListenConf.EnableAdmin
	web.BConfig.Listen.AdminAddr = beegoConfig.ListenConf.AdminAddr
	web.BConfig.Listen.AdminPort = beegoConfig.ListenConf.AdminPort
	fmt.Println("web.BConfig.Listen", web.BConfig.Listen)

	web.Router("/test_post", new(testController))
}
