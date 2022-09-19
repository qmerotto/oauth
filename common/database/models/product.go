package models

type Product struct {
	ID    string  `gorm:"primaryKey;type:integer" json:"id"`
	Name  string  `gorm:"type:varchar(64)" json:"name"`
	Price float32 `gorm:"type:numeric" json:"price"`
}
