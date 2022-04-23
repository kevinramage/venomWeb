package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// https://w3c.github.io/webdriver/#dismiss-alert
func (api WebDriverApi) DismissAlert() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/alert/dismiss", api.SessionId), nil)
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

// https://w3c.github.io/webdriver/#accept-alert
func (api WebDriverApi) AcceptAlert() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/alert/accept", api.SessionId), nil)
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

// https://w3c.github.io/webdriver/#get-alert-text
func (api WebDriverApi) GetAlertText() (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/alert/text", api.SessionId))
	if err != nil {
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return "", fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#send-alert-text
func (api WebDriverApi) SendAlertText(alertText string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request body
	type SendAlertTextBody struct {
		Text string `json:"text"`
	}
	requestBody := SendAlertTextBody{
		Text: alertText,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/alert/text", api.SessionId), requestBody)
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
