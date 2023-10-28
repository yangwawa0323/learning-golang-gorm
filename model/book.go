package model

import "gorm.io/gorm"

// normal form
type Book struct {
	gorm.Model
	BookName string `gorm:"type:VARCHAR(200)"`
	Author   Author `gorm:"embedded"`
}

type Author struct {
	AuthorName string
	Email      string
}
