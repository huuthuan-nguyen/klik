package handler

import (
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/klik-dokter/app/model"
	"github.com/huuthuan-nguyen/klik-dokter/app/render"
	"github.com/huuthuan-nguyen/klik-dokter/app/request"
	"github.com/huuthuan-nguyen/klik-dokter/app/transformer"
	"github.com/huuthuan-nguyen/klik-dokter/app/utils"
	"log"
	"net/http"
	"strconv"
)

// ProductStore /**
func (handler *Handler) ProductStore(w http.ResponseWriter, r *http.Request) {
	productRequest := request.Product{}
	productModel := model.Product{}

	if err := productRequest.Bind(r, &productModel); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := productModel.Create(r.Context(), handler.db); err != nil {
		log.Println(err)
		utils.PanicInternalServerError()
		return
	}

	productTransformer := &transformer.ProductTransformer{}
	productItem := transformer.NewItem(productModel, productTransformer)

	// render product item to JSON
	render.JSON(w, r, productItem)
}

// ProductUpdate /**
func (handler *Handler) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	productRequest := request.Product{}
	productModel := model.Product{}
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		utils.PanicNotFound()
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		utils.PanicInternalServerError()
		return
	}

	productModel.ID = productID
	if err := productRequest.Bind(r, &productModel); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := productModel.Update(r.Context(), handler.db); err != nil {
		utils.PanicInternalServerError()
		return
	}

	productTransformer := &transformer.ProductTransformer{}
	productItem := transformer.NewItem(productModel, productTransformer)

	// render product item to JSON
	render.JSON(w, r, productItem)
}

// ProductDestroy /**
func (handler *Handler) ProductDestroy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		utils.PanicNotFound()
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		utils.PanicInternalServerError()
		return
	}

	productModel := model.Product{}
	productModel.ID = productID

	if err := productModel.Delete(r.Context(), handler.db); err != nil {
		utils.PanicInternalServerError()
		return
	}

	render.NoContent(w, r)
}

// ProductIndex /**
func (handler *Handler) ProductIndex(w http.ResponseWriter, r *http.Request) {
	// prepare data response
	products, err := model.FindProducts(r.Context(), handler.db, r.URL.Query())
	if err != nil {
		render.Error(w, r, err)
		return
	}

	productTransformer := &transformer.ProductTransformer{}
	productCollection := transformer.NewCollection(products, productTransformer)

	// render Collection to JSON
	render.JSON(w, r, productCollection)
}

// ProductShow /**
func (handler *Handler) ProductShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.PanicNotFound()
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		utils.PanicInternalServerError()
		return
	}

	product, err := model.FindOneProductByID(r.Context(), productID, handler.db)
	if err != nil {
		render.Error(w, r, err)
		return
	}

	productTransformer := &transformer.ProductTransformer{}
	productItem := transformer.NewItem(product, productTransformer)

	// render product item to JSON
	render.JSON(w, r, productItem)
}
