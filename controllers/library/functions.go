package library

import (
	"github.com/astaxie/beego"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

type Json struct {
	Code string
	Msg  string
	Data string
}

type Functions struct {
	beego.Controller
}

func (this *Functions) ToJson(j Json) {
	this.Data["json"] = j
	this.ServeJSON()
	return
}

func (this *Functions) Test1() {
	fmt.Println("this is only a test 01")
}

func (this *Functions) MD5(b []byte) string {
	h := md5.New()
	h.Write(b)
	x := h.Sum(nil)
	y := make([]byte, 32)
	hex.Encode(y, x)
	return string(y)
}


