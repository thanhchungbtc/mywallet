package service

import (
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type ExpenseService struct {
	db *database.DB
}

func NewExpenseService(db *database.DB) *ExpenseService {
	return &ExpenseService{db}
}

func (s *ExpenseService) FindAll(query interface{}, args ...interface{}) ([]*model.Expense, error) {
	var objects []*model.Expense
	err := s.db.
		Where(query, args).
		Find(&objects).
		Error
	return objects, err
}

func (s *ExpenseService) FindOne(query interface{}, args ...interface{}) (*model.Expense, error) {
	var instance model.Expense
	err := s.db.
		Where(query, args).
		First(&instance).
		Error
	return &instance, err
}

func (s *ExpenseService) Save(data *model.Expense) error {
	return s.db.
		Save(data).
		Error
}

func (s *ExpenseService) Create(data *model.Expense) error {
	return s.db.
		Create(data).
		Error
}

func (s *ExpenseService) Delete(query interface{}, args ...interface{}) error {
	// permanently delete
	return s.db.
		Where(query, args).
		Unscoped().
		Delete(&model.Expense{}).
		Error
}
