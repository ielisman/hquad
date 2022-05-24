package connectMaster

import (
	"fmt"
	"context"
	"biostar/service/connectMaster"
)

const (
	QUEUE_SIZE = 32
)

func (s *ConnectMasterSvc) Subscribe() (connectMaster.ConnectMaster_SubscribeStatusClient, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	req := &connectMaster.SubscribeStatusRequest{
		QueueSize: QUEUE_SIZE,
	}

	statusStream, err := s.client.SubscribeStatus(ctx, req)
	if err != nil {
		fmt.Printf("Cannot subscribe: %v", err)
		return nil, nil, err
	}

	return statusStream, cancel, nil
}


