package web

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type testController struct {
	web.Controller
}

func (s *testController) Post() {
	// request := new(Request)
	fmt.Println("test controller input:", string(s.Ctx.Input.RequestBody))
	// json.Unmarshal(s.Ctx.Input.RequestBody, &request)

	var res = make(map[string]interface{})
	res["msg"] = "success"
	res["code"] = 200

	s.Data["json"] = res
	_ = s.ServeJSON()
	return
}

func (s *testController) Get() {
	s.Post()
}
