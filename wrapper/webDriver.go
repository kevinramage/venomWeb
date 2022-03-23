package venomWeb

import (
	"strings"
	"time"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
	"github.com/kevinramage/venomWeb/service"
	log "github.com/sirupsen/logrus"
)

type WebDriver struct {
	Driver  common.WebDriverOptions
	Service service.WebDriverService
	Api     api.WebDriverApi
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
	DefineLogLevel(w.Driver.LogLevel)
	log.Info("WewDriver.Start")
	w.Service = service.New()
	err := w.Service.Start(w.Driver.Command, w.Driver.LogLevel)
	if err != nil {
		return err
	}
	err = w.Service.Wait(w.Driver.Timeout, w.Driver.Url)
	return err
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

///TODO New page called before start
///TODO Manage error
func (w *WebDriver) NewPage() (Page, error) {
	log.Info("WewDriver.NewPage")
	err := w.Api.CreateSession(w.Driver.Args)
	return Page{w.Api}, err
}

func NewWebDriver(webDriver *WebDriver) WebDriver {

	webDriver.Driver.LogLevel = "WARN"
	webDriver.Driver.Url = "http://localhost:9515"
	webDriver.Driver.Timeout = time.Second * 60
	webDriver.Driver.Command = webDriver.Driver.WebDriverBinary
	webDriver.Api = api.New(webDriver.Driver.Url)
	return *webDriver
}

func ChromeDriver(args []string) WebDriver {
	webDriver := WebDriver{}
	webDriver.Driver.Args = args
	webDriver.Driver.WebDriverBinary = "chromedriver.exe"
	return NewWebDriver(&webDriver)
}

func GeckoDriver() WebDriver {
	return WebDriver{}
}
