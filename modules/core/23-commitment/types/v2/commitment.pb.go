// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/commitment/v2/commitment.proto

package v2

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// MerklePath is the path used to verify commitment proofs, which can be an
// arbitrary structured object (defined by a commitment type).
// ICS-23 verification supports membership proofs for nested merkle trees.
// The ICS-24 standard provable keys MUST be stored in the lowest level tree with an optional prefix.
// The IC24 provable tree may then be stored in a higher level tree(s) that hash up to the root hash
// stored in the consensus state of the client.
// Each element of the path represents the key of a merkle tree from the root to the leaf.
// The elements of the path before the final element must be the path to the tree that contains
// the ICS24 provable store. Thus, it should remain constant for all ICS24 proofs.
// The final element of the path is the key of the leaf in the ICS24 provable store,
// Thus IBC core will append the ICS24 path to the final element of the MerklePath
// stored in the counterparty to create the full path to the leaf for proof verification.
// Examples:
// Cosmos SDK:
// The Cosmos SDK commits to a multi-tree where each store is an IAVL tree and all store hashes
// are hashed in a simple merkle tree to get the final root hash. Thus, the MerklePath in the counterparty
// MerklePrefix has the following structure: ["ibc", ""]
// The core IBC handler will append the ICS24 path to the final element of the MerklePath
// like so: ["ibc", "{packetCommitmentPath}"] which will then be used for final verification.
// Ethereum:
// The Ethereum client commits to a single Patricia merkle trie. The ICS24 provable store is managed
// by the smart contract state. Each smart contract has a specific prefix reserved within the global trie.
// Thus the MerklePath in the counterparty is the prefix to the smart contract state in the global trie.
// Since there is only one tree in the commitment structure of ethereum the MerklePath in the counterparty
// MerklePrefix has the following structure: ["IBCCoreContractAddressStoragePrefix"]
// The core IBC handler will append the ICS24 path to the final element of the MerklePath
// like so: ["IBCCoreContractAddressStoragePrefix{packetCommitmentPath}"] which will then be used for final
// verification. Thus the MerklePath in the counterparty MerklePrefix is the nested key path from the root hash of the
// consensus state down to the ICS24 provable store. The IBC handler retrieves the counterparty key path to the ICS24
// provable store from the MerklePath and appends the ICS24 path to get the final key path to the value being verified
// by the client against the root hash in the client's consensus state.
type MerklePath struct {
	KeyPath [][]byte `protobuf:"bytes,1,rep,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
}

func (m *MerklePath) Reset()         { *m = MerklePath{} }
func (m *MerklePath) String() string { return proto.CompactTextString(m) }
func (*MerklePath) ProtoMessage()    {}
func (*MerklePath) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f65a9eb5e4ee5fc, []int{0}
}
func (m *MerklePath) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerklePath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerklePath.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MerklePath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerklePath.Merge(m, src)
}
func (m *MerklePath) XXX_Size() int {
	return m.Size()
}
func (m *MerklePath) XXX_DiscardUnknown() {
	xxx_messageInfo_MerklePath.DiscardUnknown(m)
}

var xxx_messageInfo_MerklePath proto.InternalMessageInfo

func (m *MerklePath) GetKeyPath() [][]byte {
	if m != nil {
		return m.KeyPath
	}
	return nil
}

func init() {
	proto.RegisterType((*MerklePath)(nil), "ibc.core.commitment.v2.MerklePath")
}

func init() {
	proto.RegisterFile("ibc/core/commitment/v2/commitment.proto", fileDescriptor_8f65a9eb5e4ee5fc)
}

var fileDescriptor_8f65a9eb5e4ee5fc = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0x4c, 0x4a, 0xd6,
	0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0x2c, 0xc9, 0x4d, 0xcd, 0x2b, 0xd1,
	0x2f, 0x33, 0x42, 0xe2, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x65, 0x26, 0x25, 0xeb,
	0x81, 0x14, 0xea, 0x21, 0x49, 0x95, 0x19, 0x29, 0xa9, 0x73, 0x71, 0xf9, 0xa6, 0x16, 0x65, 0xe7,
	0xa4, 0x06, 0x24, 0x96, 0x64, 0x08, 0x49, 0x72, 0x71, 0x64, 0xa7, 0x56, 0xc6, 0x17, 0x24, 0x96,
	0x64, 0x48, 0x30, 0x2a, 0x30, 0x6b, 0xf0, 0x04, 0xb1, 0x67, 0xa7, 0x56, 0x82, 0xa4, 0x9c, 0xa2,
	0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x21, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x09, 0x64, 0xb0, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x66, 0x52, 0xb2, 0x6e,
	0x7a, 0xbe, 0x7e, 0x99, 0xa1, 0x81, 0x7e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0x31, 0xc4, 0x91,
	0x46, 0xc6, 0xba, 0x48, 0xee, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x33, 0x4a, 0x62, 0x03,
	0xbb, 0xd1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x01, 0x78, 0xcd, 0xbf, 0xce, 0x00, 0x00, 0x00,
}

func (m *MerklePath) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerklePath) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerklePath) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyPath) > 0 {
		for iNdEx := len(m.KeyPath) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.KeyPath[iNdEx])
			copy(dAtA[i:], m.KeyPath[iNdEx])
			i = encodeVarintCommitment(dAtA, i, uint64(len(m.KeyPath[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommitment(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommitment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MerklePath) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.KeyPath) > 0 {
		for _, b := range m.KeyPath {
			l = len(b)
			n += 1 + l + sovCommitment(uint64(l))
		}
	}
	return n
}

func sovCommitment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommitment(x uint64) (n int) {
	return sovCommitment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MerklePath) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommitment
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
			return fmt.Errorf("proto: MerklePath: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerklePath: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyPath", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommitment
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
				return ErrInvalidLengthCommitment
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommitment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyPath = append(m.KeyPath, make([]byte, postIndex-iNdEx))
			copy(m.KeyPath[len(m.KeyPath)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommitment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommitment
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
func skipCommitment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommitment
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
					return 0, ErrIntOverflowCommitment
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
					return 0, ErrIntOverflowCommitment
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
				return 0, ErrInvalidLengthCommitment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommitment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommitment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommitment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommitment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommitment = fmt.Errorf("proto: unexpected end of group")
)
