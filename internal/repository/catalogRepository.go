package repository

import (
	"Region-Simulator/internal/domain"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type CatalogRepository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryByID(id int) (*domain.Category, error)
	EditCategory(e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
}

type catalogRepository struct {
	db *gorm.DB
}

func (c catalogRepository) CreateCategory(e *domain.Category) error {
	err := c.db.Create(&e).Error
	fmt.Println("This is the domain input", e)
	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("Create category failed")
	}
	return nil
}

func (c catalogRepository) FindCategories() ([]*domain.Category, error) {
	var categories []*domain.Category
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c catalogRepository) FindCategoryByID(id int) (*domain.Category, error) {
	var category *domain.Category
	err := c.db.First(&category, id).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("Category does not exist")
	}
	return category, nil
}

func (c catalogRepository) EditCategory(e *domain.Category) (*domain.Category, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("Failed to update category")
	}
	return e, nil
}

func (c catalogRepository) DeleteCategory(id int) error {
	err := c.db.Delete(&domain.Category{}, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("Failed to delete category")
	}
	return nil
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}
