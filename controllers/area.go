package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}

func (c *AreaController)ReturnData(resp map[string]interface{})  {
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connect success")
	resp := make(map[string]interface{})

	defer c.ReturnData(resp)
	// 从session中拿到数据

	//从数据库拿到数据
	var areas []models.Area
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil{
		beego.Info("数据错误")
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"
		return
	}
	if num == 0{
		beego.Info("没有数据")
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}

	// 打包成json返回给前端
	resp["errno"] = 0
	resp["errmsg"] = "ok"
	resp["data"] = &areas
	beego.Info("query data success, resp =", resp, "num =", num)
}