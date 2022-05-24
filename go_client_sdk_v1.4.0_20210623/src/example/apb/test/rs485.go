package main

import(
	"fmt"
	"biostar/service/rs485"
	"time"
)

const (
	MAX_WAIT_TRY = 10
	WAIT_SEC = 2
)

func checkSlaves(deviceID uint32) ([]*rs485.SlaveDeviceInfo, []*rs485.SlaveDeviceInfo, bool) {
	config, err := rs485Svc.GetConfig(deviceID)

	rs485Master := false

	for _, ch := range config.Channels {
		if ch.Mode == rs485.Mode_MASTER {
			rs485Master = true
			break
		}
	}

	if !rs485Master {
		fmt.Printf("!! Only a master device can have slave devices. Skip the test.\n")
		return nil, nil, false
	}

	slaves, err := rs485Svc.SearchSlave(deviceID)

	if err != nil || len(slaves) == 0 {
		fmt.Printf("!! No slave device is configured. Configure and wire the slave devices first.\n")
		return nil, nil, false
	}

	fmt.Printf("Found Slaves: %v\n", slaves)

	registeredSlaves, _ := rs485Svc.GetSlave(deviceID)

	fmt.Printf("Registered Slaves: %v\n", registeredSlaves)

	if len(registeredSlaves) == 0 {
		slaves[0].Enabled = true
		rs485Svc.SetSlave(deviceID, slaves)
	}

	for i := 0; i < MAX_WAIT_TRY; i++ {
		newSlaves, _ := rs485Svc.SearchSlave(deviceID)
		if newSlaves[0].Connected {
			fmt.Printf("Test Slaves: %v\n", newSlaves)
			break
		}

		fmt.Printf("Waiting for the slave to be connected %v...\n", i)
		time.Sleep(2 * time.Second)
	}

	return slaves, registeredSlaves, true
}

