// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/client.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import strconv "strconv"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// ReviewingState enum defines all the possible admin reviewing states that a
// request can be at. For example a third-party client creation request.
type ReviewingState int32

const (
	// Denotes that the request is pending to review.
	STATE_PENDING ReviewingState = 0
	// Denotes that the request has been reviewed and approved by an admin.
	STATE_APPROVED ReviewingState = 1
	// Denotes that the request has been reviewed and rejected by an admin.
	STATE_REJECTED ReviewingState = 2
)

var ReviewingState_name = map[int32]string{
	0: "STATE_PENDING",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
}
var ReviewingState_value = map[string]int32{
	"STATE_PENDING":  0,
	"STATE_APPROVED": 1,
	"STATE_REJECTED": 2,
}

func (ReviewingState) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

// Client is the message that defines a third-party client on the network.
type Client struct {
	// id is the client's unique and immutable ID.
	ClientIdentifiers `protobuf:"bytes,1,opt,name=id,embedded=id" json:""`
	// description is the description of the client.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// secret is the secret used to prove the client identity.
	// This is a read-only field.
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	// redirect_uri is the OAuth redirect URI of the client.
	RedirectURI string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// state denotes the reviewing state of the client by admin.
	// This field can only be modified by admins.
	State ReviewingState `protobuf:"varint,5,opt,name=state,proto3,enum=ttn.v3.ReviewingState" json:"state,omitempty"`
	// official_labeled denotes if a client has been labeled as an official
	// third-party client by the tenant admin.
	// This field can only be modified by admins.
	OfficialLabeled bool `protobuf:"varint,6,opt,name=official_labeled,json=officialLabeled,proto3" json:"official_labeled,omitempty"`
	// grants denotes which OAuth2 flows can the client use to get a token.
	// This field can only be modified by admins.
	Grants []GrantType `protobuf:"varint,7,rep,packed,name=grants,enum=ttn.v3.GrantType" json:"grants,omitempty"`
	// Rights denotes what rights the client will have access to.
	Rights []Right `protobuf:"varint,8,rep,packed,name=rights,enum=ttn.v3.Right" json:"rights,omitempty"`
	// creator is the user ID of the user who created the client.
	// This is an immutable and read-only field.
	Creator UserIdentifiers `protobuf:"bytes,9,opt,name=creator" json:"creator"`
	// created_at denotes when the client was created.
	// This a read-only field.
	CreatedAt time.Time `protobuf:"bytes,10,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	// updated_at is the last time the client was updated.
	// This is a read-only field.
	UpdatedAt time.Time `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

func (m *Client) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Client) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Client) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *Client) GetState() ReviewingState {
	if m != nil {
		return m.State
	}
	return STATE_PENDING
}

func (m *Client) GetOfficialLabeled() bool {
	if m != nil {
		return m.OfficialLabeled
	}
	return false
}

func (m *Client) GetGrants() []GrantType {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *Client) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *Client) GetCreator() UserIdentifiers {
	if m != nil {
		return m.Creator
	}
	return UserIdentifiers{}
}

