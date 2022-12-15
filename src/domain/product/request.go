package product

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Sku         string `json:"sku" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"image_url" validate:"required"`
	Price       uint8  `json:"price" validate:"required"`
	Stock       uint8  `json:"stock" validate:"required"`
}
