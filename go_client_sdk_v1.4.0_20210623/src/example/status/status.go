package status

import (
	"fmt"
	"context"
	"biostar/service/status"
	"google.golang.org/grpc"
)

type StatusSvc struct {
	client status.StatusClient
}

func NewStatusSvc(conn *grpc.ClientConn) *StatusSvc {
	return &StatusSvc{
		client: status.NewStatusClient(conn),
	}
}


func (s *StatusSvc) GetConfig(deviceID uint32) (*status.StatusConfig, error) {
	req := &status.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the status config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *StatusSvc) SetConfig(deviceID uint32, config *status.StatusConfig) error {
	req := &status.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the status config: %v\n", err)

		return err
	}

	return nil
}
