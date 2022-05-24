package access

import (
	"fmt"
	"context"
	"biostar/service/access"
	"google.golang.org/grpc"
)

type AccessSvc struct {
	client access.AccessClient
}

func NewAccessSvc(conn *grpc.ClientConn) *AccessSvc {
	return &AccessSvc{
		client: access.NewAccessClient(conn),
	}
}

func (s *AccessSvc) GetList(deviceID uint32) ([]*access.AccessGroup, error) {
	req := &access.GetListRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the access group list: %v\n", err)

		return nil, err
	}

	return resp.GetGroups(), nil
}

func (s *AccessSvc) Add(deviceID uint32, groups []*access.AccessGroup) error {
	req := &access.AddRequest{
		DeviceID: deviceID,
		Groups: groups,
	}

	_, err := s.client.Add(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add the access groups: %v\n", err)

		return err
	}

	return nil
}

func (s *AccessSvc) Delete(deviceID uint32, groupIDs []uint32) error {
	req := &access.DeleteRequest{
		DeviceID: deviceID,
		GroupIDs: groupIDs,
	}

	_, err := s.client.Delete(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete the access groups: %v\n", err)

		return err
	}

	return nil
}


func (s *AccessSvc) DeleteAll(deviceID uint32) error {
	req := &access.DeleteAllRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAll(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the access groups: %v\n", err)

		return err
	}

	return nil
}


func (s *AccessSvc) GetLevelList(deviceID uint32) ([]*access.AccessLevel, error) {
	req := &access.GetLevelListRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetLevelList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the access level list: %v\n", err)

		return nil, err
	}

	return resp.GetLevels(), nil
}

func (s *AccessSvc) AddLevel(deviceID uint32, levels []*access.AccessLevel) error {
	req := &access.AddLevelRequest{
		DeviceID: deviceID,
		Levels: levels,
	}

	_, err := s.client.AddLevel(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add the access levels: %v\n", err)

		return err
	}

	return nil
}

func (s *AccessSvc) DeleteLevel(deviceID uint32, levelIDs []uint32) error {
	req := &access.DeleteLevelRequest{
		DeviceID: deviceID,
		LevelIDs: levelIDs,
	}

	_, err := s.client.DeleteLevel(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete the access levels: %v\n", err)

		return err
	}

	return nil
}


func (s *AccessSvc) DeleteAllLevel(deviceID uint32) error {
	req := &access.DeleteAllLevelRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAllLevel(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the access levels: %v\n", err)

		return err
	}

	return nil
}


func (s *AccessSvc) GetFloorLevelList(deviceID uint32) ([]*access.FloorLevel, error) {
	req := &access.GetFloorLevelListRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetFloorLevelList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the floor level list: %v\n", err)

		return nil, err
	}

	return resp.GetLevels(), nil
}

func (s *AccessSvc) AddFloorLevel(deviceID uint32, levels []*access.FloorLevel) error {
	req := &access.AddFloorLevelRequest{
		DeviceID: deviceID,
		Levels: levels,
	}

	_, err := s.client.AddFloorLevel(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add the floor levels: %v\n", err)

		return err
	}

	return nil
}

func (s *AccessSvc) DeleteFloorLevel(deviceID uint32, levelIDs []uint32) error {
	req := &access.DeleteFloorLevelRequest{
		DeviceID: deviceID,
		LevelIDs: levelIDs,
	}

	_, err := s.client.DeleteFloorLevel(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete the floor levels: %v\n", err)

		return err
	}

	return nil
}

func (s *AccessSvc) DeleteAllFloorLevel(deviceID uint32) error {
	req := &access.DeleteAllFloorLevelRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAllFloorLevel(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the floor levels: %v\n", err)

		return err
	}

	return nil
}