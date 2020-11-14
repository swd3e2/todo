package service

import (
	"errors"
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

type Token struct {
	Token string
}

func (this *UserService) Authorize(login string, password string) (token *Token, err error) {
	var user *domain.User

	if user, err = this.userRepository.FindByLogin(login); err != nil || user == nil {
		return nil, errors.New("Пользователь не найден")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Credentials.Password), []byte(password)); err != nil {
		return nil, errors.New("Неправильный пароль")
	}

	token = &Token{
		Token: "sadasdasda",
	}

	return token, nil
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

	password := string(bytes)

	user := &domain.User{
		Name:     payload.Name,
		LastName: payload.LastName,
		Credentials: domain.Credentials{
			Login:    payload.Login,
			Password: password,
		},
	}

	if _, err := this.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}
