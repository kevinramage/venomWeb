package venomWeb

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Session struct {
	api api.WebDriverApi
}

/*
----------------------------------------
  Session
----------------------------------------
*/

// Delete session method allow to destroy web driver session
// Return nil if operation proceed with success, return an error else
func (s Session) DeleteSession() error {
	log.Info("Session.DeleteSession")
	err := s.api.DeleteSession()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during delete session action")
		log.Error(err)
	}
	return err
}

// Reset current session allow to reset web driver state
// Return nil if operation proceed with sucess, return an error else
func (s Session) Reset() error {
	log.Info("Session.Reset")

	// Close alert
	s.api.AcceptAlert()

	// Delete cookies
	if err := s.api.DeleteAllCookies(); err != nil {
		err = errors.Wrapf(err, "an error occured during reset action")
		log.Error(err)
		return err
	}

	// Clean url
	if err := s.api.Navigate("about:blank"); err != nil {
		err = errors.Wrapf(err, "an error occured during reset action")
		log.Error(err)
		return err
	}

	return nil
}

/*
----------------------------------------
  Timeouts
----------------------------------------
*/

// GetTimeouts method allow to get timeout values
// Return timeout object if operation succeed, return an error else
func (s Session) GetTimeouts() (common.Timeouts, error) {
	log.Info("Session.GetTimeouts")
	timeouts, err := s.api.GetSessionTimeout()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get timeouts action")
		log.Error(err)
	}
	return timeouts, err
}

// SetTimeouts method allow to set timeout value
// Return nil if operation succeed, return an error else
func (s Session) SetTimeouts(timeouts common.Timeouts) error {
	log.Info("Session.SetTimeouts")
	err := s.api.SetSessionTimeout(timeouts)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during set timeouts action")
		log.Error(err)
	}
	return err
}

/*
----------------------------------------
  Url
----------------------------------------
*/

// Navigate method allow to change the current web driver url
// Return nil if operation succeed, return an error else
func (s Session) Navigate(url string) error {
	log.Info("Session.Navigate")
	err := s.api.Navigate(url)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during navigate action")
		log.Error(err)
	}
	return err
}

// GetURL method allow to get the current web driver URL
// Return URL if operation succeed, return an error else
func (s Session) GetURL() (string, error) {
	log.Info("Session.GetURL")
	url, err := s.api.GetCurrentUrl()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get url action")
		log.Error(err)
	}
	return url, err
}

func (s Session) Forward() error {
	log.Info("Session.Forward")
	err := s.api.Forward()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during forward action")
		log.Error(err)
	}
	return err
}

func (s Session) Back() error {
	log.Info("Session.Back")
	err := s.api.Back()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during back action")
		log.Error(err)
	}
	return err
}

func (s Session) Refresh() error {
	log.Info("Session.Refresh")
	err := s.api.Refresh()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during refresh action")
		log.Error(err)
	}
	return err
}

func (s Session) GetTitle() (string, error) {
	log.Info("Session.GetTitle")

	title, err := s.api.GetTitle()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get title action")
		log.Error(err)
	}
	return title, err
}

/*
----------------------------------------
  Window / Context
----------------------------------------
*/

func (s Session) GetWindow() (Window, error) {
	log.Info("Session.GetWindow")
	handle, err := s.api.GetWindowHandle()
	if err == nil {
		window := Window{handleId: handle}
		return window, nil
	} else {
		err = errors.Wrapf(err, "an error occured during get window action")
		log.Error(err)
		return Window{}, err
	}
}

func (s Session) CloseWindow() error {
	log.Info("Session.CloseWindow")
	err := s.api.CloseWindow()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during close window action")
		log.Error(err)
	}
	return err
}

func (s Session) GetWindows() ([]Window, error) {
	log.Info("Session.GetWindows")
	handles, err := s.api.GetWindowHandles()
	windows := []Window{}
	if err == nil {
		for i := 0; i < len(handles); i++ {
			windows = append(windows, Window{handleId: handles[i], api: s.api})
		}
		return windows, nil
	} else {
		err = errors.Wrapf(err, "an error occured during get windows action")
		log.Error(err)
		return []Window{}, err
	}
}

func (s Session) NextWindow() error {
	log.Info("Session.NextWindow")
	handles, err := s.api.GetWindowHandles()
	if err == nil {
		handle, err := s.api.GetWindowHandle()
		if err == nil {
			var index = -1
			for i, val := range handles {
				if val == handle {
					index = i
				}
			}
			if index != -1 {
				newHandle := handles[index%len((handles))]
				s.api.SwitchWindow(newHandle)
			} else {
				err = fmt.Errorf("invalid index")
				err = errors.Wrapf(err, "an error occured during next window action")
				log.Error(err)
				return err
			}
		}
	}

	if err != nil {
		err = errors.Wrapf(err, "an error occured during next window action")
		log.Error(err)
	}

	return err
}

