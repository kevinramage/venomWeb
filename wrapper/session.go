package venomWeb

import (
	"fmt"
	"io/ioutil"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
)

type Session struct {
	api api.WebDriverApi
}

/*
----------------------------------------
  Session
----------------------------------------
*/
func (s Session) DeleteSession() error {
	return s.api.DeleteSession()
}

func (s Session) Reset() error {

	// Close alert
	s.api.AcceptAlert()

	// Delete cookies
	if err := s.api.DeleteAllCookies(); err != nil {
		return err
	}

	// Clean url
	if err := s.api.Navigate("about:blank"); err != nil {
		return err
	}

	return nil
}

/*
----------------------------------------
  Timeouts
----------------------------------------
*/
func (s Session) GetTimeouts() (common.Timeouts, error) {
	return s.api.GetSessionTimeout()
}

func (s Session) SetTimeouts(timeouts common.Timeouts) error {
	return s.api.SetSessionTimeout(timeouts)
}

/*
----------------------------------------
  Url
----------------------------------------
*/

func (s Session) Navigate(url string) error {
	return s.api.Navigate(url)
}

func (s Session) GetURL() (string, error) {
	return s.api.GetCurrentUrl()
}

func (s Session) Forward() error {
	return s.api.Forward()
}

func (s Session) Back() error {
	return s.api.Back()
}

func (s Session) Refresh() error {
	return s.api.Refresh()
}

func (s Session) GetTitle() (string, error) {
	return s.api.GetTitle()
}

/*
----------------------------------------
  Window / Context
----------------------------------------
*/

func (s Session) GetWindow() (Window, error) {
	handle, err := s.api.GetWindowHandle()
	if err == nil {
		window := Window{handleId: handle}
		return window, nil
	} else {
		return Window{}, err
	}
}

func (s Session) CloseWindow() error {
	return s.api.CloseWindow()
}

func (s Session) GetWindows() ([]Window, error) {
	handles, err := s.api.GetWindowHandles()
	windows := []Window{}
	if err == nil {
		for i := 0; i < len(handles); i++ {
			windows = append(windows, Window{handleId: handles[i], api: s.api})
		}
		return windows, nil
	} else {
		return []Window{}, err
	}
}

func (s Session) NextWindow() error {
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
			}
		}
	}
	return err
}

func (s Session) NewWindow(windowType string) (Window, error) {
	handle, err := s.api.NewWindows(windowType)
	if err == nil {
		return Window{handleId: handle, api: s.api}, nil
	} else {
		return Window{}, err
	}
}

func (s Session) SwitchToFrame(element Element) error {
	return s.api.SwitchToFrame(element.elementId)
}

func (s Session) SwitchToIndexFrame(index int) error {
	return s.api.SwitchToIndexFrame(index)
}

func (s Session) SwitchToParentFrame() error {
	return s.api.SwitchToParentFrame()
}

func (s Session) GetSize() (common.Rect, error) {
	return s.api.GetWindowRect()
}

func (s Session) Size(width int, height int) error {
	return s.api.SetWindowRect(width, height)
}

func (s Session) Maximize() error {
	return s.api.Maximize()
}

func (s Session) Minimize() error {
	return s.api.Minimize()
}

func (s Session) Fullscreen() error {
	return s.api.Fullscreen()
}

/*
----------------------------------------
  Element
----------------------------------------
*/

func (s Session) FindElement(selector string, locatorStrategy string) (Element, error) {
	fmt.Printf("%s", s.api.SessionId)
	eltId, err := s.api.FindElement(selector, locatorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: s.api}, nil
	} else {
		return Element{}, nil
	}
}

func (s Session) FindElements(selector string, locatorStrategy string) ([]Element, error) {
	eltsId, err := s.api.FindElements(selector, locatorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltsId); i++ {
			elements = append(elements, Element{elementId: eltsId[i], api: s.api})
		}
		return elements, nil
	} else {
		return []Element{}, nil
	}
}

func (s Session) GetActiveElement() (Element, error) {
	eltId, err := s.api.GetActiveElement()
	if err == nil {
		return Element{elementId: eltId, api: s.api}, nil
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
	return s.api.GetPageSource()
}

func (s Session) ExecuteScript(script string, args []string) error {
	return s.api.ExecuteScript(script, args)
}

func (s Session) ExecuteAsyncScript(script string, args []string) error {
	return s.api.ExecuteAsyncScript(script, args)
}

/*
----------------------------------------
  Cookie
----------------------------------------
*/

func (s Session) GetAllCookies() ([]string, error) {
	return s.api.GetAllCookies()
}

func (s Session) GetNamedCookie(cookieName string) (string, error) {
	return s.api.GetNamedCookie(cookieName)
}

func (s Session) AddCookie(cookie common.Cookie) error {
	return s.api.AddCookie(cookie)
}

func (s Session) DeleteCookie(cookieName string) error {
	return s.api.DeleteCookie(cookieName)
}

func (s Session) DeleteAllCookies() error {
	return s.api.DeleteAllCookies()
}

/*
----------------------------------------
  User prompts
----------------------------------------
*/

func (s Session) DismissAlert() error {
	return s.api.DismissAlert()
}

func (s Session) AcceptAlert() error {
	return s.api.AcceptAlert()
}

func (s Session) GetAlertText() (string, error) {
	return s.api.GetAlertText()
}

func (s Session) SendAlertText(alertText string) error {
	return s.api.SendAlertText(alertText)
}

/*
----------------------------------------
  Screen capture
----------------------------------------
*/

func (s Session) TakeScreenshot(fileName string) error {
	content, err := s.api.TakeScreenShot()
	if err == nil {
		return ioutil.WriteFile(fileName, content, 0666)
	} else {
		return err
	}
}

func (s Session) String() string {
	return "page"
}
