package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/adresource","views/admin")

	beego.Run()
}

