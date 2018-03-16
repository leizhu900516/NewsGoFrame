package routers

import (
	"newsWechat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})//主页
    beego.Router("/api/news", &controllers.TestController{}) //新闻列表api
    beego.Router("/news/:newsid:int.html", &controllers.NewsDetailController{}) //新闻详情
    beego.Router("/request", &controllers.TestJsonController{})
    beego.Router("/search", &controllers.TestJsonController{})
    beego.Router("/news/type/:newstype:int", &controllers.NewsTypeListController{})//新闻分类列表页
    beego.Router("/api/news/hot", &controllers.HotNewsApiController{})//热门新闻接口
}
