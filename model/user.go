package model

import "gorm.io/gorm"

// golang struct field default value
type User struct {
	// ID uint //   `ID` field primary key
	// // ID 		     uint   `gorm:"primarykey"`
	// Deleted   gorm.DeletedAt
	gorm.Model
	Name      string
	Email     *string
	Age       uint8
	Languages []Language `gorm:"many2many:user_languages"`
	// many2many:<the future auto migrate generate intermediated table>
}

type Language struct {
	gorm.Model
	Name string
}
