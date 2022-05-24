package main

import (
	"os"
	"fmt"
	
	devSvc "biostar/service/device"

	"example/client"
	"example/connect"
	"example/device"
	"example/status"
)

const (
	GATEWAY_CA_FILE = "../../../../../cert/gateway/ca.crt"
	GATEWAY_IP = "192.168.0.2"
	GATEWAY_PORT = 4000

	DEV_IP = "192.168.0.109"
	DEV_PORT = 51211
	USE_SSL = false
)

var (
	connectSvc *connect.ConnectSvc
	deviceSvc *device.DeviceSvc

	statusSvc *status.StatusSvc
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
		connectSvc.Disconnect([]uint32{ deviceID })
		gatewayClient.Close()
		os.Exit(1)
	}	

	if !isHeadless(capInfo.Type) {
		fmt.Printf("Status configuration is effective only for headless devices: %v\n", capInfo.Type)
		connectSvc.Disconnect([]uint32{ deviceID })
		gatewayClient.Close()
		os.Exit(1)
	}

	statusSvc = status.NewStatusSvc(gatewayClient.GetConn())

	testStatus(deviceID)

	connectSvc.Disconnect([]uint32{ deviceID })
	gatewayClient.Close()
}


func isHeadless(devType devSvc.Type) bool {
	switch devType {
	case devSvc.Type_BIOENTRY_P2,
		devSvc.Type_BIOENTRY_R2,
		devSvc.Type_BIOENTRY_W2,
		devSvc.Type_XPASS2,
		devSvc.Type_XPASS2_KEYPAD,
		devSvc.Type_XPASS_D2,
		devSvc.Type_XPASS_D2_KEYPAD,
		devSvc.Type_XPASS_S2:
		return true
	default:
		return false
	}
}