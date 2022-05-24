package main

import(
	"fmt"

	"biostar/service/door"
	"biostar/service/device"
)

const (
	RELAY_PORT = 0 // 1st relay
	SENSOR_PORT = 0 // 1st input port
	EXIT_BUTTON_PORT = 1 // 2nd input port
	AUTO_LOCK_TIMEOUT = 3 // locked after 3 seconds
	HELD_OPEN_TIMEOUT = 10 // held open alarm after 10 seconds
)

func makeSingleDoor(deviceID, doorID uint32) *door.DoorInfo {
	singleDoor := &door.DoorInfo{
		DoorID: doorID,
		Name: fmt.Sprintf("Test Door %v", doorID),
		EntryDeviceID: deviceID,
		Relay: &door.Relay{
			DeviceID: deviceID,
			Port: RELAY_PORT,
		},
		Sensor: &door.Sensor{
			DeviceID: deviceID,
			Port: SENSOR_PORT,
			Type: device.SwitchType_NORMALLY_OPEN,
		},
		Button: &door.ExitButton{
			DeviceID: deviceID,
			Port: EXIT_BUTTON_PORT,
			Type: device.SwitchType_NORMALLY_OPEN,
		},
		AutoLockTimeout: AUTO_LOCK_TIMEOUT,
		HeldOpenTimeout: HELD_OPEN_TIMEOUT,
	}

	return singleDoor
}


func backupDoor(deviceID uint32) ([]*door.DoorInfo, error) {
	doors, err := doorSvc.GetList(deviceID)

	if err != nil {
		return nil, err
	}

	return doors, nil
}

func cleanUpDoor(deviceID uint32) error {
	err := doorSvc.DeleteAll(deviceID)
	if err != nil {
		return err
	}

	return nil
}

func restoreDoor(deviceID uint32, doors []*door.DoorInfo) error {
	err := cleanUpDoor(deviceID)
	if err != nil {
		return err
	}

	err = doorSvc.Add(deviceID, doors)
	if err != nil {
		return err
	} 

	return nil
}