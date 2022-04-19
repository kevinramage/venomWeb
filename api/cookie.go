package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
)

type GetAllCookiesResponse struct {
	SessionId string   `json:"sessionId"`
	Status    int      `json:"status"`
	Value     []string `json:"value"`
}

// https://w3c.github.io/webdriver/#get-all-cookies
func (api WebDriverApi) GetAllCookies() ([]string, error) {

	// Security
	if api.SessionId == "" {
		return []string{}, fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedGetRequest(api, path)
	if err != nil {
		return []string{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return []string{}, fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := GetAllCookiesResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return []string{}, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-named-cookie
func (api WebDriverApi) GetNamedCookie(cookieName string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/cookie/%s", api.SessionId, cookieName)
	resp, err := ProceedGetRequest(api, path)
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

// https://w3c.github.io/webdriver/#add-cookie
func (api WebDriverApi) AddCookie(cookie common.Cookie) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Prepare body
	type AddCookieRequest struct {
		Cookie common.Cookie `json:"Cookie"`
	}
	addCookieRequest := AddCookieRequest{Cookie: cookie}

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedPostRequest(api, path, addCookieRequest)
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

// https://w3c.github.io/webdriver/#delete-cookie
func (api WebDriverApi) DeleteCookie(cookieName string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/cookie/%s", api.SessionId, cookieName)
	resp, err := ProceedDeleteRequest(api, path)
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

// https://w3c.github.io/webdriver/#delete-all-cookies
func (api WebDriverApi) DeleteAllCookies() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, err := ProceedDeleteRequest(api, path)
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
