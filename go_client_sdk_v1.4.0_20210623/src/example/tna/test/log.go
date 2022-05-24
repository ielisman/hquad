package main

import (
	"fmt"
	"time"
	"biostar/service/tna"

)

func testLog(deviceID uint32) error {
	events, err := tnaSvc.GetTNALog(deviceID, firstEventID, 0)

	if err != nil {
		return err
	}

	config, err := tnaSvc.GetConfig(deviceID)

	if err != nil {
		return err
	}

	fmt.Printf("\n===== TNA Log Events =====\n\n")	

	for _, eventLog := range events {
		printEvent(eventLog, config)
	}

	return nil
}

func printEvent(eventLog *tna.TNALog, config *tna.TNAConfig) {
	fmt.Printf("%v: Device %v, User %v, %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode), getTNALabel(eventLog.TNAKey, config))
}


func getTNALabel(key tna.Key, config *tna.TNAConfig) string {
	if len(config.Labels) > int(key - 1) {
		return fmt.Sprintf("%v(%v)", config.Labels[key - 1], key)
	} else {
		return fmt.Sprintf("%v", key)
	}
}