package service

import (
	"context"
	"errors"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/accounts"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	tools "github.com/rusystem/web-api-gateway/tool"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetById(ctx context.Context, id int64) (domain.User, error)
	UpdateProfile(ctx context.Context, user domain.UserProfileUpdate, info domain.JWTInfo) error
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, user domain.UserUpdate, info domain.JWTInfo) error
	Delete(ctx context.Context, id int64, info domain.JWTInfo) error
	GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.UserResponse, error)
}

type UserService struct {
	cfg        *config.Config
	userClient *grpc.UserAccountsClient
}

func NewUserServices(cfg *config.Config, userClient *grpc.UserAccountsClient) *UserService {
	return &UserService{
		cfg:        cfg,
		userClient: userClient,
	}
}

func (su *UserService) GetById(ctx context.Context, id int64) (domain.User, error) {
	return su.userClient.GetById(ctx, id)
}

func (su *UserService) UpdateProfile(ctx context.Context, req domain.UserProfileUpdate, info domain.JWTInfo) error {
	user, err := su.userClient.GetById(ctx, info.UserId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return domain.ErrUserNotFound
		}

		return err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if req.Password != nil {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return domain.ErrUpdateUser
		}

		user.PasswordHash = string(hashedPass)
	}

	if req.Country != nil {
		user.Country = *req.Country
	}

	return su.userClient.Update(ctx, user)
}

func (su *UserService) Create(ctx context.Context, user domain.User) (int64, error) {
	return su.userClient.Create(ctx, user)
}

func (su *UserService) Update(ctx context.Context, req domain.UserUpdate, info domain.JWTInfo) error {
	user, err := su.userClient.GetById(ctx, *req.ID)
	if err != nil {
		return domain.ErrUserNotFound
	}

	if user.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if req.Password != nil {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return domain.ErrUpdateUser
		}

		user.PasswordHash = string(hashedPass)
	}

	if req.Language != nil {
		user.Language = *req.Language
	}

	if req.Country != nil {
		user.Country = *req.Country
	}

	if req.Position != nil {
		user.Position = *req.Position
	}

	if req.IsSendSystemNotification != nil {
		user.IsSendSystemNotification = *req.IsSendSystemNotification
	}

	if req.Sections != nil {
		if tools.IsFullAccessSection(*req.Sections) {
			return domain.ErrNotAllowed
		}

		user.Sections = *req.Sections
	}

	if req.Role != nil {
		if !tools.IsFullAccessSection(info.Sections) {
			return domain.ErrNotAllowed
		}

		user.Role = *req.Role
	}

	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if req.IsApproved != nil {
		user.IsApproved = *req.IsApproved
	}

	return su.userClient.Update(ctx, user)
}

func (su *UserService) Delete(ctx context.Context, id int64, info domain.JWTInfo) error {
	user, err := su.userClient.GetById(ctx, id)
	if err != nil {
		return err
	}

	if user.CompanyID != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return su.userClient.Delete(ctx, id)
}

func (su *UserService) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.UserResponse, error) {
	var resp []domain.UserResponse

	users, err := su.userClient.GetListByCompanyId(ctx, companyId)
	if err != nil {
		return resp, err
	}

	for _, v := range users {
		resp = append(resp, domain.UserResponse{
			ID:                       v.ID,
			CompanyID:                v.CompanyID,
			Username:                 v.Username,
			Name:                     v.Name,
			Email:                    v.Email,
			Phone:                    v.Phone,
			CreatedAt:                v.CreatedAt,
			UpdatedAt:                v.UpdatedAt,
			LastLogin:                v.LastLogin.Time,
			IsActive:                 v.IsActive,
			Role:                     v.Role,
			Language:                 v.Language,
			Country:                  v.Country,
			IsApproved:               v.IsApproved,
			IsSendSystemNotification: v.IsSendSystemNotification,
			Sections:                 v.Sections,
			Position:                 v.Position,
		})
	}

	return resp, nil
}
