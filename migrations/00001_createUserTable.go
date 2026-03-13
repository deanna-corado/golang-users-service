package migrations

import (
	"user-service/models"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateUserTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "00001_createUserTable",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
