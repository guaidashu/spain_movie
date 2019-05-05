package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MvList struct {
	Id                int
	Url               string
	ImgSrc            string
	Description       string
	DescriptionPoster string
	OriginSrc         string
	Star              string
	Title             string
	PageViews         int
	Director          string
	Label             string
	CategoryId        int
	ImgCover          string
}

func (mvList *MvList) TableName() string {
	return "mv_list"
}

func GetMvListById(id int) (*MvList, error) {
	o := orm.NewOrm()
	mvList := &MvList{Id: id}
	err := o.QueryTable("mv_list").Filter("id", id).One(mvList)
	if err != nil {
		logs.Error("无此电影 id", id)
	}
	return mvList, err
}

func GetMvListByCategory(category int, size int) (*[]*MvList, int, error) {
	o := orm.NewOrm()
	mvList := &[]*MvList{}
	num, err := o.QueryTable("mv_list").Filter("category_id", category).Limit(size).All(mvList)
	if err != nil {
		logs.Error("无此电影 id", category)
	}
	return mvList, (int(num)/size + 1), err
}

func GetMvListByCategoryWithPage(category int, size int, page int) (*[]*MvList, int, error) {
	o := orm.NewOrm()
	mvList := &[]*MvList{}
	num, err := o.QueryTable("mv_list").Filter("category_id", category).Limit(size, getPosWithPage(page, size)).All(mvList)
	if err != nil {
		logs.Error("无此电影 id", category)
	}
	return mvList, (int(num)/size + 1), err
}

func GetMvListByTitleWithPage(title string, size int, page int) (*[]*MvList, int, error) {
	o := orm.NewOrm()
	mvList := &[]*MvList{}
	num, err := o.QueryTable("mv_list").Filter("title__icontains", title).Limit(size, getPosWithPage(page, size)).All(mvList)
	if err != nil {
		logs.Error("无此电影 id", title)
	}
	return mvList, (int(num)/size + 1), err
}

func getPosWithPage(page int, size int) int {
	return (page - 1) * size
}

func GetMvList(conf map[string]interface{}) (*[]*MvList, error) {
	o := orm.NewOrm()
	mvList := &[]*MvList{}
	if _, ok := conf["limit"]; ok {
		_, err := o.QueryTable("mv_list").Limit(conf["num"], conf["offset"]).All(mvList)
		if err != nil {
			logs.Error("电影列表获取出错")
		}
		return mvList, err
	} else {
		_, err := o.QueryTable("mv_list").All(mvList)
		if err != nil {
			logs.Error("电影列表获取出错")
		}
		return mvList, err
	}
}
