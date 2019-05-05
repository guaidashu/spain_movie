package controllers

import (
	"fmt"
	"github.com/yy/spain_movie/models"
	"math/rand"
	"time"
)

type VideoController struct {
	BaseController
}

func (v *VideoController) Play() {
	//url := "https://r1---sn-ogul7n7s.c.drive.google.com/videoplayback?id=f613d896805ce857&itag=22&source=webdrive&requiressl=yes&mm=30&mn=sn-ogul7n7s&ms=nxu&mv=u&pl=21&sc=yes&ttl=transient&ei=ZWbBXLeQC5TtugKPt4CgBw&susc=dr&driveid=12Iz2tI2UD9att4GlcqX3zMZjw3q2Bai4&app=texmex&mime=video/mp4&dur=6001.824&lmt=1550443046742483&mt=1556178207&ip=198.13.58.90&ipbits=0&expire=1556192933&cp=QVNKWEJfUVlRRFhOOjNfa0lQU0t3b1ND&sparams=ip,ipbits,expire,id,itag,source,requiressl,mm,mn,ms,mv,pl,sc,ttl,ei,susc,driveid,app,mime,dur,lmt,cp&signature=ABBAED9332C53E1A4A6588EB9D12073B66117E23BA4A11D88B53CC94DFC74E82.A9591DDE2D98DED2C8C3D856E1996464B1412C014D95265C9CF249AE7ED54D2B&key=us0&cpn=cQvQZC0oofZQK3I3&c=WEB_EMBEDDED_PLAYER&cver=20190423"
	mv_id, _ := v.GetInt("mv_id")
	fmt.Println(mv_id)
	mv_content, _ := models.GetContentByParentId(mv_id)
	url := mv_content.VideoSrc
	mv_list, _ := models.GetMvListById(mv_id)
	rand.Seed(time.Now().Unix())
	recommend, _, _ := models.GetMvListByCategory(rand.Intn(24), 12)
	v.Data["PlayData"] = mv_list
	v.Data["PlayContent"] = mv_content
	v.Data["Recommend"] = recommend
	v.Data["Url"] = url
	v.Layout = "layout/default.html"
	v.TplName = "page/play.html"
}