func (m *Client) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Client) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Client)(nil), "ttn.v3.Client")
	golang_proto.RegisterType((*Client)(nil), "ttn.v3.Client")
	proto.RegisterEnum("ttn.v3.ReviewingState", ReviewingState_name, ReviewingState_value)
	golang_proto.RegisterEnum("ttn.v3.ReviewingState", ReviewingState_name, ReviewingState_value)
}
func (x ReviewingState) String() string {
	s, ok := ReviewingState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Client) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Client)
	if !ok {
		that2, ok := that.(Client)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Client")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Client but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Client but is not nil && this == nil")
	}
	if !this.ClientIdentifiers.Equal(&that1.ClientIdentifiers) {
		return fmt.Errorf("ClientIdentifiers this(%v) Not Equal that(%v)", this.ClientIdentifiers, that1.ClientIdentifiers)
	}
	if this.Description != that1.Description {
		return fmt.Errorf("Description this(%v) Not Equal that(%v)", this.Description, that1.Description)
	}
	if this.Secret != that1.Secret {
		return fmt.Errorf("Secret this(%v) Not Equal that(%v)", this.Secret, that1.Secret)
	}
	if this.RedirectURI != that1.RedirectURI {
		return fmt.Errorf("RedirectURI this(%v) Not Equal that(%v)", this.RedirectURI, that1.RedirectURI)
	}
	if this.State != that1.State {
		return fmt.Errorf("State this(%v) Not Equal that(%v)", this.State, that1.State)
	}
	if this.OfficialLabeled != that1.OfficialLabeled {
		return fmt.Errorf("OfficialLabeled this(%v) Not Equal that(%v)", this.OfficialLabeled, that1.OfficialLabeled)
	}
	if len(this.Grants) != len(that1.Grants) {
		return fmt.Errorf("Grants this(%v) Not Equal that(%v)", len(this.Grants), len(that1.Grants))
	}
	for i := range this.Grants {
		if this.Grants[i] != that1.Grants[i] {
			return fmt.Errorf("Grants this[%v](%v) Not Equal that[%v](%v)", i, this.Grants[i], i, that1.Grants[i])
		}
	}
	if len(this.Rights) != len(that1.Rights) {
		return fmt.Errorf("Rights this(%v) Not Equal that(%v)", len(this.Rights), len(that1.Rights))
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return fmt.Errorf("Rights this[%v](%v) Not Equal that[%v](%v)", i, this.Rights[i], i, that1.Rights[i])
		}
	}
	if !this.Creator.Equal(&that1.Creator) {
		return fmt.Errorf("Creator this(%v) Not Equal that(%v)", this.Creator, that1.Creator)
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return fmt.Errorf("CreatedAt this(%v) Not Equal that(%v)", this.CreatedAt, that1.CreatedAt)
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return fmt.Errorf("UpdatedAt this(%v) Not Equal that(%v)", this.UpdatedAt, that1.UpdatedAt)
	}
	return nil
}
func (this *Client) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Client)
	if !ok {
		that2, ok := that.(Client)
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
	if !this.ClientIdentifiers.Equal(&that1.ClientIdentifiers) {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Secret != that1.Secret {
		return false
	}
	if this.RedirectURI != that1.RedirectURI {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if this.OfficialLabeled != that1.OfficialLabeled {
		return false
	}
	if len(this.Grants) != len(that1.Grants) {
		return false
	}
	for i := range this.Grants {
		if this.Grants[i] != that1.Grants[i] {
			return false
		}
	}
	if len(this.Rights) != len(that1.Rights) {
		return false
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return false
		}
	}
	if !this.Creator.Equal(&that1.Creator) {
		return false
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	return true
}
func (m *Client) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Client) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintClient(dAtA, i, uint64(m.ClientIdentifiers.Size()))
	n1, err := m.ClientIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Secret) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Secret)))
		i += copy(dAtA[i:], m.Secret)
	}
	if len(m.RedirectURI) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.RedirectURI)))
		i += copy(dAtA[i:], m.RedirectURI)
	}
	if m.State != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.State))
	}
	if m.OfficialLabeled {
		dAtA[i] = 0x30
		i++
		if m.OfficialLabeled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Grants) > 0 {
		dAtA3 := make([]byte, len(m.Grants)*10)
		var j2 int
		for _, num := range m.Grants {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintClient(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.Rights) > 0 {
		dAtA5 := make([]byte, len(m.Rights)*10)
		var j4 int
		for _, num := range m.Rights {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x42
		i++
		i = encodeVarintClient(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintClient(dAtA, i, uint64(m.Creator.Size()))
	n6, err := m.Creator.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x52
	i++
	i = encodeVarintClient(dAtA, i, uint64(types.SizeOfStdTime(m.CreatedAt)))
	n7, err := types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x5a
	i++
	i = encodeVarintClient(dAtA, i, uint64(types.SizeOfStdTime(m.UpdatedAt)))
	n8, err := types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedClient(r randyClient, easy bool) *Client {
	this := &Client{}
	v1 := NewPopulatedClientIdentifiers(r, easy)
	this.ClientIdentifiers = *v1
	this.Description = randStringClient(r)
	this.Secret = randStringClient(r)
	this.RedirectURI = randStringClient(r)
	this.State = ReviewingState([]int32{0, 1, 2}[r.Intn(3)])
	this.OfficialLabeled = bool(r.Intn(2) == 0)
	v2 := r.Intn(10)
	this.Grants = make([]GrantType, v2)
	for i := 0; i < v2; i++ {
		this.Grants[i] = GrantType([]int32{0, 1, 2}[r.Intn(3)])
	}
	v3 := r.Intn(10)
	this.Rights = make([]Right, v3)
	for i := 0; i < v3; i++ {
		this.Rights[i] = Right([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 17, 18, 31, 32, 33, 34, 35, 36, 37, 38, 39, 51, 52, 53, 54, 55, 56, 57, 58, 71, 72, 73, 74, 75, 76, 77, 78, 79}[r.Intn(40)])
	}
	v4 := NewPopulatedUserIdentifiers(r, easy)
	this.Creator = *v4
	v5 := types.NewPopulatedStdTime(r, easy)
	this.CreatedAt = *v5
	v6 := types.NewPopulatedStdTime(r, easy)
	this.UpdatedAt = *v6
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyClient interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneClient(r randyClient) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringClient(r randyClient) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneClient(r)
	}
	return string(tmps)
}
func randUnrecognizedClient(r randyClient, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldClient(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldClient(dAtA []byte, r randyClient, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateClient(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateClient(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateClient(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Client) Size() (n int) {
	var l int
	_ = l
	l = m.ClientIdentifiers.Size()
	n += 1 + l + sovClient(uint64(l))
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.RedirectURI)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovClient(uint64(m.State))
	}
	if m.OfficialLabeled {
		n += 2
	}
	if len(m.Grants) > 0 {
		l = 0
		for _, e := range m.Grants {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	if len(m.Rights) > 0 {
		l = 0
		for _, e := range m.Rights {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	l = m.Creator.Size()
	n += 1 + l + sovClient(uint64(l))
	l = types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovClient(uint64(l))
	l = types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovClient(uint64(l))
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *Client) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: Client: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Client: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedirectURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedirectURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ReviewingState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfficialLabeled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OfficialLabeled = bool(v != 0)
		case 7:
			if wireType == 0 {
				var v GrantType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (GrantType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Grants = append(m.Grants, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v GrantType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (GrantType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Grants = append(m.Grants, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
			}
		case 8:
			if wireType == 0 {
				var v Right
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Right(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rights = append(m.Rights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Right
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Right(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rights = append(m.Rights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rights", wireType)
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Creator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
				return 0, ErrInvalidLengthClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClient
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
				next, err := skipClient(dAtA[start:])
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
	ErrInvalidLengthClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/client.proto", fileDescriptorClient)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/client.proto", fileDescriptorClient)
}

var fileDescriptorClient = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x4c, 0xdb, 0x4e,
	0x14, 0xbf, 0x0b, 0x10, 0xe0, 0x02, 0x21, 0x78, 0xe0, 0xef, 0x7f, 0x86, 0x97, 0xa8, 0x52, 0xa5,
	0x50, 0xb5, 0x4e, 0x15, 0x54, 0x75, 0x4e, 0x20, 0x42, 0x54, 0x15, 0xa5, 0x47, 0xe8, 0xd0, 0x25,
	0x72, 0xe2, 0x8b, 0x73, 0x22, 0xd8, 0xd6, 0xf9, 0x02, 0xea, 0xc6, 0xc8, 0xc8, 0xd8, 0xad, 0x95,
	0xba, 0x30, 0x32, 0x32, 0x32, 0x32, 0x32, 0x32, 0xa5, 0xe4, 0xbc, 0x20, 0x26, 0x46, 0xc6, 0xca,
	0x5f, 0x85, 0x4e, 0xa5, 0x93, 0xef, 0xfd, 0xbe, 0xf4, 0xf4, 0x7b, 0x32, 0x79, 0x6d, 0x73, 0xd9,
	0x1f, 0x76, 0x8c, 0xae, 0xbb, 0x57, 0x6d, 0xf5, 0x59, 0xab, 0xcf, 0x1d, 0xdb, 0xdf, 0x64, 0xf2,
	0xc0, 0x15, 0xbb, 0x55, 0x29, 0x9d, 0xaa, 0xe9, 0xf1, 0x6a, 0x77, 0xc0, 0x99, 0x23, 0x0d, 0x4f,
	0xb8, 0xd2, 0xd5, 0xb2, 0x52, 0x3a, 0xc6, 0xfe, 0x4a, 0xd1, 0x78, 0x8a, 0xd3, 0x1c, 0xca, 0x7e,
	0xec, 0x2b, 0xbe, 0x79, 0x8a, 0x9e, 0x5b, 0xcc, 0x91, 0xbc, 0xc7, 0x99, 0xf0, 0x13, 0xdb, 0x93,
	0x16, 0x14, 0xdc, 0xee, 0xcb, 0xd4, 0xf1, 0xea, 0x91, 0xc3, 0x76, 0x6d, 0xb7, 0x1a, 0xc1, 0x9d,
	0x61, 0x2f, 0x9a, 0xa2, 0x21, 0x7a, 0x25, 0xf2, 0x92, 0xed, 0xba, 0xf6, 0x80, 0x3d, 0xa8, 0x24,
	0xdf, 0x63, 0xbe, 0x34, 0xf7, 0xbc, 0x58, 0xf0, 0xec, 0x64, 0x92, 0x64, 0x57, 0xa3, 0x06, 0xb4,
	0xb7, 0x24, 0xc3, 0x2d, 0x1d, 0x97, 0x71, 0x25, 0x57, 0xfb, 0xdf, 0x88, 0x8b, 0x30, 0x62, 0x6e,
	0xe3, 0x61, 0xf3, 0xc6, 0xdc, 0xc5, 0xa8, 0x84, 0x2e, 0x47, 0x25, 0x7c, 0x3b, 0x2a, 0x21, 0x9a,
	0xe1, 0x96, 0x56, 0x26, 0x39, 0x8b, 0xf9, 0x5d, 0xc1, 0x3d, 0xc9, 0x5d, 0x47, 0xcf, 0x94, 0x71,
	0x65, 0x96, 0x3e, 0x86, 0xb4, 0x25, 0x92, 0xf5, 0x59, 0x57, 0x30, 0xa9, 0x4f, 0x44, 0x64, 0x32,
	0x69, 0x35, 0x32, 0x27, 0x98, 0xc5, 0x05, 0xeb, 0xca, 0xf6, 0x50, 0x70, 0x7d, 0x32, 0x64, 0x1b,
	0x0b, 0x6a, 0x54, 0xca, 0xd1, 0x04, 0xdf, 0xa1, 0x1b, 0x34, 0x97, 0x8a, 0x76, 0x04, 0xd7, 0x5e,
	0x92, 0x29, 0x5f, 0x9a, 0x92, 0xe9, 0x53, 0x65, 0x5c, 0xc9, 0xd7, 0x96, 0xd2, 0x4d, 0x29, 0xdb,
	0xe7, 0xec, 0x80, 0x3b, 0xf6, 0x76, 0xc8, 0xd2, 0x58, 0xa4, 0x2d, 0x93, 0x82, 0xdb, 0xeb, 0xf1,
	0x2e, 0x37, 0x07, 0xed, 0x81, 0xd9, 0x61, 0x03, 0x66, 0xe9, 0xd9, 0x32, 0xae, 0xcc, 0xd0, 0x85,
	0x14, 0x7f, 0x1f, 0xc3, 0xda, 0x32, 0xc9, 0xda, 0xc2, 0x74, 0xa4, 0xaf, 0x4f, 0x97, 0x27, 0x2a,
	0xf9, 0xda, 0x62, 0x9a, 0xbc, 0x1e, 0xa2, 0xad, 0x2f, 0x1e, 0xa3, 0x89, 0x40, 0x7b, 0x4e, 0xb2,
	0xf1, 0x55, 0xf4, 0x99, 0x48, 0x3a, 0xff, 0x7b, 0x89, 0x10, 0xa5, 0x09, 0xa9, 0x35, 0xc8, 0x74,
	0x57, 0x30, 0x53, 0xba, 0x42, 0x9f, 0x8d, 0x6a, 0xfd, 0x2f, 0xd5, 0xed, 0xf8, 0x4c, 0x3c, 0x2e,
	0x75, 0x21, 0x2c, 0xf5, 0x76, 0x54, 0x4a, 0xf5, 0x34, 0x7d, 0x68, 0xab, 0x84, 0x44, 0x4f, 0x66,
	0xb5, 0x4d, 0xa9, 0x93, 0x28, 0xa6, 0x68, 0xc4, 0x67, 0x35, 0xd2, 0xb3, 0x1a, 0xad, 0xf4, 0xac,
	0x8d, 0x99, 0x30, 0xe9, 0xf8, 0x67, 0x09, 0xd3, 0xd9, 0xc4, 0x57, 0x97, 0x61, 0xc8, 0xd0, 0xb3,
	0xd2, 0x90, 0xdc, 0xbf, 0x84, 0x24, 0xbe, 0xba, 0x7c, 0xf1, 0x91, 0xe4, 0xff, 0xec, 0x58, 0x5b,
	0x24, 0xf3, 0xdb, 0xad, 0x7a, 0xab, 0xd9, 0xde, 0x6a, 0x6e, 0xae, 0x6d, 0x6c, 0xae, 0x17, 0x90,
	0xa6, 0x91, 0x7c, 0x0c, 0xd5, 0xb7, 0xb6, 0xe8, 0x87, 0x4f, 0xcd, 0xb5, 0x02, 0x7e, 0xc0, 0x68,
	0xf3, 0x5d, 0x73, 0xb5, 0xd5, 0x5c, 0x2b, 0x64, 0x8a, 0x93, 0x47, 0x3f, 0x00, 0x35, 0xbe, 0xe1,
	0x8b, 0x31, 0xe0, 0xcb, 0x31, 0xe0, 0xab, 0x31, 0xe0, 0xeb, 0x31, 0xe0, 0x9b, 0x31, 0xa0, 0xbb,
	0x31, 0xa0, 0xfb, 0x31, 0xe0, 0x43, 0x05, 0xe8, 0x48, 0x01, 0x3a, 0x51, 0x80, 0x4f, 0x15, 0xa0,
	0x33, 0x05, 0xf8, 0x5c, 0x01, 0xbe, 0x50, 0x80, 0x2f, 0x15, 0xe0, 0x2b, 0x05, 0xe8, 0x5a, 0x01,
	0xbe, 0x51, 0x80, 0xee, 0x14, 0xe0, 0x7b, 0x05, 0xe8, 0x30, 0x00, 0x74, 0x14, 0x00, 0x3e, 0x0e,
	0x00, 0x7d, 0x0d, 0x00, 0x7f, 0x0f, 0x00, 0x9d, 0x04, 0x80, 0x4e, 0x03, 0xc0, 0x67, 0x01, 0xe0,
	0xf3, 0x00, 0xf0, 0xe7, 0xe5, 0xbf, 0xfd, 0x76, 0xde, 0xae, 0x1d, 0x7e, 0xbd, 0x4e, 0x27, 0x1b,
	0xb5, 0xb3, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x05, 0x52, 0x34, 0x3e, 0x4b, 0x04, 0x00, 0x00,
}
