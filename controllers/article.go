package controllers

import (
	. "blog/base"
	. "blog/models"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
)

//AddArticleController ...
type AddArticleController struct {
	beego.Controller
}

//ArticleController ...
type ArticleController struct {
	beego.Controller
}

//Get AddArticleController ...
func (c *AddArticleController) Get() {
	// var art Article
	// art.Status = "1"
	// c.Data["art"] = art

	tplname := "admin/template/article/add-form.tpl"
	c.Layout = "admin/template/layout/default.tpl"
	c.TplName = tplname
}

//Post AddArticleController ...
func (c *AddArticleController) Post() {
	var art Article
	local, _ := time.LoadLocation("Asia/Chongqing")
	tm2, _ := time.ParseInLocation("2006-01-02 15:04", c.GetString("date"), local)
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

//Get ArticleController 博文详情方法
func (c *ArticleController) Get() {
	id := c.Ctx.Input.Param(":id")
	cid, _ := strconv.ParseInt(id, 10, 64)
	art, err := GetArticle(cid)

	if err == nil {
		//评论列表
		cmtlist := CommentsList(cid)
		c.Data["cmtlist"] = cmtlist
		//上一篇
		preCid, preTitle := GetPreArticle(cid)
		//下一篇
		nextCid, nextTitle := GetNextArticle(cid)
		text := string(blackfriday.MarkdownBasic([]byte(art.Text)))
		c.Data["preCid"] = preCid
		c.Data["preTitle"] = preTitle
		c.Data["nextCid"] = nextCid
		c.Data["nextTitle"] = nextTitle
		c.Data["text"] = text
		c.Data["art"] = art

		theme := GetTheme()
		tplname := "themes/" + theme + "/article.tpl"
		c.TplName = tplname
	} else {
		c.Abort("401")
	}
}

//CommentsList 获取博文列表
func CommentsList(Cid int64) string {
	_, list, err := ListComments(Cid, 1, 10)
	if err == nil && len(list) > 0 {
		s := SubCommentsList(list, 5)

		return s
	}
	return ""
}

//SubCommentsList 根据一级评论列表组装详细子评论列表HTML
func SubCommentsList(list []Comments, max int) string {
	s := `<ol class="comment-list">`
	for _, comment := range list {
		s += `<li id="comment-` + string(comment.Coid) + `">`
		avatarMd5 := ToMd5(comment.Author)
		s += `<img class="avatar" src="http://www.gravatar.com/avatar/`+ avatarMd5 +`?s=150&amp;r=G&amp;d=robohash" alt="duomi" width="150" height="150" />`
		s += `<div class="comment-meta">`
		s += `<span class="comment-author">`
		if comment.Url == "" {
			s += comment.Author
		} else {
			s += `<a href="` + comment.Url + `" rel="external nofollow">` + comment.Author + `</a>`
		}
		s += `</span>`
		s += `<time class="comment-time">18.03.14</time>`
		line := fmt.Sprintf("<span class=\"comment-reply\"><a href=\"/article/%d/?replyTo=%d#respond-post-%d\" rel=\"nofollow\" onclick=\"return TypechoComment.reply('comment-%d', %d);\">回复</a></span>", comment.Cid, comment.Coid, comment.Cid, comment.Coid, comment.Coid)
		s += line
		s += `</div>`
		s += `<div class="comment-content">`
		s += `<p>` + comment.Text + `</p>`
		s += `</div>`
		//是否还有子评论列表
		_, listSub, err := GetCommentsByid(comment.Coid)
		if err == nil && len(listSub) > 0 && max >= 1 {
			n := max - 1
			s += `<div class="comment-children">`
			s += SubCommentsList(listSub, n)
			s += `</div>`
		}
		s += `</li>`
	}
	s += `</ol>`
	return s
}
