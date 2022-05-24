package apb

import (
	"fmt"
	"context"
	"biostar/service/apbZone"
	"biostar/service/zone"
	"google.golang.org/grpc"
)

type APBZoneSvc struct {
	client apbZone.APBZoneClient
}

func NewAPBZoneSvc(conn *grpc.ClientConn) *APBZoneSvc {
	return &APBZoneSvc{
		client: apbZone.NewAPBZoneClient(conn),
	}
}


func (s *APBZoneSvc) Get(deviceID uint32) ([]*apbZone.ZoneInfo, error) {
	req := &apbZone.GetRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.Get(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get APB zones: %v\n", err)

		return nil, err
	}

	return resp.Zones, nil
}


func (s *APBZoneSvc) GetStatus(deviceID uint32, zoneIDs []uint32) ([]*zone.ZoneStatus, error) {
	req := &apbZone.GetStatusRequest{
		DeviceID: deviceID,
		ZoneIDs: zoneIDs,
	}

	resp, err := s.client.GetStatus(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the status: %v\n", err)

		return nil, err
	}

	return resp.Status, nil
}


func (s *APBZoneSvc) Add(deviceID uint32, zoneInfos []*apbZone.ZoneInfo) error {
	req := &apbZone.AddRequest{
		DeviceID: deviceID,
		Zones: zoneInfos,
	}

	_, err := s.client.Add(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add APB zones: %v\n", err)

		return err
	}

	return nil	
}


func (s *APBZoneSvc) Delete(deviceID uint32, zoneIDs []uint32) error {
	req := &apbZone.DeleteRequest{
		DeviceID: deviceID,
		ZoneIDs: zoneIDs,
	}

	_, err := s.client.Delete(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete APB zones: %v\n", err)

		return err
	}

	return nil	
}


func (s *APBZoneSvc) DeleteAll(deviceID uint32) error {
	req := &apbZone.DeleteAllRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAll(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the APB zones: %v\n", err)

		return err
	}

	return nil	
}


func (s *APBZoneSvc) Clear(deviceID uint32, zoneID uint32, userIDs []string) error {
	req := &apbZone.ClearRequest{
		DeviceID: deviceID,
		ZoneID: zoneID,
		UserIDs: userIDs,
	}

	_, err := s.client.Clear(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot clear APB records: %v\n", err)

		return err
	}

	return nil	
}


func (s *APBZoneSvc) ClearAll(deviceID uint32, zoneID uint32) error {
	req := &apbZone.ClearAllRequest{
		DeviceID: deviceID,
		ZoneID: zoneID,
	}

	_, err := s.client.ClearAll(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot clear all the APB records: %v\n", err)

		return err
	}

	return nil	
}