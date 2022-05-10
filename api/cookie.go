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
	resp, errResp := ProceedGetRequest(api, path)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return []string{}, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return []string{}, errResp
	}

	// Manage response
	responseBody := GetAllCookiesResponse{}
	err := mapstructure.Decode(resp, &responseBody)
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
	resp, errResp := ProceedPostRequest(api, path, addCookieRequest)

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

// https://w3c.github.io/webdriver/#delete-cookie
func (api WebDriverApi) DeleteCookie(cookieName string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/cookie/%s", api.SessionId, cookieName)
	resp, errResp := ProceedDeleteRequest(api, path)

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

// https://w3c.github.io/webdriver/#delete-all-cookies
func (api WebDriverApi) DeleteAllCookies() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s/cookie", api.SessionId)
	resp, errResp := ProceedDeleteRequest(api, path)

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
