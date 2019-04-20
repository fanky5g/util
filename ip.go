package util

import (
	"errors"
	"net/http"
	"strings"
)

func GetUserIpAddress(req *http.Request) (string, error) {
	IPAddress := req.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		forwarded := req.Header.Get("X-Forwarded-For")
		ips := strings.Split(forwarded, ",")
		if len(ips) > 0 {
			IPAddress = ips[0]
		}
	}

	if IPAddress == "" {
		IPAddress = req.RemoteAddr
	}

	if IPAddress == "" {
		return "", errors.New("ip address not found")
	}

	return IPAddress, nil
}
