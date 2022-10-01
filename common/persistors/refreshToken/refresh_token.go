package refresh_token

import (
	"oauth/common/database"
	"oauth/common/database/models"

	"gorm.io/gorm"
)

type Persistor interface {
	Create(refreshToken *models.RefreshToken) error
}

type persistor struct {
	Conn *gorm.DB
}

func GetPersistor() *persistor {
	return &persistor{Conn: database.DB}
}

func (p *persistor) Create(refreshToken *models.RefreshToken) error {
	return p.Conn.Create(refreshToken).Error
}
