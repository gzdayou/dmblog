package controllers

import (
	"strconv"
	"blog/models"
	"fmt"
	//. "blog/base"
	"github.com/astaxie/beego"
)

//AddCommentsController 增加评论结构体
type AddCommentsController struct {
	beego.Controller
}

//Post AddCommentsController 添加评论函数
func (c *AddCommentsController) Post() {
	req := c.Ctx.Request
	cid, _ := strconv.ParseInt(c.GetString("cid"), 10, 64)
	parent, _ := strconv.ParseInt(c.GetString("parent"), 10, 64)

	var cmt models.Comments
	cmt.Cid = cid
	cmt.Author = c.GetString("author")
	cmt.Authorid = 0
	cmt.Ownerid = 1
	cmt.Mail = c.GetString("mail")
	cmt.Url = c.GetString("url")
	cmt.Ip = req.RemoteAddr
	cmt.Agent = req.UserAgent()
	cmt.Text = c.GetString("text")
	cmt.Parent = parent

	id, err := models.AddComments(cmt)
	if err == nil {
		rURL := fmt.Sprintf("/article/%d/#comment-%d", cid, id)
		c.Redirect(rURL, 302)
	} else {
		c.Abort("401")
	}
}