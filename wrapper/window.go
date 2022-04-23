package venomWeb

import (
	"github.com/kevinramage/venomWeb/api"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Window struct {
	api      api.WebDriverApi
	handleId string
}

func (w Window) SwitchWindow() error {
	log.Info("Window.SwitchWindow")
	err := w.api.SwitchWindow(w.handleId)
	if err != nil {
		err = errors.Wrapf(err, "an error occured during switch window action")
		log.Error(err)
	}
	return err
}
