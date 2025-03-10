// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/client/v2/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryCounterpartyInfoRequest is the request type for the Query/CounterpartyInfo RPC
// method
type QueryCounterpartyInfoRequest struct {
	// client state unique identifier
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *QueryCounterpartyInfoRequest) Reset()         { *m = QueryCounterpartyInfoRequest{} }
func (m *QueryCounterpartyInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCounterpartyInfoRequest) ProtoMessage()    {}
func (*QueryCounterpartyInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2f0dd58f99aaa6, []int{0}
}
func (m *QueryCounterpartyInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterpartyInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterpartyInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterpartyInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterpartyInfoRequest.Merge(m, src)
}
func (m *QueryCounterpartyInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterpartyInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterpartyInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterpartyInfoRequest proto.InternalMessageInfo

func (m *QueryCounterpartyInfoRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

// QueryCounterpartyInfoResponse is the response type for the
// Query/CounterpartyInfo RPC method.
type QueryCounterpartyInfoResponse struct {
	CounterpartyInfo *CounterpartyInfo `protobuf:"bytes,1,opt,name=counterparty_info,json=counterpartyInfo,proto3" json:"counterparty_info,omitempty"`
}

func (m *QueryCounterpartyInfoResponse) Reset()         { *m = QueryCounterpartyInfoResponse{} }
func (m *QueryCounterpartyInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCounterpartyInfoResponse) ProtoMessage()    {}
func (*QueryCounterpartyInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2f0dd58f99aaa6, []int{1}
}
func (m *QueryCounterpartyInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCounterpartyInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCounterpartyInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCounterpartyInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCounterpartyInfoResponse.Merge(m, src)
}
func (m *QueryCounterpartyInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCounterpartyInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCounterpartyInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCounterpartyInfoResponse proto.InternalMessageInfo

func (m *QueryCounterpartyInfoResponse) GetCounterpartyInfo() *CounterpartyInfo {
	if m != nil {
		return m.CounterpartyInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCounterpartyInfoRequest)(nil), "ibc.core.client.v2.QueryCounterpartyInfoRequest")
	proto.RegisterType((*QueryCounterpartyInfoResponse)(nil), "ibc.core.client.v2.QueryCounterpartyInfoResponse")
}

func init() { proto.RegisterFile("ibc/core/client/v2/query.proto", fileDescriptor_ce2f0dd58f99aaa6) }

var fileDescriptor_ce2f0dd58f99aaa6 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4c, 0x4a, 0xd6,
	0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0xd2, 0x2f,
	0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xca, 0x4c, 0x4a, 0xd6,
	0x03, 0xc9, 0xeb, 0x41, 0xe4, 0xf5, 0xca, 0x8c, 0xa4, 0x54, 0xb1, 0xe8, 0x49, 0xce, 0x2f, 0xcd,
	0x2b, 0x49, 0x2d, 0x2a, 0x48, 0x2c, 0x2a, 0x81, 0x6a, 0x95, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0x86, 0xc8, 0x2a, 0x59, 0x73, 0xc9, 0x04, 0x82, 0xec, 0x71, 0x46, 0xd2, 0xe8, 0x99, 0x97,
	0x96, 0x1f, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xcd, 0xc5, 0x09, 0x31, 0x3d, 0x3e,
	0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x03, 0x22, 0xe0, 0x99, 0xa2, 0x54, 0xc4,
	0x25, 0x8b, 0x43, 0x73, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x50, 0x20, 0x97, 0x20, 0xb2, 0x8b,
	0xe2, 0x33, 0xf3, 0xd2, 0xf2, 0xc1, 0xa6, 0x70, 0x1b, 0xa9, 0xe8, 0x61, 0x7a, 0x49, 0x0f, 0xc3,
	0x20, 0x81, 0x64, 0x34, 0x11, 0xa3, 0x3d, 0x8c, 0x5c, 0xac, 0x60, 0x4b, 0x85, 0x36, 0x31, 0x72,
	0x09, 0xa0, 0x6b, 0x10, 0x32, 0xc0, 0x66, 0x2c, 0x3e, 0x1f, 0x4a, 0x19, 0x92, 0xa0, 0x03, 0xe2,
	0x2d, 0x25, 0xcb, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x19, 0x0b, 0x19, 0xea, 0x13, 0x88, 0x02, 0xb0,
	0x87, 0xf5, 0xab, 0xe1, 0x21, 0x58, 0xeb, 0x14, 0x76, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x36, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9,
	0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x20, 0xd3, 0x75, 0xd3, 0xf3, 0xf5, 0xcb, 0x0c, 0x0d, 0xf4, 0x73,
	0xf3, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0x21, 0x96, 0x19, 0x18, 0xe9, 0x22, 0xec, 0x2b, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x47, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x43, 0x0f,
	0xb5, 0x49, 0x02, 0x00, 0x00,
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
	// CounterpartyInfo queries an IBC light counter party info.
	CounterpartyInfo(ctx context.Context, in *QueryCounterpartyInfoRequest, opts ...grpc.CallOption) (*QueryCounterpartyInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CounterpartyInfo(ctx context.Context, in *QueryCounterpartyInfoRequest, opts ...grpc.CallOption) (*QueryCounterpartyInfoResponse, error) {
	out := new(QueryCounterpartyInfoResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.client.v2.Query/CounterpartyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CounterpartyInfo queries an IBC light counter party info.
	CounterpartyInfo(context.Context, *QueryCounterpartyInfoRequest) (*QueryCounterpartyInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CounterpartyInfo(ctx context.Context, req *QueryCounterpartyInfoRequest) (*QueryCounterpartyInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CounterpartyInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CounterpartyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCounterpartyInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CounterpartyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.client.v2.Query/CounterpartyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CounterpartyInfo(ctx, req.(*QueryCounterpartyInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.client.v2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CounterpartyInfo",
			Handler:    _Query_CounterpartyInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/client/v2/query.proto",
}

func (m *QueryCounterpartyInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterpartyInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterpartyInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCounterpartyInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCounterpartyInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCounterpartyInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CounterpartyInfo != nil {
		{
			size, err := m.CounterpartyInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
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
func (m *QueryCounterpartyInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCounterpartyInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CounterpartyInfo != nil {
		l = m.CounterpartyInfo.Size()
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
func (m *QueryCounterpartyInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCounterpartyInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterpartyInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
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
			m.ClientId = string(dAtA[iNdEx:postIndex])
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
func (m *QueryCounterpartyInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCounterpartyInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCounterpartyInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CounterpartyInfo == nil {
				m.CounterpartyInfo = &CounterpartyInfo{}
			}
			if err := m.CounterpartyInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
