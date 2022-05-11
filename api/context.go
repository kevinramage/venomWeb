package api

import (
	"fmt"
	"reflect"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
)

type SwitchWindowRequest struct {
	Handle string `json:"handle"`
}

type SetWindowRectRequest struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type RectResponse struct {
	SessionId string      `json:"sessionId"`
	Status    int         `json:"status"`
	Value     common.Rect `json:"value"`
}

// https://w3c.github.io/webdriver/#get-window-handle
func (api WebDriverApi) GetWindowHandle() (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/window", api.SessionId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return "", fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return "", errResp
	}

	// Manage response
	responseBody := StringResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#close-window
func (api WebDriverApi) CloseWindow() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/window", api.SessionId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	return errResp
}

// https://w3c.github.io/webdriver/#switch-to-window
func (api WebDriverApi) SwitchWindow(handle string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request body
	request := SwitchWindowRequest{
		Handle: handle,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/window", api.SessionId), request)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#get-window-handles
func (api WebDriverApi) GetWindowHandles() ([]string, error) {

	// Security
	if api.SessionId == "" {
		return []string{}, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/window/handles", api.SessionId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return []string{}, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return []string{}, errResp
	}

	// Manage response
	responseBody := ElementsResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return []string{}, err
	}

	ids := []string{}
	for i := 0; i < len(responseBody.Value); i++ {
		keys := reflect.ValueOf(responseBody.Value[i]).MapKeys()
		key := keys[0].String()
		ids = append(ids, responseBody.Value[i][key])
	}

	return ids, nil
}

// https://w3c.github.io/webdriver/#new-window
// WindowType: tab or window
func (api WebDriverApi) NewWindows(windowType string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Create request
	type ReqStruct struct {
		Type string `json:"type"`
	}
	request := ReqStruct{Type: windowType}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/new", api.SessionId), request)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return "", fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return "", errResp
	}

	// Manage response
	responseBody := StringResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#switch-to-frame
// https://source.chromium.org/chromium/chromium/src/+/master:chrome/test/chromedriver/element_util.cc;l=309;drc=7fb345a0da63049b102e1c0bcdc8d7831110e324
// https://source.chromium.org/chromium/chromium/src/+/master:chrome/test/chromedriver/element_util.cc;drc=7fb345a0da63049b102e1c0bcdc8d7831110e324;l=31
func (api WebDriverApi) SwitchToFrame(id string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request
	type SwitchToFrameRequest struct {
		Id struct {
			Element string `json:"element-6066-11e4-a52e-4f735466cecf"`
		} `json:"id"`
	}
	request := SwitchToFrameRequest{}
	request.Id.Element = id

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/frame", api.SessionId), request)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#switch-to-frame
func (api WebDriverApi) SwitchToIndexFrame(index int) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request
	type SwitchToFrameRequest struct {
		Id int `json:"id"`
	}
	request := SwitchToFrameRequest{Id: index}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/frame", api.SessionId), request)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#switch-to-parent-frame
func (api WebDriverApi) SwitchToParentFrame() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/frame/parent", api.SessionId), nil)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#get-window-rect
func (api WebDriverApi) GetWindowRect() (common.Rect, error) {

	// Security
	if api.SessionId == "" {
		return common.Rect{}, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/window/rect", api.SessionId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return common.Rect{}, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return common.Rect{}, errResp
	}

	// Manage response
	responseBody := RectResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return common.Rect{}, err
	}

	rect := common.Rect{
		X:      responseBody.Value.X,
		Y:      responseBody.Value.Y,
		Width:  responseBody.Value.Width,
		Height: responseBody.Value.Height,
	}

	return rect, nil
}

// https://w3c.github.io/webdriver/#set-window-rect
func (api WebDriverApi) SetWindowRect(width int, height int) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request body
	request := SetWindowRectRequest{
		Width:  width,
		Height: height,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/rect", api.SessionId), request)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#maximize-window
func (api WebDriverApi) Maximize() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/maximize", api.SessionId), nil)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#minimize-window
func (api WebDriverApi) Minimize() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/minimize", api.SessionId), nil)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#fullscreen-window
func (api WebDriverApi) Fullscreen() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/fullscreen", api.SessionId), nil)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}
