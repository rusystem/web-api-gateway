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
