package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/session"
)

type BaseController struct {
	beego.Controller
}

var GlobalSessions *session.Manager

func init() {
	// 初始化数据库
	DbUser := beego.AppConfig.String("mysql::DbUser")
	DbPass := beego.AppConfig.String("mysql::DbPass")
	DbIP := beego.AppConfig.String("mysql::DbIP")
	DbName := beego.AppConfig.String("mysql::DbName")
	DbCharset := beego.AppConfig.String("mysql::DbCharset")

	conn := DbUser + ":" + DbPass + "@tcp(" + DbIP + ")" + "/" + DbName + "?charset" + DbCharset

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", conn, 30)

	// create table if not exists
	orm.RunSyncdb("default", false, true)
	orm.Debug = true

	//	初始化session
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()
}
