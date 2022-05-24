package event

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

var (
	codeMap *EventCodeMap
)

type EventCodeEntry struct {
	EventCode int `json:"event_code"`
	EventCodeStr string `json:"event_code_str`
	SubCode int `json:"sub_code"`
	SubCodeStr string `json:"sub_code_str`
	Desc string `json:"desc"`
}

type EventCodeMap struct {
	Title string `json:"title"`
	Version string `json:"version"`
	Date string `json:"date"`
	Entries []EventCodeEntry `json:"entries"`
}

func (s *EventSvc) InitCodeMap(filename string) error {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Cannot read the code file: %v\n", err)
		return err
	}

	codeMap = &EventCodeMap{}

	err = json.Unmarshal(buf, codeMap)

	if err != nil {
		fmt.Printf("Invalid code map file: %v\n", err)		
		return err
	}

	return nil
}

func (s *EventSvc) GetEventString(eventCode, subCode uint32) string {
	if codeMap == nil {
		return fmt.Sprintf("No code map(%#X)", eventCode | subCode)	
	}

	for _, entry := range codeMap.Entries {
		if eventCode == uint32(entry.EventCode) && subCode == uint32(entry.SubCode) {
			return entry.Desc
		}
	}

	return fmt.Sprintf("Unknown event(%#X)", eventCode | subCode)
}
