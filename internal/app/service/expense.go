package service

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/thanhchungbtc/mywallet/internal/app/model"
	"github.com/thanhchungbtc/mywallet/internal/database"
)

type Expense interface {
	All() ([]*model.Expense, error)
	ByID(id int) (*model.Expense, error)
	Save(*model.Expense) error
	Delete(*model.Expense) error
}

type expense struct{ DB *database.DB }

func (s expense) Save(expense *model.Expense) error {
	var category model.Category
	if expense.CategoryID > 0 {
		if err := s.DB.Where("id = ?", expense.CategoryID).First(&category).Error; err != nil {
			return errors.Wrap(err, fmt.Sprintf("expense.Save: category with id = %d not exists", expense.CategoryID))
		}
	}
	var account model.Account
	if expense.AccountID > 0 {
		if err := s.DB.Where("id = ?", expense.AccountID).First(&account).Error; err != nil {
			return errors.Wrap(err, fmt.Sprintf("expense.Save: account with id = %d not exists", expense.AccountID))
		}
	}

	if err := s.DB.Save(expense).Error; err != nil {
		return errors.Wrap(err, "expense.Save: could not save")
	}
	expense.Category = category
	expense.Account = account
	return nil
}

func (s expense) Delete(expense *model.Expense) error {
	return s.DB.Delete(expense).Error
}

func (s expense) All() ([]*model.Expense, error) {
	var objects []*model.Expense
	if err := s.DB.
		Preload("Category").
		Preload("Account").
		Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func (s *expense) ByID(id int) (*model.Expense, error) {
	var object model.Expense
	if err := s.DB.
		Where("id = ?", id).
		First(&object).Error; err != nil {
		return nil, err
	}
	return &object, nil
}
