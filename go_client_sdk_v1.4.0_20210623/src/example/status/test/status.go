package main

import (
	"fmt"
	"biostar/service/status"
	"biostar/service/action"
	"biostar/service/device"
	"github.com/golang/protobuf/proto"
	"example/cli"
)


func testStatus(deviceID uint32) error {
	origConfig, err := statusSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("Original Config: %v\n", origConfig)

	// Copy the original configuration
	testConfig := proto.Clone(origConfig)
	testLEDStatus(deviceID, testConfig.(*status.StatusConfig))
	testBuzzerStatus(deviceID, testConfig.(*status.StatusConfig))

	err = statusSvc.SetConfig(deviceID, origConfig)
	if err != nil {
		return err
	} 
	
	return nil
}


func testLEDStatus(deviceID uint32, config *status.StatusConfig) error {
	fmt.Printf("\n===== LED Status Test =====\n\n")

	// change the LED color of the normal status to yellow
	for _, ledStatus := range config.LEDState {
		if ledStatus.DeviceStatus == status.DeviceStatus_DEVICE_STATUS_NORMAL {
			ledStatus.Count = 0 
			ledStatus.Signals = []*action.LEDSignal{
				&action.LEDSignal{
					Color: device.LEDColor_LED_COLOR_YELLOW, 
					Duration: 2000,
					Delay: 0,
				},
			}

			break
		}
	}

	err := statusSvc.SetConfig(deviceID, config)
	if err != nil {
		return err
	}	

	newConfig, err := statusSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("New Config: %v\n\n", newConfig)	

	fmt.Printf(">> The LED color of the normal status is changed to yellow.\n")
	cli.PressEnter(">> Press ENTER for the next test.\n")

	return nil
}

const (
	NUM_OF_FAIL_BUZZER = 2
)

func testBuzzerStatus(deviceID uint32, config *status.StatusConfig) error {
	fmt.Printf("\n===== Buzzer Status Test =====\n\n")

	// change the buzzer signal for FAIL
	for _, buzzerStatus := range config.BuzzerState {
		if buzzerStatus.DeviceStatus == status.DeviceStatus_DEVICE_STATUS_FAIL {
			buzzerStatus.Count = 1 
			buzzerStatus.Signals = []*action.BuzzerSignal{}
			
			for i := 0; i < NUM_OF_FAIL_BUZZER; i++ { // 2 x 500ms beeps
				signal := &action.BuzzerSignal{
					Tone: device.BuzzerTone_BUZZER_TONE_HIGH,
					Duration: 500,
					Delay: 2,
				}
				buzzerStatus.Signals = append(buzzerStatus.Signals, signal)
			}

			break
		}
	}

	err := statusSvc.SetConfig(deviceID, config)
	if err != nil {
		return err
	}	

	newConfig, err := statusSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("New Config: %v\n\n", newConfig)	

	fmt.Printf(">> The buzzer for the FAIL status is changed to two 500ms beeps. Try to authenticate unregistered credentials for the test.\n")
	cli.PressEnter(">> Press ENTER for the next test.\n")

	return nil
}
