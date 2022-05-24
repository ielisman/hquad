package biostar

import (
	"biostar/service/connect"
	"biostar/service/device"
	"biostar/service/display"
	"biostar/service/face"
	"biostar/service/user"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"hquad/user/flow"
	"hquad/user/images"
	"hquad/user/parameters"
	"hquad/user/parser"
	"runtime"
	"strings"
	"sync"
	"time"
)

const ConnectionTimeout = 10000

var Conn *grpc.ClientConn
var ConnectClient connect.ConnectClient
var UserClient user.UserClient
var FaceClient face.FaceClient
var DeviceClient device.DeviceClient
var DisplayClient display.DisplayClient
var DefaultStartTime = uint32(time.Now().Add(-24 * 365 * time.Hour).Unix()) //uint32(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.Local).Unix())
var DefaultEndTime = uint32(time.Date(2029, time.December, 25, 01, 00, 00, 0, time.Local).Unix())

func ConnectToGateway(p *parameters.EnrollParameters) error {
	fmt.Printf("verifying gateway key exchange via certificate %v:%v ...\n", p.GatewayIP, p.GatewayPort)
	cred, err := credentials.NewClientTLSFromFile(p.CertificateFile, "")
	if err != nil {
		return err
	}

	fmt.Printf("making connection to gateway %v:%v ...\n", p.GatewayIP, p.GatewayPort)
	Conn, err = grpc.Dial(fmt.Sprintf("%v:%v", p.GatewayIP, p.GatewayPort),
		grpc.WithTransportCredentials(cred), grpc.WithBlock(), grpc.WithTimeout(ConnectionTimeout*time.Millisecond))
	if err != nil {
		return err
	}

	ConnectClient = connect.NewConnectClient(Conn)
	UserClient = user.NewUserClient(Conn)
	FaceClient = face.NewFaceClient(Conn)
	DeviceClient = device.NewDeviceClient(Conn)
	DisplayClient = display.NewDisplayClient(Conn)

	return nil
}

func DisconnectFromGateway() error {
	if Conn != nil {
		return Conn.Close()
	}
	return errors.New("gateway connection is empty")
}

func GetAllDevices() ([]*connect.SearchDeviceInfo, error) {

	fmt.Print("Searching for devices ...\n")
	searchReq := &connect.SearchDeviceRequest{
		Timeout: ConnectionTimeout,
	}
	searchResp, err := ConnectClient.SearchDevice(context.Background(), searchReq)
	if err != nil {
		return nil, err
	}

	return searchResp.GetDeviceInfos(), nil
}

func ConnectToDevice(dvc *connect.SearchDeviceInfo) error {
	fmt.Printf("Connecting to %v %v with address %v:%v ...\n", dvc.GetType(), dvc.GetDeviceID(), dvc.GetIPAddr(), dvc.GetPort())
	_, err := ConnectClient.Connect(context.Background(), &connect.ConnectRequest{
		ConnectInfo: &connect.ConnectInfo{
			IPAddr: dvc.GetIPAddr(),
			Port:   dvc.GetPort(),
			UseSSL: dvc.GetUseSSL(),
		},
	})
	return err
}

func GetDeviceCapability(deviceInfo *connect.SearchDeviceInfo) (*device.CapabilityInfo, error) {
	fmt.Printf("Retrieving Device Capability for device %v ...\n", deviceInfo.GetDeviceID())
	respCap, err := DeviceClient.GetCapabilityInfo(context.Background(), &device.GetCapabilityInfoRequest{
		DeviceID: deviceInfo.GetDeviceID(),
	})
	if err != nil {
		return nil, err
	}
	return respCap.GetCapInfo(), nil
}

