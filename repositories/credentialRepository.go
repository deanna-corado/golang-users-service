package repositories

import (
	"user-service/models"

	"gorm.io/gorm"
)

type CredentialRepository struct {
	db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) *CredentialRepository {
	return &CredentialRepository{db: db}
}

func (r *CredentialRepository) FindByClientID(clientID string) (*models.Credential, error) {
	var cred models.Credential
	result := r.db.First(&cred, "client_id = ?", clientID)
	return &cred, result.Error
}
