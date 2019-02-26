// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/specs/FeatureSpec.proto

package specs // import "github.com/gojek/feast/protos/generated/go/feast/specs"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gojek/feast/protos/generated/go/feast/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeatureSpec struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string                 `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Description          string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Uri                  string                 `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	Granularity          types.Granularity_Enum `protobuf:"varint,6,opt,name=granularity,proto3,enum=feast.types.Granularity_Enum" json:"granularity,omitempty"`
	ValueType            types.ValueType_Enum   `protobuf:"varint,7,opt,name=valueType,proto3,enum=feast.types.ValueType_Enum" json:"valueType,omitempty"`
	Entity               string                 `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Group                string                 `protobuf:"bytes,9,opt,name=group,proto3" json:"group,omitempty"`
	Tags                 []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Options              map[string]string      `protobuf:"bytes,11,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DataStores           *DataStores            `protobuf:"bytes,12,opt,name=dataStores,proto3" json:"dataStores,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FeatureSpec) Reset()         { *m = FeatureSpec{} }
func (m *FeatureSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSpec) ProtoMessage()    {}
func (*FeatureSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f1468f11147fbe, []int{0}
}
func (m *FeatureSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSpec.Unmarshal(m, b)
}
func (m *FeatureSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSpec.Marshal(b, m, deterministic)
}
func (m *FeatureSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSpec.Merge(m, src)
}
func (m *FeatureSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSpec.Size(m)
}
func (m *FeatureSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSpec proto.InternalMessageInfo

func (m *FeatureSpec) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FeatureSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSpec) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FeatureSpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FeatureSpec) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *FeatureSpec) GetGranularity() types.Granularity_Enum {
	if m != nil {
		return m.Granularity
	}
	return types.Granularity_NONE
}

func (m *FeatureSpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_UNKNOWN
}

func (m *FeatureSpec) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *FeatureSpec) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *FeatureSpec) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FeatureSpec) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *FeatureSpec) GetDataStores() *DataStores {
	if m != nil {
		return m.DataStores
	}
	return nil
}

type DataStores struct {
	Serving              *DataStore `protobuf:"bytes,1,opt,name=serving,proto3" json:"serving,omitempty"`
	Warehouse            *DataStore `protobuf:"bytes,2,opt,name=warehouse,proto3" json:"warehouse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DataStores) Reset()         { *m = DataStores{} }
func (m *DataStores) String() string { return proto.CompactTextString(m) }
func (*DataStores) ProtoMessage()    {}
func (*DataStores) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f1468f11147fbe, []int{1}
}
func (m *DataStores) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStores.Unmarshal(m, b)
}
func (m *DataStores) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStores.Marshal(b, m, deterministic)
}
func (m *DataStores) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStores.Merge(m, src)
}
func (m *DataStores) XXX_Size() int {
	return xxx_messageInfo_DataStores.Size(m)
}
func (m *DataStores) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStores.DiscardUnknown(m)
}

var xxx_messageInfo_DataStores proto.InternalMessageInfo

func (m *DataStores) GetServing() *DataStore {
	if m != nil {
		return m.Serving
	}
	return nil
}

func (m *DataStores) GetWarehouse() *DataStore {
	if m != nil {
		return m.Warehouse
	}
	return nil
}

type DataStore struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Options              map[string]string `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataStore) Reset()         { *m = DataStore{} }
func (m *DataStore) String() string { return proto.CompactTextString(m) }
func (*DataStore) ProtoMessage()    {}
func (*DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f1468f11147fbe, []int{2}
}
func (m *DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStore.Unmarshal(m, b)
}
func (m *DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStore.Marshal(b, m, deterministic)
}
func (m *DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStore.Merge(m, src)
}
func (m *DataStore) XXX_Size() int {
	return xxx_messageInfo_DataStore.Size(m)
}
func (m *DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_DataStore proto.InternalMessageInfo

func (m *DataStore) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataStore) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*FeatureSpec)(nil), "feast.specs.FeatureSpec")
	proto.RegisterMapType((map[string]string)(nil), "feast.specs.FeatureSpec.OptionsEntry")
	proto.RegisterType((*DataStores)(nil), "feast.specs.DataStores")
	proto.RegisterType((*DataStore)(nil), "feast.specs.DataStore")
	proto.RegisterMapType((map[string]string)(nil), "feast.specs.DataStore.OptionsEntry")
}

func init() { proto.RegisterFile("feast/specs/FeatureSpec.proto", fileDescriptor_b8f1468f11147fbe) }

