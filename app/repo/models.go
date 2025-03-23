package repoModels

import (
	"github.com/go-openapi/strfmt"
)

type Card struct {
	ID         int64           `json:"id" gorm:"primaryKey;autoIncrement;type:bigserial"`
	CardNumber int64           `json:"card_number" gorm:"type:numeric;not null"`
	CreateDate strfmt.DateTime `json:"create_date" gorm:"type:timestamp without time zone;not null"`
	ExpireDate strfmt.Date     `json:"expire_date" gorm:"type:timestamp without time zone;not null"`
	Total      int64           `json:"total" gorm:"type:numeric;not null;default:0"`
	BankID     int64           `json:"bankid" gorm:"not null"`
	Bank       Bank            `json:"bank" gorm:"foreignKey:BankID"`
	ClientID   int64           `json:"clientid" gorm:"not null"`
	Client     Client          `json:"client" gorm:"foreignKey:ClientID"`
}

type Client struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement;type:bigserial"`
	FirstName string `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(50);not null"`
	Email     string `json:"email" gorm:"type:varchar(50);unique;not null"`
	Banks     []Bank `gorm:"many2many:bank_clients;"`
	Cards     []Card `gorm:"foreignKey:ClientID"`
}

type Bank struct {
	ID      int64    `json:"id" gorm:"primaryKey;autoIncrement;type:bigserial"`
	Name    string   `json:"name" gorm:"type:varchar(50);not null"`
	Clients []Client `gorm:"many2many:bank_clients;"`
	Cards   []Card   `gorm:"foreignKey:BankID"`
}
