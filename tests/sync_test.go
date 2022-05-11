package venomTest

import (
	"os"
	"testing"

	"github.com/kevinramage/venomWeb/common"
	venomWeb "github.com/kevinramage/venomWeb/wrapper"
)

func TestSyncTestNominal(t *testing.T) {
	os.Setenv("GO_TEST", "true")
	prefs := make(map[string]interface{})
	webDriver := venomWeb.ChromeDriver([]string{"headless"}, prefs)
	err := webDriver.Start()
	if err != nil {
		t.Fatalf("Impossible to start chrome driver: %s", err)
	}
	page, err := webDriver.NewSession()
	if err != nil {
		t.Fatalf("Impossible to create a new page: %s", err)
	}

	// Navigate
	err = page.Navigate("https://www.w3schools.com/w3css/w3css_progressbar.asp")
	if err != nil {
		t.Fatalf("Impossible to navigate: %s", err)
	}

	// Accept cookie
	btnAcceptCookie, err := page.FindElement("#accept-choices", common.CSS_SELECTOR)
	if err != nil {
		t.Fatalf("Impossible to find accept cookie button: %s", err)
	}
	err = btnAcceptCookie.Click()
	if err != nil {
		t.Fatalf("Impossible to click on accept cookie button: %s", err)
	}

	// Click on progress bar button
	startProgressBarBtn, err := page.FindElement(".w3-button.w3-green", common.CSS_SELECTOR)
	if err != nil {
		t.Fatalf("Impossible to find start progress bar button: %s", err)
	}
	err = startProgressBarBtn.Click()
	if err != nil {
		t.Fatalf("Impossible to click on start progress bar button: %s", err)
	}

	// Sync
	err = page.SyncElementText("#demo", common.CSS_SELECTOR, 10000, "100%")
	if err != nil {
		t.Fatalf("Impossible to synchronize: %s", err)
	}

	// Stop
	err = webDriver.Stop()
	if err != nil {
		t.Fatalf("Impossible to stop chrome driver: %s", err)
	}
}
