package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"fmt"
)
/*
*beego默认主页
*/
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.TplName = "index.html"
}
type TestJsonController struct{
	beego.Controller
}
func (this *TestJsonController) Get(){
	fmt.Println(beego.AppConfig.String("mysqluser"))
	data := make(map[string]interface{})
	data["code"] = 0
	data["msg"] = "success"
	this.Data["json"]=data
	time.Sleep(time.Second*2)
	this.ServeJSON()
}

