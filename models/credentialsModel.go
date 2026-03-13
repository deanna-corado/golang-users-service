package models

type Credential struct {
	ID        uint
	AppName   string
	ClientID  string `gorm:"unique"`
	SecretKey string
}
