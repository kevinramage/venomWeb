package venomWeb

import (
	"io/ioutil"

	"github.com/kevinramage/venomWeb/api"
	"github.com/kevinramage/venomWeb/common"
)

type Element struct {
	ElementId string
	Api       api.WebDriverApi
}

func (elt Element) Click() error {
	return elt.Api.Click(elt.ElementId)
}

func (elt Element) SendKeys(text string) error {
	return elt.Api.SendKeys(elt.ElementId, text)
}

func (elt Element) Clear() error {
	return elt.Api.Clear(elt.ElementId)
}

func (elt Element) FindElement(selector string, selectorStrategy string) (Element, error) {
	eltId, err := elt.Api.FindElementFromElement(elt.ElementId, selector, selectorStrategy)
	if err == nil {
		return Element{ElementId: eltId}, nil
	} else {
		return Element{}, nil
	}
}

func (elt Element) FindElements(selector string, selectorStrategy string) ([]Element, error) {
	eltIds, err := elt.Api.FindElementsFromElement(elt.ElementId, selector, selectorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltIds); i++ {
			elements = append(elements, Element{ElementId: eltIds[i]})
		}
		return elements, nil
	} else {
		return []Element{}, nil
	}
}

func (elt Element) GetElementShadowRoot() (Element, error) {
	eltId, err := elt.Api.GetElementShadowRoot(elt.ElementId)
	if err == nil {
		return Element{ElementId: eltId}, nil
	} else {
		return Element{}, nil
	}
}

func (elt Element) IsElementSelected() (bool, error) {
	return elt.Api.IsElementSelected(elt.ElementId)
}

func (elt Element) GetElementProperty(propertyName string) (string, error) {
	return elt.Api.GetElementProperty(elt.ElementId, propertyName)
}

func (elt Element) GetElementCSSValue(propertyName string) (string, error) {
	return elt.Api.GetElementCSSValue(elt.ElementId, propertyName)
}

func (elt Element) GetElementText() (string, error) {
	return elt.Api.GetElementText(elt.ElementId)
}

func (elt Element) GetElementTagName() (string, error) {
	return elt.Api.GetElementTagName(elt.ElementId)
}

func (elt Element) GetElementRect() (common.Rect, error) {
	return elt.Api.GetElementRect(elt.ElementId)
}

func (elt Element) IsElementEnabled() (bool, error) {
	return elt.Api.IsElementEnabled(elt.ElementId)
}

func (elt Element) GetComputedRole() (string, error) {
	return elt.Api.GetComputedRole(elt.ElementId)
}

func (elt Element) GetComputedLabel() (string, error) {
	return elt.Api.GetComputedLabel(elt.ElementId)
}

func (elt Element) TakeScreenshot(fileName string) error {
	content, err := elt.Api.TakeElementScreenShot(elt.ElementId)
	if err == nil {
		return ioutil.WriteFile(fileName, content, 0666)
	} else {
		return err
	}
}

func (elt Element) Select(text string) {

}

func (elt Element) UploadFile(file string) {

}
