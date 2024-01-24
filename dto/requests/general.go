package requests

type HeaderAccessToken struct {
	AccessToken string
}

type GeneralParams struct {
	UserId int `json:"user_id"`
}

type GeneralPaginationQuery struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	OrderWith string `json:"order_with"`
	OrderBy   string `json:"order_by"`
}
