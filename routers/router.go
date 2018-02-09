package routers

import (
	"newWechat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.TestController{})
    beego.Router("/request", &controllers.TestJsonController{})
}
