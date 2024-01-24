package requests

type ProfileRequest struct {
	Bio             string `json:"bio,omitempty"`
	ProfilePictures string `json:"profile_pictures,omitempty"`
}

// TODO: UPLOAD IMAGE OBJECT
type UploadImageRequestURL struct {
}

type ProfileFilter struct {
	Username string `json:"username,omitempty"`
}
