package main

import (
	"fmt"
	"time"
	"biostar/service/event"
)

const (
	FIRST_APB_EVENT = 0x6000 // BS2_EVENT_ZONE_APB_VIOLATION
	LAST_APB_EVENT = 0x6200 // BS2_EVENT_ZONE_APB_ALARM_CLEAR
)


func printZoneEvent(eventLog *event.EventLog) {
	if eventLog.EventCode >= FIRST_APB_EVENT && eventLog.EventCode <= LAST_APB_EVENT {
		fmt.Printf("%v: APB Zone %v, User %v,  %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.EntityID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
	} else {
		fmt.Printf("%v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
	}
}
