package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M20220905185106 = gormigrate.Migration{
	ID: "20220905185106",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
			CREATE TABLE users (
				uuid uuid PRIMARY KEY,
				name varchar(32)
			);
			CREATE TABLE clients (
				id uuid PRIMARY KEY,
				name varchar(32) NOT NULL
			);
			CREATE TABLE ressource_servers (
				id uuid PRIMARY KEY,
				name varchar(32) NOT NULL
			);

			INSERT INTO users VALUES(uuid_generate_v4 (), 'user');
			INSERT INTO users VALUES(uuid_generate_v4 (), 'admin');

			INSERT INTO clients VALUES(uuid_generate_v4 (), 'web');
			INSERT INTO clients VALUES(uuid_generate_v4 (), 'mobile');

			INSERT INTO ressource_servers VALUES(uuid_generate_v4 (), 'ms_1');
			INSERT INTO ressource_servers VALUES(uuid_generate_v4 (), 'ms_2');
			INSERT INTO ressource_servers VALUES(uuid_generate_v4 (), 'front');
		`).Error; err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Exec(`
			DROP TABLE users;
			DROP TABLE clients;
			DROP TABLE ressource_servers;
		`).Error; err != nil {
			return err
		}
		return nil
	},
}
