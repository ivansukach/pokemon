// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pokemon.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Pokemon struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Price                int32    `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	HealthPoints         int32    `protobuf:"varint,4,opt,name=healthPoints,proto3" json:"healthPoints,omitempty"`
	HealthRegen          int32    `protobuf:"varint,5,opt,name=healthRegen,proto3" json:"healthRegen,omitempty"`
	Mana                 int32    `protobuf:"varint,6,opt,name=mana,proto3" json:"mana,omitempty"`
	ManaRegen            int32    `protobuf:"varint,7,opt,name=manaRegen,proto3" json:"manaRegen,omitempty"`
	Armor                float32  `protobuf:"fixed32,8,opt,name=armor,proto3" json:"armor,omitempty"`
	Damage               int32    `protobuf:"varint,9,opt,name=damage,proto3" json:"damage,omitempty"`
	MovementSpeed        int32    `protobuf:"varint,10,opt,name=movementSpeed,proto3" json:"movementSpeed,omitempty"`
	AttackSpeed          int32    `protobuf:"varint,11,opt,name=attackSpeed,proto3" json:"attackSpeed,omitempty"`
	AttackRange          int32    `protobuf:"varint,12,opt,name=attackRange,proto3" json:"attackRange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pokemon) Reset()         { *m = Pokemon{} }
func (m *Pokemon) String() string { return proto.CompactTextString(m) }
func (*Pokemon) ProtoMessage()    {}
func (*Pokemon) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{0}
}

func (m *Pokemon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pokemon.Unmarshal(m, b)
}
func (m *Pokemon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pokemon.Marshal(b, m, deterministic)
}
func (m *Pokemon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pokemon.Merge(m, src)
}
func (m *Pokemon) XXX_Size() int {
	return xxx_messageInfo_Pokemon.Size(m)
}
func (m *Pokemon) XXX_DiscardUnknown() {
	xxx_messageInfo_Pokemon.DiscardUnknown(m)
}

var xxx_messageInfo_Pokemon proto.InternalMessageInfo

func (m *Pokemon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pokemon) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Pokemon) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Pokemon) GetHealthPoints() int32 {
	if m != nil {
		return m.HealthPoints
	}
	return 0
}

func (m *Pokemon) GetHealthRegen() int32 {
	if m != nil {
		return m.HealthRegen
	}
	return 0
}

func (m *Pokemon) GetMana() int32 {
	if m != nil {
		return m.Mana
	}
	return 0
}

func (m *Pokemon) GetManaRegen() int32 {
	if m != nil {
		return m.ManaRegen
	}
	return 0
}

func (m *Pokemon) GetArmor() float32 {
	if m != nil {
		return m.Armor
	}
	return 0
}

func (m *Pokemon) GetDamage() int32 {
	if m != nil {
		return m.Damage
	}
	return 0
}

func (m *Pokemon) GetMovementSpeed() int32 {
	if m != nil {
		return m.MovementSpeed
	}
	return 0
}

func (m *Pokemon) GetAttackSpeed() int32 {
	if m != nil {
		return m.AttackSpeed
	}
	return 0
}

func (m *Pokemon) GetAttackRange() int32 {
	if m != nil {
		return m.AttackRange
	}
	return 0
}

type BuyRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Cost                 int32    `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyRequest) Reset()         { *m = BuyRequest{} }
func (m *BuyRequest) String() string { return proto.CompactTextString(m) }
func (*BuyRequest) ProtoMessage()    {}
func (*BuyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{1}
}

func (m *BuyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyRequest.Unmarshal(m, b)
}
func (m *BuyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyRequest.Marshal(b, m, deterministic)
}
func (m *BuyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyRequest.Merge(m, src)
}
func (m *BuyRequest) XXX_Size() int {
	return xxx_messageInfo_BuyRequest.Size(m)
}
func (m *BuyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyRequest proto.InternalMessageInfo

func (m *BuyRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *BuyRequest) GetCost() int32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

type GetByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIdRequest) Reset()         { *m = GetByIdRequest{} }
func (m *GetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetByIdRequest) ProtoMessage()    {}
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{2}
}

