package main

import (
	"fmt"
	"biostar/service/action"
	"example/cli"
)


func testAction(deviceID uint32) error {
	origConfig, err := actionSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("Original config: %v\n", origConfig)

	testEventTrigger(deviceID)

	err = actionSvc.SetConfig(deviceID, origConfig)
	if err != nil {
		return err
	}
	
	return nil
}


const (
	BS2_EVENT_VERIFY_FAIL = 0x1100
	BS2_EVENT_IDENTIFY_FAIL = 0x1400

	BS2_SUB_EVENT_CREDENTIAL_CARD = 0x02
	BS2_SUB_EVENT_CREDENTIAL_FINGER = 0x04

	FAIL_SIGNAL_COUNT = 3
	ON_DURATION_MS = 500
	OFF_DURATION_MS = 500

	RELAY_INDEX = 0 
)

func testEventTrigger(deviceID uint32) error {
	cardFailTrigger := makeEventTrigger(deviceID, BS2_EVENT_VERIFY_FAIL | BS2_SUB_EVENT_CREDENTIAL_CARD)
	fingerFailTrigger := makeEventTrigger(deviceID, BS2_EVENT_IDENTIFY_FAIL | BS2_SUB_EVENT_CREDENTIAL_FINGER)
	
	failSignal := &action.Signal{
		Count: FAIL_SIGNAL_COUNT,
		OnDuration: ON_DURATION_MS,
		OffDuration: OFF_DURATION_MS,
	}

	failRelayAction := makeRelayAction(deviceID, RELAY_INDEX, failSignal)

	testConfig := &action.TriggerActionConfig{
		TriggerActions: []*action.TriggerActionConfig_TriggerAction{
			&action.TriggerActionConfig_TriggerAction{
				Trigger: cardFailTrigger, Action: failRelayAction,
			},
			&action.TriggerActionConfig_TriggerAction{
				Trigger: fingerFailTrigger, Action: failRelayAction,
			},
		},
	}

	err := actionSvc.SetConfig(deviceID, testConfig)
	if err != nil {
		return err
	}
	
	config, err := actionSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("Test Config: %v\n", config)

	fmt.Printf("\n===== Trigger & Action Test =====\n\n")	
	fmt.Printf(">> Try to authenticate a unregistered card or finger. It should trigger a relay signal.\n")
	cli.PressEnter(">> Press ENTER if you finish the test.\n")
	
	return nil
}