package main

import (
	"os"
	"fmt"
	
	"example/client"
	"example/connect"
	"example/device"
	"example/tna"
	"example/event"
)

const (
	GATEWAY_CA_FILE = "../../../../../cert/gateway/ca.crt"
	GATEWAY_IP = "192.168.0.2"
	GATEWAY_PORT = 4000

	DEV_IP = "192.168.0.110"
	DEV_PORT = 51211
	USE_SSL = false

	CODE_MAP_FILE = "../../event/event_code.json"
)

var (
	connectSvc *connect.ConnectSvc
	deviceSvc *device.DeviceSvc

	tnaSvc *tna.TNASvc
	eventSvc *event.EventSvc

	firstEventID uint32 = 0
)


func main() {
	gatewayClient := &client.GatewayClient{}

	err := gatewayClient.Connect(GATEWAY_CA_FILE, GATEWAY_IP, GATEWAY_PORT)
	if err != nil {
		fmt.Printf("Cannot connect the device gateway: %v", err)
		os.Exit(1)
	}

	connectSvc = connect.NewConnectSvc(gatewayClient.GetConn())

	deviceID, err := connectSvc.Connect(DEV_IP, DEV_PORT, USE_SSL)

	if err != nil {
		fmt.Printf("Cannot connect the device: %v", err)
		gatewayClient.Close()
		os.Exit(1)
	}	

	deviceSvc = device.NewDeviceSvc(gatewayClient.GetConn())
	capInfo, err := deviceSvc.GetCapabilityInfo(deviceID)

	if err != nil {
		fmt.Printf("Cannot get the device info: %v", err)
		gatewayClient.Close()
		os.Exit(1)
	}	

	if !capInfo.TNASupported {
		fmt.Printf("T&A service is not supported by the device %v: %v", deviceID, capInfo)
		connectSvc.Disconnect([]uint32{ deviceID })
		gatewayClient.Close()

		os.Exit(1)		
	}

	tnaSvc = tna.NewTNASvc(gatewayClient.GetConn())
	eventSvc = event.NewEventSvc(gatewayClient.GetConn())
	eventSvc.InitCodeMap(CODE_MAP_FILE)

	origConfig, _ := testConfig(deviceID)

	if err != nil {
		fmt.Printf("Cannot test TNA config: %v", err)

		if origConfig != nil {
			tnaSvc.SetConfig(deviceID, origConfig)
		}
		gatewayClient.Close()
		os.Exit(1)
	}		
	
	testLog(deviceID)

	tnaSvc.SetConfig(deviceID, origConfig)

	connectSvc.Disconnect([]uint32{ deviceID })
	gatewayClient.Close()
}