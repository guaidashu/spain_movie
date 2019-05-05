package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MvContent struct {
	Id       int
	Url      string
	VideoSrc string
	ParentId int
}

func (mv *MvContent) TableName() string {
	return "mv_content"
}

func GetContentByParentId(id int) (*MvContent, error) {
	o := orm.NewOrm()
	var mvContent MvContent
	err:=o.QueryTable("mv_content").Filter("parent_id",id).One(&mvContent)
	if err != nil {
		logs.Error("电影内容获取出错: ",err.Error())
	}
	return &mvContent, err
}