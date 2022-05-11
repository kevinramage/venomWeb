package common

const (
	DEBUG string = "DEBUG"
	INFO  string = "INFO"
	WARN  string = "WARN"
	ERROR string = "ERROR"
)

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WebDriverOptions struct {
	Debug           bool
	Command         string
	CommandLineArgs []string
	CommandPort     string
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

// Firefox
// Args   https://wiki.mozilla.org/Firefox/CommandLineOptions
// Prefs  https://searchfox.org/mozilla-central/source/modules/libpref/init/all.js

// Chrome
// Args https://peter.sh/experiments/chromium-command-line-switches/

// Edge
// https://docs.microsoft.com/fr-fr/microsoft-edge/webdriver-chromium/capabilities-edge-options

type ChromeWebDriverSession struct {
	Capabilities struct {
		AlwaysMatch struct {
			BrowserName         string        `json:"browserName,omitempty"`
			AcceptInsecureCerts bool          `json:"acceptInsecureCerts,omitempty"`
			ChromeOptions       ChromeOptions `json:"goog:chromeOptions,omitempty"`
		} `json:"alwaysMatch,omitempty"`
	} `json:"capabilities,omitempty"`
}

type GeckoWebDriverSession struct {
	Capabilities struct {
		AlwaysMatch struct {
			BrowserName         string         `json:"browserName,omitempty"`
			AcceptInsecureCerts bool           `json:"acceptInsecureCerts,omitempty"`
			FirefoxOptions      FirefoxOptions `json:"moz:firefoxOptions,omitempty"`
		} `json:"alwaysMatch,omitempty"`
	} `json:"capabilities,omitempty"`
}

type EdgeWebDriverSession struct {
	Capabilities struct {
		AlwaysMatch struct {
			BrowserName         string      `json:"browserName,omitempty"`
			AcceptInsecureCerts bool        `json:"acceptInsecureCerts,omitempty"`
			EdgeOptions         EdgeOptions `json:"ms:edgeOptions,omitempty"`
		} `json:"alwaysMatch,omitempty"`
	} `json:"capabilities,omitempty"`
}

type ChromeOptions struct {
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
}

type FirefoxOptions struct {
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
}

type EdgeOptions struct {
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
	WdpAddress       string                 `json:"wdpAddress,omitempty"`
	WdpPassword      string                 `json:"wdpPassword,omitempty"`
	WdpUsername      string                 `json:"wdpUsername,omitempty"`
	WebviewOptions   map[string]interface{} `json:"webviewOptions,omitempty"`
	WindowsApp       string                 `json:"windowsApp,omitempty"`
}

type Action struct {
	Id         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"` // "key", "pointer", "wheel", or "none"
	Parameters struct {
		PointerType string `json:"pointerType,omitempty"` // "mouse", "pen", or "touch"
	} `json:"parameters,omitempty"`
	Actions []SubAction `json:"actions,omitempty"`
}

type SubAction struct {
	Type     string `json:"type,omitempty"` // pause, pointerDown, pointerUp, pointerMove, pointerCancel
	Duration int    `json:"duration,omitempty"`
	X        int    `json:"x,omitempty"`
	Y        int    `json:"y,omitempty"`
	Origin   string `json:"origin,omitempty"`
	Button   int    `json:"button"`
}
