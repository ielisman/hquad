package cli

import (
	"example/cli"
	"example/connectMaster"
)

var (
	mainMenuItems []*cli.MenuItem
	connectMasterSvc *connectMaster.ConnectMasterSvc
	
	exitCh chan interface{}
	gatewayID string
)

func InitMainMenu(svc *connectMaster.ConnectMasterSvc, targetGatewayID string) {
	connectMasterSvc = svc
	gatewayID = targetGatewayID

	mainMenuItems = []*cli.MenuItem{
		&cli.MenuItem{
			"1", "Search devices", searchDevice, false,
		},
		&cli.MenuItem{
			"2", "Connect to a device synchronously", connectSync, false,
		},
		&cli.MenuItem{
			"3", "Manage asynchronous connections", showAsyncMenu, false,
		},
		&cli.MenuItem{
			"4", "Accept devices", showAcceptMenu, false,
		},
		&cli.MenuItem{
			"5", "Device menu", showDeviceMenu, false,
		},
		&cli.MenuItem{
			"q", "Quit", nil, true,
		},
	}
}

func ShowMainMenu(done chan interface{}) {
	cli.ShowMenu("Main Menu", mainMenuItems)
	done <- 0
}



