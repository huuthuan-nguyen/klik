package migration

import (
	"context"
	"github.com/huuthuan-nguyen/klik-dokter/app/model"
	"github.com/uptrace/bun"
)

type InitSchema struct {
	Migration
	Version string
}

func (m *InitSchema) Up() (err error) {
	db := m.GetDB()
	ctx := context.Background()
	return model.WithTransaction(ctx, db, func(tx bun.Tx) error {

		// "users" table
		const users = `CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			is_active BOOLEAN DEFAULT false NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)`
		if _, err = tx.Exec(users); err != nil {
			return err
		}

		// "products" table
		const products = `CREATE TABLE IF NOT EXISTS products(
			id SERIAL PRIMARY KEY,
			sku VARCHAR(255) NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL,
			quantity INT NOT NULL DEFAULT 0,
			price NUMERIC(11,0) NOT NULL DEFAULT 0,
			unit VARCHAR(255) NOT NULL,
			status SMALLINT NOT NULL DEFAULT 0,
			created_by INT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			CONSTRAINT fk_owner
      		FOREIGN KEY(created_by) 
	  			REFERENCES users(id)
	  				ON DELETE RESTRICT
		)`
		if _, err = tx.Exec(products); err != nil {
			return err
		}

		return nil
	})
}

func (m *InitSchema) Down() (err error) {
	db := m.GetDB()
	ctx := context.Background()
	return model.WithTransaction(ctx, db, func(tx bun.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS products, users`)
		return err
	})
}

func (m *InitSchema) GetVersion() string {
	return "20220517000001"
}
