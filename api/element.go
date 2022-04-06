package api

import (
	"fmt"
	"reflect"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
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
		Message string `json:"message"`
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

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during find element request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to find object: ", responseError.Value.Message)
		noFoundErr := fmt.Errorf("impossible to find object: %s (method %s)", selector, selectorType)
		return "", noFoundErr
	}

	// Manage response
	responseBody := ElementResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#find-elements
func (api WebDriverApi) FindElements(selector string, selectorType string) ([]string, error) {

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/elements", api.SessionId), request)
	if err != nil {
		log.Error("An error occured during find elements request: ", err)
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

// https://w3c.github.io/webdriver/#find-element-from-element
func (api WebDriverApi) FindElementFromElement(elementId string, selector string, selectorType string) (string, error) {

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/element", api.SessionId, elementId), request)
	if err != nil {
		log.Error("An error occured during find element from element request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to find object: ", responseError.Value.Message)
		noFoundErr := fmt.Errorf("impossible to find object: %s (method %s)", selector, selectorType)
		return "", noFoundErr
	}

	// Manage response
	responseBody := ElementResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#find-elements-from-element
func (api WebDriverApi) FindElementsFromElement(elementId string, selector string, selectorType string) ([]string, error) {

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/elements", api.SessionId, elementId), request)
	if err != nil {
		log.Error("An error occured during find elements from element request: ", err)
		return []string{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to find object: ", responseError.Value.Message)
		noFoundErr := fmt.Errorf("impossible to find object: %s (method %s)", selector, selectorType)
		return []string{}, noFoundErr
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

// https://w3c.github.io/webdriver/#find-element-from-shadow-root
func (api WebDriverApi) FindElementFromShadow(shadowId string, selector string, selectorType string) (string, error) {

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/shadow/%s/element", api.SessionId, shadowId), request)
	if err != nil {
		log.Error("An error occured during find element from shadow request: ", err)
		return "", err
	}

	// Manage response
	responseBody := ElementResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#find-elements-from-shadow-root
func (api WebDriverApi) FindElementsFromShadow(shadowId string, selector string, selectorType string) ([]string, error) {

	// Create request body
	request := ElementRequest{
		Using: selectorType,
		Value: selector,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/shadow/%s/elements", api.SessionId, shadowId), request)
	if err != nil {
		log.Error("An error occured during find elements from shadow request: ", err)
		return []string{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to find elements from shadow", responseError.Value.Message)
		errQuery := fmt.Errorf("impossible to find elements from shadow %s (selector: %s, shadow:%s", selector, selectorType, shadowId)
		return []string{}, errQuery
	}

	// Manage response
	responseBody := ElementsResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
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

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/active", api.SessionId))
	if err != nil {
		log.Error("An error occured during get active element request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get active element: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get active element")
	}

	// Manage response
	responseBody := ElementResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	// Get element id
	keys := reflect.ValueOf(responseBody.Value).MapKeys()
	key := keys[0].String()

	return responseBody.Value[key], nil
}

// https://w3c.github.io/webdriver/#get-element-shadow-root
func (api WebDriverApi) GetElementShadowRoot(elementId string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/shadow", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element shadow root request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element shadow root: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get element shadow root: %s", elementId)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#is-element-selected
func (api WebDriverApi) IsElementSelected(elementId string) (bool, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/selected", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during is element selected request: ", err)
		return false, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to determine if element selected: ", responseError.Value.Message)
		return false, fmt.Errorf("impossible to determine if element selected")
	}

	// Manage response
	responseBody := BooleanResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return false, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-attribute
func (api WebDriverApi) GetElementAttribute(elementId string, name string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/attribute/%s", api.SessionId, elementId, name))
	if err != nil {
		log.Error("An error occured during get element attribute request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element attribute: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get element attribute: elt: %s / name: %s", elementId, name)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-property
func (api WebDriverApi) GetElementProperty(elementId string, name string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/property/%s", api.SessionId, elementId, name))
	if err != nil {
		log.Error("An error occured during get element property request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element property: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get element property: elt: %s / property: %s", elementId, name)
	}

	// Manage element response
	responseElement := StringResponse{}
	err = mapstructure.Decode(resp, &responseElement)
	if err == nil && responseElement.Value != "" {
		return responseElement.Value, nil
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-css-value
func (api WebDriverApi) GetElementCSSValue(elementId string, name string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/css/%s", api.SessionId, elementId, name))
	if err != nil {
		log.Error("An error occured during get element CSS value request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element CSS value: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get element css value: elt: %s / name: %s", elementId, name)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-text
func (api WebDriverApi) GetElementText(elementId string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/text", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element text request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element text: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get element text: elt: %s", elementId)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-tag-name
func (api WebDriverApi) GetElementTagName(elementId string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/name", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element tag name request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element tag name: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get element tag name: elt: %s", elementId)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-element-rect
func (api WebDriverApi) GetElementRect(elementId string) (common.Rect, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/rect", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element rect request: ", err)
		return common.Rect{}, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get element rect: ", responseError.Value.Message)
		return common.Rect{}, fmt.Errorf("impossible to get element rect: elt: %s", elementId)
	}

	// Manage response
	responseBody := RectResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return common.Rect{}, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#is-element-enabled
func (api WebDriverApi) IsElementEnabled(elementId string) (bool, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/enabled", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during is element enabled request: ", err)
		return false, err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to determine if element selected: ", responseError.Value.Message)
		return false, fmt.Errorf("impossible to determine if element selected")
	}

	// Manage response
	responseBody := BooleanResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return false, err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-computed-role
func (api WebDriverApi) GetComputedRole(elementId string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/computedrole", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get computed role request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get computed role: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get computed role: elt: %s", elementId)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
		return "", err
	}

	return responseBody.Value, nil
}

// https://w3c.github.io/webdriver/#get-computed-label
func (api WebDriverApi) GetComputedLabel(elementId string) (string, error) {

	// Send request
	resp, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/computedlabel", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get computed label request: ", err)
		return "", err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to get computed label: ", responseError.Value.Message)
		return "", fmt.Errorf("impossible to get computed label: elt: %s", elementId)
	}

	// Manage response
	responseBody := StringResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding: ", err)
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
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/click", api.SessionId, elementId), nil)
	if err != nil {
		log.Error("An error occured during click request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to click: ", responseError.Value.Message)
		return fmt.Errorf("impossible to click: elt: %s", elementId)
	}

	return nil
}

// https://w3c.github.io/webdriver/#element-clear
func (api WebDriverApi) Clear(elementId string) error {

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/clear", api.SessionId, elementId), nil)
	if err != nil {
		log.Error("An error occured during clear request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to clear: ", responseError.Value.Message)
		return fmt.Errorf("impossible to clear: elt: %s", elementId)
	}

	return nil
}

// https://w3c.github.io/webdriver/#element-send-keys
func (api WebDriverApi) SendKeys(elementId string, text string) error {

	// Create request body
	request := SendKeysRequest{
		Text: text,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/value", api.SessionId, elementId), request)
	if err != nil {
		log.Error("An error occured during click request: ", err)
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		log.Error("Impossible to send keys: ", responseError.Value.Message)
		return fmt.Errorf("impossible to send keys: elt: %s", elementId)
	}

	return nil
}
