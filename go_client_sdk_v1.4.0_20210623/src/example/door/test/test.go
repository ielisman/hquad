package main

import (
	"os"
	"fmt"

	"example/client"
	"example/connect"
	"example/door"
	"example/event"
	"example/user"
	"example/access"
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

	doorSvc *door.DoorSvc
	eventSvc *event.EventSvc
	userSvc *user.UserSvc
	accessSvc *access.AccessSvc

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

	doorSvc = door.NewDoorSvc(gatewayClient.GetConn())
	userSvc = user.NewUserSvc(gatewayClient.GetConn())
	accessSvc = access.NewAccessSvc(gatewayClient.GetConn())
	eventSvc = event.NewEventSvc(gatewayClient.GetConn())
	eventSvc.InitCodeMap(CODE_MAP_FILE)
	eventSvc.StartMonitoring(deviceID)
	eventSvc.SetEventCallback(printDoorEvent)

	testDoor(deviceID)

	eventSvc.StopMonitoring(deviceID)

	connectSvc.Disconnect([]uint32{ deviceID })
	gatewayClient.Close()
}

