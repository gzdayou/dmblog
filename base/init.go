package base

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"strings"
)

func init() {
	//InitSQL()
}

//初始化数据库连接
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

//获取当前模板主题名称
func GetTheme() string {
	str := "default"
	return str
}

//自定义模板处理函数
func Strreplace(in string, search string, replace string)(out string){
	out = strings.Replace(in, search, replace, -1)
	return
}

//时间戳转日期模板处理函数
func StampToDatetime(input int64) string {
	timeLayout := "2006-01-02 15:04:05" 
	tm := time.Unix(input, 0)
	dataTimeStr := tm.Format(timeLayout)
	return dataTimeStr
}