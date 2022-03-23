package venomWeb

import (
	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
)

type Page struct {
	Api api.WebDriverApi
}

func (p Page) Navigate(url string) error {
	return p.Api.Navigate(url)
}

func (p Page) Reset() {
}

func (p Page) GetURL() (string, error) {
	return p.Api.GetCurrentUrl()
}

func (p Page) GetSize() (common.Rect, error) {
	return p.Api.GetWindowRect()
}

func (p Page) Size(width int, height int) error {
	return p.Api.SetWindowRect(width, height)
}

func (p Page) Maximize() error {
	return p.Api.Maximize()
}

func (p Page) Minimize() error {
	return p.Api.Minimize()
}

func (p Page) Fullscreen() error {
	return p.Api.Fullscreen()
}

func (p Page) FindElement(selector string, locatorStrategy string) (string, error) {
	return p.Api.FindElement(selector, locatorStrategy)
}

func (p Page) FindElements(selector string, locatorStrategy string) ([]string, error) {
	return p.Api.FindElements(selector, locatorStrategy)
}

///TODO
func (p Page) Screenshot(fileName string) error {
	return nil
}

// https://w3c.github.io/webdriver/#get-title
func (p Page) GetTitle() (string, error) {
	return p.Api.GetTitle()
}

func (p Page) Title(title string) error {
	return nil
}

func (p Page) GetHTML() (string, error) {
	return "", nil
}

func (p Page) GetPopupText() (string, error) {
	return "", nil
}

func (p Page) EnterPopupText(text string) error {
	return nil
}

func (p Page) ConfirmPopupText() error {
	return nil
}

func (p Page) CancelPopupText() error {
	return nil
}

// https://w3c.github.io/webdriver/#forward
func (p Page) Forward() error {
	return p.Api.Forward()
}

func (p Page) Back() error {
	return p.Api.Back()
}

func (p Page) Refresh() error {
	return p.Api.Refresh()
}

func (p Page) SwitchToParentFrame() error {
	return nil
}

func (p Page) SwitchToRootFrame() error {
	return nil
}

func (p Page) SwitchToWindow(name string) error {
	return nil
}

func (p Page) GetWindow() (string, error) {
	return p.Api.GetWindowHandle()
}

func (p Page) NextWindow() error {
	return nil
}

func (p Page) CloseWindow() error {
	return p.Api.CloseWindow()
}

// https://w3c.github.io/webdriver/#dfn-timeouts-configuration
func (p Page) GetTimeouts() (common.Timeouts, error) {
	return p.Api.GetSessionTimeout()
}

func (p Page) SetTimeouts(timeouts common.Timeouts) error {
	return p.Api.SetSessionTimeout(timeouts)
}

// https://w3c.github.io/webdriver/#get-page-source
func (p Page) HTML() (string, error) {
	return p.Api.GetPageSource()
}

func (p Page) String() string {
	return "page"
}
