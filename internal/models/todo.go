package models

import "gorm.io/gorm"

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})
}
