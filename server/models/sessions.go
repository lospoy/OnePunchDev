package models

import "gorm.io/gorm"

type Timespan struct {
	Start			*int64		`json:"start"`
	End				*int64		`json:"end"`
	Duration	*int64		`json:"duration"`
}

type Code struct {
	Time			Timespan	`json:"timespan"`
}

type Sessions struct {
	ID				uint			`gorm:"primary key;autoIncrement" json:"id"`
	Code			Code			`json:"code"`
	// Workout		*string		`json:"workout"`
	// Rest			*string		`json:"rest"`
}

func MigrateSessions(db *gorm.DB) error {
	err := db.AutoMigrate(&Sessions{})
	return err
}