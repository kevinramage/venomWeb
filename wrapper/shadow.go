package venomWeb

import "github.com/kevinramage/venomWeb/api"

type Shadow struct {
	Api      api.WebDriverApi
	ShadowId string
}

func (s Shadow) FindElementFromShadowId(selector string, locatorStrategy string) (Element, error) {
	eltId, err := s.Api.FindElementFromShadow(s.ShadowId, selector, locatorStrategy)
	if err == nil {
		return Element{ElementId: eltId}, nil
	} else {
		return Element{}, nil
	}
}

func (s Shadow) FindElementsFromShadow(shadowId string, selector string, locatorStrategy string) ([]Element, error) {
	eltsId, err := s.Api.FindElementsFromShadow(s.ShadowId, selector, locatorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltsId); i++ {
			elements = append(elements, Element{ElementId: eltsId[i]})
		}
		return elements, nil
	} else {
		return []Element{}, nil
	}
}
