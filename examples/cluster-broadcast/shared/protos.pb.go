// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos.proto

package shared

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Noop struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Noop) Reset()      { *m = Noop{} }
func (*Noop) ProtoMessage() {}
func (*Noop) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_c6fbb75b0b9817be, []int{0}
}
func (m *Noop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Noop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Noop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Noop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Noop.Merge(dst, src)
}
func (m *Noop) XXX_Size() int {
	return m.Size()
}
func (m *Noop) XXX_DiscardUnknown() {
	xxx_messageInfo_Noop.DiscardUnknown(m)
}

var xxx_messageInfo_Noop proto.InternalMessageInfo

type NumberRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberRequest) Reset()      { *m = NumberRequest{} }
func (*NumberRequest) ProtoMessage() {}
func (*NumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_c6fbb75b0b9817be, []int{1}
}
func (m *NumberRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NumberRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberRequest.Merge(dst, src)
}
func (m *NumberRequest) XXX_Size() int {
	return m.Size()
}
func (m *NumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumberRequest proto.InternalMessageInfo

func (m *NumberRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CountResponse struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()      { *m = CountResponse{} }
func (*CountResponse) ProtoMessage() {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_c6fbb75b0b9817be, []int{2}
}
func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(dst, src)
}
func (m *CountResponse) XXX_Size() int {
	return m.Size()
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type RegisterMessage struct {
	GrainId              string   `protobuf:"bytes,1,opt,name=grain_id,json=grainId,proto3" json:"grain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterMessage) Reset()      { *m = RegisterMessage{} }
func (*RegisterMessage) ProtoMessage() {}
func (*RegisterMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_c6fbb75b0b9817be, []int{3}
}
func (m *RegisterMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *RegisterMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterMessage.Merge(dst, src)
}
func (m *RegisterMessage) XXX_Size() int {
	return m.Size()
}
func (m *RegisterMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterMessage proto.InternalMessageInfo

func (m *RegisterMessage) GetGrainId() string {
	if m != nil {
		return m.GrainId
	}
	return ""
}

type TotalsResponse struct {
	Totals               map[string]int64 `protobuf:"bytes,1,rep,name=totals" json:"totals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TotalsResponse) Reset()      { *m = TotalsResponse{} }
func (*TotalsResponse) ProtoMessage() {}
func (*TotalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_protos_c6fbb75b0b9817be, []int{4}
}
func (m *TotalsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TotalsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalsResponse.Merge(dst, src)
}
func (m *TotalsResponse) XXX_Size() int {
	return m.Size()
}
func (m *TotalsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TotalsResponse proto.InternalMessageInfo

func (m *TotalsResponse) GetTotals() map[string]int64 {
	if m != nil {
		return m.Totals
	}
	return nil
}

func init() {
	proto.RegisterType((*Noop)(nil), "shared.Noop")
	proto.RegisterType((*NumberRequest)(nil), "shared.NumberRequest")
	proto.RegisterType((*CountResponse)(nil), "shared.CountResponse")
	proto.RegisterType((*RegisterMessage)(nil), "shared.RegisterMessage")
	proto.RegisterType((*TotalsResponse)(nil), "shared.TotalsResponse")
	proto.RegisterMapType((map[string]int64)(nil), "shared.TotalsResponse.TotalsEntry")
}
func (this *Noop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Noop)
	if !ok {
		that2, ok := that.(Noop)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *NumberRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NumberRequest)
	if !ok {
		that2, ok := that.(NumberRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	return true
}
func (this *CountResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CountResponse)
	if !ok {
		that2, ok := that.(CountResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	return true
}
func (this *RegisterMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisterMessage)
	if !ok {
		that2, ok := that.(RegisterMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GrainId != that1.GrainId {
		return false
	}
	return true
}
func (this *TotalsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TotalsResponse)
	if !ok {
		that2, ok := that.(TotalsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Totals) != len(that1.Totals) {
		return false
	}
	for i := range this.Totals {
		if this.Totals[i] != that1.Totals[i] {
			return false
		}
	}
	return true
}
func (this *Noop) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&shared.Noop{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NumberRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.NumberRequest{")
	s = append(s, "Number: "+fmt.Sprintf("%#v", this.Number)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CountResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.CountResponse{")
	s = append(s, "Number: "+fmt.Sprintf("%#v", this.Number)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RegisterMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.RegisterMessage{")
	s = append(s, "GrainId: "+fmt.Sprintf("%#v", this.GrainId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TotalsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shared.TotalsResponse{")
	keysForTotals := make([]string, 0, len(this.Totals))
	for k, _ := range this.Totals {
		keysForTotals = append(keysForTotals, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForTotals)
	mapStringForTotals := "map[string]int64{"
	for _, k := range keysForTotals {
		mapStringForTotals += fmt.Sprintf("%#v: %#v,", k, this.Totals[k])
	}
	mapStringForTotals += "}"
	if this.Totals != nil {
		s = append(s, "Totals: "+mapStringForTotals+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtos(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Noop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Noop) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *NumberRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumberRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtos(dAtA, i, uint64(m.Number))
	}
	return i, nil
}

func (m *CountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CountResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtos(dAtA, i, uint64(m.Number))
	}
	return i, nil
}

func (m *RegisterMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.GrainId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtos(dAtA, i, uint64(len(m.GrainId)))
		i += copy(dAtA[i:], m.GrainId)
	}
	return i, nil
}

func (m *TotalsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Totals) > 0 {
		for k, _ := range m.Totals {
			dAtA[i] = 0xa
			i++
			v := m.Totals[k]
			mapSize := 1 + len(k) + sovProtos(uint64(len(k))) + 1 + sovProtos(uint64(v))
			i = encodeVarintProtos(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintProtos(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintProtos(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func encodeVarintProtos(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Noop) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *NumberRequest) Size() (n int) {
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovProtos(uint64(m.Number))
	}
	return n
}

func (m *CountResponse) Size() (n int) {
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovProtos(uint64(m.Number))
	}
	return n
}

func (m *RegisterMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.GrainId)
	if l > 0 {
		n += 1 + l + sovProtos(uint64(l))
	}
	return n
}

func (m *TotalsResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Totals) > 0 {
		for k, v := range m.Totals {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovProtos(uint64(len(k))) + 1 + sovProtos(uint64(v))
			n += mapEntrySize + 1 + sovProtos(uint64(mapEntrySize))
		}
	}
	return n
}

func sovProtos(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtos(x uint64) (n int) {
	return sovProtos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Noop) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Noop{`,
		`}`,
	}, "")
	return s
}
func (this *NumberRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NumberRequest{`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CountResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CountResponse{`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RegisterMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterMessage{`,
		`GrainId:` + fmt.Sprintf("%v", this.GrainId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TotalsResponse) String() string {
	if this == nil {
		return "nil"
	}
	keysForTotals := make([]string, 0, len(this.Totals))
	for k, _ := range this.Totals {
		keysForTotals = append(keysForTotals, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForTotals)
	mapStringForTotals := "map[string]int64{"
	for _, k := range keysForTotals {
		mapStringForTotals += fmt.Sprintf("%v: %v,", k, this.Totals[k])
	}
	mapStringForTotals += "}"
	s := strings.Join([]string{`&TotalsResponse{`,
		`Totals:` + mapStringForTotals + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringProtos(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Noop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Noop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Noop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *NumberRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NumberRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumberRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *CountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *RegisterMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GrainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func (m *TotalsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtos
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TotalsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Totals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Totals == nil {
				m.Totals = make(map[string]int64)
			}
			var mapkey string
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProtos
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProtos
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthProtos
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProtos
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipProtos(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthProtos
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Totals[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtos
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
func skipProtos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtos
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
					return 0, ErrIntOverflowProtos
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtos
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtos
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtos
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtos(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtos = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtos   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos.proto", fileDescriptor_protos_c6fbb75b0b9817be) }

var fileDescriptor_protos_c6fbb75b0b9817be = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x31, 0x0f, 0xd2, 0x40,
	0x18, 0xed, 0x51, 0x2d, 0xf8, 0x01, 0x62, 0x2e, 0x8a, 0xc8, 0x70, 0x21, 0x5d, 0x64, 0x20, 0x0c,
	0x10, 0x13, 0x24, 0x2e, 0x82, 0x86, 0x38, 0xc8, 0x50, 0xd9, 0xcd, 0xd1, 0x7e, 0x41, 0x42, 0xed,
	0xd5, 0xbb, 0xab, 0x09, 0x9b, 0x9b, 0xab, 0x3f, 0xc3, 0x1f, 0xe0, 0xea, 0xee, 0xc8, 0xe8, 0x28,
	0x75, 0x71, 0xe4, 0x27, 0x98, 0x1e, 0x85, 0x58, 0x22, 0x83, 0x4e, 0xed, 0x7b, 0x79, 0xef, 0xfb,
	0x5e, 0xde, 0x77, 0x50, 0x8b, 0xa5, 0xd0, 0x42, 0xf5, 0xcd, 0x87, 0x3a, 0xea, 0x0d, 0x97, 0x18,
	0xb8, 0x0e, 0xdc, 0x98, 0x0b, 0x11, 0xbb, 0x0f, 0xa1, 0x3e, 0x4f, 0xde, 0x2e, 0x51, 0x7a, 0xf8,
	0x2e, 0x41, 0xa5, 0x69, 0x13, 0x9c, 0xc8, 0x10, 0x2d, 0xd2, 0x21, 0x5d, 0xdb, 0xcb, 0x51, 0x26,
	0x9c, 0x8a, 0x24, 0xd2, 0x1e, 0xaa, 0x58, 0x44, 0x0a, 0xaf, 0x0a, 0x7b, 0xd0, 0xf0, 0x70, 0xb5,
	0x56, 0x1a, 0xe5, 0x4b, 0x54, 0x8a, 0xaf, 0x90, 0x3e, 0x80, 0xca, 0x4a, 0xf2, 0x75, 0xf4, 0x7a,
	0x1d, 0x18, 0xf1, 0x2d, 0xaf, 0x6c, 0xf0, 0x8b, 0xc0, 0xfd, 0x48, 0xe0, 0xf6, 0x42, 0x68, 0x1e,
	0xaa, 0xf3, 0xe0, 0x31, 0x38, 0xda, 0x30, 0x2d, 0xd2, 0xb1, 0xbb, 0xd5, 0x81, 0xdb, 0x3f, 0x66,
	0xee, 0x17, 0x75, 0x39, 0x7c, 0x1e, 0x69, 0xb9, 0xf5, 0x72, 0x47, 0xfb, 0x31, 0x54, 0xff, 0xa0,
	0xe9, 0x1d, 0xb0, 0x37, 0xb8, 0xcd, 0x77, 0x66, 0xbf, 0xf4, 0x2e, 0xdc, 0x7c, 0xcf, 0xc3, 0x04,
	0x5b, 0x25, 0x13, 0xfa, 0x08, 0xc6, 0xa5, 0x11, 0x19, 0x7c, 0x21, 0x00, 0x53, 0x1e, 0xfa, 0x49,
	0xc8, 0xb5, 0x90, 0xf4, 0x11, 0xd8, 0x4f, 0x83, 0x80, 0xde, 0x3b, 0x2d, 0x2f, 0xb4, 0xd4, 0x3e,
	0xd3, 0x85, 0x4e, 0x5c, 0x8b, 0x8e, 0xa1, 0xf2, 0x2a, 0x59, 0x6a, 0xc9, 0x7d, 0xfd, 0xcf, 0xde,
	0x21, 0xc0, 0x0c, 0xf5, 0x34, 0x91, 0x12, 0x23, 0x4d, 0x6b, 0x67, 0xb7, 0x10, 0xf1, 0x55, 0xd3,
	0xe0, 0x2b, 0x81, 0xf2, 0x42, 0x72, 0x7f, 0x83, 0x92, 0x8e, 0xa0, 0x7e, 0xaa, 0x7e, 0x96, 0xf5,
	0x4b, 0xef, 0x9f, 0x5c, 0x17, 0x17, 0x69, 0x17, 0x86, 0x9b, 0xd8, 0x8d, 0x67, 0x28, 0xff, 0xcf,
	0xfb, 0x04, 0xe8, 0x44, 0x0a, 0x1e, 0xf8, 0x5c, 0xe9, 0x2c, 0x7f, 0x16, 0x50, 0x5d, 0xc4, 0x6f,
	0xfe, 0xfd, 0x86, 0xae, 0x35, 0xe9, 0xed, 0xf6, 0xcc, 0xfa, 0xbe, 0x67, 0xd6, 0x61, 0xcf, 0xc8,
	0x87, 0x94, 0x91, 0xcf, 0x29, 0x23, 0xdf, 0x52, 0x46, 0x76, 0x29, 0x23, 0x3f, 0x52, 0x46, 0x7e,
	0xa5, 0xcc, 0x3a, 0xa4, 0x8c, 0x7c, 0xfa, 0xc9, 0xac, 0xa5, 0x63, 0x5e, 0xf1, 0xf0, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x73, 0xd6, 0x0c, 0xd5, 0x02, 0x00, 0x00,
}
