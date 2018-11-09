package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"unique;not null"`
}

type Category struct {
	gorm.Model
	Name    string `gorm:"unique;not null"`
	Memo    string `sql:"text;"`
	Balance int

	UserID uint
	User   *User

	Expenses []Expense
}
type Account struct {
	gorm.Model
	Name   string `gorm:"unique;not null"`
	Memo   string `sql:"text;"`
	Budget int

	UserID uint
	User   *User

	Expenses []Expense
}

type Expense struct {
	gorm.Model
	UseDate  time.Time
	Amount   int
	Location string
	Memo     string `sql:"text;"`

	Account   Account
	AccountID int

	Category   Category
	CategoryID int

	UserID uint
	User   *User
}
