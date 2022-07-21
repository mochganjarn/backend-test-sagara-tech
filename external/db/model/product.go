package model

import (
	"time"

	"github.com/mochganjarn/go-template-project/external/db"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Price     int
	Stock     int
	Filename  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Product) create(dbconn *db.Client) error {
	db := dbconn.DbConnection
	db.AutoMigrate(p)
	result := db.Create(p)

	if result.RowsAffected > 0 {
		return nil
	} else {
		return result.Error
	}
}
