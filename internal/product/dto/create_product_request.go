package dto

import "order-server/internal/product/entity"

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	Description string   `json:"description" validate:"required,min=10"`
	Images      []string `json:"images" validate:"dive,required,url"`
}

func (r CreateProductRequest) ToEntity() entity.Product {
	return entity.Product{
		Name:        r.Name,
		Description: r.Description,
		Images:      r.Images,
	}
}
