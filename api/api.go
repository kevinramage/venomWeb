package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func ProceedGetRequest(api WebDriverApi, path string) (interface{}, error) {
	log.Debug(fmt.Sprintf("Api.GetResponse: %s", path))
	resp, err := api.Client.Get(fmt.Sprintf("%s/%s", api.Url, path))
	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during get request: %s", path)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Limit the body size logged
	bodyText := string(body)
	if len(bodyText) > 1000 {
		log.Debug(fmt.Sprintf("Response body: %s...", bodyText[:1000]))
	} else {
		log.Debug(fmt.Sprintf("Response body: %s", bodyText))
	}

	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during get request: %s", path)
	}
	if resp.StatusCode == 400 {
		return nil, errors.Wrapf(err, "an error occured during get request: %s - Invalid request", path)
	}
	if resp.StatusCode == 404 {
		return nil, errors.Wrapf(err, "an error occured during get request: %s - Resource not found", path)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Wrapf(err, "an error occured during get request: %s - Invalid status code %d", path, resp.StatusCode)
	}
	var bodyJSON interface{}
	errJSON := json.Unmarshal(body, &bodyJSON)
	if errJSON != nil {
		return nil, errors.Wrapf(err, "an error occured during get request: %s", path)
	}
	return bodyJSON, nil
}

func ProceedPostRequest(api WebDriverApi, path string, requestBody interface{}) (interface{}, error) {
	log.Debug(fmt.Sprintf("Api.PostResponse: %s", path))
	var reqBodyJSON []byte
	reqBodyJSON, err := json.Marshal(requestBody)
	log.Debug(fmt.Sprintf("Request body: %s", string(reqBodyJSON)))
	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during post request: %s", path)
	}
	if string(reqBodyJSON) == "null" {
		reqBodyJSON = []byte("{}")
	}
	reader := bytes.NewReader(reqBodyJSON)
	resp, err := api.Client.Post(fmt.Sprintf("%s/%s", api.Url, path), "application/json", reader)
	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during post request: %s", path)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Limit the body size logged
	bodyText := string(body)
	if len(bodyText) > 1000 {
		log.Debug(fmt.Sprintf("Response body: %s...", bodyText[:1000]))
	} else {
		log.Debug(fmt.Sprintf("Response body: %s", bodyText))
	}

	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during post request: %s", path)
	}
	if resp.StatusCode == 400 {
		return nil, errors.Errorf("an error occured during post request: %s - Invalid request", path)
	}
	if resp.StatusCode == 404 {
		return nil, errors.Errorf("an error occured during post request: %s - Resource not found", path)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Errorf("an error occured during post request: %s - Invalid status code %d", path, resp.StatusCode)
	}
	var bodyJSON interface{}
	errJSON := json.Unmarshal(body, &bodyJSON)
	if errJSON != nil {
		return nil, errors.Wrapf(err, "an error occured during post request: %s", path)
	}
	return bodyJSON, nil
}

func ProceedDeleteRequest(api WebDriverApi, path string) (interface{}, error) {
	log.Debug(fmt.Sprintf("Api.ProceedDeleteRequest: %s", path))
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", api.Url, path), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s", path)
	}
	resp, err := api.Client.Do(request)
	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s", path)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// Limit the body size logged
	bodyText := string(body)
	if len(bodyText) > 1000 {
		log.Debug(fmt.Sprintf("Response body: %s...", bodyText[:1000]))
	} else {
		log.Debug(fmt.Sprintf("Response body: %s", bodyText))
	}

	if err != nil {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s", path)
	}
	if resp.StatusCode == 400 {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s - Invalid request", path)
	}
	if resp.StatusCode == 404 {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s - Resource not found", path)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s - Invalid status code %d", path, resp.StatusCode)
	}

	var bodyJSON interface{}
	errJSON := json.Unmarshal(body, &bodyJSON)
	if errJSON != nil {
		return nil, errors.Wrapf(err, "an error occured during delete request: %s", path)
	}

	return bodyJSON, nil
}
