// Package venomWeb provide class to interact easily with web driver
package venomWeb

import (
	"fmt"
	"io/ioutil"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Element struct {
	elementId string
	api       api.WebDriverApi
}

// Click method allow to click on a web element
// Return nil if operation proceed with success, return an error else
// "invalid session id" error occured when session not found
// "invalid element id" error occured when element not found
func (elt Element) Click() error {
	log.Info("Element.Click")

	err := elt.api.Click(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during click action")
		log.Error(err)
	}
	return err
}

// Double click method allow to proceed a double click action on a web element
// Return nil if operation proceed with success, return an error else
func (elt Element) DoubleClick() error {
	log.Info("Element.DoubleClick")
	rect, err := elt.api.GetElementRect(elt.elementId)
	if err == nil {
		x := rect.X + rect.Width/2
		y := rect.Y + rect.Height/2 + 70
		actions := createDoubleClickActions(x, y)
		err = elt.api.PerformActions(actions)
		if err != nil {
			err = errors.Wrapf(err, "an error occured during double click action")
		}

	} else {
		err = errors.Wrapf(err, "an error occured during double click action")
	}

	if err != nil {
		log.Error(err)
	}

	return err
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
	log.Info("Element.SendKeys")
	err := elt.api.SendKeys(elt.elementId, text)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during send keys action")
		log.Error(err)
	}
	return err
}

// Clear method allow to clear text value of a web element
// Return nil if operation proceed with success, return an error else
func (elt Element) Clear() error {
	log.Info("Element.Clear")
	err := elt.api.Clear(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during clear action")
		log.Error(err)
	}
	return err
}

// FindElement method allow to search a web element from a parent element
// Return the element if operation proceed with success, return an error else
func (elt Element) FindElement(selector string, selectorStrategy string) (Element, error) {
	log.Info("Element.FindElement")
	eltId, err := elt.api.FindElementFromElement(elt.elementId, selector, selectorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: elt.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during find element action")
		log.Error(err)
		return Element{}, err
	}
}

// FindElements method allow to search web elements from a parent element
// Return element list if operation proceed with success, return an error else
func (elt Element) FindElements(selector string, selectorStrategy string) ([]Element, error) {
	log.Info("Element.FindElements")
	eltIds, err := elt.api.FindElementsFromElement(elt.elementId, selector, selectorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltIds); i++ {
			elements = append(elements, Element{elementId: eltIds[i]})
		}
		return elements, nil
	} else {
		err = errors.Wrapf(err, "an error occured during find elements action")
		log.Error(err)
		return []Element{}, err
	}
}

// GetElementShadowRoot method allow to identify root shadow element from a parent element
// Return element if operation proceed with success, return an error else
func (elt Element) GetElementShadowRoot() (Element, error) {
	log.Info("Element.GetElementShadowRoot")
	eltId, err := elt.api.GetElementShadowRoot(elt.elementId)
	if err == nil {
		return Element{elementId: eltId, api: elt.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during get element shadow root action")
		log.Error(err)
		return Element{}, err
	}
}

// IsElementSelected method allow to identify if the current element is selected or not
// Return a boolean to indicate if element selected, return an error else
func (elt Element) IsElementSelected() (bool, error) {
	log.Info("Element.IsElementSelected")
	selected, err := elt.api.IsElementSelected(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during is element selected action")
		log.Error(err)
	}
	return selected, err
}

// GetElementProperty method allow to get HTML property of the current element
// Return property value if operation proceed with success, return an error else
func (elt Element) GetElementProperty(propertyName string) (string, error) {
	log.Info("Element.GetElementProperty")
	property, err := elt.api.GetElementProperty(elt.elementId, propertyName)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get element property action")
		log.Error(err)
	}
	return property, err
}

// GetElementCSSValue method allow to get CSS property of the current element
// Return property value if operation proceed with success, return an error else
func (elt Element) GetElementCSSValue(propertyName string) (string, error) {
	log.Info("Element.GetElementCSSValue")
	property, err := elt.api.GetElementCSSValue(elt.elementId, propertyName)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get element CSS action")
		log.Error(err)
	}
	return property, err
}

