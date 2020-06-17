// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/api/v2/core/config_source.proto

package envoy_api_v2_core

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ApiVersion int32

const (
	ApiVersion_AUTO ApiVersion = 0
	ApiVersion_V2   ApiVersion = 1
	ApiVersion_V3   ApiVersion = 2
)

// Enum value maps for ApiVersion.
var (
	ApiVersion_name = map[int32]string{
		0: "AUTO",
		1: "V2",
		2: "V3",
	}
	ApiVersion_value = map[string]int32{
		"AUTO": 0,
		"V2":   1,
		"V3":   2,
	}
)

func (x ApiVersion) Enum() *ApiVersion {
	p := new(ApiVersion)
	*p = x
	return p
}

func (x ApiVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_api_v2_core_config_source_proto_enumTypes[0].Descriptor()
}

func (ApiVersion) Type() protoreflect.EnumType {
	return &file_envoy_api_v2_core_config_source_proto_enumTypes[0]
}

func (x ApiVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiVersion.Descriptor instead.
func (ApiVersion) EnumDescriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{0}
}

type ApiConfigSource_ApiType int32

const (
	// Deprecated: Do not use.
	ApiConfigSource_UNSUPPORTED_REST_LEGACY ApiConfigSource_ApiType = 0
	ApiConfigSource_REST                    ApiConfigSource_ApiType = 1
	ApiConfigSource_GRPC                    ApiConfigSource_ApiType = 2
	ApiConfigSource_DELTA_GRPC              ApiConfigSource_ApiType = 3
)

// Enum value maps for ApiConfigSource_ApiType.
var (
	ApiConfigSource_ApiType_name = map[int32]string{
		0: "UNSUPPORTED_REST_LEGACY",
		1: "REST",
		2: "GRPC",
		3: "DELTA_GRPC",
	}
	ApiConfigSource_ApiType_value = map[string]int32{
		"UNSUPPORTED_REST_LEGACY": 0,
		"REST":                    1,
		"GRPC":                    2,
		"DELTA_GRPC":              3,
	}
)

func (x ApiConfigSource_ApiType) Enum() *ApiConfigSource_ApiType {
	p := new(ApiConfigSource_ApiType)
	*p = x
	return p
}

func (x ApiConfigSource_ApiType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiConfigSource_ApiType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_api_v2_core_config_source_proto_enumTypes[1].Descriptor()
}

func (ApiConfigSource_ApiType) Type() protoreflect.EnumType {
	return &file_envoy_api_v2_core_config_source_proto_enumTypes[1]
}

func (x ApiConfigSource_ApiType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiConfigSource_ApiType.Descriptor instead.
func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{0, 0}
}

type ApiConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiType                   ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	TransportApiVersion       ApiVersion              `protobuf:"varint,8,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.api.v2.core.ApiVersion" json:"transport_api_version,omitempty"`
	ClusterNames              []string                `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	GrpcServices              []*GrpcService          `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	RefreshDelay              *duration.Duration      `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	RequestTimeout            *duration.Duration      `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	RateLimitSettings         *RateLimitSettings      `protobuf:"bytes,6,opt,name=rate_limit_settings,json=rateLimitSettings,proto3" json:"rate_limit_settings,omitempty"`
	SetNodeOnFirstMessageOnly bool                    `protobuf:"varint,7,opt,name=set_node_on_first_message_only,json=setNodeOnFirstMessageOnly,proto3" json:"set_node_on_first_message_only,omitempty"`
}

func (x *ApiConfigSource) Reset() {
	*x = ApiConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiConfigSource) ProtoMessage() {}

func (x *ApiConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiConfigSource.ProtoReflect.Descriptor instead.
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{0}
}

func (x *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if x != nil {
		return x.ApiType
	}
	return ApiConfigSource_UNSUPPORTED_REST_LEGACY
}

func (x *ApiConfigSource) GetTransportApiVersion() ApiVersion {
	if x != nil {
		return x.TransportApiVersion
	}
	return ApiVersion_AUTO
}

func (x *ApiConfigSource) GetClusterNames() []string {
	if x != nil {
		return x.ClusterNames
	}
	return nil
}

func (x *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if x != nil {
		return x.GrpcServices
	}
	return nil
}

func (x *ApiConfigSource) GetRefreshDelay() *duration.Duration {
	if x != nil {
		return x.RefreshDelay
	}
	return nil
}

func (x *ApiConfigSource) GetRequestTimeout() *duration.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

