package util

import (
	"errors"
	"net/http"
	"strings"
)

func GetUserIpAddress(req *http.Request) (string, error) {
	forwarded := req.Header.Get("X-FORWARDED-FOR")
	ips := strings.Split(forwarded, ",")
	if len(ips) > 0 {
		return ips[0], nil
	}

	return "", errors.New("ip address not found")
}
