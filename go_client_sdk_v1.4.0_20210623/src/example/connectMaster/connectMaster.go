package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
	"biostar/service/connect"
	"google.golang.org/grpc"
)

type ConnectMasterSvc struct {
	client connectMaster.ConnectMasterClient
}

func NewConnectMasterSvc(conn *grpc.ClientConn) *ConnectMasterSvc {
	return &ConnectMasterSvc{
		client: connectMaster.NewConnectMasterClient(conn),
	}
}

func (s *ConnectMasterSvc) GetDeviceList(gatewayID string) ([]*connect.DeviceInfo, error) {
	req := &connectMaster.GetDeviceListRequest{
		GatewayID: gatewayID,
	}

	resp, err := s.client.GetDeviceList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the device list: %v\n", err)

		return nil, err
	}

	return resp.GetDeviceInfos(), nil
}


func (s *ConnectMasterSvc) SearchDevice(gatewayID string, timeout uint32) ([]*connect.SearchDeviceInfo, error) {
	req := &connectMaster.SearchDeviceRequest{
		GatewayID: gatewayID, 
		Timeout: timeout,
	}

	resp, err := s.client.SearchDevice(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot search the devices: %v\n", err)
		return nil, err
	}

	return resp.GetDeviceInfos(), nil
}