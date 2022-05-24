package main

import (
	"os"
	"fmt"
	
	"example/client"
	"example/connect"
	"example/apb"
	"example/rs485"
	"example/event"
)

const (
	GATEWAY_CA_FILE = "../../../../../cert/gateway/ca.crt"
	GATEWAY_IP = "192.168.0.2"
	GATEWAY_PORT = 4000

	DEV_IP = "192.168.0.120"
	DEV_PORT = 51211
	USE_SSL = false

	CODE_MAP_FILE = "../../event/event_code.json"
)

var (
	connectSvc *connect.ConnectSvc

	apbZoneSvc *apb.APBZoneSvc
	rs485Svc *rs485.RS485Svc
	eventSvc *event.EventSvc
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

	apbZoneSvc = apb.NewAPBZoneSvc(gatewayClient.GetConn())
	rs485Svc = rs485.NewRS485Svc(gatewayClient.GetConn())

	slaves, origSlaves, hasSlave := checkSlaves(deviceID)

	if !hasSlave {
		connectSvc.Disconnect([]uint32{ deviceID })
		gatewayClient.Close()
		return
	}

	eventSvc = event.NewEventSvc(gatewayClient.GetConn())
	eventSvc.InitCodeMap(CODE_MAP_FILE)
	eventSvc.StartMonitoring(deviceID)
	eventSvc.SetEventCallback(printZoneEvent)	

	testAPB(deviceID, slaves)

	rs485Svc.SetSlave(deviceID, origSlaves)

	eventSvc.StopMonitoring(deviceID)

	connectSvc.Disconnect([]uint32{ deviceID })
	gatewayClient.Close()
}