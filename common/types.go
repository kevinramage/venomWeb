package common

import (
	"time"
)

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WebDriverOptions struct {
	Timeout         time.Duration
	Debug           bool
	Command         string
	CommandLineArgs []string
	Args            []string
	Prefs           map[string]interface{}
	WebDriverBinary string
	Url             string
	BrowserName     string

	//Internal WebDriverInternal
}

/*
type WebDriverInternal struct {
	Command   *exec.Cmd
	Client    http.Client
	SessionId string
	Version   WebDriverVersion
	Ready     bool
}*/

type WebDriverVersion struct {
	DriverVersion string
	Osname        string
	Osarch        string
	Osversion     string
}

type Element struct {
	Id      string
	Text    string
	TagName string
}

// https://w3c.github.io/webdriver/#dfn-timeouts-object
type Timeouts struct {
	Implicit int `json:"implicit"`
	PageLoad int `json:"pageLoad"`
	Script   int `json:"script"`
}

// https://w3c.github.io/webdriver/#dfn-status
type DriverStatus struct {
	Value struct {
		Build struct {
			Version string `json:"version"`
		} `json:"build"`
		Message string `json:"message"`
		Os      struct {
			Arch    string `json:"arch"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		Ready bool `json:"ready"`
	} `json:"value"`
}

// https://w3c.github.io/webdriver/#dfn-table-of-standard-capabilities
type Proxy struct {
	ProxyType          string   `json:"proxyType,omitempty"`
	ProxyAutoconfigUrl string   `json:"proxyAutoconfigUrl,omitempty"`
	FtpProxy           string   `json:"ftpProxy,omitempty"`
	HttpProxy          string   `json:"httpProxy,omitempty"`
	NoProxy            []string `json:"noProxy,omitempty"`
	SslProxy           string   `json:"sslProxy,omitempty"`
	SocksProxy         string   `json:"socksProxy,omitempty"`
	SocksVersion       int      `json:"socksVersion,omitempty"`
}

// https://w3c.github.io/webdriver/#dfn-table-of-standard-capabilities
type Capability struct {
	AcceptInsecureCerts       bool      `json:"acceptInsecureCerts,omitempty"`
	PageLoadStrategy          string    `json:"pageLoadStrategy,omitempty"`
	Proxy                     *Proxy    `json:"proxy,omitempty"`
	SetWindowRect             bool      `json:"setWindowRect,omitempty"`
	Timeouts                  *Timeouts `json:"timeouts,omitempty"`
	StrictFileInteractability bool      `json:"strictFileInteractability,omitempty"`
	UnhandledPromptBehavior   string    `json:"unhandledPromptBehavior,omitempty"`
}

// https://w3c.github.io/webdriver/#dfn-table-of-location-strategies
const (
	CSS_SELECTOR             string = "css selector"
	LINKTEXT_SELECTOR        string = "link text"
	PARTIALLINKTEXT_SELECTOR string = "partial link text"
	TAGNAME_SELECTOR         string = "tag name"
	XPATH_SELECTOR           string = "xpath"
)

// https://w3c.github.io/webdriver/#dfn-table-for-cookie-conversion
type Cookie struct {
	Name           string `json:"Name"`
	Value          string `json:"Value,omitempty"`
	Path           string `json:"Path,omitempty"`
	Domain         string `json:"Domain,omitempty"`
	SecureOnlyFlag string `json:"SecureOnlyFlag,omitempty"`
	HttpOnlyFlag   string `json:"HttpOnlyFlag,omitempty"`
	ExpireTimeFlag string `json:"ExpireTimeFlag,omitempty"`
}
