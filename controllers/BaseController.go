package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/yy/spain_movie/models"
)

type HeaderData struct {
	CategoryList *[]models.MvType
}

type BaseController struct {
	beego.Controller
	Header HeaderData
}

func (c *BaseController) Prepare() {

	c.initHeaderData()

	c.Data["Header"] = c.Header
}

func (c *BaseController) initHeaderData() {
	// 初始化Category 列表
	c.Header.CategoryList, _ = models.GetCategory()
}

func (c *BaseController) PageUtil(totalPage, current int, URL, search string, ) string {
	// search  ==   ?s=xxx   用于搜索页面
	pageHTML := `<ul class="pagination">`
	if totalPage == 0 {
		return ""
	}
	if current != 1 {
		pageHTML += fmt.Sprintf(`<li><a class="previouspostslink'" rel="nofollow" href="%s%s">< First</a>`, URL, search)
	}
	if current+2 <= totalPage && current-2 >= 1 {
		for i := current - 2; i <= current+2; i++ {
			fmt.Println(i)
			if i == current {
				pageHTML += fmt.Sprintf(`<li class="active"><a class="">%d</a></li>`, i)
			} else {
				pageHTML += fmt.Sprintf(`<li><a rel="nofollow" class="page larger" href="%s&page=%d%s">%d</a></li>`, URL, i, search, i)
			}
		}
	} else if current-2 < 1 {
		endPage := 5
		if totalPage <= 5 {
			endPage = totalPage
		}
		for i := 1; i <= endPage; i++ {
			fmt.Println(i)
			if i == current {
				pageHTML += fmt.Sprintf(`<li class="active"><a class="">%d</a></li>`, i)
			} else {
				pageHTML += fmt.Sprintf(`<li><a rel="nofollow" class="page larger" href="%s&page=%d%s">%d</a></li>`, URL, i, search, i)
			}
		}
	} else if current+2 >= totalPage {
		startPage := totalPage - 5
		if totalPage-5 <= 0 {
			startPage = 1
		}
		for i := startPage; i <= totalPage; i++ {
			fmt.Println(i)
			if i == current {
				pageHTML += fmt.Sprintf(`<li class="active"><a class="">%d</a></li>`, i)
			} else {
				pageHTML += fmt.Sprintf(`<li><a rel="nofollow" class="page larger" href="%s&page=%d%s">%d</a></li>`, URL, i, search, i)
			}
		}
	}

	if current != totalPage {
		pageHTML += fmt.Sprintf(`<li><a rel="nofollow" href="%s&page=%d%s">Last ></a></li></ul>`, URL, totalPage, search)
	} else {
		pageHTML += `</ul>`
	}
	return pageHTML
}