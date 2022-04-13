package venomWeb

import "github.com/kevinramage/venomWeb/api"

type Shadow struct {
	api      api.WebDriverApi
	shadowId string
}

func (s Shadow) FindElementFromShadowId(selector string, locatorStrategy string) (Element, error) {
	eltId, err := s.api.FindElementFromShadow(s.shadowId, selector, locatorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: s.api}, nil
	} else {
		return Element{}, err
	}
}

func (s Shadow) FindElementsFromShadow(shadowId string, selector string, locatorStrategy string) ([]Element, error) {
	eltsId, err := s.api.FindElementsFromShadow(s.shadowId, selector, locatorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltsId); i++ {
			elements = append(elements, Element{elementId: eltsId[i], api: s.api})
		}
		return elements, nil
	} else {
		return []Element{}, err
	}
}
