package models

import (
    "errors"
    "fmt"
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    Role string
    Name string
    Age uint
    Birthday time.Time
    CreditCard CreditCard
}

func (user *User) BeforeSave(db *gorm.DB) error {
    fmt.Println("User BeforeSave")

    return nil
}

func (user *User) BeforeCreate(db *gorm.DB) error {
    fmt.Println("User BeforeCreate")

    if user.Role == "Admin" {
        return errors.New("invalid role")
    }

    return nil
}

func (user *User) AfterCreate(db *gorm.DB) error {
    fmt.Println("User AfterCreate")

    return nil
}

func (user *User) AfterSave(db *gorm.DB) error {
    fmt.Println("User AfterSave")

    return nil
}
