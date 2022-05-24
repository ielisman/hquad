package main

import (
	"fmt"
	"time"
	"context"
	"google.golang.org/grpc/status"	
	"google.golang.org/grpc/codes"	

	eventEx "example/event"
	"biostar/service/event"
)

type EventMgr struct {
	eventSvc *eventEx.EventSvc
	testConfig *TestConfig

	cancelFunc context.CancelFunc
}

func NewEventMgr(eventSvc *eventEx.EventSvc, testConfig *TestConfig) *EventMgr {
	return &EventMgr {
		eventSvc: eventSvc,
		testConfig: testConfig,
	}
}

func (m *EventMgr) HandleEvent(queueSize int, callback func(eventLog *event.EventLog) error) error {
	subReq := &event.SubscribeRealtimeLogRequest{
		QueueSize: int32(queueSize),
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	eventStream, err := m.eventSvc.GetClient().SubscribeRealtimeLog(ctx, subReq)

	if err != nil {
		fmt.Printf("Cannot subscribe: %v\n", err)
		return err
	}

	m.cancelFunc = cancelFunc

	go func() {
		for {
			eventLog, err := eventStream.Recv()

			if err != nil {
				status, ok := status.FromError(err)
				if ok && status.Code() == codes.Canceled {
					fmt.Printf("Real-time monitoring is cancelled\n")
				} else {
					fmt.Printf("Cannot receive real-time events: %v\n", err)
				}

				return
			}

			fmt.Printf("[EVENT] %v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, m.eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))

			if callback != nil {
				callback(eventLog)
			}
		}		
	} ()

	return nil
}

func (m *EventMgr) StopHandleEvent() {
	if m.cancelFunc != nil {
		m.cancelFunc()
	}
}


func (m *EventMgr) ReadNewLog(devInfo *DeviceInfo, maxNumOfLog uint32) ([]*event.EventLog, error) {
	eventLogs, err := m.eventSvc.GetLog(devInfo.DeviceID, devInfo.LastEventID + 1, maxNumOfLog)

	if err != nil {
		return nil, err
	}

	// update the last event ID
	if len(eventLogs) > 0 {
		m.testConfig.UpdateLastEventID(devInfo.DeviceID, eventLogs[len(eventLogs) - 1].ID)
	}

	return eventLogs, nil
}

const (
	MAX_NUM_OF_LOG = 16384
)

func (m *EventMgr) ConnectCallback(devID uint32) error {
	fmt.Printf("***** Device %v is connected\n", devID)

	dev := m.testConfig.GetDeviceInfo(devID)
	if dev == nil {
		fmt.Printf("!!! Device %v is not in the configuration file\n", devID)
		return nil
	}

	// read new logs
	eventLogs := []*event.EventLog{}

	for {
		fmt.Printf("[%v] Reading log records starting from %v...\n", devID, dev.LastEventID)

		events, err := m.ReadNewLog(dev, MAX_NUM_OF_LOG)

		if err != nil {
			return nil
		}

		fmt.Printf("[%v] Read %v events\n", devID, len(events))

		eventLogs = append(eventLogs, events...)

		if len(events) < MAX_NUM_OF_LOG { // read the last log
			break
		}
	}

	// do something with the event logs
	// ...
	fmt.Printf("[%v] The total number of new events: %v\n", devID, len(eventLogs))

	// enable real-time monitoring
	return m.eventSvc.EnableMonitoring(devID)
}


func (m *EventMgr) PrintEvent(eventLog *event.EventLog) {
	fmt.Printf("%v: Device %v, User %v, %v\n", time.Unix(int64(eventLog.Timestamp), 0), eventLog.DeviceID, eventLog.UserID, m.eventSvc.GetEventString(eventLog.EventCode, eventLog.SubCode))
}




