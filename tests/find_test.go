package venomTest

import (
	"os"
	"testing"

	venomWeb "github.com/kevinramage/venomWeb/wrapper"
)

func TestFindCSSSelector(t *testing.T) {
	os.Setenv("GO_TEST", "true")
	prefs := make(map[string]interface{})
	webDriver := venomWeb.ChromeDriver("", "", []string{"headless"}, prefs, "")
	err := webDriver.Start()
	if err != nil {
		t.Fatalf("Impossible to start chrome driver: %s", err)
	}
	page, err := webDriver.NewSession()
	if err != nil {
		t.Fatalf("Impossible to create a new page: %s", err)
	}
	err = page.Navigate("https://github.com/")
	if err != nil {
		t.Fatalf("Impossible to navigate: %s", err)
	}
	//elt, _ := page.FindElement(".flex-lg-items-center a[href='/team']", api.CSS_SELECTOR)
	err = webDriver.Stop()
	if err != nil {
		t.Fatalf("Impossible to stop chrome driver: %s", err)
	}
}
