package venomTest

import (
	"fmt"
	"os"
	"testing"

	venomWeb "github.com/kevinramage/venomWeb/wrapper"
)

func TestResizeWindow(t *testing.T) {
	fmt.Printf("Resize\n")
	os.Setenv("GO_TEST", "true")
	webDriver := venomWeb.ChromeDriver([]string{"headless"})
	err := webDriver.Start()
	if err != nil {
		t.Fatalf("Impossible to start chrome driver: %s", err)
	}
	page, err := webDriver.NewPage()
	if err != nil {
		t.Fatalf("Impossible to create a new page: %s", err)
	}
	err = page.Navigate("https://github.com/")
	if err != nil {
		t.Fatalf("Impossible to navigate: %s", err)
	}

	err = page.Size(1024, 768)
	if err != nil {
		t.Fatalf("Impossible to resize: %s", err)
	}

	err = page.Maximize()
	if err != nil {
		t.Fatalf("Impossible to maximize: %s", err)
	}

	err = page.Minimize()
	if err != nil {
		t.Fatalf("Impossible to minimize: %s", err)
	}

	err = page.Fullscreen()
	if err != nil {
		t.Fatalf("Impossible to change to fullscreen mode: %s", err)
	}

	err = webDriver.Stop()
	if err != nil {
		t.Fatalf("Impossible to stop chrome driver: %s", err)
	}
}
