package user

import (
	"oauth/common/database"
	"oauth/common/database/models"

	"gorm.io/gorm"
)

type Persistor interface {
	Create(user *models.User) error
}

type persistor struct {
	Conn *gorm.DB
}

func GetPersistor() *persistor {
	return &persistor{Conn: database.DB}
}

func (u *persistor) Create(user *models.User) error {
	return u.Conn.Create(user).Error
}
