package main

import (
	"fmt"
	userEx "example/user"
	cardEx "example/card"
	errEx "example/err"

	"biostar/service/event"
	"biostar/service/user"
	"biostar/service/card"
)

const (
	BS2_EVENT_USER_ENROLL_SUCCESS			= 0x2000
	BS2_EVENT_USER_UPDATE_SUCCESS			= 0x2200
	BS2_EVENT_USER_DELETE_SUCCESS			= 0x2400
	BS2_EVENT_USER_DELETE_ALL_SUCCESS	= 0x2600
)

type UserMgr struct {
	userSvc *userEx.UserSvc
	cardSvc *cardEx.CardSvc

	deviceMgr *DeviceMgr
	testConfig *TestConfig

	enrolledIDs []string
}

func NewUserMgr(userSvc *userEx.UserSvc, cardSvc *cardEx.CardSvc, deviceMgr *DeviceMgr, testConfig *TestConfig) *UserMgr {
	return &UserMgr {
		userSvc: userSvc,
		cardSvc: cardSvc,
		deviceMgr: deviceMgr,
		testConfig: testConfig,

		enrolledIDs: []string{},
	}
}


func (m *UserMgr) GetNewUser(deviceID uint32) ([]*user.UserInfo, error) {
	if len(m.enrolledIDs) == 0 {
		fmt.Printf("No new user\n")
		return []*user.UserInfo{}, nil
	}

	return m.userSvc.GetUser(deviceID, m.enrolledIDs)
}


func (m *UserMgr) SyncUser(eventLog *event.EventLog) error {
	// Handle only the events of the enrollment device
	if eventLog.DeviceID != m.testConfig.GetEnrollDeviceID() {
		return nil
	}

	connectedIDs, _ := m.deviceMgr.GetConnectedDevices(false)

	targetDeviceIDs := m.testConfig.GetTargetDeviceIDs(connectedIDs)

	if eventLog.EventCode == BS2_EVENT_USER_ENROLL_SUCCESS || eventLog.EventCode == BS2_EVENT_USER_UPDATE_SUCCESS {
		fmt.Printf("Trying to synchronize the enrolled user %v...\n", eventLog.UserID)

		newUserInfos, err := m.userSvc.GetUser(eventLog.DeviceID, []string{eventLog.UserID})
		if err != nil {
			return err
		}

		err = m.userSvc.EnrollMulti(targetDeviceIDs, newUserInfos)		
		if err != nil {
			deviceErrs := errEx.GetMultiError(err)

			fmt.Printf("Enroll errors: %v\n", deviceErrs)
			return err
		}

		m.updateEnrolledIDs(eventLog.UserID)

		// Generate a MultiErrorResponse 
		// It should fail since the users are duplicated.
/*		err = m.userSvc.EnrollMulti(targetDeviceIDs, newUserInfos)		
		if err != nil {
			deviceErrs := errEx.GetMultiError(err)

			fmt.Printf("Multi errors: %v\n", deviceErrs)
			return err
		} */
	} else if eventLog.EventCode == BS2_EVENT_USER_DELETE_SUCCESS {
		fmt.Printf("Trying to synchronize the deleted user %v...\n", eventLog.UserID)

		err := m.userSvc.DeleteMulti(targetDeviceIDs, []string{eventLog.UserID})		
		if err != nil {
			deviceErrs := errEx.GetMultiError(err)

			fmt.Printf("Delete errors: %v\n", deviceErrs)			
			return err
		}

		m.updateDeletedIDs(eventLog.UserID)
	} else if eventLog.EventCode == BS2_EVENT_USER_DELETE_ALL_SUCCESS {
		err := m.userSvc.DeleteAllMulti(targetDeviceIDs)
		if err != nil {
			deviceErrs := errEx.GetMultiError(err)

			fmt.Printf("Delete errors: %v\n", deviceErrs)			
			return err
		}

		m.enrolledIDs = []string{}
	}

	return nil
}


func (m *UserMgr) updateEnrolledIDs(userID string) {
	for _, enrolledID := range m.enrolledIDs {
		if enrolledID == userID { // already enrolled
			return 
		}
	}

	m.enrolledIDs = append(m.enrolledIDs, userID)
}


func (m *UserMgr) updateDeletedIDs(userID string) {
	for i := 0; i < len(m.enrolledIDs); i++ {
		if m.enrolledIDs[i] == userID { // remove the user ID from the enrolled list
			m.enrolledIDs = append(m.enrolledIDs[:i], m.enrolledIDs[i+1:]...)
			break
		}
	}
}


func (m *UserMgr) EnrollUser(userID string) error {
	fmt.Printf(">>> Place a unregistered CSN card on the device %v...\n", m.testConfig.configData.EnrollDevice.DeviceID)
	cardInfo, err := m.cardSvc.Scan(m.testConfig.GetEnrollDeviceID())
	if err != nil {
		return err
	}

	if cardInfo.CSNCardData == nil {
		err = fmt.Errorf("This test does not support a smart card\n")
		fmt.Println(err)
		return err
	}

	userInfo := &user.UserInfo{
		Hdr: &user.UserHdr{
			ID: userID,
			NumOfCard: 1,
		},
		Cards: []*card.CSNCardData{
			cardInfo.CSNCardData,
		},
	}

	// This example calls Enroll to demonstrate the user synchronization. Without user synchronization, EnrollMulti might be a better option
	return m.userSvc.Enroll(m.testConfig.GetEnrollDeviceID(), []*user.UserInfo{ userInfo })
}


func (m *UserMgr) DeleteUser(userID string) error {
	// This example calls Delete to demonstrate the user synchronization. Without user synchronization, DeleteMulti might be a better option
	return m.userSvc.Delete(m.testConfig.GetEnrollDeviceID(), []string{ userID })
}
