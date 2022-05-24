package main

import (
	"fmt"
	"biostar/service/server"
	"biostar/service/auth"
	"biostar/service/user"
	"biostar/service/card"
	"biostar/service/finger"
	"github.com/golang/protobuf/proto"
	"example/cli"
	"google.golang.org/grpc/status"	
	"google.golang.org/grpc/codes"		
)


func testServerMatching(deviceID uint32) error {
	// backup the auth config
	origAuthConfig, err := authSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	// enable server matching for test
	testConfig := proto.Clone(origAuthConfig).(*auth.AuthConfig)
	testConfig.UseServerMatching = true

	err = authSvc.SetConfig(deviceID, testConfig)
	if err != nil {
		return err
	}	

	testVerify(deviceID)
	testIdentify(deviceID)

	// restore the auth config
	authSvc.SetConfig(deviceID, origAuthConfig)

	return nil
}

const (
	QUEUE_SIZE = 16

	TEST_USER_ID = "1234"
)

func testVerify(deviceID uint32) error {
	reqStream, cancelFunc, err := serverSvc.Subscribe(QUEUE_SIZE)
	if err != nil {
		return err
	}

	fmt.Printf("\n===== Server Matching: Verify Test =====\n\n")	
	fmt.Printf(">> Try to authenticate a card. It should fail since the device gateway will return an error code to the request.\n")

	returnError := true

	go func() {
		for {
			req, err := reqStream.Recv()

			if err != nil {
				status, ok := status.FromError(err)
				if ok && status.Code() == codes.Canceled {
					fmt.Printf("Server matching is cancelled\n")
				} else {
					fmt.Printf("Cannot receive server matching requests: %v\n", err)
				}

				return
			}

			if req.ReqType != server.RequestType_VERIFY_REQUEST {
				fmt.Printf("!! Request type is not VERIFY_REQUEST. Just ignore it.\n")
				continue
			}

			if returnError {
				fmt.Printf("## Gateway returns VERIFY_FAIL.\n")
				serverSvc.HandleVerify(req, server.ServerErrorCode_VERIFY_FAIL, nil)
			} else {
				fmt.Printf("## Gateway returns SUCCESS with user information.\n")
				userInfo := &user.UserInfo{
					Hdr: &user.UserHdr{
						ID: TEST_USER_ID,
						NumOfCard: 1,
					},
					Cards: []*card.CSNCardData{
						&card.CSNCardData{
							Data: req.VerifyReq.CardData,
						},
					},
				}

				serverSvc.HandleVerify(req, server.ServerErrorCode_SUCCESS, userInfo)
			}
		}
	} ()

	cli.PressEnter(">> Press ENTER for the next test.\n")

	returnError = false
	fmt.Printf(">> Try to authenticate a card. The gateway will return SUCCESS with user information this time. The result will vary according to the authentication modes of the devices.\n")	

	cli.PressEnter(">> Press ENTER for the next test.\n")	

	cancelFunc()
	serverSvc.Unsubscribe()

	return nil
}


func testIdentify(deviceID uint32) error {
	reqStream, cancelFunc, err := serverSvc.Subscribe(QUEUE_SIZE)
	if err != nil {
		return err
	}

	fmt.Printf("\n===== Server Matching: Identify Test =====\n\n")	
	fmt.Printf(">> Try to authenticate a fingerprint. It should fail since the device gateway will return an error code to the request.\n")

	returnError := true

	go func() {
		for {
			req, err := reqStream.Recv()

			if err != nil {
				status, ok := status.FromError(err)
				if ok && status.Code() == codes.Canceled {
					fmt.Printf("Server matching is cancelled\n")
				} else {
					fmt.Printf("Cannot receive server matching requests: %v\n", err)
				}

				return
			}

			if req.ReqType != server.RequestType_IDENTIFY_REQUEST {
				fmt.Printf("!! Request type is not IDENTIFY_REQUEST. Just ignore it.\n")
				continue
			}

			if returnError {
				fmt.Printf("## Gateway returns IDENTIFY_FAIL.\n")
				serverSvc.HandleIdentify(req, server.ServerErrorCode_IDENTIFY_FAIL, nil)
			} else {
				fmt.Printf("## Gateway returns SUCCESS with user information.\n")
				userInfo := &user.UserInfo{
					Hdr: &user.UserHdr{
						ID: TEST_USER_ID,
						NumOfFinger: 1,
					},
					Fingers: []*finger.FingerData{
						&finger.FingerData{
							Templates: [][]byte{
								req.IdentifyReq.TemplateData,
								req.IdentifyReq.TemplateData,
							},
						},
					},
				}

				serverSvc.HandleIdentify(req, server.ServerErrorCode_SUCCESS, userInfo)
			}
		}
	} ()

	cli.PressEnter(">> Press ENTER for the next test.\n")

	returnError = false
	fmt.Printf(">> Try to authenticate a fingerprint. The gateway will return SUCCESS with user information this time. The result will vary according to the authentication modes of the devices.\n")	

	cli.PressEnter(">> Press ENTER for the next test.\n")	

	cancelFunc()
	serverSvc.Unsubscribe()

	return nil
}

