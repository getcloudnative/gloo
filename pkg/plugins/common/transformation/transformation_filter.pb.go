// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transformation_filter.proto

/*
Package transformation is a generated protocol buffer package.

It is generated from these files:
	transformation_filter.proto

It has these top-level messages:
	Transformations
	Transformation
	Extraction
	TransformationTemplate
	InjaTemplate
	Passthrough
	MergeExtractorsToBody
*/
package transformation

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Transformations struct {
	Transformations   map[string]*Transformation `protobuf:"bytes,1,rep,name=transformations" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	AdvancedTemplates bool                       `protobuf:"varint,2,opt,name=advanced_templates,json=advancedTemplates,proto3" json:"advanced_templates,omitempty"`
}

func (m *Transformations) Reset()         { *m = Transformations{} }
func (m *Transformations) String() string { return proto.CompactTextString(m) }
func (*Transformations) ProtoMessage()    {}
func (*Transformations) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{0}
}

func (m *Transformations) GetTransformations() map[string]*Transformation {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *Transformations) GetAdvancedTemplates() bool {
	if m != nil {
		return m.AdvancedTemplates
	}
	return false
}

// [#proto-status: experimental]
type Transformation struct {
	// Extractors are in the origin request language domain
	Extractors map[string]*Extraction `protobuf:"bytes,1,rep,name=extractors" json:"extractors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Template is in the transformed request language domain
	// currently both are JSON
	TransformationTemplate *TransformationTemplate `protobuf:"bytes,2,opt,name=transformation_template,json=transformationTemplate" json:"transformation_template,omitempty"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{1}
}

func (m *Transformation) GetExtractors() map[string]*Extraction {
	if m != nil {
		return m.Extractors
	}
	return nil
}

func (m *Transformation) GetTransformationTemplate() *TransformationTemplate {
	if m != nil {
		return m.TransformationTemplate
	}
	return nil
}

type Extraction struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// what information to extract. if extraction fails the result is
	// an empty value.
	Regex    string `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	Subgroup uint32 `protobuf:"varint,3,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
}

func (m *Extraction) Reset()                    { *m = Extraction{} }
func (m *Extraction) String() string            { return proto.CompactTextString(m) }
func (*Extraction) ProtoMessage()               {}
func (*Extraction) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{2} }

func (m *Extraction) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Extraction) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

func (m *Extraction) GetSubgroup() uint32 {
	if m != nil {
		return m.Subgroup
	}
	return 0
}

type TransformationTemplate struct {
	Headers map[string]*InjaTemplate `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to BodyTransformation:
	//	*TransformationTemplate_Body
	//	*TransformationTemplate_Passthrough
	//	*TransformationTemplate_MergeExtractorsToBody
	BodyTransformation isTransformationTemplate_BodyTransformation `protobuf_oneof:"body_transformation"`
}

func (m *TransformationTemplate) Reset()         { *m = TransformationTemplate{} }
func (m *TransformationTemplate) String() string { return proto.CompactTextString(m) }
func (*TransformationTemplate) ProtoMessage()    {}
func (*TransformationTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{3}
}

type isTransformationTemplate_BodyTransformation interface {
	isTransformationTemplate_BodyTransformation()
}

type TransformationTemplate_Body struct {
	Body *InjaTemplate `protobuf:"bytes,2,opt,name=body,oneof"`
}
type TransformationTemplate_Passthrough struct {
	Passthrough *Passthrough `protobuf:"bytes,3,opt,name=passthrough,oneof"`
}
type TransformationTemplate_MergeExtractorsToBody struct {
	MergeExtractorsToBody *MergeExtractorsToBody `protobuf:"bytes,4,opt,name=merge_extractors_to_body,json=mergeExtractorsToBody,oneof"`
}

func (*TransformationTemplate_Body) isTransformationTemplate_BodyTransformation()                  {}
func (*TransformationTemplate_Passthrough) isTransformationTemplate_BodyTransformation()           {}
func (*TransformationTemplate_MergeExtractorsToBody) isTransformationTemplate_BodyTransformation() {}

func (m *TransformationTemplate) GetBodyTransformation() isTransformationTemplate_BodyTransformation {
	if m != nil {
		return m.BodyTransformation
	}
	return nil
}

func (m *TransformationTemplate) GetHeaders() map[string]*InjaTemplate {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TransformationTemplate) GetBody() *InjaTemplate {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Body); ok {
		return x.Body
	}
	return nil
}

func (m *TransformationTemplate) GetPassthrough() *Passthrough {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Passthrough); ok {
		return x.Passthrough
	}
	return nil
}

func (m *TransformationTemplate) GetMergeExtractorsToBody() *MergeExtractorsToBody {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_MergeExtractorsToBody); ok {
		return x.MergeExtractorsToBody
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransformationTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransformationTemplate_OneofMarshaler, _TransformationTemplate_OneofUnmarshaler, _TransformationTemplate_OneofSizer, []interface{}{
		(*TransformationTemplate_Body)(nil),
		(*TransformationTemplate_Passthrough)(nil),
		(*TransformationTemplate_MergeExtractorsToBody)(nil),
	}
}

