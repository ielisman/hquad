package main

import (
	"fmt"
	"biostar/service/door"
	"example/cli"
)

const (
	TEST_DOOR_ID = 1
)

func testDoor(deviceID uint32) error {
	origDoors, err := backupDoor(deviceID)
	if err != nil {
		return err
	}

	testSampleDoor(deviceID)
	testAccessGroup(deviceID)	

	err = restoreDoor(deviceID, origDoors)
	if err != nil {
		return err
	}

	return nil
}

func testSampleDoor(deviceID uint32) error {
	singleDoor := makeSingleDoor(deviceID, TEST_DOOR_ID)
	err := doorSvc.Add(deviceID, []*door.DoorInfo{ singleDoor })
	if err != nil {
		return err
	}

	doors, err := doorSvc.GetList(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("\nTest Doors: %v\n", doors)	

	fmt.Printf("\n===== Door Test =====\n\n")

	fmt.Printf(">> Try to authenticate a registered credential. It should fail since you can access a door only with a proper access group.\n")
	cli.PressEnter(">> Press ENTER for the next test.\n")

	return nil
}


func testLock(deviceID uint32) error {
	cli.PressEnter(">> Press ENTER to unlock the door.\n")	
	err := doorSvc.Unlock(deviceID, []uint32{ TEST_DOOR_ID }, door.DoorFlag_OPERATOR)
	if err != nil {
		return err
	}

	doorStatus, err := doorSvc.GetStatus(deviceID);
	if err != nil {
		return err
	}

	fmt.Printf("\nStatus after unlocking the door: %v\n", doorStatus)	

	cli.PressEnter(">> Press ENTER to lock the door.\n")	
	err = doorSvc.Lock(deviceID, []uint32{ TEST_DOOR_ID }, door.DoorFlag_OPERATOR)
	if err != nil {
		return err
	}

	doorStatus, err = doorSvc.GetStatus(deviceID);
	if err != nil {
		return err
	}

	fmt.Printf("\nStatus after locking the door: %v\n", doorStatus)	

	fmt.Printf(">> Try to authenticate the same registered credential. The relay should not work since the door is locked by the operator with the higher priority.\n")
	cli.PressEnter(">> Press ENTER to release the door flag.\n")

	err = doorSvc.Release(deviceID, []uint32{ TEST_DOOR_ID }, door.DoorFlag_OPERATOR)
	if err != nil {
		return err
	}

	doorStatus, err = doorSvc.GetStatus(deviceID);
	if err != nil {
		return err
	}

	fmt.Printf("\nStatus after releasing the door flag: %v\n", doorStatus)	

	fmt.Printf(">> Try to authenticate the same registered credential. The relay should work since the door flag is cleared.\n")
	cli.PressEnter(">> Press ENTER for the next test.\n")

	return nil
}

