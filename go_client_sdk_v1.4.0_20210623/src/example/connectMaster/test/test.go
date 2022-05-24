package main

import (
	"os"
	"fmt"
	"flag"
	connectMasterService "biostar/service/connectMaster"
	connectService "biostar/service/connect"
	"example/master"
	"example/connectMaster"
	"example/connectMaster/test/cli"
)

const (
	MASTER_CA_FILE = "../../../../../cert/master/ca.crt"
	MASTER_IP = "192.168.0.2"
	MASTER_PORT = 4010
	ADMIN_CERT_FILE = "../../../../../cert/master/admin.crt"
	ADMIN_KEY_FILE = "../../../../../cert/master/admin_key.pem"
	TENANT_CERT_FILE = "../../../../../cert/master/tenant1.crt"
	TENANT_KEY_FILE = "../../../../../cert/master/tenant1_key.pem"	

	TENANT_ID = "tenant1"
	GATEWAY_ID = "gateway1"	
)

func connectToMaster() *master.MasterClient{
	masterClient := &master.MasterClient{}

	err := masterClient.ConnectTenant(MASTER_CA_FILE, TENANT_CERT_FILE, TENANT_KEY_FILE, MASTER_IP, MASTER_PORT)

	if err != nil {
		fmt.Printf("Cannot connect the master gateway: %v", err)
		return nil
	}

	return masterClient
}


func initMaster() error{
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


func main() {
	initTenant := flag.Bool("i", false, "Initialize the tenant DB. You only have to do the initialization only once")
	flag.Parse()

	if *initTenant {
		initMaster()
		os.Exit(0)
	} 

	grpcClient := connectToMaster()

	if grpcClient == nil {
		os.Exit(1)
	}

	connectMasterSvc := connectMaster.NewConnectMasterSvc(grpcClient.GetConn())

	statusStream, cancelFunc, err := connectMasterSvc.Subscribe()

	if err != nil {
		grpcClient.Close()
		os.Exit(1)
	}

	defer cancelFunc()

	go receiveDeviceStatus(statusStream) 

	done := make(chan interface{})

	cli.InitMainMenu(connectMasterSvc, GATEWAY_ID)

	go func() {
		cli.ShowMainMenu(done)
	} ()

	<- done

	grpcClient.Close()

	return
}


func receiveDeviceStatus(statusStream connectMasterService.ConnectMaster_SubscribeStatusClient) {
	for {
		statusChange, err := statusStream.Recv()

		if err != nil {
			fmt.Printf("Cannot get device status: %v", err)
			return
		}

		if statusChange.Status != connectService.Status_TCP_NOT_ALLOWED && statusChange.Status != connectService.Status_TLS_NOT_ALLOWED {
			fmt.Printf("\n[STATUS] Device status change: %v\n", *statusChange)
		}
	}
}