func _TransformationTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransformationTemplate)
	// body_transformation
	switch x := m.BodyTransformation.(type) {
	case *TransformationTemplate_Body:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Body); err != nil {
			return err
		}
	case *TransformationTemplate_Passthrough:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Passthrough); err != nil {
			return err
		}
	case *TransformationTemplate_MergeExtractorsToBody:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MergeExtractorsToBody); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransformationTemplate.BodyTransformation has unexpected type %T", x)
	}
	return nil
}

func _TransformationTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransformationTemplate)
	switch tag {
	case 2: // body_transformation.body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InjaTemplate)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_Body{msg}
		return true, err
	case 3: // body_transformation.passthrough
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Passthrough)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_Passthrough{msg}
		return true, err
	case 4: // body_transformation.merge_extractors_to_body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MergeExtractorsToBody)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_MergeExtractorsToBody{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransformationTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransformationTemplate)
	// body_transformation
	switch x := m.BodyTransformation.(type) {
	case *TransformationTemplate_Body:
		s := proto.Size(x.Body)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransformationTemplate_Passthrough:
		s := proto.Size(x.Passthrough)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransformationTemplate_MergeExtractorsToBody:
		s := proto.Size(x.MergeExtractorsToBody)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//
// custom functions:
// header_value(name) -> from the original headers
// extracted_value(name, index) -> from the extracted values
type InjaTemplate struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *InjaTemplate) Reset()                    { *m = InjaTemplate{} }
func (m *InjaTemplate) String() string            { return proto.CompactTextString(m) }
func (*InjaTemplate) ProtoMessage()               {}
func (*InjaTemplate) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{4} }

func (m *InjaTemplate) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Passthrough struct {
}

func (m *Passthrough) Reset()                    { *m = Passthrough{} }
func (m *Passthrough) String() string            { return proto.CompactTextString(m) }
func (*Passthrough) ProtoMessage()               {}
func (*Passthrough) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{5} }

type MergeExtractorsToBody struct {
}

func (m *MergeExtractorsToBody) Reset()         { *m = MergeExtractorsToBody{} }
func (m *MergeExtractorsToBody) String() string { return proto.CompactTextString(m) }
func (*MergeExtractorsToBody) ProtoMessage()    {}
func (*MergeExtractorsToBody) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{6}
}

func init() {
	proto.RegisterType((*Transformations)(nil), "envoy.api.v2.filter.http.Transformations")
	proto.RegisterType((*Transformation)(nil), "envoy.api.v2.filter.http.Transformation")
	proto.RegisterType((*Extraction)(nil), "envoy.api.v2.filter.http.Extraction")
	proto.RegisterType((*TransformationTemplate)(nil), "envoy.api.v2.filter.http.TransformationTemplate")
	proto.RegisterType((*InjaTemplate)(nil), "envoy.api.v2.filter.http.InjaTemplate")
	proto.RegisterType((*Passthrough)(nil), "envoy.api.v2.filter.http.Passthrough")
	proto.RegisterType((*MergeExtractorsToBody)(nil), "envoy.api.v2.filter.http.MergeExtractorsToBody")
}

func init() { proto.RegisterFile("transformation_filter.proto", fileDescriptorTransformationFilter) }

