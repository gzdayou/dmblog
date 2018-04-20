package controllers

import (
	"blog/models"
	"blog/tools"
	"github.com/astaxie/beego"
	//"github.com/russross/blackfriday"
)

//CategoryController 分类列表
type CategoryController struct {
	beego.Controller
}

//Get *CategoryController 分类列表
func (c *CategoryController) Get() {
	p := c.Ctx.Input.Param(":p")
	condition := make(map[string]interface{})
	condition["slug"] = p
	cat, err := models.GetCategory(condition)

	if err == nil {
		r,_ := models.GetCateArticles(cat.Mid)
		var cids []int64
		for _, row := range(r) {
			cids = append(cids, row.Cid)
		}

		if len(cids) == 0 {
			cids = append(cids, 0)//防止数据库报错
		}

		condition := make(map[string]interface{})
		condition["cidin"] = cids 

		if _, list, e := models.ListArticle(condition, 1, 10); e == nil {
			c.Data["list"] = list
			c.Data["cat"] = cat

			theme := tools.GetTheme()
			tplname := "themes/" + theme + "/category.tpl"
			c.TplName = tplname
		} 
	} else {
		c.Abort("401")
	}
}