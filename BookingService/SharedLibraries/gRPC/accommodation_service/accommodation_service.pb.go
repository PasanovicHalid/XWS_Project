// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: accommodation_service/accommodation_service.proto

package accommodation

import (
	common "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type NewAccomodation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location          string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Wifi              bool     `protobuf:"varint,4,opt,name=wifi,proto3" json:"wifi,omitempty"`
	Kitchen           bool     `protobuf:"varint,5,opt,name=kitchen,proto3" json:"kitchen,omitempty"`
	AirConditioner    bool     `protobuf:"varint,6,opt,name=air_conditioner,json=airConditioner,proto3" json:"air_conditioner,omitempty"`
	Parking           bool     `protobuf:"varint,7,opt,name=parking,proto3" json:"parking,omitempty"`
	MinNumberOfGuests int32    `protobuf:"varint,8,opt,name=min_number_of_guests,json=minNumberOfGuests,proto3" json:"min_number_of_guests,omitempty"`
	MaxNumberOfGuests int32    `protobuf:"varint,9,opt,name=max_number_of_guests,json=maxNumberOfGuests,proto3" json:"max_number_of_guests,omitempty"`
	Images            []string `protobuf:"bytes,10,rep,name=images,proto3" json:"images,omitempty"`
	OwnerId           string   `protobuf:"bytes,11,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
}

func (x *NewAccomodation) Reset() {
	*x = NewAccomodation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAccomodation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAccomodation) ProtoMessage() {}

func (x *NewAccomodation) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAccomodation.ProtoReflect.Descriptor instead.
func (*NewAccomodation) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewAccomodation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewAccomodation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewAccomodation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *NewAccomodation) GetWifi() bool {
	if x != nil {
		return x.Wifi
	}
	return false
}

func (x *NewAccomodation) GetKitchen() bool {
	if x != nil {
		return x.Kitchen
	}
	return false
}

func (x *NewAccomodation) GetAirConditioner() bool {
	if x != nil {
		return x.AirConditioner
	}
	return false
}

func (x *NewAccomodation) GetParking() bool {
	if x != nil {
		return x.Parking
	}
	return false
}

func (x *NewAccomodation) GetMinNumberOfGuests() int32 {
	if x != nil {
		return x.MinNumberOfGuests
	}
	return 0
}

func (x *NewAccomodation) GetMaxNumberOfGuests() int32 {
	if x != nil {
		return x.MaxNumberOfGuests
	}
	return 0
}

func (x *NewAccomodation) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *NewAccomodation) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetFilteredAccommodationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilteredAccommodations []*NewAccomodation `protobuf:"bytes,1,rep,name=filtered_accommodations,json=filteredAccommodations,proto3" json:"filtered_accommodations,omitempty"`
}

func (x *GetFilteredAccommodationsResponse) Reset() {
	*x = GetFilteredAccommodationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilteredAccommodationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilteredAccommodationsResponse) ProtoMessage() {}

func (x *GetFilteredAccommodationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilteredAccommodationsResponse.ProtoReflect.Descriptor instead.
func (*GetFilteredAccommodationsResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFilteredAccommodationsResponse) GetFilteredAccommodations() []*NewAccomodation {
	if x != nil {
		return x.FilteredAccommodations
	}
	return nil
}

type CreateOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId  string                 `protobuf:"bytes,1,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	StartDateTimeUtc *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date_time_utc,json=startDateTimeUtc,proto3" json:"start_date_time_utc,omitempty"`
	EndDateTimeUtc   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_date_time_utc,json=endDateTimeUtc,proto3" json:"end_date_time_utc,omitempty"`
	Price            int32                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	PerGuest         bool                   `protobuf:"varint,5,opt,name=per_guest,json=perGuest,proto3" json:"per_guest,omitempty"`
}

func (x *CreateOfferRequest) Reset() {
	*x = CreateOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfferRequest) ProtoMessage() {}

func (x *CreateOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfferRequest.ProtoReflect.Descriptor instead.
func (*CreateOfferRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOfferRequest) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *CreateOfferRequest) GetStartDateTimeUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDateTimeUtc
	}
	return nil
}

func (x *CreateOfferRequest) GetEndDateTimeUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDateTimeUtc
	}
	return nil
}

func (x *CreateOfferRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOfferRequest) GetPerGuest() bool {
	if x != nil {
		return x.PerGuest
	}
	return false
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{3}
}

type AccommodationOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccommodationId  string                 `protobuf:"bytes,2,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	StartDateTimeUtc *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date_time_utc,json=startDateTimeUtc,proto3" json:"start_date_time_utc,omitempty"`
	EndDateTimeUtc   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_date_time_utc,json=endDateTimeUtc,proto3" json:"end_date_time_utc,omitempty"`
	Price            int32                  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	PerGuest         bool                   `protobuf:"varint,6,opt,name=per_guest,json=perGuest,proto3" json:"per_guest,omitempty"`
}

func (x *AccommodationOffer) Reset() {
	*x = AccommodationOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationOffer) ProtoMessage() {}

func (x *AccommodationOffer) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationOffer.ProtoReflect.Descriptor instead.
func (*AccommodationOffer) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{4}
}

func (x *AccommodationOffer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccommodationOffer) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *AccommodationOffer) GetStartDateTimeUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDateTimeUtc
	}
	return nil
}

func (x *AccommodationOffer) GetEndDateTimeUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDateTimeUtc
	}
	return nil
}

func (x *AccommodationOffer) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AccommodationOffer) GetPerGuest() bool {
	if x != nil {
		return x.PerGuest
	}
	return false
}

type AccommodationSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location         string                 `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	StartDateTimeUtc *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date_time_utc,json=startDateTimeUtc,proto3" json:"start_date_time_utc,omitempty"`
	EndDateTimeUtc   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_date_time_utc,json=endDateTimeUtc,proto3" json:"end_date_time_utc,omitempty"`
	GuestNumber      int32                  `protobuf:"varint,4,opt,name=guest_number,json=guestNumber,proto3" json:"guest_number,omitempty"`
}

func (x *AccommodationSearch) Reset() {
	*x = AccommodationSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationSearch) ProtoMessage() {}

func (x *AccommodationSearch) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationSearch.ProtoReflect.Descriptor instead.
func (*AccommodationSearch) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{5}
}

func (x *AccommodationSearch) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AccommodationSearch) GetStartDateTimeUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDateTimeUtc
	}
	return nil
}

func (x *AccommodationSearch) GetEndDateTimeUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDateTimeUtc
	}
	return nil
}

func (x *AccommodationSearch) GetGuestNumber() int32 {
	if x != nil {
		return x.GuestNumber
	}
	return 0
}

type GetOwnerIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOwnerIdRequest) Reset() {
	*x = GetOwnerIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerIdRequest) ProtoMessage() {}

