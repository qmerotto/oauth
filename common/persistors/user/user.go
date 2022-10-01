package user

import (
	"oauth/common/database"
	"oauth/common/database/models"

	"gorm.io/gorm"
)

type Persistor interface {
	Create(user *models.User) error
	GetUserByMail(email string) (*models.User, error)
}

type persistor struct {
	Conn *gorm.DB
}

func GetPersistor() *persistor {
	return &persistor{Conn: database.DB}
}

func (p *persistor) Create(user *models.User) error {
	return p.Conn.Create(user).Error
}

func (p *persistor) GetUserByMail(email string) (*models.User, error) {
	var user models.User
	tx := p.Conn.Where("email = ?", email).Find(&user)

	return &user, tx.Error
}
