package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository Repository
}

func NewService(userRepository Repository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

type Token struct {
	Token string
}

func (s *UserService) Authorize(login string, password string) (token *Token, err error) {
	var user *User

	if user, err = s.userRepository.FindByLogin(login); err != nil || user == nil {
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

func (s *UserService) Register(payload RegisterPayload) error {
	var bytes []byte
	var err error

	if bytes, err = bcrypt.GenerateFromPassword([]byte(payload.Password), 14); err != nil {
		return err
	}

	password := string(bytes)

	user := &User{
		Name:     payload.Name,
		LastName: payload.LastName,
		Credentials: Credentials{
			Login:    payload.Login,
			Password: password,
		},
	}

	if _, err := s.userRepository.Save(user); err != nil {
		return err
	}

	return nil
}
