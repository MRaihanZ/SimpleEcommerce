package entities

// Product represents a product in the database
type ProductResponse struct {
	Code string `json:"code"`
	Data struct {
		ID          int     `json:"id"`
		Title       string  `json:"title"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Image       string  `json:"image"`
		Rating      float64 `json:"rating"`
		TotalRating int     `json:"total_rating"`
	} `json:"data"`
}