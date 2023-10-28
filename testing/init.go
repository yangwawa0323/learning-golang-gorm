package testing

import (
	"log"

	"github.com/yangwawa0323/learning-golang-gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	// db, err = gorm.Open(sqlite.Open("demo.sqlite"), &gorm.Config{})
	var dsn = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect database.")
	}

	db.AutoMigrate(&model.Student{},
		&model.Product{},
		&model.User{},
		&model.VipUser{},
		&model.Employee{},
		&model.Book{},
		&model.Author{},
	)
}
