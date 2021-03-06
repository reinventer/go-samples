// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pbgogo/test.proto

package pbgogo

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestObject struct {
	Query                string                `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int64                 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int64                 `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	Flag                 bool                  `protobuf:"varint,4,opt,name=flag,proto3" json:"flag,omitempty"`
	Score                float32               `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	List                 []*TestNested         `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	Map                  map[int64]*TestNested `protobuf:"bytes,7,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Nested               *TestNested           `protobuf:"bytes,8,opt,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TestObject) Reset()         { *m = TestObject{} }
func (m *TestObject) String() string { return proto.CompactTextString(m) }
func (*TestObject) ProtoMessage()    {}
func (*TestObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_960966ae456c4ea5, []int{0}
}
func (m *TestObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestObject.Merge(m, src)
}
func (m *TestObject) XXX_Size() int {
	return m.Size()
}
func (m *TestObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TestObject.DiscardUnknown(m)
}

var xxx_messageInfo_TestObject proto.InternalMessageInfo

func (m *TestObject) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *TestObject) GetPageNumber() int64 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *TestObject) GetResultPerPage() int64 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func (m *TestObject) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *TestObject) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *TestObject) GetList() []*TestNested {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *TestObject) GetMap() map[int64]*TestNested {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *TestObject) GetNested() *TestNested {
	if m != nil {
		return m.Nested
	}
	return nil
}

type TestNested struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestNested) Reset()         { *m = TestNested{} }
func (m *TestNested) String() string { return proto.CompactTextString(m) }
func (*TestNested) ProtoMessage()    {}
func (*TestNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_960966ae456c4ea5, []int{1}
}
func (m *TestNested) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestNested.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestNested.Merge(m, src)
}
func (m *TestNested) XXX_Size() int {
	return m.Size()
}
func (m *TestNested) XXX_DiscardUnknown() {
	xxx_messageInfo_TestNested.DiscardUnknown(m)
}

var xxx_messageInfo_TestNested proto.InternalMessageInfo

func (m *TestNested) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*TestObject)(nil), "pbgogo.TestObject")
	proto.RegisterMapType((map[int64]*TestNested)(nil), "pbgogo.TestObject.MapEntry")
	proto.RegisterType((*TestNested)(nil), "pbgogo.TestNested")
}

func init() { proto.RegisterFile("pbgogo/test.proto", fileDescriptor_960966ae456c4ea5) }

