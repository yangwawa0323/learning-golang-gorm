package model

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// golang struct field default value
type User struct {
	// ID uint //   `ID` field primary key
	// // ID 		     uint   `gorm:"primarykey"`
	// Deleted   gorm.DeletedAt
	gorm.Model
	Name  string
	Email *string
	Age   uint8
	UUID  string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()
	if u.Age > 65 {
		err = errors.New("user is too old")
	}
	return
}
