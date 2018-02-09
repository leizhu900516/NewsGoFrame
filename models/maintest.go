package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_"os/user"
)

// Model Struct
type News_wechat struct {
	Newid   int  `orm:"auto"`
	Cateid   int8
	Title   string `orm:"size(100)"`
	Abstract  string `orm:"type(text)"`
	Show_url  string `orm:size(100)`
	Content   string `orm:"type(text)"`
	State   int8
	Pubtime   int
	Addtime   int
	Url string `orm:"size(100)"`
}
func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "wordpress_hc:gc895316@tcp(113.10.195.169:3306)/wechat?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(News_wechat))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	fmt.Println(o)
	//user := User{Name: "slene"}

	// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	//user.Name = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	//u := User{Id: user.Id}
	//err := o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//
	//// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}