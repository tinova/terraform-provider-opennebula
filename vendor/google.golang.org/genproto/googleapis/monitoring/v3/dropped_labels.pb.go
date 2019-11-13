// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/dropped_labels.proto

package monitoring

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A set of (label, value) pairs which were dropped during aggregation, attached
// to google.api.Distribution.Exemplars in google.api.Distribution values during
// aggregation.
//
// These values are used in combination with the label values that remain on the
// aggregated Distribution timeseries to construct the full label set for the
// exemplar values.  The resulting full label set may be used to identify the
// specific task/job/instance (for example) which may be contributing to a
// long-tail, while allowing the storage savings of only storing aggregated
// distribution values for a large group.
//
// Note that there are no guarantees on ordering of the labels from
// exemplar-to-exemplar and from distribution-to-distribution in the same
// stream, and there may be duplicates.  It is up to clients to resolve any
// ambiguities.
type DroppedLabels struct {
	// Map from label to its value, for all labels dropped in any aggregation.
	Label                map[string]string `protobuf:"bytes,1,rep,name=label,proto3" json:"label,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DroppedLabels) Reset()         { *m = DroppedLabels{} }
func (m *DroppedLabels) String() string { return proto.CompactTextString(m) }
func (*DroppedLabels) ProtoMessage()    {}
func (*DroppedLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_15749142c06d7f43, []int{0}
}

func (m *DroppedLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DroppedLabels.Unmarshal(m, b)
}
func (m *DroppedLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DroppedLabels.Marshal(b, m, deterministic)
}
func (m *DroppedLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroppedLabels.Merge(m, src)
}
func (m *DroppedLabels) XXX_Size() int {
	return xxx_messageInfo_DroppedLabels.Size(m)
}
func (m *DroppedLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_DroppedLabels.DiscardUnknown(m)
}

var xxx_messageInfo_DroppedLabels proto.InternalMessageInfo

func (m *DroppedLabels) GetLabel() map[string]string {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*DroppedLabels)(nil), "google.monitoring.v3.DroppedLabels")
	proto.RegisterMapType((map[string]string)(nil), "google.monitoring.v3.DroppedLabels.LabelEntry")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/dropped_labels.proto", fileDescriptor_15749142c06d7f43)
}

var fileDescriptor_15749142c06d7f43 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0xd7, 0x2f, 0x33,
	0xd6, 0x4f, 0x29, 0xca, 0x2f, 0x28, 0x48, 0x4d, 0x89, 0xcf, 0x49, 0x4c, 0x4a, 0xcd, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81, 0x28, 0xd5, 0x43, 0x28, 0xd5, 0x2b, 0x33, 0x56,
	0xea, 0x67, 0xe4, 0xe2, 0x75, 0x81, 0x28, 0xf7, 0x01, 0xab, 0x16, 0x72, 0xe1, 0x62, 0x05, 0xeb,
	0x93, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xd2, 0xd3, 0xc3, 0xa6, 0x4f, 0x0f, 0x45, 0x8f, 0x1e,
	0x98, 0x72, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x82, 0x68, 0x96, 0xb2, 0xe0, 0xe2, 0x42, 0x08, 0x0a,
	0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42,
	0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60, 0x31, 0x08, 0xc7, 0x8a, 0xc9,
	0x82, 0xd1, 0xc9, 0x21, 0xca, 0x0e, 0x6a, 0x63, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e,
	0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0xd8, 0x17, 0xfa, 0x10, 0xa9, 0xc4, 0x82, 0xcc, 0x62, 0x54,
	0x3f, 0x5b, 0x23, 0x78, 0x49, 0x6c, 0x60, 0xa5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97,
	0xe6, 0x1c, 0x2d, 0x1d, 0x01, 0x00, 0x00,
}