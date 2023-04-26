package main

import (
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
	//braveBinary := "C:\\Program Files\\BraveSoftware\\Brave-Browser\\Application\\brave.exe"
	prefs := make(map[string]interface{})
	webDriver := venomWeb.BraveDriver("", "", []string{}, prefs, "")
	webDriver.LogLevel = common.INFO
	webDriver.Detach = true

	webDriver.Start()
	page, _ := webDriver.NewSession()
	page.Navigate("https://github.com/")
	webDriver.Stop()

	/*
		page, _ := webDriver.NewSession()
		page.Navigate("https://www.w3schools.com/tags/tryit.asp?filename=tryhtml_textarea")
		btn, _ := page.FindElement("#accept-choices", common.CSS_SELECTOR)
		btn.Click()

		//textField, _ := page.FindElement("#textareaCode", common.CSS_SELECTOR)
		//err := textField.SendKeys("coucou")
		script := "window.editor.setValue(\"<!DOCTYPE html> <html> <body>  <h2>JavaScript Timing</h2>  <p>Click \\\"Try it\\\". Wait 3 seconds, and the page will alert \\\"Hello\\\".</p>  <button id='btnFunc' onclick=\\\"setTimeout(myFunction, 3000);\\\">Try it</button> <div id=\\\"test\\\">test</div>  <script> function myFunction() {   document.getElementById(\\\"test\\\").innerHTML = \\\"<button id='btnTest'>Coucou</button>\\\" } </script>  </body> </html>\"); window.editor.save();"
		page.ExecuteScript(script, []string{})

		btnRun, _ := page.FindElement("#runbtn", common.CSS_SELECTOR)
		btnRun.Click()

		frame, _ := page.FindElement("#iframeResult", common.CSS_SELECTOR)
		page.SwitchToFrame(frame)

		btnFunc, _ := page.FindElement("#btnFunc", common.CSS_SELECTOR)
		btnFunc.Click()

		err := page.SyncElement("#btnTest", common.CSS_SELECTOR, 4000)
		fmt.Printf("err: %v\n", err)

		time.Sleep(5 * time.Second)
	*/

	//webDriver.Stop()
	//webDriver.Start()
	//webDriver.Stop()

	/*
		page, _ := webDriver.NewSession()
		page.Navigate("https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_alert")

		page.SyncElement("#accept-choices", common.CSS_SELECTOR, 1000)
		btnAcceptCookie, _ := page.FindElement("#accept-choices", common.CSS_SELECTOR)
		btnAcceptCookie.Click()
		frame, _ := page.FindElement("#iframeResult", common.CSS_SELECTOR)
		page.SwitchToFrame(frame)
		btn, _ := page.FindElement("//button", common.XPATH_SELECTOR)
		btn.Click()
		time.Sleep(2000)
		title, err := page.GetTitle()
		time.Sleep(2000)
		webDriver.Stop()
		fmt.Printf("title: %v\n", title)
		fmt.Printf("err: %v\n", err)
	*/
	/*
		page, _ := webDriver.NewSession()
		page.Navigate("https://www.w3schools.com/w3css/w3css_progressbar.asp")

		eltCookie, _ := page.FindElement("#accept-choices", common.CSS_SELECTOR)
		eltCookie.Click()

		elt, _ := page.FindElement(".w3-button.w3-green", common.CSS_SELECTOR)
		elt.Click()
		page.SyncElementText("#demo", common.CSS_SELECTOR, 10000, "100%")
		//page.SyncElementCSSValue("#myBar", common.CSS_SELECTOR, 20000, "width", "100%")
	*/
	//webDriver.Stop()

	//page.Navigate("https://doubleclicktest.com/")
	//elt, _ := page.FindElement("#textarea", common.CSS_SELECTOR)
	//elt.DoubleClick()
	//time.Sleep(1 * time.Second)
	/*
		//page.Navigate("https://github.com/")
		page.Navigate("https://www.w3schools.com/html/tryit.asp?filename=tryhtml_id_css")

		time.Sleep(1 * time.Second)
		btn, _ := page.FindElement("#accept-choices", common.CSS_SELECTOR)
		btn.Click()

		elt, _ := page.FindElement("#iframeResult", common.CSS_SELECTOR)
		page.SwitchToFrame(elt)
		//page.SwitchToIndexFrame(0)
		page.FindElement("#myHeader", common.CSS_SELECTOR)
	*/
	//webDriver.Stop()

	//webDriver.Headless = true
	//webDriver.Detach = true
	//webDriver.Proxy = "localhost:8888"

	//webDriver.Start()
	// err := webDriver.Start()
	//if err == nil {
	//webDriver.NewPage()
	//webDriver.Status()
	//page, _ := webDriver.NewSession()
	//page.Navigate("https://github.com/")
	//webDriver.Stop()

	//elt, err := page.FindElements("input", common.CSS_SELECTOR)

	//elt, _ := page.FindElement("input[name=q]", common.CSS_SELECTOR)
	//err = elt.SendKeys("venom")

	//time.Sleep(2 * time.Second)

	// #jump-to-suggestion-search-global > a

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
	//fmt.Printf("err: %v\n", err)

	//webDriver.Stop()
	//} else {
	// log.Fatal("Impossible to start web driver")
	//}
}
