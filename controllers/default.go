package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	"time"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"log"
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

func (self *TestController) Post() {
	param := self.Input()
	page :=param.Get("page")
	limit :=6
	//sqlOrm := orm.NewOrm()
	//sqlOrm.Using("wechat")
	//var lists []orm.ParamsList
	//res,err := sqlOrm.Raw("select * from news_wechat limit ?,?",page,limit).ValuesList(&lists)
	//if err==nil && res >0{
	//	fmt.Println(lists)
	//}
	db, err := sql.Open("mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":"+beego.AppConfig.String("mysqlport")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	//在这里进行一些数据库操作
	defer db.Close()
	rows, err := db.Query("select * from news_wechat limit ?,?",page,limit)
	for rows.Next() {
		fmt.Println(rows)
	}
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

