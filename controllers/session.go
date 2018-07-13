package controllers

import "github.com/astaxie/beego"

type SessionController struct {
	beego.Controller
}

func (s *SessionController)ReturnData(resp map[string]interface{})  {
	s.Data["json"] = resp
	s.ServeJSON()
}

func (s *SessionController)GetSessionData()  {
	resp := make(map[string]interface{})
	resp["errno"] = 4001
	resp["errmsg"] = "查询成功"
	s.ReturnData(resp)
}
