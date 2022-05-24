package main

import (
	"flag"
	"fmt"
	"os"

	"example/card"
	"example/client"
	"example/connect"
	"example/connectMaster"
	"example/device"
	"example/event"
	"example/finger"
	"example/master"
	"example/user"
)

const (
	GATEWAY_CA_FILE = "../../../cert/ca.crt"
	GATEWAY_IP      = "192.168.1.110"
	GATEWAY_PORT    = 4000

	MASTER_CA_FILE   = "../../../../cert/master/ca.crt"
	MASTER_IP        = "192.168.0.2"
	MASTER_PORT      = 4010
	ADMIN_CERT_FILE  = "../../../../cert/master/admin.crt"
	ADMIN_KEY_FILE   = "../../../../cert/master/admin_key.pem"
	TENANT_CERT_FILE = "../../../../cert/master/tenant1.crt"
	TENANT_KEY_FILE  = "../../../../cert/master/tenant1_key.pem"

	TENANT_ID  = "tenant1"
	GATEWAY_ID = "gateway1"

	A2_IP   = "192.168.1.89"
	A2_PORT = 51211
	USE_SSL = false
)

var (
	connectSvc       *connect.ConnectSvc
	connectMasterSvc *connectMaster.ConnectMasterSvc

	deviceSvc *device.DeviceSvc
	userSvc   *user.UserSvc
	fingerSvc *finger.FingerSvc
	cardSvc   *card.CardSvc
	eventSvc  *event.EventSvc
)

func connectToGateway() client.GrpcClient {
	gatewayClient := &client.GatewayClient{}

	err := gatewayClient.Connect(GATEWAY_CA_FILE, GATEWAY_IP, GATEWAY_PORT)
	if err != nil {
		fmt.Printf("Cannot connect the device gateway: %v", err)
		return nil
	}

	return gatewayClient
}

func connectToMaster() client.GrpcClient {
	masterClient := &master.MasterClient{}

	err := masterClient.ConnectTenant(MASTER_CA_FILE, TENANT_CERT_FILE, TENANT_KEY_FILE, MASTER_IP, MASTER_PORT)

	if err != nil {
		fmt.Printf("Cannot connect the master gateway: %v", err)
		return nil
	}

	return masterClient
}

func initMaster() error {
	masterClient := &master.MasterClient{}

	err := masterClient.ConnectAdmin(MASTER_CA_FILE, ADMIN_CERT_FILE, ADMIN_KEY_FILE, MASTER_IP, MASTER_PORT)

	if err != nil {
		fmt.Printf("Cannot connect the master gateway: %v", err)
		return err
	}

	err = masterClient.InitTenant(TENANT_ID, GATEWAY_ID)

	if err != nil {
		fmt.Printf("Cannot initialize the tenant DB: %v", err)
		return err
	}

	return nil
}

// To enable gRPC debugging
// $ export GRPC_GO_LOG_VERBOSITY_LEVEL=99 && GRPC_GO_LOG_SEVERITY_LEVEL=info ./quick

func main() {
	masterMode := flag.Bool("m", false, "Connect to the master gateway")
	masterInitMode := flag.Bool("mi", false, "Connect to the master gateway and initialize the tenant DB. You have to do the initialization only once.")

	flag.Parse()

	if *masterInitMode {
		initMaster()
		os.Exit(0)
	}

	var grpcClient client.GrpcClient

	if *masterMode {
		grpcClient = connectToMaster()
		connectMasterSvc = connectMaster.NewConnectMasterSvc(grpcClient.GetConn())
	} else {
		grpcClient = connectToGateway()
		connectSvc = connect.NewConnectSvc(grpcClient.GetConn())
	}

	if grpcClient == nil {
		os.Exit(1)
	}

	deviceSvc = device.NewDeviceSvc(grpcClient.GetConn())
	userSvc = user.NewUserSvc(grpcClient.GetConn())
	fingerSvc = finger.NewFingerSvc(grpcClient.GetConn())
	cardSvc = card.NewCardSvc(grpcClient.GetConn())
	eventSvc = event.NewEventSvc(grpcClient.GetConn())

	var deviceID uint32
	var err error

	if *masterInitMode || *masterMode {
		deviceID, err = testConnectMaster(GATEWAY_ID)
	} else {
		deviceID, err = testConnect()
	}

	if err != nil {
		fmt.Printf("Cannot connect to the device: %v\n", err)
		os.Exit(1)
	}

	capInfo, _ := testDevice(deviceID)

	if capInfo != nil {
		if capInfo.FingerSupported {
			testFinger(deviceID)
		}

		if capInfo.CardSupported {
			testCard(deviceID)
		}
	}

	testUser(deviceID)
	testEvent(deviceID)

	if *masterMode {
		connectMasterSvc.Disconnect([]uint32{deviceID})
	} else {
		connectSvc.Disconnect([]uint32{deviceID})
	}

	grpcClient.Close()

	return
}
