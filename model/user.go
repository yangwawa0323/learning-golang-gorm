package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// golang struct field default value
type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString // golang struct , database table
}

type VipUser struct {
	User     User `gorm:"embedded"`
	PostCode string
}

type Employee struct {
	User         User `gorm:"embedded"`
	SocialNumber string
}
