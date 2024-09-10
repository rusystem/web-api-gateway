package domain

import "time"

// Material представляет структуру товара
type Material struct {
	ID                     int64                  `json:"id"`                                             // Уникальный идентификатор записи
	WarehouseID            int64                  `json:"warehouse_id" example:"10"`                      // Id склада
	ItemID                 int64                  `json:"item_id" example:"1001"`                         // Идентификатор товара
	Name                   string                 `json:"name" example:"Steel Beam"`                      // Наименование товара
	ByInvoice              string                 `json:"by_invoice" example:"INV-123456"`                // Накладная на товар
	Article                string                 `json:"article" example:"SB-101"`                       // Артикул товара
	ProductCategory        string                 `json:"product_category" example:"Construction"`        // Категория товара
	Unit                   string                 `json:"unit" example:"pcs"`                             // Единица измерения
	TotalQuantity          int64                  `json:"total_quantity" example:"500"`                   // Общее количество товара
	Volume                 int64                  `json:"volume" example:"20"`                            // Объем товара
	PriceWithoutVAT        float64                `json:"price_without_vat" example:"100.50"`             // Цена без НДС
	TotalWithoutVAT        float64                `json:"total_without_vat" example:"50250.00"`           // Общая стоимость без НДС
	SupplierID             int64                  `json:"supplier_id" example:"2001"`                     // Поставщик товара
	Location               string                 `json:"location" example:"A-12"`                        // Локация на складе
	Contract               time.Time              `json:"contract" example:"2022-01-01T00:00:00Z"`        // Дата договора
	File                   string                 `json:"file" example:"contract.pdf"`                    // Файл, связанный с товаром
	Status                 string                 `json:"status" example:"active"`                        // Статус товара
	Comments               string                 `json:"comments" example:"Urgent delivery"`             // Комментарии
	Reserve                string                 `json:"reserve" example:"50"`                           // Резерв товара
	ReceivedDate           time.Time              `json:"received_date" example:"2022-01-01T00:00:00Z"`   // Дата поступления товара
	LastUpdated            time.Time              `json:"last_updated" example:"2022-01-01T00:00:00Z"`    // Дата последнего обновления информации о товаре
	MinStockLevel          int64                  `json:"min_stock_level" example:"20"`                   // Минимальный уровень запаса
	ExpirationDate         time.Time              `json:"expiration_date" example:"2022-01-01T00:00:00Z"` // Срок годности товара
	ResponsiblePerson      string                 `json:"responsible_person" example:"John Doe"`          // Ответственное лицо за товар
	StorageCost            float64                `json:"storage_cost" example:"500.00"`                  // Стоимость хранения товара
	WarehouseSection       string                 `json:"warehouse_section" example:"Section B"`          // Секция склада, где хранится товар
	Barcode                string                 `json:"barcode" example:"BAR-102938"`                   // Штрих-код товара
	IncomingDeliveryNumber string                 `json:"incoming_delivery_number" example:"DEL-45678"`   // Входящий номер поставки
	OtherFields            map[string]interface{} `json:"other_fields"`                                   // Дополнительные пользовательские поля
	CompanyID              int64                  `json:"company_id"`                                     // Кабинет компании к кому привязан товар
}
