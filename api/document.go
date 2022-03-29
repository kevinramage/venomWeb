package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// https://w3c.github.io/webdriver/#get-page-source
func (api WebDriverApi) GetPageSource() (string, error) {

	// Send request
	path := fmt.Sprintf("session/%s/source", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get page source request")
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get page source: ", err)
		return "", fmt.Errorf("impossible to get page source")
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

// https://w3c.github.io/webdriver/#executing-script
func (api WebDriverApi) ExecuteScript(script string, args []string) error {

	// Create request
	type executeScriptRequest struct {
		Script string   `json:"script"`
		Args   []string `json:"args"`
	}
	request := executeScriptRequest{Script: script, Args: args}

	// Send request
	path := fmt.Sprintf("session/%s/execute/sync", api.SessionId)
	resp, err := ProceedPostRequest(api, path, request)
	if err != nil {
		log.Error("An error occured during execute script request")
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to execute script: ", err)
		return fmt.Errorf("impossible to execute script: %s", script)
	}

	return nil
}

// https://w3c.github.io/webdriver/#execute-async-script
func (api WebDriverApi) ExecuteAsyncScript(script string, args []string) error {

	// Create request
	type executeScriptRequest struct {
		Script string   `json:"script"`
		Args   []string `json:"args"`
	}
	request := executeScriptRequest{Script: script, Args: args}

	// Send request
	path := fmt.Sprintf("session/%s/execute/async", api.SessionId)
	resp, err := ProceedPostRequest(api, path, request)
	if err != nil {
		log.Error("An error occured during execute script request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to execute script: ", err)
		return fmt.Errorf("impossible to execute script: %s", script)
	}

	return nil
}
