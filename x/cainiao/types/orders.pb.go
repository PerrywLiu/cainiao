// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cainiao/orders.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Orders struct {
	Id             uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SendName       string   `protobuf:"bytes,2,opt,name=sendName,proto3" json:"sendName,omitempty"`
	SendAddress    string   `protobuf:"bytes,3,opt,name=sendAddress,proto3" json:"sendAddress,omitempty"`
	SendTel        string   `protobuf:"bytes,4,opt,name=sendTel,proto3" json:"sendTel,omitempty"`
	TargetName     string   `protobuf:"bytes,5,opt,name=targetName,proto3" json:"targetName,omitempty"`
	TargetAddress  string   `protobuf:"bytes,6,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
	TargetTel      string   `protobuf:"bytes,7,opt,name=targetTel,proto3" json:"targetTel,omitempty"`
	State          string   `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	Station        string   `protobuf:"bytes,9,opt,name=station,proto3" json:"station,omitempty"`
	LocationRouter []string `protobuf:"bytes,10,rep,name=locationRouter,proto3" json:"locationRouter,omitempty"`
}

func (m *Orders) Reset()         { *m = Orders{} }
func (m *Orders) String() string { return proto.CompactTextString(m) }
func (*Orders) ProtoMessage()    {}
func (*Orders) Descriptor() ([]byte, []int) {
	return fileDescriptor_07da8e30002bb3a6, []int{0}
}
func (m *Orders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Orders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Orders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Orders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orders.Merge(m, src)
}
func (m *Orders) XXX_Size() int {
	return m.Size()
}
func (m *Orders) XXX_DiscardUnknown() {
	xxx_messageInfo_Orders.DiscardUnknown(m)
}

var xxx_messageInfo_Orders proto.InternalMessageInfo

func (m *Orders) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Orders) GetSendName() string {
	if m != nil {
		return m.SendName
	}
	return ""
}

func (m *Orders) GetSendAddress() string {
	if m != nil {
		return m.SendAddress
	}
	return ""
}

func (m *Orders) GetSendTel() string {
	if m != nil {
		return m.SendTel
	}
	return ""
}

func (m *Orders) GetTargetName() string {
	if m != nil {
		return m.TargetName
	}
	return ""
}

func (m *Orders) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *Orders) GetTargetTel() string {
	if m != nil {
		return m.TargetTel
	}
	return ""
}

func (m *Orders) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Orders) GetStation() string {
	if m != nil {
		return m.Station
	}
	return ""
}

func (m *Orders) GetLocationRouter() []string {
	if m != nil {
		return m.LocationRouter
	}
	return nil
}

func init() {
	proto.RegisterType((*Orders)(nil), "cosmonaut.cainiao.cainiao.Orders")
}

func init() { proto.RegisterFile("cainiao/orders.proto", fileDescriptor_07da8e30002bb3a6) }

