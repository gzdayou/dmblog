package controllers

import (
	"github.com/astaxie/beego"
	. "blog/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var art Article
	art.Title = "测试标题"

	id, err := AddArticle(art)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "博客添加成功", "id": id}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	}
	c.ServeJSON()

	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}
