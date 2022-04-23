package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
)

// https://w3c.github.io/webdriver/#get-timeouts
func (api WebDriverApi) GetSessionTimeout() (common.Timeouts, error) {

	// Security
	if api.SessionId == "" {
		return common.Timeouts{}, fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/timeouts", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		return common.Timeouts{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return common.Timeouts{}, fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := GetTimeoutResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return common.Timeouts{}, err
	}

	return responseBody.Value, err
}

// https://w3c.github.io/webdriver/#set-timeouts
func (api WebDriverApi) SetSessionTimeout(timeouts common.Timeouts) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/timeouts", api.SessionId)
	resp, err := ProceedPostRequest(api, path, timeouts)
	if err != nil {
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return fmt.Errorf(responseError.Value.Message)
	}

	return nil
}
