package api

import (
	"net/http"
)

// https://chromium.googlesource.com/chromium/src/+/refs/heads/main/chrome/test/chromedriver/
// https://chromium.googlesource.com/chromium/src/+/master/docs/chromedriver_status.md
// https://chromium.googlesource.com/chromium/src/+/refs/heads/main/chrome/test/chromedriver/client/chromedriver.py
// https://chromedriver.chromium.org/capabilities

type WebDriverApi struct {
	SessionId string
	Url       string
	Client    http.Client
}

func New(url string) WebDriverApi {
	api := WebDriverApi{Url: url}
	api.Client = http.Client{}
	return api
}
