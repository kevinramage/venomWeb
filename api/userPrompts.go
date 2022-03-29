package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// https://w3c.github.io/webdriver/#dismiss-alert
func (api WebDriverApi) DismissAlert() error {

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/alert/dismiss", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during dismiss alert request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to dismiss alert: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#accept-alert
func (api WebDriverApi) AcceptAlert() error {

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/alert/accept", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during accept alert request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to accept alert: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-alert-text
func (api WebDriverApi) GetAlertText() (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/alert/text", api.SessionId))
	if err != nil {
		log.Error("An error occured during get alert text request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get alert text: ", err)
		return "", err
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#send-alert-text
func (api WebDriverApi) SendAlertText(alertText string) error {

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
		log.Error("An error during url retrieve call: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get current url: ", err)
		return fmt.Errorf("impossible to get current url")
	}

	return nil
}
