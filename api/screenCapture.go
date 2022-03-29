package api

import (
	"encoding/base64"
	"fmt"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// https://w3c.github.io/webdriver/#take-screenshot
func (api WebDriverApi) TakeScreenShot() ([]byte, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/screenshot", api.SessionId))
	if err != nil {
		log.Error("An error occured during take screenshot request: ", err)
		return []byte{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to take a screenshot: ", err)
		return []byte{}, err
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return []byte{}, err
	}

	decodeString, err := base64.StdEncoding.DecodeString(responseBody.Value)
	if err != nil {
		log.Error("An error occured during base 64 decoding process: ", err)
		return []byte{}, err
	}

	return decodeString, nil
}

// https://w3c.github.io/webdriver/#take-element-screenshot
func (api WebDriverApi) TakeElementScreenShot(elementId string) ([]byte, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/screenshot", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during take element screenshot request: ", err)
		return []byte{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to take element screenshot: ", err)
		return []byte{}, err
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return []byte{}, err
	}

	decodeString, err := base64.StdEncoding.DecodeString(responseBody.Value)
	if err != nil {
		log.Error("An error occured during base 64 decoding process: ", err)
		return []byte{}, err
	}

	return decodeString, nil
}
