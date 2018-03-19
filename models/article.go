package models

import (
	. "blog/base"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Article 博文model结构体类型
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
	Parent       int64
	Views        int64
}

//TableName 表名
func (a *Article) TableName() string {
	return "article"
}

func init() {
	InitSQL()
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(Article))
}

//AddArticle 添加博文
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

//ListArticle 获取博文列表
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
	num, err1 := qs.OrderBy("-Cid").Limit(limit, start).All(&articles)
	return num, articles, err1
}

//GetArticle 根据博文ID获取博文详情
func GetArticle(id int64) (Article, error) {
	o := orm.NewOrm()
	art := Article{Cid: id}
	err := o.Read(&art)

	return art, err
}

//GetPreArticle 获取上一篇博文
func GetPreArticle(id int64) (int64, string) {
	o := orm.NewOrm()
	var art Article
	err := o.Raw("SELECT cid, title FROM "+beego.AppConfig.String("dbprefix")+"article WHERE cid = (SELECT max(cid) FROM "+beego.AppConfig.String("dbprefix")+"article WHERE cid < ?)", id).QueryRow(&art)
	if err == nil {
		return art.Cid, art.Title
	}

	return 0, ""
}

//GetNextArticle 获取下一篇博文
func GetNextArticle(id int64) (int64, string) {

	o := orm.NewOrm()
	var art Article
	err := o.Raw("SELECT cid, title FROM "+beego.AppConfig.String("dbprefix")+"article WHERE cid = (SELECT min(cid) FROM "+beego.AppConfig.String("dbprefix")+"article WHERE cid > ?)", id).QueryRow(&art)
	if err == nil {
		return art.Cid, art.Title
	}

	return 0, ""
}
