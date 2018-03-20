package base

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	InitSQL()
}

//InitSQL 初始化数据库连接
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

//ToMd5 md5转换
func ToMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}