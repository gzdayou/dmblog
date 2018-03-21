package models

import (
	//. "blog/base"
	"strconv"
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
	//InitSQL()
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(Article))
}

//AddArticle 添加博文
func AddArticle(updArt Article, category []string) (int64, error) {
	o := orm.NewOrm()
	err := o.Begin() //开启事务

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

	if err != nil {
		err = o.Rollback()
	}

	//保存分类
	for i := 0; i < len(category); i++ {
		var r Relationships
		r.Cid = id
		mid, _ := strconv.ParseInt(category[i], 10, 64)
		r.Mid = mid
		_, err := o.Insert(&r)
		if err != nil {
			err = o.Rollback()
		}
	}
	err = o.Commit()

	return id, err
}

//EditArticle 编辑博文
func EditArticle(updArt Article, category []string) error {
	o := orm.NewOrm()
	err := o.Begin() //开启事务

	art := Article{Cid: updArt.Cid}
	if o.Read(&art) == nil {
		art.Title = updArt.Title
		art.Slug = "test3"
		art.Created = updArt.Created
		art.Modified = updArt.Modified
		art.Text = "<!--markdown-->" + updArt.Text
		if _, e := o.Update(&art); e != nil {
			err = o.Rollback()
		}

		//删除旧分类关系，写入新分类关系
		if len(category) > 0 {
			if e := DeleteRelate(updArt.Cid); e != nil {
				err = o.Rollback()
			}
			//保存分类
			for i := 0; i < len(category); i++ {
				var r Relationships
				r.Cid = updArt.Cid
				mid, _ := strconv.ParseInt(category[i], 10, 64)
				r.Mid = mid
				_, err := o.Insert(&r)
				if err != nil {
					err = o.Rollback()
				}
			}
		}
		
		err = o.Commit()
	}

	return err
}

//ListArticle 获取博文列表
func ListArticle(condition map[string]interface{}, page int, limit int) (num int64, list []Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "article")
	cond := orm.NewCondition()
	if condition["keyword"] != nil {
		cond = cond.And("title__icontains", condition["keyword"])
	}
	if condition["status"] != nil {
		cond = cond.And("status", condition["status"])
	}
	if condition["type"] != nil {
		cond = cond.And("type", condition["type"])
	}
	if condition["cidin"] != nil {
		cond = cond.And("cid__in", condition["cidin"])
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

//UpdateViews 更新浏览次数
func UpdateViews(upd *Article) error {

	o := orm.NewOrm()
	_, err := o.Update(upd)

	return err
}

//GetArticleByslug 根据slug获取page类型详情
func GetArticleByslug(s string) (Article, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "article")
	cond := orm.NewCondition()
	cond = cond.And("slug", s)
	cond = cond.And("type", "page")
	qs = qs.SetCond(cond)
	var articles Article
	err := qs.One(&articles)
	return articles, err
}