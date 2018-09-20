// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modbus.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{8, 0}
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{8, 1}
}

type ReadRegRequest struct {
	SlaveAddr            int32    `protobuf:"varint,1,opt,name=slave_addr,json=slaveAddr,proto3" json:"slave_addr,omitempty"`
	StartAddr            int32    `protobuf:"varint,2,opt,name=start_addr,json=startAddr,proto3" json:"start_addr,omitempty"`
	NumOfReg             int32    `protobuf:"varint,3,opt,name=num_of_reg,json=numOfReg,proto3" json:"num_of_reg,omitempty"`
	ReadInput            bool     `protobuf:"varint,4,opt,name=read_input,json=readInput,proto3" json:"read_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRegRequest) Reset()         { *m = ReadRegRequest{} }
func (m *ReadRegRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRegRequest) ProtoMessage()    {}
func (*ReadRegRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{0}
}
func (m *ReadRegRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRegRequest.Unmarshal(m, b)
}
func (m *ReadRegRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRegRequest.Marshal(b, m, deterministic)
}
func (dst *ReadRegRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegRequest.Merge(dst, src)
}
func (m *ReadRegRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRegRequest.Size(m)
}
func (m *ReadRegRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRegRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRegRequest proto.InternalMessageInfo

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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{1}
}
func (m *ReadRegReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRegReply.Unmarshal(m, b)
}
func (m *ReadRegReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRegReply.Marshal(b, m, deterministic)
}
func (dst *ReadRegReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegReply.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{1, 0}
}
func (m *ReadRegReply_RegisterValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRegReply_RegisterValue.Unmarshal(m, b)
}
func (m *ReadRegReply_RegisterValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRegReply_RegisterValue.Marshal(b, m, deterministic)
}
func (dst *ReadRegReply_RegisterValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegReply_RegisterValue.Merge(dst, src)
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
	SlaveAddr            int32    `protobuf:"varint,1,opt,name=slave_addr,json=slaveAddr,proto3" json:"slave_addr,omitempty"`
	StartAddr            int32    `protobuf:"varint,2,opt,name=start_addr,json=startAddr,proto3" json:"start_addr,omitempty"`
	NumOfCoil            int32    `protobuf:"varint,3,opt,name=num_of_coil,json=numOfCoil,proto3" json:"num_of_coil,omitempty"`
	ReadInput            bool     `protobuf:"varint,4,opt,name=read_input,json=readInput,proto3" json:"read_input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadCoilRequest) Reset()         { *m = ReadCoilRequest{} }
func (m *ReadCoilRequest) String() string { return proto.CompactTextString(m) }
func (*ReadCoilRequest) ProtoMessage()    {}
func (*ReadCoilRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{2}
}
func (m *ReadCoilRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCoilRequest.Unmarshal(m, b)
}
func (m *ReadCoilRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCoilRequest.Marshal(b, m, deterministic)
}
func (dst *ReadCoilRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCoilRequest.Merge(dst, src)
}
func (m *ReadCoilRequest) XXX_Size() int {
	return xxx_messageInfo_ReadCoilRequest.Size(m)
}
func (m *ReadCoilRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCoilRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCoilRequest proto.InternalMessageInfo

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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{3}
}
func (m *ReadCoilReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCoilReply.Unmarshal(m, b)
}
func (m *ReadCoilReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCoilReply.Marshal(b, m, deterministic)
}
func (dst *ReadCoilReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCoilReply.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{3, 0}
}
func (m *ReadCoilReply_CoilValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCoilReply_CoilValue.Unmarshal(m, b)
}
func (m *ReadCoilReply_CoilValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCoilReply_CoilValue.Marshal(b, m, deterministic)
}
func (dst *ReadCoilReply_CoilValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCoilReply_CoilValue.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{4}
}
func (m *OpenTCPConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenTCPConnectionRequest.Unmarshal(m, b)
}
func (m *OpenTCPConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenTCPConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *OpenTCPConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenTCPConnectionRequest.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{5}
}
func (m *OpenSerialConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenSerialConnectionRequest.Unmarshal(m, b)
}
func (m *OpenSerialConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenSerialConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *OpenSerialConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenSerialConnectionRequest.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{6}
}
func (m *CloseConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConnectionRequest.Unmarshal(m, b)
}
func (m *CloseConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *CloseConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConnectionRequest.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{7}
}
func (m *GetConnectionStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionStatusRequest.Unmarshal(m, b)
}
func (m *GetConnectionStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionStatusRequest.Marshal(b, m, deterministic)
}
func (dst *GetConnectionStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionStatusRequest.Merge(dst, src)
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
	return fileDescriptor_modbus_cafb66b4f02cb8d6, []int{8}
}
func (m *ConnectionStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionStatusReply.Unmarshal(m, b)
}
func (m *ConnectionStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionStatusReply.Marshal(b, m, deterministic)
}
func (dst *ConnectionStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionStatusReply.Merge(dst, src)
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
	proto.RegisterEnum("api.ConnectionStatusReply_Type", ConnectionStatusReply_Type_name, ConnectionStatusReply_Type_value)
	proto.RegisterEnum("api.ConnectionStatusReply_Status", ConnectionStatusReply_Status_name, ConnectionStatusReply_Status_value)
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

func init() { proto.RegisterFile("modbus.proto", fileDescriptor_modbus_cafb66b4f02cb8d6) }

var fileDescriptor_modbus_cafb66b4f02cb8d6 = []byte{
	// 759 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xda, 0x4a,
	0x14, 0xc5, 0x98, 0x2f, 0xdf, 0xe4, 0x11, 0x32, 0xc9, 0x8b, 0x78, 0x90, 0xbc, 0xc7, 0xb3, 0x54,
	0x89, 0x45, 0xc5, 0x22, 0xe9, 0x47, 0xd2, 0xae, 0x12, 0x40, 0x51, 0xd4, 0x14, 0xd0, 0x40, 0x53,
	0x75, 0x85, 0x06, 0x3c, 0x41, 0xa3, 0x82, 0xed, 0x8e, 0xc7, 0x91, 0xfc, 0x13, 0xaa, 0x4a, 0xdd,
	0xf6, 0x8f, 0x74, 0xd9, 0x5d, 0xff, 0x58, 0x35, 0xe3, 0x81, 0x84, 0xaf, 0xa0, 0x4a, 0xdd, 0xcd,
	0xbd, 0xf7, 0xdc, 0x33, 0x97, 0x73, 0x7d, 0x06, 0xd8, 0x9e, 0x78, 0xce, 0x20, 0x0c, 0x6a, 0x3e,
	0xf7, 0x84, 0x87, 0x4c, 0xe2, 0x33, 0xfb, 0x8b, 0x01, 0x79, 0x4c, 0x89, 0x83, 0xe9, 0x08, 0xd3,
	0x4f, 0x21, 0x0d, 0x04, 0x3a, 0x02, 0x08, 0xc6, 0xe4, 0x8e, 0xf6, 0x89, 0xe3, 0xf0, 0xa2, 0x51,
	0x31, 0xaa, 0x69, 0x6c, 0xa9, 0xcc, 0xb9, 0xe3, 0x70, 0x55, 0x16, 0x84, 0x8b, 0xb8, 0x9c, 0xd4,
	0x65, 0x99, 0x51, 0xe5, 0x43, 0x00, 0x37, 0x9c, 0xf4, 0xbd, 0xdb, 0x3e, 0xa7, 0xa3, 0xa2, 0xa9,
	0xca, 0x39, 0x37, 0x9c, 0xb4, 0x6f, 0x31, 0x1d, 0xc9, 0x66, 0x4e, 0x89, 0xd3, 0x67, 0xae, 0x1f,
	0x8a, 0x62, 0xaa, 0x62, 0x54, 0x73, 0xd8, 0x92, 0x99, 0x2b, 0x99, 0xb0, 0x7f, 0x18, 0xb0, 0x3d,
	0x9b, 0xc6, 0x1f, 0x47, 0x12, 0x3f, 0x88, 0x04, 0xed, 0x0f, 0xbd, 0xd0, 0x15, 0xd3, 0x59, 0x64,
	0xa6, 0x2e, 0x13, 0xe8, 0x1f, 0xc8, 0xb1, 0x40, 0x93, 0x25, 0x15, 0x59, 0x96, 0x05, 0x8a, 0x0a,
	0xbd, 0x84, 0xcc, 0x1d, 0x19, 0x87, 0x34, 0x28, 0x9a, 0x15, 0xb3, 0xba, 0x75, 0xfc, 0x5f, 0x8d,
	0xf8, 0xac, 0xf6, 0x90, 0xbc, 0x86, 0xe9, 0x88, 0x05, 0x82, 0xf2, 0x1b, 0x89, 0xc3, 0x1a, 0x5e,
	0x3a, 0x83, 0xbf, 0xe6, 0x0a, 0x08, 0x41, 0xea, 0x81, 0x12, 0xea, 0x8c, 0xf6, 0x21, 0xad, 0xe0,
	0xea, 0x56, 0x0b, 0xc7, 0x81, 0xfd, 0xd5, 0x80, 0x1d, 0x79, 0x43, 0xdd, 0x63, 0xe3, 0x3f, 0xa3,
	0xe6, 0xbf, 0xb0, 0xa5, 0xd5, 0x1c, 0x7a, 0x6c, 0xac, 0xe5, 0xb4, 0x94, 0x9c, 0xf2, 0x92, 0x4d,
	0x7a, 0x7e, 0x37, 0xe4, 0x8f, 0x99, 0x0e, 0x24, 0x05, 0x2d, 0x83, 0x35, 0x60, 0x62, 0x4e, 0xcf,
	0xdc, 0x80, 0x89, 0x8d, 0x72, 0x3e, 0x5b, 0x90, 0xf3, 0x70, 0x26, 0xe7, 0x8c, 0xbb, 0x26, 0x4f,
	0xf3, 0x5a, 0x3e, 0x07, 0x6b, 0x96, 0xfc, 0x0d, 0x1d, 0x2f, 0xa0, 0xd8, 0xf6, 0xa9, 0xdb, 0xab,
	0x77, 0xea, 0x9e, 0xeb, 0xd2, 0xa1, 0x60, 0x9e, 0x3b, 0xd5, 0xf3, 0x21, 0x8b, 0xa5, 0x59, 0x10,
	0xa4, 0x7c, 0x8f, 0x0b, 0x2d, 0x9f, 0x3a, 0xdb, 0xdf, 0x0c, 0x28, 0x4b, 0x92, 0x2e, 0xe5, 0x8c,
	0x8c, 0x97, 0x79, 0x0e, 0x20, 0xe3, 0xd0, 0x3b, 0x36, 0xa4, 0x9a, 0x49, 0x47, 0x92, 0x6b, 0x40,
	0x42, 0x67, 0xca, 0x25, 0xcf, 0x12, 0xeb, 0x13, 0xce, 0x44, 0xa4, 0x16, 0x60, 0x61, 0x1d, 0x49,
	0xbd, 0x1c, 0x22, 0x48, 0x7f, 0xc0, 0x62, 0xed, 0xd3, 0x38, 0x2b, 0xe3, 0x0b, 0xa6, 0xa4, 0x0c,
	0x84, 0xe7, 0xab, 0x52, 0x3a, 0x2e, 0xc9, 0xf8, 0x82, 0x09, 0xbb, 0x0a, 0x07, 0xf5, 0xb1, 0x17,
	0xd0, 0xe5, 0x99, 0xf2, 0x90, 0x64, 0x8e, 0x9e, 0x27, 0xc9, 0x1c, 0xfb, 0x29, 0x94, 0x2e, 0xa9,
	0xb8, 0xc7, 0x75, 0x05, 0x11, 0x61, 0xb0, 0x0e, 0xfd, 0xd9, 0x84, 0xbf, 0x97, 0xb1, 0x72, 0xe9,
	0x0b, 0x48, 0xf4, 0x1a, 0xd2, 0x3e, 0xf7, 0xfc, 0xa0, 0x98, 0x54, 0xbb, 0x7c, 0xa2, 0x76, 0xb9,
	0xb2, 0xb5, 0xd6, 0x91, 0xb8, 0xa6, 0x2b, 0x78, 0x84, 0xe3, 0x1e, 0x74, 0x02, 0x29, 0x11, 0xf9,
	0x54, 0x49, 0x91, 0xd7, 0xb6, 0x5a, 0xdd, 0xdb, 0x8b, 0x7c, 0x8a, 0x15, 0x18, 0x9d, 0x41, 0x26,
	0x50, 0x15, 0xa5, 0x53, 0xfe, 0xf8, 0xff, 0x47, 0xda, 0xf4, 0x59, 0x37, 0x48, 0xf1, 0x29, 0xe7,
	0x93, 0x60, 0xa4, 0x74, 0xb4, 0xb0, 0x8e, 0x4a, 0xa7, 0x00, 0xf7, 0xc3, 0xa1, 0x02, 0x98, 0x1f,
	0x69, 0xa4, 0x7f, 0xa3, 0x3c, 0xae, 0xfe, 0xb4, 0x5e, 0x25, 0x4f, 0x0d, 0xbb, 0x0c, 0x29, 0x39,
	0x1a, 0xca, 0x82, 0xd9, 0xab, 0x77, 0x0a, 0x09, 0x04, 0x90, 0xe9, 0x36, 0xf1, 0xd5, 0xf9, 0x75,
	0xc1, 0xb0, 0x1b, 0x90, 0x89, 0x07, 0x40, 0x16, 0xa4, 0x9b, 0x18, 0xb7, 0x71, 0x0c, 0xa8, 0x5f,
	0xb7, 0xbb, 0xcd, 0x46, 0xc1, 0x40, 0x5b, 0x90, 0xed, 0x34, 0x5b, 0x8d, 0xab, 0xd6, 0x65, 0x21,
	0x29, 0x83, 0x77, 0xad, 0x37, 0xad, 0xf6, 0xfb, 0x56, 0xc1, 0x94, 0x0d, 0xb8, 0x79, 0xde, 0xf8,
	0x50, 0x48, 0x1d, 0xff, 0x34, 0x21, 0xf3, 0x56, 0x3d, 0xb6, 0xe8, 0x04, 0xb2, 0xfa, 0xd5, 0x41,
	0x7b, 0xf3, 0x6f, 0x90, 0x5a, 0x63, 0x69, 0x77, 0xe9, 0x61, 0xb2, 0x13, 0xe8, 0x05, 0xe4, 0xa6,
	0xde, 0x42, 0xfb, 0x0b, 0x56, 0x8b, 0xdb, 0xd0, 0xb2, 0x01, 0xed, 0x04, 0xea, 0xc0, 0xee, 0x92,
	0x73, 0xd0, 0x91, 0x82, 0xae, 0x73, 0x54, 0xa9, 0xb4, 0x7e, 0x17, 0x76, 0x02, 0xdd, 0xc0, 0xfe,
	0x2a, 0x1b, 0xa1, 0xca, 0x8c, 0x74, 0x8d, 0xc3, 0x36, 0xf0, 0x5e, 0xc3, 0xce, 0x82, 0x0b, 0x50,
	0x39, 0x6e, 0x58, 0xe9, 0x8d, 0x0d, 0x6c, 0x3d, 0xd8, 0x5b, 0xe1, 0x14, 0x14, 0x7f, 0x9d, 0xeb,
	0x3d, 0xf4, 0x38, 0xeb, 0x20, 0xa3, 0xfe, 0x28, 0x4f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x11,
	0x2d, 0xdd, 0x42, 0x38, 0x07, 0x00, 0x00,
}
