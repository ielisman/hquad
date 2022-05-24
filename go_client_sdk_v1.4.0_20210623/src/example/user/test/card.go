package main

import (
	"fmt"
	"example/cli"
	"biostar/service/user"
	"biostar/service/card"
)

func testCard(deviceID uint32, userID string) error {
	fmt.Printf("\n===== Card Test =====\n\n")

	fmt.Printf(">> Place a unregistered card on the device...\n")
	
	cardData, err := cardSvc.Scan(deviceID)

	if err != nil {
		return err
	}

	if cardData.CSNCardData == nil {
		fmt.Printf("!! The card is a smart card. For this test, you have to use a CSN card. Skip the card test.\n")
		return nil
	}

	userCard := &user.UserCard{
		UserID: userID,
		Cards: []*card.CSNCardData{
			cardData.CSNCardData,
		},
	}

	err = userSvc.SetCard(deviceID, []*user.UserCard{ userCard })

	if err != nil {
		return err
	}

	cli.PressEnter(">> Try to authenticate the enrolled card. And, press ENTER to end the test.\n")	

	return nil
}
