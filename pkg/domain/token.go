package domain

import "time"

type JWTInfo struct {
	UserId      string
	CompanyId   string
	Role        string
	Fingerprint string
}

type RefreshSession struct {
	ID        int       `msgpack:"id"`
	UserID    string    `msgpack:"user_id"`
	CompanyID string    `msgpack:"company_id"`
	Role      string    `msgpack:"role"`
	Token     string    `msgpack:"token"`
	ExpiresAt time.Time `msgpack:"expires_at"`
	Ip        string    `msgpack:"ip"`
}

type TokenResponse struct {
	AccessToken  string   `json:"token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int64    `json:"expires_in"`
	Sections     []string `json:"sections,omitempty"`
}
