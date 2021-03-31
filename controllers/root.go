package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RootController struct {
	beego.Controller
}

// @Title Get
// @Description ping
// @Success 200
// @router / [get]
func (rc *RootController) Get() {
	rc.Data["json"] = "Duuit!!"
	rc.ServeJSON()
}
