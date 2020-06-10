package modbus

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	modbus "github.com/hitachi-vantara-edge/modbus-service/api"
	"github.com/hitachi-vantara-edge/modbus-service/pkg/libmodbus"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const gRPCServerPort = 10000

type modbusServer struct {
	connections                    map[string]*modbusConnection
	maxConnectionIdleTimeInMinutes int
}

type modbusConnection struct {
	Context               []libmodbus.T
	Id                    string
	Props                 map[string]string
	Type                  modbus.ConnectionStatusReply_Type
	Status                modbus.ConnectionStatusReply_Status
	ErrorMsg              string
	LastSuccessfulRequest time.Time
}

func NewModbusServer(addr string, port int, maxConnIdleTimeInMins int) *modbusServer {
	return &modbusServer{connections: map[string]*modbusConnection{}, maxConnectionIdleTimeInMinutes: maxConnIdleTimeInMins}
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
		connection.LastSuccessfulRequest = time.Now()

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
		connection.LastSuccessfulRequest = time.Now()

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
	log.Info().Msgf("OpenTCPConnection, address: %s, port: %v", req.Addr, req.Port)

	mctx := libmodbus.NewTcp(req.Addr, req.Port)

	reply := &modbus.ConnectionStatusReply{}

	props := make(map[string]string)
	props["addr"] = req.Addr
	props["port"] = strconv.Itoa(int(req.Port))
	id := uuid.New()
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
		connection.LastSuccessfulRequest = time.Now()
		log.Info().Msgf("OpenTCPConnection succeed, address: %s, port: %v, connectionId: %s", req.Addr, req.Port, connection.Id)
	}
	reply.Status = connection.Status

	s.connections[connection.Id] = connection
	return reply, nil
}

func (s *modbusServer) OpenSerialConnection(ctx context.Context, req *modbus.OpenSerialConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
	return &modbus.ConnectionStatusReply{}, nil
}

//ToDo: revisit workaround
//Note: Calling 'Free' from modbus.go after 'Close' caused exception/crash
/*
	//*** Error in `./main': free(): invalid size: 0x000000c4200a5d40 ***
	//
	//SIGABRT: abort
	//PC=0x7f6a13343fff m=11 sigcode=18446744073709551610
	//
	//goroutine 0 [idle]:
	//runtime: unknown pc 0x7f6a13343ff
*/
//- Tried updating libmodbus to support a new method: libmodbus.CloseAndFree (i.e. modbus_close_and_free in modbus.c) to do both modbus_close and modbus_free simultaneously, did not resolve the crash issue
//- Workaround for now: do not 'Free' - need to further investigate potential memory leak
func (s *modbusServer) CloseConnection(ctx context.Context, req *modbus.CloseConnectionRequest) (*modbus.
	ConnectionStatusReply, error) {
	log.Info().Msgf("CloseConnection, connectionId: %s", req.Id)
	if connection, ok := s.connections[req.Id]; ok {
		if connection.Context != nil {
			if connection.Type == modbus.ConnectionStatusReply_TCP {
				libmodbus.Close(connection.Context)
			}
			//libmodbus.Free(connection.Context)
			connection.Status = modbus.ConnectionStatusReply_CLOSED
			log.Info().Msgf("CloseConnection succeed, connectionId: %s", req.Id)
		}
		delete(s.connections, req.Id)
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

	// Periodic cleanup of modbus connection to device, e.g. in case of modbus-adapter crashed/restarted hence could not close connections properly by making CloseConnection calls to this modbus-service
	var ticker *time.Ticker
	var done chan bool
	if s.maxConnectionIdleTimeInMinutes > 0 {
		maxIdleDuration := time.Duration(s.maxConnectionIdleTimeInMinutes) * time.Minute
		ticker = time.NewTicker(maxIdleDuration)
		done = make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					for _, connection := range s.connections {
						if t.Sub(connection.LastSuccessfulRequest) > maxIdleDuration {
							log.Info().Msgf("Closing idle connection, connectionId: %s, lastSuccessfulRequest: %+v, currentTickTime: %+v, maxConnectionIdleTimeInMinutes: %+v", connection.Id, connection.LastSuccessfulRequest, t, s.maxConnectionIdleTimeInMinutes)
							s.CloseConnection(context.Background(), &modbus.CloseConnectionRequest{Id: connection.Id})
						}
					}
				}
			}
		}()
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	log.Info().Msgf("Stopping gRPC Server")
	grpcServer.GracefulStop()

	if ticker != nil {
		ticker.Stop()
		done <- true
	}

	log.Info().Msgf("Shutting down Modbus Library Service")
	s.Shutdown()

	os.Exit(0)
}

func (s *modbusServer) Shutdown() {
	for _, connection := range s.connections {
		if connection.Context != nil {
			libmodbus.Close(connection.Context)
			//libmodbus.Free(connection.Context)
		}
	}
}