var fileDescriptor_960966ae456c4ea5 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xa4, 0x8d, 0x75, 0x8a, 0x58, 0x07, 0x0f, 0x8b, 0x42, 0x0c, 0x3d, 0x94, 0x20,
	0x18, 0x41, 0x2f, 0xe2, 0x51, 0xf0, 0x22, 0x58, 0xcb, 0xe2, 0xbd, 0x24, 0x75, 0x0c, 0xd5, 0xb4,
	0x59, 0x77, 0x37, 0x42, 0xdf, 0xc4, 0x47, 0xf2, 0xe8, 0x23, 0x48, 0xf4, 0x41, 0x64, 0x77, 0x2d,
	0x78, 0xd0, 0xdb, 0xcc, 0x3f, 0xdf, 0xcc, 0xce, 0xfc, 0x0b, 0xbb, 0xb2, 0x28, 0xeb, 0xb2, 0x3e,
	0x31, 0xa4, 0x4d, 0x26, 0x55, 0x6d, 0x6a, 0x8c, 0xbc, 0x34, 0xfc, 0x0a, 0x00, 0xee, 0x48, 0x9b,
	0xdb, 0xe2, 0x91, 0x66, 0x06, 0xf7, 0xa0, 0xfb, 0xdc, 0x90, 0x5a, 0x71, 0x96, 0xb0, 0x74, 0x4b,
	0xf8, 0x04, 0x0f, 0xa1, 0x2f, 0xf3, 0x92, 0xa6, 0xcb, 0x66, 0x51, 0x90, 0xe2, 0x41, 0xc2, 0xd2,
	0x50, 0x80, 0x95, 0xc6, 0x4e, 0xc1, 0x11, 0xec, 0x28, 0xd2, 0x4d, 0x65, 0xa6, 0x92, 0xd4, 0xd4,
	0x16, 0x78, 0xe8, 0xa0, 0x6d, 0x2f, 0x4f, 0x48, 0x4d, 0xf2, 0x92, 0x10, 0xa1, 0xf3, 0x50, 0xe5,
	0x25, 0xef, 0x24, 0x2c, 0xed, 0x09, 0x17, 0xdb, 0x27, 0xf5, 0xac, 0x56, 0xc4, 0xbb, 0x09, 0x4b,
	0x03, 0xe1, 0x13, 0x1c, 0x41, 0xa7, 0x9a, 0x6b, 0xc3, 0xa3, 0x24, 0x4c, 0xfb, 0xa7, 0x98, 0xf9,
	0x75, 0x33, 0xbb, 0xea, 0x98, 0xb4, 0xa1, 0x7b, 0xe1, 0xea, 0x78, 0x0c, 0xe1, 0x22, 0x97, 0x7c,
	0xd3, 0x61, 0x07, 0xbf, 0x31, 0x7f, 0x51, 0x76, 0x93, 0xcb, 0xab, 0xa5, 0x51, 0x2b, 0x61, 0x39,
	0x3c, 0x82, 0x68, 0xe9, 0xda, 0x79, 0x2f, 0x61, 0xff, 0x0c, 0xfe, 0x21, 0xf6, 0xaf, 0xa1, 0xb7,
	0x6e, 0xc6, 0x01, 0x84, 0x4f, 0xe4, 0x5d, 0x09, 0x85, 0x0d, 0x31, 0x85, 0xee, 0x4b, 0x5e, 0x35,
	0xe4, 0xdc, 0xf8, 0x7b, 0x90, 0x07, 0x2e, 0x82, 0x73, 0x36, 0x1c, 0x7a, 0x97, 0x7d, 0xc1, 0x9e,
	0x6c, 0xe6, 0xa6, 0xa2, 0xb5, 0xcb, 0x2e, 0xb9, 0x1c, 0xbc, 0xb5, 0x31, 0x7b, 0x6f, 0x63, 0xf6,
	0xd1, 0xc6, 0xec, 0xf5, 0x33, 0xde, 0x28, 0x22, 0xf7, 0x57, 0x67, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x72, 0x90, 0x02, 0x6a, 0xc0, 0x01, 0x00, 0x00,
}

func (m *TestObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Nested != nil {
		{
			size, err := m.Nested.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Map) > 0 {
		for k := range m.Map {
			v := m.Map[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTest(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintTest(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintTest(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTest(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Score != 0 {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Score))))
		i--
		dAtA[i] = 0x2d
	}
	if m.Flag {
		i--
		if m.Flag {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.ResultPerPage != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.ResultPerPage))
		i--
		dAtA[i] = 0x18
	}
	if m.PageNumber != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.PageNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestNested) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestNested) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestNested) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.PageNumber != 0 {
		n += 1 + sovTest(uint64(m.PageNumber))
	}
	if m.ResultPerPage != 0 {
		n += 1 + sovTest(uint64(m.ResultPerPage))
	}
	if m.Flag {
		n += 2
	}
	if m.Score != 0 {
		n += 5
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovTest(uint64(l))
		}
	}
	if len(m.Map) > 0 {
		for k, v := range m.Map {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTest(uint64(l))
			}
			mapEntrySize := 1 + sovTest(uint64(k)) + l
			n += mapEntrySize + 1 + sovTest(uint64(mapEntrySize))
		}
	}
	if m.Nested != nil {
		l = m.Nested.Size()
		n += 1 + l + sovTest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TestNested) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: TestObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageNumber", wireType)
			}
			m.PageNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultPerPage", wireType)
			}
			m.ResultPerPage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResultPerPage |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
			m.Flag = bool(v != 0)
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Score = float32(math.Float32frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &TestNested{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Map", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Map == nil {
				m.Map = make(map[int64]*TestNested)
			}
			var mapkey int64
			var mapvalue *TestNested
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTest
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTest
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTest
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &TestNested{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTest(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTest
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Map[mapkey] = mapvalue
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nested", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nested == nil {
				m.Nested = &TestNested{}
			}
			if err := m.Nested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestNested) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: TestNested: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestNested: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
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
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTest
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
				next, err := skipTest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTest
				}
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
	ErrInvalidLengthTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest   = fmt.Errorf("proto: integer overflow")
)
