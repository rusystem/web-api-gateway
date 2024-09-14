// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: proto/materials/materials.proto

package materials

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                                         // Уникальный идентификатор записи
	WarehouseId            int64                  `protobuf:"varint,2,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`                                    // Id склада
	ItemId                 int64                  `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`                                                   // Идентификатор товара
	Name                   string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                                                                      // Наименование товара
	ByInvoice              string                 `protobuf:"bytes,5,opt,name=by_invoice,json=byInvoice,proto3" json:"by_invoice,omitempty"`                                           // Накладная на товар
	Article                string                 `protobuf:"bytes,6,opt,name=article,proto3" json:"article,omitempty"`                                                                // Артикул товара
	ProductCategory        string                 `protobuf:"bytes,7,opt,name=product_category,json=productCategory,proto3" json:"product_category,omitempty"`                         // Категория товара
	Unit                   string                 `protobuf:"bytes,8,opt,name=unit,proto3" json:"unit,omitempty"`                                                                      // Единица измерения
	TotalQuantity          int64                  `protobuf:"varint,9,opt,name=total_quantity,json=totalQuantity,proto3" json:"total_quantity,omitempty"`                              // Общее количество товара
	Volume                 int64                  `protobuf:"varint,10,opt,name=volume,proto3" json:"volume,omitempty"`                                                                // Объем товара
	PriceWithoutVat        float64                `protobuf:"fixed64,11,opt,name=price_without_vat,json=priceWithoutVat,proto3" json:"price_without_vat,omitempty"`                    // Цена без НДС
	TotalWithoutVat        float64                `protobuf:"fixed64,12,opt,name=total_without_vat,json=totalWithoutVat,proto3" json:"total_without_vat,omitempty"`                    // Общая стоимость без НДС
	SupplierId             int64                  `protobuf:"varint,13,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`                                      // Поставщик товара
	Location               string                 `protobuf:"bytes,14,opt,name=location,proto3" json:"location,omitempty"`                                                             // Локация на складе
	Contract               *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=contract,proto3" json:"contract,omitempty"`                                                             // Дата договора в формате строки
	File                   string                 `protobuf:"bytes,16,opt,name=file,proto3" json:"file,omitempty"`                                                                     // Файл, связанный с товаром
	Status                 string                 `protobuf:"bytes,17,opt,name=status,proto3" json:"status,omitempty"`                                                                 // Статус товара
	Comments               string                 `protobuf:"bytes,18,opt,name=comments,proto3" json:"comments,omitempty"`                                                             // Комментарии
	Reserve                string                 `protobuf:"bytes,19,opt,name=reserve,proto3" json:"reserve,omitempty"`                                                               // Резерв товара
	ReceivedDate           *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=received_date,json=receivedDate,proto3" json:"received_date,omitempty"`                                 // Дата поступления товара в формате строки
	LastUpdated            *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`                                    // Дата последнего обновления информации о товаре в формате строки
	MinStockLevel          int64                  `protobuf:"varint,22,opt,name=min_stock_level,json=minStockLevel,proto3" json:"min_stock_level,omitempty"`                           // Минимальный уровень запаса
	ExpirationDate         *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`                           // Срок годности товара в формате строки
	ResponsiblePerson      string                 `protobuf:"bytes,24,opt,name=responsible_person,json=responsiblePerson,proto3" json:"responsible_person,omitempty"`                  // Ответственное лицо за товар
	StorageCost            float64                `protobuf:"fixed64,25,opt,name=storage_cost,json=storageCost,proto3" json:"storage_cost,omitempty"`                                  // Стоимость хранения товара
	WarehouseSection       string                 `protobuf:"bytes,26,opt,name=warehouse_section,json=warehouseSection,proto3" json:"warehouse_section,omitempty"`                     // Секция склада, где хранится товар
	IncomingDeliveryNumber string                 `protobuf:"bytes,27,opt,name=incoming_delivery_number,json=incomingDeliveryNumber,proto3" json:"incoming_delivery_number,omitempty"` // Входящий номер поставки
	OtherFields            string                 `protobuf:"bytes,28,opt,name=other_fields,json=otherFields,proto3" json:"other_fields,omitempty"`                                    // Дополнительные пользовательские поля
	CompanyId              int64                  `protobuf:"varint,29,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`                                         // Кабинет компании к кому привязан товар
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_materials_materials_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_proto_materials_materials_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_proto_materials_materials_proto_rawDescGZIP(), []int{0}
}

func (x *Material) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Material) GetWarehouseId() int64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *Material) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Material) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Material) GetByInvoice() string {
	if x != nil {
		return x.ByInvoice
	}
	return ""
}

func (x *Material) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *Material) GetProductCategory() string {
	if x != nil {
		return x.ProductCategory
	}
	return ""
}

