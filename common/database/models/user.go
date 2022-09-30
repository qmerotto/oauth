package models

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	Username string    `gorm:"type:varchar(32)" json:"username"`
	Password string    `gorm:"type:varchar(32)" json:"password"`
}
