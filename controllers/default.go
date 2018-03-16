package controllers

import (
	"github.com/astaxie/beego"
	. "blog/models"
	. "blog/base"
	//"encoding/json"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	condition := make(map[string]string)
	condition["title"] = "abcdefg"
	
	num, list, err := ListArticle(condition, 1, 10)
	if err == nil {
		c.Data["num"] = num
		// b,_ := json.Marshal(list)
		// c.Ctx.WriteString(string(b))
		c.Data["list"] = list
		theme := GetTheme()
		tplname := "themes/"+ theme +"/index.tpl"
		c.TplName = tplname
	} else {
		c.Abort("401")
	}
}
