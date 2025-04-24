package service

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/helper"
	"Region-Simulator/internal/repository"
)

type CatalogService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}
