package response

type SignInResponse struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Lat         string `json:"lat"`
	Lng         string `json:"lang"`
}
