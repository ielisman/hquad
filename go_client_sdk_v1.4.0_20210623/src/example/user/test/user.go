package main

import (
	"fmt"
	"time"
	"biostar/service/user"
	"biostar/service/device"
	"biostar/service/auth"
	"biostar/service/event"
)

const (
)


func enrollUser(deviceID uint32, deviceType device.Type) (string, error) {
	eventSvc.StartMonitoring(deviceID)
	eventSvc.SetEventCallback(eventCallback)

	userHdrs, err := userSvc.GetList(deviceID)

	if err != nil {
		return "", err
	}

	fmt.Printf("\nExisting user list: %v\n\n", userHdrs)

	newUserID := fmt.Sprintf("%v", time.Now().Unix())
	newUser := &user.UserInfo{
		Hdr: &user.UserHdr{
			ID: newUserID,
		},
		Setting: &user.UserSetting{
		},
	}

	// Set authentication modes for test
	if deviceType == device.Type_FACESTATION_F2 || deviceType == device.Type_FACESTATION_F2_FP {
		newUser.Setting.CardAuthExtMode = uint32(auth.AuthMode_AUTH_EXT_MODE_CARD_ONLY)
		newUser.Setting.FingerAuthExtMode = uint32(auth.AuthMode_AUTH_EXT_MODE_FINGERPRINT_ONLY)
		newUser.Setting.FaceAuthExtMode = uint32(auth.AuthMode_AUTH_EXT_MODE_FACE_ONLY)
	} else {
		newUser.Setting.CardAuthMode = uint32(auth.AuthMode_AUTH_MODE_CARD_ONLY)
		newUser.Setting.BiometricAuthMode = uint32(auth.AuthMode_AUTH_MODE_BIOMETRIC_ONLY)
	}

	err = userSvc.Enroll(deviceID, []*user.UserInfo{ newUser })

	if err != nil {
		return "", err
	}

	userInfos, err := userSvc.GetUser(deviceID, []string{ newUserID })
	if err != nil {
		return "", err
	}

	fmt.Printf("\nTest User %v: %+v %+v\n\n", newUserID, *userInfos[0].Hdr, *userInfos[0].Setting)

	return newUserID, nil
}

func deleteUser(deviceID uint32, userID string) error {
	return userSvc.Delete(deviceID, []string{ userID })
}


func eventCallback(eventLog *event.EventLog) {
	fmt.Printf("	Realtime Event: %v\n", eventLog)
}