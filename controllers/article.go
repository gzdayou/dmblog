package controllers

import (
	//"blog/base"
	"blog/models"
	"blog/tools"
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

//EditArticleController 编辑博文页面展示
type EditArticleController struct {
	beego.Controller
}

//DoEditArticleController 编辑博文保存处理
type DoEditArticleController struct {
	beego.Controller
}

//Get AddArticleController ...
func (c *AddArticleController) Get() {
	if c.GetSession("is_login") != 1 {
		c.Redirect("/login", 302)
	}

	//加载分类列表（暂时只用二级）
	_, list, _ := models.ListCategories(0)
	c.Data["list"] = list
	c.Data["xsrfdata"]=c.XSRFFormHTML()

	tplname := "admin/template/article/add-form.tpl"
	c.Layout = "admin/template/layout/default.tpl"
	c.TplName = tplname
}

//Post AddArticleController 添加博客文章
func (c *AddArticleController) Post() {
	if c.GetSession("is_login") != 1 {
		c.Redirect("/login", 302)
	}
	var art models.Article
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

	id, err := models.AddArticle(art, c.GetStrings("category[]"))
	if err == nil {
		c.Redirect(fmt.Sprintf("/article/%d", id), 302)
	} else {
		c.Redirect("/", 302)
	}

	//c.ServeJSON()
}

//Get ArticleController 博文详情方法
func (c *ArticleController) Get() {
	id := c.Ctx.Input.Param(":id")
	cid, _ := strconv.ParseInt(id, 10, 64)
	art, err := models.GetArticle(cid)

	if err == nil {
		//评论列表
		cmtlist := CommentsList(cid)
		c.Data["cmtlist"] = cmtlist
		//上一篇
		preCid, preTitle := models.GetPreArticle(cid)
		//下一篇
		nextCid, nextTitle := models.GetNextArticle(cid)
		text := string(blackfriday.MarkdownBasic([]byte(art.Text)))
		//浏览次数+1
		views := art.Views + 1
		art.Views = views
		models.UpdateViews(&art)

		c.Data["preCid"] = preCid
		c.Data["preTitle"] = preTitle
		c.Data["nextCid"] = nextCid
		c.Data["nextTitle"] = nextTitle
		c.Data["text"] = text
		c.Data["art"] = art
		c.Data["xsrfdata"]=c.XSRFFormHTML()

		theme := tools.GetTheme()
		tplname := "themes/" + theme + "/article.tpl"
		c.TplName = tplname
	} else {
		c.Abort("401")
	}
}

//CommentsList 获取博文列表
func CommentsList(Cid int64) string {
	_, list, err := models.ListComments(Cid, 1, 10)
	if err == nil && len(list) > 0 {
		s := SubCommentsList(list, 5)

		return s
	}
	return ""
}

//SubCommentsList 根据一级评论列表组装详细子评论列表HTML
func SubCommentsList(list []models.Comments, max int) string {
	s := `<ol class="comment-list">`
	for _, comment := range list {
		sLi := fmt.Sprintf("<li id=\"comment-%d\">", comment.Coid)
		s += sLi
		avatarMd5 := tools.ToMd5(comment.Author)
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
		if max > 1 {
			line := fmt.Sprintf("<span class=\"comment-reply\"><a href=\"javascript::void(0);\" rel=\"nofollow\" onclick=\"return TypechoComment.reply('comment-%d', %d);\">回复</a></span>", comment.Coid, comment.Coid)
			s += line
		}
		s += `</div>`
		s += `<div class="comment-content">`
		s += `<p>` + comment.Text + `</p>`
		s += `</div>`
		//是否还有子评论列表
		_, listSub, err := models.GetCommentsByid(comment.Coid)
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

//Get EditArticleController 编辑博文
func (c *EditArticleController) Get() {
	if c.GetSession("is_login") != 1 {
		c.Redirect("/login", 302)
	}
	id := c.Ctx.Input.Param(":id")
	cid, _ := strconv.ParseInt(id, 10, 64)
	art, err := models.GetArticle(cid)

	if err == nil {
		c.Data["art"] = art
		c.Data["created"] = tools.StampToDatetime(art.Created)

		//加载分类列表（暂时只用二级）
		_, list, _ := models.ListCategories(0)
		c.Data["list"] = list

		//文章目录
		r, _ := models.GetArticleCats(cid)
		relate := make(map[int64]bool)
		for _, row := range(r) {
			relate[row.Mid] = true
		}
		c.Data["relate"] = relate
		c.Data["xsrfdata"]=c.XSRFFormHTML()

		tplname := "admin/template/article/edit-form.tpl"
		c.Layout = "admin/template/layout/default.tpl"
		c.TplName = tplname
	} else {
		beego.Error(err)
		c.Abort("401")
	}
}

//Post EditArticleController 处理编辑博文
func (c * EditArticleController) Post() {
	if c.GetSession("is_login") != 1 {
		c.Redirect("/login", 302)
	}
	id := c.Ctx.Input.Param(":id")
	cid, _ := strconv.ParseInt(id, 10, 64)

	var art models.Article
	local, _ := time.LoadLocation("Asia/Chongqing")
	tm2, _ := time.ParseInLocation("2006-01-02 15:04", c.GetString("date"), local)
	art.Cid = cid
	art.Title = c.GetString("title")
	art.Slug = "test2"
	art.Created = tm2.Unix()
	art.Modified = time.Now().Unix()
	art.Text = c.GetString("text")
	if err := models.EditArticle(art, c.GetStrings("category[]")); err == nil {
		c.Redirect(fmt.Sprintf("/article/%d", cid), 302)
	}

	c.Redirect(fmt.Sprintf("/article/edit/%d", cid), 302)
}