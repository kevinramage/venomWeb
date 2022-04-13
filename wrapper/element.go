// Package venomWeb provide class to interact easily with web driver
package venomWeb

import (
	"fmt"
	"io/ioutil"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
)

type Element struct {
	elementId string
	api       api.WebDriverApi
}

// Click method allow to click on web element
// Return nil if operation proceed with success, return an error else
// "invalid session id" error occured when session not found
// "invalid element id" error occured when element not found
func (elt Element) Click() error {
	return elt.api.Click(elt.elementId)
}

func (elt Element) DoubleClick() error {

	rect, err := elt.api.GetElementRect(elt.elementId)
	if err == nil {
		x := rect.X + rect.Width/2
		y := rect.Y + rect.Height/2 + 70
		actions := createDoubleClickActions(x, y)
		return elt.api.PerformActions(actions)

	} else {
		return err
	}
}

func createDoubleClickActions(x int, y int) []common.Action {
	var subActions = []common.SubAction{}
	subActions = append(subActions, common.SubAction{
		Type:     "pointerMove",
		Duration: 0,
		X:        x,
		Y:        y,
	})
	subActions = append(subActions, common.SubAction{
		Type:   "pointerDown",
		Button: 0,
	})
	subActions = append(subActions, common.SubAction{
		Type:   "pointerUp",
		Button: 0,
	})
	subActions = append(subActions, common.SubAction{
		Type:     "pause",
		Duration: 100,
	})
	subActions = append(subActions, common.SubAction{
		Type:   "pointerDown",
		Button: 0,
	})
	subActions = append(subActions, common.SubAction{
		Type:   "pointerUp",
		Button: 0,
	})
	action := common.Action{
		Id:      "test",
		Type:    "pointer",
		Actions: subActions,
	}
	action.Parameters.PointerType = "mouse"

	var actions = []common.Action{}
	actions = append(actions, action)
	return actions
}

// SendKeys method allow to send text value on web element
// Return nil if operation proceed with success, return an error else
func (elt Element) SendKeys(text string) error {
	return elt.api.SendKeys(elt.elementId, text)
}

// Clear method allow to clear text value of a web element
// Return nil if operation proceed with success, return an error else
func (elt Element) Clear() error {
	return elt.api.Clear(elt.elementId)
}

// FindElement method allow to search a web element
// Return the element if operation proceed with success, return an error else
func (elt Element) FindElement(selector string, selectorStrategy string) (Element, error) {
	eltId, err := elt.api.FindElementFromElement(elt.elementId, selector, selectorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: elt.api}, nil
	} else {
		return Element{}, err
	}
}

// FindElements method allow to search web elements
// Return element list if operation proceed with success, return an error else
func (elt Element) FindElements(selector string, selectorStrategy string) ([]Element, error) {
	eltIds, err := elt.api.FindElementsFromElement(elt.elementId, selector, selectorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltIds); i++ {
			elements = append(elements, Element{elementId: eltIds[i]})
		}
		return elements, nil
	} else {
		return []Element{}, err
	}
}

// GetElementShadowRoot method allow to identify root shadow element
// Return element if operation proceed with success, return an error else
func (elt Element) GetElementShadowRoot() (Element, error) {
	eltId, err := elt.api.GetElementShadowRoot(elt.elementId)
	if err == nil {
		return Element{elementId: eltId, api: elt.api}, nil
	} else {
		return Element{}, err
	}
}

// IsElementSelected method allow to identify if the current element is selected or not
// Return a boolean to indicate if element selected, return an error else
func (elt Element) IsElementSelected() (bool, error) {
	return elt.api.IsElementSelected(elt.elementId)
}

// GetElementProperty method allow to get HTML property of the current element
// Return property value if operation proceed with success, return an error else
func (elt Element) GetElementProperty(propertyName string) (string, error) {
	return elt.api.GetElementProperty(elt.elementId, propertyName)
}

// GetElementCSSValue method allow to get CSS property of the current element
// Return property value if operation proceed with success, return an error else
func (elt Element) GetElementCSSValue(propertyName string) (string, error) {
	return elt.api.GetElementCSSValue(elt.elementId, propertyName)
}

// GetElementText method allow to get text of the current element
// Return text if operation proceed with success, return an error else
func (elt Element) GetElementText() (string, error) {
	return elt.api.GetElementText(elt.elementId)
}

// GetElementTagName method allow to get tag name of the current element
// Return tag name if operation proceed with success, return an error else
func (elt Element) GetElementTagName() (string, error) {
	return elt.api.GetElementTagName(elt.elementId)
}

// GetElementRect method allow to get element position
// Return element position if operation proceed with success, return an error else
func (elt Element) GetElementRect() (common.Rect, error) {
	return elt.api.GetElementRect(elt.elementId)
}

// IsElementEnabled method allow to identify if element enabled or not
// Return if element is enabled if operation proceed with success, return an error else
func (elt Element) IsElementEnabled() (bool, error) {
	return elt.api.IsElementEnabled(elt.elementId)
}

// GetComputedRole method allow to get computed role of the current element
// Return computed role if operation proceed with success, return an error else
func (elt Element) GetComputedRole() (string, error) {
	return elt.api.GetComputedRole(elt.elementId)
}

// GetComputedLabel method allow to get computed label of the current element
// Return computed label if operation proceed with success, return an error else
func (elt Element) GetComputedLabel() (string, error) {
	return elt.api.GetComputedLabel(elt.elementId)
}

// TakeScreenshot method allow to save the screenshot of the current element
// Return nil if operation proceed with success, return an error else
func (elt Element) TakeScreenshot(fileName string) error {
	content, err := elt.api.TakeElementScreenShot(elt.elementId)
	if err == nil {
		return ioutil.WriteFile(fileName, content, 0666)
	} else {
		return err
	}
}

// Select method allow to select an option of a list of values element (<select>)
// Return otpion element if operation proceed with success, return an error else
func (elt Element) Select(text string) (Element, error) {
	expression := fmt.Sprintf(`./option[normalize-space()="%s"]`, text)
	newElt, err := elt.FindElement(expression, common.XPATH_SELECTOR)
	if err == nil {
		return Element{elementId: newElt.elementId, api: elt.api}, nil
	} else {
		return Element{}, err
	}
}

// UploadFile method allow to upload file on an upload web element
// Return nil if operation proceed with success, return an error else
func (elt Element) UploadFile(file string) error {
	return elt.SendKeys(file)
}

// Check method allow to check a checkbox web element
// Return nil if operation proceed with success, return an error else
func (elt Element) Check() error {
	// GetAttribute("type") => checbox
	//elt.api.GetElementAttribute("value")
	return elt.Click()
}

// Uncheck method allow to uncheck a checkbox web element
// Return nil if operation proceed with success, return an error else
func (elt Element) Uncheck() error {
	return elt.Click()
}
