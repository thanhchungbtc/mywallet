package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
	Category Category
	User     User
	Expense  Expense
}

func New(driver, databaseURL string) (*DB, error) {
	db, err := gorm.Open(driver, databaseURL)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.LogMode(gin.Mode() == gin.DebugMode)

	s := &DB{
		DB:       db,
		Category: &category{db},
		User:     &user{db},
		Expense:  &expense{db},
	}

	return s, nil
}
