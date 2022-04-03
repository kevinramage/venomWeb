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

// Click method allow you to simulate click on element
// Return nil if operation complete, return an error else
// "invalid session id" error occured when session not found
// "invalid element id" error occured when element not found
func (elt Element) Click() error {
	return elt.api.Click(elt.elementId)
}

func (elt Element) SendKeys(text string) error {
	return elt.api.SendKeys(elt.elementId, text)
}

func (elt Element) Clear() error {
	return elt.api.Clear(elt.elementId)
}

func (elt Element) FindElement(selector string, selectorStrategy string) (Element, error) {
	eltId, err := elt.api.FindElementFromElement(elt.elementId, selector, selectorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: elt.api}, nil
	} else {
		return Element{}, nil
	}
}

func (elt Element) FindElements(selector string, selectorStrategy string) ([]Element, error) {
	eltIds, err := elt.api.FindElementsFromElement(elt.elementId, selector, selectorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltIds); i++ {
			elements = append(elements, Element{elementId: eltIds[i]})
		}
		return elements, nil
	} else {
		return []Element{}, nil
	}
}

func (elt Element) GetElementShadowRoot() (Element, error) {
	eltId, err := elt.api.GetElementShadowRoot(elt.elementId)
	if err == nil {
		return Element{elementId: eltId, api: elt.api}, nil
	} else {
		return Element{}, nil
	}
}

func (elt Element) IsElementSelected() (bool, error) {
	return elt.api.IsElementSelected(elt.elementId)
}

func (elt Element) GetElementProperty(propertyName string) (string, error) {
	return elt.api.GetElementProperty(elt.elementId, propertyName)
}

func (elt Element) GetElementCSSValue(propertyName string) (string, error) {
	return elt.api.GetElementCSSValue(elt.elementId, propertyName)
}

func (elt Element) GetElementText() (string, error) {
	return elt.api.GetElementText(elt.elementId)
}

func (elt Element) GetElementTagName() (string, error) {
	return elt.api.GetElementTagName(elt.elementId)
}

func (elt Element) GetElementRect() (common.Rect, error) {
	return elt.api.GetElementRect(elt.elementId)
}

func (elt Element) IsElementEnabled() (bool, error) {
	return elt.api.IsElementEnabled(elt.elementId)
}

func (elt Element) GetComputedRole() (string, error) {
	return elt.api.GetComputedRole(elt.elementId)
}

func (elt Element) GetComputedLabel() (string, error) {
	return elt.api.GetComputedLabel(elt.elementId)
}

func (elt Element) TakeScreenshot(fileName string) error {
	content, err := elt.api.TakeElementScreenShot(elt.elementId)
	if err == nil {
		return ioutil.WriteFile(fileName, content, 0666)
	} else {
		return err
	}
}

func (elt Element) Select(text string) (Element, error) {
	expression := fmt.Sprintf(`./option[normalize-space()="%s"]`, text)
	newElt, err := elt.FindElement(expression, common.XPATH_SELECTOR)
	if err == nil {
		return Element{elementId: newElt.elementId, api: elt.api}, err
	} else {
		return Element{}, nil
	}
}

func (elt Element) UploadFile(file string) error {
	return elt.SendKeys(file)
}

func (elt Element) Check() error {
	// GetAttribute("type") => checbox
	//elt.api.GetElementAttribute("value")
	return elt.Click()
}

func (elt Element) Uncheck() error {
	return elt.Click()
}
