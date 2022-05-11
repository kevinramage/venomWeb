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
	resp, errResp := ProceedPostRequest(api, fmt.Sprintf("session/%s/actions", api.SessionId), request)

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}

// https://w3c.github.io/webdriver/#release-actions
func (api WebDriverApi) ReleaseActions() error {

	// Security
	if api.SessionId == "" {
		return fmt.Errorf("invalid session id")
	}

	// Send request
	resp, errResp := ProceedDeleteRequest(api, fmt.Sprintf("session/%s/actions", api.SessionId))

	// Manage functionnal error
	responseError := ElementErrorResponse{}
	if resp != nil {
		err := mapstructure.Decode(resp, &responseError)
		if err == nil && responseError.Value.Error != "" {
			return fmt.Errorf(responseError.Value.Error)
		}
	}

	// Manage technical error
	return errResp
}