func (s Session) NewWindow(windowType string) (Window, error) {
	log.Info("Session.NewWindow")
	handle, err := s.api.NewWindows(windowType)
	if err == nil {
		return Window{handleId: handle, api: s.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during new window action")
		log.Error(err)
		return Window{}, err
	}
}

func (s Session) SwitchToFrame(element Element) error {
	log.Info("Session.SwitchToFrame")
	err := s.api.SwitchToFrame(element.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during switch to frame action")
		log.Error(err)
	}
	return err
}

func (s Session) SwitchToIndexFrame(index int) error {
	log.Info("Session.SwitchToIndexFrame")
	err := s.api.SwitchToIndexFrame(index)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during switch to index frame action")
		log.Error(err)
	}
	return err
}

func (s Session) SwitchToParentFrame() error {
	log.Info("Session.SwitchToParentFrame")
	err := s.api.SwitchToParentFrame()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during switch to parent frame action")
		log.Error(err)
	}
	return err
}

func (s Session) GetSize() (common.Rect, error) {
	log.Info("Session.GetSize")
	size, err := s.api.GetWindowRect()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get size action")
		log.Error(err)
	}
	return size, err
}

func (s Session) Size(width int, height int) error {
	log.Info("Session.Size")
	err := s.api.SetWindowRect(width, height)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during set size action")
		log.Error(err)
	}
	return err
}

func (s Session) Maximize() error {
	log.Info("Session.Maximize")
	err := s.api.Maximize()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during maximize action")
		log.Error(err)
	}
	return err
}

func (s Session) Minimize() error {
	log.Info("Session.Minimize")
	err := s.api.Minimize()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during minimize action")
		log.Error(err)
	}
	return err
}

func (s Session) Fullscreen() error {
	log.Info("Session.Fullscreen")
	err := s.api.Fullscreen()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during fullscreen action")
		log.Error(err)
	}
	return err
}

/*
----------------------------------------
  Element
----------------------------------------
*/

func (s Session) FindElement(selector string, locatorStrategy string) (Element, error) {
	log.Info("Session.FindElement")
	eltId, err := s.api.FindElement(selector, locatorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: s.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during find element action")
		log.Error(err)
		return Element{}, err
	}
}

func (s Session) FindElements(selector string, locatorStrategy string) ([]Element, error) {
	log.Info("Session.FindElements")
	eltsId, err := s.api.FindElements(selector, locatorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltsId); i++ {
			elements = append(elements, Element{elementId: eltsId[i], api: s.api})
		}
		return elements, nil
	} else {
		err = errors.Wrapf(err, "an error occured during find elements action")
		log.Error(err)
		return []Element{}, err
	}
}

func (s Session) GetActiveElement() (Element, error) {
	log.Info("Session.GetActiveElement")
	eltId, err := s.api.GetActiveElement()
	if err == nil {
		return Element{elementId: eltId, api: s.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during get active element action")
		log.Error(err)
		return Element{}, err
	}
}

// SyncElement allow to wait the web element creation
// Return nil if operation succeed, return an error else
func (s Session) SyncElement(selector string, locatorStrategy string, timeout int64) error {
	log.Info("Session.SyncElement")
	now := time.Now()
	_, err := s.api.FindElement(selector, locatorStrategy)
	for err != nil {
		time.Sleep(100 * time.Millisecond)
		_, err = s.api.FindElement(selector, locatorStrategy)
		duration := time.Until(now).Milliseconds() * -1
		if err != nil && duration >= timeout {
			return fmt.Errorf("impossible to synchronize element due to a timeout")
		}
	}
	return nil
}

// SyncElementAbsence allow to wait the web element deletion
// Return nil if operation succeed, return an error else
func (s Session) SyncElementAbsence(selector string, locatorStrategy string, timeout int64) error {
	log.Info("Session.SyncElementAbsence")
	now := time.Now()
	_, err := s.api.FindElement(selector, locatorStrategy)
	for err == nil {
		time.Sleep(250 * time.Millisecond)
		_, err = s.api.FindElement(selector, locatorStrategy)
		duration := time.Until(now).Milliseconds() * -1
		if duration >= timeout {
			return fmt.Errorf("impossible to synchronize element absence due to a timeout")
		}
	}
	return nil
}

// SyncElementText allow to wait text value of an element
// Return nil if operation succeed, return an error else
// WebSite to test this feature: https://www.w3schools.com/w3css/w3css_progressbar.asp
func (s Session) SyncElementText(selector string, locatorStrategy string, timeout int64, expectedText string) error {
	log.Info("Session.SyncElementText")
	now := time.Now()
	text := ""

	// Wait element presence
	s.SyncElement(selector, locatorStrategy, timeout)

	eltId, err := s.api.FindElement(selector, locatorStrategy)
	if err == nil {
		text, err = s.api.GetElementText(eltId)
	}
	for text != expectedText || err != nil {
		time.Sleep(250 * time.Millisecond)
		text, err = s.api.GetElementText(eltId)
		duration := time.Until(now).Milliseconds() * -1
		if duration >= int64(timeout) {
			return fmt.Errorf("impossible to synchronize element due to a timeout")
		}
	}
	return nil
}

// SyncElementCSSValue allow to wait a specific CSS value of an element
// Return nil if operation succeed, return an error else
func (s Session) SyncElementCSSValue(selector string, locatorStrategy string, timeout int64, CSSPropertyName string, expectedValue string) error {
	log.Info("Session.SyncElementCSSValue")
	now := time.Now()
	propertyValue := ""

	// Wait element presence
	s.SyncElement(selector, locatorStrategy, timeout)

	eltId, err := s.api.FindElement(selector, locatorStrategy)
	if err == nil {
		propertyValue, err = s.api.GetElementCSSValue(eltId, CSSPropertyName)
	}
	for propertyValue != expectedValue || err != nil {
		time.Sleep(250 * time.Millisecond)
		propertyValue, err = s.api.GetElementCSSValue(eltId, CSSPropertyName)
		duration := time.Until(now).Milliseconds() * -1
		if duration >= timeout {
			return fmt.Errorf("impossible to synchronize element due to a timeout")
		}
	}
	return nil
}

// SyncElementPropertyValue allow to wait a specific property value of an element
// Return nil if operation succeed, return an error else
func (s Session) SyncElementProperyValue(selector string, locatorStrategy string, timeout int64, CSSPropertyName string, expectedValue string) error {
	log.Info("Session.SyncElementProperyValue")
	now := time.Now()
	propertyValue := ""

	// Wait element presence
	s.SyncElement(selector, locatorStrategy, timeout)

	eltId, err := s.api.FindElement(selector, locatorStrategy)
	if err == nil {
		propertyValue, err = s.api.GetElementProperty(eltId, CSSPropertyName)
	}
	for propertyValue != expectedValue || err != nil {
		time.Sleep(250 * time.Millisecond)
		propertyValue, err = s.api.GetElementProperty(eltId, CSSPropertyName)
		duration := time.Until(now).Milliseconds() * -1
		if duration >= timeout {
			return fmt.Errorf("impossible to synchronize element due to a timeout")
		}
	}
	return nil
}

/*
----------------------------------------
  Document
----------------------------------------
*/

func (s Session) GetPageSource() (string, error) {
	log.Info("Session.GetPageSource")
	code, err := s.api.GetPageSource()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get page source action")
		log.Error(err)
	}
	return code, err
}

