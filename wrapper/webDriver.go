package venomWeb

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
	"github.com/kevinramage/venomWeb/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type WebDriver struct {
	driver    common.WebDriverOptions
	service   service.WebDriverService
	api       api.WebDriverApi
	isStarted bool
	Headless  bool
	Proxy     string
	Detach    bool
	LogLevel  string
	Timeout   time.Duration
}

func DefineLogLevel(logLevel string) {
	logLevelUpperCase := strings.ToUpper(logLevel)

	if logLevelUpperCase == common.DEBUG {
		log.SetLevel(log.DebugLevel)
	} else if logLevelUpperCase == common.INFO {
		log.SetLevel(log.InfoLevel)
	} else if logLevelUpperCase == common.WARN {
		log.SetLevel(log.WarnLevel)
	} else if logLevelUpperCase == common.ERROR {
		log.SetLevel(log.ErrorLevel)
	}
}

// Start web driver
func (w *WebDriver) Start() error {
	log.Info("WebDriver.Start")
	DefineLogLevel(w.LogLevel)
	if !w.isStarted {
		w.service = service.New()
		err := w.service.Start(w.driver.Command, w.driver.CommandPort, w.LogLevel, w.driver.CommandLineArgs)
		if err != nil {
			err = errors.Wrapf(err, "an error occured during session starting")
			log.Error(err)
			return err
		}
		w.isStarted = true
		err = w.service.Wait(w.Timeout, w.driver.Url)
		if err != nil {
			err = errors.Wrapf(err, "an error occured during session starting")
			log.Error(err)
		}
		return err

	} else {
		err := errors.New("a web driver is already running")
		err = errors.Wrapf(err, "an error occured during session starting")
		log.Error(err)
		return err
	}
}

// Stop service and delete session
func (w *WebDriver) Stop() error {
	log.Info("WebDriver.Stop")

	if w.api.SessionId != "" {
		err := w.api.DeleteSession()
		if err != nil {
			err = errors.Wrapf(err, "an error occured during delete session action")
			log.Error(err)
			return err
		}
	}
	err := w.service.Stop()
	w.isStarted = false
	if err != nil {
		err = errors.Wrapf(err, "an error occured during service stopping")
		log.Error(err)
	}
	return err
}

func (w *WebDriver) Status() (common.DriverStatus, error) {
	log.Info("WebDriver.Status")
	status, err := w.api.CheckStatus()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during check status action")
		log.Error(err)
	}
	return status, err
}

// Instanciate a new session to browser server
func (w *WebDriver) NewSession() (Session, error) {
	log.Info("WebDriver.NewSession")

	// Headless & detach & proxy
	w.defineHeadless()
	w.defineProxy()

	// Create session
	_, err := w.api.CreateSession(w.driver.BrowserName, w.driver.Binary, w.driver.Args, w.driver.Prefs, w.Detach)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during create session action")
		log.Error(err)
	}

	return Session{w.api}, err
}

func (w *WebDriver) defineHeadless() {
	log.Info("WebDriver.DefineHeadless")
	if w.Headless {
		if w.driver.BrowserName == "firefox" {
			if !common.SliceContains(w.driver.Args, "-headless") {
				w.driver.Args = append(w.driver.Args, "-headless")
			}
		} else {
			if !common.SliceContains(w.driver.Args, "headless") {
				w.driver.Args = append(w.driver.Args, "headless")
			}
		}
	}
}

func (w *WebDriver) defineProxy() {
	log.Info("WebDriver.DefineProxy")
	if w.Proxy != "" {
		if w.driver.BrowserName == "firefox" {
			proxy_values := strings.Split(w.Proxy, ":")
			if len(proxy_values) == 2 {
				port, err := strconv.Atoi(proxy_values[1])
				if err == nil {
					w.driver.Prefs["network.proxy.type"] = 1
					w.driver.Prefs["network.websocket.allowInsecureFromHTTPS"] = true
					w.driver.Prefs["network.proxy.http"] = proxy_values[0]
					w.driver.Prefs["network.proxy.http_port"] = port
					w.driver.Prefs["network.proxy.ssl"] = proxy_values[0]
					w.driver.Prefs["network.proxy.ssl_port"] = port
				}
			}
		} else {
			if !common.SliceContains(w.driver.Args, "ignore-certificate-errors") {
				w.driver.Args = append(w.driver.Args, "ignore-certificate-errors")
			}
			if !common.SliceContains(w.driver.Args, "ignore-ssl-errors") {
				w.driver.Args = append(w.driver.Args, "ignore-ssl-errors")
			}
			w.driver.Args = append(w.driver.Args, "-proxy-server="+w.Proxy)
		}
	}
}

func NewWebDriver(webDriver *WebDriver) WebDriver {

	webDriver.driver.Prefs = make(map[string]interface{})
	webDriver.LogLevel = "WARN"
	webDriver.Timeout = time.Second * 60
	webDriver.driver.Command = webDriver.driver.WebDriverBinary
	webDriver.api = api.New(webDriver.driver.Url)
	return *webDriver
}

