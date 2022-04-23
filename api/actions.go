package api

import (
	"fmt"

	"github.com/kevinramage/venomWeb/common"
	"github.com/mitchellh/mapstructure"
)

type PerformActionsRequest struct {
	Actions []common.Action `json:"actions,omitempty"`
}

// https://w3c.github.io/webdriver/#perform-actions
func (api WebDriverApi) PerformActions(actions []common.Action) error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	request := PerformActionsRequest{
		Actions: actions,
	}

	// Send request
	resp, err := ProceedPostRequest(api, fmt.Sprintf("session/%s/actions", api.SessionId), request)
	if err != nil {
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return fmt.Errorf(responseError.Value.Message)
	}

	return nil
}

// https://w3c.github.io/webdriver/#release-actions
func (api WebDriverApi) ReleaseActions() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, err := ProceedDeleteRequest(api, fmt.Sprintf("session/%s/actions", api.SessionId))
	if err != nil {
		return err
	}

	// Manage error
	responseError := ElementErrorResponse{}
	err = mapstructure.Decode(resp, &responseError)
	if err == nil && responseError.Value.Message != "" {
		return fmt.Errorf(responseError.Value.Message)
	}

	return nil
}
