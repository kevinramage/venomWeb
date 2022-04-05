package api

import (
	"fmt"
	"reflect"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
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

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/window", api.SessionId))
	if err != nil {
		log.Error("An error occured during get window handle request: ", err)
		return "", err
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#close-window
func (api WebDriverApi) CloseWindow() error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/window", api.SessionId))
	if err != nil {
		log.Error("An error occured during close window request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#switch-to-window
func (api WebDriverApi) SwitchWindow(handle string) error {

	// Create request body
	request := SwitchWindowRequest{
		Handle: handle,
	}

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/window", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during switch window request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-window-handles
func (api WebDriverApi) GetWindowHandles() ([]string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/window/handles", api.SessionId))
	if err != nil {
		log.Error("An error occured during get window handles request: ", err)
		return []string{}, err
	}

	// Manage response
	responseBody := ElementsResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
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

	// Create request
	type ReqStruct struct {
		Type string `json:"type"`
	}
	request := ReqStruct{Type: windowType}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/new", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during get window request: ", err)
		return "", err
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#switch-to-frame
// https://source.chromium.org/chromium/chromium/src/+/master:chrome/test/chromedriver/element_util.cc;l=309;drc=7fb345a0da63049b102e1c0bcdc8d7831110e324
// https://source.chromium.org/chromium/chromium/src/+/master:chrome/test/chromedriver/element_util.cc;drc=7fb345a0da63049b102e1c0bcdc8d7831110e324;l=31
func (api WebDriverApi) SwitchToFrame(id string) error {

	// Create request
	type SwitchToFrameRequest struct {
		Id struct {
			Element string `json:"element-6066-11e4-a52e-4f735466cecf"`
		} `json:"id"`
	}
	request := SwitchToFrameRequest{}
	request.Id.Element = id

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/frame", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during switch to frame request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#switch-to-frame
func (api WebDriverApi) SwitchToIndexFrame(index int) error {

	// Create request
	type SwitchToFrameRequest struct {
		Id int `json:"id"`
	}
	request := SwitchToFrameRequest{Id: index}

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/frame", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during switch to frame request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#switch-to-parent-frame
func (api WebDriverApi) SwitchToParentFrame() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/frame/parent", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during switch to parent frame request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-window-rect
func (api WebDriverApi) GetWindowRect() (common.Rect, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/window/rect", api.SessionId))
	if err != nil {
		log.Error("An error occured during get size request: ", err)
		return common.Rect{}, err
	}

	// Manage response
	responseBody := RectResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
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

	// Create request body
	request := SetWindowRectRequest{
		Width:  width,
		Height: height,
	}

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/rect", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during set size request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#maximize-window
func (api WebDriverApi) Maximize() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/maximize", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during maximize request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#minimize-window
func (api WebDriverApi) Minimize() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/minimize", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during minimize request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#fullscreen-window
func (api WebDriverApi) Fullscreen() error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/window/fullscreen", api.SessionId), nil)
	if err != nil {
		log.Error("An error occured during fullscreen request: ", err)
		return err
	} else {
		return nil
	}
}
