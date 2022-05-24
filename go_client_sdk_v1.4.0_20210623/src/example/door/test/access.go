package main

import (
	"fmt"
	"biostar/service/access"
	"biostar/service/user"
	"example/cli"
)

const (
	TEST_ACCESS_LEVEL_ID = 1
	TEST_ACCESS_GROUP_ID = 1
	ALWAYS_SCHEDULE_ID = 1
)

func testAccessGroup(deviceID uint32) error {
	events, err := eventSvc.GetLog(deviceID, firstEventID, 16)
	if err != nil {
		return err
	}

	userID := ""

	for _, eventLog := range events {
		if eventLog.EventCode == 0x1900 {// BS2_EVENT_ACCESS_DENIED
			userID = eventLog.UserID
			break
		}
	}

	if userID == "" {
		err = fmt.Errorf("!! Cannot find ACCESS_DENIED event. You should have tried to authenticate a registered credentail for the test.")
		fmt.Println(err.Error())
		return err
	}

	origGroups, origLevels, origFloorLevels, err := backupAccessGroup(deviceID)
	if err != nil {
		return err
	}

	err = cleanUpAccessGroup(deviceID)
	if err != nil {
		return err
	}

	origUserAccessGroups, err := userSvc.GetAccessGroup(deviceID, []string{ userID })
	if err != nil {
		return err
	}

	fmt.Printf("Original User Access Groups: %v\n", origUserAccessGroups)	

	accessLevel := &access.AccessLevel{
		ID: TEST_ACCESS_LEVEL_ID,
		DoorSchedules: []*access.DoorSchedule{
			&access.DoorSchedule{
				DoorID: TEST_DOOR_ID,
				ScheduleID: ALWAYS_SCHEDULE_ID, // always
			},
		},
	}

	err = accessSvc.AddLevel(deviceID, []*access.AccessLevel{ accessLevel })
	if err != nil {
		return err
	}

	accessGroup := &access.AccessGroup{
		ID: TEST_ACCESS_GROUP_ID,
		LevelIDs: []uint32{
			TEST_ACCESS_LEVEL_ID,
		},
	}

	err = accessSvc.Add(deviceID, []*access.AccessGroup{ accessGroup })
	if err != nil {
		return err
	}

	testAccessGroups, err := accessSvc.GetList(deviceID)
	if err != nil {
		return err
	}

	testAccessLevels, err := accessSvc.GetLevelList(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("\nTest Access Groups: %v\n", testAccessGroups)	
	fmt.Printf("Test Access Levels: %v\n", testAccessLevels)	

	userAccessGroups := []*user.UserAccessGroup{
		&user.UserAccessGroup{
			UserID: userID,
			AccessGroupIDs: []uint32{ TEST_ACCESS_GROUP_ID },
		},
	}

	err = userSvc.SetAccessGroup(deviceID, userAccessGroups)
	if err != nil {
		return err
	}	

	fmt.Printf("Test User Access Groups: %v\n", userAccessGroups)	

	fmt.Printf("\n>> Try to authenticate the same registered credential. It should succeed since the access group is added.\n")
	cli.PressEnter(">> Press ENTER for the next test.\n")

	testLock(deviceID)

	userSvc.SetAccessGroup(deviceID, origUserAccessGroups)
	restoreAccessGroup(deviceID, origGroups, origLevels, origFloorLevels)

	return nil
}


func backupAccessGroup(deviceID uint32) ([]*access.AccessGroup, []*access.AccessLevel, []*access.FloorLevel, error) {
	groups, err := accessSvc.GetList(deviceID)

	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Printf("Original Access Groups: %v\n", groups)

	levels, err := accessSvc.GetLevelList(deviceID)

	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Printf("Original Access Levels: %v\n", levels)

	floorLevels, err := accessSvc.GetFloorLevelList(deviceID)

	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Printf("Original Floor Levels: %v\n", floorLevels)

	return groups, levels, floorLevels, nil
}

func cleanUpAccessGroup(deviceID uint32) error {
	err := accessSvc.DeleteAll(deviceID)
	if err != nil {
		return err
	}

	err = accessSvc.DeleteAllLevel(deviceID)
	if err != nil {
		return err
	}

	err = accessSvc.DeleteAllFloorLevel(deviceID)
	if err != nil {
		return err
	}

	return nil
}

func restoreAccessGroup(deviceID uint32, groups []*access.AccessGroup, levels []*access.AccessLevel, floorLevels []*access.FloorLevel) error {
	err := cleanUpAccessGroup(deviceID)
	if err != nil {
		return err
	}

	err = accessSvc.AddLevel(deviceID, levels)
	if err != nil {
		return err
	}

	err = accessSvc.AddFloorLevel(deviceID, floorLevels)
	if err != nil {
		return err
	}

	err = accessSvc.Add(deviceID, groups)
	if err != nil {
		return err
	}

	return nil
}