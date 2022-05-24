package main

import (
	"fmt"
	"context"
	"google.golang.org/grpc/status"	
	"google.golang.org/grpc/codes"	

	connectEx "example/connect"
	"biostar/service/connect"
)

type DeviceMgr struct {
	connectSvc *connectEx.ConnectSvc
	testConfig *TestConfig

	connectedIDs []uint32
	cancelFunc context.CancelFunc
}

func NewDeviceMgr(connectSvc *connectEx.ConnectSvc, testConfig *TestConfig) *DeviceMgr {
	return &DeviceMgr {
		connectSvc: connectSvc,
		testConfig: testConfig,

		connectedIDs: []uint32{},
	}
}

func (m *DeviceMgr) ConnectToDevices() error {
	connInfos := m.testConfig.GetAsyncConnectInfo()

	err := m.connectSvc.AddAsyncConnection(connInfos)

	if err != nil {
		return err
	}

	return nil
}


func (m *DeviceMgr) GetConnectedDevices(refreshList bool) ([]uint32, error) {
	if !refreshList {
		deviceIDs := make([]uint32, len(m.connectedIDs))
		copy(deviceIDs, m.connectedIDs)

		return deviceIDs, nil
	}

	devInfos, err := m.connectSvc.GetDeviceList()

	if err != nil {
		return nil, err
	}

	deviceIDs := []uint32{}

	for _, dev := range devInfos {
		if dev.Status == connect.Status_TCP_CONNECTED || dev.Status == connect.Status_TLS_CONNECTED {
			deviceIDs = append(deviceIDs, dev.DeviceID)
		}
	}

	m.connectedIDs = deviceIDs
	
	return m.connectedIDs, nil
}

func (m *DeviceMgr) HandleConnection(callback func(devID uint32) error) error {
	statusStream, cancelFunc, err := m.connectSvc.Subscribe()

	if err != nil {
		return err
	}

	m.cancelFunc = cancelFunc

	go func() {
		for {
			devStatus, err := statusStream.Recv()

			if err != nil {
				status, ok := status.FromError(err)
				if ok && status.Code() == codes.Canceled {
					fmt.Printf("Subscription is cancelled\n")
				} else {
					fmt.Printf("Cannot receive device status: %v\n", err)
				}

				return
			}

			if devStatus.Status == connect.Status_TCP_CONNECTED || devStatus.Status == connect.Status_TLS_CONNECTED {
				m.connectedIDs = append(m.connectedIDs, devStatus.DeviceID)

				if callback != nil {
					callback(devStatus.DeviceID)
				}
			} else if devStatus.Status == connect.Status_DISCONNECTED {
				for i := 0; i < len(m.connectedIDs); i++ {
					if m.connectedIDs[i] == devStatus.DeviceID {
						m.connectedIDs = append(m.connectedIDs[:i], m.connectedIDs[i+1:]...)
						break
					}
				}
			}

			if devStatus.Status != connect.Status_TCP_NOT_ALLOWED && devStatus.Status != connect.Status_TLS_NOT_ALLOWED {
				fmt.Printf("\n[STATUS] Device status change: %v\n", *devStatus)
			}
		}
	} ()

	return nil
}


func (m *DeviceMgr) DeleteConnection() error {
	if len(m.connectedIDs) > 0 {
		return m.connectSvc.DeleteAsyncConnection(m.connectedIDs)
	}

	if m.cancelFunc != nil {
		m.cancelFunc()
	}

	return nil
}