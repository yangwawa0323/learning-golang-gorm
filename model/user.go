package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	// IMPORANT!!!
	CompanyID int // belong-to  foreighKey
	Company   Company
}

type Company struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string
}
