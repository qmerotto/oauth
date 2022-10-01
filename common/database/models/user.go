package models

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	Name     string    `gorm:"type:varchar(32)" json:"name"`
	Password string    `gorm:"type:varchar(32)" json:"password"`
	Email    string    `gorm:"type:varchar(32)" json:"email"`
}
