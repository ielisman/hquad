package main

import (
	"fmt"
	"time"
	"example/cli"
	"biostar/service/event"
)

const (
	MAX_NUM_OF_EVENT = 32
)

func testLog(deviceID uint32) error {
	eventSvc.StartMonitoring(deviceID)
	eventSvc.SetEventCallback(eventCallback)

	fmt.Printf("\n===== Event Test =====\n\n")	

	cli.PressEnter(">> Try to authenticate credentials to check real-time monitoring. And, press ENTER to read the generated event logs.\n")

	if firstEventID == 0 {
		fmt.Printf("\n>> There is no new event. Just read %v event logs from the start.\n", MAX_NUM_OF_EVENT)
	} else {
		fmt.Printf("\n>> Read new events starting from %v\n", firstEventID)		
	}

	events, err := eventSvc.GetLog(deviceID, firstEventID, MAX_NUM_OF_EVENT)

	if err != nil {
		return err
	}

	for _, eventLog := range events {
		printEvent(eventLog)
	}

	if len(events) > 0 && firstEventID != 0 {
		eventFilter := &event.EventFilter{
			EventCode: events[0].EventCode,
		}

		filteredEvents, err := eventSvc.GetLogWithFilter(deviceID, firstEventID, MAX_NUM_OF_EVENT, eventFilter)
		if err != nil {
			return err
		}

		fmt.Printf("\n>> Filter with event code %v\n", eventFilter.EventCode)

		for _, eventLog := range filteredEvents {
			printEvent(eventLog)
		}
	}

	return nil
}

func printEvent(eventLog *event.EventLog) {
	fmt.Printf("%v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
}


func eventCallback(eventLog *event.EventLog) {
	if firstEventID == 0 {
		firstEventID = eventLog.ID
	}

	fmt.Printf("	Realtime Event: %v\n", eventLog)
}