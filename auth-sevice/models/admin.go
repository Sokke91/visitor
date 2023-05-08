package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID       int    `gorm:"primaryKey; not null"`
	Username string `gorm:"size:255;not null"`
	Password string `gorm:"size:255; not null"`
}
