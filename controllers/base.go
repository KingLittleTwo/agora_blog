package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego/config"
)

type BaseController struct {
	beego.Controller
}

func init() {
	iniconf, err := config.NewConfig("ini", "app.conf")
	if err != nil {
		fmt.Println("err: ", err)
	}
	// set default database
	DbUser := iniconf.String("DbUser")
	DbPass := iniconf.String("DbPass")
	DbIP := iniconf.String("DbIP")
	DbName := iniconf.String("DbName")
	DbCharset := iniconf.String("DbCharset")


	str := DbUser + ":" + DbPass + "@tcp(" + DbIP + ")" + "/" + DbName + "?charset" + DbCharset
	fmt.Println(str)

	orm.RegisterDataBase("default", "mysql", str, 30)
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/codeagora?charset=utf8", 30)
	// create table
	orm.RunSyncdb("default", false, true)
}
