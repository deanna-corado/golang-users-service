package services

import (
	"errors"
	"user-service/repositories"
)

type CredentialService struct {
	repo *repositories.CredentialRepository
}

func NewCredentialService(repo *repositories.CredentialRepository) *CredentialService {
	return &CredentialService{repo: repo}
}

// find cred and compare
func (s *CredentialService) Validate(clientID, secret string) error {
	cred, err := s.repo.FindByClientID(clientID)
	if err != nil {
		return errors.New("invalid client")
	}

	if cred.SecretKey != secret {
		return errors.New("invalid secret")
	}

	return nil
}
