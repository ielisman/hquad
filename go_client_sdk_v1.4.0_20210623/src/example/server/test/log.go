package main

import (
	"fmt"
	"time"
	"biostar/service/event"
)

func printEvent(eventLog *event.EventLog) {
	fmt.Printf("%v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
}
