package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rusystem/web-api-gateway/pkg/domain"
)

type Company interface {
	GetById(ctx context.Context, id int64) (domain.Company, error)
	IsExist(ctx context.Context, id int64) (bool, error)
}

type CompanyDatabaseRepository struct {
	db *sql.DB
}

func NewCompanyDatabase(db *sql.DB) *CompanyDatabaseRepository {
	return &CompanyDatabaseRepository{
		db: db,
	}
}

func (cdr *CompanyDatabaseRepository) GetById(ctx context.Context, id int64) (domain.Company, error) {
	query := fmt.Sprintf(`SELECT id, name_ru, name_en, country, address, phone, email, website, is_active, created_at, updated_at, is_approved, timezone FROM %s WHERE id = $1`,
		domain.CompaniesTable)

	var company domain.Company
	err := cdr.db.QueryRowContext(ctx, query, id).Scan(
		&company.ID,
		&company.NameRU,
		&company.NameEN,
		&company.Country,
		&company.Address,
		&company.Phone,
		&company.Email,
		&company.Website,
		&company.IsActive,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.IsApproved,
		&company.Timezone,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return company, domain.ErrCompanyNotFound
		}

		return company, err
	}

	return company, nil
}

func (cdr *CompanyDatabaseRepository) IsExist(ctx context.Context, id int64) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)`, domain.CompaniesTable)

	var exists bool
	if err := cdr.db.QueryRowContext(ctx, query, id).Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, domain.ErrCompanyNotFound
		}

		return false, err
	}

	return exists, nil
}
