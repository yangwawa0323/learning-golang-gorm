package model

import (
	"gorm.io/gorm"
)

// Capitalized 1st letter
type Student struct {
	gorm.Model
	Name      string
	Age       int
	ClassName string // snake   class_name
}
