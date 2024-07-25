package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/internal/repository"
	"github.com/rusystem/web-api-gateway/pkg/auth"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type Auth interface {
	SignIn(c *gin.Context, input domain.SignIn) (domain.TokenResponse, error)
	SignOut(c *gin.Context, u domain.User) error
	SignUp(c *gin.Context, input domain.SignUp) (string, bool, error)
	RefreshTokens(c *gin.Context, refreshToken string) (domain.TokenResponse, error)
	ValidateAccessToken(c *gin.Context, token, userAgent, ip string) (domain.JWTInfo, bool, error)
	GetAuthUserByToken(c *gin.Context, accessToken, userAgent, ip string) (domain.User, bool, error)
	GetAuthUser(c *gin.Context, userId, companyId string) (domain.User, error)
}

type AuthServices struct {
	cfg          *config.Config
	repo         *repository.Repository
	tokenManager auth.TokenManager
}

func NewAuthServices(cfg *config.Config, repo *repository.Repository, tokenManager auth.TokenManager) *AuthServices {
	return &AuthServices{
		cfg:          cfg,
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (as *AuthServices) SignIn(c *gin.Context, input domain.SignIn) (domain.TokenResponse, error) {
	//todo
	return domain.TokenResponse{}, nil
}

func (as *AuthServices) SignOut(c *gin.Context, u domain.User) error {
	//todo
	return nil
}

func (as *AuthServices) SignUp(c *gin.Context, input domain.SignUp) (string, bool, error) {
	//todo
	return "", false, nil
}

func (as *AuthServices) RefreshTokens(c *gin.Context, refreshToken string) (domain.TokenResponse, error) {
	//todo return new access and refresh tokens

	return domain.TokenResponse{}, nil
}

func (as *AuthServices) ValidateAccessToken(c *gin.Context, token, userAgent, ip string) (domain.JWTInfo, bool, error) {
	//todo

	return domain.JWTInfo{}, false, nil
}

func (as *AuthServices) GetAuthUserByToken(c *gin.Context, accessToken, userAgent, ip string) (domain.User, bool, error) {
	//todo

	return domain.User{}, false, nil
}

func (as *AuthServices) GetAuthUser(c *gin.Context, userId, companyId string) (domain.User, error) {
	//todo

	return domain.User{}, nil
}