func EnrollUsers(dvc *connect.SearchDeviceInfo, faceUsers []*parser.ImageUser) error {
	fmt.Printf("enrolling %d users for %v %v\n", len(faceUsers), dvc.GetType(), dvc.GetDeviceID())

	start := time.Now().Unix()
	for _, faceUser := range faceUsers {
		retrieveTime := time.Now().UnixMilli()
		faceImage, err := images.GetFaceImage(faceUser.ImageFile)
		if ok := flow.ErrorFlowStandard(err); !ok {
			// save row error and return with message can't get image
		}
		fmt.Printf("It took %v milliseconds to retrieve %s image\n", time.Now().UnixMilli()-retrieveTime, faceUser.ImageFile)
		normalizeTime := time.Now().UnixMilli()
		_, err = GetNormalized(dvc, faceUser, faceImage)
		if ok := flow.ErrorFlowStandard(err); !ok {
			// save row error and return with message can't normalize image
		}
		fmt.Printf("It took %v milliseconds to normalize %s image\n", time.Now().UnixMilli()-normalizeTime, faceUser.ImageFile)
	}
	fmt.Printf("It took %v seconds to process %d records sequentially\n", time.Now().Unix()-start, len(faceUsers))

	return nil
}

func EnrollUsersConcurrent(dvc *connect.SearchDeviceInfo, faceUsers []*parser.ImageUser) error {
	fmt.Printf("enrolling %d users for %v %v\n", len(faceUsers), dvc.GetType(), dvc.GetDeviceID())

	start := time.Now().UnixMilli()
	g := runtime.NumCPU() // ~ 13 ms per retrieving image via concurrent (50 images ~ 30mb) - batch should be 200 images max (120Mb)
	var wg sync.WaitGroup
	wg.Add(g)
	chFaceUser := make(chan *parser.ImageUser, g)
	chFaceImage := make(chan *images.FaceImage)
	l := len(faceUsers)
	numRec := l
	var mx sync.Mutex
	errors := make([]*flow.EnrollError, 0)

	for i := 0; i < g; i++ {
		go func() {
			defer wg.Done()
			for faceUser := range chFaceUser { // RETRIEVING IMAGE
				faceImage, err := images.GetFaceImage(faceUser.ImageFile)
				if ok := flow.ErrorFlowStandardMessage(err, fmt.Sprintf("finished retrieving image %v for user %v", faceImage.GetImgFile(), faceUser.GetUserName())); !ok {
					mx.Lock()
					errors = append(errors, &flow.EnrollError{ErrorType: flow.RETRIEVE_IMAGE, ErrorMessage: fmt.Sprintf("Cant retrieve %v", faceUser.GetImageFile())})
					mx.Unlock()
				}
				go func(fi *images.FaceImage) {
					chFaceImage <- fi
				}(faceImage)
			}
		}()
	}

	// submit all users to channel to parse images
	faceUsersMap := make(map[string]*parser.ImageUser, len(faceUsers))
	for _, faceUser := range faceUsers {
		faceUsersMap[faceUser.ImageFile] = faceUser
		chFaceUser <- faceUser
	}
	close(chFaceUser)
	wg.Wait() // finished retrieving images
	fmt.Printf("It took %v milliseconds to process %d images concurrently using %d go routines\n", time.Now().UnixMilli()-start, len(faceUsers), g)

	// enrolling
	start = time.Now().Unix()
	h := 3
	var wgEnroll sync.WaitGroup
	chFaceImageEnrolled := make(chan *images.FaceImage, h)
	wgEnroll.Add(h)

	for i := 0; i < h; i++ {
		go func() {
			defer wgEnroll.Done()
			for faceImage := range chFaceImageEnrolled {
				faceUser := faceUsersMap[faceImage.GetImgFile()]
				normalizedImg, err := GetNormalized(dvc, faceUser, faceImage) // NORMALIZING IMAGE
				if ok := flow.ErrorFlowStandardMessage(err, fmt.Sprintf("finished normalizing image %v for user %v", faceImage.GetImgFile(), faceUser.GetUserName())); !ok {
					mx.Lock()
					errors = append(errors, &flow.EnrollError{ErrorType: flow.NORMALIZE_IMAGE, ErrorMessage: fmt.Sprintf("Cant normalize %v", faceUser.GetImageFile())})
					mx.Unlock()
				} else { // enroll user
					err = EnrollUserWithNormalizedFace(dvc, faceUser, normalizedImg, DefaultStartTime, DefaultEndTime)
					if ok := flow.ErrorFlowStandardMessage(err, fmt.Sprintf("finished enrolling user %v with normalized image %v", faceUser.GetUserName(), faceImage.GetImgFile())); !ok {
						mx.Lock()
						errors = append(errors, &flow.EnrollError{ErrorType: flow.ENROLL_USER, ErrorMessage: fmt.Sprintf("Cant enroll user %v (%v)", faceUser.GetUserName(), faceUser.GetUserId())})
						mx.Unlock()
					}
				}
			}
		}()
	}

	for l > 0 {
		faceImage := <-chFaceImage
		if faceImage != nil {
			chFaceImageEnrolled <- faceImage
		}
		l--
		fmt.Println("Retrieved images remains =>", l)
	}
	close(chFaceImageEnrolled)
	wgEnroll.Wait() // finished with retrieving images
	fmt.Printf("It took %v seconds to enroll %d users concurrently using %d go routines\n", time.Now().Unix()-start, len(faceUsers), h)

	numErrors := len(errors)
	if numErrors > 0 {
		fmt.Printf("There are %v errors occurred when processing %d records\n", numErrors, numRec)
		for _, enrollError := range errors {
			fmt.Printf("%v error: %v\n", flow.EnrollErrorString(enrollError.ErrorType), enrollError.ErrorMessage)
		}
	}

	return nil
}

