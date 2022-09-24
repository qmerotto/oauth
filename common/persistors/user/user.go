package user

import (
	"oauth/common/database"
	"oauth/common/database/models"

	"gorm.io/gorm"
)

type Persistor interface {
	Create(order *models.User) error
}

type user struct {
	Conn *gorm.DB
}

func GetPersistor() *user {
	return &user{Conn: database.DB}
}

func (u *user) Create(order *models.User) error {
	return u.Conn.Create(order).Error
}
