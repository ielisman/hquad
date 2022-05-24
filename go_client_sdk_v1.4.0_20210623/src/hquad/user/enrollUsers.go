package main

import (
	"biostar/service/user"
	"context"
	"fmt"
	"hquad/user/biostar"
	"hquad/user/images"
	"hquad/user/parameters"
)

var param parameters.EnrollParameters = parameters.NewEnrollParameters()

func main() {

	fmt.Println()
	fmt.Printf("%-35s %s:%s\n", "G-SDK IP address and port:", param.GatewayIP, param.GatewayPort)
	fmt.Printf("%-35s %-20s\n", "Image Dir with user images:", param.ImageDir)
	fmt.Printf("%-35s %-20s\n", "CA Certification File location:", param.CertificateFile)

	err := biostar.ConnectToGateway(&param)
	ifErrorExitElseOk(err)

	devices, err := biostar.GetAllDevices()
	ifErrorExitElseOk(err)

	for _, device := range devices {
		fmt.Printf("Device ID=%v Type=%v DHCP=%v Mode=%v\n", device.GetDeviceID(), device.GetType(), device.GetUseDHCP(), device.GetConnectionMode())

		err = biostar.ConnectToDevice(device)
		ifErrorExitElseOk(err)

		capInfo, err := biostar.GetDeviceCapability(device)
		ifErrorExitElseOk(err)
		fmt.Printf("\nCapability Info\n %v\n\n", capInfo)

		err = biostar.EnrollUser(device, "stu-1", "User-stu-1", biostar.DefaultStartTime, biostar.DefaultEndTime)
		ifErrorExitElseOk(err)

		var imgFile string = "data/users/001.jpg"
		faceImage, err := images.GetFaceImage(imgFile)
		biostar.EnrollVisualFaceWithNormalization(device, "stu-1", faceImage)
		ifErrorExitElseOk(err)

		err = biostar.DisconnectFromDevice(device)
		ifErrorExitElseOk(err)
	}

}

func getUserIds(userClient user.UserClient) {
	userIds := []string{"stu-714"}
	respUsers, err := userClient.Get(context.Background(), &user.GetRequest{
		DeviceID: 0,
		UserIDs:  userIds,
	})
	if err != nil {
		fmt.Println("Cant retrieve user ids", userIds)
	} else {
		users := respUsers.GetUsers()
		for _, u := range users {
			fmt.Println("Retrieved", u.GetHdr().GetID())
		}
	}
}
