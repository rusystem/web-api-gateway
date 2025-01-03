package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rusystem/web-api-gateway/pkg/domain"
	"time"
)

type User interface {
	GetByUsername(ctx context.Context, username string) (domain.User, error)
	GetSections(ctx context.Context, id int64) ([]string, error)
	GetById(ctx context.Context, id int64) (domain.User, error)
	UpdateLastLogin(ctx context.Context, id int64) error
}

type UserDatabaseRepository struct {
	db *sql.DB
}

func NewUserDatabase(db *sql.DB) *UserDatabaseRepository {
	return &UserDatabaseRepository{
		db: db,
	}
}

func (udr *UserDatabaseRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	query := fmt.Sprintf(`
        SELECT id, company_id, username, email, phone, password_hash, created_at, updated_at, last_login, is_active,
               role, language, country, is_approved, is_send_system_notification, sections, position
        FROM %s
        WHERE username = $1`, domain.UsersTable)

	var user domain.User
	var b []byte
	err := udr.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.CompanyID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
		&user.IsActive,
		&user.Role,
		&user.Language,
		&user.Country,
		&user.IsApproved,
		&user.IsSendSystemNotification,
		&b,
		&user.Position,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, err
	}

	if err = json.Unmarshal(b, &user.Sections); err != nil {
		return domain.User{}, fmt.Errorf("error unmarshalling sections: %v", err)
	}

	return user, nil
}

func (udr *UserDatabaseRepository) GetSections(ctx context.Context, id int64) ([]string, error) {
	sections := make([]string, 0)

	query := fmt.Sprintf(`
        SELECT sections
        FROM %s
        WHERE id = $1`, domain.UsersTable)

	var b []byte
	err := udr.db.QueryRowContext(ctx, query, id).Scan(&b)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sections, domain.ErrUserNotFound
		}

		return sections, err
	}

	if err = json.Unmarshal(b, &sections); err != nil {
		return sections, fmt.Errorf("error unmarshalling sections: %v", err)
	}

	return sections, nil
}

func (udr *UserDatabaseRepository) GetById(ctx context.Context, id int64) (domain.User, error) {
	query := fmt.Sprintf(`
        SELECT id, company_id, username, email, phone, password_hash, created_at, updated_at, last_login, is_active,
               role, language, country, is_approved, is_send_system_notification, sections, position
        FROM %s
        WHERE id = $1`, domain.UsersTable)

	var user domain.User
	var b []byte
	err := udr.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CompanyID,
		&user.Username,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
		&user.IsActive,
		&user.Role,
		&user.Language,
		&user.Country,
		&user.IsApproved,
		&user.IsSendSystemNotification,
		&b,
		&user.Position,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, err
	}

	if err = json.Unmarshal(b, &user.Sections); err != nil {
		return domain.User{}, fmt.Errorf("error unmarshalling sections: %v", err)
	}

	return user, nil
}

func (udr *UserDatabaseRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`UPDATE %s SET last_login = $1 WHERE id = $2`, domain.UsersTable)

	_, err := udr.db.Exec(query, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	return nil
}
