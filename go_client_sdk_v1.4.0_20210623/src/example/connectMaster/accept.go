package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
	"biostar/service/connect"
)

func (s *ConnectMasterSvc) GetPendingList(gatewayID string) ([]*connect.PendingDeviceInfo, error) {
	req := &connectMaster.GetPendingListRequest{
		GatewayID: gatewayID,
	}
	
	resp, err := s.client.GetPendingList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the pending device list: %v", err)
		return nil, err
	}

	return resp.GetDeviceInfos(), nil
}


func (s *ConnectMasterSvc) GetAcceptFilter(gatewayID string) (*connect.AcceptFilter, error) {
	req := &connectMaster.GetAcceptFilterRequest{
		GatewayID: gatewayID,
	}

	resp, err := s.client.GetAcceptFilter(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get accept filter: %v\n", err)
		return nil, err
	}

	return resp.GetFilter(), nil
}


func (s *ConnectMasterSvc) SetAcceptFilter(gatewayID string, filter *connect.AcceptFilter) error {
	req := &connectMaster.SetAcceptFilterRequest{
		GatewayID: gatewayID,
		Filter: filter,
	}

	_, err := s.client.SetAcceptFilter(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set accept filter: %v\n", err)
		return err
	}

	return nil
}


func (s *ConnectMasterSvc) AddDeviceToAcceptFilter(gatewayID string, deviceIDs []uint32) error {
	getReq := &connectMaster.GetAcceptFilterRequest{
		GatewayID: gatewayID,
	}

	getResp, err := s.client.GetAcceptFilter(context.Background(), getReq)

	if err != nil {
		fmt.Printf("Cannot get accept filter: %v\n", err)
		return err
	}

	filter := getResp.GetFilter()
	filter.AllowAll = false

	for _, deviceID := range deviceIDs {
		exist := false;

		for i := 0; i < len(filter.DeviceIDs); i++ {
			if filter.DeviceIDs[i] == deviceID {
				exist = true;
				break
			}
		}

		if !exist {
			filter.DeviceIDs = append(filter.DeviceIDs, deviceID)
		}
	}

	setReq := &connectMaster.SetAcceptFilterRequest{
		GatewayID: gatewayID,
		Filter: filter,
	}

	_, err = s.client.SetAcceptFilter(context.Background(), setReq)

	if err != nil {
		fmt.Printf("Cannot set accept filter: %v\n", err)
		return err
	}

	return nil
}


func (s *ConnectMasterSvc) DeleteDeviceFromAcceptFilter(deviceIDs []uint32) error {
	getReq := &connectMaster.GetAcceptFilterRequest{}

	getResp, err := s.client.GetAcceptFilter(context.Background(), getReq)

	if err != nil {
		fmt.Printf("Cannot get accept filter: %v\n", err)
		return err
	}

	filter := getResp.GetFilter()
	filter.AllowAll = false

	for _, deviceID := range deviceIDs {
		for i := 0; i < len(filter.DeviceIDs); i++ {
			if filter.DeviceIDs[i] == deviceID {
				filter.DeviceIDs = append(filter.DeviceIDs[:i], filter.DeviceIDs[i+1:]...)
				break
			}
		}
	}

	setReq := &connectMaster.SetAcceptFilterRequest{
		Filter: filter,
	}

	_, err = s.client.SetAcceptFilter(context.Background(), setReq)

	if err != nil {
		fmt.Printf("Cannot set accept filter: %v", err)
		return err
	}

	return nil
}
