package main

import(
	"biostar/service/action"
	"biostar/service/device"
)

const (
)

func makeEventTrigger(deviceID uint32, eventCode uint32) *action.Trigger {
	return &action.Trigger{
		DeviceID: deviceID,
		Type: action.TriggerType_TRIGGER_EVENT,
		Event: &action.EventTrigger{
			EventCode: eventCode,
		},
	}
}

func makeInputTrigger(deviceID, port uint32, switchType device.SwitchType, durationMS uint32, scheduleID uint32) *action.Trigger {
	return &action.Trigger{
		DeviceID: deviceID,
		Type: action.TriggerType_TRIGGER_INPUT,
		Input: &action.InputTrigger{
			Port: port,
			SwitchType: switchType,
			Duration: durationMS,
			ScheduleID: scheduleID,
		},
	}
}

func makeScheduleTrigger(deviceID uint32, triggerType action.ScheduleTriggerType, scheduleID uint32) *action.Trigger {
	return &action.Trigger {
		DeviceID: deviceID,
		Type: action.TriggerType_TRIGGER_SCHEDULE,
		Schedule: &action.ScheduleTrigger{
			Type: triggerType,
			ScheduleID: scheduleID,	
		},
	}
}

func makeRelayAction(deviceID uint32, relayIndex uint32, signal *action.Signal) *action.Action{
	return &action.Action{
		DeviceID: deviceID,
		Type: action.ActionType_ACTION_RELAY,
		Relay: &action.RelayAction{
			RelayIndex: relayIndex,
			Signal: signal,
		},
	}
}

func makeLEDAction(deviceID uint32, signals []*action.LEDSignal) *action.Action {
	return &action.Action{
		DeviceID: deviceID,
		Type: action.ActionType_ACTION_LED,
		LED: &action.LEDAction{
			Signals: signals,
		},
	}
}

func makeAction(deviceID uint32, actionType action.ActionType) *action.Action {
	return &action.Action{
		DeviceID: deviceID,
		Type: actionType,
	}
}
