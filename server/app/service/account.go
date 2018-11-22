package service

import (
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type AccountService struct {
	db *database.DB
}

func NewAccountService(db *database.DB) *AccountService {
	return &AccountService{db}
}

func (s *AccountService) FindAll(query interface{}, args ...interface{}) ([]*model.Account, error) {
	var objects []*model.Account
	err := s.db.
		Where(query, args).
		Find(&objects).
		Error
	return objects, err
}

func (s *AccountService) FindOne(query interface{}, args ...interface{}) (*model.Account, error) {
	var instance model.Account
	err := s.db.
		Where(query, args).
		First(&instance).
		Error
	return &instance, err
}

func (s *AccountService) Save(data *model.Account) error {
	return s.db.
		Save(data).
		Error
}

func (s *AccountService) Create(data *model.Account) error {
	return s.db.
		Create(data).
		Error
}

func (s *AccountService) Delete(query interface{}, args ...interface{}) error {
	// permanently delete
	return s.db.
		Where(query, args).
		Unscoped().
		Delete(&model.Account{}).
		Error
}
