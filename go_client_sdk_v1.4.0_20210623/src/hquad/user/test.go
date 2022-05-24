package main

import (
	"biostar/service/connect"
	"biostar/service/face"
	"biostar/service/user"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"hquad/user/biostar"
	"hquad/user/flow"
	"hquad/user/images"
	"hquad/user/parameters"
	"hquad/user/parser"
	"image/jpeg"
	"os"
	"time"
)

var cert string = "data/cert/ca.crt"
var server string = "192.168.1.110"
var port string = "4000"
var f2 string = "192.168.1.89"
var f2port int32 = 51211
var f2id uint32 = 543664380
var uid string = "Igor2"
var imgFile string = "data/users/001.jpg"

var connectClient connect.ConnectClient
var userClient user.UserClient
var faceClient face.FaceClient

var params parameters.EnrollParameters

func init() {
	params = parameters.NewEnrollParameters()
}

func main() {
	//original()
	//current()
	updateSlideImages()
}

func updateSlideImages() {
	var slideImages [][]byte
	test := 3
	if test == 1 {
		imgWhite, err := images.GetPNGImage("data/slideshow/img1.white.png")
		flow.ErrorFlowStandardExit(err)
		imgBlack, err := images.GetPNGImage("data/slideshow/img2.black.png")
		flow.ErrorFlowStandardExit(err)
		slideImages = [][]byte{
			imgWhite, imgBlack,
			imgWhite, imgBlack,
			imgWhite, imgBlack,
			imgWhite, imgBlack,
			imgWhite, imgBlack,
			imgWhite, imgBlack,
			imgWhite, imgBlack,
		}
	} else if test == 2 {
		ilen := 11
		slideImages = make([][]byte, ilen-1)
		for i := 1; i < ilen; i++ {
			img, err := images.GetPNGImage(fmt.Sprintf("data/slideshow/Screen%d.1024.800.png", i))
			flow.ErrorFlowStandardExit(err)
			slideImages[i-1] = img
		}
	} else if test == 3 {
		imgWhite, err := images.GetPNGImage("data/slideshow/img1.white.png")
		flow.ErrorFlowStandardExit(err)
		imgBlack, err := images.GetPNGImage("data/slideshow/img2.black.png")
		flow.ErrorFlowStandardExit(err)
		ilen := 10
		slideImages = make([][]byte, ilen+4)
		slideImages[0] = imgWhite
		slideImages[1] = imgBlack
		slideImages[2] = imgWhite
		slideImages[3] = imgBlack
		for i := 4; i < ilen+4; i++ {
			img, err := images.GetPNGImage(fmt.Sprintf("data/slideshow/Screen%d.1024.800.png", i-3))
			flow.ErrorFlowStandardExit(err)
			slideImages[i] = img
		}
	}

	err := biostar.ConnectToGateway(&params)
	flow.ErrorFlowStandardMessage(err, fmt.Sprintf("Connected to gateway %v:%v successfully", params.GatewayIP, params.GatewayPort))

	devices, err := biostar.GetAllDevices()
	flow.ErrorFlowStandardMessage(err, fmt.Sprintf("Got list of devices via gateway %v:%v successfully", params.GatewayIP, params.GatewayPort))

	for _, device := range devices {
		err = biostar.ConnectToDevice(device)
		flow.ErrorFlowStandardExit(err)

		err = biostar.SetSlideImages(device, slideImages)
		flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)

		err = biostar.DisconnectFromDevice(device)
		flow.ErrorFlowStandardExit(err)
	}

	err = biostar.DisconnectFromGateway()
	flow.ErrorFlowStandardExit(err)
}

