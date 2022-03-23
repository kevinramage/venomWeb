package api

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

// https://w3c.github.io/webdriver/#dfn-table-of-location-strategies
const (
	CSS_SELECTOR             string = "css selector"
	LINKTEXT_SELECTOR        string = "link text"
	PARTIALLINKTEXT_SELECTOR string = "partial link text"
	TAGNAME_SELECTOR         string = "tag name"
	XPATH_SELECTOR           string = "xpath"
)

// https://chromium.googlesource.com/chromium/src/+/master/docs/chromedriver_status.md
type ElementRequest struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

type ElementResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     struct {
		Element string `json:"ELEMENT"`
	} `json:"value"`
}

type ElementsResponse struct {
	SessionId string `json:"sessionId"`
	Status    int    `json:"status"`
	Value     []struct {
		Element string `json:"ELEMENT"`
	} `json:"value"`
}

type SendKeysRequest struct {
	Value []string `json:"value"`
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

	// Manage response
	responseBody := ElementResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", nil
	}

	return responseBody.Value.Element, nil
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
		log.Error("An error occured during the response decoding")
		return []string{}, nil
	}

	ids := []string{}
	for i := 0; i < len(responseBody.Value); i++ {
		ids = append(ids, responseBody.Value[i].Element)
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

	// Manage response
	responseBody := ElementResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return "", nil
	}

	return responseBody.Value.Element, nil
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

	// Manage response
	responseBody := ElementsResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return []string{}, nil
	}

	ids := []string{}
	for i := 0; i < len(responseBody.Value); i++ {
		ids = append(ids, responseBody.Value[i].Element)
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
		return "", nil
	}

	return responseBody.Value.Element, nil
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

	// Manage response
	responseBody := ElementsResponse{}
	err = mapstructure.Decode(resp, &responseBody)
	if err != nil {
		log.Error("An error occured during the response decoding")
		return []string{}, nil
	}

	ids := []string{}
	for i := 0; i < len(responseBody.Value); i++ {
		ids = append(ids, responseBody.Value[i].Element)
	}

	return ids, nil
}

// https://w3c.github.io/webdriver/#get-active-element
func (api WebDriverApi) GetActiveElement() error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/active", api.SessionId))
	if err != nil {
		log.Error("An error occured during get active element request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-shadow-root
func (api WebDriverApi) GetElementShadowRoot(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/shadow", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element shadow root request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#is-element-selected
func (api WebDriverApi) IsElementSelected(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/selected", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during is element selected request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-attribute
func (api WebDriverApi) GetElementAttribute(elementId string, name string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/attribute/%s", api.SessionId, elementId, name))
	if err != nil {
		log.Error("An error occured during get element attribute request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-property
func (api WebDriverApi) GetElementProperty(elementId string, name string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/property/%s", api.SessionId, elementId, name))
	if err != nil {
		log.Error("An error occured during get element property request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-css-value
func (api WebDriverApi) GetElementCSSValue(elementId string, name string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/css/%s", api.SessionId, elementId, name))
	if err != nil {
		log.Error("An error occured during get element CSS value request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-text
func (api WebDriverApi) GetElementText(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/text", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element text request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-tag-name
func (api WebDriverApi) GetElementTagName(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/name", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element tag name request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-element-rect
func (api WebDriverApi) GetElementRect(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/rect", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get element rect request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#is-element-enabled
func (api WebDriverApi) IsElementEnabled(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/enabled", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during is element enabled request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-computed-role
func (api WebDriverApi) GetComputedRole(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/computedrole", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get computed role request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#get-computed-label
func (api WebDriverApi) GetComputedLabel(elementId string) error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/element/%s/computedlabel", api.SessionId, elementId))
	if err != nil {
		log.Error("An error occured during get computed label request: ", err)
		return err
	}

	return nil
}

// https://w3c.github.io/webdriver/#element-click
func (api WebDriverApi) Click(elementId string) error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/click", api.SessionId, elementId), nil)
	if err != nil {
		log.Error("An error occured during click request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#element-clear
func (api WebDriverApi) Clear(elementId string) error {

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/clear", api.SessionId, elementId), nil)
	if err != nil {
		log.Error("An error occured during clear request: ", err)
		return err
	} else {
		return nil
	}
}

// https://w3c.github.io/webdriver/#element-send-keys
func (api WebDriverApi) SendKeys(elementId string, text string) error {

	// Create request body
	request := SendKeysRequest{
		Value: strings.Split(text, ""),
	}

	// Send request
	_, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/element/%s/value", api.SessionId, elementId), request)
	if err != nil {
		log.Error("An error occured during click request: ", err)
		return err
	} else {
		return nil
	}
}
