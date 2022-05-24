package main

import (
	"fmt"
	"biostar/service/tna"
	"biostar/service/event"
	"example/cli"
)


func testConfig(deviceID uint32) (*tna.TNAConfig, error) {
	eventSvc.StartMonitoring(deviceID)
	eventSvc.SetEventCallback(eventCallback)

	// backup the original configuration
	origConfig, err := tnaSvc.GetConfig(deviceID)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Original Config: %v\n\n", origConfig)

	// (1) BY_USER
	newConfig := &tna.TNAConfig{
		Mode: tna.Mode_BY_USER,
		Labels: []string{
			"In",
			"Out",
			"Scheduled In",
			"Fixed Out",
		},
	}

	err = tnaSvc.SetConfig(deviceID, newConfig)
	if err != nil {
		return origConfig, err
	}

	fmt.Printf("\n===== Test for TNAConfig =====\n\n")

	fmt.Printf("(1) The T&A mode is set to BY_USER(optional). You can select a T&A key before authentication. Try to authenticate after selecting a T&A key.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")

	// (2) IsRequired
	newConfig.IsRequired = true

	err = tnaSvc.SetConfig(deviceID, newConfig)
	if err != nil {
		return origConfig, err
	}

	fmt.Printf("(2) The T&A mode is set to BY_USER(mandatory). Try to authenticate without selecting a T&A key.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")

	// (3) LAST_CHOICE
	newConfig.Mode = tna.Mode_LAST_CHOICE

	err = tnaSvc.SetConfig(deviceID, newConfig)
	if err != nil {
		return origConfig, err
	}

	fmt.Printf("(3) The T&A mode is set to LAST_CHOICE. The T&A key selected by the previous user will be used. Try to authenticate multiple users.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")	

	// (4) BY_SCHEDULE
	newConfig.Mode = tna.Mode_BY_SCHEDULE
	newConfig.Schedules = []uint32{ 0, 0, 1 } // Always for KEY_3 (Scheduled In)

	err = tnaSvc.SetConfig(deviceID, newConfig)
	if err != nil {
		return origConfig, err
	}

	fmt.Printf("(4) The T&A mode is set to BY_SCHEDULE. The T&A key will be determined automatically by schedule. Try to authenticate without selecting a T&A key.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")


	// (5) FIXED
	newConfig.Mode = tna.Mode_FIXED
	newConfig.Key = tna.Key_KEY_4

	err = tnaSvc.SetConfig(deviceID, newConfig)
	if err != nil {
		return origConfig, err
	}

	fmt.Printf("(5) The T&A mode is set to FIXED(KEY_4). Try to authenticate without selecting a T&A key.\n\n")
	cli.PressEnter(">> Press ENTER if you finish testing this mode.\n")

	return origConfig, nil
}

func eventCallback(eventLog *event.EventLog) {
	if firstEventID == 0 {
		firstEventID = eventLog.ID
	}

	fmt.Printf("	Realtime Event: %v\n", eventLog)
}

