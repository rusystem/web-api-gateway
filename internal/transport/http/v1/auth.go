package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	tools "github.com/rusystem/web-api-gateway/tools"
	"net/http"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/", h.signIn)
		auth.POST("/refresh", h.refresh)

		authenticated := auth.Group("/", h.userIdentity)
		{
			authenticated.GET("/logout", h.signOut)
		}
	}

	register := api.Group("/register", h.adminIdentity)
	{
		register.POST("/", h.signUp)
	}
}

// @Summary Sign in
// @Tags auth
// @Description Аутентификация пользователя.
// @Description Только у super admin есть возможность авторизоваться под определенной компанией.
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body domain.SignIn true "Необходимо указать данные для аутентификации пользователя."
// @Success 200 {object} domain.TokenResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /auth [POST]
func (h *Handler) signIn(c *gin.Context) {
	var inp domain.SignIn
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	res, err := h.services.Auth.SignIn(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrUserIsNotActive) || errors.Is(err, domain.ErrUserIsNotApproved) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		if errors.Is(err, domain.ErrLoginCredentials) {
			newResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if errors.Is(err, domain.ErrCompanyNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.TokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.ExpiresIn,
	})
}

// @Summary Refresh tokens
// @Tags auth
// @Description Обновление access токена
// @ID refresh-tokens
// @Accept json
// @Produce json
// @Param input body domain.TokensRequest true "Необходимо указать текущие refresh token и sections."
// @Success 200 {object} domain.TokenResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /auth/refresh [POST]
func (h *Handler) refresh(c *gin.Context) {
	var inp domain.TokensRequest
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	res, err := h.services.Auth.RefreshTokens(c, inp.RefreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidRefreshToken) || errors.Is(err, domain.ErrGetIpAddress) {
			newResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.TokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		ExpiresIn:    res.ExpiresIn,
	})
}

// @Summary Sign up
// @Security ApiKeyAuth
// @Tags auth
// @Description Регистрация нового пользователя.
// @Description Только у super admin есть возможность добавлять пользователя в другие компании.
// @Description Только у super admin есть возможность давать роль admin пользователю
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body domain.SignUp true "Необходимо указать данные для регистрации нового пользователя."
// @Success 200 {object} domain.SignUpResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /register [POST]
func (h *Handler) signUp(c *gin.Context) {
	var inp domain.SignUp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInputBody.Error())
		return
	}

	if !tools.IsAllowedRole(inp.Role) {
		newResponse(c, http.StatusBadRequest, domain.ErrRoleNotAllowed.Error())
		return
	}

	info, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, isAdmin, err := h.services.Auth.SignUp(c, inp, info)
	if err != nil {
		if errors.Is(err, domain.ErrUserIsNotActive) || errors.Is(err, domain.ErrUserIsNotApproved) {
			newResponse(c, http.StatusForbidden, err.Error())
			return
		}

		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newResponse(c, http.StatusConflict, err.Error())
			return
		}

		if errors.Is(err, domain.ErrSectionsNotAllowed) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, domain.ErrCompanyNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, domain.SignUpResponse{
		ID:      userId,
		IsAdmin: isAdmin,
	})
}

// @Summary Sign out
// @Security ApiKeyAuth
// @Tags auth
// @Description Выход из аккаунта пользователя
// @ID sign-out
// @Accept json
// @Produce json
// @Success 200 {object} domain.MessageResponse
// @Failure 400,404 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /auth/logout [GET]
func (h *Handler) signOut(c *gin.Context) {
	userInfo, err := getUserInfo(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = h.services.Auth.SignOut(c, userInfo.UserId, userInfo.CompanyId); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, domain.MessageResponse{Message: "success"})
}
