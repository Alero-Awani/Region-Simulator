package service

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/domain"
	"Region-Simulator/internal/dto"
	"Region-Simulator/internal/helper"
	"Region-Simulator/internal/repository"
	"github.com/pkg/errors"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {
	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})
	return err
}

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {

	return nil, nil
}

func (s CatalogService) DeleteCategory(input any) error {
	return nil
}

func (s CatalogService) GetCategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {
	category, err := s.Repo.FindCategoryByID(id)
	if err != nil {
		return nil, errors.New("category does not exist")
	}
	return category, nil
}
