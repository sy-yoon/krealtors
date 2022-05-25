// protos/v1/realtor/realtor.proto
//protoc -I . --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative protos/v1/realtor/realtor.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/v1/realtor/realtor.proto

package realtor

import (
	common "github.com/sy-yoon/krealtors/protos/v1/common"
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

type ReItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ReType        int32               `protobuf:"varint,3,opt,name=re_type,json=reType,proto3" json:"re_type,omitempty"`
	TxType        int32               `protobuf:"varint,4,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	AvailableDate string              `protobuf:"bytes,5,opt,name=available_date,json=availableDate,proto3" json:"available_date,omitempty"`
	UserId        string              `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CityId        string              `protobuf:"bytes,7,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Address       string              `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Thumbnail     string              `protobuf:"bytes,9,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Price         int64               `protobuf:"varint,10,opt,name=price,proto3" json:"price,omitempty"`
	Bedroom       int32               `protobuf:"varint,11,opt,name=bedroom,proto3" json:"bedroom,omitempty"`
	Bathroom      int32               `protobuf:"varint,12,opt,name=bathroom,proto3" json:"bathroom,omitempty"`
	Parking       int32               `protobuf:"varint,13,opt,name=parking,proto3" json:"parking,omitempty"`
	Garage        bool                `protobuf:"varint,14,opt,name=garage,proto3" json:"garage,omitempty"`
	Furnished     bool                `protobuf:"varint,15,opt,name=furnished,proto3" json:"furnished,omitempty"`
	Facing        int32               `protobuf:"varint,16,opt,name=facing,proto3" json:"facing,omitempty"`
	Area          int32               `protobuf:"varint,17,opt,name=area,proto3" json:"area,omitempty"`
	Content       string              `protobuf:"bytes,18,opt,name=content,proto3" json:"content,omitempty"`
	GeoLocation   *common.GeoLocation `protobuf:"bytes,19,opt,name=geo_location,json=geoLocation,proto3" json:"geo_location,omitempty"`
	Images        string              `protobuf:"bytes,20,opt,name=images,proto3" json:"images,omitempty"`
	CreatedDate   string              `protobuf:"bytes,21,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	UpdatedDate   string              `protobuf:"bytes,22,opt,name=updated_date,json=updatedDate,proto3" json:"updated_date,omitempty"`
}


func (x *ReItem) Reset() {
	*x = ReItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_realtor_realtor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReItem) ProtoMessage() {}

func (x *ReItem) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_realtor_realtor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReItem.ProtoReflect.Descriptor instead.
func (*ReItem) Descriptor() ([]byte, []int) {
	return file_protos_v1_realtor_realtor_proto_rawDescGZIP(), []int{0}
}

func (x *ReItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReItem) GetReType() int32 {
	if x != nil {
		return x.ReType
	}
	return 0
}

func (x *ReItem) GetTxType() int32 {
	if x != nil {
		return x.TxType
	}
	return 0
}

func (x *ReItem) GetAvailableDate() string {
	if x != nil {
		return x.AvailableDate
	}
	return ""
}

func (x *ReItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReItem) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *ReItem) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ReItem) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *ReItem) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ReItem) GetBedroom() int32 {
	if x != nil {
		return x.Bedroom
	}
	return 0
}

func (x *ReItem) GetBathroom() int32 {
	if x != nil {
		return x.Bathroom
	}
	return 0
}

func (x *ReItem) GetParking() int32 {
	if x != nil {
		return x.Parking
	}
	return 0
}

func (x *ReItem) GetGarage() bool {
	if x != nil {
		return x.Garage
	}
	return false
}

func (x *ReItem) GetFurnished() bool {
	if x != nil {
		return x.Furnished
	}
	return false
}

func (x *ReItem) GetFacing() int32 {
	if x != nil {
		return x.Facing
	}
	return 0
}

func (x *ReItem) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *ReItem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReItem) GetGeoLocation() *common.GeoLocation {
	if x != nil {
		return x.GeoLocation
	}
	return nil
}

func (x *ReItem) GetImages() string {
	if x != nil {
		return x.Images
	}
	return ""
}

func (x *ReItem) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *ReItem) GetUpdatedDate() string {
	if x != nil {
		return x.UpdatedDate
	}
	return ""
}

type ReItemHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ReType        int32               `protobuf:"varint,3,opt,name=re_type,json=reType,proto3" json:"re_type,omitempty"`
	TxType        int32               `protobuf:"varint,4,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	UserId        string              `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CityId        string              `protobuf:"bytes,6,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Thumbnail     string              `protobuf:"bytes,7,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Price         int64               `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
	Bedroom       int32               `protobuf:"varint,9,opt,name=bedroom,proto3" json:"bedroom,omitempty"`
	Bathroom      int32               `protobuf:"varint,10,opt,name=bathroom,proto3" json:"bathroom,omitempty"`
	AvailableDate string              `protobuf:"bytes,11,opt,name=available_date,json=availableDate,proto3" json:"available_date,omitempty"`
	Address       string              `protobuf:"bytes,12,opt,name=address,proto3" json:"address,omitempty"`
	GeoLocation   *common.GeoLocation `protobuf:"bytes,13,opt,name=geo_location,json=geoLocation,proto3" json:"geo_location,omitempty"`
	CreatedDate   string              `protobuf:"bytes,14,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	UpdatedDate   string              `protobuf:"bytes,15,opt,name=updated_date,json=updatedDate,proto3" json:"updated_date,omitempty"`
}

func (x *ReItemHeader) Reset() {
	*x = ReItemHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_realtor_realtor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReItemHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReItemHeader) ProtoMessage() {}

func (x *ReItemHeader) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_realtor_realtor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReItemHeader.ProtoReflect.Descriptor instead.
func (*ReItemHeader) Descriptor() ([]byte, []int) {
	return file_protos_v1_realtor_realtor_proto_rawDescGZIP(), []int{1}
}

func (x *ReItemHeader) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReItemHeader) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReItemHeader) GetReType() int32 {
	if x != nil {
		return x.ReType
	}
	return 0
}

func (x *ReItemHeader) GetTxType() int32 {
	if x != nil {
		return x.TxType
	}
	return 0
}

func (x *ReItemHeader) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReItemHeader) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *ReItemHeader) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *ReItemHeader) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ReItemHeader) GetBedroom() int32 {
	if x != nil {
		return x.Bedroom
	}
	return 0
}

func (x *ReItemHeader) GetBathroom() int32 {
	if x != nil {
		return x.Bathroom
	}
	return 0
}

func (x *ReItemHeader) GetAvailableDate() string {
	if x != nil {
		return x.AvailableDate
	}
	return ""
}

func (x *ReItemHeader) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ReItemHeader) GetGeoLocation() *common.GeoLocation {
	if x != nil {
		return x.GeoLocation
	}
	return nil
}

func (x *ReItemHeader) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *ReItemHeader) GetUpdatedDate() string {
	if x != nil {
		return x.UpdatedDate
	}
	return ""
}

type GetReItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReItemRequest) Reset() {
	*x = GetReItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_realtor_realtor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReItemRequest) ProtoMessage() {}

func (x *GetReItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_realtor_realtor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReItemRequest.ProtoReflect.Descriptor instead.
func (*GetReItemRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_realtor_realtor_proto_rawDescGZIP(), []int{2}
}

func (x *GetReItemRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListReItemHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query       *common.ListParameters `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PrevId      int64                  `protobuf:"varint,2,opt,name=prev_id,json=prevId,proto3" json:"prev_id,omitempty"`
	ReType      int32                  `protobuf:"varint,3,opt,name=re_type,json=reType,proto3" json:"re_type,omitempty"`
	TxType      int32                  `protobuf:"varint,4,opt,name=txType,proto3" json:"txType,omitempty"`
	CityId      string                 `protobuf:"bytes,5,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Recommanded bool                   `protobuf:"varint,6,opt,name=recommanded,proto3" json:"recommanded,omitempty"`
}

func (x *ListReItemHeaderRequest) Reset() {
	*x = ListReItemHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_realtor_realtor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReItemHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReItemHeaderRequest) ProtoMessage() {}

func (x *ListReItemHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_realtor_realtor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReItemHeaderRequest.ProtoReflect.Descriptor instead.
func (*ListReItemHeaderRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_realtor_realtor_proto_rawDescGZIP(), []int{3}
}

func (x *ListReItemHeaderRequest) GetQuery() *common.ListParameters {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ListReItemHeaderRequest) GetPrevId() int64 {
	if x != nil {
		return x.PrevId
	}
	return 0
}

func (x *ListReItemHeaderRequest) GetReType() int32 {
	if x != nil {
		return x.ReType
	}
	return 0
}

func (x *ListReItemHeaderRequest) GetTxType() int32 {
	if x != nil {
		return x.TxType
	}
	return 0
}

func (x *ListReItemHeaderRequest) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *ListReItemHeaderRequest) GetRecommanded() bool {
	if x != nil {
		return x.Recommanded
	}
	return false
}

type ListReItemHeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ReItemHeader `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListReItemHeaderResponse) Reset() {
	*x = ListReItemHeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_realtor_realtor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReItemHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReItemHeaderResponse) ProtoMessage() {}

