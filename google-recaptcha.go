package util

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"encoding/json"

	"github.com/fanky5g/logger"
)

// CheckRecaptchaToken checks validates google recapture token
func CheckRecaptchaToken(clientIP, secret, responseToken string) (bool, error) {
	client := GetClient()

	apiUrl := "https://www.google.com"
	resource := "/recaptcha/api/siteverify"
	data := url.Values{}
	data.Set("secret", secret)
	data.Set("response", responseToken)
	data.Set("remoteip", clientIP)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		logger.SetLogLevel(logger.ErrorLevel)
		logger.Error(err)
		return false, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.SetLogLevel(logger.ErrorLevel)
		logger.Error(err)
		return false, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.SetLogLevel(logger.ErrorLevel)
		logger.Error(err)
		return false, err
	}

	if _, ok := response["error-codes"]; ok {
		if _, ok := response["error-codes"].([]interface{}); ok {
			errs := response["error-codes"].([]interface{})
			return false, errors.New(errs[0].(string))
		}
	} else if _, ok := response["success"]; ok {
		return response["success"].(bool), nil
	}

	return false, nil
}
