package controllers

import (
	"github.com/yy/spain_movie/models"
)

//
//func init() {
//	orm.RegisterModel(new(models.User))
//}

type IndexController struct {
	BaseController
}

func (c *IndexController) Get(){
	SugerenciaData, _, _ := models.GetMvListByCategory(1, 16)
	LastAdded, _, _ := models.GetMvListByCategory(2, 16)
	DownloadLeaderboard, _, _ := models.GetMvListByCategory(3, 16)
	c.Data["SugerenciaData"] = SugerenciaData
	c.Data["LastAdded"] = LastAdded
	c.Data["DownloadLeaderboard"] = DownloadLeaderboard
	c.Layout = "layout/default.html"
	c.TplName = "page/home.html"
}

func (c *IndexController) Index() {
	//o := orm.NewOrm()
	//user := models.User{Id: 1}
	//user := new(models.User)
	//var users = [] *models.User{}
	//user.Username = "guaidashu"
	//user.Password = "wyysdsa!"
	//result, _ := o.Insert(user)
	//_ = o.Read(&user)
	//result, _ := o.QueryTable("mv_user").All(&users)
	//for _, item := range users{
	//	fmt.Println(item)
	//	fmt.Println(item.Id)
	//}
	//fmt.Println(result)
	//conf := map[string]interface{}{}
	//conf["num"] = 30
	//conf["offset"] = 10
	//conf["limit"] = true
	//mvList, _ := models.GetMvList(conf)
	//for _, item := range *mvList {
	//	fmt.Println(item.Url)
	//}

	c.Data["data"] = "我是数据"
	c.TplName = "index/index.html"
}



