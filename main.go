package main

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-create/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "time"
)

func InitDb() *gorm.DB {
    db, err := gorm.Open(
        sqlite.Open("db.sqlite"),
        &gorm.Config{},
    )

    if err != nil {
        fmt.Printf("gorm.Open error %v\n", err)
        panic(err)
    }

    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.CreditCard{})

    return db
}

func CreateAdminUser(db *gorm.DB) {
    fmt.Println("CreateAdminUser")

    user := models.User {
        Name: "Jinzhu",
        Role: "Admin",
        Age: 18,
        Birthday: time.Now(),
    }
    result := db.Create(&user)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    user.Println()
}

func CreateAdminUserSkipHooks(db *gorm.DB) {
    fmt.Println("CreateAdminUserSkipHooks")

    user := models.User {
        Name: "Jinzhu",
        Role: "Admin",
        Age: 18,
        Birthday: time.Now(),
    }
    result := db.Session(&gorm.Session{
        SkipHooks: true,
    }).Create(&user)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    user.Println()
}

func CreateUserWithCreditCard(db *gorm.DB) {
    fmt.Println("CreateUserWithCreditCard")

    user := models.User {
        Name: "jinzhu",
        CreditCard: models.CreditCard {
            Number: "411111111111",
        },
    }
    result := db.Create(&user)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    user.Println()

    creditCard := user.CreditCard
    creditCard.Println()
}

func CreateUser(db *gorm.DB) {
    fmt.Println("CreateUser")

    user := models.User {
        Name: "Jinzhu",
        Age: 18,
        Birthday: time.Now(),
    }
    result := db.Create(&user)

    fmt.Printf("result.Error: %v\n", result.Error)
    fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)

    user.Println()
}

func CreateUsers(db *gorm.DB) {
    fmt.Println("CreateUsers")

    users := []models.User {
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
        user.Println()
    }
}

func CreateInBatchesUsers(db *gorm.DB) {
    fmt.Println("CreateInBatchesUsers")

    users := []models.User {
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
        user.Println()
    }
}

func main() {
    db := InitDb()
    // CreateAdminUser(db)
    // CreateAdminUserSkipHooks(db)
    // CreateUser(db)
    // CreateUsers(db)
    // CreateInBatchesUsers(db)

    CreateUserWithCreditCard(db)
}
