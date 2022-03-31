package main

import (
	"fmt"
	"log"

	"github.com/kevinramage/venomWeb/common"
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
	//	webDriver := venomWeb.ChromeDriver([]string{"headless", "ignore-certificate-errors", "ignore-ssl-errors", "proxy-server=localhost:8888"})
	webDriver := venomWeb.GeckoDriver([]string{})
	webDriver.Headless = true
	//webDriver.Detach = true
	//webDriver.Proxy = "localhost:8888"
	webDriver.LogLevel = "DEBUG"

	//webDriver.Start()
	err := webDriver.Start()
	if err == nil {
		//webDriver.NewPage()
		//webDriver.Status()
		page, _ := webDriver.NewSession()
		page.Navigate("https://github.com/")

		_, err := page.FindElement("input[name=q]", common.CSS_SELECTOR)
		//fmt.Printf("err: %v\n", err)
		//elts, _ := page.FindElements("inputbchzhf", common.CSS_SELECTOR)
		//fmt.Printf("elts: %v\n", elts)

		//elt, _ := page.FindElement("div .HeaderMenu.HeaderMenu", common.CSS_SELECTOR)
		//subElt, err := page.Api.FindElementFromElement(elt, "input[name=qe]", common.CSS_SELECTOR)
		//page.FindElement("div .HeaderMenu.HeaderMenu", common.CSS_SELECTOR)
		//subElts, err := page.Api.FindElementsFromElement("", "input", common.CSS_SELECTOR)

		//elt, _ := page.Api.GetActiveElement()
		//check, err := page.Api.IsElementSelected(elt)
		//check, err := page.Api.IsElementEnabled(elt)
		//fmt.Printf("check: %v\n", check)

		/*
			att, _ := page.Api.GetElementAttribute(elt, "class")
			fmt.Printf("att: %v\n", att)

			lbl, _ := page.Api.GetComputedLabel(elt)
			fmt.Printf("lbl: %v\n", lbl)

			role, _ := page.Api.GetComputedRole(elt)
			fmt.Printf("role: %v\n", role)

			css, _ := page.Api.GetElementCSSValue(elt, "background-color")
			fmt.Printf("css: %v\n", css)

			tagName, _ := page.Api.GetElementTagName(elt)
			fmt.Printf("tagName: %v\n", tagName)

			txt, _ := page.Api.GetElementText(elt)
			fmt.Printf("txt: %v\n", txt)
		*/

		//property, _ := page.Api.GetElementProperty(elt, "parentNode")
		//property, _ := page.Api.GetElementProperty(elt, "ariaAutoComplete")
		//fmt.Printf("property: %v\n", property)

		//rect, err := page.Api.GetElementRect(elt)
		// fmt.Printf("rect: %v\n", rect)

		// fmt.Printf("subElt: %v\n", elt)
		//test, _ := page.Api.TakeScreenShot()
		//ioutil.WriteFile("test.png", test, 0666)
		fmt.Printf("err: %v\n", err)

		//webDriver.Stop()
	} else {
		log.Fatal("Impossible to start web driver")
	}
}
