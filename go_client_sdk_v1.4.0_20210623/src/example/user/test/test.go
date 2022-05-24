package main

import (
	"os"
	"fmt"
	
	"example/client"
	"example/device"
	"example/connect"
	"example/user"
	"example/card"
	"example/auth"
	"example/event"
	"example/finger"
	"example/face"
)

const (
	GATEWAY_CA_FILE = "../../../../../cert/gateway/ca.crt"
	GATEWAY_IP = "192.168.0.2"
	GATEWAY_PORT = 4000

	DEV_IP = "192.168.0.135"
	DEV_PORT = 51211
	USE_SSL = false

	CODE_MAP_FILE = "../../event/event_code.json"
)

var (
	connectSvc *connect.ConnectSvc
	deviceSvc *device.DeviceSvc

	userSvc *user.UserSvc
	cardSvc *card.CardSvc
	authSvc *auth.AuthSvc
	eventSvc *event.EventSvc
	fingerSvc *finger.FingerSvc
	faceSvc *face.FaceSvc
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

	userSvc = user.NewUserSvc(gatewayClient.GetConn())
	deviceSvc = device.NewDeviceSvc(gatewayClient.GetConn())
	authSvc = auth.NewAuthSvc(gatewayClient.GetConn())
	eventSvc = event.NewEventSvc(gatewayClient.GetConn())
	eventSvc.InitCodeMap(CODE_MAP_FILE)

	capInfo, err := deviceSvc.GetCapabilityInfo(deviceID)

	if err != nil {
		fmt.Printf("Cannot get the device info: %v", err)
		gatewayClient.Close()
		os.Exit(1)
	}	

	origAuthConfig, err := prepareAuthConfig(deviceID)

	if err != nil {
		fmt.Printf("Cannot set the auth config: %v", err)
		gatewayClient.Close()
		os.Exit(1)
	}

	testUserID, err := enrollUser(deviceID, capInfo.Type)

	if err != nil {
		fmt.Printf("Cannot enroll a test user: %v", err)
		restoreAuthConfig(deviceID, origAuthConfig)
		gatewayClient.Close()
		os.Exit(1)
	}	

	if capInfo.CardSupported {
		cardSvc = card.NewCardSvc(gatewayClient.GetConn())
		testCard(deviceID, testUserID)
	} else {
		fmt.Printf("!! The device %v does not support cards. Skip the card test.\n", deviceID)
	}

	if capInfo.FingerSupported {
		fingerSvc = finger.NewFingerSvc(gatewayClient.GetConn())
		testFinger(deviceID, testUserID)
	} else {
		fmt.Printf("!! The device %v does not support fingerprints. Skip the fingerprint test.\n", deviceID)
	}		

	if capInfo.FaceSupported {
		faceSvc = face.NewFaceSvc(gatewayClient.GetConn())
		testFace(deviceID, testUserID)
	} else {
		fmt.Printf("!! The device %v does not support faces. Skip the face test.\n", deviceID)
	}	

	testAuthMode(deviceID, capInfo.Type)

	printUserLog(deviceID, testUserID)

	restoreAuthConfig(deviceID, origAuthConfig)
	deleteUser(deviceID, testUserID)

	connectSvc.Disconnect([]uint32{ deviceID })
	gatewayClient.Close()
}