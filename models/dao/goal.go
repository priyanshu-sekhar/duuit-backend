package dao

import (
	"duuit/models/common"
	"duuit/models/request"
	"duuit/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Goal struct {
	gorm.Model
	request.GoalRequest
	Status common.Status
	Tracking []Tracking
	Tags []*GoalTag
}

func (g *Goal) GetGoal(db *gorm.DB) Goal {
	result := db.
		Preload(clause.Associations).
		First(&g, g.ID)
	utils.LogDBError(result)
	return *g
}

func (g *Goal) CreateGoal(db *gorm.DB) Goal {
	result := db.Create(&g)
	utils.LogDBError(result)
	return *g
}

func (g *Goal) UpdateGoal(db *gorm.DB, newOb *Goal) Goal {
	result := db.Model(&g).Updates(newOb)
	utils.LogDBError(result)
	return *g
}

func (g *Goal) DeleteGoal(db *gorm.DB) {
	result := db.Delete(&g, g.ID)
	utils.LogDBError(result)
}
