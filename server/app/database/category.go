package database

import (
	"github.com/jinzhu/gorm"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type Category interface {
	All() ([]*model.Category, error)
	ByID(id int) (*model.Category, error)
	Save(*model.Category) error
	Delete(*model.Category) error
}

type category struct{ DB *gorm.DB }

func (s category) Save(category *model.Category) error {
	return s.DB.Save(category).Error
}

func (s category) Delete(category *model.Category) error {
	return s.DB.Delete(category).Error
}

func (s category) All() ([]*model.Category, error) {
	var objects []*model.Category
	if err := s.DB.Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func (s *category) ByID(id int) (*model.Category, error) {
	var object model.Category
	if err := s.DB.
		Where("id = ?", id).
		First(&object).Error; err != nil {
		return nil, err
	}
	return &object, nil
}
