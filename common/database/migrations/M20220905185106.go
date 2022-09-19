package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M20220905185106 = gormigrate.Migration{
	ID: "20220905185106",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.Exec(`
			CREATE TABLE orders (
				id BIGSERIAL PRIMARY KEY, value INT,
				vat numeric NOT NULL DEFAULT 0.0,
				total numeric NOT NULL DEFAULT 0.0
			);
			CREATE TABLE products (
				id varchar(32) PRIMARY KEY,
				name varchar(32) NOT NULL,
				price numeric NOT NULL
			);
			CREATE TABLE products_orders (
				product_id varchar(32) NOT NULL,
				order_id integer NOT NULL
			);
		`).Error; err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Exec(`
			DROP TABLE orders;
			DROP TABLE products;
			DROP TABLE products_orders;
		`).Error; err != nil {
			return err
		}
		return nil
	},
}
