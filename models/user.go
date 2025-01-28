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

func (user *User) Println() {
    fmt.Println("\n---USER---")
    fmt.Printf("ID: %d\n", user.ID)
    fmt.Printf("CreatedAt: %v\n", user.CreatedAt)
    fmt.Printf("UpdatedAt: %v\n", user.UpdatedAt)
    fmt.Printf("DeletedAt: %v\n", user.DeletedAt)
    fmt.Printf("Role: %s\n", user.Role)
    fmt.Printf("Name: %s\n", user.Name)
    fmt.Printf("Age: %d\n", user.Age)
    fmt.Printf("Birthday: %v\n", user.Birthday)
    fmt.Println()
}

// Hooks
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
