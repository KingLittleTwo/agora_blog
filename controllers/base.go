package controllers
//
//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	_ "github.com/go-sql-driver/mysql"
//	//"fmt"
//	"os"
//)
//
//type BaseController struct {
//	beego.Controller
//}
//
//func init() {
//	// set default database
//	DbUser := beego.AppConfig.String("mysql::DbUser")
//	DbPass := beego.AppConfig.String("mysql::DbPass")
//	DbIP := beego.AppConfig.String("mysql::DbIP")
//	DbName := beego.AppConfig.String("mysql::DbName")
//	DbCharset := beego.AppConfig.String("mysql::DbCharset")
//
//	beego.Info(DbName)
//	os.Stdout.WriteString(DbName)
//
//	conn := DbUser + ":" + DbPass + "@tcp(" + DbIP + ")" + "/" + DbName + "?charset" + DbCharset
//
//	orm.RegisterDriver("mysql", orm.DRMySQL)
//
//	orm.RegisterDataBase("default", "mysql", conn, 30)
//
//	orm.Debug = true
//}
