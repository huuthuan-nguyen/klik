package transformer

import (
	"github.com/huuthuan-nguyen/klik-dokter/app/model"
	"time"
)

type ProductTransformer struct {
	ID        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Quantity  uint64    `json:"quantity"`
	Price     float64   `json:"price"`
	Unit      string    `json:"unit"`
	Status    int       `json:"status"`
	CreatedBy int       `json:"created_by"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Transform /**
func (product *ProductTransformer) Transform(e any) any {
	productModel, ok := e.(model.Product)
	if !ok {
		return e
	}

	product.ID = productModel.ID
	product.SKU = productModel.SKU
	product.Name = productModel.Name
	product.Quantity = productModel.Quantity
	product.Price = productModel.Price
	product.Unit = productModel.Unit
	product.Status = productModel.Status
	product.CreatedBy = productModel.CreatedBy
	product.CreateAt = productModel.CreatedAt
	product.UpdatedAt = productModel.UpdatedAt
	return *product
}
