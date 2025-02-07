// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/phosae/kube-code-generator/example/apis/comic/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *Hero) Reset()      { *m = Hero{} }
func (*Hero) ProtoMessage() {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7c7b1ccb0563d8, []int{0}
}
func (m *Hero) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(m, src)
}
func (m *Hero) XXX_Size() int {
	return m.Size()
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *HeroList) Reset()      { *m = HeroList{} }
func (*HeroList) ProtoMessage() {}
func (*HeroList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7c7b1ccb0563d8, []int{1}
}
func (m *HeroList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeroList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HeroList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroList.Merge(m, src)
}
func (m *HeroList) XXX_Size() int {
	return m.Size()
}
func (m *HeroList) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroList.DiscardUnknown(m)
}

var xxx_messageInfo_HeroList proto.InternalMessageInfo

func (m *HeroSpec) Reset()      { *m = HeroSpec{} }
func (*HeroSpec) ProtoMessage() {}
func (*HeroSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7c7b1ccb0563d8, []int{2}
}
func (m *HeroSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeroSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HeroSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroSpec.Merge(m, src)
}
func (m *HeroSpec) XXX_Size() int {
	return m.Size()
}
func (m *HeroSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroSpec.DiscardUnknown(m)
}

var xxx_messageInfo_HeroSpec proto.InternalMessageInfo

func (m *HeroStatus) Reset()      { *m = HeroStatus{} }
func (*HeroStatus) ProtoMessage() {}
func (*HeroStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7c7b1ccb0563d8, []int{3}
}
func (m *HeroStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeroStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HeroStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroStatus.Merge(m, src)
}
func (m *HeroStatus) XXX_Size() int {
	return m.Size()
}
func (m *HeroStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HeroStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Hero)(nil), "github.com.phosae.kube_code_generator.example.apis.comic.v1.Hero")
	proto.RegisterType((*HeroList)(nil), "github.com.phosae.kube_code_generator.example.apis.comic.v1.HeroList")
	proto.RegisterType((*HeroSpec)(nil), "github.com.phosae.kube_code_generator.example.apis.comic.v1.HeroSpec")
	proto.RegisterType((*HeroStatus)(nil), "github.com.phosae.kube_code_generator.example.apis.comic.v1.HeroStatus")
}

func init() {
	proto.RegisterFile("github.com/phosae/kube-code-generator/example/apis/comic/v1/generated.proto", fileDescriptor_2e7c7b1ccb0563d8)
}

