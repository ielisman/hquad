package main

import (
	"fmt"
	"time"
	"strconv"
	"biostar/service/thermal"
	"biostar/service/event"
)

func testLog(deviceID uint32) error {
	events, err := thermalSvc.GetTemperatureLog(deviceID, firstEventID, 0)

	if err != nil {
		return err
	}

	fmt.Printf("\n===== Log Events with Temperature =====\n\n")	

	for _, eventLog := range events {
		printEvent(eventLog)
	}

	return nil
}

func printEvent(eventLog *thermal.TemperatureLog) {
	userID, err := strconv.Atoi(eventLog.UserID)

	if err != nil || userID == 0xFFFFFFFF { // no user ID
		fmt.Printf("%v: Device %v, %v, Temperature %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode), eventLog.Temperature)
	} else {
		fmt.Printf("%v: Device %v, User %v, %v, Temperature %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode), eventLog.Temperature)
	}
}

func eventCallback(eventLog *event.EventLog) {
	if firstEventID == 0 {
		firstEventID = eventLog.ID
	}

	fmt.Printf("	Realtime Event: %v\n", eventLog)
}