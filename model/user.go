package model

import "gorm.io/gorm"

// golang struct field default value
type User struct {
	// ID uint //   `ID` field primary key
	// // ID 		     uint   `gorm:"primarykey"`
	// Deleted   gorm.DeletedAt
	gorm.Model
	Name  string
	Email *string
	Age   uint8
}

type VipUser struct {
	User     User `gorm:"embedded"`
	PostCode string
}

type Employee struct {
	User         User `gorm:"embedded"`
	SocialNumber string
}
