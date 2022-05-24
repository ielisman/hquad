package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"biostar/service/connect"
)

type DeviceInfo struct {
	DeviceID uint32 `json:"device_id"`
	IPAddr string `json:"ip_addr"`
	Port int32 `json:"port"`
	UseSSL bool `json:"use_ssl"`

	LastEventID uint32 `json:"last_event_id"`
}

type ConfigData struct {
	EnrollDevice DeviceInfo `json:"enroll_device"`
	Devices []DeviceInfo `json:"devices"`
}


type TestConfig struct {
	configData ConfigData

	configFile string
}

func (tc *TestConfig) GetJson() ([]byte, error) {
	return json.MarshalIndent(&tc.configData, "", "  ")
}

func (tc *TestConfig) Read(configFile string) error {
	tc.configFile = configFile

	configData, err := ioutil.ReadFile(configFile)

	if err != nil {
		fmt.Printf("Cannot read the test config file %v: %v\n", configFile, err)
		return err
	}

	err = json.Unmarshal(configData, &tc.configData)

	if err != nil {
		fmt.Printf("Cannot unmarshal the test config file %v: %v", configFile, err)
		return err
	}
	
	return nil
}


func (tc *TestConfig) Write() error {
	configData, err := tc.GetJson()

	if err != nil {
		fmt.Printf("Cannot marshall the test config file %v: %v", tc.configFile, err)
		return err
	}

	err = ioutil.WriteFile(tc.configFile, configData, 0644)

	if err != nil {
		fmt.Printf("Cannot write the test config file %v: %v\n", tc.configFile, err)
		return err
	}
	
	return nil
}


func (tc *TestConfig) GetAsyncConnectInfo() []*connect.AsyncConnectInfo {
	connectInfos := []*connect.AsyncConnectInfo{}

	enrollDeviceInfo := &connect.AsyncConnectInfo{
		DeviceID: tc.configData.EnrollDevice.DeviceID,
		IPAddr: tc.configData.EnrollDevice.IPAddr,
		Port: tc.configData.EnrollDevice.Port,
		UseSSL: tc.configData.EnrollDevice.UseSSL,
	}

	connectInfos = append(connectInfos, enrollDeviceInfo)

	for _, dev := range tc.configData.Devices {
		devInfo := &connect.AsyncConnectInfo{
			DeviceID: dev.DeviceID,
			IPAddr: dev.IPAddr,
			Port: dev.Port,
			UseSSL: dev.UseSSL,
		}
	
		connectInfos = append(connectInfos, devInfo)
	}

	return connectInfos
}


func (tc *TestConfig) GetTargetDeviceIDs(connectedIDs []uint32) []uint32 {
	devIDs := []uint32{}

	for _, connID := range connectedIDs {
		for _, dev := range tc.configData.Devices {
			if connID == dev.DeviceID {
				devIDs = append(devIDs, dev.DeviceID)
				break
			}
		}
	}

	return devIDs
}

func (tc *TestConfig) GetEnrollDeviceID() uint32 {
	return tc.configData.EnrollDevice.DeviceID
}

func (tc *TestConfig) GetDeviceInfo(devID uint32) *DeviceInfo {
	if devID == tc.GetEnrollDeviceID() {
		return &tc.configData.EnrollDevice
	}

	for _, dev := range tc.configData.Devices {
		if devID == dev.DeviceID {
			return &dev
		}
	}

	return nil
}

func (tc *TestConfig) UpdateLastEventID(devID, lastEventID uint32) {
	updated := false
	
	if devID == tc.GetEnrollDeviceID() {
		tc.configData.EnrollDevice.LastEventID = lastEventID
		updated = true
	} else {
		for i := 0; i < len(tc.configData.Devices); i++ {
			if devID == tc.configData.Devices[i].DeviceID {
				tc.configData.Devices[i].LastEventID = lastEventID
				updated = true
				break
			}
		}
	}

	if updated {
		tc.Write()
	}
}
