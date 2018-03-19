package controllers

import (
	//"strconv"
	. "blog/models"
	//"fmt"
	//. "blog/base"
	"github.com/astaxie/beego"
)

//LoginController 登录控制器结构体
type LoginController struct {
	beego.Controller
}

//DologinController 登录验证控制器结构体
type DologinController struct {
	beego.Controller
}

//Get LoginController 后台登陆页面
func (c *LoginController) Get() {
	c.TplName = "admin/template/login/form.tpl"
}

//Post DologinController 登录验证控制器
func (c *DologinController) Post() {
	name := c.Ctx.Input.Query("name")
	password := c.Ctx.Input.Query("password")

	if name == "" || password == "" {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名或密码为空"}
		c.ServeJSON()
	}

	condition := make(map[string]string)
	condition["name"] = name
	condition["password"] = password
	user, err := GetUser(condition)

	if err == nil && user.Uid > 0 {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "账号或密码错误"}
	}

	c.ServeJSON()
}