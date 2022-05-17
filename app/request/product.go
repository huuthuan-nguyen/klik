package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/huuthuan-nguyen/klik-dokter/app/model"
	"net/http"
)

type Product struct {
	SKU      string  `json:"sku" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Quantity uint64  `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Unit     string  `json:"unit" validate:"required"`
	Status   int     `json:"status" validate:"required"`
}

// Validate /**
func (product *Product) Validate(r *http.Request) error {
	if validate, ok := r.Context().Value("validate").(*validator.Validate); ok {
		return validate.Struct(product)
	}

	return nil
}

// Bind /**
func (product *Product) Bind(r *http.Request, productModel *model.Product) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		return err
	}

	if err := product.Validate(r); err != nil {
		return err
	}

	productModel.SKU = product.SKU
	productModel.Name = product.Name
	productModel.Quantity = product.Quantity
	productModel.Price = product.Price
	productModel.Unit = product.Unit
	productModel.Status = product.Status
	return nil
}
