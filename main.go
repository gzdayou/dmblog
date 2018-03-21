package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"blog/tools"
)

func main() {
	beego.SetStaticPath("/static","static")
	beego.SetStaticPath("/adresource","views/admin")
	theme := tools.GetTheme()
	beego.SetStaticPath("/themepth","views/themes/" + theme)
	
	//注册自定义模板函数
	tools.AddSelfTemplateFuncs()

	beego.Run()
}

