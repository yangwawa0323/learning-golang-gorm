package model

import (
	"database/sql"
	"time"
)

// golang struct field default value
type User struct {
	ID uint //   `ID` field primary key
	// ID 		     uint   `gorm:"primarykey"`
	Name         string  `gorm:"type:VARCHAR(250);not null"`
	Email        *string `gorm:"not null"`
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString // golang struct , database table
	Expired      sql.NullBool   `gorm:"default:false"`
}

type VipUser struct {
	User     User `gorm:"embedded"`
	PostCode string
}

type Employee struct {
	User         User `gorm:"embedded"`
	SocialNumber string
}
