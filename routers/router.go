package routers

import (
	"newsWechat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})//主页
    beego.Router("/api/news", &controllers.TestController{}) //新闻列表api
    beego.Router("/news/:newsid:int", &controllers.NewsDetailController{}) //新闻详情
    beego.Router("/request", &controllers.TestJsonController{})
}
