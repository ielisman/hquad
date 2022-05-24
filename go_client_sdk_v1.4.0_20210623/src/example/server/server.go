package server

import (
	"fmt"
	"context"
	"biostar/service/server"
	"biostar/service/user"
	"google.golang.org/grpc"
)

type ServerSvc struct {
	client server.ServerClient
}

func NewServerSvc(conn *grpc.ClientConn) *ServerSvc {
	return &ServerSvc{
		client: server.NewServerClient(conn),
	}
}

func (s *ServerSvc) Subscribe(queueSize int) (server.Server_SubscribeClient, context.CancelFunc, error) {
	req := &server.SubscribeRequest{
		QueueSize: int32(queueSize),
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	reqStream, err := s.client.Subscribe(ctx, req)

	if err != nil {
		fmt.Printf("Cannot subscribe to the server request stream: %v\n", err)

		return nil, nil, err
	}

	return reqStream, cancelFunc, nil
}

func (s *ServerSvc) Unsubscribe() error {
	req := &server.UnsubscribeRequest{
	}

	_, err := s.client.Unsubscribe(context.Background(), req)

	if err != nil {
		fmt.Printf("Cannot unsubscribe to the server request stream: %v\n", err)

		return err
	}

	return nil
}

func (s *ServerSvc) HandleVerify(serverReq *server.ServerRequest, errCode server.ServerErrorCode, userInfo *user.UserInfo) error {
	req := &server.HandleVerifyRequest{
		DeviceID: serverReq.DeviceID,
		SeqNO: serverReq.SeqNO,
		ErrCode: errCode,
		User: userInfo,
	}

	_, err := s.client.HandleVerify(context.Background(), req)
	return err
}


func (s *ServerSvc) HandleIdentify(serverReq *server.ServerRequest, errCode server.ServerErrorCode, userInfo *user.UserInfo) error {
	req := &server.HandleIdentifyRequest{
		DeviceID: serverReq.DeviceID,
		SeqNO: serverReq.SeqNO,
		ErrCode: errCode,
		User: userInfo,
	}

	_, err := s.client.HandleIdentify(context.Background(), req)
	return err
}
