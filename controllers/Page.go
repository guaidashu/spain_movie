package controllers

import (
	"fmt"
	"github.com/yy/spain_movie/models"
)

type PageController struct {
	BaseController
}

func (b *PageController) Category() {
	id, _ := b.GetInt("id")
	page, _ := b.GetInt("page")
	search := ""
	if page == 0 {
		page = 1
	}
	URL := fmt.Sprintf("/page/category?id=%d", id)
	movieData, totalPage, _ := models.GetMvListByCategoryWithPage(id, 16, page)
	category := models.GetCategoryById(id)
	b.Data["TitleName"] = category.Name
	b.Data["pagination"] = b.PageUtil(totalPage, page, URL, search)
	b.Data["OtherMovie"] = movieData
	b.Layout = "layout/default.html"
	b.TplName = "page/other_page.html"
}

func (b *PageController) Search() {
	// 接收get参数(搜索关键词)
	keyword := b.GetString("s")
	page, _ := b.GetInt("page")
	search := ""
	if page == 0 {
		page = 1
	}
	URL := fmt.Sprintf("/page/search?s=%s", keyword)
	b.Data["TitleName"] = keyword
	movieData, totalPage, _ := models.GetMvListByTitleWithPage(keyword, 16, page)
	b.Data["pagination"] = b.PageUtil(totalPage, page, URL, search)
	b.Data["OtherMovie"] = movieData
	b.Layout = "layout/default.html"
	b.TplName = "page/other_page.html"
}
