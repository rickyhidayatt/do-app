package usecase

import (
	"errors"

	"github.com/rickyhidayatt/do-app/model"
	"github.com/rickyhidayatt/do-app/repository"
	"github.com/rickyhidayatt/do-app/utils"
)

type UserUseCase interface {
	ViewAllUser(page int, totalRows int) ([]model.User, error)
	CreateNewUser(newUser *model.User) error
	UpdateUser(user model.User) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) ViewAllUser(page int, totalRows int) ([]model.User, error) {
	return u.userRepo.ViewAll(page, totalRows)
}

func (u *userUseCase) CreateNewUser(newUser *model.User) error {
	newUser.Id = utils.GenerateId()
	return u.userRepo.CreateNew(newUser)
}

func (u *userUseCase) UpdateUser(user model.User) error {
	if len(user.Name) < 3 || len(user.Name) > 20 {
		return errors.New("Nama Minimal 3 Sampai 20 karakter")
	}
	return nil
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepository,
	}
}