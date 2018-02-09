package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"log"
)

func dbOperate(){
	db,err := sql.Open("mysql",
		beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":"+beego.AppConfig.String("mysqlport")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8",)
	if err!=nil{
		log.Println(err)
	}
	defer db.Close()
	}