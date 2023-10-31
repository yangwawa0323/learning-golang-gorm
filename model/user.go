package model

<<<<<<< HEAD
<<<<<<< HEAD
import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)
=======
import "gorm.io/gorm"
>>>>>>> has-many

type User struct {
	gorm.Model
<<<<<<< HEAD
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString // golang struct , database table
=======
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
>>>>>>> delete-data
=======
	Name string
	// IMPORANT!!!
	CompanyID   int // belong-to  foreighKey
	Company     Company
	CreditCards []CreditCard //preload owner field
>>>>>>> has-many
}

type CreditCard struct { // credit_card(s)
	gorm.Model
	Number string
	UserID int // XxxID   Xxx struct name ,ID field ;
	// table field  xxx_id
}

type Company struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string
}
