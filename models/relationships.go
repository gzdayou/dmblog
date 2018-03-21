package models

import (
	//"fmt"
	//. "blog/base"
	//"golang.org/x/crypto/scrypt" 
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Relationships 文章分类表结构体
type Relationships struct {
	Cid 		int64 `orm:"pk"`
	Mid 		int64
}


//TableName 表名
func (r *Relationships) TableName() string {
	return "relationships"
}

//TableUnique 多字段唯一键
func (r *Relationships) TableUnique() [][]string {
    return [][]string{
        []string{"Cid", "Mid"},
    }
}

func init() {
	//InitSQL()
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(Relationships))
}

//GetArticleCats 获取文章关联分类
func GetArticleCats(cid int64) ([]Relationships, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "relationships")
	cond := orm.NewCondition()
	cond = cond.And("cid", cid)
	qs = qs.SetCond(cond)

	var r []Relationships
	_, err := qs.All(&r)

	return r,err
}

//DeleteRelate 删除文章关联的分类
func DeleteRelate(cid int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Relationships{Cid: cid});

	return err
}

//GetCateArticles 获取分类下的文章ID
func GetCateArticles(mid int64)([]Relationships, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "relationships")
	cond := orm.NewCondition()
	cond = cond.And("mid", mid)
	qs = qs.SetCond(cond)

	var r []Relationships
	_, err := qs.All(&r)

	return r,err
}