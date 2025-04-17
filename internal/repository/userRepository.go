package repository

import (
	"Region-Simulator/internal/domain"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id int) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	err := r.db.Create(&usr).Error

	if err != nil {
		log.Printf("Create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}
	return usr, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}
func (r userRepository) FindUserById(id int) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).First(&user).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", u.ID).Updates(u).Error
	if err != nil {
		log.Printf("Update user error %v", err)
		return domain.User{}, errors.New("failed to update user")
	}
	return domain.User{}, nil
}
