package dao

import (
	"duuit/models/request"
	"duuit/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type User struct {
	UID string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	request.UserRequest
	Goals []Goal
}

func (u *User) GetUser(db *gorm.DB) User {
	result := db.
		Preload("Goals.Tracking").
		Preload(clause.Associations).
		First(&u, u.UID)
	utils.LogDBError(result)
	return *u
}

func (u *User) CreateUser(db *gorm.DB) User {
	result := db.Create(&u)
	utils.LogDBError(result)
	return *u
}

func (u *User) UpdateUser(db *gorm.DB, newOb *User) User {
	result := db.Model(&u).Updates(newOb)
	utils.LogDBError(result)
	return *u
}

func (u *User) DeleteUser(db *gorm.DB) {
	result := db.Delete(&u, u.UID)
	utils.LogDBError(result)
}