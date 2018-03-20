package models

import (
	//"fmt"
	. "blog/base"
	//"golang.org/x/crypto/scrypt" 
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
	//InitSQL()
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
	if condition["uid"] != "" {
		cond = cond.And("uid", condition["uid"])
	}
	qs = qs.SetCond(cond)

	var user Users
	err1 := qs.One(&user)
	return user, err1
}

//CheckPassword 验证密码是否正确
func CheckPassword(u Users, p string) bool {
	dk := ToMd5(p + u.AuthCode)
	if string(dk) == u.Password {
		return true
	}
	return false
}