var fileDescriptor_2e7c7b1ccb0563d8 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xc7, 0xe3, 0xd4, 0xad, 0x92, 0xf3, 0xaf, 0x3f, 0xd0, 0xb1, 0x58, 0x1d, 0x9c, 0x28, 0x03,
	0xaa, 0x10, 0x3d, 0x93, 0x0a, 0x10, 0x52, 0x27, 0x5c, 0x10, 0xa0, 0x52, 0x40, 0x6e, 0x25, 0x24,
	0x84, 0x54, 0x2e, 0xce, 0xab, 0x73, 0x98, 0xf3, 0x59, 0xf6, 0x39, 0x90, 0x8d, 0x91, 0x91, 0xbf,
	0x89, 0xa9, 0x63, 0xc7, 0x4e, 0x11, 0x35, 0xff, 0x05, 0x13, 0xba, 0xb3, 0x53, 0x5b, 0x14, 0x44,
	0xa5, 0x6c, 0xb9, 0xf7, 0xbe, 0xef, 0xf3, 0x7d, 0xf7, 0x3d, 0x2b, 0x68, 0x2f, 0x64, 0x72, 0x92,
	0x8f, 0x48, 0x20, 0xb8, 0x9b, 0x4c, 0x44, 0x46, 0xc1, 0x8d, 0xf2, 0x11, 0x6c, 0x05, 0x62, 0x0c,
	0x5b, 0x21, 0xc4, 0x90, 0x52, 0x29, 0x52, 0x17, 0x3e, 0x51, 0x9e, 0x7c, 0x00, 0x97, 0x26, 0x2c,
	0x73, 0x03, 0xc1, 0x59, 0xe0, 0x4e, 0x87, 0x6e, 0xd5, 0x87, 0x31, 0x49, 0x52, 0x21, 0x05, 0xde,
	0xa9, 0x61, 0xa4, 0x84, 0x11, 0x05, 0x3b, 0x52, 0xb0, 0xa3, 0x0b, 0x18, 0xa9, 0x60, 0x44, 0xc1,
	0x88, 0x86, 0x91, 0xe9, 0x70, 0x63, 0xab, 0xb1, 0x49, 0x28, 0x42, 0xe1, 0x6a, 0xe6, 0x28, 0x3f,
	0xd6, 0x27, 0x7d, 0xd0, 0xbf, 0x4a, 0xaf, 0x8d, 0xbb, 0xd1, 0x83, 0x8c, 0x30, 0xa1, 0x36, 0xe2,
	0x34, 0x98, 0xb0, 0x18, 0xd2, 0x99, 0x9b, 0x44, 0x61, 0xb9, 0x22, 0x07, 0x49, 0xff, 0xb0, 0xe1,
	0x86, 0xfb, 0xb7, 0xa9, 0x34, 0x8f, 0x25, 0xe3, 0x70, 0x69, 0xe0, 0xfe, 0xbf, 0x06, 0xb2, 0x60,
	0x02, 0x9c, 0xfe, 0x3e, 0x37, 0xf8, 0xd6, 0x46, 0xe6, 0x53, 0x48, 0x05, 0x7e, 0x87, 0x3a, 0x6a,
	0x99, 0x31, 0x95, 0xd4, 0x36, 0xfa, 0xc6, 0xa6, 0xb5, 0x7d, 0x87, 0x94, 0x4c, 0xd2, 0x64, 0x92,
	0x24, 0x0a, 0xcb, 0x40, 0x94, 0x9a, 0x4c, 0x87, 0xe4, 0xe5, 0xe8, 0x3d, 0x04, 0x72, 0x1f, 0x24,
	0xf5, 0xf0, 0xc9, 0xbc, 0xd7, 0x2a, 0xe6, 0x3d, 0x54, 0xd7, 0xfc, 0x0b, 0x2a, 0x0e, 0x91, 0x99,
	0x25, 0x10, 0xd8, 0x6d, 0x4d, 0x7f, 0x4c, 0x96, 0x78, 0x04, 0xa2, 0x56, 0x3e, 0x48, 0x20, 0xf0,
	0xfe, 0xab, 0x2c, 0x4d, 0x75, 0xf2, 0xb5, 0x01, 0x16, 0x68, 0x2d, 0x93, 0x54, 0xe6, 0x99, 0xbd,
	0xa2, 0xad, 0x9e, 0x2c, 0x6f, 0xa5, 0x71, 0xde, 0xff, 0x95, 0xd9, 0x5a, 0x79, 0xf6, 0x2b, 0x9b,
	0xc1, 0xa9, 0x81, 0x3a, 0x4a, 0xf6, 0x9c, 0x65, 0x12, 0xbf, 0xbd, 0x14, 0x24, 0xb9, 0x5a, 0x90,
	0x6a, 0x5a, 0xc7, 0x78, 0xbd, 0xb2, 0xe9, 0x2c, 0x2a, 0x8d, 0x10, 0x8f, 0xd1, 0x2a, 0x93, 0xc0,
	0x33, 0xbb, 0xdd, 0x5f, 0xd9, 0xb4, 0xb6, 0x1f, 0x2e, 0x7d, 0x35, 0x6f, 0xbd, 0x72, 0x5b, 0x7d,
	0xa6, 0xb8, 0x7e, 0x89, 0x1f, 0x7c, 0x69, 0x97, 0x57, 0x52, 0xb1, 0xe2, 0x3e, 0x32, 0x63, 0xca,
	0x41, 0x5f, 0xa7, 0x5b, 0x47, 0xfe, 0x82, 0x72, 0xf0, 0x75, 0x47, 0x29, 0x02, 0x26, 0x67, 0xfa,
	0x6d, 0x1b, 0x8a, 0x5d, 0x26, 0x67, 0xbe, 0xee, 0xe0, 0xdb, 0xc8, 0x8c, 0x58, 0x3c, 0xd6, 0x4f,
	0xd2, 0xf5, 0xec, 0x85, 0x62, 0x8f, 0xc5, 0xe3, 0x9f, 0xf3, 0x9e, 0xf6, 0x3a, 0x9c, 0x25, 0xe0,
	0x6b, 0x15, 0x7e, 0x8d, 0xba, 0x23, 0x96, 0xca, 0xc9, 0x23, 0x2a, 0xc1, 0x36, 0x75, 0x8a, 0xb7,
	0xae, 0x96, 0xe2, 0x21, 0xe3, 0xe0, 0xad, 0x17, 0xf3, 0x5e, 0xd7, 0x5b, 0x00, 0xfc, 0x9a, 0x85,
	0x87, 0xc8, 0xca, 0xf2, 0x04, 0xd2, 0x57, 0xe2, 0x23, 0xa4, 0x99, 0xbd, 0xda, 0x5f, 0xd9, 0xec,
	0x7a, 0xd7, 0x8a, 0x79, 0xcf, 0x3a, 0xa8, 0xcb, 0x7e, 0x53, 0x33, 0x88, 0x10, 0xaa, 0xbf, 0x01,
	0x7c, 0x13, 0xad, 0x71, 0x31, 0x65, 0x71, 0xa8, 0xd3, 0xe8, 0xd4, 0xdf, 0xc4, 0xbe, 0xae, 0xfa,
	0x55, 0x17, 0xdf, 0x43, 0x56, 0x90, 0xa7, 0x29, 0xc4, 0x72, 0xb7, 0x0e, 0xe6, 0x46, 0x25, 0xb6,
	0x76, 0xeb, 0x96, 0xdf, 0xd4, 0x79, 0xf4, 0xe4, 0xdc, 0x69, 0x9d, 0x9e, 0x3b, 0xad, 0xb3, 0x73,
	0xa7, 0xf5, 0xb9, 0x70, 0x8c, 0x93, 0xc2, 0x31, 0x4e, 0x0b, 0xc7, 0x38, 0x2b, 0x1c, 0xe3, 0x7b,
	0xe1, 0x18, 0x5f, 0x7f, 0x38, 0xad, 0x37, 0x3b, 0x4b, 0xfc, 0x1b, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xf8, 0xf6, 0xaa, 0x4b, 0x05, 0x00, 0x00,
}

