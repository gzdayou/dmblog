package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"blog/tools"
	//"encoding/json"
)

//MainController 首页控制器结构体
type MainController struct {
	beego.Controller
}

//Get MainController 首页控制器
func (c *MainController) Get() {
	condition := make(map[string]string)
	condition["title"] = "abcdefg"
	
	num, list, err := models.ListArticle(condition, 1, 10)
	if err == nil {
		c.Data["num"] = num
		// b,_ := json.Marshal(list)
		// c.Ctx.WriteString(string(b))
		c.Data["list"] = list
		theme := tools.GetTheme()
		tplname := "themes/"+ theme +"/index.tpl"
		c.TplName = tplname
	} else {
		c.Abort("401")
	}
}
