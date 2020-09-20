package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   int
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

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
