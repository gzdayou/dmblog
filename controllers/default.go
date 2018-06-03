package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"blog/tools"
	"strconv"
)

//MainController 首页控制器结构体
type MainController struct {
	beego.Controller
}

//Get MainController 首页控制器
func (c *MainController) Get() {
	pageStr := c.GetString("page")
	page, err := strconv.Atoi(pageStr)
	limit := 2
	condition := make(map[string]interface{})
	condition["type"] = "post"
	
	num, list, err := models.ListArticle(condition, page, limit)
	
	if err == nil {
		hasNext, hasPre := false, false
		if int64(page * limit) < num {
			hasNext = true
		}
		if page > 1 {
			hasPre = true
		}

		c.Data["num"] = num
		c.Data["list"] = list
		c.Data["hasNext"] = hasNext
		c.Data["hasPre"] = hasPre
		c.Data["prePage"] = page - 1
		c.Data["nextPage"] = page + 1
		theme := tools.GetTheme()
		tplname := "themes/"+ theme +"/index.tpl"
		c.TplName = tplname
	} else {
		c.Abort("401")
	}
}
