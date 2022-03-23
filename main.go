package main

import (
	venomWeb "github.com/kevinramage/venomWeb/wrapper"
)

/*
func main() {
	log.SetLevel(log.DebugLevel)
	webDriverService := service.New()
	webDriverService.Start()
	webDriverService.Wait(time.Second * 60)
	//webDriverService.Client.CheckStatus()
	webDriverService.Client.CreateSession()
	webDriverService.Client.SetUrl("https://github.com/")
	//url, _ := webDriverService.Client.GetUrl()
	//fmt.Printf("Url: %s", url)

	//webDriverService.Client.FindElement(".gLFyf.gsfi", api.CSS_SELECTOR)
	//webDriverService.Client.FindElements(".gLFyf.gsfi", api.CSS_SELECTOR)

	//
	//webDriverService.Client.ClickOnElement(elt)

	//elt, _ := webDriverService.Client.FindElement("input[name='q']", api.CSS_SELECTOR)
	//webDriverService.Client.SendKey(elt, "coucou")
	//time.Sleep(time.Second * 2)

	size, _ := webDriverService.Client.GetSize()
	fmt.Printf("size: %v\n", size)

	time.Sleep(time.Second * 2)

	//webDriverService.Client.RefreshPage()
	//webDriverService.Client.Back()
	//webDriverService.Client.Forward()

	webDriverService.Client.DeleteSession()
	webDriverService.Stop()
}
*/

func main() {
	webDriver := venomWeb.ChromeDriver([]string{"headless"})
	webDriver.Driver.LogLevel = "DEBUG"
	webDriver.Start()
	// page, _ := webDriver.NewPage()
	//page.Navigate("https://github.com/")
	// timeouts, _ := page.GetTimeouts()

	webDriver.Stop()
}
