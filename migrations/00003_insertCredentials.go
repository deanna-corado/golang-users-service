package migrations

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func InsertCredentialsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "00003_insert_credentials",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
		       INSERT INTO credentials (app_name, client_id, secret_key)
VALUES ('userService', 'user-service-client', 'secretKey');

            `).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec(`
                DELETE FROM credentials WHERE client_id = 'user-service-client'
            `).Error
		},
	}
}
