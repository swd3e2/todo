package application

import (
	"errors"
	"github.com/swd3e2/todo/internal/application/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

//go:generate mockgen -source=user.go -destination=generated_mocks/user_mock.go

// User Пользователь
type User struct {
	Id          int64
	Name        string
	LastName    string
	Age         int
	LastUpdated time.Time
	Login       string
	Password    string
}

// UserRepository Репозиторий пользователей
type UserRepository interface {
	// FindByLogin Поиск пользователя по логину
	FindByLogin(login string) (*User, error)

	// Save Сохранение пользователя
	Save(user *User) error
}

var (
	UserNotFound        = errors.New("пользователь не найден")
	WrongPassword       = errors.New("неправильный пароль")
	LoginIsAlreadyInUse = errors.New("логин уже занят")
)

type UserService interface {
	// Authorize Получение jwt токена
	Authorize(login string, password string) (token *Token, err error)

	// Register Регистрация нового пользователя
	Register(Name string, LastName string, Login string, Password string) error
}

// UserServiceImpl Сервис пользователей
type UserServiceImpl struct {
	userRepository  UserRepository
	jwtTokenBuilder jwt.Builder
}

func NewUserService(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

// Token Jwt токен
type Token struct {
	Token string
}

// Authorize Авторизация пользователя, выдача токена
func (s *UserServiceImpl) Authorize(login string, password string) (token *Token, err error) {
	var user *User
	if user, err = s.userRepository.FindByLogin(login); err != nil || user == nil {
		return nil, UserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, WrongPassword
	}

	t, err := s.jwtTokenBuilder.CreateToken(strconv.Itoa(int(user.Id)))
	if err != nil {
		return nil, err
	}

	token = &Token{t}

	return token, nil
}

// Register Регистрация нового пользователя
func (s *UserServiceImpl) Register(Name string, LastName string, Login string, Password string) error {
	var err error

	existingUser, err := s.userRepository.FindByLogin(Login)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return LoginIsAlreadyInUse
	}

	var bytes []byte

	if bytes, err = bcrypt.GenerateFromPassword([]byte(Password), 14); err != nil {
		return err
	}

	password := string(bytes)

	user := &User{
		Name:     Name,
		LastName: LastName,
		Login:    Login,
		Password: password,
	}

	return s.userRepository.Save(user)
}
