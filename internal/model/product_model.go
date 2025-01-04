package model

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Description string  `json:"description" validate:"required,max=255"`
	Price       float64 `json:"price" validate:"required"`
}