func (x *ApiConfigSource) GetRateLimitSettings() *RateLimitSettings {
	if x != nil {
		return x.RateLimitSettings
	}
	return nil
}

func (x *ApiConfigSource) GetSetNodeOnFirstMessageOnly() bool {
	if x != nil {
		return x.SetNodeOnFirstMessageOnly
	}
	return false
}

type AggregatedConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AggregatedConfigSource) Reset() {
	*x = AggregatedConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedConfigSource) ProtoMessage() {}

func (x *AggregatedConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedConfigSource.ProtoReflect.Descriptor instead.
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{1}
}

type SelfConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SelfConfigSource) Reset() {
	*x = SelfConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfConfigSource) ProtoMessage() {}

func (x *SelfConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfConfigSource.ProtoReflect.Descriptor instead.
func (*SelfConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{2}
}

type RateLimitSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxTokens *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	FillRate  *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=fill_rate,json=fillRate,proto3" json:"fill_rate,omitempty"`
}

func (x *RateLimitSettings) Reset() {
	*x = RateLimitSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitSettings) ProtoMessage() {}

func (x *RateLimitSettings) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitSettings.ProtoReflect.Descriptor instead.
func (*RateLimitSettings) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitSettings) GetMaxTokens() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxTokens
	}
	return nil
}

func (x *RateLimitSettings) GetFillRate() *wrappers.DoubleValue {
	if x != nil {
		return x.FillRate
	}
	return nil
}

type ConfigSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	//	*ConfigSource_Self
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	InitialFetchTimeout   *duration.Duration                   `protobuf:"bytes,4,opt,name=initial_fetch_timeout,json=initialFetchTimeout,proto3" json:"initial_fetch_timeout,omitempty"`
	ResourceApiVersion    ApiVersion                           `protobuf:"varint,6,opt,name=resource_api_version,json=resourceApiVersion,proto3,enum=envoy.api.v2.core.ApiVersion" json:"resource_api_version,omitempty"`
}

func (x *ConfigSource) Reset() {
	*x = ConfigSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSource) ProtoMessage() {}

