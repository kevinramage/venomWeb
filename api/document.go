package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type GetSourceResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

// https://w3c.github.io/webdriver/#get-page-source
func (api WebDriverApi) GetPageSource() (string, error) {

	// Send request
	path := fmt.Sprintf("session/%s/source", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get page source request")
		return "", nil
	}

	// Manage response
	responseBody := GetSourceResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
	}

	return responseBody.Value, err
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
	_, err := ProceedPostRequest(api, path, request)
	if err != nil {
		log.Error("An error occured during execute script request")
		return nil
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
	_, err := ProceedPostRequest(api, path, request)
	if err != nil {
		log.Error("An error occured during execute script request")
		return nil
	}

	return nil
}
