package venomWeb

import "github.com/kevinramage/venomWeb/api"

type Window struct {
	Api      api.WebDriverApi
	HandleId string
}

func (w Window) SwitchWindow() error {
	return w.Api.SwitchWindow(w.HandleId)
}
