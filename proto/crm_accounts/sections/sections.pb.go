// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: proto/sections.proto

package sections

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

type SectionsId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *SectionsId) Reset() {
	*x = SectionsId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sections_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionsId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionsId) ProtoMessage() {}

func (x *SectionsId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sections_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionsId.ProtoReflect.Descriptor instead.
func (*SectionsId) Descriptor() ([]byte, []int) {
	return file_proto_sections_proto_rawDescGZIP(), []int{0}
}

func (x *SectionsId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Section struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // Уникальный идентификатор секции
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Название секции
}

func (x *Section) Reset() {
	*x = Section{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sections_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Section) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Section) ProtoMessage() {}

func (x *Section) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sections_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Section.ProtoReflect.Descriptor instead.
func (*Section) Descriptor() ([]byte, []int) {
	return file_proto_sections_proto_rawDescGZIP(), []int{1}
}

func (x *Section) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Section) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SectionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sections []*Section `protobuf:"bytes,1,rep,name=sections,proto3" json:"sections,omitempty"`
}

func (x *SectionList) Reset() {
	*x = SectionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sections_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionList) ProtoMessage() {}

func (x *SectionList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sections_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionList.ProtoReflect.Descriptor instead.
func (*SectionList) Descriptor() ([]byte, []int) {
	return file_proto_sections_proto_rawDescGZIP(), []int{2}
}

func (x *SectionList) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

var File_proto_sections_proto protoreflect.FileDescriptor

var file_proto_sections_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0x8a, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x64,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x17, 0x5a, 0x15, 0x2e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sections_proto_rawDescOnce sync.Once
	file_proto_sections_proto_rawDescData = file_proto_sections_proto_rawDesc
)

func file_proto_sections_proto_rawDescGZIP() []byte {
	file_proto_sections_proto_rawDescOnce.Do(func() {
		file_proto_sections_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sections_proto_rawDescData)
	})
	return file_proto_sections_proto_rawDescData
}

var file_proto_sections_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_sections_proto_goTypes = []any{
	(*SectionsId)(nil),    // 0: proto.SectionsId
	(*Section)(nil),       // 1: proto.Section
	(*SectionList)(nil),   // 2: proto.SectionList
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_proto_sections_proto_depIdxs = []int32{
	1, // 0: proto.SectionList.sections:type_name -> proto.Section
	0, // 1: proto.SectionsService.GetById:input_type -> proto.SectionsId
	1, // 2: proto.SectionsService.Create:input_type -> proto.Section
	1, // 3: proto.SectionsService.Update:input_type -> proto.Section
	0, // 4: proto.SectionsService.Delete:input_type -> proto.SectionsId
	3, // 5: proto.SectionsService.GetList:input_type -> google.protobuf.Empty
	1, // 6: proto.SectionsService.GetById:output_type -> proto.Section
	0, // 7: proto.SectionsService.Create:output_type -> proto.SectionsId
	3, // 8: proto.SectionsService.Update:output_type -> google.protobuf.Empty
	3, // 9: proto.SectionsService.Delete:output_type -> google.protobuf.Empty
	2, // 10: proto.SectionsService.GetList:output_type -> proto.SectionList
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_sections_proto_init() }
func file_proto_sections_proto_init() {
	if File_proto_sections_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sections_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SectionsId); i {
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
		file_proto_sections_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Section); i {
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
		file_proto_sections_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SectionList); i {
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
			RawDescriptor: file_proto_sections_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sections_proto_goTypes,
		DependencyIndexes: file_proto_sections_proto_depIdxs,
		MessageInfos:      file_proto_sections_proto_msgTypes,
	}.Build()
	File_proto_sections_proto = out.File
	file_proto_sections_proto_rawDesc = nil
	file_proto_sections_proto_goTypes = nil
	file_proto_sections_proto_depIdxs = nil
}