func (x *ConfigSource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_core_config_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSource.ProtoReflect.Descriptor instead.
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_core_config_source_proto_rawDescGZIP(), []int{4}
}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (x *ConfigSource) GetPath() string {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (x *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (x *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

func (x *ConfigSource) GetSelf() *SelfConfigSource {
	if x, ok := x.GetConfigSourceSpecifier().(*ConfigSource_Self); ok {
		return x.Self
	}
	return nil
}

func (x *ConfigSource) GetInitialFetchTimeout() *duration.Duration {
	if x != nil {
		return x.InitialFetchTimeout
	}
	return nil
}

func (x *ConfigSource) GetResourceApiVersion() ApiVersion {
	if x != nil {
		return x.ResourceApiVersion
	}
	return ApiVersion_AUTO
}

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

type ConfigSource_Self struct {
	Self *SelfConfigSource `protobuf:"bytes,5,opt,name=self,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Self) isConfigSource_ConfigSourceSpecifier() {}

var File_envoy_api_v2_core_config_source_proto protoreflect.FileDescriptor

var file_envoy_api_v2_core_config_source_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x05,
	0x0a, 0x0f, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x4f, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x70, 0x69, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x5b, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x54, 0x0a, 0x13, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x11, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x41, 0x0a,
	0x1e, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x73, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x6e,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x54, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x17, 0x55,
	0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x5f,
	0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x00, 0x1a, 0x08, 0x08, 0x01, 0xa8, 0xf7, 0xb4, 0x8b,
	0x02, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x47, 0x52, 0x50, 0x43, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x4c, 0x54, 0x41, 0x5f,
	0x47, 0x52, 0x50, 0x43, 0x10, 0x03, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x6c, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x12, 0x09, 0x21,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x6c, 0x52, 0x61,
	0x74, 0x65, 0x22, 0xba, 0x03, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x50, 0x0a, 0x11, 0x61, 0x70, 0x69,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x70, 0x69, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x61,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x03, 0x61, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x65,
	0x6c, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x6c,
	0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x4d, 0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x59, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x1e, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x2a,
	0x26, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a,
	0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x56, 0x33, 0x10, 0x02, 0x42, 0x5a, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x11, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2,
	0x98, 0xfe, 0x8f, 0x05, 0x16, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_api_v2_core_config_source_proto_rawDescOnce sync.Once
	file_envoy_api_v2_core_config_source_proto_rawDescData = file_envoy_api_v2_core_config_source_proto_rawDesc
)

func file_envoy_api_v2_core_config_source_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_core_config_source_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_core_config_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_core_config_source_proto_rawDescData)
	})
	return file_envoy_api_v2_core_config_source_proto_rawDescData
}

var file_envoy_api_v2_core_config_source_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_api_v2_core_config_source_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_api_v2_core_config_source_proto_goTypes = []interface{}{
	(ApiVersion)(0),                // 0: envoy.api.v2.core.ApiVersion
	(ApiConfigSource_ApiType)(0),   // 1: envoy.api.v2.core.ApiConfigSource.ApiType
	(*ApiConfigSource)(nil),        // 2: envoy.api.v2.core.ApiConfigSource
	(*AggregatedConfigSource)(nil), // 3: envoy.api.v2.core.AggregatedConfigSource
	(*SelfConfigSource)(nil),       // 4: envoy.api.v2.core.SelfConfigSource
	(*RateLimitSettings)(nil),      // 5: envoy.api.v2.core.RateLimitSettings
	(*ConfigSource)(nil),           // 6: envoy.api.v2.core.ConfigSource
	(*GrpcService)(nil),            // 7: envoy.api.v2.core.GrpcService
	(*duration.Duration)(nil),      // 8: google.protobuf.Duration
	(*wrappers.UInt32Value)(nil),   // 9: google.protobuf.UInt32Value
	(*wrappers.DoubleValue)(nil),   // 10: google.protobuf.DoubleValue
}
var file_envoy_api_v2_core_config_source_proto_depIdxs = []int32{
	1,  // 0: envoy.api.v2.core.ApiConfigSource.api_type:type_name -> envoy.api.v2.core.ApiConfigSource.ApiType
	0,  // 1: envoy.api.v2.core.ApiConfigSource.transport_api_version:type_name -> envoy.api.v2.core.ApiVersion
	7,  // 2: envoy.api.v2.core.ApiConfigSource.grpc_services:type_name -> envoy.api.v2.core.GrpcService
	8,  // 3: envoy.api.v2.core.ApiConfigSource.refresh_delay:type_name -> google.protobuf.Duration
	8,  // 4: envoy.api.v2.core.ApiConfigSource.request_timeout:type_name -> google.protobuf.Duration
	5,  // 5: envoy.api.v2.core.ApiConfigSource.rate_limit_settings:type_name -> envoy.api.v2.core.RateLimitSettings
	9,  // 6: envoy.api.v2.core.RateLimitSettings.max_tokens:type_name -> google.protobuf.UInt32Value
	10, // 7: envoy.api.v2.core.RateLimitSettings.fill_rate:type_name -> google.protobuf.DoubleValue
	2,  // 8: envoy.api.v2.core.ConfigSource.api_config_source:type_name -> envoy.api.v2.core.ApiConfigSource
	3,  // 9: envoy.api.v2.core.ConfigSource.ads:type_name -> envoy.api.v2.core.AggregatedConfigSource
	4,  // 10: envoy.api.v2.core.ConfigSource.self:type_name -> envoy.api.v2.core.SelfConfigSource
	8,  // 11: envoy.api.v2.core.ConfigSource.initial_fetch_timeout:type_name -> google.protobuf.Duration
	0,  // 12: envoy.api.v2.core.ConfigSource.resource_api_version:type_name -> envoy.api.v2.core.ApiVersion
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_core_config_source_proto_init() }
func file_envoy_api_v2_core_config_source_proto_init() {
	if File_envoy_api_v2_core_config_source_proto != nil {
		return
	}
	file_envoy_api_v2_core_grpc_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_core_config_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiConfigSource); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregatedConfigSource); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfConfigSource); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitSettings); i {
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
		file_envoy_api_v2_core_config_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSource); i {
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
	file_envoy_api_v2_core_config_source_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
		(*ConfigSource_Self)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_api_v2_core_config_source_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_core_config_source_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_core_config_source_proto_depIdxs,
		EnumInfos:         file_envoy_api_v2_core_config_source_proto_enumTypes,
		MessageInfos:      file_envoy_api_v2_core_config_source_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_core_config_source_proto = out.File
	file_envoy_api_v2_core_config_source_proto_rawDesc = nil
	file_envoy_api_v2_core_config_source_proto_goTypes = nil
	file_envoy_api_v2_core_config_source_proto_depIdxs = nil
}
