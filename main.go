package main

import (
	_ "agora_blog/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// set default database
	DbUser := beego.AppConfig.String("mysql::DbUser")
	DbPass := beego.AppConfig.String("mysql::DbPass")
	DbIP := beego.AppConfig.String("mysql::DbIP")
	DbName := beego.AppConfig.String("mysql::DbName")
	DbCharset := beego.AppConfig.String("mysql::DbCharset")

	conn := DbUser + ":" + DbPass + "@tcp(" + DbIP + ")" + "/" + DbName + "?charset" + DbCharset

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", conn, 30)

	// create table if not exists
	orm.RunSyncdb("default", false, false)
	orm.Debug = true
}

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	beego.Run()
}
