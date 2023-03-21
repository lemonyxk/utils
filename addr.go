package utils

import (
	"encoding/binary"
	"errors"
	"net"
	"strconv"
	"strings"
)

var defaultIP = ""

type addr int

const Addr addr = iota

func (a addr) GetLocalhostIp() string {

	if defaultIP != "" {
		return defaultIP
	}

	addresses, err := net.InterfaceAddrs()

	if err != nil {
		defaultIP = "127.0.0.1"
		return defaultIP
	}

	for i := 0; i < len(addresses); i++ {
		// check the address type and if it is not a LoopBack the display it
		if ipNet, ok := addresses[i].(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				defaultIP = ipNet.IP.String()
				return defaultIP
			}
		}
	}

	defaultIP = "127.0.0.1"
	return defaultIP
}

func (a addr) Ip2long(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

func (a addr) IsLocalIP(ip string) bool {
	return a.IsLocalNet(net.ParseIP(ip))
}

// var localNetworks = []string{
// 	"10.0.0.0/8",
// 	"169.254.0.0/16",
// 	"172.16.0.0/12",
// 	"172.17.0.0/12",
// 	"172.18.0.0/12",
// 	"172.19.0.0/12",
// 	"172.20.0.0/12",
// 	"172.21.0.0/12",
// 	"172.22.0.0/12",
// 	"172.23.0.0/12",
// 	"172.24.0.0/12",
// 	"172.25.0.0/12",
// 	"172.26.0.0/12",
// 	"172.27.0.0/12",
// 	"172.28.0.0/12",
// 	"172.29.0.0/12",
// 	"172.30.0.0/12",
// 	"172.31.0.0/12",
// 	"192.168.0.0/16",
// }

func (a addr) IsLocalNet(ip net.IP) bool {

	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16

}

func (a addr) Parse(host string) (string, int, error) {

	var u = strings.Split(host, ":")

	if len(u) == 0 || len(u) > 2 {
		return "", 0, errors.New("invalid host")
	}

	if len(u) == 1 {
		return u[0], 80, nil
	}

	var ip = u[0]
	var port, _ = strconv.Atoi(u[1])

	return ip, port, nil
}
