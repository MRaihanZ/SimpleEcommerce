package entities

type CartResponse struct {
	Id        int     `json:"id"`
	IdProduct int     `json:"id_product"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Image     string  `json:"image"`
	Rating    float64 `json:"rating"`
	Quantity  int     `json:"quantity"`
}