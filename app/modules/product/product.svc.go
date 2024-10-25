package product

import (
	productdto "app/app/modules/product/dto"
	productent "app/app/modules/product/ent"

	"github.com/uptrace/bun"
)

type ProductService struct {
	db *bun.DB
}

func newService(db *bun.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

func ProducttEntityReturn() *productent.ProductEntity {
	return &productent.ProductEntity{
		ProductName: "product1",
		Price:       500.00,
		Stock:       100,
	}
}

func (service *ProductService) Product() *productdto.ProductDTOResponse {
	req := ProducttEntityReturn()
	return &productdto.ProductDTOResponse{
		ProductName: req.ProductName,
		Price:       int64(req.Price),
	}
}
