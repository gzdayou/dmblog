package models

import (
	//. "blog/base"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Categories 博文model结构体类型
type Categories struct {
	Mid				int64 `orm:"pk"`
	Name			string
	Slug			string
	Type			string
	Description		string
	Count			int64
	Order			int64
	Parent			int64
}

//TableName 表名
func (c *Categories) TableName() string {
	return beego.AppConfig.String("dbprefix") + "categories"
}

func init() {
	//InitSQL()
	orm.RegisterModel(new(Categories))
}

//ListCategories 分类列表
//@param p 父级分类ID
func ListCategories(p int64) (num int64, list []Categories, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "categories")
	cond := orm.NewCondition()
	cond = cond.And("parent", p)
	qs = qs.SetCond(cond)

	var cat []Categories
	num, err = qs.OrderBy("order").All(&cat)
	return num, cat, err
}