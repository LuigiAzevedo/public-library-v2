package usecase

import (
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/LuigiAzevedo/public-library-v2/internal/domain/entity"
	"github.com/LuigiAzevedo/public-library-v2/internal/errs"
	r "github.com/LuigiAzevedo/public-library-v2/internal/ports/repository"
	u "github.com/LuigiAzevedo/public-library-v2/internal/ports/usecase"
)

type userService struct {
	userRepo r.UserRepository
}

// NewUserService creates a new instance of userService
func NewUserService(repository r.UserRepository) u.UserUsecase {
	return &userService{
		userRepo: repository,
	}
}

func (s *userService) GetUser(id int) (*entity.User, error) {
	user, err := s.userRepo.Get(id)
	if err != nil {
		return nil, errors.Wrap(err, errs.ErrGetUser)
	}

	return user, nil
}

func (s *userService) CreateUser(u *entity.User) (int, error) {
	user, err := entity.NewUser(u.Username, u.Password, u.Email)
	if err != nil {
		return 0, errors.Wrap(err, errs.ErrCreateUser)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.Wrap(err, "error has occurred while hashing the password")
	}

	user.Password = string(hashedPassword)

	id, err := s.userRepo.Create(user)
	if err != nil {
		return 0, errors.Wrap(err, errs.ErrCreateUser)
	}

	return id, nil
}

func (s *userService) UpdateUser(u *entity.User) error {
	u.UpdatedAt = time.Now()

	err := u.Validate()
	if err != nil {
		return errors.Wrap(err, errs.ErrUpdateUser)
	}

	err = s.userRepo.Update(u)
	if err != nil {
		return errors.Wrap(err, errs.ErrUpdateUser)
	}

	return nil
}

func (s *userService) DeleteUser(id int) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return errors.Wrap(err, errs.ErrDeleteUser)
	}

	return nil
}