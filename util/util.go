package util

import (
	"errors"
	"net"
)

//math.MaxUint32
func EnsureRange(value *int32, min int32, max int32) {
	if *value < min {
		*value = min
	}

	if *value > max {
		*value = max
	}
}

//config Split Char
const (
	OneSplitChar   string = "#"
	TwoSplitChar   string = "|"
	ThreeSplitChar string = "^"
)

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Local Ip Not Find")
}
