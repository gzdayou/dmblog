package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Admin
	beego.Router("/AddArticle", &controllers.AddArticleController{})
	//网站
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{})
}
