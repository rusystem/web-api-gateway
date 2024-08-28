package domain

type IdResponse struct {
	ID interface{} `json:"id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Response struct {
	ID                string   `json:"id,omitempty"`
	AllowRegistration bool     `json:"allow_registration,omitempty"`
	Message           string   `json:"message,omitempty"`
	Token             string   `json:"token,omitempty"`
	Sections          []string `json:"sections,omitempty"`
	Status            string   `json:"status,omitempty"`
	Name              string   `json:"name,omitempty,omitempty"`
	IsAdmin           bool     `json:"is_admin,omitempty"`
	CreatedUserID     string   `json:"created_user_id,omitempty"`
	Avatar            string   `json:"avatar,omitempty"`
}

type AvatarResponse struct {
	Avatar string `json:"avatar"`
}

type OperatorStatusResponse struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

type UserStatusResponse struct {
	Status string `json:"status"`
}

type AllowRegistrationResponse struct {
	AllowRegistration bool `json:"allow_registration"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type SignUpResponse struct {
	ID      int64 `json:"id"`
	IsAdmin bool  `json:"is_admin"`
}

type CreateUserResponse struct {
	ID            string `json:"id"`
	CreatedUserID string `json:"created_user_id"`
}