func EnrollUser(dvc *connect.SearchDeviceInfo, userId, userName string, startTime, endTime uint32) error {

	fmt.Printf("Enrolling user %v (id %v) into %v ...\n", userName, userId, dvc.GetDeviceID())
	req := &user.EnrollRequest{
		DeviceID: dvc.GetDeviceID(),
		Users: []*user.UserInfo{
			&user.UserInfo{
				Hdr:  &user.UserHdr{ID: userId, NumOfFace: 0},
				Name: userName,
				Setting: &user.UserSetting{
					StartTime: startTime, // uint32(time.Date(2001, time.January, 1, 0, 0, 0, 0, time.Local).Unix()), // 0: no restrictions
					EndTime:   endTime,
					//SecurityLevel: 3, // normal
				},
			},
		},
		Overwrite: true,
	}
	_, err := UserClient.Enroll(context.Background(), req)
	return err
}

func EnrollUserWithFace(dvc *connect.SearchDeviceInfo, faceUser *parser.ImageUser, startTime, endTime uint32) error {

	fmt.Printf("Enrolling user %v (id %v) into %v ...\n", faceUser.GetUserName(), faceUser.GetUserId(), dvc.GetDeviceID())
	faceImage, err := images.GetFaceImage(faceUser.GetImageFile())
	if err != nil {
		return err
	}
	normalizedImgData, err := GetNormalized(dvc, faceUser, faceImage)
	if err != nil {
		return err
	}

	return EnrollUserWithNormalizedFace(dvc, faceUser, normalizedImgData, startTime, endTime)

}

