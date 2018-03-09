package api

import (
	. "agora_blog/models"
	//"agora_blog/controllers"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

type json struct {
	code int
	msg  string
	data string
}

type AuthController struct {
	//controllers.BaseController
	beego.Controller
}

// @Title Get
// @Description get all
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [get]
func (this *AuthController) Get() {
	a := map[string]string{
		"a": "bbb",
		"b": "ccc",
	}
	this.Data["json"] = a
	this.ServeJSON()
	return
}

// @Title SignUp
// @Description signup
// @Param	email	formData	string	true	"The email for login"
// @Param	name	formData	string	true	"The name for login"
// @Param	pass	formData	string	true	"The pass for login"
// @Param	repass	formData	string	true	"The repass for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /signup [post]
func (this *AuthController) SignUp() {
	data := make(map[string]string)
	data["email"] = this.GetString("email")
	data["name"] = this.GetString("name")
	data["pass"] = this.GetString("pass")
	data["repass"] = this.GetString("repass")

	for k, v := range data {
		if len(v) == 0 {
			j := map[string]string{
				"code": "0",
				"msg":  k + " is required",
				"data": "false",
			}
			this.toJson(j)
			break
		}
	}
	if data["pass"] != data["repass"] {
		this.toJson(map[string]string{
			"code": "0",
			"msg":  "pass must equal to repass",
			"data": "false",
		})
	}
	// md5 加密
	h := md5.New()
	h.Write([]byte(data["pass"]))
	data["pass"] = hex.EncodeToString(h.Sum(nil))

	delete(data, "repass")
	data["sex"] = this.GetString("sex")
	data["age"] = this.GetString("age")
	id := Add(data)
	if id == "0" {
		this.toJson(map[string]string{
			"code": "1",
			"msg":  "insert failed",
			"data": "false",
		})
	}
	this.toJson(map[string]string{
		"code": "1",
		"msg":  "success",
		"data": "id:" + id,
	})
}

func (this *AuthController) toJson(j map[string]string) {
	this.Data["json"] = j
	this.ServeJSON()
	return
}
