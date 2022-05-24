package schedule

import (
	"fmt"
	"context"
	"biostar/service/schedule"
	"google.golang.org/grpc"
)

type ScheduleSvc struct {
	client schedule.ScheduleClient
}

func NewScheduleSvc(conn *grpc.ClientConn) *ScheduleSvc {
	return &ScheduleSvc{
		client: schedule.NewScheduleClient(conn),
	}
}

func (s *ScheduleSvc) GetList(deviceID uint32) ([]*schedule.ScheduleInfo, error) {
	req := &schedule.GetListRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the schedule list: %v\n", err)

		return nil, err
	}

	return resp.GetSchedules(), nil
}


func (s *ScheduleSvc) Add(deviceID uint32, schedules []*schedule.ScheduleInfo) error {
	req := &schedule.AddRequest{
		DeviceID: deviceID,
		Schedules: schedules,
	}

	_, err := s.client.Add(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add the schedules: %v\n", err)

		return err
	}

	return nil
}


func (s *ScheduleSvc) Delete(deviceID uint32, scheduleIDs []uint32) error {
	req := &schedule.DeleteRequest{
		DeviceID: deviceID,
		ScheduleIDs: scheduleIDs,
	}

	_, err := s.client.Delete(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete the schedules: %v\n", err)

		return err
	}

	return nil
}


func (s *ScheduleSvc) DeleteAll(deviceID uint32) error {
	req := &schedule.DeleteAllRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAll(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the schedules: %v\n", err)

		return err
	}

	return nil
}


func (s *ScheduleSvc) GetHolidayList(deviceID uint32) ([]*schedule.HolidayGroup, error) {
	req := &schedule.GetHolidayListRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetHolidayList(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the holiday list: %v\n", err)

		return nil, err
	}

	return resp.GetGroups(), nil
}


func (s *ScheduleSvc) AddHoliday(deviceID uint32, groups []*schedule.HolidayGroup) error {
	req := &schedule.AddHolidayRequest{
		DeviceID: deviceID,
		Groups: groups,
	}

	_, err := s.client.AddHoliday(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot add the holiday groups: %v\n", err)

		return err
	}

	return nil
}


func (s *ScheduleSvc) DeleteAllHoliday(deviceID uint32) error {
	req := &schedule.DeleteAllHolidayRequest{
		DeviceID: deviceID,
	}

	_, err := s.client.DeleteAllHoliday(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot delete all the holiday groups: %v\n", err)

		return err
	}

	return nil
}




