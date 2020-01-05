package watch

import (
	"github.com/hitachi-vantara-edge/modbus-service/internal/config"
	"github.com/rs/zerolog/log"
	"gopkg.in/fsnotify/fsnotify.v1"
)

func WatchServiceConfigFile(svcConfigPath string, configurator config.Configurator) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error().Err(err).Msg("Error creating watcher for service configuration")
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Error().Msg("Unable to receive Event from channel")
					continue
				}
				log.Debug().Msgf("config changed: %v", event)
				serviceConfig, err := config.GetServiceConfigurationFromFile(svcConfigPath)
				if err != nil {
					log.Error().Err(err).Msg("Error reading service configuration from file")
					continue
				}
				configurator.UpdateConfig(serviceConfig)
			}
		}
	}()

	err = watcher.Add(svcConfigPath)
	if err != nil {
		log.Error().Err(err).Msg("Error setting file watcher")
	}
}
