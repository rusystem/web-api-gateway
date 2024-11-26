package service

import (
	"context"
	"errors"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/warehouse"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	tools "github.com/rusystem/web-api-gateway/tools"
)

type Warehouse interface {
	GetById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Warehouse, error)
	Create(ctx context.Context, wh domain.Warehouse) (int64, error)
	Update(ctx context.Context, wh domain.WarehouseUpdate, info domain.JWTInfo) error
	Delete(ctx context.Context, id int64, info domain.JWTInfo) error
	GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.Warehouse, error)
	GetResponsibleUsers(ctx context.Context, companyId int64) ([]domain.UserResponse, error)
}

type WarehouseServices struct {
	cfg *config.Config
	wc  *warehouse.WarehouseClient
}

func NewWarehouseServices(cfg *config.Config, warehouseClient *warehouse.WarehouseClient) *WarehouseServices {
	return &WarehouseServices{
		cfg: cfg,
		wc:  warehouseClient,
	}
}

func (s *WarehouseServices) GetById(ctx context.Context, id int64, info domain.JWTInfo) (domain.Warehouse, error) {
	wh, err := s.wc.GetById(ctx, id)
	if err != nil {
		return domain.Warehouse{}, err
	}

	if wh.CompanyId != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.Warehouse{}, domain.ErrNotAllowed
	}

	return wh, nil
}

func (s *WarehouseServices) Create(ctx context.Context, wh domain.Warehouse) (int64, error) {
	return s.wc.Create(ctx, wh)
}

func (s *WarehouseServices) Update(ctx context.Context, inp domain.WarehouseUpdate, info domain.JWTInfo) error {
	wh, err := s.wc.GetById(ctx, inp.ID)
	if err != nil {
		return err
	}

	if wh.CompanyId != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	if inp.Name != nil {
		wh.Name = *inp.Name
	}

	if inp.Address != nil {
		wh.Address = *inp.Address
	}

	if inp.ResponsiblePerson != nil {
		wh.ResponsiblePerson = *inp.ResponsiblePerson
	}

	if inp.Phone != nil {
		wh.Phone = *inp.Phone
	}

	if inp.Email != nil {
		wh.Email = *inp.Email
	}

	if inp.MaxCapacity != nil {
		wh.MaxCapacity = *inp.MaxCapacity
	}

	if inp.CurrentOccupancy != nil {
		wh.CurrentOccupancy = *inp.CurrentOccupancy
	}

	if inp.OtherFields != nil {
		wh.OtherFields = *inp.OtherFields
	}

	if inp.Country != nil {
		wh.Country = *inp.Country
	}

	return s.wc.Update(ctx, wh)
}

func (s *WarehouseServices) Delete(ctx context.Context, id int64, info domain.JWTInfo) error {
	wh, err := s.wc.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrWarehouseNotFound) {
			return domain.ErrWarehouseNotFound
		}

		return err
	}

	if wh.CompanyId != info.CompanyId && !tools.IsFullAccessSection(info.Sections) {
		return domain.ErrNotAllowed
	}

	return s.wc.Delete(ctx, id)
}

func (s *WarehouseServices) GetListByCompanyId(ctx context.Context, companyId int64) ([]domain.Warehouse, error) {
	return s.wc.GetList(ctx, companyId)
}

func (s *WarehouseServices) GetResponsibleUsers(ctx context.Context, companyId int64) ([]domain.UserResponse, error) {
	var resp []domain.UserResponse

	users, err := s.wc.GetResponsiblePerson(ctx, companyId)
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
