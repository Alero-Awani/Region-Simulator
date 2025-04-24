package service

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/domain"
	"Region-Simulator/internal/dto"
	"Region-Simulator/internal/helper"
	"Region-Simulator/internal/repository"
	"Region-Simulator/pkg/notification"
	"errors"
	"fmt"
	"log"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {

	user, err := s.Repo.FindUser(email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s UserService) Signup(input dto.UserSignUp) (string, error) {

	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}
	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}
	// generate token

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

// isVerifiedUser Check if the user is verified by checking their id in the database

func (s UserService) isVerifiedUser(id uint) bool {
	currentUser, err := s.Repo.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) error {
	// if user already verified
	if s.isVerifiedUser(e.ID) {
		return errors.New("user already verified")
	}

	// generate the verification code
	code, err := s.Auth.GenerateCode()

	if err != nil {
		return err
	}

	// Update user
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)

	if err != nil {
		return errors.New("unable to update the verification code")
	}

	user, _ = s.Repo.FindUserById(e.ID)

	// send SMS
	notificationClient := notification.NewNotificationClient(s.Config)

	msg := fmt.Sprintf("Your verification code is: %v", code)

	err = notificationClient.SendSMS(user.Phone, msg)

	if err != nil {
		return errors.New("error sending SMS")
	}

	// Return verification code

	return nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	if s.isVerifiedUser(id) {
		log.Println("verified...")
		return errors.New("user already verified")
	}
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}
	if user.Code != code {
		return errors.New("verification code incorrect")
	}
	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}
	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)

	if err != nil {
		return errors.New("unable to verify the user")
	}

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {

	// Find the existing user
	user, _ := s.Repo.FindUserById(id)

	if user.UserType == domain.SELLER {
		return "", errors.New("you have already joined the seller program")
	}

	// Update the user
	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})
	if err != nil {
		return "", err
	}
	token, err := s.Auth.GenerateToken(id, user.Email, seller.UserType)

	// create bank information
	account := domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
		UserId:      id,
	}

	err = s.Repo.CreateBankAccount(account)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId int) ([]interface{}, error) {
	return nil, nil
}