func (x *Material) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Material) GetTotalQuantity() int64 {
	if x != nil {
		return x.TotalQuantity
	}
	return 0
}

func (x *Material) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Material) GetPriceWithoutVat() float64 {
	if x != nil {
		return x.PriceWithoutVat
	}
	return 0
}

func (x *Material) GetTotalWithoutVat() float64 {
	if x != nil {
		return x.TotalWithoutVat
	}
	return 0
}

func (x *Material) GetSupplierId() int64 {
	if x != nil {
		return x.SupplierId
	}
	return 0
}

func (x *Material) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Material) GetContract() *timestamppb.Timestamp {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *Material) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Material) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Material) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

func (x *Material) GetReserve() string {
	if x != nil {
		return x.Reserve
	}
	return ""
}

func (x *Material) GetReceivedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedDate
	}
	return nil
}

func (x *Material) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Material) GetMinStockLevel() int64 {
	if x != nil {
		return x.MinStockLevel
	}
	return 0
}

func (x *Material) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

func (x *Material) GetResponsiblePerson() string {
	if x != nil {
		return x.ResponsiblePerson
	}
	return ""
}

func (x *Material) GetStorageCost() float64 {
	if x != nil {
		return x.StorageCost
	}
	return 0
}

func (x *Material) GetWarehouseSection() string {
	if x != nil {
		return x.WarehouseSection
	}
	return ""
}

func (x *Material) GetIncomingDeliveryNumber() string {
	if x != nil {
		return x.IncomingDeliveryNumber
	}
	return ""
}

func (x *Material) GetOtherFields() string {
	if x != nil {
		return x.OtherFields
	}
	return ""
}

func (x *Material) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type MaterialId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ItemId int64 `protobuf:"varint,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (x *MaterialId) Reset() {
	*x = MaterialId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_materials_materials_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialId) ProtoMessage() {}

func (x *MaterialId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_materials_materials_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialId.ProtoReflect.Descriptor instead.
func (*MaterialId) Descriptor() ([]byte, []int) {
	return file_proto_materials_materials_proto_rawDescGZIP(), []int{1}
}

func (x *MaterialId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MaterialId) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

type MaterialList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Materials []*Material `protobuf:"bytes,1,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *MaterialList) Reset() {
	*x = MaterialList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_materials_materials_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialList) ProtoMessage() {}

