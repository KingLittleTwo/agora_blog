package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["agora_blog/controllers/api:AuthController"] = append(beego.GlobalControllerRouter["agora_blog/controllers/api:AuthController"],
		beego.ControllerComments{
			Method: "SignIn",
			Router: `/signin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["agora_blog/controllers/api:AuthController"] = append(beego.GlobalControllerRouter["agora_blog/controllers/api:AuthController"],
		beego.ControllerComments{
			Method: "SignUp",
			Router: `/signup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
