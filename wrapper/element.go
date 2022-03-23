package venomWeb

import "github.com/kevinramage/venomWeb/api"

type Element struct {
	ElementId string
	Api       api.WebDriverApi
}

// https://w3c.github.io/webdriver/#element-click
func (elt Element) Click() error {
	return elt.Api.Click(elt.ElementId)
}

// https://w3c.github.io/webdriver/#element-send-keys
func (elt Element) Fill(text string) error {
	return elt.Api.SendKeys(elt.ElementId, text)
}

// https://w3c.github.io/webdriver/#element-clear
func (elt Element) Clear() error {
	return elt.Api.Clear(elt.ElementId)
}

func (elt Element) Select(text string) {

}

func (elt Element) UploadFile(file string) {

}
