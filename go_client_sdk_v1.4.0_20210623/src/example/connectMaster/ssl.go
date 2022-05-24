package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
)

func (s *ConnectMasterSvc) EnableSSL(deviceIDs []uint32) error {
	req := &connectMaster.EnableSSLMultiRequest{
		DeviceIDs:      deviceIDs,
	}

	resp, err := s.client.EnableSSLMulti(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot enable SSL: %v %v", err, resp)
		return err
	}

	return nil
}


func (s *ConnectMasterSvc) DisableSSL(deviceIDs []uint32) error {
	req := &connectMaster.DisableSSLMultiRequest{
		DeviceIDs:      deviceIDs,
	}

	resp, err := s.client.DisableSSLMulti(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot disable SSL: %v %v", err, resp)
		return err
	}

	return nil
}

