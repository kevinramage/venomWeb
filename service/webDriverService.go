package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	log "github.com/sirupsen/logrus"
)

type WebDriverService struct {
	Command *exec.Cmd
}

func New() WebDriverService {
	s := WebDriverService{}
	return s
}

func (s *WebDriverService) Start(command string, logLevel string, args []string) error {
	log.Debug("WebDriverService.Start")
	if os.Getenv("GO_TEST") != "true" {
		s.Command = exec.Command(command)
		s.Command.Args = args
		if logLevel == "DEBUG" {
			s.Command.Stdout = os.Stdout
			s.Command.Stderr = os.Stderr
		}
		err := s.Command.Start()
		if err != nil {
			log.Error("An error occured during web driver starting: ", err)
		}
		return err
	}
	return nil
}

func (s WebDriverService) Wait(timeout time.Duration, url string) error {
	log.Debug("WebDriverService.Wait")
	timeoutChan := time.After(timeout)
	failedChan := make(chan struct{}, 1)
	startedChan := make(chan struct{})

	go func() {
		up := s.CheckStatus(url)
		for !up {
			select {
			case <-failedChan:
				return
			default:
				time.Sleep(500 * time.Millisecond)
				up = s.CheckStatus(url)
			}
		}
		startedChan <- struct{}{}
	}()

	select {
	case <-timeoutChan:
		failedChan <- struct{}{}
		return errors.New("failed to start before timeout")
	case <-startedChan:
		return nil
	}
}

func (s WebDriverService) Stop() error {
	if os.Getenv("GO_TEST") != "true" {
		log.Debug("WebDriverService.Stop")
		if s.Command != nil {
			err := s.Command.Process.Kill()
			return err
		}
	}
	return nil
}

func (s WebDriverService) CheckStatus(url string) bool {
	log.Debug("WebDriverService.CheckStatus")
	client := http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf("%s/status", url), nil)
	response, err := client.Do(request)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == 200
}
