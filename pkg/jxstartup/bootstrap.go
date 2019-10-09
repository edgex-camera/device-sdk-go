// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package jxstartup

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/edgexfoundry/device-sdk-go"
	"github.com/edgexfoundry/device-sdk-go/internal/common"
	"github.com/edgexfoundry/device-sdk-go/internal/remote"
	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

var (
	confProfile string
	confDir     string
	useRegistry string = "consul://172.17.0.1:8500"

	Service *device.Service
	err     error
)

func init() {
	flag.StringVar(&useRegistry, "registry", useRegistry, "Indicates the service should use the registry and provide the registry url.")
	flag.StringVar(&useRegistry, "r", useRegistry, "Indicates the service should use registry and provide the registry path.")
	flag.StringVar(&confProfile, "profile", confProfile, "Specify a profile other than default.")
	flag.StringVar(&confProfile, "p", confProfile, "Specify a profile other than default.")
	flag.StringVar(&confDir, "confdir", confDir, "Specify an alternate configuration directory.")
	flag.StringVar(&confDir, "c", confDir, "Specify an alternate configuration directory.")
}

func StartService(serviceName string, serviceVersion string, driver dsModels.ProtocolDriver) error {
	if !flag.Parsed() {
		flag.Parse()
	}

	Service, err = device.NewService(serviceName, serviceVersion, confProfile, confDir, useRegistry, driver)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Calling service.Start.\n")
	errChan := make(chan error, 2)
	go listenForInterrupt(errChan)
	err = Service.Start(errChan)
	if err != nil {
		return err
	}
	go listenForConfigChanges()
	go updateDeviceLocation()

	err = <-errChan
	fmt.Fprintf(os.Stdout, "Terminating: %v.\n", err)

	return Service.Stop(false)
}

func listenForInterrupt(errChan chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errChan <- fmt.Errorf("%s", <-c)
}

func updateDeviceLocation() {
	nodeInfo, err := remote.GetNodeInfo()
	if err != nil {
		common.LoggingClient.Error("Failed to get node info from gateway")
		return
	}
	location := map[string]string{"nodeid": nodeInfo.WorkId}

	for _, d := range Service.Devices() {
		if d.Location == nil {
			d.Location = location
			err := Service.UpdateDevice(d)

			var msg string
			if err != nil {
				msg = fmt.Sprintf("Device location update failed: %s", err)
			} else {
				msg = fmt.Sprintf("Device location updated: %s", location)
			}
			common.LoggingClient.Info(msg)
		}
	}
}
