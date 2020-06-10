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
	maxConnIdleTimeInMins := getEnvInt("MODBUS_MAX_CONNECTION_IDLE_TIME_IN_MINUTES", 0)

	if port == 0 {
		port = 5440
	}

	server := modbus.NewModbusServer("127.0.0.1", port, maxConnIdleTimeInMins)

	log.Info().Msg("Starting up Modbus Library Service...")

	server.Run()

}

// Function GetEnvInt is os.getenv with default value.
// If environment value is not set for key, it returns default value as integer.
func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Error().Str("key", key).Err(err).Str("value", value).Msg("Unable to parse given value")
		return fallback
	}
	return i
}
