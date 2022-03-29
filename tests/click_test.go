package venomTest

import (
	"os"
	"testing"

	venomWeb "github.com/kevinramage/venomWeb/wrapper"
)

func TestClick(t *testing.T) {
	os.Setenv("GO_TEST", "true")
	webDriver := venomWeb.ChromeDriver([]string{"headless"})
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

	err = webDriver.Stop()
	if err != nil {
		t.Fatalf("Impossible to stop chrome driver: %s", err)
	}
}
