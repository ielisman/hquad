package main

import (
	"fmt"
	"example/cli"
	"biostar/service/auth"
	"biostar/service/device"
	"github.com/golang/protobuf/proto"
)

const (
)

func prepareAuthConfig(deviceID uint32) (*auth.AuthConfig, error) {
	origConfig, err := authSvc.GetConfig(deviceID)

	if err != nil {
		return nil, err
	}

	// Enable private authentication for test
	newConfig := proto.Clone(origConfig).(*auth.AuthConfig)
	newConfig.UsePrivateAuth = true

	err = authSvc.SetConfig(deviceID, newConfig)

	if err != nil {
		return origConfig, err
	}

	return origConfig, nil
}


func restoreAuthConfig(deviceID uint32, config *auth.AuthConfig) error {
	return authSvc.SetConfig(deviceID, config)
}


func testAuthMode(deviceID uint32, deviceType device.Type) error {
	fmt.Printf("\n===== Auth Mode Test =====\n\n")
	
	authConfig := &auth.AuthConfig{
		MatchTimeout: 10,
		AuthTimeout: 15,
	}

	if deviceType == device.Type_FACESTATION_F2 || deviceType == device.Type_FACESTATION_F2_FP {
		authConfig.AuthSchedules = []*auth.AuthSchedule{ 
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_EXT_MODE_FACE_ONLY, ScheduleID: 1 }, // Face Only, Always
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_EXT_MODE_FINGERPRINT_ONLY, ScheduleID: 1 }, // Fingerprint Only, Always
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_EXT_MODE_CARD_ONLY, ScheduleID: 1 }, // Card Only, Always
		}
	} else {
		authConfig.AuthSchedules = []*auth.AuthSchedule{
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_MODE_BIOMETRIC_ONLY, ScheduleID: 1 }, // Biometric Only, Always
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_MODE_CARD_ONLY, ScheduleID: 1 }, // Card Only, Always
		}
	}

	err := authSvc.SetConfig(deviceID, authConfig)
	if err != nil {
		return err
	}

	cli.PressEnter(">> Try to authenticate card or fingerprint or face. And, press ENTER for the next test.\n")

	if deviceType == device.Type_FACESTATION_F2 || deviceType == device.Type_FACESTATION_F2_FP {
		authConfig.AuthSchedules = []*auth.AuthSchedule{ 
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_EXT_MODE_CARD_FACE, ScheduleID: 1 }, // Card + Face, Always
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_EXT_MODE_CARD_FINGERPRINT, ScheduleID: 1 }, // Card + Fingerprint, Always
		}
	} else {
		authConfig.AuthSchedules = []*auth.AuthSchedule{
			&auth.AuthSchedule{ Mode: auth.AuthMode_AUTH_MODE_CARD_BIOMETRIC, ScheduleID: 1 }, // Card + Biometric, Always
		}
	}

	err = authSvc.SetConfig(deviceID, authConfig)
	if err != nil {
		return err
	}	

	cli.PressEnter(">> Try to authenticate (card + fingerprint) or (card + face). And, press ENTER to end the test.\n")

	return nil
}