func (x *GetOwnerIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerIdRequest.ProtoReflect.Descriptor instead.
func (*GetOwnerIdRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetOwnerIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOwnerIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOwnerIdResponse) Reset() {
	*x = GetOwnerIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerIdResponse) ProtoMessage() {}

func (x *GetOwnerIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerIdResponse.ProtoReflect.Descriptor instead.
func (*GetOwnerIdResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetOwnerIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_accommodation_service_accommodation_service_proto protoreflect.FileDescriptor

var file_accommodation_service_accommodation_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69,
	0x66, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x77, 0x69, 0x66, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x69, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x14, 0x6d,
	0x69, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x14,
	0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x7c, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x17, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x84, 0x02,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x49, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74, 0x63, 0x12, 0x45, 0x0a, 0x11, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74,
	0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x5f, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x65, 0x72, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74,
	0x63, 0x12, 0x45, 0x0a, 0x11, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x65, 0x72, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x70, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x13,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x49, 0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74, 0x63, 0x12, 0x45, 0x0a, 0x11, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x74, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x74,
	0x63, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xc6, 0x05, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e,
	0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x15,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x22, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01,
	0x0a, 0x14, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x30, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x33, 0x22, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8e, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x73, 0x61, 0x6e, 0x6f, 0x76, 0x69, 0x63,
	0x48, 0x61, 0x6c, 0x69, 0x64, 0x2f, 0x58, 0x57, 0x53, 0x5f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accommodation_service_accommodation_service_proto_rawDescOnce sync.Once
	file_accommodation_service_accommodation_service_proto_rawDescData = file_accommodation_service_accommodation_service_proto_rawDesc
)

func file_accommodation_service_accommodation_service_proto_rawDescGZIP() []byte {
	file_accommodation_service_accommodation_service_proto_rawDescOnce.Do(func() {
		file_accommodation_service_accommodation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accommodation_service_accommodation_service_proto_rawDescData)
	})
	return file_accommodation_service_accommodation_service_proto_rawDescData
}

var file_accommodation_service_accommodation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_accommodation_service_accommodation_service_proto_goTypes = []interface{}{
	(*NewAccomodation)(nil),                   // 0: accommodation.newAccomodation
	(*GetFilteredAccommodationsResponse)(nil), // 1: accommodation.GetFilteredAccommodationsResponse
	(*CreateOfferRequest)(nil),                // 2: accommodation.CreateOfferRequest
	(*EmptyMessage)(nil),                      // 3: accommodation.EmptyMessage
	(*AccommodationOffer)(nil),                // 4: accommodation.AccommodationOffer
	(*AccommodationSearch)(nil),               // 5: accommodation.AccommodationSearch
	(*GetOwnerIdRequest)(nil),                 // 6: accommodation.GetOwnerIdRequest
	(*GetOwnerIdResponse)(nil),                // 7: accommodation.GetOwnerIdResponse
	(*timestamppb.Timestamp)(nil),             // 8: google.protobuf.Timestamp
	(*common.RequestResult)(nil),              // 9: common.RequestResult
}
var file_accommodation_service_accommodation_service_proto_depIdxs = []int32{
	0,  // 0: accommodation.GetFilteredAccommodationsResponse.filtered_accommodations:type_name -> accommodation.newAccomodation
	8,  // 1: accommodation.CreateOfferRequest.start_date_time_utc:type_name -> google.protobuf.Timestamp
	8,  // 2: accommodation.CreateOfferRequest.end_date_time_utc:type_name -> google.protobuf.Timestamp
	8,  // 3: accommodation.AccommodationOffer.start_date_time_utc:type_name -> google.protobuf.Timestamp
	8,  // 4: accommodation.AccommodationOffer.end_date_time_utc:type_name -> google.protobuf.Timestamp
	8,  // 5: accommodation.AccommodationSearch.start_date_time_utc:type_name -> google.protobuf.Timestamp
	8,  // 6: accommodation.AccommodationSearch.end_date_time_utc:type_name -> google.protobuf.Timestamp
	0,  // 7: accommodation.AccommodationService.CreateAccomodation:input_type -> accommodation.newAccomodation
	2,  // 8: accommodation.AccommodationService.CreateAccomodationOffer:input_type -> accommodation.CreateOfferRequest
	4,  // 9: accommodation.AccommodationService.UpdateAccomodationOffer:input_type -> accommodation.AccommodationOffer
	5,  // 10: accommodation.AccommodationService.FilterAccommodations:input_type -> accommodation.AccommodationSearch
	6,  // 11: accommodation.AccommodationService.GetOwnerIdByAccommodationId:input_type -> accommodation.GetOwnerIdRequest
	9,  // 12: accommodation.AccommodationService.CreateAccomodation:output_type -> common.RequestResult
	9,  // 13: accommodation.AccommodationService.CreateAccomodationOffer:output_type -> common.RequestResult
	9,  // 14: accommodation.AccommodationService.UpdateAccomodationOffer:output_type -> common.RequestResult
	1,  // 15: accommodation.AccommodationService.FilterAccommodations:output_type -> accommodation.GetFilteredAccommodationsResponse
	7,  // 16: accommodation.AccommodationService.GetOwnerIdByAccommodationId:output_type -> accommodation.GetOwnerIdResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_accommodation_service_accommodation_service_proto_init() }
func file_accommodation_service_accommodation_service_proto_init() {
	if File_accommodation_service_accommodation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accommodation_service_accommodation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAccomodation); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilteredAccommodationsResponse); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOfferRequest); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationOffer); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationSearch); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerIdRequest); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerIdResponse); i {
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
			RawDescriptor: file_accommodation_service_accommodation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accommodation_service_accommodation_service_proto_goTypes,
		DependencyIndexes: file_accommodation_service_accommodation_service_proto_depIdxs,
		MessageInfos:      file_accommodation_service_accommodation_service_proto_msgTypes,
	}.Build()
	File_accommodation_service_accommodation_service_proto = out.File
	file_accommodation_service_accommodation_service_proto_rawDesc = nil
	file_accommodation_service_accommodation_service_proto_goTypes = nil
	file_accommodation_service_accommodation_service_proto_depIdxs = nil
}
