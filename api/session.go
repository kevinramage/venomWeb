package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type CreateSessionRequest struct {
	DesiredCapabilities DesiredCapabilitiesStruct `json:"desiredCapabilities"`
}
type DesiredCapabilitiesStruct struct {
	AcceptSslCerts bool                `json:"acceptSslCerts"`
	ChromeOptions  ChromeOptionsStruct `json:"chromeOptions"`
}
type ChromeOptionsStruct struct {
	Args  []string    `json:"args"`
	Prefs PrefsStruct `json:"prefs"`
}
type PrefsStruct struct {
}

type CreateSessionResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
}

type GetTimeoutResponse struct {
	SessionId string          `json:"sessionId"`
	Status    int             `json:"status"`
	Value     common.Timeouts `json:"value"`
}

// https://w3c.github.io/webdriver/#new-session
func (api *WebDriverApi) CreateSession(args []string) error {

	// Create request body
	requestBody := CreateSessionRequest{
		DesiredCapabilities: DesiredCapabilitiesStruct{
			AcceptSslCerts: true,
			ChromeOptions:  ChromeOptionsStruct{},
		},
	}
	if len(args) > 0 {
		requestBody.DesiredCapabilities.ChromeOptions.Args = args
	} else {
		requestBody.DesiredCapabilities.ChromeOptions.Args = []string{}
	}

	// Send request
	resp, err := ProceedPostRequest(*api, "session", requestBody)
	if err != nil {
		log.Error("An error during the session creation: ", err)
		return err
	}

	// Manage response
	responseBody := CreateSessionResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return err
	}
	api.SessionId = responseBody.SessionId

	return nil
}

// https://w3c.github.io/webdriver/#delete-session
func (api WebDriverApi) DeleteSession() error {

	// Send request
	path := fmt.Sprintf("session/%s", api.SessionId)
	_, err := ProceedDeleteRequest(api, path)

	return err
}

// https://w3c.github.io/webdriver/#status
func (api WebDriverApi) CheckStatus() error {

	// Send request
	_, err := ProceedGetRequest(api, "status")
	return err
}
