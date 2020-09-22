package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func OpenDB() (*gorm.DB, error) {
	dsn := "host=db user=admin password=admin dbname=admin port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func greet(w http.ResponseWriter, r *http.Request) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	var user User
	stratTime := time.Now()
	fmt.Fprintf(w, "%s\n", stratTime)
	db.Take(&user)
	endTime := time.Now()
	fmt.Fprintf(w, "%s\n", endTime)
	fmt.Fprintf(w, "%s\n", endTime.Sub(stratTime))

	fmt.Println(user.Name)
}

func create(w http.ResponseWriter, r *http.Request) {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	var users []User
	stratTime1 := time.Now()
	fmt.Fprintf(w, "%s\n", stratTime1)
	for i := 1; i <= 10000; i++ {
		user := User{Name: "kenshin" + strconv.Itoa(i)}
		users = append(users, user)
	}
	endTime1 := time.Now()
	fmt.Fprintf(w, "%s\n", endTime1)
	fmt.Fprintf(w, "%s\n", endTime1.Sub(stratTime1))

	stratTime := time.Now()
	fmt.Fprintf(w, "%s\n", stratTime)
	db.Create(&users)
	endTime := time.Now()
	fmt.Fprintf(w, "%s\n", endTime)
	fmt.Fprintf(w, "%s\n", endTime.Sub(stratTime))
}
func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}
