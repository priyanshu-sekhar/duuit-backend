package dao

import (
	"duuit/models/request"
	"duuit/utils"
	"gorm.io/gorm"
)

type GoalTag struct {
	request.GoalTagRequest
	Goals []*Goal
}

func (gt *GoalTag) DeleteGoalTag(db *gorm.DB) {
	result := db.Delete(&gt, gt.Tag)
	utils.LogDBError(result)
}