func ChromeDriver(browserBinary string, driverBinary string, args []string, prefs map[string]interface{}, port string) WebDriver {
	log.Info("ChromeDriver.New")
	webDriver := WebDriver{}
	webDriver.driver.BrowserName = "chrome"
	webDriver.driver.Binary = browserBinary
	webDriver.driver.Args = args
	webDriver.driver.Prefs = prefs
	if driverBinary == "" {
		if runtime.GOOS == "windows" {
			webDriver.driver.WebDriverBinary = ".\\chromedriver.exe"
		} else {
			webDriver.driver.WebDriverBinary = "./chromedriver"
		}
	} else {
		webDriver.driver.WebDriverBinary = driverBinary
	}
	if port == "" {
		port = "9515"
	}
	webDriver.driver.CommandPort = port
	webDriver.driver.CommandLineArgs = []string{"--log=WARN", "--port=" + port}
	webDriver.driver.Url = "http://localhost:" + port
	return NewWebDriver(&webDriver)
}

func GeckoDriver(browserBinary string, driverBinary string, args []string, prefs map[string]interface{}, port string) WebDriver {
	log.Info("GeckoDriver.New")
	webDriver := WebDriver{}
	webDriver.driver.BrowserName = "firefox"
	webDriver.driver.Binary = browserBinary
	webDriver.driver.Args = args
	webDriver.driver.Prefs = prefs
	if driverBinary == "" {
		if runtime.GOOS == "windows" {
			webDriver.driver.WebDriverBinary = ".\\geckodriver.exe"
		} else {
			webDriver.driver.WebDriverBinary = "./geckodriver"
		}
	} else {
		webDriver.driver.WebDriverBinary = driverBinary
	}
	if port == "" {
		port = "4444"
	}
	webDriver.driver.CommandPort = port
	webDriver.driver.CommandLineArgs = []string{"--log=WARN", "--port=" + port}
	webDriver.driver.Url = "http://localhost:" + port
	return NewWebDriver(&webDriver)
}

func EdgeChroniumDriver(browserBinary string, driverBinary string, args []string, prefs map[string]interface{}, port string) WebDriver {
	log.Info("EdgeChroniumDriver.New")
	webDriver := WebDriver{}
	webDriver.driver.BrowserName = "msedge"
	webDriver.driver.Binary = browserBinary
	webDriver.driver.Args = args
	webDriver.driver.Prefs = prefs
	if driverBinary == "" {
		if runtime.GOOS == "windows" {
			webDriver.driver.WebDriverBinary = ".\\msedgedriver.exe"
		} else {
			webDriver.driver.WebDriverBinary = "./msedgedriver"
		}
	} else {
		webDriver.driver.WebDriverBinary = driverBinary
	}
	if port == "" {
		port = "9515"
	}
	webDriver.driver.CommandPort = port
	webDriver.driver.CommandLineArgs = []string{"--log=WARN", "--port=" + port}
	webDriver.driver.Url = "http://localhost:" + port
	return NewWebDriver(&webDriver)
}

func OperaDriver(browserBinary string, driverBinary string, args []string, port string) WebDriver {
	log.Info("OperaDriver.New")
	webDriver := WebDriver{}
	webDriver.driver.BrowserName = "opera"
	webDriver.driver.Binary = browserBinary
	webDriver.driver.Args = args
	if driverBinary == "" {
		if runtime.GOOS == "windows" {
			webDriver.driver.WebDriverBinary = ".\\operadriver.exe"
		} else {
			webDriver.driver.WebDriverBinary = "./operadriver"
		}
	} else {
		webDriver.driver.WebDriverBinary = driverBinary
	}
	if port == "" {
		port = "9515"
	}
	webDriver.driver.CommandPort = port
	webDriver.driver.CommandLineArgs = []string{"--log=WARN", "--port=" + port}
	webDriver.driver.Url = "http://localhost:" + port
	return NewWebDriver(&webDriver)
}

func BraveDriver(browserBinary string, driverBinary string, args []string, prefs map[string]interface{}, port string) WebDriver {
	log.Info("Brave.New")
	webDriver := WebDriver{}
	webDriver.driver.BrowserName = "brave"
	webDriver.driver.Binary = browserBinary
	webDriver.driver.Args = args
	webDriver.driver.Prefs = prefs
	if driverBinary == "" {
		if runtime.GOOS == "windows" {
			webDriver.driver.WebDriverBinary = ".\\chromedriver.exe"
		} else {
			webDriver.driver.WebDriverBinary = "./chromedriver"
		}
	} else {
		webDriver.driver.WebDriverBinary = driverBinary
	}
	if port == "" {
		port = "9515"
	}
	webDriver.driver.CommandPort = port
	webDriver.driver.CommandLineArgs = []string{"--log=WARN", "--port=" + port}
	webDriver.driver.Url = "http://localhost:" + port
	return NewWebDriver(&webDriver)
}
