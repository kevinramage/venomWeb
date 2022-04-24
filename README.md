# Venom Web
Venom web is a package dedicated to venom integration test.
It allow to manipulate web driver to test workflow of your web application.

# Overview
Venom web is a go library to interact with web driver.
It use REST api to communicate with web driver.

# Web Drivers
Use web driver to instanciate a session with the browser.
To configure the browser behaviour, you can use:
* args: Browser arguments
Chrome: https://peter.sh/experiments/chromium-command-line-switches/
Firefox: https://wiki.mozilla.org/Firefox/CommandLineOptions

* prefs: Browser preferences
* shortcuts: Common browser features

** Proxy: Force the proxy of the web driver (default value: "")
```go
webDriver := venomWeb.ChromeDriver([]string{})
webDriver.Proxy = "localhost:8888"
page, _ := webDriver.NewSession()
page.Navigate("https://github.com/")
webDriver.Stop()
```

** Headless: Enable browser headless mode, usefull for CI/CD integration (default value: false)
```go
webDriver := venomWeb.ChromeDriver([]string{})
webDriver.Headless = true
page, _ := webDriver.NewSession()
page.Navigate("https://github.com/")
webDriver.Stop()
```

** LogLevel: Define log level of the component, values possible: DEBUG, INFO, WARN, ERROR (default value: WARN)
The DEBUG mode provide REST communication between the client and the web driver
```go
webDriver := venomWeb.ChromeDriver([]string{})
webDriver.LogLevel = "DEBUG"
page, _ := webDriver.NewSession()
page.Navigate("https://github.com/")
webDriver.Stop()
```

** Detach: Enable browser detach mode, usefull for debug (default value: false)
```go
webDriver := venomWeb.ChromeDriver([]string{})
webDriver.Detach = true
page, _ := webDriver.NewSession()
page.Navigate("https://github.com/")
webDriver.Stop()
```

## Chrome

### Installation
To use chrome browser, the web driver must be installed (chrome browser and driver must use the same version).
https://chromedriver.chromium.org/downloads

### Usage
Call venomWeb.ChromeDriver method to instanciate a chrome driver.

## Firefox

### Installation
To use firefox browser, the web driver must be installed (firefox browser and driver must use the same version).
https://github.com/mozilla/geckodriver/releases

Call venomWeb.GeckoDriver method to instanciate a chrome driver

## Microsoft Edge
Call venomWeb.EdgeChroniumDriver method to instanciate a chrome driver


# Session

# Element

# Window

# Links

https://w3c.github.io/webdriver
https://github.com/jlipps/simple-wd-spec/ 
https://www.guru99.com/chrome-options-desiredcapabilities.html
https://www.selenium.dev/documentation/legacy/json_wire_protocol/#webelement

https://source.chromium.org/chromium/chromium/src/+/master:chrome/test/chromedriver/window_commands.cc;l=865;drc=7fb345a0da63049b102e1c0bcdc8d7831110e324
https://webdriver.io/docs/api/webdriver/#performactions
https://github.com/jlipps/simple-wd-spec#perform-actions

## Tests

Alert:
https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_alert

Sync:
https://www.w3schools.com/w3css/w3css_progressbar.asp