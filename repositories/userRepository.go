package repositories

import (
	"errors"
	"user-service/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// REGISTER USER
func (r *UserRepository) Register(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FOR USERS
func (r *UserRepository) FindMeByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) UpdateMe(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	return user, result.Error
}

// // GET ALL USERS - WALA PA
// func (r *UserRepository) GetAllUsers() ([]models.User, error) {
// 	var users []models.User
// 	result := r.db.Find(&users)
// 	return users, result.Error
// }

// // ADMIN SANA KASO NOT SURE PA YUN
// func (r *UserRepository) FindByID(id uint) (models.User, error) {
// 	var user models.User
// 	result := r.db.First(&user, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return models.User{}, errors.New("user not found")
// 	}
// 	return user, result.Error
// }

// // UPDATE SA ADMIN? NEED BA YUN
// func (r *UserRepository) Update(user *models.User) error {
// 	result := r.db.Save(user)
// 	return result.Error
// }

// // DELETE FOR ADMIN
// func (r *UserRepository) Delete(user *models.User) error {
// 	result := r.db.Delete(user)
// 	return result.Error
// }
