package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "blog/base"
)

type Article struct {
	Cid          int64 `orm:"pk"`
	Title        string
	Slug         string
	Created      int64
	Modified     int64
	Text         string
	Order        int64
	Authorid     int64
	Template     string
	Type         string
	Status       string
	Password     string
	CommentsNum  int64
	AllowComment int
	AllowPing    int
	AllowFeed    int
	Parent      int64
	Views        int64
}

func (a *Article) TableName() string {
	return "article"
}

func init() {
	InitSQL()
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(Article))
}

func AddArticle(updArt Article) (int64, error) {
	o := orm.NewOrm()
	art := new(Article)
	art.Title = updArt.Title
	art.Slug = "test3"
	art.Created = updArt.Created
	art.Modified = updArt.Modified
	art.Text = "<!--markdown-->" + updArt.Text
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

	id, err := o.Insert(art)
	return id, err
}

func ListArticle(condition map[string]string, page int, limit int) (num int64, list []Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "article")
	cond := orm.NewCondition()
	if condition["keyword"] != "" {
		cond = cond.And("title__icontains", condition["keyword"])
	}
	if condition["status"] != "" {
		cond = cond.And("status", condition["status"])
	}
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	start := (page - 1) * limit

	var articles []Article
	num, err1 := qs.Limit(limit, start).All(&articles)
	return num, articles, err1
}

func GetArticle(id int64) (Article, error) {
	o := orm.NewOrm()
	art := Article{Cid: id}

	err := o.Read(&art)

	return art, err
}