package api

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (api WebDriverApi) Screenshot() error {

	// Send request
	_, err := ProceedGetRequest(api, fmt.Sprintf("session/%s/screenshot", api.SessionId))
	if err != nil {
		log.Error("An error occured during fullscreen request: ", err)
		return err
	} else {
		return nil
	}
}
