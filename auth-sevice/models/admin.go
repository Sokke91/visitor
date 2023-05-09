package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID             int    `gorm:"primaryKey; not null"`
	PersonalNumber string `gorm:"not null; size:255"`
	Name           string `gorm:"not null; size:255"`
	Prename        string `gorm:"not null; size:255"`
	Department     string `gorm:"not null; size:255"`
	Password       string `gorm:"size:255; not null"`
}
