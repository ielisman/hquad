package wiegand

import (
	"fmt"
	"context"
	"biostar/service/wiegand"
	"google.golang.org/grpc"
)

type WiegandSvc struct {
	client wiegand.WiegandClient
}

func NewWiegandSvc(conn *grpc.ClientConn) *WiegandSvc {
	return &WiegandSvc{
		client: wiegand.NewWiegandClient(conn),
	}
}


func (s *WiegandSvc) GetConfig(deviceID uint32) (*wiegand.WiegandConfig, error) {
	req := &wiegand.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the wiegand config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *WiegandSvc) SetConfig(deviceID uint32, config *wiegand.WiegandConfig) error {
	req := &wiegand.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the wiegand config: %v\n", err)

		return err
	}

	return nil
}

