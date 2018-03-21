package tools

import (
	"fmt"
	"strings"
	"time"
	"blog/models"
	"github.com/astaxie/beego"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/base64"
	"io"
	"github.com/russross/blackfriday"
	"regexp"
)

//AddSelfTemplateFuncs 注册自定义模板函数
func AddSelfTemplateFuncs() {
	beego.AddFuncMap("strreplace", Strreplace)
	beego.AddFuncMap("stampToDatetime", StampToDatetime)
	beego.AddFuncMap("getHeaderCatlist", GetHeaderCatlist)
	beego.AddFuncMap("getCatHTML", GetCatHTML)
	beego.AddFuncMap("toMarkdown", ToMarkdown)
	beego.AddFuncMap("trimHTML", TrimHTML)
	beego.AddFuncMap("subList", SubList)
}

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

//GetHeaderCatlist 前端页面头部分类列表
func GetHeaderCatlist() string {
	s := `<ul class="sub-menu">`
	_, list, _ := models.ListCategories(0)
	for _, cat := range(list) {
		s += `<li><a href="`+ fmt.Sprintf("/category/%s", cat.Slug) +`">`+ cat.Name +`</a></li>`
	}
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

//GetCatHTML 获取文章所属分类的HTML
func GetCatHTML(cid int64) string {
	s := `<span class="category">`
	//获取分类信息
	idx := 0
	relat, _ := models.GetArticleCats(cid)
	for _, r := range(relat) {
		if idx > 0  {
			s += `， `
		}
		condition := make(map[string]interface{})
		condition["mid"] = r.Mid
		c,_ := models.GetCategory(condition)
		s += `<a href="/category/`+ c.Slug +`">`+ c.Name +`</a>`
		idx++
	}
	s += `</span>`

	return s
}

//SubString 字串截取
func SubString(str string, begin, length int) string {
    rs := []rune(str)
    lth := len(rs)
    if begin < 0 {
        begin = 0
    }
    if begin >= lth {
        begin = lth
    }
    end := begin + length

    if end > lth {
        end = lth
    }
    return string(rs[begin:end])
}

//GetMd5String md5方法
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//GetGUID Guid方法
func GetGUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//ToMarkdown ...
func ToMarkdown(str string) string {
	return string(blackfriday.MarkdownBasic([]byte(str)))
}

//TrimHTML 过滤HTML标签
func TrimHTML(src string) string {
    //将HTML标签全转换成小写
    re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
    src = re.ReplaceAllStringFunc(src, strings.ToLower)
    //去除STYLE
    re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
    src = re.ReplaceAllString(src, "")
    //去除SCRIPT
    re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
    src = re.ReplaceAllString(src, "")
    //去除所有尖括号内的HTML代码，并换成换行符
    re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
    src = re.ReplaceAllString(src, "\n")
    //去除连续的换行符
    re, _ = regexp.Compile("\\s{2,}")
    src = re.ReplaceAllString(src, "\n")
    return strings.TrimSpace(src)
}

//SubList 文章列表截取部分字符串
func SubList(str string) string {
	return SubString(str, 0, 130)
}