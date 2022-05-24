package main

import(
	"fmt"
	"biostar/service/apbZone"
	"biostar/service/rs485"
	"biostar/service/action"
	"example/cli"
)

const (
	TEST_ZONE_ID = 1
)

func testAPB(deviceID uint32, slaves []*rs485.SlaveDeviceInfo) error {
	// backup the existing zones
	origZones, err := apbZoneSvc.Get(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("Original APB Zones: %v\n", origZones)
	apbZoneSvc.DeleteAll(deviceID)

	// make the test zone
	zoneInfo := makeAPBZone(deviceID, slaves)
	err = apbZoneSvc.Add(deviceID, []*apbZone.ZoneInfo{ zoneInfo })
	if err != nil {
		return err
	}

	fmt.Printf("\n===== Anti Passback Zone Test =====\n\n")
	fmt.Printf("Test Zone: %v\n\n", zoneInfo)

	fmt.Printf(">> Authenticate a regsistered credential on the entry device(%v) and/or the exit device(%v) to test if the APB zone works correctly.\n", deviceID, slaves[0].DeviceID)
	cli.PressEnter(">> Press ENTER for the next test.\n")

	cli.PressEnter(">> Press ENTER after generating an APB violation.\n")
	
	apbZoneSvc.ClearAll(deviceID, TEST_ZONE_ID)

	fmt.Printf(">> The APB records are cleared. Try to authenticate the last credential which caused the APB violation. It should succeed since the APB records are cleared.\n")
	cli.PressEnter(">> Press ENTER to finish the test.\n")

	// restore the existing zones
	apbZoneSvc.DeleteAll(deviceID)
	if len(origZones) > 0 {
		apbZoneSvc.Add(deviceID, origZones)
	} 

	return nil
}

const (
	RELAY_ACTION_ON_MS = 500
	RELAY_ACTION_OFF_MS = 500
)

func makeAPBZone(deviceID uint32, slaves []*rs485.SlaveDeviceInfo) *apbZone.ZoneInfo {
  // Make a zone with the master device and the 1st slave device

	return &apbZone.ZoneInfo{
		ZoneID: TEST_ZONE_ID,
		Name: "Test APB Zone",
		Type: apbZone.Type_HARD,
		ResetDuration: 0, // indefinite
		Members: []*apbZone.Member{
			&apbZone.Member{
				DeviceID: deviceID,
				ReaderType: apbZone.ReaderType_ENTRY,
			},
			&apbZone.Member{
				DeviceID: slaves[0].DeviceID,
				ReaderType: apbZone.ReaderType_EXIT,
			},
		},
		Actions: []*action.Action{ // Activate the 1st relay of the master device when an alarm is detected
			&action.Action{
				DeviceID: deviceID,
				Type: action.ActionType_ACTION_RELAY,
				Relay: &action.RelayAction{
					RelayIndex: 0, // 1st relay of the master device
					Signal: &action.Signal{ 
						Count: 3,
						OnDuration: RELAY_ACTION_ON_MS,
						OffDuration: RELAY_ACTION_OFF_MS,
					},
				},
			},
		},
	}
}