func current() {

	imageUsers, err := parser.ProcessImageDir(params.ImageDir)
	flow.ErrorFlowStandardMessage(err, fmt.Sprintf("Processed image directory %v successfully", params.ImageDir))

	err = biostar.ConnectToGateway(&params)
	flow.ErrorFlowStandardMessage(err, fmt.Sprintf("Connected to gateway %v:%v successfully", params.GatewayIP, params.GatewayPort))

	devices, err := biostar.GetAllDevices()
	flow.ErrorFlowStandardMessage(err, fmt.Sprintf("Got list of devices via gateway %v:%v successfully", params.GatewayIP, params.GatewayPort))

	for _, device := range devices {
		err = biostar.ConnectToDevice(device)
		flow.ErrorFlowStandardMessage(err, fmt.Sprintf("Connected to device %v (%v:%v) successfully", device.GetDeviceID(), device.GetIPAddr(), device.GetPort()))

		err = biostar.EnrollUsersConcurrent(device, imageUsers)
		flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)

		/*		users, err := biostar.GetUsersList(device.GetDeviceID(), int(user.InfoMask_USER_MASK_NAME), "Image")
				flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)

				for _, user := range users {
					fmt.Println("User: ", user.GetName(), user.GetCards())
				}*/

		// search directory for images, rename them all, use channels and go routines to verify images
		// equal to number of threads in F2 // 1.8 dual-core and 1.4 quad-core
		// check image normalization speed sequentially 30 times in a row, after that
		// in parallel: at 2 go routines, at 4, 6, 8, 12, 16, 24, 32, 48, 64, 96, 128

		//err = biostar.EnrollUsersConcurrent(device, imageUsers)
		//flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
		//err = biostar.EnrollUser(device, imageUsers[0].GetUserId(), imageUsers[0].GetUserName(), biostar.DefaultStartTime, biostar.DefaultEndTime)
		//flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)

		//err = biostar.EnrollUsersConcurrent(device, imageUsers)
		//flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)

		/*if a != 3 {
			faceImage, err := images.GetFaceImage(imageUsers[0].GetImageFile())
			flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
			imgNormalized, err := biostar.GetNormalized(device, imageUsers[0], faceImage)
			flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
			err = biostar.EnrollUserWithNormalizedFace(device, imageUsers[0], imgNormalized, biostar.DefaultStartTime, biostar.DefaultEndTime)
			flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
			a = 3
		}*/

		/*		err = biostar.EnrollUser(device, "Igor3", "Igor 3", biostar.DefaultStartTime, biostar.DefaultEndTime)
				flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
				img, err := images.GetFaceImage(imageUsers[0].GetImageFile())
				flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
				biostar.EnrollVisualFaceWithNormalization(device, "Igor3", img)
				flow.IfErrorExitElseOk(err, biostar.DisconnectFromDevice, device)
				imageUsers[0] = nil*/

		err = biostar.DisconnectFromDevice(device)
		flow.ErrorFlowStandardExit(err)
	}

}

func original() {
	connectToGateway()
	connectToF2()
	imgData := getImageFromFile()
	enrollUser()
	enrollVisualFace(imgData)
	disconnectFromF2()
}

func connectToGateway() {
	fmt.Printf("making connection to gateway %v:%v ...", server, port)
	cred, err := credentials.NewClientTLSFromFile(cert, "")
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", server, port), grpc.WithTransportCredentials(cred), grpc.WithBlock(), grpc.WithTimeout(10000*time.Millisecond))
	ifErrorExitElseOk(err)

	connectClient = connect.NewConnectClient(conn)
	userClient = user.NewUserClient(conn)
	faceClient = face.NewFaceClient(conn)
}

func connectToF2() {
	fmt.Printf("making connection to F2 %v:%v ...", f2, f2port)
	_, err := connectClient.Connect(context.Background(), &connect.ConnectRequest{
		ConnectInfo: &connect.ConnectInfo{
			IPAddr: f2,
			Port:   f2port,
			UseSSL: false,
		},
	})
	flow.IfErrorExitElseOkId(err, nil, f2id)
}

func getImageFromFile() []byte {
	fmt.Printf("getting image data from file %v ...", imgFile)
	reader, err := os.Open(imgFile)
	defer reader.Close()
	img, err := jpeg.Decode(reader)
	//buf := new(bytes.Buffer)
	buf := bytes.NewBuffer([]byte{})
	jpeg.Encode(buf, img, nil)
	flow.IfErrorExitElseOkId(err, biostar.DisconnectFromDeviceId, f2id)
	return buf.Bytes()
}

