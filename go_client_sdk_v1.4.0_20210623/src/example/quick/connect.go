package main

import (
	"fmt"
)


func testConnect() (uint32, error) {
	devList, err := connectSvc.GetDeviceList()

	if err != nil {
		return 0, err
	}

	fmt.Printf("Device list before connection: %v\n\n", devList)

	deviceID, err := connectSvc.Connect(A2_IP, A2_PORT, USE_SSL)

	if err != nil {
		return 0, err
	}	

	devList, err = connectSvc.GetDeviceList()

	if err != nil {
		return 0, err
	}

	fmt.Printf("Device list after connection %v\n\n", devList)

	return deviceID, nil
}


func testConnectMaster(gatewayID string) (uint32, error) {
	devList, err := connectMasterSvc.GetDeviceList(gatewayID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Device list before connection: %v\n\n", devList)

	deviceID, err := connectMasterSvc.Connect(gatewayID, A2_IP, A2_PORT, USE_SSL)

	if err != nil {
		return 0, err
	}	

	devList, err = connectMasterSvc.GetDeviceList(gatewayID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Device list after connection %v\n\n", devList)

	return deviceID, nil
}
