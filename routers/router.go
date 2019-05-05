package routers

import (
	"github.com/astaxie/beego"
	"github.com/yy/spain_movie/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.VideoController{})
	beego.AutoRouter(&controllers.PageController{})
}
