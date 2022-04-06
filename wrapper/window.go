package venomWeb

import "github.com/kevinramage/venomWeb/api"

type Window struct {
	api      api.WebDriverApi
	handleId string
}

func (w Window) SwitchWindow() error {
	return w.api.SwitchWindow(w.handleId)
}
