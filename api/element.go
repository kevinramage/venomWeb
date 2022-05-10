package api

import (
	"fmt"
	"reflect"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
)

// https://chromium.googlesource.com/chromium/src/+/master/docs/chromedriver_status.md
type ElementRequest struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

type ElementResponse struct {
	SessionId string            `json:"sessionId"`
	Status    int               `json:"status"`
	Value     map[string]string `json:"value"`
}

type ElementErrorResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		StackTrace string `json:"stacktrace"`
	} `json:"value"`
}

type ElementsResponse struct {
	SessionId string              `json:"sessionId"`
	Status    int                 `json:"status"`
	Value     []map[string]string `json:"value"`
}

type SendKeysRequest struct {
	Text string `json:"text"`
}

type StringResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     string `json:"value"`
}

type StringListResponse struct {
	SessionId string   `json:"sessionId"`
	Status    int      `json:"status"`
	Value     []string `json:"value"`
}

type BooleanResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     bool   `json:"value"`
}

// https://w3c.github.io/webdriver/#find-element
func (api WebDriverApi) FindElement(selector string, selectorType string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/element", api.SessionId), request)

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
	responseBody := ElementResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#find-elements
func (api WebDriverApi) FindElements(selector string, selectorType string) ([]string, error) {

	// Security
	if api.SessionId == "" {
		return []string{}, fmt.Errorf("invalid session id")
	}

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/elements", api.SessionId), request)

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

// https://w3c.github.io/webdriver/#find-element-from-element
func (api WebDriverApi) FindElementFromElement(elementId string, selector string, selectorType string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/element", api.SessionId, elementId), request)

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
	responseBody := ElementResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#find-elements-from-element
func (api WebDriverApi) FindElementsFromElement(elementId string, selector string, selectorType string) ([]string, error) {

	// Security
	if api.SessionId == "" {
		return []string{}, fmt.Errorf("invalid session id")
	}

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/elements", api.SessionId, elementId), request)

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

// https://w3c.github.io/webdriver/#find-element-from-shadow-root
func (api WebDriverApi) FindElementFromShadow(shadowId string, selector string, selectorType string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/shadow/%s/element", api.SessionId, shadowId), request)

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
	responseBody := ElementResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#find-elements-from-shadow-root
func (api WebDriverApi) FindElementsFromShadow(shadowId string, selector string, selectorType string) ([]string, error) {

	// Security
	if api.SessionId == "" {
		return []string{}, fmt.Errorf("invalid session id")
	}

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/shadow/%s/elements", api.SessionId, shadowId), request)

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

// https://w3c.github.io/webdriver/#get-active-element
func (api WebDriverApi) GetActiveElement() (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/active", api.SessionId))

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
	responseBody := ElementResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#get-element-shadow-root
func (api WebDriverApi) GetElementShadowRoot(elementId string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/shadow", api.SessionId, elementId))

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

// https://w3c.github.io/webdriver/#is-element-selected
func (api WebDriverApi) IsElementSelected(elementId string) (bool, error) {

	// Security
	if api.SessionId == "" {
		return false, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/selected", api.SessionId, elementId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return false, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return false, errResp
	}

	// Manage response
	responseBody := BooleanResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return false, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-attribute
func (api WebDriverApi) GetElementAttribute(elementId string, name string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/attribute/%s", api.SessionId, elementId, name))

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

// https://w3c.github.io/webdriver/#get-element-property
func (api WebDriverApi) GetElementProperty(elementId string, name string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/property/%s", api.SessionId, elementId, name))

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

// https://w3c.github.io/webdriver/#get-element-css-value
func (api WebDriverApi) GetElementCSSValue(elementId string, name string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/css/%s", api.SessionId, elementId, name))

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

// https://w3c.github.io/webdriver/#get-element-text
func (api WebDriverApi) GetElementText(elementId string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/text", api.SessionId, elementId))

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

// https://w3c.github.io/webdriver/#get-element-tag-name
func (api WebDriverApi) GetElementTagName(elementId string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/name", api.SessionId, elementId))

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

// https://w3c.github.io/webdriver/#get-element-rect
func (api WebDriverApi) GetElementRect(elementId string) (common.Rect, error) {

	// Security
	if api.SessionId == "" {
		return common.Rect{}, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/rect", api.SessionId, elementId))

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

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#is-element-enabled
func (api WebDriverApi) IsElementEnabled(elementId string) (bool, error) {

	// Security
	if api.SessionId == "" {
		return false, fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/enabled", api.SessionId, elementId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return false, fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	if errResp != nil {
		return false, errResp
	}

	// Manage response
	responseBody := BooleanResponse{}
	err := mapstructure.Decode(resp, &responseBody)
	if err != nil {
		return false, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-computed-role
func (api WebDriverApi) GetComputedRole(elementId string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/computedrole", api.SessionId, elementId))

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

// https://w3c.github.io/webdriver/#get-computed-label
func (api WebDriverApi) GetComputedLabel(elementId string) (string, error) {

	// Security
	if api.SessionId == "" {
		return "", fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/computedlabel", api.SessionId, elementId))

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

// https://w3c.github.io/webdriver/#element-click
func (api WebDriverApi) Click(elementId string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}
	if elementId == "" {
		return fmt.Errorf("invalid element id")
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/click", api.SessionId, elementId), nil)

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

// https://w3c.github.io/webdriver/#element-clear
func (api WebDriverApi) Clear(elementId string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/clear", api.SessionId, elementId), nil)

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

// https://w3c.github.io/webdriver/#element-send-keys
func (api WebDriverApi) SendKeys(elementId string, text string) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Create request body
	request := SendKeysRequest{
		Text: text,
	}

	// Send request
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/value", api.SessionId, elementId), request)

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
