package controllers

import (
	//"time"
	//"strconv"
	//"os"
	//"blog/models"
	"github.com/astaxie/beego"
	//"strings"
)

//PageController 登录控制器结构体
type PageController struct {
	beego.Controller
}

//Get *PageController 显示page类型单独页
func (c *PageController) Get() {
	// p := c.Ctx.Input.Param(":p")

	// art := models.Article{Slug: p}

	// if err == nil {
	// 	//评论列表
	// 	cmtlist := models.CommentsList(art.Cid)
	// 	c.Data["cmtlist"] = cmtlist
	// 	//浏览次数+1
	// 	views := art.Views + 1
	// 	art.Views = views
	// 	models.UpdateViews(&art)

	// 	c.Data["text"] = text
	// 	c.Data["art"] = art
	// 	c.Data["xsrfdata"]=c.XSRFFormHTML()

	// 	theme := tools.GetTheme()
	// 	tplname := "themes/" + theme + "/article.tpl"
	// 	c.TplName = tplname
	// } else {
	// 	c.Abort("401")
	// }
}