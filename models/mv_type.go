package models

import "github.com/astaxie/beego/orm"

type MvType struct {
	Id          int
	Name        string
	Url         string
	ImgSrc      string
	IconImgSrc  string
	Description string
	Status      int
}

func (mvType *MvType) TableName() string {
	return "mv_type"
}

func GetCategory() (*[]MvType, error) {
	o := orm.NewOrm()
	mvType := &[]MvType{}
	_, err := o.QueryTable("mv_type").All(mvType)
	return mvType, err
}

func GetCategoryById(id int) *MvType {
	o := orm.NewOrm()
	mvType := &MvType{}
	_ = o.QueryTable("mv_type").Filter("id", id).One(mvType)
	return mvType
}
