package main

import (
	"fmt"
	"strings"
	"example/cli"
)

type TestMenu struct {
	menuItems []*cli.MenuItem
	devMgr *DeviceMgr
	eventMgr *EventMgr
	userMgr *UserMgr
}

func NewTestMenu(devMgr *DeviceMgr, eventMgr *EventMgr, userMgr *UserMgr) *TestMenu {
	return &TestMenu{
		devMgr: devMgr,
		eventMgr: eventMgr,
		userMgr: userMgr,
	}
}

func (m *TestMenu) Show() {
	m.menuItems = []*cli.MenuItem{
		&cli.MenuItem{
			"1", "Show test devices", m.showDevice, false,
		},
		&cli.MenuItem{
			"2", "Show new events", m.showEvent, false,
		},
		&cli.MenuItem{
			"3", "Show new users", m.showUser, false,
		},
		&cli.MenuItem{
			"4", "Enroll a user", m.enrollUser, false,
		},
		&cli.MenuItem{
			"5", "Delete a user", m.deleteUser, false,
		},
		&cli.MenuItem{
			"q", "Quit", nil, true,
		},
	}

	cli.ShowMenu("Test Menu", m.menuItems)
}

func (m *TestMenu) showDevice() error {
	configData, err := m.devMgr.testConfig.GetJson()
	if err != nil {
		return err
	}

	fmt.Printf("***** Test Configuration:\n%v\n\n", string(configData))

	devIDs, err := m.devMgr.GetConnectedDevices(true)
	if err != nil {
		return nil
	}

	fmt.Printf("***** Connected Devices: %v\n", devIDs)

	return nil
}


const (
	MAX_NEW_LOG = 16
)


func (m *TestMenu) showEvent() error {
	devIDs, _ := m.devMgr.GetConnectedDevices(false)

	for _, devID := range devIDs {
		devInfo := m.devMgr.testConfig.GetDeviceInfo(devID)

		if devInfo == nil {
			fmt.Printf("Device %v is not in the configuration file\n", devID)
			continue
		}

		fmt.Printf("Read new event logs from device %v...\n", devID)

		eventLogs, err := m.eventMgr.ReadNewLog(devInfo, MAX_NUM_OF_LOG)
		if err != nil {
			return nil
		}

		fmt.Printf("Read %v event logs\n", len(eventLogs))

		numOfLog := len(eventLogs)
		if numOfLog > MAX_NEW_LOG {
			numOfLog = MAX_NEW_LOG
		}

		if numOfLog > 0 {
			fmt.Printf("Show the last %v events...\n", numOfLog)
			for i := 0; i < numOfLog; i++ {
				m.eventMgr.PrintEvent((eventLogs[numOfLog - i - 1]))
			}
		}
	}

	return nil
}

func (m *TestMenu) showUser() error {
	devIDs, _ := m.devMgr.GetConnectedDevices(false)

	for _, devID := range devIDs {
		fmt.Printf("Read new users from device %v...\n", devID)

		userInfos, err := m.userMgr.GetNewUser(devID)
		if err != nil {
			return nil
		}

		if len(userInfos) > 0	{	
			fmt.Printf("New users: %v\n", userInfos)
		}
	}

	return nil
}


func (m *TestMenu) enrollUser() error {
	return m.userMgr.EnrollUser(m.getUserID())
}

func (m *TestMenu) deleteUser() error {
	return m.userMgr.DeleteUser(m.getUserID())
}

const (
	DEFAULT_USER_ID = "1234"
)

func (m TestMenu) getUserID() string {
	var userIDStr string
	
	inputs := []*cli.UserInput{
		&cli.UserInput{
			"Enter the user ID", DEFAULT_USER_ID, &userIDStr,
		},
	}

	err := cli.GetUserInput(inputs)
	if err != nil {
		return DEFAULT_USER_ID
	}		

	return strings.TrimSpace(userIDStr)
}
