package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Admin
	beego.Router("/AddArticle", &controllers.AddArticleController{})//添加文章
	//网站
	beego.Router("/", &controllers.MainController{})//首页
	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{})//文章详情
	beego.Router("/article/edit/:id([0-9]+)", &controllers.EditArticleController{})
	//beego.Router("/article/edit/:id([0-9]+)", &controllers.DoEditArticleController{})
	beego.Router("/AddComments", &controllers.AddCommentsController{})//保存评论
	beego.Router("/login", &controllers.LoginController{})//登录页面
	beego.Router("/dologin", &controllers.DologinController{})//登录处理
	beego.Router("/logout", &controllers.LogoutController{})//登出
	beego.Router("/upload", &controllers.UploadController{})//上传图片
	beego.Router("/page/:p:string", &controllers.PageController{})//page类型单独页
}
