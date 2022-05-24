package master

import (
	"time"
	"fmt"
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"biostar/service/login"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	ADMIN_TENANT_ID = "administrator"

	CONN_TIMEOUT = 10000 * time.Millisecond
)


type MasterClient struct {
	conn *grpc.ClientConn
}


func (c *MasterClient) GetConn() *grpc.ClientConn {
	return c.conn
}

func (c *MasterClient) ConnectAdmin(caCertFile, adminCertFile, adminKeyFile, masterIP string, masterPort int) error {
	clientCert, err := tls.LoadX509KeyPair(adminCertFile, adminKeyFile)
	if err != nil {
		fmt.Printf("Cannot load the admin certificate: %v", err)
		return err
	}

	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		fmt.Printf("Cannot load the root CA certificate: %v", err)
		return err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config {
		Certificates: []tls.Certificate{ clientCert },
		RootCAs: caCertPool,
	}

	var tokenCreds JWTCredential

	c.conn, err = grpc.Dial(fmt.Sprintf("%v:%v", masterIP, masterPort), grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), grpc.WithPerRPCCredentials(&tokenCreds))
	if err != nil {
		fmt.Printf("Cannot dial to the master gateway: %v", err)
		return err
	}

	adminCertData, _ := ioutil.ReadFile(adminCertFile)

	loginClient := login.NewLoginClient(c.conn)

	loginReq := &login.LoginAdminRequest{
		AdminTenantCert: string(adminCertData),
		TenantID: ADMIN_TENANT_ID,
	}

	loginResp, err := loginClient.LoginAdmin(context.Background(), loginReq)
	if err != nil {
		fmt.Printf("Cannot login as an administrator: %v", err)
		return err
	}

	tokenCreds.Token = loginResp.JwtToken	
	
	return nil
}


func (c *MasterClient) ConnectTenant(caCertFile, tenantCertFile, tenantKeyFile, masterIP string, masterPort int) error {
	clientCert, err := tls.LoadX509KeyPair(tenantCertFile, tenantKeyFile)
	if err != nil {
		fmt.Printf("Cannot load the tenant certificate: %v", err)
		return err
	}

	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		fmt.Printf("Cannot load the root CA certificate: %v", err)
		return err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config {
		Certificates: []tls.Certificate{ clientCert },
		RootCAs: caCertPool,
	}

	var tokenCreds JWTCredential

	c.conn, err = grpc.Dial(fmt.Sprintf("%v:%v", masterIP, masterPort), grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), grpc.WithPerRPCCredentials(&tokenCreds))
	if err != nil {
		fmt.Printf("Cannot dial to the master gateway: %v", err)
		return err
	}

	tenantCertData, _ := ioutil.ReadFile(tenantCertFile)

	loginClient := login.NewLoginClient(c.conn)

	loginReq := &login.LoginRequest{
		TenantCert: string(tenantCertData),
	}

	loginResp, err := loginClient.Login(context.Background(), loginReq)
	if err != nil {
		fmt.Printf("Cannot login as a tenant: %v", err)
		return err
	}

	tokenCreds.Token = loginResp.JwtToken

	return nil
}


func (c *MasterClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	c.conn = nil
}