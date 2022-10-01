package models

import "github.com/google/uuid"

type RefreshToken struct {
	UUID     uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	Token    string    `gorm:"type:varchar(50)" json:"token"`
	UserUUID uuid.UUID `gorm:"type:uuid" json:"token"`
	//User     User      `gorm:"foreignKey:UserUUID;type:uuid" json:"user""`
}
