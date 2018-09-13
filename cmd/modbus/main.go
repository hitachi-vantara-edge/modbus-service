package main

import (
	"github.com/hitachi-vantara-edge/modbus-service/internal/modbus"
	"os"
	"strconv"
)

func main() {

	port, _ := strconv.Atoi(os.Getenv("MODBUS_PORT"))
	//addr := os.Getenv("MODBUS_ADDR")

	if port == 0 {
		port = 5440
	}

	server := modbus.NewModbusServer("127.0.0.1", port)

	server.Run()

}




