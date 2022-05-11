package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// https://w3c.github.io/webdriver/#get-page-source
func (api WebDriverApi) GetPageSource() (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/source", api.SessionId)
	resp, errResp := ProceedGetRequest(api, path)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return "", fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return "", errResp
	}

	// Manage response
	responseBody := StringResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#executing-script
func (api WebDriverApi) ExecuteScript(script string, args []string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request
	type executeScriptRequest struct {
		Script string   `json:"script"`
		Args   []string `json:"args"`
	}
	request := executeScriptRequest{Script: script, Args: args}

	// Send request
	path := fmt.Sprintf("session/%s/execute/sync", api.SessionId)
	resp, errResp := ProceedPostRequest(api, path, request)

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

// https://w3c.github.io/webdriver/#execute-async-script
func (api WebDriverApi) ExecuteAsyncScript(script string, args []string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request
	type executeScriptRequest struct {
		Script string   `json:"script"`
		Args   []string `json:"args"`
	}
	request := executeScriptRequest{Script: script, Args: args}

	// Send request
	path := fmt.Sprintf("session/%s/execute/async", api.SessionId)
	resp, errResp := ProceedPostRequest(api, path, request)

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
