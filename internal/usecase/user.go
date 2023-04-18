package usecase

import (
	model "GO-Payment/internal/model/entity"
	"GO-Payment/internal/repository"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUser(userID int) (*model.User, error)
	CreateUser(name, email, password string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
}

type userUsecase struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
}

type USConfig struct {
	UserRepository   repository.UserRepository
	WalletRepository repository.WalletRepository
}

func NewUserUsecase(c *USConfig) UserUsecase {
	return &userUsecase{
		userRepository:   c.UserRepository,
		walletRepository: c.WalletRepository,
	}
}

func (uu *userUsecase) GetUser(userID int) (*model.User, error) {
	user, err := uu.userRepository.FindById(uint(userID))
	if err != nil {
		return user, ErrUsecaseInternal
	}
	return user, nil
}

func (uu *userUsecase) CreateUser(name, email, password string) (*model.User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return &model.User{}, ErrUsecaseInvalidAuth
	}

	user, err := uu.userRepository.FindByEmail(email)
	if err != nil {
		return user, ErrUsecaseInternal
	}
	if user.ID != 0 {
		return user, ErrUsecaseExistsUsername
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return user, ErrUsecaseInternal
	}

	newUser := &model.User{
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
	}

	newUser, err = uu.userRepository.Save(newUser)
	if err != nil {
		return newUser, ErrUsecaseInternal
	}

	return newUser, nil
}

func (uu *userUsecase) GetAllUsers() ([]model.User, error) {
	res, err := uu.userRepository.FindAll()
	if err == repository.ErrNotFound {
		return nil, ErrUsecaseNoData
	}

	var users []model.User
	for _, u := range res {
		users = append(users, *u)
	}

	return users, nil
}
