package venomWeb

import (
	"github.com/kevinramage/venomWeb/api"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Shadow struct {
	api      api.WebDriverApi
	shadowId string
}

func (s Shadow) FindElementFromShadowId(selector string, locatorStrategy string) (Element, error) {
	log.Info("Shadow.FindElementFromShadowId")
	eltId, err := s.api.FindElementFromShadow(s.shadowId, selector, locatorStrategy)
	if err == nil {
		return Element{elementId: eltId, api: s.api}, nil
	} else {
		err = errors.Wrapf(err, "an error occured during find element from shadow id action")
		log.Error(err)
		return Element{}, err
	}
}

func (s Shadow) FindElementsFromShadow(shadowId string, selector string, locatorStrategy string) ([]Element, error) {
	log.Info("Shadow.FindElementsFromShadow")
	eltsId, err := s.api.FindElementsFromShadow(s.shadowId, selector, locatorStrategy)
	if err == nil {
		elements := []Element{}
		for i := 0; i < len(eltsId); i++ {
			elements = append(elements, Element{elementId: eltsId[i], api: s.api})
		}
		return elements, nil
	} else {
		err = errors.Wrapf(err, "an error occured during find elements from shadow action")
		log.Error(err)
		return []Element{}, err
	}
}
