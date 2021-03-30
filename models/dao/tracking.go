package dao

import (
	"duuit/models/request"
	"duuit/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tracking struct {
	gorm.Model
	request.TrackingRequest
	Journal Journal
}

func (t *Tracking) GetTracking(db *gorm.DB) Tracking {
	result := db.
		Preload(clause.Associations).
		First(&t, t.ID)
	utils.LogDBError(result)
	return *t
}

func (t *Tracking) CreateTracking(db *gorm.DB) Tracking {
	result := db.Create(&t)
	utils.LogDBError(result)
	return *t
}

func (t *Tracking) UpdateTracking(db *gorm.DB, newOb *Tracking) Tracking {
	result := db.Model(&t).Updates(newOb)
	utils.LogDBError(result)
	return *t
}

func (t *Tracking) DeleteTracking(db *gorm.DB) {
	result := db.Delete(&t, t.ID)
	utils.LogDBError(result)
}