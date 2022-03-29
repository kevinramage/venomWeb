package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type CreateSessionRequest struct {
	DesiredCapabilities common.ChromeCapability `json:"desiredCapabilities"`
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
	Value     struct {
		AcceptInsecureCerts      bool   `json:"acceptInsecureCerts,omitempty"`
		AcceptSslCerts           bool   `json:"acceptSslCerts"`
		BrowserConnectionEnabled bool   `json:"browserConnectionEnabled"`
		BrowserName              string `json:"browserName"`
		Chrome                   struct {
			ChromeDriverVersion string `json:"chromedriverVersion"`
			UserDataDir         string `json:"userDataDir"`
		} `json:"chrome"`
		CssSelectorsEnabled bool `json:"cssSelectorsEnabled"`
		DatabaseEnabled     bool `json:"databaseEnabled"`
		ChromeOptions       struct {
			DebuggerAddress string `json:"debuggerAddress"`
		} `json:"goog:chromeOptions"`
		HandlesAlerts             bool            `json:"handlesAlerts"`
		HasTouchScreen            bool            `json:"hasTouchScreen"`
		JavascriptEnabled         bool            `json:"javascriptEnabled"`
		LocationContextEnabled    bool            `json:"locationContextEnabled"`
		MobileEmulationEnabled    bool            `json:"mobileEmulationEnabled"`
		NativeEvents              bool            `json:"nativeEvents"`
		NetworkConnectionEnabled  bool            `json:"networkConnectionEnabled"`
		PageLoadStrategy          string          `json:"pageLoadStrategy,omitempty"`
		Proxy                     common.Proxy    `json:"proxy,omitempty"`
		Platform                  string          `json:"platform"`
		Rotatable                 bool            `json:"rotatable"`
		SetWindowRect             bool            `json:"setWindowRect,omitempty"`
		StrictFileInteractability bool            `json:"strictFileInteractability,omitempty"`
		TakesHeapSnapshot         bool            `json:"takesHeapSnapshot"`
		TakesScreenshot           bool            `json:"takesScreenshot"`
		Timeouts                  common.Timeouts `json:"timeouts,omitempty"`
		UnhandledPromptBehavior   string          `json:"unhandledPromptBehavior,omitempty"`
		Version                   string          `json:"version"`
		WebStorageEnabled         bool            `json:"webStorageEnabled"`
		CredBlob                  bool            `json:"webauthn:extension:credBlob"`
		LargeBlob                 bool            `json:"webauthn:extension:largeBlob"`
		VirtualAuthenticators     bool            `json:"webauthn:virtualAuthenticators"`
	}
}

type GetTimeoutResponse struct {
	SessionId string          `json:"sessionId"`
	Status    int             `json:"status"`
	Value     common.Timeouts `json:"value"`
}

// https://w3c.github.io/webdriver/#new-session
func (api *WebDriverApi) CreateSession(args []string) (CreateSessionResponse, error) {

	// Create request body
	requestBody := CreateSessionRequest{}
	//requestBody.DesiredCapabilities.AcceptInsecureCerts = true
	requestBody.DesiredCapabilities.ChromeOptions.Args = args

	// Send request
	resp, err := ProceedPostRequest(*api, "session", requestBody)
	if err != nil {
		log.Error("An error during the session creation: ", err)
		return CreateSessionResponse{}, err
	}

	// Manage response
	responseBody := CreateSessionResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return CreateSessionResponse{}, err
	}
	api.SessionId = responseBody.SessionId

	// Display debug informations
	version := ""
	if responseBody.Value.BrowserName == "chrome" {
		version = responseBody.Value.Chrome.ChromeDriverVersion
	}
	log.Info(fmt.Sprintf("Session created %s version: %s - Load strategy: %s\n", responseBody.Value.BrowserName, version, responseBody.Value.PageLoadStrategy))

	return responseBody, nil
}

// https://w3c.github.io/webdriver/#delete-session
func (api WebDriverApi) DeleteSession() error {

	// Send request
	path := fmt.Sprintf("session/%s", api.SessionId)
	_, err := ProceedDeleteRequest(api, path)

	return err
}

// https://w3c.github.io/webdriver/#status
func (api WebDriverApi) CheckStatus() (common.DriverStatus, error) {

	// Send request
	resp, err := ProceedGetRequest(api, "status")
	if err != nil {
		log.Error("An error during the session creation: ", err)
		return common.DriverStatus{}, err
	}

	responseBody := common.DriverStatus{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return common.DriverStatus{}, err
	}

	return responseBody, nil
}
