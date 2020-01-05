package main

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hitachi-vantara-edge/modbus-service/internal/config"
	"github.com/hitachi-vantara-edge/modbus-service/internal/modbus"
	"github.com/hitachi-vantara-edge/modbus-service/internal/watch"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	ServiceConfigFilePath = "/res/service-config/config.yaml"
)

func main() {
	// RFC3339Nano has format "2006-01-02T15:04:05.999999999Z07:00"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.With().Caller().Logger()

	var serviceConfigPath = filepath.FromSlash(ServiceConfigFilePath)
	serviceConfig, err := config.GetServiceConfigurationFromFile(serviceConfigPath)
	if err != nil {
		log.Error().Err(err).Msg("Unable to get service configurations")
	}

	configurator := config.NewConfigurator(&serviceConfig)

	watch.WatchServiceConfigFile(serviceConfigPath, configurator)

	port, _ := strconv.Atoi(os.Getenv("MODBUS_PORT"))
	// addr := os.Getenv("MODBUS_ADDR")

	if port == 0 {
		port = 5440
	}

	server := modbus.NewModbusServer("127.0.0.1", port)

	server.Run()

}
