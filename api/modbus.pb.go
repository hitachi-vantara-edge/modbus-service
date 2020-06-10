// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modbus.proto

package api

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

type ConnectionStatusReply_Type int32

const (
	ConnectionStatusReply_TCP    ConnectionStatusReply_Type = 0
	ConnectionStatusReply_SERIAL ConnectionStatusReply_Type = 1
)

var ConnectionStatusReply_Type_name = map[int32]string{
	0: "TCP",
	1: "SERIAL",
}

var ConnectionStatusReply_Type_value = map[string]int32{
	"TCP":    0,
	"SERIAL": 1,
}

func (x ConnectionStatusReply_Type) String() string {
	return proto.EnumName(ConnectionStatusReply_Type_name, int32(x))
}

func (ConnectionStatusReply_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{8, 0}
}

type ConnectionStatusReply_Status int32

const (
	ConnectionStatusReply_ERROR   ConnectionStatusReply_Status = 0
	ConnectionStatusReply_CLOSED  ConnectionStatusReply_Status = 1
	ConnectionStatusReply_PENDING ConnectionStatusReply_Status = 2
	ConnectionStatusReply_UNKNOWN ConnectionStatusReply_Status = 3
	ConnectionStatusReply_READY   ConnectionStatusReply_Status = 4
)

var ConnectionStatusReply_Status_name = map[int32]string{
	0: "ERROR",
	1: "CLOSED",
	2: "PENDING",
	3: "UNKNOWN",
	4: "READY",
}

var ConnectionStatusReply_Status_value = map[string]int32{
	"ERROR":   0,
	"CLOSED":  1,
	"PENDING": 2,
	"UNKNOWN": 3,
	"READY":   4,
}

func (x ConnectionStatusReply_Status) String() string {
	return proto.EnumName(ConnectionStatusReply_Status_name, int32(x))
}

func (ConnectionStatusReply_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{8, 1}
}

