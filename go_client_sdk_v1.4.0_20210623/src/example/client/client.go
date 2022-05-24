package client

import (
	"time"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	CONN_TIMEOUT = 10000 * time.Millisecond
)


type GrpcClient interface {
	GetConn() *grpc.ClientConn
	Close()
}

type GatewayClient struct {
	conn *grpc.ClientConn
}


func (c *GatewayClient) GetConn() *grpc.ClientConn {
	return c.conn
}


func (c *GatewayClient) Connect(certFile string, serverIP string, serverPort int) error {
	creds, err := credentials.NewClientTLSFromFile(certFile, "")

	if err != nil {
		fmt.Printf("Cannot make the client: %v\n", err)
		return err
	}

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", serverIP, serverPort), grpc.WithTransportCredentials(creds), grpc.WithBlock(), grpc.WithTimeout(CONN_TIMEOUT))

	if err != nil {
		fmt.Printf("Cannot connect to the server %v:%v %v\n", serverIP, serverPort, err)
		return err
	}

	c.conn = conn

	return nil
}


func (c *GatewayClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	c.conn = nil
}