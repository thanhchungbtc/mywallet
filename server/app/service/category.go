package service

import (
	"github.com/thanhchungbtc/mywallet/server/app/database"
	"github.com/thanhchungbtc/mywallet/server/app/model"
)

type CategoryService struct {
	db *database.DB
}

func NewCategoryService(db *database.DB) *CategoryService {
	return &CategoryService{db}
}

func (s *CategoryService) FindAll(query interface{}, args ...interface{}) ([]*model.Category, error) {
	var objects []*model.Category
	err := s.db.
		Preload("Expenses").
		Where(query, args).
		Find(&objects).Error
	return objects, err
}

func (s *CategoryService) FindOne(query interface{}, args ...interface{}) (*model.Category, error) {
	var instance model.Category
	err := s.db.
		Preload("Expenses").
		Where(query, args).
		First(&instance).Error
	return &instance, err
}

func (s *CategoryService) Save(data *model.Category) error {
	return s.db.Save(data).Error
}

func (s *CategoryService) Create(data *model.Category) error {
	return s.db.Create(data).Error
}

func (s *CategoryService) Delete(query interface{}, args ...interface{}) error {
	// permanently delete
	return s.db.Where(query, args).Unscoped().Delete(&model.Category{}).Error
}
