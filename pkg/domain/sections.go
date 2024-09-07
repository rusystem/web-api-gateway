package domain

const (
	FullAllAccessSection            = "full_all_access"             // полный доступ
	FullCompanyAccessSection        = "full_company_access"         // полный доступ по управлению в рамках компании
	FullAccessSection               = "full_access"                 // Начальник цеха, Зам начальника цеха, Просчетчик. Доступ ко всем данным.
	OrderCardAccessSection          = "order_card_access"           // Конструктор.Доступ только к информации по карточке заказа: описание изделия, количество, уникальные номера заказа, ответственный менеджер, задачи по заказу.
	ProductionDataAccessSection     = "production_data_access"      // Работник 1, Работник 2, Работник N. Доступ только к информации, необходимой для производства: описание изделия, количество, уникальные номера.
	StatusAndCalculateAccessSection = "status_and_calculate_access" // Менеджер. Доступ к статусу заказа и калькуляции.
	PurchasePlanningAccessSection   = "purchase_planning_access"    // Снабженец. Доступ к заявкам на закуп и планированию закупок.
)

type Section struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SectionCreate struct {
	Name string `json:"name" binding:"required" example:"purchase_planning_access"`
}

type SectionUpdate struct {
	Name string `json:"name" binding:"required" example:"status_and_calculation_access"`
}
