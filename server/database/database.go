package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DB struct {
	*gorm.DB
}

func New() (*DB, error) {
	db, err := gorm.Open("postgres", "postgres://thanhchungbui@localhost/mywallet?sslmode=disable")
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.LogMode(gin.Mode() == gin.DebugMode)
	appDB := &DB{DB: db}
	return appDB, nil
}
