package modbus

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"

	modbus "github.com/hitachi-vantara-edge/modbus-service/api"
	"github.com/hitachi-vantara-edge/modbus-service/pkg/libmodbus"
	"github.com/rs/zerolog/log"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc"
)

const gRPCServerPort = 10000

type modbusServer struct {
	connections map[string]*modbusConnection
}

type modbusConnection struct {
	Context  []libmodbus.T
	Id       string
	Props    map[string]string
	Type     modbus.ConnectionStatusReply_Type
	Status   modbus.ConnectionStatusReply_Status
	ErrorMsg string
}

func NewModbusServer(addr string, port int) *modbusServer {
	return &modbusServer{connections: map[string]*modbusConnection{}}
}

func (s *modbusServer) ReadReg(ctx context.Context, req *modbus.ReadRegRequest) (*modbus.
	ReadRegReply, error) {
	if connection, ok := s.connections[req.ConnectionId]; ok {
		var ret int32
		var data = make([]uint16, 65535)
		ret = libmodbus.SetSlave(connection.Context, int32(req.SlaveAddr))
		if ret == -1 {
			return nil, fmt.Errorf("fail to set slave addr. ret=%v", ret)
		}
		if req.ReadInput {
			ret = libmodbus.ReadInputRegisters(connection.Context, req.StartAddr, req.NumOfReg, data)
		} else {
			ret = libmodbus.ReadRegisters(connection.Context, req.StartAddr, req.NumOfReg, data)
		}
		if ret == -1 {
			return nil, fmt.Errorf("fail to read register. ret=%v", ret)
		}

		values := make([]*modbus.ReadRegReply_RegisterValue, 0)
		for i, v := range data {
			if i >= int(req.NumOfReg) {
				break
			}
			addr := i + int(req.StartAddr)
			log.Debug().Msgf("reg[%v]=%v (0x%X)\n", addr, v, v)
			values = append(values, &modbus.ReadRegReply_RegisterValue{
				Addr:  int32(addr),
				Value: fmt.Sprint(v),
			})
		}

		return &modbus.ReadRegReply{ByteCount: int32(len(values)), IsInput: req.ReadInput, Values: values}, nil
	}

	err := errors.New("no modbus connection Id: " + req.ConnectionId)
	log.Error().Err(err)
	return nil, err
}

func (s *modbusServer) ReadCoil(ctx context.Context, req *modbus.ReadCoilRequest) (*modbus.ReadCoilReply, error) {
	if connection, ok := s.connections[req.ConnectionId]; ok {
		var ret int32
		var data = make([]byte, 65535)
		ret = libmodbus.SetSlave(connection.Context, req.SlaveAddr)
		if ret == -1 {
			return nil, fmt.Errorf("fail to set slave addr. ret=%v", ret)
		}
		if req.ReadInput {
			ret = libmodbus.ReadInputBits(connection.Context, req.StartAddr, req.NumOfCoil, data)
		} else {
			ret = libmodbus.ReadBits(connection.Context, req.StartAddr, req.NumOfCoil, data)
		}
		if ret == -1 {
			return nil, fmt.Errorf("fail to read coil. ret=%v", ret)
		}

		values := make([]*modbus.ReadCoilReply_CoilValue, 0)
		for i, v := range data {

			if i >= int(req.NumOfCoil) {
				break
			}
			addr := i + int(req.StartAddr)
			log.Debug().Msgf("coil[%v]=%v (0x%X)\n", addr, v, v)
			values = append(values, &modbus.ReadCoilReply_CoilValue{
				Addr:  int32(addr),
				Value: fmt.Sprint(v),
			})
		}

		return &modbus.ReadCoilReply{BitCount: int32(len(values)), Values: values}, nil
	}

	err := errors.New("no modbus connection Id: " + req.ConnectionId)
	log.Error().Err(err)
	return nil, err
}

func (s *modbusServer) OpenTCPConnection(ctx context.Context, req *modbus.OpenTCPConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {

	mctx := libmodbus.NewTcp(req.Addr, req.Port)

	reply := &modbus.ConnectionStatusReply{}

	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	props := make(map[string]string)
	props["addr"] = req.Addr
	props["port"] = strconv.Itoa(int(req.Port))
	id, err := sf.NextID()
	if err != nil {
		return nil, err
	}
	connection := &modbusConnection{
		Context: []libmodbus.T{*mctx},
		Id:      fmt.Sprint(id),
		Props:   props,
		Type:    modbus.ConnectionStatusReply_TCP,
	}

	reply.Id = connection.Id
	reply.Props = connection.Props
	reply.Type = connection.Type
	reply.Errmsg = connection.ErrorMsg

	rcode := libmodbus.Connect(connection.Context)

	if rcode == -1 {
		errstr := libmodbus.GetErrnoStr()
		connection.ErrorMsg = fmt.Sprint(errstr)
		connection.Status = modbus.ConnectionStatusReply_ERROR
		return reply, fmt.Errorf("Open Connection failed: %s\n", errstr)
	} else {
		connection.Status = modbus.ConnectionStatusReply_READY
	}
	reply.Status = connection.Status

	s.connections[connection.Id] = connection
	return reply, nil
}

func (s *modbusServer) OpenSerialConnection(ctx context.Context, req *modbus.OpenSerialConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
	return &modbus.ConnectionStatusReply{}, nil
}

func (s *modbusServer) CloseConnection(ctx context.Context, req *modbus.CloseConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
	if connection, ok := s.connections[req.Id]; ok {
		if connection.Context != nil {
			if connection.Type == modbus.ConnectionStatusReply_TCP {
				libmodbus.Close(connection.Context)
			}
			libmodbus.Free(connection.Context)
			connection.Status = modbus.ConnectionStatusReply_CLOSED
		}
		return &modbus.ConnectionStatusReply{
			Id:     connection.Id,
			Props:  connection.Props,
			Type:   connection.Type,
			Status: connection.Status,
			Errmsg: connection.ErrorMsg,
		}, nil
	}

	err := errors.New("no modbus connection Id: " + req.Id)
	log.Error().Err(err)
	return &modbus.ConnectionStatusReply{}, err
}

func (s *modbusServer) GetConnectionStatus(ctx context.Context, req *modbus.GetConnectionStatusRequest) (*modbus.
	ConnectionStatusReply, error) {
	if connection, ok := s.connections[req.Id]; ok {
		return &modbus.ConnectionStatusReply{
			Id:     connection.Id,
			Props:  connection.Props,
			Type:   connection.Type,
			Status: connection.Status,
			Errmsg: connection.ErrorMsg,
		}, nil
	}

	err := errors.New("no modbus connection Id: " + req.Id)
	log.Error().Err(err)
	return &modbus.ConnectionStatusReply{}, err
}

func (s *modbusServer) Run() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	modbus.RegisterModbusServer(grpcServer, s)

	log.Info().Msgf("Starting gRPC Listener on port :%d", gRPCServerPort)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCServerPort))
		if err != nil {
			log.Error().Err(err).Msgf("Error on net.Listen tcp address :%d", gRPCServerPort)
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Error().Err(err).Msg("Error on grpc server Serve of listener")
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	log.Info().Msgf("Stopping gRPC Server")
	grpcServer.GracefulStop()

	log.Info().Msgf("Shutting down Modbus Library Service")
	s.Shutdown()

	os.Exit(0)
}

func (s *modbusServer) Shutdown() {
	for _, connection := range s.connections {
		if connection.Context != nil {
			libmodbus.Close(connection.Context)
			libmodbus.Free(connection.Context)
		}
	}
}
