package web

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	fmt.Println("432423432")
	this.TplName = "index.html"
}
