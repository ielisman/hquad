package main

import (
	"fmt"
	"example/cli"
	"biostar/service/user"
	"biostar/service/face"
)

const (
	ENROLL_THRESHOLD = face.FaceEnrollThreshold_BS2_FACE_ENROLL_THRESHOLD_DEFAULT
)

func testFace(deviceID uint32, userID string) error {
	fmt.Printf("\n===== Face Test =====\n\n")

	fmt.Printf(">> Enroll a face on the device...\n")

	faceData, err := faceSvc.Scan(deviceID, ENROLL_THRESHOLD)

	if err != nil {
		return err
	}

	userFace := &user.UserFace{
		UserID: userID,
		Faces: []*face.FaceData{
			faceData,
		},
	}

	err = userSvc.SetFace(deviceID, []*user.UserFace{ userFace })

	if err != nil {
		return err
	}

	cli.PressEnter(">> Try to authenticate the enrolled face. And, press ENTER to end the test.\n")

	return nil
}
