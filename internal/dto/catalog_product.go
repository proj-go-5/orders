package dto

type ListProductsResponse struct {
	Products []*Product `json:"page"`
}

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}
