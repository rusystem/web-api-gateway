package domain

type Warehouse struct {
	ID                int64                  `gorm:"primaryKey" json:"id"` // Уникальный идентификатор склада
	Name              string                 `json:"name"`                 // Название склада
	Address           string                 `json:"address"`              // Адрес склада
	ResponsiblePerson string                 `json:"responsible_person"`   // Ответственное лицо за склад
	Phone             string                 `json:"phone"`                // Контактный телефон склада
	Email             string                 `json:"email"`                // Электронная почта для связи
	MaxCapacity       int64                  `json:"max_capacity"`         // Максимальная вместимость склада
	CurrentOccupancy  int64                  `json:"current_occupancy"`    // Текущая заполняемость склада
	OtherFields       map[string]interface{} `json:"other_fields"`         // Дополнительные пользовательские поля
	Country           string                 `json:"country"`              // Страна склада
	CompanyId         int64                  `json:"company_id"`           // Уникальный идентификатор компании
}

type InputWarehouse struct {
	Name              string                 `json:"name" binding:"required,min=5,max=140" example:"Название склада"`                           // Название склада
	Address           string                 `json:"address" binding:"required,min=5,max=140" example:"Адрес склада"`                           // Адрес склада
	ResponsiblePerson string                 `json:"responsible_person" binding:"required,min=5,max=140" example:"Ответственное лицо за склад"` // Ответственное лицо за склад
	Phone             string                 `json:"phone" binding:"required,min=5" example:"Контактный телефон склада"`                        // Контактный телефон склада
	Email             string                 `json:"email" example:"Электронная почта для связи"`                                               // Электронная почта для связи
	MaxCapacity       int64                  `json:"max_capacity" example:"100"`                                                                // Максимальная вместимость склада
	CurrentOccupancy  int64                  `json:"current_occupancy" example:"50"`                                                            // Текущая заполняемость склада
	OtherFields       map[string]interface{} `json:"other_fields"`                                                                              // Дополнительные пользовательские поля
	Country           string                 `json:"country" example:"Страна склада"`                                                           // Страна склада
}
