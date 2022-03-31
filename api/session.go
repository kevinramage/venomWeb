package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// Firefox
// Args   https://wiki.mozilla.org/Firefox/CommandLineOptions
// Prefs  https://searchfox.org/mozilla-central/source/modules/libpref/init/all.js

// Chrome
// Args https://peter.sh/experiments/chromium-command-line-switches/

type CreateSessionRequest struct {
	Capabilities struct {
		AlwaysMatch struct {
			BrowserName         string `json:"browserName,omitempty"`
			AcceptInsecureCerts bool   `json:"acceptInsecureCerts,omitempty"`
			FirefoxOptions      struct {
				Binary  string                 `json:"binary,omitempty"`
				Profile string                 `json:"profile,omitempty"`
				Args    []string               `json:"args,omitempty"`
				Prefs   map[string]interface{} `json:"prefs,omitempty"`
				Log     struct {
					Level string `json:"level,omitempty"`
				} `json:"log,omitempty"`
				Env struct {
					Log  string `json:"MOZ_LOG,omitempty"`
					File string `json:"MOZ_LOG_FILE,omitempty"`
				} `json:"env,omitempty"`
				AndroidPackage         string   `json:"androidPackage,omitempty"`
				AndroidActivity        string   `json:"androidActivity,omitempty"`
				AndroidDeviceSerial    string   `json:"androidDeviceSerial,omitempty"`
				AndroidIntentArguments []string `json:"androidIntentArguments,omitempty"`
			} `json:"moz:firefoxOptions,omitempty"`

			ChromeOptions struct {
				Args             []string               `json:"args,omitempty"`
				Binary           string                 `json:"binary,omitempty"`
				Extensions       []string               `json:"extensions,omitempty"`
				LocalState       map[string]interface{} `json:"localState,omitempty"`
				Prefs            map[string]interface{} `json:"prefs,omitempty"`
				Detach           bool                   `json:"detach,omitempty"`
				DebuggerAddress  string                 `json:"debuggerAddress,omitempty"`
				ExcludeSwitches  []string               `json:"excludeSwitches,omitempty"`
				MinidumpPath     string                 `json:"minidumpPath,omitempty"`
				MobileEmulation  map[string]interface{} `json:"mobileEmulation,omitempty"`
				PerfLoggingPrefs map[string]interface{} `json:"perfLoggingPrefs,omitempty"`
				WindowTypes      []string               `json:"windowTypes,omitempty"`
			} `json:"goog:chromeOptions,omitempty"`
		} `json:"alwaysMatch,omitempty"`
	} `json:"capabilities,omitempty"`
	//DesiredCapabilities common.ChromeCapability `json:"desiredCapabilities"`
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

	// Create request body
	requestBody := CreateSessionRequest{}
	requestBody.Capabilities.AlwaysMatch.AcceptInsecureCerts = true
	requestBody.Capabilities.AlwaysMatch.BrowserName = browserName

	if browserName == "firefox" {
		requestBody.Capabilities.AlwaysMatch.FirefoxOptions.Args = args
		requestBody.Capabilities.AlwaysMatch.FirefoxOptions.Prefs = prefs

	} else {

		requestBody.Capabilities.AlwaysMatch.ChromeOptions.Args = args
		requestBody.Capabilities.AlwaysMatch.ChromeOptions.Prefs = prefs
		if detach {
			requestBody.Capabilities.AlwaysMatch.ChromeOptions.Detach = detach
		}
	}

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
	if api.SessionId == "" {
		api.SessionId = responseBody.Value.SessionId
	}

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
