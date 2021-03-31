package controllers

import (
	"duuit-backend/models/dao"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
	beego.Controller
}

// @Title Ping
// @Description ping
// @Success 200
// @router /ping [get]
func (uc *UserController) Ping() {
	uc.Data["json"] = "Duuit Users!!"
	uc.ServeJSON()
}

// @Title Get
// @Description get user details
// @Param	id		path 	int	true		"The user id"
// @Success 200 {string} dao.User
// @Failure 403 body is empty
// @router /:id [get]
func (uc *UserController) Get() {
	uc.Data["json"] = "Duuit Users!!"
	uc.ServeJSON()
}

// @Title Put
// @Description update user details
// @Param	id		path 	int	true		"The user id"
// @Param course body request.UserRequest true "Include fields to be updated"
// @Success 200 {string} dao.User
// @Failure 403 body is empty
// @router /:id [put]
func (uc *UserController) Put() {
	var ob dao.User
	var newOb dao.User
	ob.UID = uc.Ctx.Input.Param(":id")
	json.Unmarshal(uc.Ctx.Input.RequestBody, &newOb)
	uc.Data["json"] = ob.UpdateUser(uc.DB, &newOb)
	uc.ServeJSON()
}

// @Title Post
// @Description add new user
// @Param course body request.UserRequest true "The user object"
// @Success 200 {user} dao.User
// @router / [post]
func (uc *UserController) Post() {
	var ob dao.User
	json.Unmarshal(uc.Ctx.Input.RequestBody, &ob)
	uc.Data["json"] = ob.CreateUser(uc.DB)
	uc.ServeJSON()
}

// @Title Delete
// @Description remove user
// @Param	id		path 	int	true		"The user id"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /:id [delete]
func (uc *UserController) Delete() {
	var ob dao.User
	ob.UID = uc.Ctx.Input.Param(":id")
	ob.DeleteUser(uc.DB)
	uc.Finish()
}




