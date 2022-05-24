package auth

import (
	"fmt"
	"context"
	"biostar/service/auth"
	"google.golang.org/grpc"
)

type AuthSvc struct {
	client auth.AuthClient
}

func NewAuthSvc(conn *grpc.ClientConn) *AuthSvc {
	return &AuthSvc{
		client: auth.NewAuthClient(conn),
	}
}


func (s *AuthSvc) GetConfig(deviceID uint32) (*auth.AuthConfig, error) {
	req := &auth.GetConfigRequest{
		DeviceID: deviceID,
	}

	resp, err := s.client.GetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot get the auth config: %v\n", err)

		return nil, err
	}

	return resp.GetConfig(), nil
}


func (s *AuthSvc) SetConfig(deviceID uint32, config *auth.AuthConfig) error {
	req := &auth.SetConfigRequest{
		DeviceID: deviceID,
		Config: config,
	}

	_, err := s.client.SetConfig(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot set the auth config: %v\n", err)

		return err
	}

	return nil
}
