// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/service.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import api1 "google.golang.org/genproto/googleapis/api"
import annotations "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/api/label"
import metric "google.golang.org/genproto/googleapis/api/metric"
import monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
import api "google.golang.org/genproto/protobuf/api"
import ptype "google.golang.org/genproto/protobuf/ptype"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// `Service` is the root object of Google service configuration schema. It
// describes basic information about a service, such as the name and the
// title, and delegates other aspects to sub-sections. Each sub-section is
// either a proto message or a repeated proto message that configures a
// specific aspect, such as auth. See each proto message definition for details.
//
// Example:
//
//     type: google.api.Service
//     config_version: 3
//     name: calendar.googleapis.com
//     title: Google Calendar API
//     apis:
//     - name: google.calendar.v3.Calendar
//     authentication:
//       providers:
//       - id: google_calendar_auth
//         jwks_uri: https://www.googleapis.com/oauth2/v1/certs
//         issuer: https://securetoken.google.com
//       rules:
//       - selector: "*"
//         requirements:
//           provider_id: google_calendar_auth
type Service struct {
	// The semantic version of the service configuration. The config version
	// affects the interpretation of the service configuration. For example,
	// certain features are enabled by default for certain config versions.
	// The latest config version is `3`.
	ConfigVersion *wrappers.UInt32Value `protobuf:"bytes,20,opt,name=config_version,json=configVersion,proto3" json:"config_version,omitempty"`
	// The service name, which is a DNS-like logical identifier for the
	// service, such as `calendar.googleapis.com`. The service name
	// typically goes through DNS verification to make sure the owner
	// of the service also owns the DNS name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a specific instance of this message, typically assigned
	// by the client for tracking purpose. If empty, the server may choose to
	// generate one instead. Must be no longer than 60 characters.
	Id string `protobuf:"bytes,33,opt,name=id,proto3" json:"id,omitempty"`
	// The product title for this service.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// The Google project that owns this service.
	ProducerProjectId string `protobuf:"bytes,22,opt,name=producer_project_id,json=producerProjectId,proto3" json:"producer_project_id,omitempty"`
	// A list of API interfaces exported by this service. Only the `name` field
	// of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration
	// author, as the remaining fields will be derived from the IDL during the
	// normalization process. It is an error to specify an API interface here
	// which cannot be resolved against the associated IDL files.
	Apis []*api.Api `protobuf:"bytes,3,rep,name=apis,proto3" json:"apis,omitempty"`
	// A list of all proto message types included in this API service.
	// Types referenced directly or indirectly by the `apis` are
	// automatically included.  Messages which are not referenced but
	// shall be included, such as types used by the `google.protobuf.Any` type,
	// should be listed here by name. Example:
	//
	//     types:
	//     - name: google.protobuf.Int32
	Types []*ptype.Type `protobuf:"bytes,4,rep,name=types,proto3" json:"types,omitempty"`
	// A list of all enum types included in this API service.  Enums
	// referenced directly or indirectly by the `apis` are automatically
	// included.  Enums which are not referenced but shall be included
	// should be listed here by name. Example:
	//
	//     enums:
	//     - name: google.someapi.v1.SomeEnum
	Enums []*ptype.Enum `protobuf:"bytes,5,rep,name=enums,proto3" json:"enums,omitempty"`
	// Additional API documentation.
	Documentation *Documentation `protobuf:"bytes,6,opt,name=documentation,proto3" json:"documentation,omitempty"`
	// API backend configuration.
	Backend *Backend `protobuf:"bytes,8,opt,name=backend,proto3" json:"backend,omitempty"`
	// HTTP configuration.
	Http *annotations.Http `protobuf:"bytes,9,opt,name=http,proto3" json:"http,omitempty"`
	// Quota configuration.
	Quota *Quota `protobuf:"bytes,10,opt,name=quota,proto3" json:"quota,omitempty"`
	// Auth configuration.
	Authentication *Authentication `protobuf:"bytes,11,opt,name=authentication,proto3" json:"authentication,omitempty"`
	// Context configuration.
	Context *Context `protobuf:"bytes,12,opt,name=context,proto3" json:"context,omitempty"`
	// Configuration controlling usage of this service.
	Usage *Usage `protobuf:"bytes,15,opt,name=usage,proto3" json:"usage,omitempty"`
	// Configuration for network endpoints.  If this is empty, then an endpoint
	// with the same name as the service is automatically generated to service all
	// defined APIs.
	Endpoints []*Endpoint `protobuf:"bytes,18,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Configuration for the service control plane.
	Control *Control `protobuf:"bytes,21,opt,name=control,proto3" json:"control,omitempty"`
	// Defines the logs used by this service.
	Logs []*LogDescriptor `protobuf:"bytes,23,rep,name=logs,proto3" json:"logs,omitempty"`
	// Defines the metrics used by this service.
	Metrics []*metric.MetricDescriptor `protobuf:"bytes,24,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Defines the monitored resources used by this service. This is required
	// by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations.
	MonitoredResources []*monitoredres.MonitoredResourceDescriptor `protobuf:"bytes,25,rep,name=monitored_resources,json=monitoredResources,proto3" json:"monitored_resources,omitempty"`
	// Billing configuration.
	Billing *Billing `protobuf:"bytes,26,opt,name=billing,proto3" json:"billing,omitempty"`
	// Logging configuration.
	Logging *Logging `protobuf:"bytes,27,opt,name=logging,proto3" json:"logging,omitempty"`
	// Monitoring configuration.
	Monitoring *Monitoring `protobuf:"bytes,28,opt,name=monitoring,proto3" json:"monitoring,omitempty"`
	// System parameter configuration.
	SystemParameters *SystemParameters `protobuf:"bytes,29,opt,name=system_parameters,json=systemParameters,proto3" json:"system_parameters,omitempty"`
	// Output only. The source information for this configuration if available.
	SourceInfo *SourceInfo `protobuf:"bytes,37,opt,name=source_info,json=sourceInfo,proto3" json:"source_info,omitempty"`
	// Experimental configuration.
	Experimental         *api1.Experimental `protobuf:"bytes,101,opt,name=experimental,proto3" json:"experimental,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e73db125e27d5d7f, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetConfigVersion() *wrappers.UInt32Value {
	if m != nil {
		return m.ConfigVersion
	}
	return nil
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Service) GetProducerProjectId() string {
	if m != nil {
		return m.ProducerProjectId
	}
	return ""
}

func (m *Service) GetApis() []*api.Api {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Service) GetTypes() []*ptype.Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Service) GetEnums() []*ptype.Enum {
	if m != nil {
		return m.Enums
	}
	return nil
}

func (m *Service) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *Service) GetBackend() *Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Service) GetHttp() *annotations.Http {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Service) GetQuota() *Quota {
	if m != nil {
		return m.Quota
	}
	return nil
}

func (m *Service) GetAuthentication() *Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Service) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Service) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetControl() *Control {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *Service) GetLogs() []*LogDescriptor {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Service) GetMetrics() []*metric.MetricDescriptor {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Service) GetMonitoredResources() []*monitoredres.MonitoredResourceDescriptor {
	if m != nil {
		return m.MonitoredResources
	}
	return nil
}

func (m *Service) GetBilling() *Billing {
	if m != nil {
		return m.Billing
	}
	return nil
}

func (m *Service) GetLogging() *Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *Service) GetMonitoring() *Monitoring {
	if m != nil {
		return m.Monitoring
	}
	return nil
}

func (m *Service) GetSystemParameters() *SystemParameters {
	if m != nil {
		return m.SystemParameters
	}
	return nil
}

func (m *Service) GetSourceInfo() *SourceInfo {
	if m != nil {
		return m.SourceInfo
	}
	return nil
}

func (m *Service) GetExperimental() *api1.Experimental {
	if m != nil {
		return m.Experimental
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.api.Service")
}

func init() { proto.RegisterFile("google/api/service.proto", fileDescriptor_service_e73db125e27d5d7f) }

var fileDescriptor_service_e73db125e27d5d7f = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x96, 0xdf, 0x6e, 0xdb, 0x36,
	0x14, 0x87, 0x61, 0xd7, 0x6e, 0x16, 0x3a, 0xcd, 0x1a, 0xc6, 0x49, 0x19, 0xd7, 0x1b, 0xd2, 0xfd,
	0x41, 0x8d, 0x0d, 0x95, 0x01, 0x17, 0xe8, 0x2e, 0x36, 0x60, 0x88, 0xdb, 0x60, 0x33, 0xd0, 0x01,
	0x1e, 0xb3, 0x16, 0xc3, 0x6e, 0x0c, 0x5a, 0xa2, 0x55, 0x6e, 0x12, 0xc9, 0x91, 0x54, 0x17, 0x3f,
	0xc7, 0xde, 0x60, 0x4f, 0x3a, 0x88, 0xa4, 0x62, 0xca, 0x52, 0xee, 0x22, 0x7e, 0xdf, 0x39, 0x38,
	0x14, 0xa9, 0x9f, 0x03, 0x50, 0x2a, 0x44, 0x9a, 0xd1, 0x29, 0x91, 0x6c, 0xaa, 0xa9, 0xfa, 0xc8,
	0x62, 0x1a, 0x49, 0x25, 0x8c, 0x80, 0xc0, 0x91, 0x88, 0x48, 0x36, 0x1a, 0x07, 0x16, 0xe1, 0x5c,
	0x18, 0x62, 0x98, 0xe0, 0xda, 0x99, 0xa3, 0xb3, 0x90, 0x16, 0xe6, 0x83, 0x5f, 0x0e, 0x5b, 0xaf,
	0x49, 0xfc, 0x17, 0xe5, 0x49, 0x1b, 0x61, 0x59, 0xc6, 0x78, 0xda, 0x42, 0x62, 0xc1, 0x0d, 0xbd,
	0x35, 0xf7, 0x10, 0x25, 0x32, 0x4f, 0x3e, 0x0f, 0x48, 0x22, 0xe2, 0x22, 0xa7, 0xdc, 0xcd, 0xe7,
	0xf9, 0x45, 0xc0, 0x29, 0x4f, 0xa4, 0x60, 0xbc, 0x6a, 0xfa, 0x4d, 0x88, 0x6e, 0x25, 0x55, 0xcc,
	0x16, 0x67, 0xb5, 0x87, 0x96, 0x5d, 0x7e, 0x30, 0x46, 0xfa, 0xe5, 0xf3, 0x60, 0x39, 0x23, 0x6b,
	0x5a, 0xe9, 0xc3, 0x70, 0x5d, 0xb4, 0xed, 0x2f, 0x13, 0x69, 0xba, 0xdb, 0xf9, 0x93, 0x80, 0xe4,
	0xd4, 0x28, 0x16, 0x7b, 0xf0, 0x65, 0x08, 0x04, 0x67, 0x46, 0x28, 0x9a, 0xac, 0x14, 0xd5, 0xa2,
	0x50, 0xd5, 0x61, 0x8d, 0x9e, 0x36, 0xa5, 0x5d, 0xeb, 0x70, 0xc4, 0xbf, 0x0b, 0x61, 0x88, 0x5f,
	0x0f, 0x4f, 0xd5, 0x75, 0x5b, 0x31, 0xbe, 0x11, 0x9e, 0x3e, 0x0b, 0xe9, 0x56, 0x1b, 0x9a, 0xaf,
	0x24, 0x51, 0x24, 0xa7, 0x86, 0xaa, 0x96, 0xc6, 0x85, 0x26, 0x29, 0xdd, 0x7b, 0xe3, 0xf6, 0x69,
	0x5d, 0x6c, 0xa6, 0x84, 0x6f, 0xef, 0x45, 0x92, 0x79, 0x34, 0xda, 0x47, 0x66, 0x2b, 0xe9, 0xde,
	0x19, 0xdf, 0xb1, 0x7f, 0x14, 0x91, 0x92, 0x2a, 0x7f, 0x05, 0xbf, 0xf8, 0x17, 0x80, 0x83, 0x1b,
	0x77, 0x7d, 0xe1, 0x6b, 0x70, 0x1c, 0x0b, 0xbe, 0x61, 0xe9, 0xea, 0x23, 0x55, 0x9a, 0x09, 0x8e,
	0x86, 0x97, 0x9d, 0xc9, 0x60, 0x36, 0x8e, 0xfc, 0x8d, 0xae, 0x9a, 0x44, 0xef, 0x16, 0xdc, 0xbc,
	0x9c, 0xbd, 0x27, 0x59, 0x41, 0xf1, 0x23, 0x57, 0xf3, 0xde, 0x95, 0x40, 0x08, 0x7a, 0x9c, 0xe4,
	0x14, 0x75, 0x2e, 0x3b, 0x93, 0x43, 0x6c, 0xff, 0x86, 0xc7, 0xa0, 0xcb, 0x12, 0xf4, 0xcc, 0xae,
	0x74, 0x59, 0x02, 0x87, 0xa0, 0x6f, 0x98, 0xc9, 0x28, 0xea, 0xda, 0x25, 0xf7, 0x00, 0x23, 0x70,
	0x2a, 0x95, 0x48, 0x8a, 0x98, 0xaa, 0x95, 0x54, 0xe2, 0x4f, 0x1a, 0x9b, 0x15, 0x4b, 0xd0, 0xb9,
	0x75, 0x4e, 0x2a, 0xb4, 0x74, 0x64, 0x91, 0xc0, 0x09, 0xe8, 0x11, 0xc9, 0x34, 0x7a, 0x70, 0xf9,
	0x60, 0x32, 0x98, 0x0d, 0x1b, 0x43, 0x5e, 0x49, 0x86, 0xad, 0x01, 0xbf, 0x05, 0xfd, 0xf2, 0x95,
	0x68, 0xd4, 0xb3, 0xea, 0x59, 0x43, 0xfd, 0x6d, 0x2b, 0x29, 0x76, 0x4e, 0x29, 0x53, 0x5e, 0xe4,
	0x1a, 0xf5, 0xef, 0x91, 0xaf, 0x79, 0x91, 0x63, 0xe7, 0xc0, 0x1f, 0xc1, 0xa3, 0xda, 0x97, 0x83,
	0x1e, 0xda, 0x37, 0x76, 0x11, 0xed, 0x32, 0x20, 0x7a, 0x13, 0x0a, 0xb8, 0xee, 0xc3, 0x17, 0xe0,
	0xc0, 0x7f, 0xe2, 0xe8, 0x13, 0x5b, 0x7a, 0x1a, 0x96, 0xce, 0x1d, 0xc2, 0x95, 0x03, 0xbf, 0x02,
	0xbd, 0xf2, 0x13, 0x42, 0x87, 0xd6, 0x7d, 0x1c, 0xba, 0x3f, 0x1b, 0x23, 0xb1, 0xa5, 0xf0, 0x39,
	0xe8, 0xdb, 0xeb, 0x8a, 0x80, 0xd5, 0x4e, 0x42, 0xed, 0xd7, 0x12, 0x60, 0xc7, 0xe1, 0x1c, 0x1c,
	0x97, 0xb9, 0x43, 0xb9, 0x61, 0xb1, 0x9b, 0x7f, 0x60, 0x2b, 0x46, 0x61, 0xc5, 0x55, 0xcd, 0xc0,
	0x7b, 0x15, 0xe5, 0x0e, 0x7c, 0xe0, 0xa0, 0xa3, 0xe6, 0x0e, 0x5e, 0x3b, 0x84, 0x2b, 0xa7, 0x9c,
	0xcd, 0xde, 0x78, 0xf4, 0x69, 0x73, 0xb6, 0x77, 0x25, 0xc0, 0x8e, 0xc3, 0x19, 0x38, 0xac, 0x42,
	0x47, 0x23, 0x58, 0x3f, 0xe3, 0x52, 0xbe, 0xf6, 0x10, 0xef, 0xb4, 0x6a, 0x16, 0x25, 0x32, 0x74,
	0xd6, 0x3e, 0x8b, 0x12, 0x19, 0xae, 0x1c, 0xf8, 0x02, 0xf4, 0x32, 0x91, 0x6a, 0xf4, 0xc4, 0x76,
	0xaf, 0x1d, 0xda, 0x5b, 0x91, 0xbe, 0xa1, 0x3a, 0x56, 0x4c, 0x1a, 0xa1, 0xb0, 0xd5, 0xe0, 0x2b,
	0x70, 0xe0, 0x02, 0x46, 0x23, 0x64, 0x2b, 0xc6, 0x61, 0xc5, 0x2f, 0x16, 0x05, 0x45, 0x95, 0x0c,
	0x7f, 0x07, 0xa7, 0xcd, 0xfc, 0xd1, 0xe8, 0xc2, 0xf6, 0x78, 0x5e, 0xeb, 0x51, 0x69, 0xd8, 0x5b,
	0x41, 0x3b, 0x98, 0xef, 0x43, 0xbb, 0x5f, 0xff, 0x33, 0x80, 0x46, 0x2d, 0xb7, 0xc7, 0x21, 0x5c,
	0x39, 0xa5, 0xee, 0xb3, 0x13, 0x3d, 0x6d, 0xea, 0x6f, 0x1d, 0xc2, 0x95, 0x03, 0x5f, 0x01, 0xb0,
	0x8b, 0x44, 0x34, 0xb6, 0x15, 0xe7, 0x2d, 0xe3, 0x96, 0x45, 0x81, 0x09, 0x17, 0xe0, 0x64, 0x3f,
	0xf7, 0x34, 0xfa, 0xac, 0x1e, 0x25, 0x65, 0xf9, 0x8d, 0x95, 0x96, 0x77, 0x0e, 0x7e, 0xac, 0xf7,
	0x56, 0xe0, 0x77, 0x60, 0x10, 0x04, 0x2c, 0xfa, 0xba, 0x39, 0xc3, 0x8d, 0xc5, 0x0b, 0xbe, 0x11,
	0x18, 0xe8, 0xbb, 0xbf, 0xe1, 0x0f, 0xe0, 0x28, 0xfc, 0x29, 0x42, 0xd4, 0x56, 0xa2, 0xda, 0x05,
	0x0a, 0x38, 0xae, 0xd9, 0x73, 0x5e, 0x26, 0x61, 0x1e, 0xc8, 0xf3, 0x23, 0x1f, 0x92, 0xcb, 0x32,
	0x05, 0x96, 0x9d, 0x3f, 0xae, 0x3d, 0x4b, 0x45, 0x46, 0x78, 0x1a, 0x09, 0x95, 0x4e, 0x53, 0xca,
	0x6d, 0x46, 0x4c, 0x1d, 0x2a, 0x93, 0x27, 0xfc, 0xef, 0xc0, 0xc5, 0xe4, 0xf7, 0xb5, 0xa7, 0xff,
	0xba, 0xbd, 0x9f, 0xae, 0x96, 0x8b, 0xf5, 0x43, 0x5b, 0xf8, 0xf2, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0x6c, 0x4b, 0xf7, 0x55, 0x08, 0x00, 0x00,
}