package controllers

import "github.com/astaxie/beego"

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController)ReturnData(resp map[string]interface{})  {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *HouseIndexController)GetHouseIndex()  {
	resp := make(map[string]interface{})
	resp["errno"] = 4001
	resp["errmsg"] = "查询成功"
	c.ReturnData(resp)
}