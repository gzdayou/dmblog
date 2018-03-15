package controllers

import (
	. "blog/models"
	"time"
	//"strconv"
	"github.com/astaxie/beego"
)

type AddArticleController struct {
	beego.Controller
}

func (c *AddArticleController) Get() {
	var art Article
	art.Status = "1"
	c.Data["art"] = art

	tplname := "admin/template/article/add-form.tpl"
	c.Layout = "admin/template/layout/default.tpl"
	c.TplName = tplname
}

func (c *AddArticleController) Post() {
	var art Article
	tm2, _ := time.Parse("2006-01-02 15:04", c.GetString("date"))
	art.Title = c.GetString("title")
	art.Slug = "test2"
	art.Created = tm2.Unix()
	art.Modified = time.Now().Unix()
	art.Text = c.GetString("text")
	art.Order = 0
	art.Authorid = 1
	art.Template = ""
	art.Type = "post"
	art.Status = "publish"
	art.Password = ""
	art.CommentsNum = 0
	art.AllowComment = 1
	art.AllowPing = 1
	art.AllowFeed = 1
	art.Parent = 0
	art.Views = 0

	id, err := AddArticle(art)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "博客添加成功", "id": id}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "博客添加出错"}
	}

	c.ServeJSON()
}
