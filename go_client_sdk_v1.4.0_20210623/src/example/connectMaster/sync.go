package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
	"biostar/service/connect"
)


func (s *ConnectMasterSvc) Connect(gatewayID string, deviceIP string, devicePort int, useSSL bool) (uint32, error) {
	req := &connectMaster.ConnectRequest{
		GatewayID: gatewayID,
		ConnectInfo: &connect.ConnectInfo{
			IPAddr: deviceIP,
			Port: int32(devicePort),
			UseSSL: useSSL,
		},
	}

	resp, err := s.client.Connect(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot connect to %v:%v: %v\n", deviceIP, devicePort, err)
		return 0, err
	}

	return resp.GetDeviceID(), nil
}


func (s *ConnectMasterSvc) Disconnect(deviceIDs []uint32) error {
	req := &connectMaster.DisconnectRequest{
		DeviceIDs: deviceIDs,
	}

	_, err := s.client.Disconnect(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot disconnect %v: %v\n", deviceIDs, err)
		return err
	}

	return nil
}


func (s *ConnectMasterSvc) DisconnectAll(gatewayID string) error {
	req := &connectMaster.DisconnectAllRequest{
		GatewayID: gatewayID, 
	}

	_, err := s.client.DisconnectAll(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot disconnect all: %v\n", err)
		return err
	}

	return nil
}
