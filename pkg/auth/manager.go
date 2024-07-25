package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"math/rand"
	"time"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(info domain.JWTInfo, ttl time.Duration) (string, error)
	Parse(accessToken string) (domain.JWTInfo, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(info domain.JWTInfo, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(ttl).Unix(),
		Id:        info.UserId,
		Audience:  info.Role,
		Subject:   info.CompanyId,
		Issuer:    info.Fingerprint,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string) (domain.JWTInfo, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return domain.JWTInfo{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return domain.JWTInfo{}, fmt.Errorf("error get user claims from token")
	}

	info := domain.JWTInfo{
		UserId:      claims["jti"].(string),
		Role:        claims["aud"].(string),
		CompanyId:   claims["sub"].(string),
		Fingerprint: claims["iss"].(string),
	}

	return info, nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