func (m *Hero) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hero) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Hero) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HeroList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeroList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeroList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HeroSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeroSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeroSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuperPowers) > 0 {
		for iNdEx := len(m.SuperPowers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SuperPowers[iNdEx])
			copy(dAtA[i:], m.SuperPowers[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.SuperPowers[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.BirthDate != nil {
		{
			size, err := m.BirthDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i -= len(m.Kind)
	copy(dAtA[i:], m.Kind)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Kind)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.City)
	copy(dAtA[i:], m.City)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.City)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HeroStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeroStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeroStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.CurrentCity)
	copy(dAtA[i:], m.CurrentCity)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.CurrentCity)))
	i--
	dAtA[i] = 0x12
	i--
	if m.Moving {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Hero) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *HeroList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *HeroSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.City)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	if m.BirthDate != nil {
		l = m.BirthDate.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.SuperPowers) > 0 {
		for _, s := range m.SuperPowers {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *HeroStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	l = len(m.CurrentCity)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Hero) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Hero{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "HeroSpec", "HeroSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "HeroStatus", "HeroStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HeroList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Hero{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Hero", "Hero", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&HeroList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *HeroSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HeroSpec{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`City:` + fmt.Sprintf("%v", this.City) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`BirthDate:` + strings.Replace(fmt.Sprintf("%v", this.BirthDate), "Time", "v1.Time", 1) + `,`,
		`SuperPowers:` + fmt.Sprintf("%v", this.SuperPowers) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HeroStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HeroStatus{`,
		`Moving:` + fmt.Sprintf("%v", this.Moving) + `,`,
		`CurrentCity:` + fmt.Sprintf("%v", this.CurrentCity) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Hero) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Hero: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hero: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeroList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeroList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeroList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Hero{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeroSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeroSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeroSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = HeroType(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BirthDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BirthDate == nil {
				m.BirthDate = &v1.Time{}
			}
			if err := m.BirthDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperPowers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuperPowers = append(m.SuperPowers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HeroStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HeroStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeroStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moving", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Moving = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentCity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentCity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
