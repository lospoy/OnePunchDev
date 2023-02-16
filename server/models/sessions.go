package models

import "gorm.io/gorm"

type Sessions struct {
	ID				uint			`gorm:"primary key;autoIncrement" json:"id"`
	Code			*string		`json:"code"`
	Workout		*string		`json:"workout"`
	Rest			*string		`json:"rest"`
}

func MigrateSessions(db *gorm.DB) error {
	err := db.AutoMigrate(&Sessions{})
	return err
}