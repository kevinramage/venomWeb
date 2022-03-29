package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// https://w3c.github.io/webdriver/#get-timeouts
func (api WebDriverApi) GetSessionTimeout() (common.Timeouts, error) {

	// Send request
	path := fmt.Sprintf("session/%s/timeouts", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get session timeout request: ", err)
		return common.Timeouts{}, nil
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get session timeout: ", responseError.Value.Message)
		return common.Timeouts{}, fmt.Errorf("impossible to get session timeout")
	}

	// Manage response
	responseBody := GetTimeoutResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return common.Timeouts{}, err
	}

	return responseBody.Value, err
}

// https://w3c.github.io/webdriver/#set-timeouts
func (api WebDriverApi) SetSessionTimeout(timeouts common.Timeouts) error {

	// Send request
	path := fmt.Sprintf("session/%s/timeouts", api.SessionId)
	resp, err := ProceedPostRequest(api, path, timeouts)
	if err != nil {
		log.Error("An error occured during set session timeout request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to set session timeout: ", responseError.Value.Message)
		return fmt.Errorf("impossible to set session timeout: %v", timeouts)
	}

	return nil
}
