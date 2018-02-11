package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)
/*
CREATE TABLE `news_wechat` (
  `newid` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '资讯id',
  `cateid` TINYINT(2) UNSIGNED NOT NULL DEFAULT '0' COMMENT '资讯分类,1:最新活动,2:微信营销,3:微商资讯,4:品牌资讯',
  `title` VARCHAR(100) NOT NULL COMMENT '资讯标题',
  `abstract` LONGTEXT NOT NULL COMMENT '资讯摘要',
  `show_url` VARCHAR(100) NOT NULL COMMENT '封面连接地址',
  `content` LONGTEXT NOT NULL COMMENT '资讯内容详情',
  `state` TINYINT(2) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态 1正常 ',
  `pubtime` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发布时间',
  `addtime` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '采集时间',
  `url` VARCHAR(100) NOT NULL COMMENT '采集来源的url',
  PRIMARY KEY (`newid`)
) ENGINE=INNODB AUTO_INCREMENT=393 DEFAULT CHARSET=utf8 COMMENT='微信新闻资讯'
*/
type News_wechat struct {
	newid   int  `orm:"auto"`
	cateid   int8
	title   string `orm:"size(100)"`
	abstract  string `orm:"type(text)"`
	show_url  string `orm:size(100)`
	content   string `orm:"type(text)"`
	state   int8
	pubtime   int
	addtime   int
	url string `orm:"size(100)"`
}

func init(){
	db_conn_str:=beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":"+beego.AppConfig.String("mysqlport")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8"
	orm.RegisterModel(new(News_wechat))
	orm.RegisterDataBase("wechat","mysql",db_conn_str)
	orm.RunSyncdb("wechat",false,true)
}
