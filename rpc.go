package util

import (
	"errors"
	"strings"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/fanky5g/logger"
	"google.golang.org/grpc"
)

const (
	// MaxConcurrentRequests represents number of parallel requests we can run
	MaxConcurrentRequests = 100
	// ErrorPercentThreshold represents ...
	ErrorPercentThreshold = 1
)

var (
	// Timeout value for when to drop connection if no response has been received
	Timeout      = 30000
	commandstore = make(map[string]interface{})
)

// SetRPCTimeout sets rpc timeout used for connections
func SetRPCTimeout(timeout int) {
	Timeout = timeout
}

// RPC calls an internal service with retry mechanisms as well as circuit breakers for resilience
func RPC(address string, defaultResponse interface{}, handle func(*grpc.ClientConn) (interface{}, error)) (interface{}, error) {
	dialParams := strings.Split(address, ":")
	if len(dialParams) < 2 {
		return nil, errors.New("address must be in valid service:port format")
	}

	name := dialParams[0]
	if dialParams[0] == "" {
		return nil, errors.New("service name missing")
	}

	if dialParams[1] == "" {
		return nil, errors.New("no port specified")
	}

	// if command hasn't been configured, do so
	if _, ok := commandstore[name]; !ok {
		hystrix.ConfigureCommand(name, hystrix.CommandConfig{
			Timeout:               Timeout,
			MaxConcurrentRequests: MaxConcurrentRequests,
			ErrorPercentThreshold: ErrorPercentThreshold,
		})

		commandstore[name] = true
	}

	var response interface{}
	err := hystrix.Do(name, func() error {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			return err
		}

		defer conn.Close()
		r, err := handle(conn)
		if err != nil {
			return err
		}

		response = r
		return nil
	}, func(e error) error {
		response = defaultResponse
		logger.SetLogLevel(logger.ErrorLevel)
		logger.Error(e)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
