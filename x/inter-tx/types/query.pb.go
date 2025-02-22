// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: intertx/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	_ "github.com/cosmos/interchain-accounts/x/ibc-account/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryIBCAccountFromAddressRequest struct {
	Port    string                                        `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Channel string                                        `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
}

func (m *QueryIBCAccountFromAddressRequest) Reset()         { *m = QueryIBCAccountFromAddressRequest{} }
func (m *QueryIBCAccountFromAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountFromAddressRequest) ProtoMessage()    {}
func (*QueryIBCAccountFromAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b2be0eec8e13d, []int{0}
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountFromAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountFromAddressRequest.Merge(m, src)
}
func (m *QueryIBCAccountFromAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountFromAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountFromAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountFromAddressRequest proto.InternalMessageInfo

type QueryIBCAccountFromAddressResponse struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
}

func (m *QueryIBCAccountFromAddressResponse) Reset()         { *m = QueryIBCAccountFromAddressResponse{} }
func (m *QueryIBCAccountFromAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIBCAccountFromAddressResponse) ProtoMessage()    {}
func (*QueryIBCAccountFromAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b2be0eec8e13d, []int{1}
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIBCAccountFromAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIBCAccountFromAddressResponse.Merge(m, src)
}
func (m *QueryIBCAccountFromAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIBCAccountFromAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIBCAccountFromAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIBCAccountFromAddressResponse proto.InternalMessageInfo

func (m *QueryIBCAccountFromAddressResponse) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIBCAccountFromAddressRequest)(nil), "intertx.QueryIBCAccountFromAddressRequest")
	proto.RegisterType((*QueryIBCAccountFromAddressResponse)(nil), "intertx.QueryIBCAccountFromAddressResponse")
}

func init() { proto.RegisterFile("intertx/query.proto", fileDescriptor_5c0b2be0eec8e13d) }

