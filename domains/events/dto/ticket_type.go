package dto

type TicketType struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Sale     int     `json:"sale"`
	Price    float64 `json:"price"`
}

type CreateTicketType struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
