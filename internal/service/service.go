package service

import (
	"github.com/thanhchungbtc/mywallet/internal/database"
)

//type Service interface {
//	Category Category
//}
type Service struct {
	*database.DB
	Category Category
}

func New(DB *database.DB) *Service {
	s := &Service{
		DB:       DB,
		Category: &category{DB},
	}

	return s
}

type wrapper struct {
}
