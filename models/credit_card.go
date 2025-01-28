package models

import (
    "gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	UserId uint
	Number string
}
