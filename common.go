package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	mathrand "math/rand"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"

	"github.com/fanky5g/logger"
)

var (
	client *http.Client
)

// CreateHMACFromStruct creates sha256 hmac string from struct
// @todo:ensure body is struct
func CreateHMACFromStruct(body interface{}, secret string) (string, error) {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)

	// encode body
	b, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	written, err := h.Write(b)
	if err != nil {
		return "", err
	}

	if written != len(b) {
		return "", fmt.Errorf("written in part")
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// CheckMAC checks signatures if matches
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

// GetCurrentTimestamp formats current time in "2006-01-02 15:04:05 Z0700" layout
func GetCurrentTimestamp() string {
	return GetCurrentTimestampWithLayout("2006-01-02 15:04:05 ZO700")
}

// GetCurrentTimestampWithLayout returns current timestamp with specified layout
func GetCurrentTimestampWithLayout(layout string) string {
	return time.Now().Format(layout)
}

func createClient(cookieJar http.CookieJar) *http.Client {
	var t = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			// Timeout:   60 * time.Second,
			// KeepAlive: 30 * time.Second,
			Timeout:   0,
			KeepAlive: 0,
		}).Dial,
		TLSHandshakeTimeout: 30 * time.Second,
	}

	var client = &http.Client{
		Transport: t,
		Jar:       cookieJar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return client
}

// GetClient returns an http client for transactions
func GetClient() *http.Client {
	if client == nil {
		client = createClient(nil)
		return client
	}

	return client
}

// GetClientWithCookieJar creates a new http client passing along cookie jar
func GetClientWithCookieJar(jar http.CookieJar) *http.Client {
	return createClient(jar)
}

// GetProxyClient returns a socks5 proxy client
func GetProxyClient(address string) (*http.Client, error) {
	proxyDialer, err := proxy.SOCKS5("tcp", address, nil, proxy.Direct)
	if err != nil {
		logger.SetLogLevel(logger.ErrorLevel)
		logger.Debug(fmt.Sprintf("Failed to get proxy: %v", err))
		return nil, err
	}

	var t = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		Dial:                proxyDialer.Dial,
		TLSHandshakeTimeout: 30 * time.Second,
	}

	var client = &http.Client{
		Transport: t,
	}

	return client, nil
}

// NewRequest creates an http request and sets connection close header
func NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, res := http.NewRequest(method, url, body)
	req.Header.Set("Connection", "close") // set a connection close header for all application scoped requests

	return req, res
}

// GenerateVerificationCode generates 6 digit verification string for mobile
func GenerateVerificationCode() string {
	vals := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = vals[int(b[i])%len(vals)]
	}
	return string(b)
}

// PickRandomSliceString returns random element from a slice of strings
func PickRandomSliceString(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	mathrand.Seed(time.Now().Unix())
	return slice[mathrand.Intn(len(slice))]
}

// RoundFloat rounds floats to the nearest places provided. roundOn is a value to round by example .5
func RoundFloat(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
