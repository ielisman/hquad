package main

import (
	"fmt"
	"time"
	"biostar/service/event"
)

const (
	FIRST_DOOR_EVENT = 0x5000 // BS2_EVENT_DOOR_UNLOCKED
	LAST_DOOR_EVENT = 0x5E00 // BS2_EVENT_DOOR_UNLOCK
)

func printDoorEvent(eventLog *event.EventLog) {
	if firstEventID == 0 {
		firstEventID = eventLog.ID
	}

	if eventLog.EventCode >= FIRST_DOOR_EVENT && eventLog.EventCode <= LAST_DOOR_EVENT {
		fmt.Printf("%v: Door %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.EntityID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
	} else {
		fmt.Printf("%v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
	}
}
