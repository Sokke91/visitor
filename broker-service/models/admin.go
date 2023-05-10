package models

import (
	"broker/database"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	ID             int    `gorm:"primaryKey; not null"`
	PersonalNumber string `gorm:"not null; size:255"`
	Name           string `gorm:"not null; size:255"`
	Prename        string `gorm:"not null; size:255"`
	Department     string `gorm:"not null; size:255"`
	Password       string `gorm:"size:255; not null"`
}

func (admin *Admin) Save() (Admin, error) {
	err := database.DB.Create(&admin).Error
	if err != nil {
		return Admin{}, err
	}
	return *admin, nil
}

func getAdminById(id int) (Admin, error) {
	var admin Admin
	err := database.DB.Find(&admin, id).Error
	if err != nil {
		return Admin{}, err
	}
	return admin, nil
}
