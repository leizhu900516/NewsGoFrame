package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	"time"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"reflect"
	"encoding/json"
	"log"
	"strconv"
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
	c.Data["Request_url"] = "/"
	c.TplName = "index.html"
}

type TestController struct {
	beego.Controller
}

func (self *TestController) Post() {
	/*新闻列表api*/
	returndata := make(map[string]interface{})
	var ob map[string]int
	json.Unmarshal(self.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob["page"])
	//fmt.Println(string(ob))
	//param := self.Input()
	page :=ob["page"]
	limit :=ob["limit"]
	db, err := sql.Open("mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":"+beego.AppConfig.String("mysqlport")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8")
	fmt.Println(reflect.TypeOf(db))
	if err != nil {
		returndata["code"] = 1
		returndata["msg"] = "success"
		returndata["data"]=""
		self.Data["json"]=returndata
		self.ServeJSON()
	}
	//在这里进行一些数据库操作
	sql := fmt.Sprintf("select * from newslist limit %d,%d",(page-1)*8,limit)
	result := selectSqlData(db,sql)
	returndata["code"] = 0
	returndata["msg"] = "success"
	returndata["data"]=result
	self.Data["json"]=returndata
	self.ServeJSON()
}
/*
*新闻详情
*/
type NewsDetailController struct {
	beego.Controller
}
func (self *NewsDetailController) Get(){
	fmt.Println(self.Ctx.Input.Param(":newsid"))
	newsid,err := strconv.Atoi(self.Ctx.Input.Param(":newsid"))
	fmt.Println(reflect.TypeOf(newsid))
	db, err := sql.Open("mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":"+beego.AppConfig.String("mysqlport")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	var (
		cateid int
		title string
		abstract string
		content string
		addtime int64
		shownumber int
		previous string
		next string
	)
	err =db.QueryRow("select cateid,title,abstract,content,addtime,shownumber from newslist where newsid=?",newsid).Scan(&cateid,&title,&abstract,&content,&addtime,&shownumber)
	if err!=nil{
		fmt.Println(err)
	}
	err = db.QueryRow("select title from newslist where newsid=?",int(newsid)+1).Scan(&next)
	err = db.QueryRow("select title from newslist where newsid=?",int(newsid)-1).Scan(&previous)
	db.Exec("update newslist set shownumber=shownumber+1 where newsid=?",newsid)
	self.Data["newsid"] = newsid
	self.Data["cateid"] = cateid
	self.Data["title"] = title
	self.Data["previous"] = previous
	self.Data["next"] = next
	self.Data["abstract"] = abstract
	self.Data["content"] = content
	self.Data["shownumber"] = shownumber
	self.Data["addtime"] = time.Unix(addtime,0).Format("2006-01-02 03:04:05 PM")
	self.TplName = "news-detail.html"
}
/*
*新闻分类接口
*/
type NewsTypeController struct{
	beego.Controller
}
func (self *NewsTypeController) Get(){
	data := make(map[string]interface{})
	param := make(map[string]int)
	json.Unmarshal(self.Ctx.Input.RequestBody,&param)
	newstype := param["newstype"]
	limit := param["limit"]
	page := param["page"]
	fmt.Println(newstype)
	selectsql := fmt.Sprintf("select * from newslist where cateid = %d limit %d,%d",newstype,(page-1)*limit,limit)
	fmt.Println(selectsql)
	db,err :=sql.Open("mysql",beego.AppConfig.String("mysqlurl"))
	checkErr(err)
	result :=selectSqlData(db,selectsql)
	data["code"]=0
	data["msg"]=""
	data["data"]=result
	self.Data["json"]=data
	self.ServeJSON()

}
/*
*新闻分类列表
*/
type NewsTypeListController struct{
	beego.Controller
}
func (self *NewsTypeListController) Get(){
	data := make(map[string]interface{})
	newstype := self.Ctx.Input.Param(":newsid")
	cate,err:=strconv.Atoi(newstype)
	fmt.Println(newstype)
	selectsql := fmt.Sprintf("select * from newslist where cateid = %d limit %d,%d",cate,0,10)
	fmt.Println(selectsql)
	db,err :=sql.Open("mysql",beego.AppConfig.String("mysqlurl"))
	checkErr(err)
	result :=selectSqlData(db,selectsql)
	data["code"]=0
	data["msg"]=""
	data["data"]=result
	self.Data["json"]=data
	self.TplName="newslist.html"

}
/*
*测试
*/
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
/*热门新闻api*/
type HotNewsApiController struct {
	beego.Controller
}
func (self *HotNewsApiController) Get(){
	data := make(map[string]interface{})
	limit:=5
	page,err:=strconv.Atoi(self.GetString("page"))
	checkErr(err)
	selectsql:=fmt.Sprintf("SELECT newsid,title FROM newslist ORDER BY shownumber LIMIT %d,%d",page,limit)
	fmt.Println(selectsql)
	db,err :=sql.Open("mysql",beego.AppConfig.String("mysqlurl"))
	checkErr(err)
	result :=selectSqlData(db,selectsql)
	data["code"]=0
	data["data"]=result
	self.Data["json"]=data
	self.ServeJSON()
}

func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}

func selectSqlData(db *sql.DB,sql string) map[int]map[string]string{
	/*sql查询返回数据*/
	defer db.Close()
	fmt.Println(sql)
	rows2, err := db.Query(sql)
	if err!=nil{
		fmt.Println("xxxxx",err)
		return map[int]map[string]string{}
	}
	//返回所有列
	cols, _ := rows2.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//fmt.Printf(string(v))
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	return result
}

