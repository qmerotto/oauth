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
				name varchar(32) NOT NULL,
				password varchar(100) NOT NULL,
				email varchar(32) UNIQUE NOT NULL
			);
			CREATE TABLE clients (
				id uuid PRIMARY KEY,
				name varchar(32) NOT NULL
			);
			CREATE TABLE resource_servers (
				id uuid PRIMARY KEY,
				name varchar(32) NOT NULL
			);
			CREATE TABLE resources (
				id uuid PRIMARY KEY,
				name varchar(32) NOT NULL
			);
			CREATE TABLE roles (
				id uuid PRIMARY KEY,
				name varchar(32) NOT NULL
			);
			CREATE TABLE refresh_tokens (
				uuid uuid PRIMARY KEY,
				token varchar NOT NULL,
				user_uuid uuid NOT NULL,
				FOREIGN KEY (user_uuid) REFERENCES users(uuid)
			);

			INSERT INTO clients VALUES(uuid_generate_v4 (), 'web');
			INSERT INTO clients VALUES(uuid_generate_v4 (), 'mobile');

			INSERT INTO resource_servers VALUES(uuid_generate_v4 (), 'ms_1');
			INSERT INTO resource_servers VALUES(uuid_generate_v4 (), 'ms_2');
			INSERT INTO resource_servers VALUES(uuid_generate_v4 (), 'front');

			INSERT INTO resources VALUES(uuid_generate_v4 (), 'account');
			INSERT INTO resources VALUES(uuid_generate_v4 (), 'obj1');
			INSERT INTO resources VALUES(uuid_generate_v4 (), 'obj2');

			INSERT INTO roles VALUES(uuid_generate_v4 (), 'basic_user');
			INSERT INTO roles VALUES(uuid_generate_v4 (), 'admin');
		`).Error; err != nil {
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Exec(`
			DROP TABLE users;
			DROP TABLE clients;
			DROP TABLE resource_servers;
			DROP TABLE resources;
			DROP TABLE roles;
			DROP TABLE refresh_tokens;
		`).Error; err != nil {
			return err
		}
		return nil
	},
}
