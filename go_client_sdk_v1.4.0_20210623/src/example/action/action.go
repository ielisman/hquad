package action

import (
	"fmt"
	"context"
	"biostar/service/action"
	"google.golang.org/grpc"
)

type ActionSvc struct {
	client action.TriggerActionClient
}

func NewActionSvc(conn *grpc.ClientConn) *ActionSvc {
	return &ActionSvc{
		client: action.NewTriggerActionClient(conn),
	}
}


func (s *ActionSvc) GetConfig(deviceID uint32) (*action.TriggerActionConfig, error) {
	req := &action.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the action config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *ActionSvc) SetConfig(deviceID uint32, config *action.TriggerActionConfig) error {
	req := &action.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the action config: %v\n", err)

		return err
	}

	return nil
}
