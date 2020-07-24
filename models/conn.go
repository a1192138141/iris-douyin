package models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ims/logs"
)


//全局mysql 链接
var DbConn *gorm.DB

//初始化数据库链接
func InitDbConn()  {
	sourceConf , err :=config.NewConfig("ini","./conf/datasource.conf")
	if err != nil {
		fmt.Print(err)
	}
	//user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
	host :=sourceConf.String("db::host")
	password :=sourceConf.String("db::password")
	user :=sourceConf.String("db::user")
	dbname :=sourceConf.String("db::dbname")
	charset :=sourceConf.String("db::charset")
	connStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",user,password,host,dbname,charset)
	DbConn , err =gorm.Open("mysql",connStr)
	DbConn.LogMode(true)
	DbConn.SetLogger(logs.NewLogs())

	if err != nil {
		fmt.Print(err)
	}
	initTable()
}

//初始化table
func initTable()  {
	user :=DbConn.HasTable(&User{})
	if user == false {
		DbConn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
	video :=DbConn.HasTable(&Video{})
	if video == false{
		DbConn.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Video{})
	}
}