package controllers

import (
	//"strconv"
	"blog/models"
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

//LogoutController 退出登录控制器结构体
type LogoutController struct {
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
	user, err := models.GetUser(condition)

	if err == nil && user.Uid > 0 && models.CheckPassword(user, password) {
		SetUserSession(user, c)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "账号或密码错误"}
	}

	c.ServeJSON()
}

//SetUserSession 保存session
func SetUserSession(u models.Users, c *DologinController) {
	c.SetSession("is_login", 1)
	c.SetSession("user_session", u)
}

//Get LoginController 后台登陆页面
func (c *LogoutController) Get() {
	c.DestroySession()
	c.Redirect("/login", 302)
}