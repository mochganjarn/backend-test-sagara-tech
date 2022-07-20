package model

import (
	"time"

	"github.com/mochganjarn/go-template-project/external/db"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) create(dbconn *db.Client) error {
	db := dbconn.DbConnection
	db.AutoMigrate(u)
	result := db.Create(u)

	if result.RowsAffected > 0 {
		return nil
	} else {
		return result.Error
	}
}

func (u *User) show(dbconn *db.Client) error {
	db := dbconn.DbConnection
	db.AutoMigrate(u)
	result := db.Where(u).First(u)

	if result.RowsAffected > 0 {
		return nil
	} else {
		return result.Error
	}
}

func (u *User) read() {

}

func (u *User) update() {

}

func (u *User) delete() {

}
