package main

import (
	"fmt"
	"example/cli"
	"biostar/service/user"
	"biostar/service/finger"
)

const (
	QUALITY_THRESHOLD uint32 = 50
	NUM_OF_TEMPLATE = 2
)

func testFinger(deviceID uint32, userID string) error {
	fmt.Printf("\n===== Finger Test =====\n\n")

	fingerData := &finger.FingerData{
		Templates: [][]byte{},
	}

	for i := 0; i < NUM_OF_TEMPLATE; i++ {
		if i == 0 {
			fmt.Printf(">> Scan a finger on the device...\n")		
		} else {
			fmt.Printf(">> Scan the same finger on the device...\n")		
		}

		templateData, _, err := fingerSvc.Scan(deviceID, finger.TemplateFormat_TEMPLATE_FORMAT_SUPREMA, QUALITY_THRESHOLD)

		if err != nil {
			return err
		}

		fingerData.Templates = append(fingerData.Templates, templateData)
	}

	userFinger := &user.UserFinger{
		UserID: userID,
		Fingers: []*finger.FingerData{
			fingerData,
		},
	}

	err := userSvc.SetFinger(deviceID, []*user.UserFinger{ userFinger })

	if err != nil {
		return err
	}

	cli.PressEnter(">> Try to authenticate the enrolled finger. And press ENTER to end the test.\n")

	return nil
}