var fileDescriptor_07da8e30002bb3a6 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4a, 0xfc, 0x30,
	0x10, 0xc7, 0x37, 0xdd, 0xbf, 0x9d, 0x1f, 0xbf, 0x3d, 0x84, 0x3d, 0x44, 0x91, 0x50, 0x44, 0xa4,
	0x07, 0x69, 0x0f, 0x3e, 0x81, 0x1e, 0x3c, 0x2a, 0x14, 0x4f, 0xde, 0xb2, 0x4d, 0x58, 0x03, 0xdb,
	0x66, 0x49, 0x52, 0xd0, 0xb7, 0xf0, 0x41, 0x7c, 0x10, 0x8f, 0x7b, 0xf4, 0x28, 0xed, 0x8b, 0x48,
	0xa7, 0xdb, 0xba, 0x7a, 0xea, 0x7c, 0x3e, 0xdf, 0x61, 0x3a, 0x64, 0x60, 0x95, 0x0b, 0x5d, 0x6a,
	0x61, 0x52, 0x63, 0xa5, 0xb2, 0x2e, 0xd9, 0x59, 0xe3, 0x0d, 0x3d, 0xc9, 0x8d, 0x2b, 0x4c, 0x29,
	0x2a, 0x9f, 0x1c, 0xf2, 0xfe, 0x7b, 0xfe, 0x1e, 0xc0, 0xec, 0x01, 0x7b, 0xe9, 0x12, 0x02, 0x2d,
	0x19, 0x89, 0x48, 0x3c, 0xc9, 0x02, 0x2d, 0xe9, 0x29, 0x2c, 0x9c, 0x2a, 0xe5, 0xbd, 0x28, 0x14,
	0x0b, 0x22, 0x12, 0x87, 0xd9, 0xc0, 0x34, 0x82, 0x7f, 0x6d, 0x7d, 0x23, 0xa5, 0x55, 0xce, 0xb1,
	0x31, 0xc6, 0xc7, 0x8a, 0x32, 0x98, 0xb7, 0xf8, 0xa8, 0xb6, 0x6c, 0x82, 0x69, 0x8f, 0x94, 0x03,
	0x78, 0x61, 0x37, 0xca, 0xe3, 0xe4, 0x29, 0x86, 0x47, 0x86, 0x5e, 0xc0, 0xff, 0x8e, 0xfa, 0xe9,
	0x33, 0x6c, 0xf9, 0x2d, 0xe9, 0x19, 0x84, 0x9d, 0x68, 0xff, 0x30, 0xc7, 0x8e, 0x1f, 0x41, 0x57,
	0x30, 0x75, 0x5e, 0x78, 0xc5, 0x16, 0x98, 0x74, 0x80, 0x3b, 0x79, 0xe1, 0xb5, 0x29, 0x59, 0x78,
	0xd8, 0xa9, 0x43, 0x7a, 0x09, 0xcb, 0xad, 0xc9, 0xb1, 0xce, 0x4c, 0xe5, 0x95, 0x65, 0x10, 0x8d,
	0xe3, 0x30, 0xfb, 0x63, 0x6f, 0xef, 0x3e, 0x6a, 0x4e, 0xf6, 0x35, 0x27, 0x5f, 0x35, 0x27, 0x6f,
	0x0d, 0x1f, 0xed, 0x1b, 0x3e, 0xfa, 0x6c, 0xf8, 0xe8, 0xe9, 0x6a, 0xa3, 0xfd, 0x73, 0xb5, 0x4e,
	0x72, 0x53, 0xa4, 0xc3, 0x73, 0xa7, 0xfd, 0x39, 0x5e, 0x86, 0xca, 0xbf, 0xee, 0x94, 0x5b, 0xcf,
	0xf0, 0x30, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x17, 0x12, 0x7f, 0xb3, 0xb0, 0x01, 0x00,
	0x00,
}

func (m *Orders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Orders) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Orders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LocationRouter) > 0 {
		for iNdEx := len(m.LocationRouter) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.LocationRouter[iNdEx])
			copy(dAtA[i:], m.LocationRouter[iNdEx])
			i = encodeVarintOrders(dAtA, i, uint64(len(m.LocationRouter[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Station) > 0 {
		i -= len(m.Station)
		copy(dAtA[i:], m.Station)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.Station)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TargetTel) > 0 {
		i -= len(m.TargetTel)
		copy(dAtA[i:], m.TargetTel)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.TargetTel)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TargetAddress) > 0 {
		i -= len(m.TargetAddress)
		copy(dAtA[i:], m.TargetAddress)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.TargetAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TargetName) > 0 {
		i -= len(m.TargetName)
		copy(dAtA[i:], m.TargetName)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.TargetName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SendTel) > 0 {
		i -= len(m.SendTel)
		copy(dAtA[i:], m.SendTel)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.SendTel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SendAddress) > 0 {
		i -= len(m.SendAddress)
		copy(dAtA[i:], m.SendAddress)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.SendAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SendName) > 0 {
		i -= len(m.SendName)
		copy(dAtA[i:], m.SendName)
		i = encodeVarintOrders(dAtA, i, uint64(len(m.SendName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintOrders(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrders(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrders(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Orders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovOrders(uint64(m.Id))
	}
	l = len(m.SendName)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.SendAddress)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.SendTel)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.TargetName)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.TargetAddress)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.TargetTel)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	l = len(m.Station)
	if l > 0 {
		n += 1 + l + sovOrders(uint64(l))
	}
	if len(m.LocationRouter) > 0 {
		for _, s := range m.LocationRouter {
			l = len(s)
			n += 1 + l + sovOrders(uint64(l))
		}
	}
	return n
}

func sovOrders(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrders(x uint64) (n int) {
	return sovOrders(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Orders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrders
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
			return fmt.Errorf("proto: Orders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Orders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendTel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendTel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetTel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetTel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Station", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Station = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocationRouter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrders
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
				return ErrInvalidLengthOrders
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrders
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocationRouter = append(m.LocationRouter, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrders(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrders
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
func skipOrders(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrders
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
					return 0, ErrIntOverflowOrders
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
					return 0, ErrIntOverflowOrders
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
				return 0, ErrInvalidLengthOrders
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrders
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrders
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrders        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrders          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrders = fmt.Errorf("proto: unexpected end of group")
)
