// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: internal/services/rate/proto/rate.proto

package rate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	InDate   string   `protobuf:"bytes,2,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate  string   `protobuf:"bytes,3,opt,name=outDate,proto3" json:"outDate,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_rate_proto_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_rate_proto_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_internal_services_rate_proto_rate_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetHotelIds() []string {
	if x != nil {
		return x.HotelIds
	}
	return nil
}

func (x *Request) GetInDate() string {
	if x != nil {
		return x.InDate
	}
	return ""
}

func (x *Request) GetOutDate() string {
	if x != nil {
		return x.OutDate
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatePlans []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans,proto3" json:"ratePlans,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_rate_proto_rate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_rate_proto_rate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_internal_services_rate_proto_rate_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetRatePlans() []*RatePlan {
	if x != nil {
		return x.RatePlans
	}
	return nil
}

type RatePlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId  string    `protobuf:"bytes,1,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	Code     string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	InDate   string    `protobuf:"bytes,3,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate  string    `protobuf:"bytes,4,opt,name=outDate,proto3" json:"outDate,omitempty"`
	RoomType *RoomType `protobuf:"bytes,5,opt,name=roomType,proto3" json:"roomType,omitempty"`
}

func (x *RatePlan) Reset() {
	*x = RatePlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_rate_proto_rate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatePlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatePlan) ProtoMessage() {}

func (x *RatePlan) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_rate_proto_rate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatePlan.ProtoReflect.Descriptor instead.
func (*RatePlan) Descriptor() ([]byte, []int) {
	return file_internal_services_rate_proto_rate_proto_rawDescGZIP(), []int{2}
}

func (x *RatePlan) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *RatePlan) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RatePlan) GetInDate() string {
	if x != nil {
		return x.InDate
	}
	return ""
}

func (x *RatePlan) GetOutDate() string {
	if x != nil {
		return x.OutDate
	}
	return ""
}

func (x *RatePlan) GetRoomType() *RoomType {
	if x != nil {
		return x.RoomType
	}
	return nil
}

type RoomType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookableRate       float64 `protobuf:"fixed64,1,opt,name=bookableRate,proto3" json:"bookableRate,omitempty"`
	TotalRate          float64 `protobuf:"fixed64,2,opt,name=totalRate,proto3" json:"totalRate,omitempty"`
	TotalRateInclusive float64 `protobuf:"fixed64,3,opt,name=totalRateInclusive,proto3" json:"totalRateInclusive,omitempty"`
	Code               string  `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	RoomDescription    string  `protobuf:"bytes,6,opt,name=roomDescription,proto3" json:"roomDescription,omitempty"`
}

func (x *RoomType) Reset() {
	*x = RoomType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_rate_proto_rate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomType) ProtoMessage() {}

func (x *RoomType) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_rate_proto_rate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomType.ProtoReflect.Descriptor instead.
func (*RoomType) Descriptor() ([]byte, []int) {
	return file_internal_services_rate_proto_rate_proto_rawDescGZIP(), []int{3}
}

func (x *RoomType) GetBookableRate() float64 {
	if x != nil {
		return x.BookableRate
	}
	return 0
}

func (x *RoomType) GetTotalRate() float64 {
	if x != nil {
		return x.TotalRate
	}
	return 0
}

func (x *RoomType) GetTotalRateInclusive() float64 {
	if x != nil {
		return x.TotalRateInclusive
	}
	return 0
}

func (x *RoomType) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RoomType) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *RoomType) GetRoomDescription() string {
	if x != nil {
		return x.RoomDescription
	}
	return ""
}

var File_internal_services_rate_proto_rate_proto protoreflect.FileDescriptor

var file_internal_services_rate_proto_rate_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22,
	0x57, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x73,
	0x22, 0x96, 0x01, 0x0a, 0x08, 0x52, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x08, 0x52, 0x6f,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x62, 0x6f,
	0x6f, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x6f, 0x6f, 0x6d,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x2f, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_services_rate_proto_rate_proto_rawDescOnce sync.Once
	file_internal_services_rate_proto_rate_proto_rawDescData = file_internal_services_rate_proto_rate_proto_rawDesc
)

func file_internal_services_rate_proto_rate_proto_rawDescGZIP() []byte {
	file_internal_services_rate_proto_rate_proto_rawDescOnce.Do(func() {
		file_internal_services_rate_proto_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_services_rate_proto_rate_proto_rawDescData)
	})
	return file_internal_services_rate_proto_rate_proto_rawDescData
}

var file_internal_services_rate_proto_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_services_rate_proto_rate_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: rate.Request
	(*Result)(nil),   // 1: rate.Result
	(*RatePlan)(nil), // 2: rate.RatePlan
	(*RoomType)(nil), // 3: rate.RoomType
}
var file_internal_services_rate_proto_rate_proto_depIdxs = []int32{
	2, // 0: rate.Result.ratePlans:type_name -> rate.RatePlan
	3, // 1: rate.RatePlan.roomType:type_name -> rate.RoomType
	0, // 2: rate.Rate.GetRates:input_type -> rate.Request
	1, // 3: rate.Rate.GetRates:output_type -> rate.Result
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_services_rate_proto_rate_proto_init() }
func file_internal_services_rate_proto_rate_proto_init() {
	if File_internal_services_rate_proto_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_services_rate_proto_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_internal_services_rate_proto_rate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_internal_services_rate_proto_rate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatePlan); i {
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
		file_internal_services_rate_proto_rate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomType); i {
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
			RawDescriptor: file_internal_services_rate_proto_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_services_rate_proto_rate_proto_goTypes,
		DependencyIndexes: file_internal_services_rate_proto_rate_proto_depIdxs,
		MessageInfos:      file_internal_services_rate_proto_rate_proto_msgTypes,
	}.Build()
	File_internal_services_rate_proto_rate_proto = out.File
	file_internal_services_rate_proto_rate_proto_rawDesc = nil
	file_internal_services_rate_proto_rate_proto_goTypes = nil
	file_internal_services_rate_proto_rate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RateClient is the client API for Rate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateClient interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type rateClient struct {
	cc grpc.ClientConnInterface
}

func NewRateClient(cc grpc.ClientConnInterface) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/rate.Rate/GetRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServer is the server API for Rate service.
type RateServer interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request) (*Result, error)
}

// UnimplementedRateServer can be embedded to have forward compatible implementations.
type UnimplementedRateServer struct {
}

func (*UnimplementedRateServer) GetRates(context.Context, *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRates not implemented")
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.Rate/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).GetRates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/services/rate/proto/rate.proto",
}
