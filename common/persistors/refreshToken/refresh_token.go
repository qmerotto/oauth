package refresh_token

import (
	"github.com/google/uuid"
	"oauth/common/database"
	"oauth/common/database/models"

	"gorm.io/gorm"
)

type Persistor interface {
	Create(refreshToken *models.RefreshToken) error
	GetByUUID(uuid uuid.UUID) (*models.RefreshToken, error)
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

func (p *persistor) GetByUUID(uuid uuid.UUID) (*models.RefreshToken, error) {
	refreshToken := &models.RefreshToken{UUID: uuid}
	return refreshToken, p.Conn.Find(&models.RefreshToken{UUID: uuid}).Error
}