var fileDescriptorTransformationFilter = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x86, 0x31, 0x90, 0x34, 0x19, 0x27, 0x21, 0x9d, 0x06, 0x62, 0xd1, 0x0b, 0xb2, 0xda, 0xca,
	0x97, 0xba, 0x15, 0xbd, 0x44, 0x51, 0x9a, 0x03, 0x52, 0x24, 0x72, 0xa8, 0x54, 0xad, 0x50, 0x5b,
	0xf5, 0x62, 0x2d, 0x78, 0x83, 0x49, 0xc1, 0x6b, 0xad, 0x17, 0x84, 0xdf, 0xa4, 0xcf, 0xd2, 0xb7,
	0xe9, 0x9b, 0x54, 0xd8, 0x5e, 0xb0, 0x2d, 0x5b, 0x72, 0x6f, 0x3b, 0x0c, 0xff, 0x37, 0x33, 0xff,
	0x8c, 0x0c, 0xaf, 0xa5, 0xa0, 0x7e, 0xf8, 0xc4, 0xc5, 0x8a, 0xca, 0x05, 0xf7, 0x9d, 0xa7, 0xc5,
	0x52, 0x32, 0x61, 0x07, 0x82, 0x4b, 0x8e, 0x06, 0xf3, 0x37, 0x3c, 0xb2, 0x69, 0xb0, 0xb0, 0x37,
	0x43, 0x3b, 0x4d, 0x79, 0x52, 0x06, 0xe6, 0xef, 0x26, 0x74, 0x26, 0x39, 0x65, 0x88, 0x1e, 0x74,
	0xf2, 0xb0, 0xd0, 0xd0, 0x06, 0x2d, 0x4b, 0x1f, 0xde, 0xdb, 0x55, 0x1c, 0xbb, 0xc0, 0x28, 0xc6,
	0x0f, 0xbe, 0x14, 0x11, 0x29, 0x62, 0xf1, 0x3d, 0x20, 0x75, 0x37, 0xd4, 0x9f, 0x31, 0xd7, 0x91,
	0x6c, 0x15, 0x2c, 0xa9, 0x64, 0xa1, 0xd1, 0x1c, 0x68, 0xd6, 0x09, 0x79, 0xa9, 0x32, 0x13, 0x95,
	0xe8, 0x2f, 0xe1, 0xaa, 0x8c, 0x8b, 0x97, 0xd0, 0xfa, 0xc5, 0x22, 0x43, 0x1b, 0x68, 0xd6, 0x29,
	0xd9, 0x3d, 0xf1, 0x1e, 0x8e, 0x36, 0x74, 0xb9, 0x66, 0x31, 0x4b, 0x1f, 0x5a, 0x75, 0x1b, 0x27,
	0x89, 0xec, 0xb6, 0x79, 0xa3, 0x99, 0x7f, 0x9a, 0x70, 0x91, 0xcf, 0xe2, 0x0f, 0x00, 0xb6, 0x95,
	0x82, 0xce, 0x24, 0x17, 0xca, 0x94, 0x9b, 0xba, 0x6c, 0xfb, 0x61, 0x2f, 0x4d, 0xec, 0xc8, 0xb0,
	0x70, 0x01, 0xd7, 0x85, 0x05, 0x2a, 0x3f, 0xd2, 0x11, 0x3e, 0xd6, 0x2d, 0xa3, 0xec, 0x22, 0x3d,
	0x59, 0xfa, 0x7b, 0x7f, 0x06, 0x9d, 0x42, 0x27, 0x25, 0x06, 0xde, 0xe6, 0x0d, 0x7c, 0x53, 0x5d,
	0x3d, 0x65, 0x15, 0xcc, 0xfb, 0x06, 0x70, 0x48, 0x60, 0x0f, 0x8e, 0x3d, 0x46, 0x5d, 0x26, 0xd2,
	0x12, 0x69, 0x84, 0x57, 0x70, 0x24, 0xd8, 0x9c, 0x6d, 0xe3, 0x2a, 0xa7, 0x24, 0x09, 0xb0, 0x0f,
	0x27, 0xe1, 0x7a, 0x3a, 0x17, 0x7c, 0x1d, 0x18, 0xad, 0x81, 0x66, 0x9d, 0x93, 0x7d, 0x6c, 0xfe,
	0x6d, 0x41, 0xaf, 0x7c, 0x5e, 0xfc, 0x0e, 0x2f, 0x12, 0xac, 0xda, 0xcc, 0xe7, 0xff, 0xb5, 0xcc,
	0x1e, 0x27, 0xfa, 0x64, 0x3d, 0x8a, 0x86, 0x77, 0xd0, 0x9e, 0x72, 0x37, 0x4a, 0xad, 0x78, 0x57,
	0x4d, 0x7d, 0xf4, 0x9f, 0xa9, 0x62, 0x8d, 0x1b, 0x24, 0x56, 0xe1, 0x23, 0xe8, 0x01, 0x0d, 0x43,
	0xe9, 0x09, 0xbe, 0x9e, 0x7b, 0xf1, 0x40, 0xfa, 0xf0, 0x6d, 0x35, 0xe4, 0xeb, 0xe1, 0xcf, 0xe3,
	0x06, 0xc9, 0x6a, 0xf1, 0x19, 0x8c, 0x15, 0x13, 0x73, 0xe6, 0x1c, 0x0e, 0xc7, 0x91, 0xdc, 0x89,
	0x9b, 0x6b, 0xc7, 0xdc, 0x0f, 0xd5, 0xdc, 0x2f, 0x3b, 0xe5, 0x61, 0xf1, 0x13, 0x3e, 0xe2, 0x6e,
	0x34, 0x6e, 0x90, 0xee, 0xaa, 0x2c, 0xd1, 0x9f, 0xc2, 0x59, 0xd6, 0x8d, 0x92, 0x13, 0xb9, 0xcb,
	0x9f, 0x48, 0x4d, 0x5f, 0x32, 0x47, 0x32, 0xea, 0xc2, 0xab, 0x5d, 0xef, 0x4e, 0xfe, 0x50, 0x4d,
	0x13, 0xce, 0xb2, 0x0a, 0x44, 0x68, 0x4b, 0xb6, 0x95, 0x69, 0xed, 0xf8, 0x6d, 0x9e, 0x83, 0x9e,
	0x31, 0xca, 0xbc, 0x86, 0x6e, 0xe9, 0x7c, 0xa3, 0xcb, 0x9f, 0x17, 0x79, 0xfa, 0xf4, 0x38, 0xfe,
	0x24, 0x7e, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x8c, 0xe6, 0xc8, 0x31, 0x05, 0x00, 0x00,
}