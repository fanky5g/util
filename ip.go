package util

import (
	"errors"
	"net/http"
	"strings"
)

// GetUserIPAddress gets ip address from request
func GetUserIPAddress(req *http.Request) (string, error) {
	var IPAddress string

	originalForwardedFor := req.Header.Get("X-Original-Forwarded-For")
	if originalForwardedFor != "" {
		ips := strings.Split(originalForwardedFor, ",")
		if len(ips) > 0 {
			IPAddress = ips[0]
		}
	}

	forwarded := req.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		if len(ips) > 0 {
			IPAddress = ips[0]
		}
	}

	if IPAddress == "" {
		IPAddress = req.Header.Get("X-Real-Ip")
	}

	if IPAddress == "" {
		IPAddress = req.RemoteAddr
	}

	if IPAddress == "" {
		return "", errors.New("ip address not found")
	}

	return IPAddress, nil
}
