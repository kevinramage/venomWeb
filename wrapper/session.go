package venomWeb

import (
	"io/ioutil"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
)

type Session struct {
	Api api.WebDriverApi
}

/*
----------------------------------------
  Session
----------------------------------------
*/
func (s Session) DeleteSession() error {
	return s.Api.DeleteSession()
}

/*
----------------------------------------
  Timeouts
----------------------------------------
*/
func (s Session) GetTimeouts() (common.Timeouts, error) {
	return s.Api.GetSessionTimeout()
}

func (s Session) SetTimeouts(timeouts common.Timeouts) error {
	return s.Api.SetSessionTimeout(timeouts)
}

/*
----------------------------------------
  Url
----------------------------------------
*/

func (s Session) Navigate(url string) error {
	return s.Api.Navigate(url)
}

func (s Session) GetURL() (string, error) {
	return s.Api.GetCurrentUrl()
}

func (s Session) Forward() error {
	return s.Api.Forward()
}

func (s Session) Back() error {
	return s.Api.Back()
}

func (s Session) Refresh() error {
	return s.Api.Refresh()
}

func (s Session) GetTitle() (string, error) {
	return s.Api.GetTitle()
}

/*
----------------------------------------
  Window / Context
----------------------------------------
*/

func (s Session) GetWindow() (Window, error) {
	handle, err := s.Api.GetWindowHandle()
	if err == nil {
		window := Window{HandleId: handle}
		return window, nil
	} else {
		return Window{}, err
	}
}

func (s Session) CloseWindow() error {
	return s.Api.CloseWindow()
}

func (s Session) GetWindows() ([]Window, error) {
	handles, err := s.Api.GetWindowHandles()
	windows := []Window{}
	if err == nil {
		for i := 0; i < len(handles); i++ {
			windows = append(windows, Window{HandleId: handles[i]})
		}
		return windows, nil
	} else {
		return []Window{}, err
	}
}

func (s Session) NewWindow(windowType string) (Window, error) {
	handle, err := s.Api.NewWindows(windowType)
	if err == nil {
		return Window{HandleId: handle}, nil
	} else {
		return Window{}, err
	}
}

func (s Session) SwitchToFrame(frameId string) error {
	return s.Api.SwitchToFrame(frameId)
}

func (s Session) SwitchToParentFrame() error {
	return s.Api.SwitchToParentFrame()
}

func (s Session) GetSize() (common.Rect, error) {
	return s.Api.GetWindowRect()
}

func (s Session) Size(width int, height int) error {
	return s.Api.SetWindowRect(width, height)
}

func (s Session) Maximize() error {
	return s.Api.Maximize()
}

func (s Session) Minimize() error {
	return s.Api.Minimize()
}

func (s Session) Fullscreen() error {
	return s.Api.Fullscreen()
}

/*
----------------------------------------
  Element
----------------------------------------
*/

func (s Session) FindElement(selector string, locatorStrategy string) (Element, error) {
	eltId, err := s.Api.FindElement(selector, locatorStrategy)
	if err == nil {
		return Element{ElementId: eltId}, nil
	} else {
		return Element{}, nil
	}
}

func (s Session) FindElements(selector string, locatorStrategy string) ([]Element, error) {
	eltsId, err := s.Api.FindElements(selector, locatorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltsId); i++ {
			elements = append(elements, Element{ElementId: eltsId[i]})
		}
		return elements, nil
	} else {
		return []Element{}, nil
	}
}

func (s Session) GetActiveElement() (Element, error) {
	eltId, err := s.Api.GetActiveElement()
	if err == nil {
		return Element{ElementId: eltId}, nil
	} else {
		return Element{}, nil
	}
}

/*
----------------------------------------
  Document
----------------------------------------
*/

func (s Session) GetPageSource() (string, error) {
	return s.Api.GetPageSource()
}

func (s Session) ExecuteScript(script string, args []string) error {
	return s.Api.ExecuteScript(script, args)
}

func (s Session) ExecuteAsyncScript(script string, args []string) error {
	return s.Api.ExecuteAsyncScript(script, args)
}

/*
----------------------------------------
  Cookie
----------------------------------------
*/

func (s Session) GetAllCookies() ([]string, error) {
	return s.Api.GetAllCookies()
}

func (s Session) GetNamedCookie(cookieName string) (string, error) {
	return s.Api.GetNamedCookie(cookieName)
}

func (s Session) AddCookie(cookie common.Cookie) error {
	return s.Api.AddCookie(cookie)
}

func (s Session) DeleteCookie(cookieName string) error {
	return s.Api.DeleteCookie(cookieName)
}

func (s Session) DeleteAllCookies() error {
	return s.Api.DeleteAllCookies()
}

/*
----------------------------------------
  User prompts
----------------------------------------
*/

func (s Session) DismissAlert() error {
	return s.Api.DismissAlert()
}

func (s Session) AcceptAlert() error {
	return s.Api.AcceptAlert()
}

func (s Session) GetAlertText() (string, error) {
	return s.Api.GetAlertText()
}

func (s Session) SendAlertText(alertText string) error {
	return s.Api.SendAlertText(alertText)
}

/*
----------------------------------------
  Screen capture
----------------------------------------
*/

func (s Session) TakeScreenshot(fileName string) error {
	content, err := s.Api.TakeScreenShot()
	if err == nil {
		return ioutil.WriteFile(fileName, content, 0666)
	} else {
		return err
	}
}

func (s Session) String() string {
	return "page"
}
