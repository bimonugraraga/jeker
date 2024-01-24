package responses

type BookResponse struct {
	UserId     int    `json:"user_id"`
	Author     string `json:"author"`
	CategoryId int    `json:"category_id"`
	Category   string `json:"category"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	Synopsis   string `json:"synopsis"`
	Status     string `json:"status"`
}
