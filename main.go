package main

import (
	"github.com/kevinramage/venomWeb/common"
	venomWeb "github.com/kevinramage/venomWeb/wrapper"
)

func main() {

	prefs := make(map[string]interface{})
	webDriver := venomWeb.OperaDriver("", "", []string{}, prefs, "")
	webDriver.LogLevel = common.DEBUG
	webDriver.Detach = true

	webDriver.Start()
	page, _ := webDriver.NewSession()
	page.Navigate("https://github.com/")
	webDriver.Stop()
}
