package routers

import (
	"newsWechat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/news", &controllers.TestController{})
    beego.Router("/request", &controllers.TestJsonController{})
}
