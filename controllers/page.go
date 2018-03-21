package controllers

import (
	"blog/models"
	"blog/tools"
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
)

//PageController 登录控制器结构体
type PageController struct {
	beego.Controller
}

//Get *PageController 显示page类型单独页
func (c *PageController) Get() {
	p := c.Ctx.Input.Param(":p")

	art, err := models.GetArticleByslug(p)

	if err == nil {
		//评论列表
		cmtlist := CommentsList(art.Cid)
		c.Data["cmtlist"] = cmtlist
		//浏览次数+1
		views := art.Views + 1
		art.Views = views
		models.UpdateViews(&art)

		c.Data["text"] = string(blackfriday.MarkdownBasic([]byte(art.Text)))
		c.Data["art"] = art
		c.Data["xsrfdata"]=c.XSRFFormHTML()

		theme := tools.GetTheme()
		tplname := "themes/" + theme + "/page.tpl"
		c.TplName = tplname
	} else {
		c.Abort("401")
	}
}