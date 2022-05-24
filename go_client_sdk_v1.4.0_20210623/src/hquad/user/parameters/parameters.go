package parameters

import (
	"flag"
	"hquad/user/network"
)

type EnrollParameters struct {
	CertificateFile string
	GatewayIP       string
	GatewayPort     string
	ImageDir        string
}

func NewEnrollParameters() EnrollParameters {

	certFile := flag.String("cert", "data/cert/ca.crt", "Specify CA certificate file location")
	gateIP := flag.String("ip", network.GetLocalIP(), "Specify G-SDK gateway IP address")
	gatePort := flag.String("port", "4000", "Specify G-SDK gateway IP port")
	imageDir := flag.String("imageDir", "data/images/", "Specify dir containing user images files to enroll")
	flag.Parse()

	return EnrollParameters{
		CertificateFile: *certFile,
		GatewayIP:       *gateIP,
		GatewayPort:     *gatePort,
		ImageDir:        *imageDir,
	}
}
