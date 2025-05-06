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
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
	CreateBankAccount(e domain.BankAccount) error

	FindCartItems(uId uint) ([]domain.Cart, error)
	FindCartItem(uId uint, pId uint) (domain.Cart, error)
	CreateCart(c domain.Cart) error
	UpdateCart(c domain.Cart) error
	DeleteCartById(id uint) error
	DeleteCartItems(uId uint) error

	CreateAddress(e domain.Address) error
	UpdateProfile(e domain.Address) error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateAddress(e domain.Address) error {
	err := r.db.Create(&e).Error
	if err != nil {
		log.Printf("Error while creating profile with address %v", err)
		return errors.New("failed to create profile")
	}
	return nil
}

func (r userRepository) UpdateProfile(e domain.Address) error {
	err := r.db.Where("user_id = ?", e.UserId).Updates(e).Error
	if err != nil {
		log.Printf("Error while updating profile with address %v", err)
		return errors.New("failed to update profile")
	}
	return nil
}

func (r userRepository) FindCartItems(uId uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id = ?", uId).Find(&carts).Error
	return carts, err
}

func (r userRepository) FindCartItem(uId uint, pId uint) (domain.Cart, error) {
	cartItem := domain.Cart{}
	err := r.db.Where("user_id = ? AND product_id=?", uId, pId).First(&cartItem).Error
	return cartItem, err
}

func (r userRepository) CreateCart(c domain.Cart) error {
	return r.db.Create(&c).Error
}

func (r userRepository) UpdateCart(c domain.Cart) error {
	var cart domain.Cart
	err := r.db.Model(&cart).Clauses(clause.Returning{}).Where("id=?", c.ID).Updates(c).Error
	return err
}

func (r userRepository) DeleteCartById(id uint) error {
	err := r.db.Delete(&domain.Cart{}, id).Error
	return err
}

func (r userRepository) DeleteCartItems(uId uint) error {
	err := r.db.Where("user_id = ?", uId).Delete(&domain.Cart{}).Error
	return err
}

func (r userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
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
	err := r.db.Preload("Address").First(&user, "email = ?", email).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}
func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.Preload("Address").First(&user, id).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(u).Error
	if err != nil {
		log.Printf("Update user error %v", err)
		return domain.User{}, errors.New("failed to update user")
	}
	return user, nil
}
