package dao

import (
	"duuit/models/request"
	"duuit/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Journal struct {
	gorm.Model
	request.JournalRequest
}

func (j *Journal) GetJournal(db *gorm.DB) Journal {
	result := db.
		Preload(clause.Associations).
		First(&j, j.ID)
	utils.LogDBError(result)
	return *j
}

func (j *Journal) CreateJournal(db *gorm.DB) Journal {
	result := db.Create(&j)
	utils.LogDBError(result)
	return *j
}

func (j *Journal) UpdateJournal(db *gorm.DB, newOb *Journal) Journal {
	result := db.Model(&j).Updates(newOb)
	utils.LogDBError(result)
	return *j
}

func (j *Journal) DeleteJournal(db *gorm.DB) {
	result := db.Delete(&j, j.ID)
	utils.LogDBError(result)
}