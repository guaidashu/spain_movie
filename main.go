package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yy/spain_movie/init"
	_ "github.com/yy/spain_movie/models"
	_ "github.com/yy/spain_movie/routers"
)

func main() {
	beego.Run()
}
