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
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/screenshot", api.SessionId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return []byte{}, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return []byte{}, errResp
	}

	// Manage response
	responseBody := StringResponse{}
	err := mapstructure.Decode(resp, &responseBody)
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
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/screenshot", api.SessionId, elementId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return []byte{}, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return []byte{}, errResp
	}

	// Manage response
	responseBody := StringResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return []byte{}, err
	}

	decodeString, err := base64.StdEncoding.DecodeString(responseBody.Value)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "an error occured during base 64 decoding process")
	}

	return decodeString, nil
}
