package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type IndexData struct {
	//大写是可以公共变量，如果想用小写后面加json转换
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) IndexHtml(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 5. 定义所需要的数据
	// 数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	// 执行处理数据
	index.WriteData(w, hr)
}