type ReadRegRequest struct {
	ConnectionId         string   `protobuf:"bytes,1,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	SlaveAddr            int32    `protobuf:"varint,2,opt,name=slave_addr,json=slaveAddr,proto3" json:"slave_addr,omitempty"`
	StartAddr            int32    `protobuf:"varint,3,opt,name=start_addr,json=startAddr,proto3" json:"start_addr,omitempty"`
	NumOfReg             int32    `protobuf:"varint,4,opt,name=num_of_reg,json=numOfReg,proto3" json:"num_of_reg,omitempty"`
	ReadInput            bool     `protobuf:"varint,5,opt,name=read_input,json=readInput,proto3" json:"read_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRegRequest) Reset()         { *m = ReadRegRequest{} }
func (m *ReadRegRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRegRequest) ProtoMessage()    {}
func (*ReadRegRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{0}
}

func (m *ReadRegRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRegRequest.Unmarshal(m, b)
}
func (m *ReadRegRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRegRequest.Marshal(b, m, deterministic)
}
func (m *ReadRegRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegRequest.Merge(m, src)
}
func (m *ReadRegRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRegRequest.Size(m)
}
func (m *ReadRegRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRegRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRegRequest proto.InternalMessageInfo

func (m *ReadRegRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *ReadRegRequest) GetSlaveAddr() int32 {
	if m != nil {
		return m.SlaveAddr
	}
	return 0
}

func (m *ReadRegRequest) GetStartAddr() int32 {
	if m != nil {
		return m.StartAddr
	}
	return 0
}

func (m *ReadRegRequest) GetNumOfReg() int32 {
	if m != nil {
		return m.NumOfReg
	}
	return 0
}

func (m *ReadRegRequest) GetReadInput() bool {
	if m != nil {
		return m.ReadInput
	}
	return false
}

type ReadRegReply struct {
	ByteCount            int32                         `protobuf:"varint,1,opt,name=byte_count,json=byteCount,proto3" json:"byte_count,omitempty"`
	IsInput              bool                          `protobuf:"varint,2,opt,name=is_input,json=isInput,proto3" json:"is_input,omitempty"`
	Values               []*ReadRegReply_RegisterValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ReadRegReply) Reset()         { *m = ReadRegReply{} }
func (m *ReadRegReply) String() string { return proto.CompactTextString(m) }
func (*ReadRegReply) ProtoMessage()    {}
func (*ReadRegReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{1}
}

func (m *ReadRegReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRegReply.Unmarshal(m, b)
}
func (m *ReadRegReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRegReply.Marshal(b, m, deterministic)
}
func (m *ReadRegReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegReply.Merge(m, src)
}
func (m *ReadRegReply) XXX_Size() int {
	return xxx_messageInfo_ReadRegReply.Size(m)
}
func (m *ReadRegReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRegReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRegReply proto.InternalMessageInfo

func (m *ReadRegReply) GetByteCount() int32 {
	if m != nil {
		return m.ByteCount
	}
	return 0
}

func (m *ReadRegReply) GetIsInput() bool {
	if m != nil {
		return m.IsInput
	}
	return false
}

func (m *ReadRegReply) GetValues() []*ReadRegReply_RegisterValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReadRegReply_RegisterValue struct {
	Addr                 int32    `protobuf:"varint,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRegReply_RegisterValue) Reset()         { *m = ReadRegReply_RegisterValue{} }
func (m *ReadRegReply_RegisterValue) String() string { return proto.CompactTextString(m) }
func (*ReadRegReply_RegisterValue) ProtoMessage()    {}
func (*ReadRegReply_RegisterValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{1, 0}
}

func (m *ReadRegReply_RegisterValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRegReply_RegisterValue.Unmarshal(m, b)
}
func (m *ReadRegReply_RegisterValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRegReply_RegisterValue.Marshal(b, m, deterministic)
}
func (m *ReadRegReply_RegisterValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegReply_RegisterValue.Merge(m, src)
}
func (m *ReadRegReply_RegisterValue) XXX_Size() int {
	return xxx_messageInfo_ReadRegReply_RegisterValue.Size(m)
}
func (m *ReadRegReply_RegisterValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRegReply_RegisterValue.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRegReply_RegisterValue proto.InternalMessageInfo

func (m *ReadRegReply_RegisterValue) GetAddr() int32 {
	if m != nil {
		return m.Addr
	}
	return 0
}

func (m *ReadRegReply_RegisterValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReadCoilRequest struct {
	ConnectionId         string   `protobuf:"bytes,1,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	SlaveAddr            int32    `protobuf:"varint,2,opt,name=slave_addr,json=slaveAddr,proto3" json:"slave_addr,omitempty"`
	StartAddr            int32    `protobuf:"varint,3,opt,name=start_addr,json=startAddr,proto3" json:"start_addr,omitempty"`
	NumOfCoil            int32    `protobuf:"varint,4,opt,name=num_of_coil,json=numOfCoil,proto3" json:"num_of_coil,omitempty"`
	ReadInput            bool     `protobuf:"varint,5,opt,name=read_input,json=readInput,proto3" json:"read_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadCoilRequest) Reset()         { *m = ReadCoilRequest{} }
func (m *ReadCoilRequest) String() string { return proto.CompactTextString(m) }
func (*ReadCoilRequest) ProtoMessage()    {}
func (*ReadCoilRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{2}
}

func (m *ReadCoilRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCoilRequest.Unmarshal(m, b)
}
func (m *ReadCoilRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCoilRequest.Marshal(b, m, deterministic)
}
func (m *ReadCoilRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCoilRequest.Merge(m, src)
}
func (m *ReadCoilRequest) XXX_Size() int {
	return xxx_messageInfo_ReadCoilRequest.Size(m)
}
func (m *ReadCoilRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCoilRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCoilRequest proto.InternalMessageInfo

func (m *ReadCoilRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *ReadCoilRequest) GetSlaveAddr() int32 {
	if m != nil {
		return m.SlaveAddr
	}
	return 0
}

func (m *ReadCoilRequest) GetStartAddr() int32 {
	if m != nil {
		return m.StartAddr
	}
	return 0
}

func (m *ReadCoilRequest) GetNumOfCoil() int32 {
	if m != nil {
		return m.NumOfCoil
	}
	return 0
}

func (m *ReadCoilRequest) GetReadInput() bool {
	if m != nil {
		return m.ReadInput
	}
	return false
}

type ReadCoilReply struct {
	BitCount             int32                      `protobuf:"varint,1,opt,name=bit_count,json=bitCount,proto3" json:"bit_count,omitempty"`
	IsInput              bool                       `protobuf:"varint,2,opt,name=is_input,json=isInput,proto3" json:"is_input,omitempty"`
	Values               []*ReadCoilReply_CoilValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ReadCoilReply) Reset()         { *m = ReadCoilReply{} }
func (m *ReadCoilReply) String() string { return proto.CompactTextString(m) }
func (*ReadCoilReply) ProtoMessage()    {}
func (*ReadCoilReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{3}
}

func (m *ReadCoilReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCoilReply.Unmarshal(m, b)
}
func (m *ReadCoilReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCoilReply.Marshal(b, m, deterministic)
}
func (m *ReadCoilReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCoilReply.Merge(m, src)
}
func (m *ReadCoilReply) XXX_Size() int {
	return xxx_messageInfo_ReadCoilReply.Size(m)
}
func (m *ReadCoilReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCoilReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCoilReply proto.InternalMessageInfo

func (m *ReadCoilReply) GetBitCount() int32 {
	if m != nil {
		return m.BitCount
	}
	return 0
}

func (m *ReadCoilReply) GetIsInput() bool {
	if m != nil {
		return m.IsInput
	}
	return false
}

func (m *ReadCoilReply) GetValues() []*ReadCoilReply_CoilValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ReadCoilReply_CoilValue struct {
	Addr                 int32    `protobuf:"varint,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadCoilReply_CoilValue) Reset()         { *m = ReadCoilReply_CoilValue{} }
func (m *ReadCoilReply_CoilValue) String() string { return proto.CompactTextString(m) }
func (*ReadCoilReply_CoilValue) ProtoMessage()    {}
func (*ReadCoilReply_CoilValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{3, 0}
}

func (m *ReadCoilReply_CoilValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCoilReply_CoilValue.Unmarshal(m, b)
}
func (m *ReadCoilReply_CoilValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCoilReply_CoilValue.Marshal(b, m, deterministic)
}
func (m *ReadCoilReply_CoilValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCoilReply_CoilValue.Merge(m, src)
}
func (m *ReadCoilReply_CoilValue) XXX_Size() int {
	return xxx_messageInfo_ReadCoilReply_CoilValue.Size(m)
}
func (m *ReadCoilReply_CoilValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCoilReply_CoilValue.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCoilReply_CoilValue proto.InternalMessageInfo

func (m *ReadCoilReply_CoilValue) GetAddr() int32 {
	if m != nil {
		return m.Addr
	}
	return 0
}

func (m *ReadCoilReply_CoilValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type OpenTCPConnectionRequest struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenTCPConnectionRequest) Reset()         { *m = OpenTCPConnectionRequest{} }
func (m *OpenTCPConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*OpenTCPConnectionRequest) ProtoMessage()    {}
func (*OpenTCPConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{4}
}

func (m *OpenTCPConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenTCPConnectionRequest.Unmarshal(m, b)
}
func (m *OpenTCPConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenTCPConnectionRequest.Marshal(b, m, deterministic)
}
func (m *OpenTCPConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenTCPConnectionRequest.Merge(m, src)
}
func (m *OpenTCPConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_OpenTCPConnectionRequest.Size(m)
}
func (m *OpenTCPConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenTCPConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenTCPConnectionRequest proto.InternalMessageInfo

func (m *OpenTCPConnectionRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *OpenTCPConnectionRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type OpenSerialConnectionRequest struct {
	Device               string   `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Baud                 int32    `protobuf:"varint,2,opt,name=baud,proto3" json:"baud,omitempty"`
	Parity               string   `protobuf:"bytes,3,opt,name=parity,proto3" json:"parity,omitempty"`
	DataBit              int32    `protobuf:"varint,4,opt,name=data_bit,json=dataBit,proto3" json:"data_bit,omitempty"`
	StopBit              int32    `protobuf:"varint,5,opt,name=stop_bit,json=stopBit,proto3" json:"stop_bit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenSerialConnectionRequest) Reset()         { *m = OpenSerialConnectionRequest{} }
func (m *OpenSerialConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*OpenSerialConnectionRequest) ProtoMessage()    {}
func (*OpenSerialConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{5}
}

func (m *OpenSerialConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenSerialConnectionRequest.Unmarshal(m, b)
}
func (m *OpenSerialConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenSerialConnectionRequest.Marshal(b, m, deterministic)
}
func (m *OpenSerialConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenSerialConnectionRequest.Merge(m, src)
}
func (m *OpenSerialConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_OpenSerialConnectionRequest.Size(m)
}
func (m *OpenSerialConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenSerialConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenSerialConnectionRequest proto.InternalMessageInfo

func (m *OpenSerialConnectionRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *OpenSerialConnectionRequest) GetBaud() int32 {
	if m != nil {
		return m.Baud
	}
	return 0
}

func (m *OpenSerialConnectionRequest) GetParity() string {
	if m != nil {
		return m.Parity
	}
	return ""
}

func (m *OpenSerialConnectionRequest) GetDataBit() int32 {
	if m != nil {
		return m.DataBit
	}
	return 0
}

func (m *OpenSerialConnectionRequest) GetStopBit() int32 {
	if m != nil {
		return m.StopBit
	}
	return 0
}

type CloseConnectionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseConnectionRequest) Reset()         { *m = CloseConnectionRequest{} }
func (m *CloseConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*CloseConnectionRequest) ProtoMessage()    {}
func (*CloseConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{6}
}

func (m *CloseConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConnectionRequest.Unmarshal(m, b)
}
func (m *CloseConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConnectionRequest.Marshal(b, m, deterministic)
}
func (m *CloseConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConnectionRequest.Merge(m, src)
}
func (m *CloseConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_CloseConnectionRequest.Size(m)
}
func (m *CloseConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseConnectionRequest proto.InternalMessageInfo

func (m *CloseConnectionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetConnectionStatusRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConnectionStatusRequest) Reset()         { *m = GetConnectionStatusRequest{} }
func (m *GetConnectionStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetConnectionStatusRequest) ProtoMessage()    {}
func (*GetConnectionStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{7}
}

func (m *GetConnectionStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionStatusRequest.Unmarshal(m, b)
}
func (m *GetConnectionStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetConnectionStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionStatusRequest.Merge(m, src)
}
func (m *GetConnectionStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetConnectionStatusRequest.Size(m)
}
func (m *GetConnectionStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectionStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectionStatusRequest proto.InternalMessageInfo

func (m *GetConnectionStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ConnectionStatusReply struct {
	Id                   string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Props                map[string]string            `protobuf:"bytes,2,rep,name=props,proto3" json:"props,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 ConnectionStatusReply_Type   `protobuf:"varint,3,opt,name=type,proto3,enum=api.ConnectionStatusReply_Type" json:"type,omitempty"`
	Status               ConnectionStatusReply_Status `protobuf:"varint,4,opt,name=status,proto3,enum=api.ConnectionStatusReply_Status" json:"status,omitempty"`
	Errmsg               string                       `protobuf:"bytes,5,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ConnectionStatusReply) Reset()         { *m = ConnectionStatusReply{} }
func (m *ConnectionStatusReply) String() string { return proto.CompactTextString(m) }
func (*ConnectionStatusReply) ProtoMessage()    {}
func (*ConnectionStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4384190c9e71792, []int{8}
}

func (m *ConnectionStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionStatusReply.Unmarshal(m, b)
}
func (m *ConnectionStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionStatusReply.Marshal(b, m, deterministic)
}
func (m *ConnectionStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionStatusReply.Merge(m, src)
}
func (m *ConnectionStatusReply) XXX_Size() int {
	return xxx_messageInfo_ConnectionStatusReply.Size(m)
}
func (m *ConnectionStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionStatusReply proto.InternalMessageInfo

func (m *ConnectionStatusReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConnectionStatusReply) GetProps() map[string]string {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *ConnectionStatusReply) GetType() ConnectionStatusReply_Type {
	if m != nil {
		return m.Type
	}
	return ConnectionStatusReply_TCP
}

func (m *ConnectionStatusReply) GetStatus() ConnectionStatusReply_Status {
	if m != nil {
		return m.Status
	}
	return ConnectionStatusReply_ERROR
}

func (m *ConnectionStatusReply) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.ConnectionStatusReply_Type", ConnectionStatusReply_Type_name, ConnectionStatusReply_Type_value)
	proto.RegisterEnum("api.ConnectionStatusReply_Status", ConnectionStatusReply_Status_name, ConnectionStatusReply_Status_value)
	proto.RegisterType((*ReadRegRequest)(nil), "api.ReadRegRequest")
	proto.RegisterType((*ReadRegReply)(nil), "api.ReadRegReply")
	proto.RegisterType((*ReadRegReply_RegisterValue)(nil), "api.ReadRegReply.RegisterValue")
	proto.RegisterType((*ReadCoilRequest)(nil), "api.ReadCoilRequest")
	proto.RegisterType((*ReadCoilReply)(nil), "api.ReadCoilReply")
	proto.RegisterType((*ReadCoilReply_CoilValue)(nil), "api.ReadCoilReply.CoilValue")
	proto.RegisterType((*OpenTCPConnectionRequest)(nil), "api.OpenTCPConnectionRequest")
	proto.RegisterType((*OpenSerialConnectionRequest)(nil), "api.OpenSerialConnectionRequest")
	proto.RegisterType((*CloseConnectionRequest)(nil), "api.CloseConnectionRequest")
	proto.RegisterType((*GetConnectionStatusRequest)(nil), "api.GetConnectionStatusRequest")
	proto.RegisterType((*ConnectionStatusReply)(nil), "api.ConnectionStatusReply")
	proto.RegisterMapType((map[string]string)(nil), "api.ConnectionStatusReply.PropsEntry")
}

func init() { proto.RegisterFile("modbus.proto", fileDescriptor_f4384190c9e71792) }

var fileDescriptor_f4384190c9e71792 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6f, 0xda, 0x4c,
	0x10, 0xc6, 0x98, 0x2f, 0x4f, 0xf2, 0x12, 0x32, 0xc9, 0x1b, 0x51, 0x48, 0x5a, 0x6a, 0xa9, 0x12,
	0x87, 0x8a, 0x43, 0xd2, 0x8f, 0xa4, 0x3d, 0x25, 0x80, 0x22, 0xd4, 0x14, 0xd0, 0x42, 0x53, 0xf5,
	0x84, 0x0c, 0xde, 0xa0, 0x55, 0xc1, 0x76, 0xed, 0x75, 0x24, 0x7e, 0x42, 0x7f, 0x41, 0x7f, 0x46,
	0xa5, 0xaa, 0xc7, 0xde, 0xfa, 0xc7, 0xaa, 0x5d, 0x2f, 0x24, 0x7c, 0x05, 0xf5, 0xd2, 0xdb, 0xce,
	0xcc, 0x33, 0x8f, 0xc7, 0xcf, 0xf8, 0x59, 0xc3, 0xf6, 0xd8, 0xb5, 0xfb, 0x61, 0x50, 0xf1, 0x7c,
	0x97, 0xbb, 0xa8, 0x5b, 0x1e, 0x33, 0xbf, 0x6b, 0x90, 0x25, 0xd4, 0xb2, 0x09, 0x1d, 0x12, 0xfa,
	0x25, 0xa4, 0x01, 0x47, 0x13, 0xb6, 0x07, 0xae, 0xe3, 0xd0, 0x01, 0x67, 0xae, 0xd3, 0xb0, 0xf3,
	0x5a, 0x49, 0x2b, 0x1b, 0x64, 0x2e, 0x87, 0x47, 0x00, 0xc1, 0xc8, 0xba, 0xa5, 0x3d, 0xcb, 0xb6,
	0xfd, 0x7c, 0xbc, 0xa4, 0x95, 0x93, 0xc4, 0x90, 0x99, 0x73, 0xdb, 0xf6, 0x65, 0x99, 0x5b, 0x3e,
	0x8f, 0xca, 0xba, 0x2a, 0x8b, 0x8c, 0x2c, 0x1f, 0x02, 0x38, 0xe1, 0xb8, 0xe7, 0xde, 0xf4, 0x7c,
	0x3a, 0xcc, 0x27, 0x64, 0x39, 0xe3, 0x84, 0xe3, 0xd6, 0x0d, 0xa1, 0x43, 0xd1, 0xec, 0x53, 0xcb,
	0xee, 0x31, 0xc7, 0x0b, 0x79, 0x3e, 0x59, 0xd2, 0xca, 0x19, 0x62, 0x88, 0x4c, 0x43, 0x24, 0xcc,
	0x5f, 0x1a, 0x6c, 0xcf, 0x26, 0xf6, 0x46, 0x13, 0x81, 0xef, 0x4f, 0x38, 0xed, 0x0d, 0xdc, 0xd0,
	0xe1, 0x72, 0xda, 0x24, 0x31, 0x44, 0xa6, 0x2a, 0x12, 0xf8, 0x08, 0x32, 0x2c, 0x50, 0x64, 0x71,
	0x49, 0x96, 0x66, 0x81, 0xa4, 0xc2, 0xd7, 0x90, 0xba, 0xb5, 0x46, 0x21, 0x0d, 0xf2, 0x7a, 0x49,
	0x2f, 0x6f, 0x1d, 0x3f, 0xa9, 0x58, 0x1e, 0xab, 0xdc, 0x27, 0xaf, 0x10, 0x3a, 0x64, 0x01, 0xa7,
	0xfe, 0xb5, 0xc0, 0x11, 0x05, 0x2f, 0x9c, 0xc1, 0x7f, 0x73, 0x05, 0x44, 0x48, 0xc8, 0x57, 0x8d,
	0x9e, 0x2e, 0xcf, 0xb8, 0x0f, 0x49, 0x09, 0x97, 0x4f, 0x35, 0x48, 0x14, 0x98, 0x3f, 0x34, 0xd8,
	0x11, 0x4f, 0xa8, 0xba, 0x6c, 0xf4, 0xef, 0x14, 0x7f, 0x0c, 0x5b, 0x4a, 0xf1, 0x81, 0xcb, 0x46,
	0x4a, 0x72, 0x43, 0x4a, 0x2e, 0x06, 0xd9, 0xa4, 0xf9, 0x4f, 0x4d, 0xbc, 0xf0, 0x74, 0x68, 0x21,
	0x7a, 0x11, 0x8c, 0x3e, 0xe3, 0x73, 0x9a, 0x67, 0xfa, 0x8c, 0x6f, 0x94, 0xfc, 0xc5, 0x82, 0xe4,
	0x87, 0x33, 0xc9, 0x67, 0xdc, 0x15, 0x71, 0x9a, 0xd7, 0xfb, 0x25, 0x18, 0xb3, 0xe4, 0x5f, 0x68,
	0x7d, 0x01, 0xf9, 0x96, 0x47, 0x9d, 0x6e, 0xb5, 0x5d, 0x9d, 0x49, 0x39, 0xd5, 0xfc, 0x3e, 0x8b,
	0xa1, 0x58, 0x10, 0x12, 0x9e, 0xeb, 0x73, 0xa5, 0xae, 0x3c, 0x9b, 0xdf, 0x34, 0x28, 0x0a, 0x92,
	0x0e, 0xf5, 0x99, 0x35, 0x5a, 0xe6, 0x39, 0x80, 0x94, 0x4d, 0x6f, 0xd9, 0x80, 0x2a, 0x26, 0x15,
	0x09, 0xae, 0xbe, 0x15, 0xda, 0x53, 0x2e, 0x71, 0x16, 0x58, 0xcf, 0xf2, 0x19, 0x9f, 0xc8, 0x05,
	0x19, 0x44, 0x45, 0x42, 0x2f, 0xdb, 0xe2, 0x56, 0xaf, 0xcf, 0xb8, 0x5a, 0x4d, 0x5a, 0xc4, 0x17,
	0x4c, 0x4a, 0x19, 0x70, 0xd7, 0x93, 0xa5, 0x64, 0x54, 0x12, 0xf1, 0x05, 0xe3, 0x66, 0x19, 0x0e,
	0xaa, 0x23, 0x37, 0xa0, 0xcb, 0x33, 0x65, 0x21, 0xce, 0xa6, 0x5f, 0x51, 0x9c, 0xd9, 0xe6, 0x73,
	0x28, 0x5c, 0x52, 0x7e, 0x87, 0xeb, 0x70, 0x8b, 0x87, 0xc1, 0x3a, 0xf4, 0x57, 0x1d, 0xfe, 0x5f,
	0xc6, 0x8a, 0xa5, 0x2f, 0x20, 0xf1, 0x2d, 0x24, 0x3d, 0xdf, 0xf5, 0x82, 0x7c, 0x5c, 0xee, 0xf2,
	0x99, 0xdc, 0xe5, 0xca, 0xd6, 0x4a, 0x5b, 0xe0, 0xea, 0x0e, 0xf7, 0x27, 0x24, 0xea, 0xc1, 0x13,
	0x48, 0xf0, 0x89, 0x47, 0xa5, 0x14, 0x59, 0x65, 0xbd, 0xd5, 0xbd, 0xdd, 0x89, 0x47, 0x89, 0x04,
	0xe3, 0x19, 0xa4, 0x02, 0x59, 0x91, 0x3a, 0x65, 0x8f, 0x9f, 0x3e, 0xd0, 0xa6, 0xce, 0xaa, 0x41,
	0x88, 0x4f, 0x7d, 0x7f, 0x1c, 0x0c, 0xa5, 0x8e, 0x06, 0x51, 0x51, 0xe1, 0x14, 0xe0, 0x6e, 0x38,
	0xcc, 0x81, 0xfe, 0x99, 0x4e, 0xd4, 0x3b, 0x8a, 0xe3, 0xea, 0x4f, 0xeb, 0x4d, 0xfc, 0x54, 0x33,
	0x8b, 0x90, 0x10, 0xa3, 0x61, 0x1a, 0xf4, 0x6e, 0xb5, 0x9d, 0x8b, 0x21, 0x40, 0xaa, 0x53, 0x27,
	0x8d, 0xf3, 0xab, 0x9c, 0x66, 0xd6, 0x20, 0x15, 0x0d, 0x80, 0x06, 0x24, 0xeb, 0x84, 0xb4, 0x48,
	0x04, 0xa8, 0x5e, 0xb5, 0x3a, 0xf5, 0x5a, 0x4e, 0xc3, 0x2d, 0x48, 0xb7, 0xeb, 0xcd, 0x5a, 0xa3,
	0x79, 0x99, 0x8b, 0x8b, 0xe0, 0x43, 0xf3, 0x5d, 0xb3, 0xf5, 0xb1, 0x99, 0xd3, 0x45, 0x03, 0xa9,
	0x9f, 0xd7, 0x3e, 0xe5, 0x12, 0xc7, 0xbf, 0x75, 0x48, 0xbd, 0x97, 0x97, 0x36, 0x9e, 0x40, 0x5a,
	0xdd, 0x4c, 0xb8, 0x37, 0x7f, 0x4f, 0xc9, 0x35, 0x16, 0x76, 0x97, 0x2e, 0x2f, 0x33, 0x86, 0xaf,
	0x20, 0x33, 0xf5, 0x16, 0xee, 0x2f, 0x58, 0x2d, 0x6a, 0xc3, 0x65, 0x03, 0x9a, 0x31, 0x6c, 0xc3,
	0xee, 0x92, 0x73, 0xf0, 0x48, 0x42, 0xd7, 0x39, 0xaa, 0x50, 0x58, 0xbf, 0x0b, 0x33, 0x86, 0xd7,
	0xb0, 0xbf, 0xca, 0x46, 0x58, 0x9a, 0x91, 0xae, 0x71, 0xd8, 0x06, 0xde, 0x2b, 0xd8, 0x59, 0x70,
	0x01, 0x16, 0xa3, 0x86, 0x95, 0xde, 0xd8, 0xc0, 0xd6, 0x85, 0xbd, 0x15, 0x4e, 0xc1, 0xe8, 0xeb,
	0x5c, 0xef, 0xa1, 0x87, 0x59, 0xfb, 0x29, 0xf9, 0xc3, 0x3d, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x72, 0x21, 0x9b, 0xd3, 0x80, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ModbusClient is the client API for Modbus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ModbusClient interface {
	ReadReg(ctx context.Context, in *ReadRegRequest, opts ...grpc.CallOption) (*ReadRegReply, error)
	ReadCoil(ctx context.Context, in *ReadCoilRequest, opts ...grpc.CallOption) (*ReadCoilReply, error)
	OpenTCPConnection(ctx context.Context, in *OpenTCPConnectionRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error)
	OpenSerialConnection(ctx context.Context, in *OpenSerialConnectionRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error)
	CloseConnection(ctx context.Context, in *CloseConnectionRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error)
	GetConnectionStatus(ctx context.Context, in *GetConnectionStatusRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error)
}

type modbusClient struct {
	cc *grpc.ClientConn
}

func NewModbusClient(cc *grpc.ClientConn) ModbusClient {
	return &modbusClient{cc}
}

func (c *modbusClient) ReadReg(ctx context.Context, in *ReadRegRequest, opts ...grpc.CallOption) (*ReadRegReply, error) {
	out := new(ReadRegReply)
	err := c.cc.Invoke(ctx, "/api.Modbus/ReadReg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modbusClient) ReadCoil(ctx context.Context, in *ReadCoilRequest, opts ...grpc.CallOption) (*ReadCoilReply, error) {
	out := new(ReadCoilReply)
	err := c.cc.Invoke(ctx, "/api.Modbus/ReadCoil", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modbusClient) OpenTCPConnection(ctx context.Context, in *OpenTCPConnectionRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error) {
	out := new(ConnectionStatusReply)
	err := c.cc.Invoke(ctx, "/api.Modbus/OpenTCPConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modbusClient) OpenSerialConnection(ctx context.Context, in *OpenSerialConnectionRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error) {
	out := new(ConnectionStatusReply)
	err := c.cc.Invoke(ctx, "/api.Modbus/OpenSerialConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modbusClient) CloseConnection(ctx context.Context, in *CloseConnectionRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error) {
	out := new(ConnectionStatusReply)
	err := c.cc.Invoke(ctx, "/api.Modbus/CloseConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modbusClient) GetConnectionStatus(ctx context.Context, in *GetConnectionStatusRequest, opts ...grpc.CallOption) (*ConnectionStatusReply, error) {
	out := new(ConnectionStatusReply)
	err := c.cc.Invoke(ctx, "/api.Modbus/GetConnectionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModbusServer is the server API for Modbus service.
type ModbusServer interface {
	ReadReg(context.Context, *ReadRegRequest) (*ReadRegReply, error)
	ReadCoil(context.Context, *ReadCoilRequest) (*ReadCoilReply, error)
	OpenTCPConnection(context.Context, *OpenTCPConnectionRequest) (*ConnectionStatusReply, error)
	OpenSerialConnection(context.Context, *OpenSerialConnectionRequest) (*ConnectionStatusReply, error)
	CloseConnection(context.Context, *CloseConnectionRequest) (*ConnectionStatusReply, error)
	GetConnectionStatus(context.Context, *GetConnectionStatusRequest) (*ConnectionStatusReply, error)
}

// UnimplementedModbusServer can be embedded to have forward compatible implementations.
type UnimplementedModbusServer struct {
}

func (*UnimplementedModbusServer) ReadReg(ctx context.Context, req *ReadRegRequest) (*ReadRegReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadReg not implemented")
}
func (*UnimplementedModbusServer) ReadCoil(ctx context.Context, req *ReadCoilRequest) (*ReadCoilReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCoil not implemented")
}
func (*UnimplementedModbusServer) OpenTCPConnection(ctx context.Context, req *OpenTCPConnectionRequest) (*ConnectionStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenTCPConnection not implemented")
}
func (*UnimplementedModbusServer) OpenSerialConnection(ctx context.Context, req *OpenSerialConnectionRequest) (*ConnectionStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSerialConnection not implemented")
}
func (*UnimplementedModbusServer) CloseConnection(ctx context.Context, req *CloseConnectionRequest) (*ConnectionStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseConnection not implemented")
}
func (*UnimplementedModbusServer) GetConnectionStatus(ctx context.Context, req *GetConnectionStatusRequest) (*ConnectionStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionStatus not implemented")
}

func RegisterModbusServer(s *grpc.Server, srv ModbusServer) {
	s.RegisterService(&_Modbus_serviceDesc, srv)
}

func _Modbus_ReadReg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusServer).ReadReg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Modbus/ReadReg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusServer).ReadReg(ctx, req.(*ReadRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modbus_ReadCoil_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCoilRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusServer).ReadCoil(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Modbus/ReadCoil",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusServer).ReadCoil(ctx, req.(*ReadCoilRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modbus_OpenTCPConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenTCPConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusServer).OpenTCPConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Modbus/OpenTCPConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusServer).OpenTCPConnection(ctx, req.(*OpenTCPConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modbus_OpenSerialConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSerialConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusServer).OpenSerialConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Modbus/OpenSerialConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusServer).OpenSerialConnection(ctx, req.(*OpenSerialConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modbus_CloseConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusServer).CloseConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Modbus/CloseConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusServer).CloseConnection(ctx, req.(*CloseConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modbus_GetConnectionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModbusServer).GetConnectionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Modbus/GetConnectionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModbusServer).GetConnectionStatus(ctx, req.(*GetConnectionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Modbus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Modbus",
	HandlerType: (*ModbusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadReg",
			Handler:    _Modbus_ReadReg_Handler,
		},
		{
			MethodName: "ReadCoil",
			Handler:    _Modbus_ReadCoil_Handler,
		},
		{
			MethodName: "OpenTCPConnection",
			Handler:    _Modbus_OpenTCPConnection_Handler,
		},
		{
			MethodName: "OpenSerialConnection",
			Handler:    _Modbus_OpenSerialConnection_Handler,
		},
		{
			MethodName: "CloseConnection",
			Handler:    _Modbus_CloseConnection_Handler,
		},
		{
			MethodName: "GetConnectionStatus",
			Handler:    _Modbus_GetConnectionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modbus.proto",
}
