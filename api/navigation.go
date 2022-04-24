package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type GetTitleResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

// https://w3c.github.io/webdriver/#navigate-to
func (api WebDriverApi) Navigate(url string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request body
	type NavigateRequest struct {
		Url string `json:"url"`
	}
	request := NavigateRequest{
		Url: url,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/url", api.SessionId), request)
	if err != nil {
		return err
	}

	// Manage error
	responseError := ErrorResponse{}
	errDecode := mapstructure.Decode(resp, &responseError)
	if errDecode == nil && responseError.Value.Message != "" {
		return fmt.Errorf(responseError.Value.Message)
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-current-url
func (api WebDriverApi) GetCurrentUrl() (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/url", api.SessionId))
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

// https://w3c.github.io/webdriver/#back
func (api WebDriverApi) Back() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/back", api.SessionId), nil)
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

// https://w3c.github.io/webdriver/#forward
func (api WebDriverApi) Forward() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/forward", api.SessionId), nil)
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

// https://w3c.github.io/webdriver/#refresh
func (api WebDriverApi) Refresh() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/refresh", api.SessionId), nil)
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

// https://w3c.github.io/webdriver/#get-title
func (api WebDriverApi) GetTitle() (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/title", api.SessionId))
	if err != nil {
		return "", err
	}

	// Manage error
	responseError := ErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return "", fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := GetTitleResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Value, nil
}
