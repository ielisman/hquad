package main

import (
	connectService "biostar/service/connect"
	"example/client"
	"example/connect"
	"example/connect/test/cli"
	"fmt"
	"os"
)

const (
	CA_FILE     = "../../../../src/hquad/user/data/cert/ca.crt" //"../../../../../cert/gateway/ca.crt"
	SERVER_IP   = "192.168.1.110"
	SERVER_PORT = 4000
)

var (
	grpcClient *client.GatewayClient

	connectSvc *connect.ConnectSvc
)

func main() {
	grpcClient = &client.GatewayClient{}

	err := grpcClient.Connect(CA_FILE, SERVER_IP, SERVER_PORT)
	if err != nil {
		os.Exit(1)
	}

	connectSvc = connect.NewConnectSvc(grpcClient.GetConn())

	statusStream, cancelFunc, err := connectSvc.Subscribe()

	if err != nil {
		grpcClient.Close()
		os.Exit(1)
	}

	defer cancelFunc()

	go receiveDeviceStatus(statusStream)

	done := make(chan interface{})

	cli.InitMainMenu(connectSvc)

	go func() {
		cli.ShowMainMenu(done)
	}()

	<-done

	grpcClient.Close()

	return
}

func receiveDeviceStatus(statusStream connectService.Connect_SubscribeStatusClient) {
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