// GetElementText method allow to get text of the current element
// Return text if operation proceed with success, return an error else
func (elt Element) GetElementText() (string, error) {
	log.Info("Element.GetElementText")
	text, err := elt.api.GetElementText(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get element text action")
		log.Error(err)
	}
	return text, err
}

// GetElementTagName method allow to get tag name of the current element
// Return tag name if operation proceed with success, return an error else
func (elt Element) GetElementTagName() (string, error) {
	log.Info("Element.GetElementTagName")
	tagName, err := elt.api.GetElementTagName(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get element tag name action")
		log.Error(err)
	}
	return tagName, err
}

// GetElementRect method allow to get element position
// Return element position if operation proceed with success, return an error else
func (elt Element) GetElementRect() (common.Rect, error) {
	log.Info("Element.GetElementRect")
	rect, err := elt.api.GetElementRect(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get element rect action")
		log.Error(err)
	}
	return rect, err
}

// IsElementEnabled method allow to identify if element enabled or not
// Return if element is enabled if operation proceed with success, return an error else
func (elt Element) IsElementEnabled() (bool, error) {
	log.Info("Element.IsElementEnabled")
	enabled, err := elt.api.IsElementEnabled(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during is element enabled action")
		log.Error(err)
	}
	return enabled, err
}

// GetComputedRole method allow to get computed role of the current element
// Return computed role if operation proceed with success, return an error else
func (elt Element) GetComputedRole() (string, error) {
	log.Info("Element.GetComputedRole")
	role, err := elt.api.GetComputedRole(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get computed role action")
		log.Error(err)
	}
	return role, err
}

// GetComputedLabel method allow to get computed label of the current element
// Return computed label if operation proceed with success, return an error else
func (elt Element) GetComputedLabel() (string, error) {
	log.Info("Element.GetComputedLabel")
	label, err := elt.api.GetComputedLabel(elt.elementId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during get computed label action")
		log.Error(err)
	}
	return label, err
}

// TakeScreenshot method allow to save the screenshot of the current element
// Return nil if operation proceed with success, return an error else
func (elt Element) TakeScreenshot(fileName string) error {
	log.Info("Element.TakeScreenshot")
	content, err := elt.api.TakeElementScreenShot(elt.elementId)
	if err == nil {
		err = ioutil.WriteFile(fileName, content, 0666)
		if err != nil {
			err = errors.Wrapf(err, "an error occured during take screenshot action")
			log.Error(err)
		}
	} else {
		err = errors.Wrapf(err, "an error occured during take screenshot action")
		log.Error(err)
	}
	return err
}

// Select method allow to select an option of a list of values element (<select>)
// Return otpion element if operation proceed with success, return an error else
func (elt Element) Select(text string) (Element, error) {
	log.Info("Element.Select")
	expression := fmt.Sprintf(`./option[normalize-space()="%s"]`, text)
	newElt, err := elt.FindElement(expression, common.XPATH_SELECTOR)
	if err == nil {
		return Element{elementId: newElt.elementId, api: elt.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during select action")
		log.Error(err)
		return Element{}, err
	}
}

// UploadFile method allow to upload file on an upload web element
// Return nil if operation proceed with success, return an error else
func (elt Element) UploadFile(file string) error {
	log.Info("Element.UploadFile")
	err := elt.SendKeys(file)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during upload file action")
		log.Error(err)
	}
	return err
}

// Check method allow to check a checkbox web element
// Return nil if operation proceed with success, return an error else
func (elt Element) Check() error {
	log.Info("Element.Check")
	// GetAttribute("type") => checbox
	//elt.api.GetElementAttribute("value")
	err := elt.Click()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during check action")
		log.Error(err)
	}
	return err
}

// Uncheck method allow to uncheck a checkbox web element
// Return nil if operation proceed with success, return an error else
func (elt Element) Uncheck() error {
	log.Info("Element.Uncheck")
	err := elt.Click()
	if err != nil {
		err = errors.Wrapf(err, "an error occured during uncheck action")
		log.Error(err)
	}
	return err
}