func EnrollUserWithNormalizedFace(dvc *connect.SearchDeviceInfo, faceUser *parser.ImageUser, normalizedImgData []byte, startTime, endTime uint32) error {
	fmt.Printf("Enrolling user %v (id %v) with normalized face image into %v ...\n", faceUser.GetUserName(), faceUser.GetUserId(), dvc.GetDeviceID())
	req := &user.EnrollRequest{
		DeviceID: dvc.GetDeviceID(),
		Users: []*user.UserInfo{
			&user.UserInfo{
				Hdr:  &user.UserHdr{ID: faceUser.GetUserId(), NumOfFace: 0},
				Name: faceUser.GetUserName(),
				Setting: &user.UserSetting{
					StartTime: startTime,
					EndTime:   endTime,
					//SecurityLevel: 3, // normal
				},
				Faces: []*face.FaceData{
					&face.FaceData{
						Flag:      uint32(257), // this is bug when fixed should be face.FaceFlag_BS2_FACE_FLAG_F2
						ImageData: normalizedImgData,
					},
				},
			},
		},
		Overwrite: true,
	}
	_, err := UserClient.Enroll(context.Background(), req)
	return err
}

func EnrollVisualFace(dvc *connect.SearchDeviceInfo, userId, userName string, normalizedImgData []byte) error {

	fmt.Printf("setting up face record for user %v (id %v) ...\n", userName, userId)
	userFace := &user.UserFace{
		UserID: userId,
		Faces: []*face.FaceData{
			&face.FaceData{
				Flag:      uint32(257), // this is bug when fixed should be face.FaceFlag_BS2_FACE_FLAG_F2
				ImageData: normalizedImgData,
			},
		},
	}
	faceReq := &user.SetFaceRequest{
		DeviceID:  dvc.GetDeviceID(),
		UserFaces: []*user.UserFace{userFace},
	}
	_, err := UserClient.SetFace(context.Background(), faceReq)
	return err
}

