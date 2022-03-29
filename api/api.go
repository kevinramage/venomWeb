package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ProceedGetRequest(api WebDriverApi, path string) (interface{}, error) {
	log.Debug(fmt.Sprintf("Api.GetResponse: %s", path))
	resp, err := api.Client.Get(fmt.Sprintf("%s/%s", api.Url, path))
	if err != nil {
		log.Error("Error during GET call: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if !strings.HasSuffix(path, "/screenshot") {
		log.Debug(fmt.Sprintf("Response body: %s", string(body)))
	}
	if err != nil {
		log.Error("Error during body reading: ", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Error("Error invalid response status: ", resp.StatusCode)
		return nil, err
	}
	var bodyJSON interface{}
	errJSON := json.Unmarshal(body, &bodyJSON)
	if errJSON != nil {
		log.Error("Error during json parsing: ", errJSON, " JSON: ", string(body))
		return nil, err
	}
	return bodyJSON, nil
}

func ProceedPostRequest(api WebDriverApi, path string, requestBody interface{}) (interface{}, error) {
	log.Debug(fmt.Sprintf("Api.PostResponse: %s", path))
	var reqBodyJSON []byte
	reqBodyJSON, err := json.Marshal(requestBody)
	log.Debug(fmt.Sprintf("Request body: %s", string(reqBodyJSON)))
	if err != nil {
		log.Error("Error during json parsing: ", err)
	}
	if string(reqBodyJSON) == "null" {
		reqBodyJSON = []byte("{}")
	}
	reader := bytes.NewReader(reqBodyJSON)
	resp, err := api.Client.Post(fmt.Sprintf("%s/%s", api.Url, path), "application/json", reader)
	if err != nil {
		log.Error("Error during GET call: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Debug(fmt.Sprintf("Response body: %s", string(body)))
	if err != nil {
		log.Error("Error during body reading: ", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Error("Error invalid response status: ", resp.StatusCode)
		return nil, err
	}
	var bodyJSON interface{}
	errJSON := json.Unmarshal(body, &bodyJSON)
	if errJSON != nil {
		log.Error("Error during json parsing: ", errJSON)
		return nil, err
	}
	return bodyJSON, nil
}

func ProceedDeleteRequest(api WebDriverApi, path string) (interface{}, error) {
	log.Debug(fmt.Sprintf("Api.ProceedDeleteRequest: %s", path))
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", api.Url, path), nil)
	if err != nil {
		log.Error("Error during request creation: ", err)
		return nil, err
	}
	resp, err := api.Client.Do(request)
	if err != nil {
		log.Error("Error during DELETE call: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Debug(fmt.Sprintf("Response body: %s", string(body)))
	if err != nil {
		log.Error("Error during body reading: ", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Error("Error invalid response status: ", resp.StatusCode)
		return nil, err
	}
	var bodyJSON interface{}
	errJSON := json.Unmarshal(body, &bodyJSON)
	if errJSON != nil {
		log.Error("Error during json parsing: ", errJSON, " JSON: ", string(body))
		return nil, err
	}
	return bodyJSON, nil
}
