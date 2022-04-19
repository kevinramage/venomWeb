package api

import (
	"encoding/base64"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// https://w3c.github.io/webdriver/#take-screenshot
func (api WebDriverApi) TakeScreenShot() ([]byte, error) {

	// Security
	if api.SessionId == "" {
		return []byte{}, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/screenshot", api.SessionId))
	if err != nil {
		return []byte{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return []byte{}, fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return []byte{}, err
	}

	decodeString, err := base64.StdEncoding.DecodeString(responseBody.Value)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "an error occured during base 64 decoding process")
	}

	return decodeString, nil
}

// https://w3c.github.io/webdriver/#take-element-screenshot
func (api WebDriverApi) TakeElementScreenShot(elementId string) ([]byte, error) {

	// Security
	if api.SessionId == "" {
		return []byte{}, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/screenshot", api.SessionId, elementId))
	if err != nil {
		return []byte{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return []byte{}, fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return []byte{}, err
	}

	decodeString, err := base64.StdEncoding.DecodeString(responseBody.Value)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "an error occured during base 64 decoding process")
	}

	return decodeString, nil
}