func EnrollVisualFaceWithNormalization(dvc *connect.SearchDeviceInfo, uid string, faceImage *images.FaceImage) error {

	fmt.Printf("Normalizing image %v for user %v ...\n", faceImage.GetImgFile(), uid)
	respNorm, err := FaceClient.Normalize(context.Background(), &face.NormalizeRequest{
		DeviceID:           dvc.GetDeviceID(),
		UnwrappedImageData: faceImage.GetImgData(),
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to normalize image %v for user id %v device %v : [%v]", faceImage.GetImgFile(), uid, dvc.GetDeviceID(), err))
	}

	fmt.Printf("setting up face record for normalized image %v and user %v ...\n", faceImage.GetImgFile(), uid)
	userFace := &user.UserFace{
		UserID: uid,
		Faces: []*face.FaceData{
			&face.FaceData{
				Flag:      uint32(257), // this is bug when fixed should be face.FaceFlag_BS2_FACE_FLAG_F2
				ImageData: respNorm.GetWrappedImageData(),
			},
		},
	}
	faceReq := &user.SetFaceRequest{
		DeviceID:  dvc.GetDeviceID(),
		UserFaces: []*user.UserFace{userFace},
	}
	_, err = UserClient.SetFace(context.Background(), faceReq)
	return err
}

func GetNormalized(dvc *connect.SearchDeviceInfo, faceUser *parser.ImageUser, faceImage *images.FaceImage) ([]byte, error) {
	fmt.Printf("Normalizing image %v for user %v ...\n", faceImage.GetImgFile(), faceUser.UserId)
	respNorm, err := FaceClient.Normalize(context.Background(), &face.NormalizeRequest{
		DeviceID:           dvc.GetDeviceID(),
		UnwrappedImageData: faceImage.GetImgData(),
	})
	if err != nil {
		return nil, errors.New(
			fmt.Sprintf("Unable to normalize image %v for user id %v device %v : [%v]",
				faceImage.GetImgFile(), faceUser.UserId, dvc.GetDeviceID(), err))
	}
	return respNorm.GetWrappedImageData(), nil

}

// GetUsersList returns list of users.
//  deviceID - id of device such as FaceStation F2
//  infoMask - if specified returns list of partial user fields. If provided number < 0, all fields are returned
//             Examples:
//               * user.InfoMask_USER_MASK_CARD | user.InfoMask_USER_MASK_PHOTO | user.InfoMask_USER_MASK_SETTING
//               * -1
//               * user.InfoMask_USER_MASK_FACE
//             If value under infoMask is provided, user.InfoMask_USER_MASK_NAME will be added to it
//	filter   - list of names. If only one is provided, it will be used as regular expression search
func GetUsersList(deviceID uint32, infoMask int, filter ...string) ([]*user.UserInfo, error) {
	fmt.Printf("Getting list of users from device %v and filter <%v> ...\n", deviceID, filter)
	resp, err := UserClient.GetList(context.Background(), &user.GetListRequest{
		DeviceID: deviceID,
	})
	if err != nil {
		return nil, err
	}
	headerLen := len(resp.GetHdrs())
	if headerLen == 0 {
		return nil, nil
	}
	ids := make([]string, headerLen)
	for i := range resp.GetHdrs() {
		ids[i] = resp.GetHdrs()[i].GetID()
	}

	var users []*user.UserInfo
	if infoMask > -1 {
		if (infoMask & int(user.InfoMask_USER_MASK_NAME)) == 0 { // include name if it is not included
			infoMask |= int(user.InfoMask_USER_MASK_NAME)
		}
		respUsers, err := UserClient.GetPartial(context.Background(), &user.GetPartialRequest{
			DeviceID: deviceID,
			UserIDs:  ids,
			InfoMask: uint32(infoMask),
		})
		if err != nil {
			return nil, err
		}
		users = respUsers.GetUsers()
	} else {
		respUsers, err := UserClient.Get(
			context.Background(),
			&user.GetRequest{
				DeviceID: deviceID,
				UserIDs:  ids,
			},
			grpc.MaxCallRecvMsgSize(1024*1024*1024*5), // 5Gb - assuming 5000 users have 1Mb of data on average
		)

		if err != nil {
			return nil, err
		}
		users = respUsers.GetUsers()
	}
	if filter != nil && len(filter) > 0 {
		filteredUsers := make([]*user.UserInfo, 0, headerLen)
		if len(filter) == 1 {
			reString := filter[0]
			for _, user := range users {
				if strings.Contains(user.GetName(), reString) {
					filteredUsers = append(filteredUsers, user)
				}
			}
		} else {
			userMap := make(map[string]bool)
			for _, val := range filter {
				userMap[val] = true
			}
			for _, user := range users {
				if _, ok := userMap[user.GetName()]; ok { // "if userMap[user.GetName()] {}" should work too
					filteredUsers = append(filteredUsers, user)
				}
			}
		}
		return filteredUsers, nil
	}
	return users, nil
}

func GetUserHeadersList(deviceID uint32, filter ...string) ([]*user.UserHdr, error) {
	fmt.Printf("Getting list of user headers from device %v and filter <%v> ...\n", deviceID, filter)
	resp, err := UserClient.GetList(context.Background(), &user.GetListRequest{
		DeviceID: deviceID,
	})
	if err != nil {
		return nil, err
	}

	hdrLen := len(resp.GetHdrs())
	if filter != nil && len(filter) > 0 && hdrLen > 0 {
		filteredUserHeaders := make([]*user.UserHdr, 0, hdrLen)
		if len(filter) == 1 { // use regular expression
			reString := filter[0]
			for _, hdr := range resp.GetHdrs() {
				if strings.Contains(hdr.GetID(), reString) {
					filteredUserHeaders = append(filteredUserHeaders, hdr)
				}
			}
		} else { // use equality
			hdrMap := make(map[string]bool)
			for _, val := range filter {
				hdrMap[val] = true
			}
			for _, hdr := range resp.GetHdrs() {
				if _, ok := hdrMap[hdr.GetID()]; ok { // "if hdrMap[hdr.GetID()] {}" should work too
					filteredUserHeaders = append(filteredUserHeaders, hdr)
				}
			}
		}
		return filteredUserHeaders, nil
	}

	return resp.GetHdrs(), nil

}

func GetDisplayConfig(dvc *connect.SearchDeviceInfo) (*display.DisplayConfig, error) {
	fmt.Printf("Getting Display Config from %v %v with address %v:%v ...\n", dvc.GetType(), dvc.GetDeviceID(), dvc.GetIPAddr(), dvc.GetPort())
	resp, err := DisplayClient.GetConfig(context.Background(), &display.GetConfigRequest{
		DeviceID: dvc.GetDeviceID(),
	})
	if err != nil {
		return nil, err
	}
	return resp.GetConfig(), nil
}

func SetMessageTimeout(dvc *connect.SearchDeviceInfo, config *display.DisplayConfig, msgTimeout uint32) error {
	fmt.Printf("Getting Display Config from %v %v with address %v:%v ...\n", dvc.GetType(), dvc.GetDeviceID(), dvc.GetIPAddr(), dvc.GetPort())
	_, err := DisplayClient.SetConfig(context.Background(), &display.SetConfigRequest{
		DeviceID: dvc.GetDeviceID(),
		Config: &display.DisplayConfig{
			Language:         config.GetLanguage(),
			Background:       config.GetBackground(), // display.BackgroundType_BS2_BG_SLIDE
			Theme:            config.GetTheme(),
			Volume:           config.GetVolume(),
			UseVoice:         config.GetUseVoice(),
			DateFormat:       config.GetDateFormat(),
			TimeFormat:       config.GetTimeFormat(),
			ShowDateTime:     config.GetShowDateTime(),
			MenuTimeout:      config.GetMenuTimeout(),
			MsgTimeout:       msgTimeout,
			BacklightTimeout: config.GetBacklightTimeout(),
			UseUserPhrase:    config.GetUseUserPhrase(),
			QueryUserPhrase:  config.GetQueryUserPhrase(),
		},
	})
	return err
}

func SetSlideImages(dvc *connect.SearchDeviceInfo, images [][]byte) error {
	config, err := GetDisplayConfig(dvc)
	if err != nil {
		return err
	}
	if config.Background != display.BackgroundType_BS2_BG_SLIDE {
		config.Background = display.BackgroundType_BS2_BG_SLIDE
		_, err = DisplayClient.SetConfig(context.Background(), &display.SetConfigRequest{
			DeviceID: dvc.GetDeviceID(),
			Config:   config,
		})
		if err != nil {
			return err
		}
	}
	fmt.Printf("Setting %v slide images for %v %v with address %v:%v ...\n", len(images), dvc.GetType(), dvc.GetDeviceID(), dvc.GetIPAddr(), dvc.GetPort())
	_, err = DisplayClient.UpdateSlideImages(context.Background(), &display.UpdateSlideImagesRequest{
		DeviceID:  dvc.GetDeviceID(),
		PNGImages: images,
	})
	return err
}

func DisconnectFromDevice(dvc *connect.SearchDeviceInfo) error {
	fmt.Printf("Disconnecting from %v %v with address %v:%v ...\n", dvc.GetType(), dvc.GetDeviceID(), dvc.GetIPAddr(), dvc.GetPort())
	_, err := ConnectClient.Disconnect(context.Background(), &connect.DisconnectRequest{DeviceIDs: []uint32{dvc.GetDeviceID()}})
	return err
}

func DisconnectFromDeviceId(deviceID uint32) error {
	fmt.Printf("Disconnecting from %v ...\n", deviceID)
	_, err := ConnectClient.Disconnect(context.Background(), &connect.DisconnectRequest{DeviceIDs: []uint32{deviceID}})
	return err
}

func RebootDevice(dvc *connect.SearchDeviceInfo) error {
	fmt.Printf("Rebooting from %v %v with address %v:%v ...\n", dvc.GetType(), dvc.GetDeviceID(), dvc.GetIPAddr(), dvc.GetPort())
	_, err := DeviceClient.Reboot(context.Background(), &device.RebootRequest{
		DeviceID: dvc.GetDeviceID(),
	})
	return err
}
