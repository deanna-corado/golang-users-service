package services

import (
	"errors"
	"net/mail"
	"strings"
	"user-service/models"
	"user-service/repositories"
	"user-service/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{repo: r}
}


// FOR USERS
func (s *UserService) PatchMe(id uint, updates map[string]any) (models.User, error) {
	existingUser, err := s.repo.FindMeByID(id)
	if err != nil || existingUser == nil {
		return models.User{}, utils.ErrUserNotFound
	}

	// YUNG BINIGAY LANG UPDAE
	if email, ok := updates["email"].(string); ok {
		trimmedEmail := strings.TrimSpace(email)
		if trimmedEmail == "" {
			return models.User{}, utils.ErrFillAllFields
		}
		// Validate email format
		_, err := mail.ParseAddress(trimmedEmail)
		if err != nil {
			return models.User{}, utils.ErrInvalidEmailFormat
		}
		existingUser.Email = trimmedEmail
	}

	if password, ok := updates["password"].(string); ok {
		trimmedPassword := strings.TrimSpace(password)
		if trimmedPassword == "" {
			return models.User{}, utils.ErrFillAllFields
		}
		// Validate password length
		if len(trimmedPassword) < 6 {
			return models.User{}, errors.New("password must be at least 6 characters long")
		}
		// hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(trimmedPassword), bcrypt.DefaultCost)
		if err != nil {
			return models.User{}, errors.New("failed to hash password")
		}
		existingUser.Password = string(hash)
	}

	if firstName, ok := updates["firstname"].(string); ok {
		trimmedFirstName := strings.TrimSpace(firstName)
		if trimmedFirstName == "" {
			return models.User{}, utils.ErrFillAllFields
		}
		existingUser.FirstName = trimmedFirstName
	}
	if lastName, ok := updates["lastname"].(string); ok {
		trimmedLastName := strings.TrimSpace(lastName)
		if trimmedLastName == "" {
			return models.User{}, utils.ErrFillAllFields
		}
		existingUser.LastName = trimmedLastName
	}

	if err := s.repo.UpdateMe(existingUser); err != nil {
		return models.User{}, err
	}

	return *existingUser, nil
}
