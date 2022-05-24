package door

import (
	"fmt"
	"context"
	"biostar/service/door"
	"google.golang.org/grpc"
)

type DoorSvc struct {
	client door.DoorClient
}

func NewDoorSvc(conn *grpc.ClientConn) *DoorSvc {
	return &DoorSvc{
		client: door.NewDoorClient(conn),
	}
}

func (s *DoorSvc) GetList(deviceID uint32) ([]*door.DoorInfo, error) {
	req := &door.GetListRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the door list: %v\n", err)

		return nil, err
	}

	return resp.GetDoors(), nil
}


func (s *DoorSvc) GetStatus(deviceID uint32) ([]*door.Status, error) {
	req := &door.GetStatusRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetStatus(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the door status: %v\n", err)

		return nil, err
	}

	return resp.GetStatus(), nil
}

func (s *DoorSvc) Add(deviceID uint32, doors []*door.DoorInfo) error {
	req := &door.AddRequest{
		DeviceID: deviceID,
		Doors: doors,
	}

	_, err := s.client.Add(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add the doors: %v\n", err)

		return err
	}

	return nil
}

func (s *DoorSvc) Delete(deviceID uint32, doorIDs []uint32) error {
	req := &door.DeleteRequest{
		DeviceID: deviceID,
		DoorIDs: doorIDs,
	}

	_, err := s.client.Delete(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete the doors: %v\n", err)

		return err
	}

	return nil
}


func (s *DoorSvc) DeleteAll(deviceID uint32) error {
	req := &door.DeleteAllRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAll(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the doors: %v\n", err)

		return err
	}

	return nil
}


func (s *DoorSvc) Lock(deviceID uint32, doorIDs []uint32, doorFlag door.DoorFlag) error {
	req := &door.LockRequest{
		DeviceID: deviceID,
		DoorIDs: doorIDs,
		DoorFlag: uint32(doorFlag),
	}

	_, err := s.client.Lock(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot lock the doors: %v\n", err)

		return err
	}

	return nil
}


func (s *DoorSvc) Unlock(deviceID uint32, doorIDs []uint32, doorFlag door.DoorFlag) error {
	req := &door.UnlockRequest{
		DeviceID: deviceID,
		DoorIDs: doorIDs,
		DoorFlag: uint32(doorFlag),
	}

	_, err := s.client.Unlock(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot unlock the doors: %v\n", err)

		return err
	}

	return nil
}


func (s *DoorSvc) Release(deviceID uint32, doorIDs []uint32, doorFlag door.DoorFlag) error {
	req := &door.ReleaseRequest{
		DeviceID: deviceID,
		DoorIDs: doorIDs,
		DoorFlag: uint32(doorFlag),
	}

	_, err := s.client.Release(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot release the doors: %v\n", err)

		return err
	}

	return nil
}


func (s *DoorSvc) SetAlarm(deviceID uint32, doorIDs []uint32, alarmFlag door.AlarmFlag) error {
	req := &door.SetAlarmRequest{
		DeviceID: deviceID,
		DoorIDs: doorIDs,
		AlarmFlag: uint32(alarmFlag),
	}

	_, err := s.client.SetAlarm(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set alarm the doors: %v\n", err)

		return err
	}

	return nil
}

