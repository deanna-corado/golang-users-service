package migrations

import (
	"user-service/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateCredentialTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "00002_createCredentialTable",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Credential{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("credentials")
		},
	}

}
