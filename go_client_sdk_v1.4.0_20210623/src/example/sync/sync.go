package main

import (
	"flag"
	"fmt"
	"os"

	"example/card"
	"example/cli"
	"example/client"
	"example/connect"
	"example/event"
	"example/user"
)

const (
	GATEWAY_CA_FILE = "../../../../cert/gateway/ca.crt"
	GATEWAY_IP = "192.168.0.2"
	GATEWAY_PORT = 4000

	CODE_MAP_FILE = "../event/event_code.json"
)


func main() {
	var configFile = flag.String("f", "config.json", "test config file")	
	
	flag.Parse()

	testConfig := &TestConfig{}
	err := testConfig.Read(*configFile)
	if err != nil {
		os.Exit(1)
	}

	gatewayClient := &client.GatewayClient{}

	err = gatewayClient.Connect(GATEWAY_CA_FILE, GATEWAY_IP, GATEWAY_PORT)
	if err != nil {
		fmt.Printf("Cannot connect the device gateway: %v", err)
		os.Exit(1)
	}

	connectSvc := connect.NewConnectSvc(gatewayClient.GetConn())
	userSvc := user.NewUserSvc(gatewayClient.GetConn())
	cardSvc := card.NewCardSvc(gatewayClient.GetConn())
	eventSvc := event.NewEventSvc(gatewayClient.GetConn())
	eventSvc.InitCodeMap(CODE_MAP_FILE)

	devMgr := NewDeviceMgr(connectSvc, testConfig)	
	eventMgr := NewEventMgr(eventSvc, testConfig)
	userMgr := NewUserMgr(userSvc, cardSvc, devMgr, testConfig)

	fmt.Printf("Trying to connect to the devices...\n")

	err = devMgr.HandleConnection(eventMgr.ConnectCallback)
	if err != nil {
		fmt.Printf("Cannot handle the connections: %v", err)
		os.Exit(1)
	}

	err = devMgr.ConnectToDevices()
	if err != nil {
		fmt.Printf("Cannot connect to the devices: %v", err)
		os.Exit(1)
	}

	err = eventMgr.HandleEvent(len(testConfig.GetAsyncConnectInfo()), userMgr.SyncUser)
	if err != nil {
		fmt.Printf("Cannot handle the events: %v", err)
		os.Exit(1)
	}	

	cli.PressEnter("\n>>> Press ENTER to show the test menu\n\n")

	NewTestMenu(devMgr, eventMgr, userMgr).Show()

	devMgr.DeleteConnection()
	eventMgr.StopHandleEvent()
	gatewayClient.Close()
}