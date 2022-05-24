package main

import (
	"fmt"
	"time"
	"biostar/service/schedule"
)

func testSchedule(deviceID uint32) error {
	origSchedules, origHolidayGroups, err := backupSchedule(deviceID)
	if err != nil {
		return err
	}

	testSampleSchedule(deviceID)

	err = restoreSchedule(deviceID, origSchedules, origHolidayGroups)
	if err != nil {
		return err
	}
	
	return nil
}


func testSampleSchedule(deviceID uint32) error {
	fmt.Printf("\n===== Test Sample Schedules =====\n\n")

	err := cleanUpSchedule(deviceID)
	if err != nil {
		return err
	}

	holidaySchedule, err := makeHolidaySchedule(deviceID)
	if err != nil {
		return err
	}

	err = makeWeeklySchedule(deviceID, holidaySchedule)
	if err != nil {
		return err
	}

	err = makeDailySchedule(deviceID)
	if err != nil {
		return err
	}

	// check new schedules
	newSchedules, err := scheduleSvc.GetList(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("\n>>> Sample Schedules\n\n")
	for _, sched := range newSchedules {
		fmt.Printf("%v\n\n", sched)
	}

	newHolidays, err := scheduleSvc.GetHolidayList(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("\n>>> Sample Holiday Group\n\n")
	for _, hg := range newHolidays {
		fmt.Printf("%v\n\n", hg)
	}

	return nil
}

const (
	SAMPLE_HOLIDAY_GROUP_ID = 1
)

func makeHolidaySchedule(deviceID uint32) (*schedule.HolidaySchedule, error) {
	holidays := []*schedule.Holiday{
		&schedule.Holiday{
			Date: 0, // Jan. 1
			Recurrence: schedule.HolidayRecurrence_RECUR_YEARLY,
		},
		&schedule.Holiday{
			Date: 358, // Dec. 25 in non leap year, 359 in leap year
			Recurrence: schedule.HolidayRecurrence_RECUR_YEARLY,
		},
	}

	holidayGroup := &schedule.HolidayGroup{
		ID: SAMPLE_HOLIDAY_GROUP_ID,
		Name: "Sample Holiday Group",
		Holidays: holidays,
	}

	holidayPeriods := []*schedule.TimePeriod {
		&schedule.TimePeriod{
			StartTime: 600, // 10 * 60 min, 10 am
			EndTime: 720, // 12 * 60 min, 12 pm
		},
	}

	holidaySchedule := &schedule.HolidaySchedule{
		GroupID: SAMPLE_HOLIDAY_GROUP_ID,
		DaySchedule: &schedule.DaySchedule{
			Periods: holidayPeriods,
		},
	}	

	return holidaySchedule, scheduleSvc.AddHoliday(deviceID, []*schedule.HolidayGroup{ holidayGroup })
}

const (
	WEEKLY_SCHEDULE_ID = 2 // 0 and 1 are reserved
)

func makeWeeklySchedule(deviceID uint32, holidaySchedule *schedule.HolidaySchedule) error {
	weekdayPeriods := []*schedule.TimePeriod {
		&schedule.TimePeriod{
			StartTime: 540, // 9 * 60 min, 9 am
			EndTime: 720, // 12 * 60 min, 12 pm
		},
		&schedule.TimePeriod{
			StartTime: 780, // 13 * 60 min, 1 pm
			EndTime: 1080, // 18 * 60 min, 6 pm
		},
	}

	weekday := &schedule.DaySchedule{
		Periods: weekdayPeriods,
	}

	weekendPeriods := []*schedule.TimePeriod {
		&schedule.TimePeriod{
			StartTime: 570, // (9 * 60) + 30 min, 9:30 am
			EndTime: 750, // (12 * 60) + 30 min, 12:30 pm
		},
	}

	weekend := &schedule.DaySchedule{
		Periods: weekendPeriods,
	}

	weeklySchedule := &schedule.WeeklySchedule{
		DaySchedules: []*schedule.DaySchedule{
			weekend, // Sunday
			weekday, // Monday
			weekday, // Tuesday
			weekday, // Wednesday
			weekday, // Thursday
			weekday, // Friday
			weekend, // Saturday
		},
	}

	weeklyScheduleInfo := &schedule.ScheduleInfo{
		ID: WEEKLY_SCHEDULE_ID,
		Name: "Sample Weekly Schedule",
		Weekly: weeklySchedule,
		Holidays: []*schedule.HolidaySchedule{ holidaySchedule },
	}

	return scheduleSvc.Add(deviceID, []*schedule.ScheduleInfo{ weeklyScheduleInfo })
}

const (
	DAILY_SCHEDULE_ID = WEEKLY_SCHEDULE_ID + 1
	NUM_OF_DAY = 30
)

func makeDailySchedule(deviceID uint32) error {
	dailySchedule := &schedule.DailySchedule{
		StartDate: uint32(time.Now().YearDay() - 1), // 30 days from today
		DaySchedules: []*schedule.DaySchedule{},
	}

	dayPeriods := []*schedule.TimePeriod {
		&schedule.TimePeriod{
			StartTime: 540, // 9 * 60 min, 9 am
			EndTime: 720, // 12 * 60 min, 12 pm
		},
		&schedule.TimePeriod{
			StartTime: 780, // 13 * 60 min, 1 pm
			EndTime: 1080, // 18 * 60 min, 6 pm
		},
	}	

	for i := 0; i < NUM_OF_DAY; i++ {
		daySchedule := &schedule.DaySchedule{
			Periods: dayPeriods,
		}			

		dailySchedule.DaySchedules = append(dailySchedule.DaySchedules, daySchedule)
	}

	dailyScheduleInfo := &schedule.ScheduleInfo{
		ID: DAILY_SCHEDULE_ID,
		Name: "Sample Daily Schedule",
		Daily: dailySchedule,
	}

	return scheduleSvc.Add(deviceID, []*schedule.ScheduleInfo{ dailyScheduleInfo })
}

func backupSchedule(deviceID uint32) ([]*schedule.ScheduleInfo, []*schedule.HolidayGroup, error) {
	schedules, err := scheduleSvc.GetList(deviceID)

	if err != nil {
		return nil, nil, err
	}

	fmt.Printf("Original Schedules: %v\n", schedules)

	holidayGroups, err := scheduleSvc.GetHolidayList(deviceID)

	if err != nil {
		return nil, nil, err
	}

	fmt.Printf("Original Holidays: %v\n", holidayGroups)

	return schedules, holidayGroups, nil
}


func cleanUpSchedule(deviceID uint32) error {
	err := scheduleSvc.DeleteAll(deviceID)
	if err != nil {
		return err
	}

	err = scheduleSvc.DeleteAllHoliday(deviceID)
	if err != nil {
		return err
	}	

	return nil
}



func restoreSchedule(deviceID uint32, schedules []*schedule.ScheduleInfo, holidayGroups []*schedule.HolidayGroup) error {
	err := cleanUpSchedule(deviceID)
	if err != nil {
		return err
	}

	err = scheduleSvc.AddHoliday(deviceID, holidayGroups)
	if err != nil {
		return err
	}

	err = scheduleSvc.Add(deviceID, schedules)
	if err != nil {
		return err
	}

	return nil
}



