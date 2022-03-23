package common

import (
	"time"
)

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type WebDriverOptions struct {
	Timeout         time.Duration
	LogLevel        string
	Debug           bool
	Command         string
	Args            []string
	WebDriverBinary string
	Url             string

	//Internal WebDriverInternal
}

/*
type WebDriverInternal struct {
	Command   *exec.Cmd
	Client    http.Client
	SessionId string
	Version   WebDriverVersion
	Ready     bool
}*/

type WebDriverVersion struct {
	DriverVersion string
	Osname        string
	Osarch        string
	Osversion     string
}

type Element struct {
	Id      string
	Text    string
	TagName string
}

// https://w3c.github.io/webdriver/#dfn-timeouts-object
type Timeouts struct {
	Implicit int `json:"implicit"`
	PageLoad int `json:"pageLoad"`
	Script   int `json:"script"`
}

type DriverStatus struct {
	Value struct {
		Build struct {
			Version string `json:"version"`
		} `json:"build"`
		Message string `json:"message"`
		Os      struct {
			Arch    string `json:"arch"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		Ready bool `json:"ready"`
	} `json:"value"`
}
