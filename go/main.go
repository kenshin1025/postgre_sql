package main

import (
	"fmt"
	"log"
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

func create(db *gorm.DB, executionTimes int) {
	var users []User
	for i := 1; i <= executionTimes; i++ {
		user := User{Name: "kenshin" + strconv.Itoa(i)}
		users = append(users, user)
	}
	stratTime := time.Now()
	db.Create(&users)
	endTime := time.Now()
	fmt.Println("create")
	fmt.Println(endTime.Sub(stratTime))
	var user User
	db.Take(&user)
	fmt.Println(user.Name)
}

func read(db *gorm.DB) {
	var users []User
	stratTime := time.Now()
	db.Find(&users)
	endTime := time.Now()
	fmt.Println("read")
	fmt.Println(endTime.Sub(stratTime))
	fmt.Println(len(users))
}

func update(db *gorm.DB, executionTimes int) {
	var IDs []int
	for i := 1; i <= executionTimes; i++ {
		IDs = append(IDs, i)
	}
	stratTime := time.Now()
	db.Model(User{}).Where("id IN ?", IDs).Updates(User{Name: "hello"})
	endTime := time.Now()
	fmt.Println("update")
	fmt.Println(endTime.Sub(stratTime))
	var user User
	db.Take(&user)
	fmt.Println(user.Name)
}

func delete(db *gorm.DB, executionTimes int) {
	var IDs []int
	for i := 1; i <= executionTimes; i++ {
		IDs = append(IDs, i)
	}
	stratTime := time.Now()
	db.Delete(&User{}, IDs)
	endTime := time.Now()
	fmt.Println("delete")
	fmt.Println(endTime.Sub(stratTime))
}

func main() {
	fmt.Println("openDB")
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("initDB")
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	executionTimes := 1
	create(db, executionTimes)
	fmt.Println("")
	read(db)
	fmt.Println("")
	fmt.Println("update")
	update(db, executionTimes)
	fmt.Println("")
	delete(db, executionTimes)
}
