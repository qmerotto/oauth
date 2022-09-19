package models

type Order struct {
	ID       int64     `gorm:"primaryKey;type:integer" json:"id"`
	VAT      float32   `gorm:"type:numeric;not null" json:"vat"`
	Total    float32   `gorm:"type:numeric;not null" json:"total"`
	Products []Product `gorm:"many2many:products_orders"`
}