func (s Session) ExecuteScript(script string, args []string) error {
	log.Info("Session.ExecuteScript")
	err := s.api.ExecuteScript(script, args)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during execute script action")
		log.Error(err)
	}
	return err
}

func (s Session) ExecuteAsyncScript(script string, args []string) error {
	log.Info("Session.ExecuteAsyncScript")
	err := s.api.ExecuteAsyncScript(script, args)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during execute async script action")
		log.Error(err)
	}
	return err
}

/*
----------------------------------------
  Cookie
----------------------------------------
*/

func (s Session) GetAllCookies() ([]string, error) {
	log.Info("Session.GetAllCookies")
	cookies, err := s.api.GetAllCookies()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get all cookies action")
		log.Error(err)
	}
	return cookies, err
}

func (s Session) GetNamedCookie(cookieName string) (string, error) {
	log.Info("Session.GetNamedCookie")
	cookieName, err := s.api.GetNamedCookie(cookieName)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get named cookie action")
		log.Error(err)
	}
	return cookieName, err
}

func (s Session) AddCookie(cookie common.Cookie) error {
	log.Info("Session.AddCookie")
	err := s.api.AddCookie(cookie)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during add cookie action")
		log.Error(err)
	}
	return err
}

func (s Session) DeleteCookie(cookieName string) error {
	log.Info("Session.DeleteCookie")
	err := s.api.DeleteCookie(cookieName)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during delete cookie action")
		log.Error(err)
	}
	return err
}

func (s Session) DeleteAllCookies() error {
	log.Info("Session.DeleteAllCookies")
	err := s.api.DeleteAllCookies()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during delete all cookies action")
		log.Error(err)
	}
	return err
}

/*
----------------------------------------
  User prompts
----------------------------------------
*/

func (s Session) DismissAlert() error {
	log.Info("Session.DismissAlert")
	err := s.api.DismissAlert()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during dismiss alert action")
		log.Error(err)
	}
	return err
}

func (s Session) AcceptAlert() error {
	log.Info("Session.AcceptAlert")
	err := s.api.AcceptAlert()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during accept alert action")
		log.Error(err)
	}
	return err
}

func (s Session) GetAlertText() (string, error) {
	log.Info("Session.GetAlertText")
	alert, err := s.api.GetAlertText()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get alert text action")
		log.Error(err)
	}
	return alert, err
}

func (s Session) SendAlertText(alertText string) error {
	log.Info("Session.SendAlertText")
	err := s.api.SendAlertText(alertText)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during send alert text action")
		log.Error(err)
	}
	return err
}

/*
----------------------------------------
  Screen capture
----------------------------------------
*/

func (s Session) TakeScreenshot(fileName string) error {
	log.Info("Session.TakeScreenshot")
	content, err := s.api.TakeScreenShot()
	if err == nil {
		err = ioutil.WriteFile(fileName, content, 0666)
	}
	if err != nil {
		err = errors.Wrapf(err, "an error occured during take screenshot action")
		log.Error(err)
	}

	return err
}

func (s Session) String() string {
	return "page"
}
