package tools

import (
	"strings"
	"time"
	"blog/models"
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
)

//Strreplace 自定义模板处理函数
func Strreplace(in string, search string, replace string) (out string) {
	out = strings.Replace(in, search, replace, -1)
	return
}

//StampToDatetime 时间戳转日期模板处理函数
func StampToDatetime(input int64) string {
	timeLayout := "2006-01-02 15:04:05"
	tm := time.Unix(input, 0)
	dataTimeStr := tm.Format(timeLayout)
	return dataTimeStr
}


//AddSelfTemplateFuncs 注册自定义模板函数
func AddSelfTemplateFuncs() {
	beego.AddFuncMap("strreplace", Strreplace)
	beego.AddFuncMap("stampToDatetime", StampToDatetime)
	beego.AddFuncMap("getHeaderCatlist", GetHeaderCatlist)
}

//GetHeaderCatlist 前端页面头部分类列表
func GetHeaderCatlist() string {
	models.ListCategories(0)

	s := `<ul class="sub-menu">`
	s += `<li><a href="/default">默认分类</a></li>`
	s += `</ul>`

	return s
}

//GetTheme 获取当前模板主题名称
func GetTheme() string {
	str := "default"
	return str
}

//ToMd5 md5转换
func ToMd5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//SliceContains slice中是否包含某个值
func SliceContains(src []string, value string) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}