func (x *MaterialList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_materials_materials_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialList.ProtoReflect.Descriptor instead.
func (*MaterialList) Descriptor() ([]byte, []int) {
	return file_proto_materials_materials_proto_rawDescGZIP(), []int{2}
}

func (x *MaterialList) GetMaterials() []*Material {
	if x != nil {
		return x.Materials
	}
	return nil
}

type MaterialParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit     int64 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset    int64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	CompanyId int64 `protobuf:"varint,3,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
}

func (x *MaterialParams) Reset() {
	*x = MaterialParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_materials_materials_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaterialParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaterialParams) ProtoMessage() {}

func (x *MaterialParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_materials_materials_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaterialParams.ProtoReflect.Descriptor instead.
func (*MaterialParams) Descriptor() ([]byte, []int) {
	return file_proto_materials_materials_proto_rawDescGZIP(), []int{3}
}

func (x *MaterialParams) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *MaterialParams) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *MaterialParams) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

var File_proto_materials_materials_proto protoreflect.FileDescriptor

var file_proto_materials_materials_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x08, 0x0a, 0x08, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x5f, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x79, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f,
	0x76, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x56, 0x61, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x76, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x56, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x12, 0x3f, 0x0a,
	0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3d,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x69, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0c, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x5c,
	0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x32, 0xdb, 0x09, 0x0a,
	0x0f, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x3d,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x12, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x17, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x49, 0x64, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49,
	0x64, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x47,
	0x0a, 0x16, 0x4d, 0x6f, 0x76, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x54,
	0x6f, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x15, 0x2e,
	0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x41, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x4c, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x64, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x47, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2e,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_materials_materials_proto_rawDescOnce sync.Once
	file_proto_materials_materials_proto_rawDescData = file_proto_materials_materials_proto_rawDesc
)

func file_proto_materials_materials_proto_rawDescGZIP() []byte {
	file_proto_materials_materials_proto_rawDescOnce.Do(func() {
		file_proto_materials_materials_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_materials_materials_proto_rawDescData)
	})
	return file_proto_materials_materials_proto_rawDescData
}

var file_proto_materials_materials_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_materials_materials_proto_goTypes = []any{
	(*Material)(nil),              // 0: materials.Material
	(*MaterialId)(nil),            // 1: materials.MaterialId
	(*MaterialList)(nil),          // 2: materials.MaterialList
	(*MaterialParams)(nil),        // 3: materials.MaterialParams
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_proto_materials_materials_proto_depIdxs = []int32{
	4,  // 0: materials.Material.contract:type_name -> google.protobuf.Timestamp
	4,  // 1: materials.Material.received_date:type_name -> google.protobuf.Timestamp
	4,  // 2: materials.Material.last_updated:type_name -> google.protobuf.Timestamp
	4,  // 3: materials.Material.expiration_date:type_name -> google.protobuf.Timestamp
	0,  // 4: materials.MaterialList.materials:type_name -> materials.Material
	0,  // 5: materials.MaterialService.CreatePlanning:input_type -> materials.Material
	0,  // 6: materials.MaterialService.UpdatePlanning:input_type -> materials.Material
	1,  // 7: materials.MaterialService.DeletePlanning:input_type -> materials.MaterialId
	1,  // 8: materials.MaterialService.GetPlanning:input_type -> materials.MaterialId
	3,  // 9: materials.MaterialService.GetListPlanning:input_type -> materials.MaterialParams
	1,  // 10: materials.MaterialService.MovePlanningToPurchased:input_type -> materials.MaterialId
	0,  // 11: materials.MaterialService.CreatePurchased:input_type -> materials.Material
	0,  // 12: materials.MaterialService.UpdatePurchased:input_type -> materials.Material
	1,  // 13: materials.MaterialService.DeletePurchased:input_type -> materials.MaterialId
	1,  // 14: materials.MaterialService.GetPurchased:input_type -> materials.MaterialId
	3,  // 15: materials.MaterialService.GetListPurchased:input_type -> materials.MaterialParams
	1,  // 16: materials.MaterialService.MovePurchasedToArchive:input_type -> materials.MaterialId
	1,  // 17: materials.MaterialService.GetPlanningArchive:input_type -> materials.MaterialId
	1,  // 18: materials.MaterialService.GetPurchasedArchive:input_type -> materials.MaterialId
	3,  // 19: materials.MaterialService.GetListPlanningArchive:input_type -> materials.MaterialParams
	3,  // 20: materials.MaterialService.GetListPurchasedArchive:input_type -> materials.MaterialParams
	1,  // 21: materials.MaterialService.DeletePlanningArchive:input_type -> materials.MaterialId
	1,  // 22: materials.MaterialService.DeletePurchasedArchive:input_type -> materials.MaterialId
	1,  // 23: materials.MaterialService.CreatePlanning:output_type -> materials.MaterialId
	5,  // 24: materials.MaterialService.UpdatePlanning:output_type -> google.protobuf.Empty
	5,  // 25: materials.MaterialService.DeletePlanning:output_type -> google.protobuf.Empty
	0,  // 26: materials.MaterialService.GetPlanning:output_type -> materials.Material
	2,  // 27: materials.MaterialService.GetListPlanning:output_type -> materials.MaterialList
	1,  // 28: materials.MaterialService.MovePlanningToPurchased:output_type -> materials.MaterialId
	1,  // 29: materials.MaterialService.CreatePurchased:output_type -> materials.MaterialId
	5,  // 30: materials.MaterialService.UpdatePurchased:output_type -> google.protobuf.Empty
	5,  // 31: materials.MaterialService.DeletePurchased:output_type -> google.protobuf.Empty
	0,  // 32: materials.MaterialService.GetPurchased:output_type -> materials.Material
	2,  // 33: materials.MaterialService.GetListPurchased:output_type -> materials.MaterialList
	5,  // 34: materials.MaterialService.MovePurchasedToArchive:output_type -> google.protobuf.Empty
	0,  // 35: materials.MaterialService.GetPlanningArchive:output_type -> materials.Material
	0,  // 36: materials.MaterialService.GetPurchasedArchive:output_type -> materials.Material
	2,  // 37: materials.MaterialService.GetListPlanningArchive:output_type -> materials.MaterialList
	2,  // 38: materials.MaterialService.GetListPurchasedArchive:output_type -> materials.MaterialList
	5,  // 39: materials.MaterialService.DeletePlanningArchive:output_type -> google.protobuf.Empty
	5,  // 40: materials.MaterialService.DeletePurchasedArchive:output_type -> google.protobuf.Empty
	23, // [23:41] is the sub-list for method output_type
	5,  // [5:23] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_materials_materials_proto_init() }
func file_proto_materials_materials_proto_init() {
	if File_proto_materials_materials_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_materials_materials_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Material); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_materials_materials_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MaterialId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_materials_materials_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MaterialList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_materials_materials_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MaterialParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_materials_materials_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_materials_materials_proto_goTypes,
		DependencyIndexes: file_proto_materials_materials_proto_depIdxs,
		MessageInfos:      file_proto_materials_materials_proto_msgTypes,
	}.Build()
	File_proto_materials_materials_proto = out.File
	file_proto_materials_materials_proto_rawDesc = nil
	file_proto_materials_materials_proto_goTypes = nil
	file_proto_materials_materials_proto_depIdxs = nil
}