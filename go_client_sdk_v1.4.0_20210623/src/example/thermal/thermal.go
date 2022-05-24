package thermal

import (
	"fmt"
	"context"
	"biostar/service/thermal"
	"google.golang.org/grpc"
)

type ThermalSvc struct {
	client thermal.ThermalClient
}

func NewThermalSvc(conn *grpc.ClientConn) *ThermalSvc {
	return &ThermalSvc{
		client: thermal.NewThermalClient(conn),
	}
}

func (s *ThermalSvc) GetConfig(deviceID uint32) (*thermal.ThermalConfig, error) {
	req := &thermal.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the thermal config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *ThermalSvc) SetConfig(deviceID uint32, config *thermal.ThermalConfig) error {
	req := &thermal.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the thermal config: %v\n", err)

		return err
	}

	return nil
}


func (s *ThermalSvc) GetTemperatureLog(deviceID, startEventID, maxNumOfLog uint32) ([]*thermal.TemperatureLog, error) {
	req := &thermal.GetTemperatureLogRequest{
		DeviceID: deviceID,
		StartEventID: startEventID, 
		MaxNumOfLog: maxNumOfLog,
	}

	resp, err := s.client.GetTemperatureLog(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the temperature log events: %v\n", err)

		return nil, err
	}

	return resp.GetTemperatureEvents(), nil
}