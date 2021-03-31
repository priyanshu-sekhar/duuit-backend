package controllers

import (
	"duuit-backend/models/dao"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

// Operations about goals
type GoalController struct {
	DB *gorm.DB
	beego.Controller
}

// @Title CreateGoal
// @Description create goal
// @Param	body		body 	request.GoalRequest	true		"body for goal content"
// @Success 200 {int} dao.Goal
// @Failure 403 body is empty
// @router / [post]
func (gc *GoalController) Post()  {
	var ob dao.Goal
	json.Unmarshal(gc.Ctx.Input.RequestBody, &ob)
	gc.Data["json"] = ob.CreateGoal(gc.DB)
	gc.ServeJSON()
}

// @Title Delete
// @Description delete goal
// @Param	id		path 	int	true		"The comment id"
// @Success 200 {string} dao.Reply
// @Failure 403 body is empty
// @routers /:id [delete]
func (gc *GoalController) Delete()  {
	var ob dao.User
	ob.UID = gc.Ctx.Input.Param(":id")
	ob.DeleteUser(gc.DB)
	gc.Finish()
}

// @Title AddGoalTracking
// @Description add goal tracking
// @Param	body		body 	request.TrackingRequest	true "body for tracking content"
// @Success 200 {int} dao.Tracking
// @Failure 403 body is empty
// @router /:id/tracking [post]
func (gc *GoalController) AddGoalTracking()  {
	var ob dao.Tracking
	json.Unmarshal(gc.Ctx.Input.RequestBody, &ob)
	gc.Data["json"] = ob.CreateTracking(gc.DB)
	gc.ServeJSON()
}