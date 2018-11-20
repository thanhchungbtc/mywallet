package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type User struct {
	Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Category struct {
	Model
	Name      string `gorm:"unique;not null" json:"name"`
	AvatarURL string `json:"avatar_url"`
	Memo      string `sql:"text;" json:"memo"`
	Budget    int    `json:"budget"`

	UserID uint  `gorm:"not null" json:"-"`
	User   *User `json:"-"`

	Expenses []Expense `json:"-"`
}
type Account struct {
	Model
	Name      string `gorm:"unique;not null" json:"name"`
	AvatarURL string `json:"avatar_url"`
	Memo      string `sql:"text;" json:"memo"`
	Balance   int    `json:"balance"`

	UserID uint  `gorm:"not null" json:"-"`
	User   *User `json:"-"`

	Expenses []Expense `json:"-"`
}

type Expense struct {
	Model
	UseDate  time.Time `json:"use_date"`
	Amount   int       `json:"amount"`
	Location string    `json:"location"`
	Memo     string    `sql:"text;" json:"memo"`

	Account   Account `json:"-"`
	AccountID int     `json:"account_id"`

	Category   Category `json:"-"`
	CategoryID int      `json:"category_id"`

	UserID uint  `gorm:"not null" json:"-"`
	User   *User `json:"-"`
}
