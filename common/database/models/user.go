package models

type User struct {
	UUID int64 `gorm:"primaryKey;type:uuid" json:"uuid"`
}
