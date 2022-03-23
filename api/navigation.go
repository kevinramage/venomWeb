package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type GetTitleResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     struct {
		Title string `json:"title"`
	} `json:"value"`
}

// https://w3c.github.io/webdriver/#navigate-to
func (api WebDriverApi) Navigate(url string) error {

	// Create request body
	request := NavigateRequest{
		Url: url,
	}

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/url", api.SessionId), request)
	if err != nil {
		log.Error("An error during the navigation: ", err)
	}
	return err
}

// https://w3c.github.io/webdriver/#get-current-url
func (api WebDriverApi) GetCurrentUrl() (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/url", api.SessionId))
	if err != nil {
		log.Error("An error during url retrieve call: ", err)
		return "", err
	}

	// Manage response
	responseBody := NavigateResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#back
func (api WebDriverApi) Back() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/back", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during click request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#forward
func (api WebDriverApi) Forward() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/forward", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during forward request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#refresh
func (api WebDriverApi) Refresh() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/refresh", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during refresh request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#get-title
func (api WebDriverApi) GetTitle() (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/title", api.SessionId))
	if err != nil {
		log.Error("An error occured during get title request: ", err)
		return "", err
	}

	// Manage response
	responseBody := GetTitleResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", err
	}

	return responseBody.Value.Title, nil
}
