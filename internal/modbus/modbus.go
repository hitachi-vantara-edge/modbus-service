package modbus

/*
#cgo LDFLAGS: -lmodbus
#include <config.h>
#include "../../pkg/libmodbus/modbus-tcp.h"
#include "../../pkg/libmodbus/modbus-private.h"
#include "../../pkg/libmodbus/modbus.h"
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
modbus_t *ctx;
uint16_t tab_reg[65535];
uint8_t coils[65535];
int rc;
int rco;
int i;

static inline uint16_t * read_regs(modbus_t *ctx, int slaveNum, int start, int end, int readInput) {

	modbus_set_slave(ctx, slaveNum);

    if (readInput) {
    	rc = modbus_read_input_registers(ctx, start, end, tab_reg);
    } else {
    	rc = modbus_read_registers(ctx, start, end, tab_reg);
    }

	if (rc == -1) {
		fprintf(stderr, "%s\n", modbus_strerror(errno));
		return tab_reg;
	}

	return tab_reg;
}

static inline uint8_t * read_coils(modbus_t *ctx, int slaveNum, int start, int numOfBits, int readInput) {

	modbus_set_slave(ctx, slaveNum);

	if (readInput) {
		rco = modbus_read_input_bits(ctx, start, numOfBits, coils);
	} else {
		rco = modbus_read_bits(ctx, start, numOfBits, coils);
	}

	if (rco == -1) {
		fprintf(stderr, "%s\n", modbus_strerror(errno));
		return coils;
	}

	return coils;
}

static inline const char * get_error() {
	return modbus_strerror(errno);
}
*/
import "C"
import (
	"context"
	"fmt"
	"github.com/hitachi-vantara-edge/modbus-service/pkg/rpc"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"unsafe"
)

type modbusServer struct {
	connection *modbusConnection
}

type modbusConnection struct {
	Context *C.modbus_t
	Id string
	Props map[string]string
	Type modbus.ConnectionStatusReply_Type
	Status modbus.ConnectionStatusReply_Status
	ErrorMsg string
}

func NewModbusServer(addr string, port int) (*modbusServer) {
	return &modbusServer{connection: &modbusConnection{}}
}

func (s *modbusServer) ReadHoldingReg(ctx context.Context, req *modbus.ReadHoldingRegRequest) (*modbus.
	ReadHoldingRegReply, error) {

	var reg *C.uint16_t = C.read_regs(s.connection.Context, C.int(req.SlaveAddr), C.int(req.StartAddr), C.int(req.NumOfReg),
		C.int(Btoi(req.ReadInput)))
	data := (*[1 << 30]C.uint16_t)(unsafe.Pointer(reg))[:65535:65535]
	values := make([]*modbus.ReadHoldingRegReply_HoldingRegisterValue, 0)
	for i, v := range data {
		ui := uint32(i)
		if ui >= req.NumOfReg {
			break
		}
		addr := ui + req.StartAddr
		fmt.Printf("reg[%v]=%v (0x%X)\n", addr, v, v)
		values = append(values, &modbus.ReadHoldingRegReply_HoldingRegisterValue{
			Addr: uint32(addr),
			Value: fmt.Sprint(v),
		})
	}

	return &modbus.ReadHoldingRegReply{ByteCount: uint32(len(values)), Values: values}, nil
}

func (s *modbusServer) ReadCoil(ctx context.Context, req *modbus.ReadCoilRequest) (*modbus.ReadCoilReply, error) {
	var coils *C.uint8_t = C.read_coils(s.connection.Context, C.int(req.SlaveAddr), C.int(req.StartAddr),
		C.int(req.NumOfCoil), C.int(Btoi(req.ReadInput)))
	data := (*[1 << 30]C.uint8_t)(unsafe.Pointer(coils))[:65535:65535]
	values := make([]*modbus.ReadCoilReply_CoilValue, 0)
	for i, v := range data {
		ui := uint32(i)
		if ui >= req.NumOfCoil {
			break
		}
		addr := ui + req.StartAddr
		fmt.Printf("coil[%v]=%v (0x%X)\n", addr, v, v)
		values = append(values, &modbus.ReadCoilReply_CoilValue{
			Addr: uint32(addr),
			Value: fmt.Sprint(v),
		})
	}
	return &modbus.ReadCoilReply{BitCount: uint32(len(values)), Values: values}, nil
}

func (s *modbusServer) OpenTCPConnection(ctx context.Context, req *modbus.OpenTCPConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
	mctx := C.modbus_new_tcp(C.CString(req.Addr), C.int(req.Port))

	reply := &modbus.ConnectionStatusReply{}

	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	props := make(map[string]string)
	props["addr"] = req.Addr
	props["port"] = strconv.Itoa(int(req.Port))
	id, err := sf.NextID()
	if err != nil {
		return nil, err
	}
	s.connection = &modbusConnection{
		Context: mctx,
		Id: fmt.Sprint(id),
		Props: props,
		Type: modbus.ConnectionStatusReply_TCP,
	}

	reply.Id = s.connection.Id
	reply.Props = s.connection.Props
	reply.Type = s.connection.Type
	reply.Errmsg = s.connection.ErrorMsg

	if C.modbus_connect(mctx) == -1 {
		error := C.get_error()
		fmt.Errorf("Connection failed: %s\n", C.GoString(error))
		s.connection.ErrorMsg = fmt.Sprint(C.GoString(error))
		s.connection.Status = modbus.ConnectionStatusReply_ERROR
		return reply, fmt.Errorf("Connection failed: %s\n", C.GoString(error))
	} else {
		s.connection.Status = modbus.ConnectionStatusReply_READY
	}
	reply.Status = s.connection.Status

	return reply, nil
}

func (s *modbusServer) OpenSerialConnection(ctx context.Context, req *modbus.OpenSerialConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
		return &modbus.ConnectionStatusReply{}, nil
}

func (s *modbusServer) CloseConnection(ctx context.Context, req *modbus.CloseConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
		if s.connection.Context != nil {
			if s.connection.Type == modbus.ConnectionStatusReply_TCP {
				C.modbus_close(s.connection.Context)
			}
			C.modbus_free(s.connection.Context)
			s.connection.Status = modbus.ConnectionStatusReply_CLOSED
		}
		return &modbus.ConnectionStatusReply{
			Id: s.connection.Id,
			Props: s.connection.Props,
			Type: s.connection.Type,
			Status: s.connection.Status,
			Errmsg: s.connection.ErrorMsg,
		}, nil

}

func (s *modbusServer) GetConnectionStatus(ctx context.Context, req *modbus.GetConnectionStatusRequest) (*modbus.
	ConnectionStatusReply, error) {
	return &modbus.ConnectionStatusReply{
		Id: s.connection.Id,
		Props: s.connection.Props,
		Type: s.connection.Type,
		Status: s.connection.Status,
		Errmsg: s.connection.ErrorMsg,
	}, nil
}

func (s *modbusServer) Run() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	modbus.RegisterModbusServer(grpcServer, s)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
		if err != nil {
			fmt.Println(err)
		}
		grpcServer.Serve(lis)
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	grpcServer.GracefulStop()

	s.Shutdown()

	os.Exit(0)
}

func (s *modbusServer) Shutdown() {
	if s.connection.Context != nil {
		C.modbus_close(s.connection.Context)
		C.modbus_free(s.connection.Context)
	}
}

func Btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
