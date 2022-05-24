package network

import (
	"net"
	"os"
)

func GetLocalIP() string {

	host, _ := os.Hostname()
	addresses, _ := net.LookupIP(host)
	localIP := ""
	for _, addr := range addresses {
		if ipv4 := addr.To4(); ipv4 != nil {
			localIP = ipv4.String()
		}
	}
	return localIP
}
