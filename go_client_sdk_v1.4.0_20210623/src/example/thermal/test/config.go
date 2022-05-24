package main

import (
	"fmt"
	"biostar/service/thermal"
	"example/cli"
	"github.com/golang/protobuf/proto"
)


func testConfig(deviceID uint32, config *thermal.ThermalConfig) error {
	// Backup the original configuration
	origConfig := proto.Clone(config)

	// Set options for the test
	config.AuditTemperature = true // write temperature logs
	config.CheckMode = thermal.CheckMode_HARD // disalllow access when temperature is too high

	// (1) Set check order to AFTER_AUTH
	config.CheckOrder = thermal.CheckOrder_AFTER_AUTH

	err := thermalSvc.SetConfig(deviceID, config)
	if err != nil {
		return err
	}

	fmt.Printf("\n===== Test for ThermalConfig =====\n\n")

	fmt.Printf("(1) The Check Order is set to AFTER_AUTH. The device will measure the temperature only after successful authentication. Try to authenticate faces.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")

	// (2) Set check order to BEFORE_AUTH
	config.CheckOrder = thermal.CheckOrder_BEFORE_AUTH

	err = thermalSvc.SetConfig(deviceID, config)
	if err != nil {
		return err
	}

	fmt.Printf("(2) The Check Order is set to BEFORE_AUTH. The device will try to authenticate a user only when the user's temperature is within the threshold. Try to authenticate faces.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")

	// (3) Set check order to WITHOUT_AUTH
	config.CheckOrder = thermal.CheckOrder_WITHOUT_AUTH

	err = thermalSvc.SetConfig(deviceID, config)
	if err != nil {
		return err
	}

	fmt.Printf("(3) The Check Order is set to WITHOUT_AUTH. Any user whose temperature is within the threshold will be allowed to access. Try to authenticate faces.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")

	// (4) Set check order to AFTER_AUTH with too low threshold
	config.CheckOrder = thermal.CheckOrder_AFTER_AUTH
	config.TemperatureThreshold = 3500 // Too low threshold. Most temperature check will fail

	err = thermalSvc.SetConfig(deviceID, config)
	if err != nil {
		return err
	}

	fmt.Printf("(4) To reproduce the case of high temperature, the Check Order is set to AFTER_AUTH with the threshold of 35 degree Celsius. Most temperature check will fail, now. Try to authenticate faces.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")	

	// Restore the original configuration
	err = thermalSvc.SetConfig(deviceID, origConfig.(*thermal.ThermalConfig))
	if err != nil {
		return err
	}

	return nil
}


