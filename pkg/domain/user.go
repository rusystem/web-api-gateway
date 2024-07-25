package domain

import "time"

// User представляет данные о пользователе
type User struct {
	ID                       int64     `db:"id" json:"id"`                                                   // Уникальный идентификатор пользователя
	CompanyID                int64     `db:"company_id" json:"company_id"`                                   // Уникальный идентификатор компании
	Username                 string    `db:"username" json:"username"`                                       // Имя пользователя, уникальное
	Email                    string    `db:"email" json:"email"`                                             // Электронная почта пользователя, уникальная
	Phone                    string    `db:"phone" json:"phone"`                                             // Телефон пользователя
	PasswordHash             string    `db:"password_hash" json:"password_hash"`                             // Хеш пароля пользователя
	CreatedAt                time.Time `db:"created_at" json:"created_at"`                                   // Дата и время создания учетной записи
	UpdatedAt                time.Time `db:"updated_at" json:"updated_at"`                                   // Дата и время последнего обновления учетной записи
	LastLogin                time.Time `db:"last_login" json:"last_login"`                                   // Дата и время последнего входа
	IsActive                 bool      `db:"is_active" json:"is_active"`                                     // Статус активности учетной записи
	Role                     string    `db:"role" json:"role"`                                               // Роль пользователя (например, 'user', 'admin')
	Language                 string    `db:"language" json:"language"`                                       // Язык пользователя
	Country                  string    `db:"country" json:"country"`                                         // Страна пользователя
	IsApproved               bool      `db:"is_approved" json:"is_approved"`                                 // Подтвержден ли пользователь
	IsSendSystemNotification bool      `db:"is_send_system_notification" json:"is_send_system_notification"` // Отправлять ли пользователю уведомления
}
