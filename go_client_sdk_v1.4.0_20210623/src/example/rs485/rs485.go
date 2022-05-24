package rs485

import (
	"fmt"
	"context"
	"biostar/service/rs485"
	"google.golang.org/grpc"
)

type RS485Svc struct {
	client rs485.RS485Client
}

func NewRS485Svc(conn *grpc.ClientConn) *RS485Svc {
	return &RS485Svc{
		client: rs485.NewRS485Client(conn),
	}
}


func (s *RS485Svc) GetConfig(deviceID uint32) (*rs485.RS485Config, error) {
	req := &rs485.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the rs485 config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *RS485Svc) SetConfig(deviceID uint32, config *rs485.RS485Config) error {
	req := &rs485.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the rs485 config: %v\n", err)

		return err
	}

	return nil
}


func (s *RS485Svc) SearchSlave(deviceID uint32) ([]*rs485.SlaveDeviceInfo, error) {
	req := &rs485.SearchDeviceRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.SearchDevice(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot search RS485 slaves: %v\n", err)

		return nil, err
	}

	return resp.SlaveInfos, nil
}


func (s *RS485Svc) GetSlave(deviceID uint32) ([]*rs485.SlaveDeviceInfo, error) {
	req := &rs485.GetDeviceRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetDevice(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get RS485 slaves: %v\n", err)

		return nil, err
	}

	return resp.SlaveInfos, nil
}


func (s *RS485Svc) SetSlave(deviceID uint32, slaveInfos []*rs485.SlaveDeviceInfo) error {
	req := &rs485.SetDeviceRequest{
		DeviceID: deviceID,
		SlaveInfos: slaveInfos,
	}

	_, err := s.client.SetDevice(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set RS485 slaves: %v\n", err)

		return err
	}

	return nil
}