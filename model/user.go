package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	// IMPORANT!!!
	CompanyID  int // belong-to  foreighKey
	Company    Company
	CreditCard CreditCard
}

type Company struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID int // XxxID   Xxx struct name ,ID field
}
