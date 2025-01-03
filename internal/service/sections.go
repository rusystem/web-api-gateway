package service

import (
	"context"
	"errors"
	"github.com/rusystem/web-api-gateway/internal/config"
	"github.com/rusystem/web-api-gateway/pkg/client/grpc/accounts"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	tools "github.com/rusystem/web-api-gateway/tools"
)

type Sections interface {
	GetById(ctx context.Context, id int64) (domain.Section, error)
	Create(ctx context.Context, section domain.Section) (int64, error)
	Update(ctx context.Context, section domain.Section) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, info domain.JWTInfo) ([]domain.Section, error)
}

type SectionsService struct {
	cfg            *config.Config
	sectionsClient *grpc.SectionsAccountsClient
}

func NewSectionsService(cfg *config.Config, sectionsClient *grpc.SectionsAccountsClient) *SectionsService {
	return &SectionsService{
		cfg:            cfg,
		sectionsClient: sectionsClient,
	}
}

func (s *SectionsService) GetById(ctx context.Context, id int64) (domain.Section, error) {
	return s.sectionsClient.GetById(ctx, id)
}

func (s *SectionsService) Create(ctx context.Context, section domain.Section) (int64, error) {
	return s.sectionsClient.Create(ctx, section)
}

func (s *SectionsService) Update(ctx context.Context, section domain.Section) error {
	oldSection, err := s.sectionsClient.GetById(ctx, section.Id)
	if err != nil {
		return err
	}

	if oldSection.Name == domain.SectionFullAllAccess {
		return domain.ErrNotAllowed
	}

	return s.sectionsClient.Update(ctx, section)
}

func (s *SectionsService) Delete(ctx context.Context, id int64) error {
	section, err := s.sectionsClient.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrSectionNotFound) {
			return domain.ErrSectionNotFound
		}

		return err
	}

	if section.Name == domain.SectionFullAllAccess {
		return domain.ErrNotAllowed
	}

	return s.sectionsClient.Delete(ctx, id)
}

func (s *SectionsService) List(ctx context.Context, info domain.JWTInfo) ([]domain.Section, error) {
	sections, err := s.sectionsClient.GetList(ctx)
	if err != nil {
		return nil, err
	}

	if !tools.IsFullAccessSection(info.Sections) {
		tools.RemoveFullAccessSection(sections, domain.SectionFullAllAccess)
	}

	return s.sectionsClient.GetList(ctx)
}
