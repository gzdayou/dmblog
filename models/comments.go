package models

import (
	. "blog/base"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Comments 评论model结构体类型
type Comments struct {
	Coid     int64 `orm:"pk"`
	Cid      int64
	Created  int64
	Author   string
	Authorid int64
	Ownerid  int64
	Mail     string
	Url      string
	Ip       string
	Agent    string
	Text     string
	Type     string
	Status   string
	Parent   int64
}

//TableName 表名
func (c *Comments) TableName() string {
	return "comments"
}

func init() {
	InitSQL()
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(Comments))
}

//AddComments 添加博文
func AddComments(comments Comments) (int64, error) {
	o := orm.NewOrm()
	cmt := new(Comments)
	cmt.Cid = comments.Cid
	cmt.Created = time.Now().Unix()
	cmt.Author = comments.Author
	cmt.Authorid = comments.Authorid
	cmt.Ownerid = comments.Ownerid
	cmt.Mail = comments.Mail
	cmt.Url = comments.Url
	cmt.Ip = comments.Ip
	cmt.Agent = comments.Agent
	cmt.Text = comments.Text
	cmt.Type = "comment"
	cmt.Status = "approved"
	cmt.Parent = comments.Parent

	id, err := o.Insert(cmt)
	return id, err
}

//ListComments 根据博文ID获取博文评论列表（一级）
func ListComments(cid int64, page int, limit int) (num int64, list []Comments, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "comments")
	cond := orm.NewCondition()
	cond = cond.And("cid", cid).And("parent", 0)
	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	start := (page - 1) * limit

	var comts []Comments
	num, err1 := qs.OrderBy("Coid").Limit(limit, start).All(&comts)
	return num, comts, err1
}

//GetCommentsByid 根据评论ID获取子评论列表
func GetCommentsByid(coid int64) (num int64, list []Comments, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "comments")
	cond := orm.NewCondition()
	cond = cond.And("parent", coid)
	qs = qs.SetCond(cond)

	var comts []Comments
	num, err1 := qs.OrderBy("Coid").All(&comts)
	return num, comts, err1
}