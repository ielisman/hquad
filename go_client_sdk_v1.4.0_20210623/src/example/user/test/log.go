package main

import (
	"fmt"
	"time"
	"biostar/service/event"
)

func printUserLog(deviceID uint32, userID string) error {
	userFilter := &event.EventFilter{
		UserID: userID,
	}

	events, err := eventSvc.GetLogWithFilter(deviceID, 0, 0, userFilter)

	if err != nil {
		return err
	}

	fmt.Printf("\n===== Log Events of User %v =====\n\n", userID)	

	for _, eventLog := range events {
		printEvent(eventLog)
	}

	eventFilter := &event.EventFilter{
		UserID: userID,
		EventCode: 0x1000, // BS2_EVENT_VERIFY_SUCCESS
	}

	events, err = eventSvc.GetLogWithFilter(deviceID, 0, 0, eventFilter)

	if err != nil {
		return err
	}

	fmt.Printf("\n===== Verify Success Events of User %v =====\n\n", userID)	

	for _, eventLog := range events {
		printEvent(eventLog)
	}

	return nil
}

func printEvent(eventLog *event.EventLog) {
	fmt.Printf("%v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
}