func (m *GetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIdRequest.Unmarshal(m, b)
}
func (m *GetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIdRequest.Merge(m, src)
}
func (m *GetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetByIdRequest.Size(m)
}
func (m *GetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIdRequest proto.InternalMessageInfo

func (m *GetByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetByNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByNameRequest) Reset()         { *m = GetByNameRequest{} }
func (m *GetByNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetByNameRequest) ProtoMessage()    {}
func (*GetByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{4}
}

func (m *GetByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByNameRequest.Unmarshal(m, b)
}
func (m *GetByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByNameRequest.Marshal(b, m, deterministic)
}
func (m *GetByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByNameRequest.Merge(m, src)
}
func (m *GetByNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetByNameRequest.Size(m)
}
func (m *GetByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByNameRequest proto.InternalMessageInfo

func (m *GetByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{5}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type FightResult struct {
	Winner               string   `protobuf:"bytes,1,opt,name=winner,proto3" json:"winner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FightResult) Reset()         { *m = FightResult{} }
func (m *FightResult) String() string { return proto.CompactTextString(m) }
func (*FightResult) ProtoMessage()    {}
func (*FightResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{6}
}

func (m *FightResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FightResult.Unmarshal(m, b)
}
func (m *FightResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FightResult.Marshal(b, m, deterministic)
}
func (m *FightResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FightResult.Merge(m, src)
}
func (m *FightResult) XXX_Size() int {
	return xxx_messageInfo_FightResult.Size(m)
}
func (m *FightResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FightResult.DiscardUnknown(m)
}

var xxx_messageInfo_FightResult proto.InternalMessageInfo

func (m *FightResult) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

type FightRequest struct {
	Id1                  string   `protobuf:"bytes,1,opt,name=id1,proto3" json:"id1,omitempty"`
	Id2                  string   `protobuf:"bytes,2,opt,name=id2,proto3" json:"id2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FightRequest) Reset()         { *m = FightRequest{} }
func (m *FightRequest) String() string { return proto.CompactTextString(m) }
func (*FightRequest) ProtoMessage()    {}
func (*FightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75142c9f6896f36f, []int{7}
}

func (m *FightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FightRequest.Unmarshal(m, b)
}
func (m *FightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FightRequest.Marshal(b, m, deterministic)
}
func (m *FightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FightRequest.Merge(m, src)
}
func (m *FightRequest) XXX_Size() int {
	return xxx_messageInfo_FightRequest.Size(m)
}
func (m *FightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FightRequest proto.InternalMessageInfo

func (m *FightRequest) GetId1() string {
	if m != nil {
		return m.Id1
	}
	return ""
}

func (m *FightRequest) GetId2() string {
	if m != nil {
		return m.Id2
	}
	return ""
}

func init() {
	proto.RegisterType((*Pokemon)(nil), "protocol.Pokemon")
	proto.RegisterType((*BuyRequest)(nil), "protocol.BuyRequest")
	proto.RegisterType((*GetByIdRequest)(nil), "protocol.GetByIdRequest")
	proto.RegisterType((*DeleteRequest)(nil), "protocol.DeleteRequest")
	proto.RegisterType((*GetByNameRequest)(nil), "protocol.GetByNameRequest")
	proto.RegisterType((*EmptyResponse)(nil), "protocol.EmptyResponse")
	proto.RegisterType((*FightResult)(nil), "protocol.FightResult")
	proto.RegisterType((*FightRequest)(nil), "protocol.FightRequest")
}

func init() { proto.RegisterFile("pokemon.proto", fileDescriptor_75142c9f6896f36f) }

var fileDescriptor_75142c9f6896f36f = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x92, 0xda, 0x69, 0x26, 0x49, 0x29, 0x4b, 0x29, 0xab, 0x08, 0x09, 0xcb, 0x02, 0x94,
	0x53, 0x24, 0x42, 0x94, 0x03, 0xc7, 0xf0, 0x25, 0x2e, 0xa8, 0x72, 0xc5, 0x0f, 0x58, 0xe2, 0x91,
	0x63, 0xd5, 0xbb, 0x6b, 0xec, 0x4d, 0x51, 0x7e, 0x0d, 0xbf, 0x8b, 0x7f, 0x83, 0x76, 0x76, 0x1b,
	0x3b, 0x6d, 0x73, 0xe0, 0xe4, 0x99, 0x37, 0xef, 0x6d, 0x5e, 0xde, 0x0c, 0x8c, 0x4b, 0x7d, 0x83,
	0x52, 0xab, 0x59, 0x59, 0x69, 0xa3, 0xd9, 0x29, 0x7d, 0xd6, 0xba, 0x88, 0xff, 0x76, 0xa1, 0x7f,
	0xe5, 0x66, 0x8c, 0xc1, 0x89, 0x12, 0x12, 0x79, 0x27, 0xea, 0x4c, 0x07, 0x09, 0xd5, 0xec, 0x02,
	0x82, 0x5c, 0x8a, 0x0c, 0x79, 0x97, 0x40, 0xd7, 0x58, 0xb4, 0xac, 0xf2, 0x35, 0xf2, 0x5e, 0xd4,
	0x99, 0x06, 0x89, 0x6b, 0x58, 0x0c, 0xa3, 0x0d, 0x8a, 0xc2, 0x6c, 0xae, 0x74, 0xae, 0x4c, 0xcd,
	0x4f, 0x68, 0x78, 0x80, 0xb1, 0x08, 0x86, 0xae, 0x4f, 0x30, 0x43, 0xc5, 0x03, 0xa2, 0xb4, 0x21,
	0xeb, 0x42, 0x0a, 0x25, 0x78, 0x48, 0x23, 0xaa, 0xd9, 0x4b, 0x18, 0xd8, 0xaf, 0xd3, 0xf4, 0x69,
	0xd0, 0x00, 0xd6, 0x8d, 0xa8, 0xa4, 0xae, 0xf8, 0x69, 0xd4, 0x99, 0x76, 0x13, 0xd7, 0xb0, 0x4b,
	0x08, 0x53, 0x41, 0xd6, 0x07, 0x24, 0xf0, 0x1d, 0x7b, 0x0d, 0x63, 0xa9, 0x6f, 0x51, 0xa2, 0x32,
	0xd7, 0x25, 0x62, 0xca, 0x81, 0xc6, 0x87, 0xa0, 0xf5, 0x29, 0x8c, 0x11, 0xeb, 0x1b, 0xc7, 0x19,
	0x3a, 0x9f, 0x2d, 0xa8, 0x61, 0x24, 0x42, 0x65, 0xc8, 0x47, 0x6d, 0x06, 0x41, 0xf1, 0x12, 0x60,
	0xb5, 0xdd, 0x25, 0xf8, 0x6b, 0x8b, 0xb5, 0xb1, 0x2e, 0x0b, 0x9d, 0xe5, 0xca, 0xc7, 0xeb, 0x1a,
	0xfb, 0x6f, 0xd7, 0xba, 0x36, 0x14, 0x6f, 0x90, 0x50, 0x1d, 0x47, 0x70, 0xf6, 0x15, 0xcd, 0x6a,
	0xf7, 0x2d, 0xbd, 0xd3, 0x9e, 0x41, 0x37, 0x4f, 0xbd, 0xb0, 0x9b, 0xa7, 0xf1, 0x2b, 0x18, 0x7f,
	0xc2, 0x02, 0x0d, 0x1e, 0x23, 0xbc, 0x85, 0x73, 0x7a, 0xe2, 0xbb, 0x90, 0x7b, 0xce, 0x23, 0xeb,
	0x8d, 0x9f, 0xc0, 0xf8, 0xb3, 0x2c, 0xcd, 0x2e, 0xc1, 0xba, 0xd4, 0xaa, 0xc6, 0xf8, 0x0d, 0x0c,
	0xbf, 0xe4, 0xd9, 0xc6, 0x24, 0x58, 0x6f, 0x0b, 0x63, 0x43, 0xfc, 0x9d, 0x2b, 0x85, 0x95, 0x57,
	0xf9, 0x2e, 0x9e, 0xc3, 0xc8, 0xd3, 0xdc, 0xdb, 0xe7, 0xd0, 0xcb, 0xd3, 0x77, 0x9e, 0x64, 0x4b,
	0x87, 0xcc, 0xfd, 0xd9, 0xd8, 0x72, 0xfe, 0xa7, 0x07, 0xcf, 0xfc, 0xa9, 0x91, 0xf6, 0x1a, 0xab,
	0x5b, 0x7b, 0x36, 0x0b, 0x08, 0x3f, 0x56, 0x28, 0x0c, 0xb2, 0xa7, 0xb3, 0xbb, 0xbb, 0x9c, 0x79,
	0xe2, 0xe4, 0x45, 0x03, 0x1d, 0x18, 0x65, 0x1f, 0x20, 0x74, 0x11, 0xb0, 0x16, 0xe5, 0x20, 0x94,
	0xe3, 0xda, 0x05, 0x84, 0x3f, 0xca, 0xf4, 0xff, 0x7f, 0x71, 0xb0, 0xcf, 0x94, 0x4d, 0x1a, 0xd6,
	0xfd, 0xa0, 0x27, 0x0f, 0x1f, 0x65, 0x4b, 0xe8, 0xfb, 0x95, 0x32, 0x7e, 0x4f, 0xb9, 0xdf, 0xf2,
	0x63, 0xba, 0x05, 0xf4, 0x56, 0xdb, 0x1d, 0xbb, 0x68, 0x26, 0xcd, 0x45, 0x1d, 0x77, 0xba, 0x84,
	0x80, 0x12, 0x66, 0x97, 0x0d, 0xa3, 0xbd, 0xae, 0xc9, 0xf3, 0x07, 0xb8, 0xdd, 0xf6, 0xcf, 0x90,
	0xd0, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xcb, 0x2c, 0x47, 0x2e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PokemonFightServiceClient is the client API for PokemonFightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PokemonFightServiceClient interface {
	Create(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*EmptyResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	Update(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetByName(ctx context.Context, in *GetByNameRequest, opts ...grpc.CallOption) (*Pokemon, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Pokemon, error)
	Buy(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	Fight(ctx context.Context, in *FightRequest, opts ...grpc.CallOption) (*FightResult, error)
}

type pokemonFightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonFightServiceClient(cc grpc.ClientConnInterface) PokemonFightServiceClient {
	return &pokemonFightServiceClient{cc}
}

func (c *pokemonFightServiceClient) Create(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonFightServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonFightServiceClient) Update(ctx context.Context, in *Pokemon, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonFightServiceClient) GetByName(ctx context.Context, in *GetByNameRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonFightServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonFightServiceClient) Buy(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/Buy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonFightServiceClient) Fight(ctx context.Context, in *FightRequest, opts ...grpc.CallOption) (*FightResult, error) {
	out := new(FightResult)
	err := c.cc.Invoke(ctx, "/protocol.PokemonFightService/Fight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonFightServiceServer is the server API for PokemonFightService service.
type PokemonFightServiceServer interface {
	Create(context.Context, *Pokemon) (*EmptyResponse, error)
	Delete(context.Context, *DeleteRequest) (*EmptyResponse, error)
	Update(context.Context, *Pokemon) (*EmptyResponse, error)
	GetByName(context.Context, *GetByNameRequest) (*Pokemon, error)
	GetById(context.Context, *GetByIdRequest) (*Pokemon, error)
	Buy(context.Context, *BuyRequest) (*EmptyResponse, error)
	Fight(context.Context, *FightRequest) (*FightResult, error)
}

// UnimplementedPokemonFightServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPokemonFightServiceServer struct {
}

func (*UnimplementedPokemonFightServiceServer) Create(ctx context.Context, req *Pokemon) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPokemonFightServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedPokemonFightServiceServer) Update(ctx context.Context, req *Pokemon) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPokemonFightServiceServer) GetByName(ctx context.Context, req *GetByNameRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (*UnimplementedPokemonFightServiceServer) GetById(ctx context.Context, req *GetByIdRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (*UnimplementedPokemonFightServiceServer) Buy(ctx context.Context, req *BuyRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buy not implemented")
}
func (*UnimplementedPokemonFightServiceServer) Fight(ctx context.Context, req *FightRequest) (*FightResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fight not implemented")
}

func RegisterPokemonFightServiceServer(s *grpc.Server, srv PokemonFightServiceServer) {
	s.RegisterService(&_PokemonFightService_serviceDesc, srv)
}

func _PokemonFightService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pokemon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).Create(ctx, req.(*Pokemon))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonFightService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonFightService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pokemon)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).Update(ctx, req.(*Pokemon))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonFightService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).GetByName(ctx, req.(*GetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonFightService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonFightService_Buy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).Buy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/Buy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).Buy(ctx, req.(*BuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonFightService_Fight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonFightServiceServer).Fight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.PokemonFightService/Fight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonFightServiceServer).Fight(ctx, req.(*FightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PokemonFightService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.PokemonFightService",
	HandlerType: (*PokemonFightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PokemonFightService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PokemonFightService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PokemonFightService_Update_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _PokemonFightService_GetByName_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _PokemonFightService_GetById_Handler,
		},
		{
			MethodName: "Buy",
			Handler:    _PokemonFightService_Buy_Handler,
		},
		{
			MethodName: "Fight",
			Handler:    _PokemonFightService_Fight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon.proto",
}
