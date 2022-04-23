package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
)

type CreateSessionResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     struct {
		SessionId                string `json:"sessionId"`
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
	} `json:"value"`
}

type GetTimeoutResponse struct {
	SessionId string          `json:"sessionId"`
	Status    int             `json:"status"`
	Value     common.Timeouts `json:"value"`
}

// https://w3c.github.io/webdriver/#new-session
func (api *WebDriverApi) CreateSession(browserName string, args []string, prefs map[string]interface{}, detach bool) (CreateSessionResponse, error) {

	var resp interface{}
	var err error

	// Chrome
	if browserName == "chrome" {

		requestBody := common.ChromeWebDriverSession{}
		requestBody.Capabilities.AlwaysMatch.AcceptInsecureCerts = true
		requestBody.Capabilities.AlwaysMatch.BrowserName = browserName
		requestBody.Capabilities.AlwaysMatch.ChromeOptions.Args = args
		requestBody.Capabilities.AlwaysMatch.ChromeOptions.Prefs = prefs
		if detach {
			requestBody.Capabilities.AlwaysMatch.ChromeOptions.Detach = detach
		}
		resp, err = ProceedPostRequest(*api, "session", requestBody)

		// Firefox
	} else if browserName == "firefox" {

		requestBody := common.GeckoWebDriverSession{}
		requestBody.Capabilities.AlwaysMatch.AcceptInsecureCerts = true
		requestBody.Capabilities.AlwaysMatch.BrowserName = browserName
		requestBody.Capabilities.AlwaysMatch.FirefoxOptions.Args = args
		requestBody.Capabilities.AlwaysMatch.FirefoxOptions.Prefs = prefs
		resp, err = ProceedPostRequest(*api, "session", requestBody)

	} else if browserName == "msedge" {

		requestBody := common.EdgeWebDriverSession{}
		requestBody.Capabilities.AlwaysMatch.AcceptInsecureCerts = true
		requestBody.Capabilities.AlwaysMatch.BrowserName = browserName
		requestBody.Capabilities.AlwaysMatch.EdgeOptions.Args = args
		requestBody.Capabilities.AlwaysMatch.EdgeOptions.Prefs = prefs
		resp, err = ProceedPostRequest(*api, "session", requestBody)
	}

	// Send request
	if err != nil {
		return CreateSessionResponse{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return CreateSessionResponse{}, fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := CreateSessionResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return CreateSessionResponse{}, err
	}
	api.SessionId = responseBody.SessionId
	if api.SessionId == "" {
		api.SessionId = responseBody.Value.SessionId
	}

	// Display debug informations
	/*
		version := ""
		if responseBody.Value.BrowserName == "chrome" {
			version = responseBody.Value.Chrome.ChromeDriverVersion
		}
		log.Info(fmt.Sprintf("Session created %s version: %s - Load strategy: %s\n", responseBody.Value.BrowserName, version, responseBody.Value.PageLoadStrategy))
	*/

	return responseBody, nil
}

// https://w3c.github.io/webdriver/#delete-session
func (api WebDriverApi) DeleteSession() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	path := fmt.Sprintf("session/%s", api.SessionId)
	resp, err := ProceedDeleteRequest(api, path)

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return fmt.Errorf(responseError.Value.Message)
	}

	return err
}

// https://w3c.github.io/webdriver/#status
func (api WebDriverApi) CheckStatus() (common.DriverStatus, error) {

	// Send request
	resp, err := ProceedGetRequest(api, "status")
	if err != nil {
		return common.DriverStatus{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return common.DriverStatus{}, fmt.Errorf(responseError.Value.Message)
	}

	// Manage response
	responseBody := common.DriverStatus{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return common.DriverStatus{}, err
	}

	return responseBody, nil
}
