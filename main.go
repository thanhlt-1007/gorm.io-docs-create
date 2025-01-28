package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "time"
)

type User struct {
	gorm.Model
	Name string
	Age uint
	Birthday time.Time
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

func CreateUser(db *gorm.DB) {
    user := User {
        Name: "Jinzhu",
        Age: 18,
        Birthday: time.Now(),
    }
    result := db.Create(&user)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    fmt.Printf("ID: %d\n", user.ID)
    fmt.Printf("CreatedAt: %v\n", user.CreatedAt)
    fmt.Printf("UpdatedAt: %v\n", user.UpdatedAt)
    fmt.Printf("DeletedAt: %v\n", user.DeletedAt)
    fmt.Printf("Name: %s\n", user.Name)
    fmt.Printf("Age: %d\n", user.Age)
    fmt.Printf("Birthday: %v\n", user.Birthday)
}

func CreateUsers(db *gorm.DB) {
    users := []User {
        {
            Name: "Jinzhu 1",
        },
        {
            Name: "Jinzhu 2",
        },
    }
    result := db.Create(&users)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    for _, user := range users {
        fmt.Printf("ID: %d\n", user.ID)
        fmt.Printf("CreatedAt: %v\n", user.CreatedAt)
        fmt.Printf("UpdatedAt: %v\n", user.UpdatedAt)
        fmt.Printf("DeletedAt: %v\n", user.DeletedAt)
        fmt.Printf("Name: %s\n", user.Name)
        fmt.Printf("Age: %d\n", user.Age)
        fmt.Printf("Birthday: %v\n", user.Birthday)
    }
}

func CreateInBatchesUsers(db *gorm.DB) {
    users := []User {
        {
            Name: "Jinzhu 1",
        },
        {
            Name: "Jinzhu 2",
        },
    }
    result := db.CreateInBatches(users, 100)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    for _, user := range users {
        fmt.Printf("ID: %d\n", user.ID)
        fmt.Printf("CreatedAt: %v\n", user.CreatedAt)
        fmt.Printf("UpdatedAt: %v\n", user.UpdatedAt)
        fmt.Printf("DeletedAt: %v\n", user.DeletedAt)
        fmt.Printf("Name: %s\n", user.Name)
        fmt.Printf("Age: %d\n", user.Age)
        fmt.Printf("Birthday: %v\n", user.Birthday)
    }
}

func main() {
    db := InitDb()
    CreateUser(db)
    CreateUsers(db)
    CreateInBatchesUsers(db)
}
