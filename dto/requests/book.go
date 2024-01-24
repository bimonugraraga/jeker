package requests

type BookRequest struct {
	CategoryId int    `json:"category_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Synopsis   string `json:"synopsis,omitempty"`
	Cover      string `json:"cover,omitempty"`
}
