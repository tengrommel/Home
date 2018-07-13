package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (u *UserController)RetData(resp map[string]interface{})  {
	u.Data["json"] = resp
	u.ServeJSON()
}

func (u *UserController)Reg()  {
	resp := make(map[string]interface{})

	json.Unmarshal(u.Ctx.Input.RequestBody, &resp)
	beego.Info(`resp["mobile"] =`, resp["mobile"])
}