func (x *ListReItemHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_realtor_realtor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReItemHeaderResponse.ProtoReflect.Descriptor instead.
func (*ListReItemHeaderResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_realtor_realtor_proto_rawDescGZIP(), []int{4}
}

func (x *ListReItemHeaderResponse) GetItems() []*ReItemHeader {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteReItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReItemRequest) Reset() {
	*x = DeleteReItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_realtor_realtor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReItemRequest) ProtoMessage() {}

func (x *DeleteReItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_realtor_realtor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteReItemRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_realtor_realtor_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReItemRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_protos_v1_realtor_realtor_proto protoreflect.FileDescriptor

var file_protos_v1_realtor_realtor_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x6c,
	0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x04, 0x0a, 0x06, 0x52, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x64, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x65, 0x64, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x62, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x75, 0x72, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x66, 0x75, 0x72, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x39, 0x0a, 0x0c, 0x67, 0x65, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x67,
	0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xc4, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x64, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x65, 0x64, 0x72, 0x6f, 0x6f,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x39,
	0x0a, 0x0c, 0x67, 0x65, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x67, 0x65,
	0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x70, 0x72, 0x65, 0x76, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe7, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x61,
	0x6c, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65,
	0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x5d, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23,
	0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x12, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x36, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x74,
	0x6f, 0x72, 0x2e, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x47, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x72,
	0x65, 0x61, 0x6c, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x79, 0x2d, 0x79, 0x6f, 0x6f, 0x6e, 0x2f, 0x6b, 0x72, 0x65, 0x61, 0x6c, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x61,
	0x6c, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_realtor_realtor_proto_rawDescOnce sync.Once
	file_protos_v1_realtor_realtor_proto_rawDescData = file_protos_v1_realtor_realtor_proto_rawDesc
)

func file_protos_v1_realtor_realtor_proto_rawDescGZIP() []byte {
	file_protos_v1_realtor_realtor_proto_rawDescOnce.Do(func() {
		file_protos_v1_realtor_realtor_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_realtor_realtor_proto_rawDescData)
	})
	return file_protos_v1_realtor_realtor_proto_rawDescData
}

var file_protos_v1_realtor_realtor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_v1_realtor_realtor_proto_goTypes = []interface{}{
	(*ReItem)(nil),                   // 0: v1.realtor.ReItem
	(*ReItemHeader)(nil),             // 1: v1.realtor.ReItemHeader
	(*GetReItemRequest)(nil),         // 2: v1.realtor.GetReItemRequest
	(*ListReItemHeaderRequest)(nil),  // 3: v1.realtor.ListReItemHeaderRequest
	(*ListReItemHeaderResponse)(nil), // 4: v1.realtor.ListReItemHeaderResponse
	(*DeleteReItemRequest)(nil),      // 5: v1.realtor.DeleteReItemRequest
	(*common.GeoLocation)(nil),       // 6: v1.common.GeoLocation
	(*common.ListParameters)(nil),    // 7: v1.common.ListParameters
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_protos_v1_realtor_realtor_proto_depIdxs = []int32{
	6, // 0: v1.realtor.ReItem.geo_location:type_name -> v1.common.GeoLocation
	6, // 1: v1.realtor.ReItemHeader.geo_location:type_name -> v1.common.GeoLocation
	7, // 2: v1.realtor.ListReItemHeaderRequest.query:type_name -> v1.common.ListParameters
	1, // 3: v1.realtor.ListReItemHeaderResponse.items:type_name -> v1.realtor.ReItemHeader
	2, // 4: v1.realtor.RealtorService.GetReItem:input_type -> v1.realtor.GetReItemRequest
	3, // 5: v1.realtor.RealtorService.ListReItemHeader:input_type -> v1.realtor.ListReItemHeaderRequest
	0, // 6: v1.realtor.RealtorService.CreateReItem:input_type -> v1.realtor.ReItem
	0, // 7: v1.realtor.RealtorService.UpdateReItem:input_type -> v1.realtor.ReItem
	5, // 8: v1.realtor.RealtorService.DeleteReItem:input_type -> v1.realtor.DeleteReItemRequest
	0, // 9: v1.realtor.RealtorService.GetReItem:output_type -> v1.realtor.ReItem
	4, // 10: v1.realtor.RealtorService.ListReItemHeader:output_type -> v1.realtor.ListReItemHeaderResponse
	0, // 11: v1.realtor.RealtorService.CreateReItem:output_type -> v1.realtor.ReItem
	0, // 12: v1.realtor.RealtorService.UpdateReItem:output_type -> v1.realtor.ReItem
	8, // 13: v1.realtor.RealtorService.DeleteReItem:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_v1_realtor_realtor_proto_init() }
func file_protos_v1_realtor_realtor_proto_init() {
	if File_protos_v1_realtor_realtor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_realtor_realtor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReItem); i {
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
		file_protos_v1_realtor_realtor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReItemHeader); i {
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
		file_protos_v1_realtor_realtor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReItemRequest); i {
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
		file_protos_v1_realtor_realtor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReItemHeaderRequest); i {
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
		file_protos_v1_realtor_realtor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReItemHeaderResponse); i {
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
		file_protos_v1_realtor_realtor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReItemRequest); i {
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
			RawDescriptor: file_protos_v1_realtor_realtor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_realtor_realtor_proto_goTypes,
		DependencyIndexes: file_protos_v1_realtor_realtor_proto_depIdxs,
		MessageInfos:      file_protos_v1_realtor_realtor_proto_msgTypes,
	}.Build()
	File_protos_v1_realtor_realtor_proto = out.File
	file_protos_v1_realtor_realtor_proto_rawDesc = nil
	file_protos_v1_realtor_realtor_proto_goTypes = nil
	file_protos_v1_realtor_realtor_proto_depIdxs = nil
}
