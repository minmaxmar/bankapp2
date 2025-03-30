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
	UserID     int64           `json:"userid" gorm:"not null"`
	User       User            `json:"user" gorm:"foreignKey:UserID"`
}

type User struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement;type:bigserial"`
	FirstName string `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(50);not null"`
	Email     string `json:"email" gorm:"type:varchar(50);unique;not null"`
	Banks     []Bank `gorm:"many2many:bank_users;"`
	Cards     []Card `gorm:"foreignKey:UserID"`
}

type Bank struct {
	ID    int64  `json:"id" gorm:"primaryKey;autoIncrement;type:bigserial"`
	Name  string `json:"name" gorm:"type:varchar(50);not null"`
	Users []User `gorm:"many2many:bank_usess;"`
	Cards []Card `gorm:"foreignKey:BankID"`
}
