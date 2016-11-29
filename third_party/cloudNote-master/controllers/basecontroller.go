package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

var Ob map[string]interface{}

func (this *BaseController) Prepare() {

}