var fileDescriptor_5c0b2be0eec8e13d = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x3f, 0x4b, 0x3b, 0x41,
	0x10, 0xbd, 0xfd, 0xfd, 0x31, 0xba, 0x58, 0x9d, 0x0a, 0x67, 0x90, 0xbb, 0x98, 0x2a, 0x28, 0x77,
	0x4b, 0x14, 0x2c, 0xec, 0x12, 0x41, 0x10, 0x1b, 0x4d, 0x69, 0xb7, 0xd9, 0x2c, 0x97, 0xc5, 0x64,
	0xe7, 0xb2, 0xbb, 0x27, 0xc9, 0x37, 0xb0, 0xf4, 0x23, 0xa4, 0xf4, 0xa3, 0x58, 0xa6, 0xb4, 0x12,
	0x49, 0x1a, 0x3f, 0x83, 0x95, 0xe4, 0x76, 0x0f, 0x04, 0x45, 0x05, 0xab, 0x99, 0x7b, 0x6f, 0x6e,
	0xde, 0xbc, 0x7d, 0x78, 0x43, 0x48, 0xc3, 0x95, 0x19, 0x93, 0x51, 0xce, 0xd5, 0x24, 0xc9, 0x14,
	0x18, 0xf0, 0x2b, 0x0e, 0xac, 0x6e, 0xa6, 0x90, 0x42, 0x81, 0x91, 0x65, 0x67, 0xe9, 0xea, 0x4e,
	0x0a, 0x90, 0x0e, 0x38, 0xa1, 0x99, 0x20, 0x54, 0x4a, 0x30, 0xd4, 0x08, 0x90, 0xda, 0xb1, 0x91,
	0xe8, 0x32, 0xc2, 0x40, 0x71, 0xc2, 0x06, 0x82, 0x4b, 0x43, 0x6e, 0x9a, 0xae, 0x73, 0x03, 0xdb,
	0xcb, 0x01, 0xca, 0x18, 0xe4, 0xd2, 0x94, 0xd5, 0x52, 0xf5, 0x7b, 0x84, 0x77, 0x2f, 0x97, 0x87,
	0x9c, 0xb5, 0x4f, 0x5a, 0x96, 0x39, 0x55, 0x30, 0x6c, 0xf5, 0x7a, 0x8a, 0x6b, 0xdd, 0xe1, 0xa3,
	0x9c, 0x6b, 0xe3, 0xfb, 0xf8, 0x5f, 0x06, 0xca, 0x04, 0xa8, 0x86, 0x1a, 0x6b, 0x9d, 0xa2, 0xf7,
	0x03, 0x5c, 0x61, 0x7d, 0x2a, 0x25, 0x1f, 0x04, 0x7f, 0x0a, 0xb8, 0xfc, 0xf4, 0xcf, 0x71, 0x85,
	0xda, 0xff, 0x83, 0xbf, 0x35, 0xd4, 0x58, 0x6f, 0x37, 0x5f, 0x9f, 0xa2, 0x38, 0x15, 0xa6, 0x9f,
	0x77, 0x13, 0x06, 0x43, 0xc2, 0x40, 0x0f, 0x41, 0xbb, 0x12, 0xeb, 0xde, 0x35, 0x31, 0x93, 0x8c,
	0xeb, 0xa4, 0xc5, 0x58, 0x29, 0x5c, 0x6e, 0x38, 0x5e, 0xbd, 0x9d, 0x46, 0xde, 0xcb, 0x34, 0xf2,
	0xea, 0x23, 0x5c, 0xff, 0xea, 0x52, 0x9d, 0x81, 0xd4, 0xfc, 0xbd, 0x38, 0xfa, 0xad, 0xf8, 0xc1,
	0x04, 0xff, 0x2f, 0x24, 0xfd, 0x0c, 0x6f, 0x7d, 0x2a, 0xeb, 0xef, 0x25, 0x2e, 0xb9, 0xe4, 0xdb,
	0x57, 0xac, 0xee, 0xff, 0x68, 0xd6, 0xfa, 0x68, 0x5f, 0x3c, 0xcc, 0x43, 0x34, 0x9b, 0x87, 0xe8,
	0x79, 0x1e, 0xa2, 0xbb, 0x45, 0xe8, 0xcd, 0x16, 0xa1, 0xf7, 0xb8, 0x08, 0xbd, 0xab, 0xa3, 0x8f,
	0x66, 0x8a, 0xbd, 0xac, 0x4f, 0x85, 0x8c, 0x5d, 0xbc, 0x9a, 0x8c, 0x2d, 0x1a, 0x9b, 0xb1, 0x35,
	0xd8, 0x5d, 0x29, 0x12, 0x3f, 0x7c, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x17, 0xdb, 0x05, 0xfd, 0x81,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	IBCAccountFromAddress(ctx context.Context, in *QueryIBCAccountFromAddressRequest, opts ...grpc.CallOption) (*QueryIBCAccountFromAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IBCAccountFromAddress(ctx context.Context, in *QueryIBCAccountFromAddressRequest, opts ...grpc.CallOption) (*QueryIBCAccountFromAddressResponse, error) {
	out := new(QueryIBCAccountFromAddressResponse)
	err := c.cc.Invoke(ctx, "/intertx.Query/IBCAccountFromAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	IBCAccountFromAddress(context.Context, *QueryIBCAccountFromAddressRequest) (*QueryIBCAccountFromAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IBCAccountFromAddress(ctx context.Context, req *QueryIBCAccountFromAddressRequest) (*QueryIBCAccountFromAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IBCAccountFromAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IBCAccountFromAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIBCAccountFromAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IBCAccountFromAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intertx.Query/IBCAccountFromAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IBCAccountFromAddress(ctx, req.(*QueryIBCAccountFromAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "intertx.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IBCAccountFromAddress",
			Handler:    _Query_IBCAccountFromAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intertx/query.proto",
}

func (m *QueryIBCAccountFromAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountFromAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountFromAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIBCAccountFromAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIBCAccountFromAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIBCAccountFromAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryIBCAccountFromAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIBCAccountFromAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryIBCAccountFromAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIBCAccountFromAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountFromAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryIBCAccountFromAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIBCAccountFromAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIBCAccountFromAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
