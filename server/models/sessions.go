package models

import "gorm.io/gorm"

type Timespan struct {
	Start					*int64		`json:"start"`
	End						*int64		`json:"end"`
	CodeID				int
	Code					Code
}

type Code struct {
	ID						int
	SessionsID		int
	Sessions			Sessions
	Times					[]Timespan	`json:"times"`
}

type Sessions struct {
	ID						int			`gorm:"autoIncrement" json:"id"`
	Coding				[]Code
	// Workout		*string		`json:"workout"`
	// Rest				*string		`json:"rest"`
}

func MigrateSessions(db *gorm.DB) error {
	err := db.AutoMigrate(&Sessions{}, &Code{}, &Timespan{})
	return err
}