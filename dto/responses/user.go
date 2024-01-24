package responses

type ResponseRegisterUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
}

type LoginUser struct {
	AccessToken string `json:"access_token"`
}
