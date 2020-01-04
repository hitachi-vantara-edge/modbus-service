package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Configurator struct {
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	LogLevel string `yaml:"logLevel"`
}

func NewConfigurator(serviceConfig *ServiceConfig) Configurator {

	// Set log level
	logLevel, err := zerolog.ParseLevel(serviceConfig.LogLevel)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse logLevel setting to default INFO")
		logLevel = zerolog.InfoLevel
	}
	if logLevel == zerolog.NoLevel {
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	return Configurator{ServiceConfig: serviceConfig}
}

func (c Configurator) UpdateConfig(serviceConfig ServiceConfig) {

	// Update log level
	if c.ServiceConfig.LogLevel != serviceConfig.LogLevel {
		logLevel, err := zerolog.ParseLevel(serviceConfig.LogLevel)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse logLevel setting to default INFO")
			logLevel = zerolog.InfoLevel
		}
		if logLevel == zerolog.NoLevel {
			logLevel = zerolog.InfoLevel
		}
		zerolog.SetGlobalLevel(logLevel)
		c.ServiceConfig.LogLevel = zerolog.GlobalLevel().String()
		log.Debug().Str("Log level", c.ServiceConfig.LogLevel).Msg("Updated log level")
	}
}

func GetServiceConfigurationFromFile(svcConfigPath string) (ServiceConfig, error) {
	var serviceConfig ServiceConfig

	source, err := ioutil.ReadFile(filepath.FromSlash(svcConfigPath))

	if err != nil {
		return ServiceConfig{}, err
	}
	err = yaml.Unmarshal(source, &serviceConfig)

	if err != nil {
		return ServiceConfig{}, err
	}
	return serviceConfig, nil
}
