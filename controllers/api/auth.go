package api

import (
	. "agora_blog/models"
	"crypto/md5"
	"encoding/hex"
	. "agora_blog/controllers/library"
	"agora_blog/controllers"
)

type json struct {
	code int
	msg  string
	data string
}

type AuthController struct {
	controllers.BaseController
	//Functions
}

func (this *AuthController) Get() {

	this.TplName = "app/index.html"
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
	this.Data["IsSignIn"] = false
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplName = "signin.html"
	} else {
		data := make(map[string]string)
		data["email"] = this.GetString("email")
		data["pass"] = this.GetString("pass")
		data["repass"] = this.GetString("repass")
		data["name"] = this.GetString("name")
		data["telephone"] = this.GetString("telephone")

		for k, v := range data {
			if len(v) == 0 {
				this.Data["json"] = Json{
					Code: "0",
					Msg:  k + " is required",
					Data: "false",
				}
				this.ServeJSON()
				break
			}
		}
		if data["pass"] != data["repass"] {
			this.Data["json"] = Json{
				Code: "0",
				Msg:  "pass must equal to repass",
				Data: "false",
			}
			this.ServeJSON()
			return
		}
		// md5 加密
		h := md5.New()
		h.Write([]byte(data["pass"]))
		data["pass"] = hex.EncodeToString(h.Sum(nil))

		delete(data, "repass")
		data["age"] = this.GetString("age")
		data["sex"] = this.GetString("sex")
		data["address"] = this.GetString("address")
		user := new(User)
		id := user.AddUser(data)

		if id == "0" {
			this.Data["json"] = Json{
				Code: "1",
				Msg:  "insert failed",
				Data: "false",
			}
			this.ServeJSON()
			return
		}
		this.Data["json"] = Json{
			Code: "1",
			Msg:  "success",
			Data: "id:" + id,
		}
		this.ServeJSON()
		return
	}
}

// @Title SignIn
// @Description signup
// @Param	email	formData	string	true	"The email for login"
// @Param	name	formData	string	true	"The name for login"
// @Param	pass	formData	string	true	"The pass for login"
// @Param	repass	formData	string	true	"The repass for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /signin [post]
func (this *AuthController) SignIn() {
	this.Data["IsSignIn"] = true
	if !this.IsAjax() {
		this.Layout = "layout.html"
		this.TplName = "signin.html"
	} else {
		 //session
		sess, _ := controllers.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		defer sess.SessionRelease(this.Ctx.ResponseWriter)

		data := make(map[string]string)
		data["email"] = this.GetString("email")
		data["pass"] = this.GetString("pass")

		for k, v := range data {
			if len(v) == 0 {
				this.Data["json"] = Json{
					Code: "0",
					Msg:  k + " is required",
					Data: "false",
				}
				this.ServeJSON()
				break
			}
		}
		user := new(User)
		userinfo, err := user.GetUserInfoByEmail(data["email"])

		if err != nil {
			this.Data["json"] = Json{
				Code: "0",
				Msg:  "invalid email address",
				Data: "false",
			}
			this.ServeJSON()
			return
		}

		// md5 加密
		h := md5.New()
		h.Write([]byte(data["pass"]))
		data["pass"] = hex.EncodeToString(h.Sum(nil))

		if userinfo.Pass != data["pass"] {
			this.Data["json"] = Json{
				Code: "0",
				Msg:  "passwprd is increditable",
				Data: "false",
			}
			this.ServeJSON()
			return
		}
		sess.Set("userinfo", userinfo)
		this.Data["json"] = Json{
			Code: "1",
			Msg:  "success",
			Data: "true",
		}
		this.ServeJSON()
		return
	}
}