func enrollUser() {

	fmt.Printf("Enrolling user id %v into F2 %v ...", uid, f2id) // works Ok
	time.Now().Unix()
	req := &user.EnrollRequest{
		DeviceID: f2id,
		Users: []*user.UserInfo{
			&user.UserInfo{
				Hdr: &user.UserHdr{ID: uid, NumOfFace: 0},
				//Setting: &user.UserSetting{FaceAuthExtMode: uint32(auth.AuthMode_AUTH_EXT_MODE_FACE_ONLY)},
				Name: fmt.Sprintf("%s Great", uid),
				//Photo:   imgData,
				Setting: &user.UserSetting{
					StartTime:     uint32(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.Local).Unix()), // 0: no restrictions
					EndTime:       uint32(time.Date(2030, time.January, 1, 0, 0, 0, 0, time.Local).Unix()), // 0: no restrictions
					SecurityLevel: uint32(3),                                                               // normal
				},
			},
		},
		Overwrite: true,
	}
	_, err := userClient.Enroll(context.Background(), req)
	flow.IfErrorExitElseOkId(err, biostar.DisconnectFromDeviceId, f2id)
}

func enrollVisualFace(imgData []byte) {

	fmt.Printf("Normalizing image %v for user %v ...", imgFile, uid)
	respNorm, err := faceClient.Normalize(context.Background(), &face.NormalizeRequest{
		DeviceID:           f2id,
		UnwrappedImageData: imgData,
	})
	flow.IfErrorExitElseOkId(err, biostar.DisconnectFromDeviceId, f2id)

	/*
		fmt.Printf("Trying to extract template from image %v for user %v ...", imgFile, uid)
		respExt, err := faceClient.Extract(context.Background(), &face.ExtractRequest{
			DeviceID:  f2id,
			ImageData: respNorm.GetWrappedImageData(),
			IsWarped:  false,
		})
		ifErrorExitElseOk(err) // if it fails, need to figure out set of actions
	*/

	fmt.Printf("setting up face record for image %v and user %v ...", imgFile, uid)
	faceRec := &face.FaceData{
		Flag:      uint32(257),
		ImageData: respNorm.GetWrappedImageData(),
		//Templates: [][]byte{respExt.GetTemplateData()},
	}
	userFace := &user.UserFace{
		UserID: uid,
		Faces:  []*face.FaceData{faceRec},
	}
	faceReq := &user.SetFaceRequest{
		DeviceID:  f2id,
		UserFaces: []*user.UserFace{userFace},
	}
	_, err = userClient.SetFace(context.Background(), faceReq)
	flow.IfErrorExitElseOkId(err, biostar.DisconnectFromDeviceId, f2id)
}

func disconnectFromF2() {
	fmt.Printf("disconnecting from F2 %v ...", f2id)
	_, err := connectClient.Disconnect(context.Background(), &connect.DisconnectRequest{DeviceIDs: []uint32{f2id}})
	flow.IfErrorExitElseOkId(err, nil, 0)
}

func ifErrorExitElseOk(err error) {
	if err != nil {
		fmt.Printf(" %v\n", err)
		connectClient.Disconnect(context.Background(), &connect.DisconnectRequest{DeviceIDs: []uint32{f2id}})
		os.Exit(0)
	} else {
		fmt.Println(" Ok")
	}
}

func test() {
	errors := make([]*flow.EnrollError, 0)
	errors = append(errors, &flow.EnrollError{ErrorType: flow.RETRIEVE_IMAGE, ErrorMessage: fmt.Sprintf("There is some error %v", params.ImageDir)})
	errors = append(errors, &flow.EnrollError{ErrorType: flow.NORMALIZE_IMAGE, ErrorMessage: fmt.Sprintf("There is another error %v", params.GatewayIP)})
	errors = append(errors, &flow.EnrollError{ErrorType: flow.ENROLL_USER, ErrorMessage: fmt.Sprintf("There is yet another error %v", params.GatewayPort)})

	numErrors := len(errors)
	if numErrors > 0 {
		fmt.Printf("There are %v errors occurred\n", numErrors)
		for _, enrollError := range errors {
			fmt.Printf("Error occurred: %v [%v]\n", flow.EnrollErrorString(enrollError.ErrorType), enrollError.ErrorMessage)
		}
	}
}
