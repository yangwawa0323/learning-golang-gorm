package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/yangwawa0323/learning-golang-gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // global

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are welcome to cloudclass gorm training course")
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello you are request %s", r.URL))
}

// student/create?name=yangkun&age=49
func student_create(w http.ResponseWriter, r *http.Request) {
	// CS get user query parameters
	name := r.URL.Query().Get("name")
	ageParam := r.URL.Query().Get("age")
	age, _ := strconv.Atoi(ageParam)

	// instance
	s := model.Student{
		Name: name,
		Age:  age,
	}

	// insert into xxx values(xxx,zzzz,yyy)
	// db.Create
	db.Create(&s)

	io.WriteString(w, fmt.Sprintf("success create user : %s", name))
}

// student/find?id=3
func student_find(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	var student model.Student

	// select * from student where id =
	db.First(&student, id)

	output := fmt.Sprintf("student: %s , age: %d , classname: %s",
		student.Name, student.Age, student.ClassName,
	)
	io.WriteString(w, output)
}

// student/update?id=2&age=24
func student_update(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id") // default string
	id, _ := strconv.Atoi(idParam)

	ageParam := r.URL.Query().Get("age")
	age, _ := strconv.Atoi(ageParam)

	var student model.Student
	db.First(&student, id)

	// update student set age = xxx where id = 3
	student.Age = age
	db.Save(&student)

	io.WriteString(w, "update successfully.")
}

// student/delete?id=2
func student_delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id") // default string
	id, _ := strconv.Atoi(idParam)

	var student model.Student
	db.Delete(&student, id)

	io.WriteString(w, "delete successfully.")
}

func init() {
	var err error
	// db, err = gorm.Open(sqlite.Open("demo.sqlite"), &gorm.Config{})
	var dsn = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to connect database.")
	}

	db.AutoMigrate(&model.Student{}, &model.Product{})
}

func main() {

	// multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/hello", hello)

	//
	mux.HandleFunc("/student/create", student_create)
	mux.HandleFunc("/student/find", student_find)
	mux.HandleFunc("/student/update", student_update)
	mux.HandleFunc("/student/delete", student_delete)

	idleConnsClosed := make(chan struct{})
	// http.Server
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// gorouting
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("HTTP server ListenAndServer: %v", err)
		}
		log.Printf("HTTP server starting...\n")
	}() // IIFE

	<-idleConnsClosed
}
