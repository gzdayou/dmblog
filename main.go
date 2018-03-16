package main

import (
	_ "blog/routers"
	. "blog/base"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/adresource","views/admin")
	theme := GetTheme()
	beego.SetStaticPath("/themepth","views/themes/" + theme)
	
	//注册自定义模板函数
	beego.AddFuncMap("strreplace", Strreplace)
	beego.AddFuncMap("stampToDatetime", StampToDatetime)

	beego.Run()
}

