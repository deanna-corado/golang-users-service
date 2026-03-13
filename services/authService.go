package services

import (
	"errors"
	"net/mail"
	"os"
	"strings"
	"time"
	"user-service/models"
	"user-service/repositories"
	"user-service/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(r *repositories.UserRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Register(email, password, first, last string) error {
	// Validate required fields
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" || strings.TrimSpace(first) == "" || strings.TrimSpace(last) == "" {
		return utils.ErrFillAllFields
	}

	// Validate email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		return utils.ErrInvalidEmailFormat
	}

	// Validate password length
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// check if email exists
	_, err = s.repo.FindByEmail(email)
	if err == nil {
		return utils.ErrEmailExists
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		Email:     email,
		Password:  string(hash),
		FirstName: first,
		LastName:  last,
	}

	return s.repo.Register(&user)
}

func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", utils.ErrInvalidCredentials
	}

	// generate jwt token; here yung info
	expDuration, _ := time.ParseDuration(os.Getenv("TOKEN_EXP"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                            // kung kanino
		"exp": time.Now().Add(expDuration).Unix(), //expiration
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", utils.ErrFailedToCreateToken
	}

	return tokenString, nil
}
