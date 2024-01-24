package responses

type UserProfileResponses struct {
	UserId          int     `json:"user_id"`
	Username        string  `json:"username"`
	Bio             *string `json:"bio"`
	ProfilePictures *string `json:"profile_pictures"`
}
