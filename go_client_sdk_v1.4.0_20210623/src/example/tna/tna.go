package tna

import (
	"fmt"
	"context"
	"biostar/service/tna"
	"google.golang.org/grpc"
)

type TNASvc struct {
	client tna.TNAClient
}

func NewTNASvc(conn *grpc.ClientConn) *TNASvc {
	return &TNASvc{
		client: tna.NewTNAClient(conn),
	}
}

func (s *TNASvc) GetConfig(deviceID uint32) (*tna.TNAConfig, error) {
	req := &tna.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the tna config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *TNASvc) SetConfig(deviceID uint32, config *tna.TNAConfig) error {
	req := &tna.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the tna config: %v\n", err)

		return err
	}

	return nil
}


func (s *TNASvc) GetTNALog(deviceID, startEventID, maxNumOfLog uint32) ([]*tna.TNALog, error) {
	req := &tna.GetTNALogRequest{
		DeviceID: deviceID,
		StartEventID: startEventID, 
		MaxNumOfLog: maxNumOfLog,
	}

	resp, err := s.client.GetTNALog(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the tna log events: %v\n", err)

		return nil, err
	}

	return resp.GetTNAEvents(), nil
}