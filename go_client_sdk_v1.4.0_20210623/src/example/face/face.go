package face

import (
	"fmt"
	"context"
	"biostar/service/face"
	"google.golang.org/grpc"
)

type FaceSvc struct {
	client face.FaceClient
}

func NewFaceSvc(conn *grpc.ClientConn) *FaceSvc {
	return &FaceSvc{
		client: face.NewFaceClient(conn),
	}
}

func (s *FaceSvc) Scan(deviceID uint32, threshold face.FaceEnrollThreshold) (*face.FaceData, error) {
	req := &face.ScanRequest{
		DeviceID: deviceID,
		EnrollThreshold: threshold,
	}

	resp, err := s.client.Scan(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot scan a face: %v\n", err)

		return nil, err
	}

	return resp.FaceData, nil
}


func (s *FaceSvc) GetConfig(deviceID uint32) (*face.FaceConfig, error) {
	req := &face.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the face config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}

