package dto

import "go_ecommerce/internal/model"

type ProductCardResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"discountPrice"`
	Brand         struct {
		Name string `json:"name"`
	} `json:"brand"`
	PrimaryImage struct {
		URL string `json:"url"`
	} `json:"primaryImage"`
	Status model.ProductStatus `json:"status"`
}
