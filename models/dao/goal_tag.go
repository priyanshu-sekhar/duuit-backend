package dao

import (
	"duuit-backend/models/request"
	"duuit-backend/utils"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	request.TagRequest
	Goals []*Goal `gorm:"many2many:goal_tags;"`
}

func (gt *Tag) DeleteGoalTag(db *gorm.DB) {
	result := db.Delete(&gt, gt.Tag)
	utils.LogDBError(result)
}
