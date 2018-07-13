package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
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
	// 获取前端传入的数据
	json.Unmarshal(u.Ctx.Input.RequestBody, &resp)
	//beego.Info(`resp["mobile"] =`, resp["mobile"])
	//beego.Info(`resp["password"] =`, resp["password"])
	//beego.Info(`resp["sms_code"] =`, resp["sms_code"])
	o := orm.NewOrm()
	user := models.User{}
	user.Password_hash = resp["password"].(string)
	user.Mobile = resp["mobile"].(string)
	user.Name = resp["mobile"].(string)
	id, err := o.Insert(&user)
	if err != nil{
		resp["errno"] = 4002
		resp["errmsg"] = "注册失败"
		return
	}
	beego.Info("reg success, id = ", id)
	resp["errno"] = 0
	resp["errmsg"] = "注册成功"
	return
}
