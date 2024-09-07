// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: proto/warehouse/warehouse.proto

package warehouse

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                       // Уникальный идентификатор склада
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                    // Название склада
	Address           string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`                                              // Адрес склада
	ResponsiblePerson string `protobuf:"bytes,4,opt,name=responsible_person,json=responsiblePerson,proto3" json:"responsible_person,omitempty"` // Ответственное лицо за склад
	Phone             string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`                                                  // Контактный телефон склада
	Email             string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`                                                  // Электронная почта для связи
	MaxCapacity       int64  `protobuf:"varint,7,opt,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`                  // Максимальная вместимость склада
	CurrentOccupancy  int64  `protobuf:"varint,8,opt,name=current_occupancy,json=currentOccupancy,proto3" json:"current_occupancy,omitempty"`   // Текущая заполняемость склада
	OtherFields       string `protobuf:"bytes,9,opt,name=other_fields,json=otherFields,proto3" json:"other_fields,omitempty"`                   // Дополнительные пользовательские поля
	Country           string `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`                                             // Страна склада
	CompanyId         int64  `protobuf:"varint,11,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`                       // Идентификатор компании
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouse.ProtoReflect.Descriptor instead.
func (*Warehouse) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *Warehouse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Warehouse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Warehouse) GetResponsiblePerson() string {
	if x != nil {
		return x.ResponsiblePerson
	}
	return ""
}

func (x *Warehouse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Warehouse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Warehouse) GetMaxCapacity() int64 {
	if x != nil {
		return x.MaxCapacity
	}
	return 0
}

func (x *Warehouse) GetCurrentOccupancy() int64 {
	if x != nil {
		return x.CurrentOccupancy
	}
	return 0
}

func (x *Warehouse) GetOtherFields() string {
	if x != nil {
		return x.OtherFields
	}
	return ""
}

func (x *Warehouse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Warehouse) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type WarehouseId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *WarehouseId) Reset() {
	*x = WarehouseId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseId) ProtoMessage() {}

func (x *WarehouseId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseId.ProtoReflect.Descriptor instead.
func (*WarehouseId) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *WarehouseId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type WarehouseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warehouses []*Warehouse `protobuf:"bytes,1,rep,name=warehouses,proto3" json:"warehouses,omitempty"`
}

func (x *WarehouseList) Reset() {
	*x = WarehouseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseList) ProtoMessage() {}

func (x *WarehouseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseList.ProtoReflect.Descriptor instead.
func (*WarehouseList) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *WarehouseList) GetWarehouses() []*Warehouse {
	if x != nil {
		return x.Warehouses
	}
	return nil
}

type WarehouseCompanyId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *WarehouseCompanyId) Reset() {
	*x = WarehouseCompanyId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarehouseCompanyId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarehouseCompanyId) ProtoMessage() {}

func (x *WarehouseCompanyId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarehouseCompanyId.ProtoReflect.Descriptor instead.
func (*WarehouseCompanyId) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *WarehouseCompanyId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_warehouse_warehouse_proto protoreflect.FileDescriptor

var file_proto_warehouse_warehouse_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x09, 0x57, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f,
	0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x6e, 0x63, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0d, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0a,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x0a, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x32, 0xb9, 0x02, 0x0a, 0x10, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x1a, 0x16, 0x2e,
	0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x16, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x77, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_warehouse_warehouse_proto_rawDescOnce sync.Once
	file_proto_warehouse_warehouse_proto_rawDescData = file_proto_warehouse_warehouse_proto_rawDesc
)

func file_proto_warehouse_warehouse_proto_rawDescGZIP() []byte {
	file_proto_warehouse_warehouse_proto_rawDescOnce.Do(func() {
		file_proto_warehouse_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_warehouse_warehouse_proto_rawDescData)
	})
	return file_proto_warehouse_warehouse_proto_rawDescData
}

var file_proto_warehouse_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_warehouse_warehouse_proto_goTypes = []any{
	(*Warehouse)(nil),          // 0: warehouse.Warehouse
	(*WarehouseId)(nil),        // 1: warehouse.WarehouseId
	(*WarehouseList)(nil),      // 2: warehouse.WarehouseList
	(*WarehouseCompanyId)(nil), // 3: warehouse.WarehouseCompanyId
	(*emptypb.Empty)(nil),      // 4: google.protobuf.Empty
}
var file_proto_warehouse_warehouse_proto_depIdxs = []int32{
	0, // 0: warehouse.WarehouseList.warehouses:type_name -> warehouse.Warehouse
	0, // 1: warehouse.WarehouseService.Create:input_type -> warehouse.Warehouse
	1, // 2: warehouse.WarehouseService.GetById:input_type -> warehouse.WarehouseId
	0, // 3: warehouse.WarehouseService.Update:input_type -> warehouse.Warehouse
	1, // 4: warehouse.WarehouseService.Delete:input_type -> warehouse.WarehouseId
	3, // 5: warehouse.WarehouseService.GetList:input_type -> warehouse.WarehouseCompanyId
	1, // 6: warehouse.WarehouseService.Create:output_type -> warehouse.WarehouseId
	0, // 7: warehouse.WarehouseService.GetById:output_type -> warehouse.Warehouse
	4, // 8: warehouse.WarehouseService.Update:output_type -> google.protobuf.Empty
	4, // 9: warehouse.WarehouseService.Delete:output_type -> google.protobuf.Empty
	2, // 10: warehouse.WarehouseService.GetList:output_type -> warehouse.WarehouseList
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_warehouse_warehouse_proto_init() }
func file_proto_warehouse_warehouse_proto_init() {
	if File_proto_warehouse_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_warehouse_warehouse_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Warehouse); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WarehouseId); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WarehouseList); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WarehouseCompanyId); i {
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
			RawDescriptor: file_proto_warehouse_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_warehouse_warehouse_proto_goTypes,
		DependencyIndexes: file_proto_warehouse_warehouse_proto_depIdxs,
		MessageInfos:      file_proto_warehouse_warehouse_proto_msgTypes,
	}.Build()
	File_proto_warehouse_warehouse_proto = out.File
	file_proto_warehouse_warehouse_proto_rawDesc = nil
	file_proto_warehouse_warehouse_proto_goTypes = nil
	file_proto_warehouse_warehouse_proto_depIdxs = nil
}
