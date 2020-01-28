package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nama_lengkap string
	Username string `gorm:"unique_index"`
	Password string
	Photo string `sql:"type:text;"`
}