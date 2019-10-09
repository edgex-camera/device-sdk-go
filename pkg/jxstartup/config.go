package jxstartup

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/edgexfoundry/device-sdk-go/internal/common"
	"github.com/edgexfoundry/device-sdk-go/internal/config"
)

type ConfigChangeListener interface {
	OnConfigChange(old map[string]string, new map[string]string)
}

const driverKey = "Driver"

func PutDriverConfig(name string, value []byte) error {
	return config.RegistryClient.PutConfigurationValue(driverKey+"/"+name, value)
}

func listenForConfigChanges(listenerCandidate interface{}) {
	errChannel := make(chan error)
	updateChannel := make(chan interface{})

	// common.CurrentConfig.Driver need to be updated manually,
	// WatchForChanges will not update for us
	config.RegistryClient.WatchForChanges(updateChannel, errChannel, &map[string]string{}, driverKey)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			// Quietly and gracefully stop when SIGINT/SIGTERM received
			return
		case ex := <-errChannel:
			common.LoggingClient.Error(ex.Error())
		case raw, ok := <-updateChannel:
			if ok {
				newDriver, ok := raw.(*map[string]string)
				if !ok {
					common.LoggingClient.Error("listenForConfigChanges() type check failed")
				}
				if listener, ok := listenerCandidate.(ConfigChangeListener); ok {
					listener.OnConfigChange(common.CurrentConfig.Driver, *newDriver)
				}
				common.CurrentConfig.Driver = *newDriver
				common.LoggingClient.Info("Driver configuration has been updated.")
			} else {
				return
			}
		}
	}
}
