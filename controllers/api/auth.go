package api

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Email   string  `orm:"size(50)"`
	Name    string  `orm:"size(50)"`
	Pass    string  `orm:"size(100)"`
	Profile Profile `orm:"rel(one)"`
}

type Profile struct {
	Id  int
	Sex int8
	Age int16
}

type AuthController struct {
	beego.Controller
}

func init() {
	// register model
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Profile))
}

// @Title Get
// @Description get all
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [get]
func (this *AuthController) Get() {

}
