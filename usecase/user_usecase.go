package usecase

import (
	"errors"
	"simple-payment/model"
	"simple-payment/repository"
	"simple-payment/util"
	"simple-payment/util/authenticator"
)

type UserUseCase interface {
	Insert(user *model.User) error
	Login(user model.UserCredential) (token string, err error)
}

type userUseCase struct {
	repo         repository.UserRepository
	tokenService authenticator.AccessToken
}

func (uu *userUseCase) Insert(user *model.User) error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return uu.repo.Insert(user)
}

func (uu *userUseCase) Login(userCredential model.UserCredential) (token string, err error) {
	user, err := uu.repo.UserLogin(&userCredential)

	if err != nil && err.Error() == "Invalid Password" {
		return "", errors.New("Invalid Password")
	}

	if user.UserId == "" {
		return "", errors.New("User does not exist, please sign up first")
	}

	userCredential.UserId = user.UserId
	if err == nil {
		token, err := uu.tokenService.CreateAccessToken(&userCredential)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", err
	}
}

func NewUserUseCase(repo repository.UserRepository, tokenService authenticator.AccessToken) UserUseCase {
	return &userUseCase{
		repo:         repo,
		tokenService: tokenService,
	}
}
