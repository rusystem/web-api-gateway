package domain

type SignIn struct {
	Email    string `json:"email" binding:"required,email,min=5,max=140" example:"dmitry@test.com"`
	Password string `json:"password" binding:"required,min=8,max=255" example:"12345678"`
}

type EmailInput struct {
	Email string `json:"email" binding:"required,email,min=5,max=140" example:"hassans@latro.com"`
}

type SignUp struct {
	Email    string `json:"email" binding:"required,email,min=5,max=140" example:"dmitry@test.com"`
	Phone    string `json:"phone" binding:"required,min=7,max=140" example:"+77777777777"`
	Name     string `json:"name" binding:"required,min=1,max=140" example:"dmitry"`
	Password string `json:"password" binding:"required,min=8,max=255" example:"12345678"`
	Avatar   string `json:"avatar" example:"/example.jpg"`
	Position string `json:"position" example:"test"`
	Language string `json:"language" example:"ru"`
	Country  string `json:"country" example:"KZ"`
}

type TokensRequest struct {
	RefreshToken string   `json:"refresh_token" binding:"required"`
	Sections     []string `json:"sections"`
}
