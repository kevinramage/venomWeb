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
		log.Error("An error occured during get session timeout request")
		return common.Timeouts{}, nil
	}

	// Manage response
	responseBody := GetTimeoutResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return common.Timeouts{}, nil
	}

	return responseBody.Value, err
}

// https://w3c.github.io/webdriver/#set-timeouts
func (api WebDriverApi) SetSessionTimeout(timeouts common.Timeouts) error {

	// Send request
	path := fmt.Sprintf("session/%s/timeouts", api.SessionId)
	_, err := ProceedPostRequest(api, path, timeouts)

	return err
}
