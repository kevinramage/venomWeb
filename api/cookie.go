package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type GetAllCookiesResponse struct {
	SessionId string   `json:"sessionId"`
	Status    int      `json:"status"`
	Value     []string `json:"value"`
}

// https://w3c.github.io/webdriver/#get-all-cookies
func (api WebDriverApi) GetAllCookies() ([]string, error) {

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get all cookies request: ", err)
		return []string{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get all cookies: ", err)
		return []string{}, err
	}

	// Manage response
	responseBody := GetAllCookiesResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return []string{}, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-named-cookie
func (api WebDriverApi) GetNamedCookie(cookieName string) (string, error) {

	// Send request
	path := fmt.Sprintf("session/%s/cookie/%s", api.SessionId, cookieName)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		log.Error("An error occured during get named cookie request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get named cookie: ", err)
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

// https://w3c.github.io/webdriver/#add-cookie
func (api WebDriverApi) AddCookie(cookie common.Cookie) error {

	// Prepare body
	type AddCookieRequest struct {
		Cookie common.Cookie `json:"Cookie"`
	}
	addCookieRequest := AddCookieRequest{Cookie: cookie}

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedPostRequest(api, path, addCookieRequest)
	if err != nil {
		log.Error("An error occured during add cookie request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to add cookie: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#delete-cookie
func (api WebDriverApi) DeleteCookie(cookieName string) error {

	// Send request
	path := fmt.Sprintf("session/%s/cookie/%s", api.SessionId, cookieName)
	resp, err := ProceedDeleteRequest(api, path)
	if err != nil {
		log.Error("An error occured during delete cookie request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to delete cookie: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#delete-all-cookies
func (api WebDriverApi) DeleteAllCookies() error {

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedDeleteRequest(api, path)
	if err != nil {
		log.Error("An error occured during delete all cookies request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to delete all cookies: ", err)
		return err
	}

	return nil
}
