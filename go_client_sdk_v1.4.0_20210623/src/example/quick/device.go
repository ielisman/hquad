package main

import (
	"fmt"
	"biostar/service/device"
)


func testDevice(deviceID uint32) (*device.CapabilityInfo, error) {
	devInfo, err := deviceSvc.GetInfo(deviceID)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Device info: %v\n\n", devInfo)

	capInfo, err := deviceSvc.GetCapabilityInfo(deviceID)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Device capability info: %v\n\n", capInfo)

	return capInfo, nil
}

