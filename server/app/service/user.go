package service

import (
	"github.com/thanhchungbtc/mywallet/server/app/model"
	"github.com/thanhchungbtc/mywallet/server/database"
)

type User interface {
	All() ([]*model.User, error)
	ByID(id int) (*model.User, error)
	ByUsername(username string) (*model.User, error)
	Save(*model.User) error
	Delete(*model.User) error
}

type user struct{ DB *database.DB }

func (s user) Save(user *model.User) error {
	return s.DB.Save(user).Error
}

func (s user) Delete(user *model.User) error {
	return s.DB.Delete(user).Error
}

func (s user) All() ([]*model.User, error) {
	var objects []*model.User
	if err := s.DB.Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func (s *user) ByID(id int) (*model.User, error) {
	var object model.User
	if err := s.DB.
		Where("id = ?", id).
		Find(&object).Error; err != nil {
		return nil, err
	}
	return &object, nil
}

func (s *user) ByUsername(username string) (*model.User, error) {
	var object model.User
	if err := s.DB.
		Where("username = ?", username).
		First(&object).Error; err != nil {
		return nil, err
	}
	return &object, nil
}
