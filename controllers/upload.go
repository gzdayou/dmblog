package controllers

import (
	"time"
	"strconv"
	"os"
	"blog/tools"
	"github.com/astaxie/beego"
	"strings"
)

//UploadController 登录控制器结构体
type UploadController struct {
	beego.Controller
}

//Prepare UploadController ...
func (c *UploadController) Prepare() {
	c.EnableXSRF = false
}

//Post UploadController ...
func (c *UploadController) Post() {
	if c.GetSession("is_login") != 1 {
		c.Redirect("/login", 302)
	}
	//imgFile
	f, h, err := c.GetFile("file")
	defer f.Close()

	//生成上传路径
	now := time.Now()
	dir := "./static/upload/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
	err1 := os.MkdirAll(dir, 0755)
	if err1 != nil {
		c.Data["json"] = map[string]interface{}{"error": 1, "message": "目录权限不够"}
		c.ServeJSON()
		return
	}
	//生成新的文件名
	filename := h.Filename
	ext := tools.SubString(filename, strings.LastIndex(filename, "."), 5)
	filename = tools.GetGUID() + ext

	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": 1, "message": err}
	} else {
		c.SaveToFile("file", dir+"/"+filename)

		info := make(map[string]interface{})
		info["cid"] = 4
		info["title"] = filename
		info["type"] = "png"
		info["size"] = 185307
		info["bytes"] = "181 Kb"
		info["isImage"] = true
		info["url"] = beego.AppConfig.String("website")+strings.Replace(dir, ".", "", 1) + "/" + filename
		info["permalink"] = "/article/4/"
		c.Data["json"] = map[string]interface{}{"0": beego.AppConfig.String("website")+strings.Replace(dir, ".", "", 1) + "/" + filename, "1": info}
	}
	c.ServeJSON()
}