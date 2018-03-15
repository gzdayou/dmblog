package base

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	//InitSQL()
}

func InitSQL() {
	user := beego.AppConfig.String("dbuser")
	passwd := beego.AppConfig.String("dbpass")
	host := beego.AppConfig.String("dbhost")
	port, err := beego.AppConfig.Int("dbport")
	dbname := beego.AppConfig.String("dbname")
	if nil != err {
		port = 3306
	}
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
}