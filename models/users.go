package models

import (
	. "blog/base"
	//"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Users 用户model结构体类型
type Users struct {
	Uid 		int64 `orm:"pk"`
	Name 		string
	Password 	string
	Mail 		string
	Url 		string
	ScreenName 	string
	Created 	int64
	Activated 	int64
	Logged 		int64
	Group 		string
	AuthCode 	string
}


//TableName 表名
func (c *Users) TableName() string {
	return "users"
}

func init() {
	InitSQL()
	orm.RegisterModelWithPrefix(beego.AppConfig.String("dbprefix"), new(Users))
}

//GetUser 获取用户信息
func GetUser(condition map[string]string) (u Users, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(beego.AppConfig.String("dbprefix") + "users")
	cond := orm.NewCondition()
	if condition["name"] != "" {
		cond = cond.And("name", condition["name"])
	}
	if condition["password"] != "" {
		cond = cond.And("password", condition["password"])
	}
	qs = qs.SetCond(cond)

	var user Users
	err1 := qs.OrderBy("-Uid").One(&user)
	return user, err1
}