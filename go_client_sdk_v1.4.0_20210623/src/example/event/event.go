package event

import (
	"biostar/service/event"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventSvc struct {
	client event.EventClient
}

func NewEventSvc(conn *grpc.ClientConn) *EventSvc {
	return &EventSvc{
		client: event.NewEventClient(conn),
	}
}

func (s *EventSvc) GetClient() event.EventClient {
	return s.client
}

func (s *EventSvc) GetLog(deviceID, startEventID, maxNumOfLog uint32) ([]*event.EventLog, error) {
	req := &event.GetLogRequest{
		DeviceID:     deviceID,
		StartEventID: startEventID,
		MaxNumOfLog:  maxNumOfLog,
	}

	resp, err := s.client.GetLog(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the log events: %v\n", err)

		return nil, err
	}

	return resp.GetEvents(), nil
}

func (s *EventSvc) GetLogWithFilter(deviceID, startEventID, maxNumOfLog uint32, filter *event.EventFilter) ([]*event.EventLog, error) {
	req := &event.GetLogWithFilterRequest{
		DeviceID:     deviceID,
		StartEventID: startEventID,
		MaxNumOfLog:  maxNumOfLog,
		Filters: []*event.EventFilter{
			filter,
		},
	}

	resp, err := s.client.GetLogWithFilter(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the log events with a filter: %v\n", err)

		return nil, err
	}

	return resp.GetEvents(), nil
}

func (s *EventSvc) GetImageLog(deviceID, startEventID, maxNumOfLog uint32) ([]*event.ImageLog, error) {
	req := &event.GetImageLogRequest{
		DeviceID:     deviceID,
		StartEventID: startEventID,
		MaxNumOfLog:  maxNumOfLog,
	}

	resp, err := s.client.GetImageLog(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the images logs: %v\n", err)

		return nil, err
	}

	return resp.GetImageEvents(), nil
}

func (s *EventSvc) EnableMonitoring(deviceID uint32) error {
	enableReq := &event.EnableMonitoringRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.EnableMonitoring(context.Background(), enableReq)

	if err != nil {
		fmt.Printf("Cannot enable log monitoring: %v\n", err)
		return err
	}

	return nil
}

const (
	MONITORING_QUEUE_SIZE = 8
)

var (
	eventStream   event.Event_SubscribeRealtimeLogClient
	eventCallback func(eventLog *event.EventLog)
)

func (s *EventSvc) SetEventCallback(callback func(eventLog *event.EventLog)) {
	eventCallback = callback
}

func (s *EventSvc) StartMonitoring(deviceID uint32) (context.CancelFunc, error) {
	err := s.EnableMonitoring(deviceID)
	if err != nil {
		return nil, err
	}

	subReq := &event.SubscribeRealtimeLogRequest{
		QueueSize: MONITORING_QUEUE_SIZE,
		DeviceIDs: []uint32{deviceID},
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	eventStream, err = s.client.SubscribeRealtimeLog(ctx, subReq)

	if err != nil {
		fmt.Printf("Cannot enable log monitoring: %v\n", err)
		return nil, err
	}

	go func() {
		fmt.Printf("Start receiving real-time events\n")

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

			if eventCallback != nil {
				eventCallback(eventLog)
			} else {
				fmt.Printf("Event: %v\n", eventLog)
			}
		}

	}()

	return cancelFunc, nil
}

func (s *EventSvc) StopMonitoring(deviceID uint32) error {
	disableReq := &event.DisableMonitoringRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DisableMonitoring(context.Background(), disableReq)

	if err != nil {
		fmt.Printf("Cannot disable log monitoring: %v\n", err)
		return err
	}

	fmt.Printf("Stop receiving real-time events\n")

	return nil
}