var fileDescriptor_b8f1468f11147fbe = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0x56, 0x62, 0x57, 0xa3, 0x10, 0xc2, 0x52, 0x92, 0xc5, 0x6d, 0x40, 0xb8, 0x14, 0x7c,
	0x92, 0x8a, 0x5b, 0xfa, 0x11, 0x28, 0x81, 0x50, 0xb7, 0xc7, 0x16, 0xa5, 0xa4, 0xd0, 0xdb, 0xc6,
	0x9a, 0x2a, 0x6a, 0x62, 0xad, 0xd8, 0x5d, 0x25, 0xe8, 0x8f, 0xf4, 0x6f, 0xf6, 0x2f, 0x94, 0x9d,
	0xb5, 0xa3, 0x35, 0x71, 0x4f, 0xb9, 0xcd, 0xec, 0x7b, 0x33, 0x9a, 0xf7, 0x46, 0x03, 0xc7, 0xbf,
	0x50, 0x68, 0x93, 0xe9, 0x06, 0x17, 0x3a, 0xfb, 0x8c, 0xc2, 0xb4, 0x0a, 0xcf, 0x1b, 0x5c, 0xa4,
	0x8d, 0x92, 0x46, 0xb2, 0x98, 0xe0, 0x94, 0xe0, 0xf1, 0x73, 0x9f, 0x3b, 0xaf, 0x4d, 0x65, 0xba,
	0x9e, 0x3a, 0xde, 0xe8, 0x74, 0x6e, 0xa4, 0x12, 0x25, 0x3e, 0x84, 0x4d, 0xd7, 0xa0, 0xce, 0xbe,
	0x28, 0x51, 0xb7, 0x37, 0x42, 0x55, 0xa6, 0x5b, 0xc1, 0x47, 0x3e, 0x7c, 0x21, 0x6e, 0x5a, 0x74,
	0xc0, 0xe4, 0x6f, 0x08, 0xb1, 0x37, 0x17, 0xdb, 0x87, 0x41, 0x55, 0xf0, 0x20, 0x09, 0xa6, 0x51,
	0x3e, 0xa8, 0x0a, 0xc6, 0x60, 0xa7, 0x16, 0x4b, 0xe4, 0x03, 0x7a, 0xa1, 0x98, 0x3d, 0x85, 0x5d,
	0x79, 0x57, 0xa3, 0xe2, 0x21, 0x3d, 0xba, 0x84, 0x25, 0x10, 0x17, 0xa8, 0x17, 0xaa, 0x6a, 0x4c,
	0x25, 0x6b, 0xbe, 0x43, 0x98, 0xff, 0xc4, 0x0e, 0x20, 0x6c, 0x55, 0xc5, 0x77, 0x09, 0xb1, 0x21,
	0x3b, 0x85, 0xb8, 0xec, 0x67, 0xe5, 0xc3, 0x24, 0x98, 0xee, 0xcf, 0x8e, 0x53, 0xe7, 0x0a, 0x0d,
	0x9b, 0xfa, 0x5a, 0xe6, 0x75, 0xbb, 0xcc, 0xfd, 0x0a, 0xf6, 0x01, 0xa2, 0x5b, 0xab, 0xe6, 0x7b,
	0xd7, 0x20, 0x1f, 0x51, 0xf9, 0xb3, 0x8d, 0xf2, 0x8b, 0x35, 0xea, 0x8a, 0x7b, 0x36, 0x3b, 0x84,
	0x21, 0x92, 0xc9, 0xfc, 0x09, 0x0d, 0xb4, 0xca, 0xac, 0xba, 0x52, 0xc9, 0xb6, 0xe1, 0x91, 0x53,
	0x47, 0x89, 0xf5, 0xc1, 0x88, 0x52, 0x73, 0x48, 0x42, 0xeb, 0x83, 0x8d, 0xd9, 0x29, 0x8c, 0x24,
	0x29, 0xd3, 0x3c, 0x4e, 0xc2, 0x69, 0x3c, 0x7b, 0x99, 0x7a, 0xfb, 0x4c, 0xfd, 0x75, 0x7f, 0x75,
	0xbc, 0x79, 0x6d, 0x54, 0x97, 0xaf, 0xab, 0xd8, 0x3b, 0x80, 0x42, 0x18, 0x61, 0xb7, 0x89, 0x9a,
	0xef, 0x25, 0xc1, 0x34, 0x9e, 0x1d, 0x6d, 0xf4, 0xf8, 0x74, 0x0f, 0xe7, 0x1e, 0x75, 0x7c, 0x02,
	0x7b, 0x7e, 0x47, 0xeb, 0xec, 0x35, 0x76, 0xab, 0xb5, 0xd9, 0xd0, 0xaa, 0x20, 0xa9, 0xab, 0xc5,
	0xb9, 0xe4, 0x64, 0xf0, 0x3e, 0x98, 0x18, 0x80, 0xbe, 0x2b, 0x7b, 0x05, 0x23, 0x8d, 0xea, 0xb6,
	0xaa, 0x4b, 0xaa, 0x8e, 0x67, 0x87, 0xdb, 0xbf, 0x9f, 0xaf, 0x69, 0xec, 0x0d, 0x44, 0x77, 0x42,
	0xe1, 0x95, 0x6c, 0xb5, 0xeb, 0xfe, 0xff, 0x9a, 0x9e, 0x38, 0xf9, 0x13, 0x40, 0x74, 0x0f, 0x3c,
	0xf8, 0xcb, 0x3e, 0xf6, 0x4e, 0x0e, 0xc8, 0xc9, 0x17, 0xdb, 0x3b, 0x6e, 0xf7, 0xf1, 0x31, 0x76,
	0x9c, 0xfd, 0x00, 0xff, 0x08, 0xcf, 0x0e, 0xbc, 0xad, 0x7d, 0xb3, 0x17, 0xf2, 0xf3, 0x6d, 0x59,
	0x99, 0xab, 0xf6, 0x32, 0x5d, 0xc8, 0x65, 0x56, 0xca, 0xdf, 0x78, 0x9d, 0xb9, 0x5b, 0xa2, 0xfb,
	0xd1, 0x59, 0x89, 0x35, 0x2a, 0x61, 0xb0, 0xc8, 0x4a, 0x99, 0x79, 0x37, 0x7a, 0x39, 0x24, 0xc2,
	0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0x00, 0x0c, 0x12, 0x03, 0x04, 0x00, 0x00,
}
