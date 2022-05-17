package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"net/url"
	"time"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:product"`
	ID            int       `bun:"id,pk,autoincrement"`
	SKU           string    `bun:"sku"`
	Name          string    `bun:"name"`
	Quantity      uint64    `bun:"quantity"`
	Price         float64   `bun:"price"`
	Unit          string    `bun:"unit"`
	Status        int       `bun:"status"`
	CreatedBy     int       `bun:"created_by,pk"`
	Owner         *User     `bun:"rel:belongs-to,join:created_by=id"`
	CreatedAt     time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,default:current_timestamp"`
}

// FindProducts /**
func FindProducts(ctx context.Context, db bun.IDB, params url.Values) ([]Product, error) {
	var products = make([]Product, 0)
	query := db.NewSelect().Model(&products)

	for k := range params {
		v := params.Get(k)

		if k == "sku" {
			query.Where("sku ILIKE ?", fmt.Sprintf("%%%s%%", v))
			continue
		}
	}

	err := query.Scan(ctx)

	return products, err
}

// FindOneProductByID /**
func FindOneProductByID(ctx context.Context, id int, db bun.IDB) (Product, error) {
	product := Product{
		ID: id,
	}

	err := db.NewSelect().Model(&product).WherePK().Scan(ctx)
	return product, err
}

// Create /**
func (product *Product) Create(ctx context.Context, db bun.IDB) error {
	user, ok := ctx.Value("user").(*User)
	if !ok {
		return errors.New("invalid user")
	}
	product.CreatedAt = time.Now().UTC()
	product.UpdatedAt = time.Now().UTC()
	product.CreatedBy = user.ID

	if _, err := db.NewInsert().Model(product).Exec(ctx); err != nil {
		return err
	}
	return nil

}

// Update /**
func (product *Product) Update(ctx context.Context, db bun.IDB) error {
	// force current time on updated_at column
	product.UpdatedAt = time.Now().UTC()
	if _, err := db.NewUpdate().
		Model(product).
		ExcludeColumn("created_at", "id", "created_by").
		Where("id=?", product.ID).
		Returning("*").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// Delete /**
func (product *Product) Delete(ctx context.Context, db bun.IDB) error {

	if _, err := db.NewDelete().
		Model(product).
		Where("id=?", product.ID).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}
