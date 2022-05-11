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
	resp, errResp := ProceedGetRequest(api, path)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return common.Timeouts{}, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return common.Timeouts{}, errResp
	}

	// Manage response
	responseBody := GetTimeoutResponse{}
	err := mapstructure.Decode(resp, &responseBody)
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
	resp, errResp := ProceedPostRequest(api, path, timeouts)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}
