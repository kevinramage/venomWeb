package api

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type GetAllCookiesResponse struct {
	SessionId string   `json:"sessionId"`
	Status    int      `json:"status"`
	Value     []string `json:"value"`
}
type GetCookieNameResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

// https://w3c.github.io/webdriver/#get-all-cookies
func (api WebDriverApi) GetAllCookies() ([]string, error) {

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get all cookies request")
		return []string{}, nil
	}

	// Manage response
	responseBody := GetAllCookiesResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return []string{}, nil
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-named-cookie
func (api WebDriverApi) GetNamedCookie(cookieName string) (string, error) {

	// Send request
	path := fmt.Sprintf("session/%s/cookie/%s", api.SessionId, cookieName)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get named cookie request")
		return "", nil
	}

	// Manage response
	responseBody := GetCookieNameResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", nil
	}

	return responseBody.Value, nil
}
