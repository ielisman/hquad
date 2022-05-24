package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
	"biostar/service/connect"
)

func (s *ConnectMasterSvc) AddAsyncConnection(gatewayID string, connInfos []*connect.AsyncConnectInfo) error {
	req := &connectMaster.AddAsyncConnectionRequest{
		GatewayID: gatewayID,
		ConnectInfos: connInfos,
	}

	_, err := s.client.AddAsyncConnection(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set async connectMasterlion list: %v\n", err)
		return err
	}

	return nil
}


func (s *ConnectMasterSvc) DeleteAsyncConnection(gatewayID string, deviceIDs []uint32) error {
	req := &connectMaster.DeleteAsyncConnectionRequest{
		GatewayID: gatewayID,
		DeviceIDs: deviceIDs,
	}

	_, err := s.client.DeleteAsyncConnection(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete async connectMasterion: %v\n", err)
		return err
	}

	return nil
}
