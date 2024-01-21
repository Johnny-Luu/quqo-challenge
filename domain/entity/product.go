package entity

type Product struct {
	BaseModel
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"image_url"`
	Stock       int64   `json:"stock"`
	Price       float64 `json:"price"`
}
