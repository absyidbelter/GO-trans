package usecase

import "GO-Payment/internal/repository"

type authUsecase struct {
	customerRepo repository.UserRepository
}

type AuthUsecase interface {
	LoginUser(username string, password string) (int64, error)
	LogoutUser(userID int64) error
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{customerRepo: userRepo}
}

func (uc *authUsecase) LoginUser(username string, password string) (int64, error) {
	users, err := uc.customerRepo.FindByName(username)
	if err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, ErrUsecaseInvalidAuth
	}

	user := users[0]
	if user.Password != password {
		return 0, ErrUsecaseInvalidAuth
	}

	return int64(user.ID), nil
}

func (uc *authUsecase) LogoutUser(userID int64) error {
	err := uc.customerRepo.Delete(int(userID))
	if err != nil {
		return err
	}
	return nil
}
