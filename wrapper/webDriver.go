package venomWeb

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
	"github.com/kevinramage/venomWeb/service"
	log "github.com/sirupsen/logrus"
)

type WebDriver struct {
	Driver    common.WebDriverOptions
	Service   service.WebDriverService
	Api       api.WebDriverApi
	isStarted bool
	Headless  bool
	Proxy     string
	Detach    bool
	LogLevel  string
}

func DefineLogLevel(logLevel string) {
	logLevelUpperCase := strings.ToUpper(logLevel)

	if logLevelUpperCase == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else if logLevelUpperCase == "INFO" {
		log.SetLevel(log.InfoLevel)
	} else if logLevelUpperCase == "WARN" {
		log.SetLevel(log.WarnLevel)
	} else if logLevelUpperCase == "ERROR" {
		log.SetLevel(log.ErrorLevel)
	}
}

///TODO What happen if the web driver start two times
///TODO Manage error
///TODO SetTimeout
func (w *WebDriver) Start() error {
	DefineLogLevel(w.LogLevel)
	log.Info("WebDriver.Start")
	if !w.isStarted {
		w.Service = service.New()
		err := w.Service.Start(w.Driver.Command, w.LogLevel, w.Driver.CommandLineArgs)
		if err != nil {
			return err
		}
		w.isStarted = true
		err = w.Service.Wait(w.Driver.Timeout, w.Driver.Url)
		return err

	} else {
		return errors.New("a web driver is already running")
	}
}

///TODO Stop when start not called
///TODO Close session when NewPage not called
///TODO Manage error
func (w *WebDriver) Stop() error {
	log.Info("WewDriver.Stop")
	err := w.Api.DeleteSession()
	if err != nil {
		return err
	}
	err = w.Service.Stop()
	return err
}

func (w *WebDriver) Status() (common.DriverStatus, error) {
	log.Info("WewDriver.Status")
	return w.Api.CheckStatus()
}

///TODO New page called before start
///TODO Manage error
func (w *WebDriver) NewSession() (Session, error) {
	log.Info("WewDriver.NewSession")

	// Headless & detach & proxy
	w.defineHeadless()
	w.defineProxy()

	// Create session
	_, err := w.Api.CreateSession(w.Driver.BrowserName, w.Driver.Args, w.Driver.Prefs, w.Detach)
	return Session{w.Api}, err
}

func (w *WebDriver) defineHeadless() {
	if w.Headless {
		if w.Driver.BrowserName == "firefox" {
			if !common.SliceContains(w.Driver.Args, "-headless") {
				w.Driver.Args = append(w.Driver.Args, "-headless")
			}
		} else {
			if !common.SliceContains(w.Driver.Args, "headless") {
				w.Driver.Args = append(w.Driver.Args, "headless")
			}
		}
	}
}

func (w *WebDriver) defineProxy() {
	if w.Proxy != "" {
		if w.Driver.BrowserName == "firefox" {
			proxy_values := strings.Split(w.Proxy, ":")
			if len(proxy_values) == 2 {
				port, err := strconv.Atoi(proxy_values[1])
				if err == nil {
					w.Driver.Prefs["network.proxy.type"] = 1
					w.Driver.Prefs["network.websocket.allowInsecureFromHTTPS"] = true
					w.Driver.Prefs["network.proxy.http"] = proxy_values[0]
					w.Driver.Prefs["network.proxy.http_port"] = port
					w.Driver.Prefs["network.proxy.ssl"] = proxy_values[0]
					w.Driver.Prefs["network.proxy.ssl_port"] = port
				}
			}
		} else {
			if !common.SliceContains(w.Driver.Args, "ignore-certificate-errors") {
				w.Driver.Args = append(w.Driver.Args, "ignore-certificate-errors")
			}
			if !common.SliceContains(w.Driver.Args, "ignore-ssl-errors") {
				w.Driver.Args = append(w.Driver.Args, "ignore-ssl-errors")
			}
			w.Driver.Args = append(w.Driver.Args, "-proxy-server="+w.Proxy)
		}
	}
}

func NewWebDriver(webDriver *WebDriver) WebDriver {

	webDriver.Driver.Prefs = make(map[string]interface{})
	webDriver.LogLevel = "WARN"
	webDriver.Driver.Timeout = time.Second * 60
	webDriver.Driver.Command = webDriver.Driver.WebDriverBinary
	webDriver.Api = api.New(webDriver.Driver.Url)
	return *webDriver
}

func ChromeDriver(args []string) WebDriver {
	webDriver := WebDriver{}
	webDriver.Driver.BrowserName = "chrome"
	webDriver.Driver.Args = args
	if runtime.GOOS == "windows" {
		webDriver.Driver.WebDriverBinary = "chromedriver.exe"
	} else {
		webDriver.Driver.WebDriverBinary = "chromedriver"
	}
	webDriver.Driver.CommandLineArgs = []string{"--log=WARN", "--port=9515"}
	webDriver.Driver.Url = "http://localhost:9515"
	return NewWebDriver(&webDriver)
}

func GeckoDriver(args []string) WebDriver {
	webDriver := WebDriver{}
	webDriver.Driver.BrowserName = "firefox"
	webDriver.Driver.Args = args
	if runtime.GOOS == "windows" {
		webDriver.Driver.WebDriverBinary = "geckodriver.exe"
	} else {
		webDriver.Driver.WebDriverBinary = "geckodriver"
	}
	webDriver.Driver.CommandLineArgs = []string{"--log=WARN", "--port=4444"}
	webDriver.Driver.Url = "http://localhost:4444"
	return NewWebDriver(&webDriver)
}

func EdgeChroniumDriver(args []string) WebDriver {
	webDriver := WebDriver{}
	webDriver.Driver.BrowserName = "msedge"
	webDriver.Driver.Args = args
	if runtime.GOOS == "windows" {
		webDriver.Driver.WebDriverBinary = "msedgedriver.exe"
	} else {
		webDriver.Driver.WebDriverBinary = "msedgedriver"
	}
	webDriver.Driver.CommandLineArgs = []string{"--log=WARN", "--port=9515"}
	webDriver.Driver.Url = "http://localhost:9515"
	return NewWebDriver(&webDriver)
}

func OperaDriver(args []string) WebDriver {
	webDriver := WebDriver{}
	webDriver.Driver.BrowserName = "opera"
	webDriver.Driver.Args = args
	if runtime.GOOS == "windows" {
		webDriver.Driver.WebDriverBinary = "operadriver.exe"
	} else {
		webDriver.Driver.WebDriverBinary = "operadriver"
	}
	webDriver.Driver.CommandLineArgs = []string{"--log=WARN", "--port=9515"}
	webDriver.Driver.Url = "http://localhost:9515"
	return NewWebDriver(&webDriver)
}
