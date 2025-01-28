package models

import (
    "fmt"
    "gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	UserId uint
	Number string
}

func (creditCard *CreditCard) Println() {
    fmt.Println("\n---CREDIT CARD---")
    fmt.Printf("ID: %d\n", creditCard.ID)
    fmt.Printf("UserId: %d\n", creditCard.UserId)
    fmt.Printf("Number: %s\n", creditCard.Number)
    fmt.Println()
}
