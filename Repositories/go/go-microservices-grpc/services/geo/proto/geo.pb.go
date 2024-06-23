// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: internal/services/geo/proto/geo.proto

package proto

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

// The latitude and longitude of the current location.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float32 `protobuf:"fixed32,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_geo_proto_geo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_geo_proto_geo_proto_msgTypes[0]
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
	return file_internal_services_geo_proto_geo_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Request) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_geo_proto_geo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_geo_proto_geo_proto_msgTypes[1]
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
	return file_internal_services_geo_proto_geo_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetHotelIds() []string {
	if x != nil {
		return x.HotelIds
	}
	return nil
}

var File_internal_services_geo_proto_geo_proto protoreflect.FileDescriptor

var file_internal_services_geo_proto_geo_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x65, 0x6f, 0x22, 0x2d, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64,
	0x73, 0x32, 0x2a, 0x0a, 0x03, 0x47, 0x65, 0x6f, 0x12, 0x23, 0x0a, 0x06, 0x4e, 0x65, 0x61, 0x72,
	0x62, 0x79, 0x12, 0x0c, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x1f, 0x5a,
	0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_services_geo_proto_geo_proto_rawDescOnce sync.Once
	file_internal_services_geo_proto_geo_proto_rawDescData = file_internal_services_geo_proto_geo_proto_rawDesc
)

func file_internal_services_geo_proto_geo_proto_rawDescGZIP() []byte {
	file_internal_services_geo_proto_geo_proto_rawDescOnce.Do(func() {
		file_internal_services_geo_proto_geo_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_services_geo_proto_geo_proto_rawDescData)
	})
	return file_internal_services_geo_proto_geo_proto_rawDescData
}

var file_internal_services_geo_proto_geo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_services_geo_proto_geo_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: geo.Request
	(*Result)(nil),  // 1: geo.Result
}
var file_internal_services_geo_proto_geo_proto_depIdxs = []int32{
	0, // 0: geo.Geo.Nearby:input_type -> geo.Request
	1, // 1: geo.Geo.Nearby:output_type -> geo.Result
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_services_geo_proto_geo_proto_init() }
func file_internal_services_geo_proto_geo_proto_init() {
	if File_internal_services_geo_proto_geo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_services_geo_proto_geo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_services_geo_proto_geo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_services_geo_proto_geo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_services_geo_proto_geo_proto_goTypes,
		DependencyIndexes: file_internal_services_geo_proto_geo_proto_depIdxs,
		MessageInfos:      file_internal_services_geo_proto_geo_proto_msgTypes,
	}.Build()
	File_internal_services_geo_proto_geo_proto = out.File
	file_internal_services_geo_proto_geo_proto_rawDesc = nil
	file_internal_services_geo_proto_geo_proto_goTypes = nil
	file_internal_services_geo_proto_geo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GeoClient is the client API for Geo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoClient interface {
	// Finds the hotels contained nearby the current lat/lon.
	Nearby(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type geoClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoClient(cc grpc.ClientConnInterface) GeoClient {
	return &geoClient{cc}
}

func (c *geoClient) Nearby(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/geo.Geo/Nearby", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoServer is the server API for Geo service.
type GeoServer interface {
	// Finds the hotels contained nearby the current lat/lon.
	Nearby(context.Context, *Request) (*Result, error)
}

// UnimplementedGeoServer can be embedded to have forward compatible implementations.
type UnimplementedGeoServer struct {
}

func (*UnimplementedGeoServer) Nearby(context.Context, *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nearby not implemented")
}

func RegisterGeoServer(s *grpc.Server, srv GeoServer) {
	s.RegisterService(&_Geo_serviceDesc, srv)
}

func _Geo_Nearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServer).Nearby(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/geo.Geo/Nearby",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServer).Nearby(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Geo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geo.Geo",
	HandlerType: (*GeoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nearby",
			Handler:    _Geo_Nearby_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/services/geo/proto/geo.proto",
}
