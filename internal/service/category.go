package service

import (
	"github.com/thanhchungbtc/mywallet/internal/database"
	"github.com/thanhchungbtc/mywallet/internal/model"
)

type Category interface {
	All() ([]*model.Category, error)
	ByID(id int) (*model.Category, error)
	Save(*model.Category) error
	Delete(*model.Category) error
}

type category struct{ DB *database.DB }

func (c category) Save(category *model.Category) error {
	return c.DB.Save(category).Error
}

func (c category) Delete(category *model.Category) error {
	return c.DB.Delete(category).Error
}

func (c category) All() ([]*model.Category, error) {
	var objects []*model.Category
	if err := c.DB.Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

func (c *category) ByID(id int) (*model.Category, error) {
	var object model.Category
	if err := c.DB.
		Where("id = ?", id).
		Find(&object).Error; err != nil {
		return nil, err
	}
	return &object, nil
}
