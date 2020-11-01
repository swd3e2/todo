package service

import (
	"github.com/swd3e2/todo/internal/domain"
	"github.com/swd3e2/todo/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (this *UserService) Authorize(login string, password string) {

}

type RegisterPayload struct {
	Name     string
	LastName string
	Login    string
	Password string
}

func (this *UserService) Register(payload RegisterPayload) error {
	var bytes []byte
	var err error

	if bytes, err = bcrypt.GenerateFromPassword([]byte(payload.Password), 14); err != nil {
		return err
	}

	user := &domain.User{
		Name:     payload.Name,
		LastName: payload.LastName,
		Credentials: domain.Credentials{
			Login:    payload.Login,
			Password: string(bytes),
		},
	}

	if _, err := this.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}
