package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
	"biostar/service/connect"
)


func (s *ConnectMasterSvc) SetConnectionMode(deviceID uint32, mode connect.ConnectionMode) error {
	req := &connectMaster.SetConnectionModeRequest{
		DeviceID:       deviceID,
		ConnectionMode: mode,
	}

	_, err := s.client.SetConnectionMode(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set connectMasterion mode: %v", err)
		return err
	}

	return nil
}

func (s *ConnectMasterSvc) SetConnectionModeMulti(deviceIDs []uint32, mode connect.ConnectionMode) error {
	req := &connectMaster.SetConnectionModeMultiRequest{
		DeviceIDs:      deviceIDs,
		ConnectionMode: mode,
	}

	resp, err := s.client.SetConnectionModeMulti(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set connectMasterion mode multi: %v %v", err, resp)
		return err
	}

	return nil
}

