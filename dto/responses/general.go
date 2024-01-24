package responses

type ResponsesHTTP struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type JWTTokenFormat struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
