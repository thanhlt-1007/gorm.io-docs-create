package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age uint
	Brithday time.Time
}

func InitDb() *gorm.DB {
    db, err := gorm.Open(
        sqlite.Open("db.sqlite"),
        &gorm.Config{},
    )

    if err != nil {
        fmt.Printf("gorm.Open error %v\n", err)
        panic(err)
    }

    db.AutoMigrate(&User{})

    return db
}

func main() {
    InitDb()
}
